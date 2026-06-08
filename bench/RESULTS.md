# shen-go Benchmark Results

Machine: Linux amd64  
Kernel: S39.2  
Date: 2026-05-05  

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

---

## Remaining gap to shen-cl

shen-cl (SBCL) typically runs `tak(18,12,6)` in under 0.002s.  
Current gap: ~6.5× (0.013s vs ~0.002s).  
Target: within 3–5× of shen-cl.

Next steps to close the gap:
- Phase 4: decision-tree pattern matching compilation
- Inline allocation pooling to reduce GC pressure
