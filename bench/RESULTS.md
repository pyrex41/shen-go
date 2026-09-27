# shen-go Benchmark Results

Machine: Linux amd64  
Kernel: S39.2  
Date: 2026-05-05  

## Open perf issues, round 2 (2026-09, issues #32 #33 #55 #57)

Measured on a 4-core linux/amd64 container, Go 1.27, user CPU / wall /
peak RSS, three interleaved runs each, `0fa60ea` (before) against this
branch. `kernel/tests/runme.shen` stdout is byte-identical before and
after, apart from its own `run time:` lines.

| Workload | before | after |
|---|---|---|
| `kernel/tests/runme.shen` | 10.8-11.2 s / 7.9-8.3 s / 212 MB | **5.0-5.3 s / 4.7-5.0 s / 122 MB** |
| trivial script (startup, StLib load) | 0.26-0.28 s / 0.21-0.23 s / 69 MB | **0.11-0.12 s / 0.13-0.14 s / 45 MB** |
| `bench/bitlist.shen` (SHA-style bit lists) | 1.55-1.66 s / 1.36-1.45 s | **1.09-1.13 s / 1.15-1.18 s** |
| BFS, put/get visited set (`bfs-portable`) | 1.46-1.67 s | **1.01-1.06 s** |
| `(map (fn inc) L)` + a left fold, 300 x 5000 | 1.26-1.31 s | **0.75-0.80 s** |
| tak(18,12,6) x 400 | 2.96-3.07 s / 70 MB | **2.47-2.52 s / 45 MB** |
| StLib `(floor X)` / `(mod X Y)` per call | ~45 us / ~110 us | **~0.15 us / ~0.12 us** |
| BFS visited set, `shen.x.map` vs put/get | (no native map) | **2.2x faster than put/get** |

What changed, in order of effect:

1. **`let` in the compiled kernel** (`src/compiler.shen`, `codegen`). A KL
   `let` was compiled as `((lambda X Z) Y)`: a closure allocated and
   tail-applied per evaluation, 47% of all bytes allocated on the kernel
   suite. It is now a Go local (`bind`), with binders renamed so repeated
   names never collide. Suite: 12.2 s -> 7.2 s user on its own.
2. **Symbol table**. Symbols were interned in a byte trie whose every node
   held `[256]*trieNode` (2 KiB) plus a symbol: ~40 KiB per long name,
   77 of the suite's 88 MB live heap. A map replaced it; the smaller live
   heap then made GOGC=100 collect 155 times instead of 24, so `cmd/shen`
   now also sets a 64 MiB pointer-free GC ballast (never scanned, never
   resident; off when GOGC or GOMEMLIMIT is set).
3. **VM**: frames sized to each function's computed operand-stack peak
   (`MaxStack`) instead of >= 48 slots cleared on every return; the step
   budget charged per activation instead of per instruction; guarded
   primitives dispatched on a compile-time id instead of a string switch
   on the symbol name; fused `OP_GP1_LOCAL` (hd/tl/cons?/... of a local)
   and `OP_EQ_JF` (`=` and its branch); closures over <= 4 upvalues in one
   allocation.
4. **Compiled kernel**: a defun's full-arity tail call of itself is a
   `goto` while the name is still bound to the running code (117 kernel
   functions); integer constants are fixnums (`MakeInteger`) rather than
   `MakeNumber`, which re-tested integrality on every evaluation; `Call`
   runs an exact-arity native directly. `MakeNumber`/`isPreciseInteger`
   test integrality with a truncation instead of `math.Ilogb`.
5. **Natives**: StLib `floor`/`ceiling`/`round`/`mod`/`div`, bit-identical
   to the library (see `stlibmath_test.go`); the `shen.x.map` host map
   keyed by any Shen value under `=` (feature `shen.x/map`).

Not done, with the reason:

- **#32, a slab/arena allocator for conses.** In Go a slab lives as long
  as any cell in it, and the GC scans its dead cells too, so their
  car/cdr keep other garbage alive transitively: a long-running program
  that keeps one cell per slab could retain everything it ever built.
  Cons allocation plus its GC scan is still ~30% of `bench/bitlist.shen`.
- **#33, compiling StLib to Go or caching its load.** Both need every
  load-time side effect (arities, types, macros, datatypes, package
  externals) replayed exactly; that is its own change. Startup is down
  from ~0.21 s to ~0.13 s wall anyway, since the load runs on the
  compiled kernel.

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
