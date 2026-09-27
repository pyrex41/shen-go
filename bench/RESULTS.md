# shen-go Benchmark Results

Machine: Linux amd64  
Kernel: S39.2  
Date: 2026-05-05  

## Wide fixnums and cheaper number construction (2026-09)

- **Fixnums widened to [-2^35, 2^35)** (`kl/types.go`, `kl/fixnum_span_*.go`).
  On 64-bit unix the fixnum span is a 2^36-byte `PROT_READ` anonymous mapping
  reserved at startup, not the 2^26-byte BSS array. Every uint32, and sums of
  several, is now unboxed. The mapping costs address space (VmSize +64 GB),
  not RSS. Other platforms, or a failed reservation, keep [-2^25, 2^25);
  `SHEN_GO_NARROW_FIXNUM=1` forces that. Fixnum products multiply as `int`
  only when both factors are below 2^31, since two 35-bit values overflow
  int64. Constant folding only folds operands that fit its `int32` fields.
- **`isPreciseInteger` is `f == math.Trunc(f)`**, not `Ilogb` plus a shift.
  `MakeNumber` and `slotFromNumber` range-check first and then narrow.
  `HasCanonicalPrimitiveBinding` drops a redundant `isFixnum` test.

`bench/word32.shen` covers 32-bit word arithmetic of the kind a pure-Shen
SHA-256 or PRNG does, using only `+ -` and comparisons. The figures are
medians of 6 interleaved runs, in seconds:

| | original | after `let`/VM work | now |
|---|---|---|---|
| fib32 ×3M | 0.634 | 0.646 | **0.494 (−22%)** |
| mix 64 words ×40k | 1.756 | 1.809 | **1.335 (−24%)** |
| word list 1M | 0.538 | 0.576 | **0.430 (−20%)** |

Allocation on that benchmark fell from 325 MB to 83 MB (`makeInteger` had been
78% of it). Across all of this work, the kernel cert suite went from 7.31 to
5.06 s wall and from 9.81 to 6.06 s CPU. Startup went from 186 to 153 ms.
Cert output is identical to the original build's.

Measured and not kept:

- **Integer-ID dispatch for `vmGuardedPrimitive`**, in place of its string
  `switch`: no gain on the cert suite, and about 2% slower on `word32`.
- **Running top-level `eval-kl` forms on the VM** instead of the tree-walker.
  A top-level `(foldl (/. A X ...) 0 L)` ran 4–5× faster. It was neutral on
  the cert suite, where the tree-walker's own time is only ~2.5%. It changes
  REPL error text, though, because the VM and the interpreter word the same
  errors differently. An unbound function gives `function F not bound` rather
  than `can't apply non function: F`. `(not 0)` gives `not: expected boolean`
  rather than `if requires a boolean`. Tests pin the interpreter's wording.
  Unify the messages first, then this is a small change.

Found, not changed: the stdlib's `floor`, and so `mod`, `round` and `ceiling`
(`cmd/shen/stlib/Maths/maths.shen`), is a pure-Shen decimal digit-guessing
loop, about 100 µs per call. Word arithmetic written with `mod` is two orders
of magnitude slower than the `+`/compare version above, for that reason
alone. Native versions would need the same equivalence audit as the
`kl/equiv.json` kernel natives.

## Kernel `let` lowering and VM frame/closure allocation (2026-09)

Profiled the kernel certification suite (`kernel/tests/runme.shen`) with
`-cpuprofile`/`-memprofile`. About 30% of its CPU time was GC. Nearly half of
all bytes allocated (1.42 of 2.94 GB) came from `MakeNative` in the compiled
kernel, because `src/compiler.shen` lowered `(let X V Body)` to
`((lambda [X] Body) V)`. Every `let` in the kernel (the type checker and Prolog
engine are full of them) allocated a Go closure plus an `scmNative`, and
entered its body through `TailApply` and the trampoline.

Changes:

- **`let` → `(bind X V)`** (`src/compiler.shen`, `codegen/codegen.go`,
  regenerated `cmd/shen/*.go`). A `let` is now `X := V` followed by the body
  inline. Binders are alpha-renamed to fresh symbols, so shadowing lets and
  lets reusing a parameter name never redeclare a Go variable in one scope.
  `bind` is always a statement and is never scalar-inlined like a `<=`
  temporary. A let variable can be used inside a lambda or an `if` branch, and
  inlining it there would defer or skip an error in V such as `(/ 4 0)`.
- **VM operand stack learns its size** (`kl/vm.go`, `kl/eval.go`). A function
  whose operand stack outgrew the fixed 16-slot frame headroom used to
  reallocate its stack off the arena on every call. It now records the
  capacity it needed, and later frames reserve it.
- **Closures with ≤4 upvalues are one allocation** (`makeBytecodeClosure`),
  not an upvals slice plus a closure object.

Kernel certification suite, 8 interleaved runs, medians. linux/amd64, go1.27.0:

| | wall | CPU (user+sys) | bytes allocated |
|---|---|---|---|
| before | 6.74 s | 9.76 s | 2.94 GB |
| `let` lowering | 4.78 s | 5.96 s | 1.18 GB |
| + VM stack/closure | **4.74 s (−30%)** | **5.84 s (−40%)** | **1.06 GB (−64%)** |

Startup (`shen eval -e '(+ 1 2)'`) went from about 170 ms to about 145 ms.
User code defined at the REPL runs on the bytecode VM, so `bench/bench.shen`
and the `BenchmarkVM*` micro-benchmarks are unchanged within noise
(interleaved geomean −3%). Cert output is byte-identical apart from timings.

## Allocation-reduction work (2026-06)

Profiled with the Go-level VM micro-benchmarks in `kl/vm_bench_test.go`
(`go test ./kl -bench BenchmarkVM -benchmem`). Two findings drove the work:

1. The per-activation VM `locals`/`stack` slices do **not** heap-allocate —
   Go stack-allocates them — so `fib`/`tak`/`ack` already run at **0 allocs/op**.
   "VM slice pooling" was therefore *not* implemented; it would be dead work.
2. The two real allocation sources, by alloc profile, were **integer boxing**
   (`makeInteger`, 99.76% of allocs in integer code) and **the interpreter's
   alist environment** (`cons` via `envExtend`, 99.89% of allocs in lambda code).

Changes:

- **Fixnum range widened** from unsigned `[0, 2^20)` to signed `[-2^25, 2^25)`
  (`kl/types.go`). The sentinel array is pure address space (never dereferenced,
  ~0 RSS). Removes boxing for all negative integers and mid-size positives.
- **Interpreter env node** (`kl/types.go`, `kl/eval.go`): `envExtend` now allocates
  one `scmEnv{sym,val,next}` per binding instead of `cons(cons(sym,val), env)`
  (two cons cells). Halves per-binding allocation for `let`/`lambda` bodies that
  run in the tree-walker.

| Benchmark | Before | After |
|---|---|---|
| `fib(24)` | 0 allocs | 0 allocs (unchanged) |
| `sum` to +32M (in-range) | 7868 allocs, 682 µs | **0 allocs, 398 µs** |
| `sum` to −32M (negatives) | 8000 allocs, 663 µs | **0 allocs, 397 µs** |
| lambda apply ×20000 | 40002 allocs, 5.16 ms | **20002 allocs, 4.36 ms** |

Truly large integers (e.g. an accumulator reaching 5e9) still box — that needs a
full number-representation change, deliberately out of scope here.

## Measurements

### `tak(18,12,6)` single call (wall time, seconds)

| Milestone | Time (s) | vs baseline |
|---|---|---|
| Baseline (tree-walker, interpreter only) | 0.088 | 1× |
| Phase 0 (float-fix only, no perf change) | 0.088 | 1× |
| Phase 2 (bytecode VM, indexed slots) | 0.013 | **6.6×** |
| Phase 3+5 (arithmetic fast paths + self-tail loop) | 0.006 | **14.7×** |

### `(sum 0 5000000)` tail-recursive integer loop (wall time in Go tests)

| Milestone | Time (s) |
|---|---|
| Baseline | 2.99 |
| Phase 2 VM | 0.24 |
| Phase 3+5 (self-tail + fast integer =,-) | 0.05 |

Speedup on tight tail-call loop: **60×**

### `fib(30)` (non-tail double recursion)

| Milestone | Time (s) |
|---|---|
| Phase 3+5 | 0.29 |

---

## Method

```
# Define tak via KL defun, then time with get-time run:
printf '(defun tak (X Y Z) ...) (get-time run) (tak 18 12 6) (get-time run)\n' \
  | ./shen-go/shen
# tak time = t2 - t1
```

---

## Native `hash` (dictionaries)

Per the Shen port-performance recommendations, the kernel's interpreted `hash`
(`sys.kl`: char-code product via `shen.hashkey` + modulo-by-repeated-subtraction
via `shen.mod`) is replaced with a native FNV-1a + true modulo (`kl.PrimHash`,
overriding `hash` in `cmd/shen` between `regist` and `shen.initialise`). Measured
with `bench/dict-bench.shen` (darwin/arm64, S41.1):

| Workload | KL hash | native hash | speedup |
|---|---|---|---|
| fill 2000 short keys | 0.0136 s | 0.0027 s | 5.0× |
| 20k short-key lookups | 0.111 s | 0.0126 s | 8.8× |
| fill 2000 long keys (~25 chars) | 0.039 s | 0.0029 s | 13.4× |
| 20k long-key lookups | 0.352 s | 0.0181 s | 19.4× |

Hash values are not persisted, so only determinism + distribution matter; the
canonical kernel certification (dict-heavy: packages, types, prolog) stays 134/134.

## Notes

- Phase 0: fixed float comparison bug (`mustInteger` → `mustNumber` for `<`, `<=`, `>`, `>=`).
  No performance change.
- Phase 2: `defun` now compiles to a bytecode VM with flat indexed slots (no alist env).
  Both REPL-defined KL `defun` and Shen-level `define` are compiled.
  `lambda`/`freeze`/`trap-error`/`let`/`cond` all compile to bytecode.
  Closures capture upvalues by value at creation time.

## Go 1.27 typed-region benchmark matrix (2026-08-26)

The Go-level matrix lives in `kl/typed_ir_bench_test.go`; equivalent Shen
workloads for interpreter and generated AOT runs live in `bench/typed-ir.shen`.
The matrix is shared by VM and AOT workloads: each case defines a KL function
once, then measures repeated calls with parsing and kernel boot outside the
timed section. The VM benchmark is executable today. AOT tests currently
verify generated-plugin compilation and ABI wiring, but do not provide
comparable wall-clock rows. Run the VM matrix on a clean tree with:

```
go test ./kl -run '^$' -bench 'BenchmarkTypedVM' -benchmem -count=10 \
  | tee /tmp/shen-go-typed-ir.txt
benchstat /tmp/shen-go-typed-ir.txt
```

Coverage includes numeric fixnum/float paths, boolean branches, Unicode string
concatenation, pair/list construction, mutable vectors, higher-order dynamic
application, and fallback-heavy mixed values. Typed IR is enabled by default;
set `SHEN_GO_TYPED_IR=off` for the dynamic control run. Do not compare
single-run `-benchtime=1x` output with `benchstat` results.

Reference run (Apple M4, Go 1.27.0, `-benchtime=100ms`, one sample; use the
ten-sample command above for decisions):

| Benchmark | ns/op | B/op | allocs/op |
|---|---:|---:|---:|
| TypedVMBoolBranch | 768,534 | 5 | 0 |
| TypedVMUnicodeString | 63,780 | 47,378 | 400 |
| TypedVMPairList | 250,816 | 48,014 | 2,000 |
| TypedVMVector | 808,741 | 47 | 2 |
| TypedVMDynamicApply | 1,408,127 | 256,098 | 8,002 |
| TypedVMFallbackHeavy | 1,072,590 | 2,177,803 | 4,000 |

These values are orientation data, not acceptance thresholds; machine, Go
version, and benchmark duration materially affect them.

Clean VM acceptance spot-check on the same Apple M4 (`-benchtime=100ms`, five
samples, medians shown) compared the default guarded path with
`SHEN_GO_TYPED_IR=off`:

| Benchmark | Typed default | Typed off | Result |
|---|---:|---:|---:|
| VMFib | 12.83 ms | 15.44 ms | 1.20x faster |
| VMTak | 4.86 ms | 6.25 ms | 1.29x faster |
| VMSum | 6.21 ms, 1 alloc | 9.81 ms, 199,328 allocs | 1.58x faster; boxing removed |
| VMSumMid | 540 µs | 695 µs | 1.29x faster |
| VMSumNeg | 612 µs | 747 µs | 1.22x faster |

These are a smoke gate rather than a substitute for the ten-sample matrix.

The VM specialization keeps finite numbers in `vmSlot` form where possible and
uses guarded primitive opcodes for arithmetic, predicates, strings, pairs, and
vectors. Every specialized operation checks the canonical primitive binding and
falls back to the ordinary `Obj` call path, preserving redefinition and dynamic
dispatch. AOT emission has the same guard/fallback shape and fuses nested,
side-effect-free scalar call trees, including safe single-use temporary chains
in generated blocks. Annotations remain advisory metadata and never suppress
the fallback.

---

## Remaining gap to shen-cl

shen-cl (SBCL) typically runs `tak(18,12,6)` in under 0.002s.  
Current gap: ~6.5× (0.013s vs ~0.002s).  
Target: within 3–5× of shen-cl.

Further work to close the gap:
- Decision-tree pattern matching compilation
- Inline allocation pooling to reduce GC pressure
