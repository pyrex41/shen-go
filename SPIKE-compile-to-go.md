# Spike: compile user functions to Go and load them at runtime via Go plugin

**Branch:** `spike/compile-to-go-plugin` (throwaway worktree off `master`)
**Date:** 2026-06-08 · **Machine:** darwin/arm64, go 1.25.6 · **Kernel:** S41.1

## Question

shen-go runs *runtime-defined* functions on the bytecode VM (`vmExec`), ~7× slower
than shen-cl (SBCL native) on recursion-heavy code. The dormant `KL → IR → Go`
pipeline that generates the kernel could also compile **user** functions to Go and
`plugin`-load them. The old design doc rejected this ("plugin ABI fragile, ~100s ms
compile latency") but never measured it. Does it actually help, and is the latency
tolerable?

## What was built (all reusing the existing pipeline)

- **Stage 0** — `cmd/plugintest/{plugin,host}`: a hand-written `tak` plugin proving Go
  plugins work on darwin/arm64 and that a plugin-provided `Obj` is callable through
  the host's trampoline. ✅ `(tak 18 12 6) = 7`.
- **Stages 1+2** — `cmd/kl/plugin.go`: a new native **`(go-build-and-load RawGoPath ExportName)`**.
  Given `bc->go`'s output it appends a self-contained `Install` hook (+ the runtime-bound
  globals `bc->go` can't know — `ns2_1set`/`try_1catch`), shells out to
  `go build -buildmode=plugin`, then `plugin.Open` + `Lookup("Install")` + runs it
  (which `defun`-binds the function). The full opt-in flow, driven from the running
  `cmd/kl` image:
  `(compile-file f.kl f.tmp)` → `(bc->go Cg "Main" true f.tmp f_raw.go)` → `(go-build-and-load …)`.
  Repro scripts: `compiled/spike-e2e.kl`, `compiled/spike-bench2.kl`.
  VM binding survives on any error (VM stays the default/fallback).

Compiled functions don't retain KL source (`BytecodeFunc` keeps only bytecode), so the
flow consumes the `defun` form directly rather than a function name.

## Results

Same workloads, three runtimes (lower is better):

| Benchmark | Bytecode VM | **Plugin (compile-to-Go)** | shen-cl (SBCL) |
|---|---|---|---|
| `tak(24,16,8)` ×20 | 2.930 s | **1.541 s** | 0.377 s |
| `fib(30)` | 0.182 s | **0.108 s** | 0.028 s |

**Speedups:** plugin is **1.90×** (tak) / **1.69×** (fib) faster than the VM — but still
**~4×** slower than SBCL (4.1× tak, 3.9× fib).

**Compile latency (per function):** `go build -buildmode=plugin` ≈ 700–965 ms +
`plugin.Open` ≈ 270–290 ms ≈ **~1 s/function** warm. (The first build in a fresh process
that must also recompile the `kl` package is the high end; steady-state is ~700 ms.)
So the doc's "hundreds of ms" was about right for the build step; ~1 s end-to-end.

Correctness: plugin-backed `tak(18,12,6)=7`, `fib(20)=6765`.

## Why only ~2×, not ~7×

The existing codegen emits Go that **still uses the trampoline (`Call`/`TailApply`) and
boxed `Obj` values**, and does **not** apply the VM's Phase 3+5 wins (inlined `OP_SUB`/`OP_LT`,
self-tail-call loop) — every arithmetic op is a `PrimNumberSubtract` call, every recursion a
trampolined `Call`. Compiling to Go removes the **opcode-dispatch loop + per-frame slice
allocations** (→ the ~2× win) but keeps **boxing + trampoline** (→ the residual ~4× gap to
SBCL's unboxed, natively-recursive code). Reaching SBCL-class speed would need a *new*
optimizing Go backend (unboxed int64 where provable, native Go recursion), which is a much
larger effort than reviving the existing pipeline.

## Verdict

Empirically: the compile-to-Go plugin path is a **real but modest ~1.8× win** over the
already-optimized VM, at **~1 s compile latency per function** plus Go's plugin constraints
(identical-toolchain builds, no unload — each redefinition is a fresh `.so`, no Windows).

Worth productionizing **only** for a precompile-once workflow (compile hot functions at
startup, behind an opt-in `-jit` flag in `cmd/shen`), never auto-on-every-`defun` (the ~1 s
latency would wreck REPL responsiveness — the design doc's original concern, now confirmed
with numbers). If the goal is closing the gap to SBCL, the higher-leverage path is a new
optimizing backend or further VM numeric specialization — not this pipeline as-is.

## Usage (productionized: hybrid VM + AOT compile-to-Go)

The spike was productionized into an opt-in **whole-file AOT** path that coexists with the
VM. The VM remains the default for everything (REPL included); a precompiled `.so` overrides
`.function` only for the functions it contains.

**1. Precompile a whole Shen (or KL) file → a Go plugin `.so` (offline, ~1s/fn):**
```
make precompile FILE=bench/hot.shen OUT=hot.so
```
This runs the `cmd/kl` toolchain: Shen→KL (`bootstrap`) → IR (`compile-file`) → Go (`bc->go`)
→ `go build -buildmode=plugin`. It emits `hot.so` plus `hot.so.fns` (a `NAME ARITY` manifest).

**2. Run shen with the precompiled functions loaded at startup (~0.27s load, no compile):**
```
./shen -precompiled hot.so        # repeatable, or comma-separated
```
At startup (after `shen.initialise`, before the REPL) each `.so` is `plugin.Open`ed, its
`Install` hook binds every function, and the `.fns` manifest is replayed through
`shen.store-arity` so the functions resolve at the Shen REPL. A REPL-typed `(defun …)` is
still VM-compiled; VM and plugin functions call each other freely. A bad/missing `.so` warns
and the REPL continues on the VM. `(load-native "x.so")` does the same from inside a session.

**Measured (S41.1, darwin/arm64):** `bench-tak(24,16,8)×20` = 3.15s plugin vs 5.02s VM
(**1.59×**) — same Shen-`define` logic both sides. The plugin path is for hot functions you
precompile once; it is never auto-applied to REPL defuns (the ~1s build latency stays at build
time, not in the REPL).

**Constraints:** the `.so` and the `shen` that loads it must be built from the same module +
Go toolchain (both from this tree); Go plugins can't be unloaded (redefine = rebuild); no
Windows. The generated Go still boxes ints + uses the trampoline, hence ~1.6–1.9× over the VM,
not SBCL-class.

## Files

- `cmd/plugintest/plugin/tak.go`, `cmd/plugintest/host/main.go` — Stage 0 ABI proof.
- `cmd/kl/plugin.go` — `go-build-plugin` (build-only, AOT), `go-build-and-load` (build+load,
  dev), `emit-arities` (manifest); bound in `cmd/kl/main.go`.
- `cmd/shen/plugin.go` — startup loader, `-precompiled` flag wiring, `load-native` primitive,
  arity replay; wired in `cmd/shen/main.go`.
- `compiled/precompile.kl` — the AOT driver; `Makefile` `precompile` target.
- `bench/hot.shen` — sample whole-file unit (tak/fib/bench-tak).
- `compiled/onefn.kl`, `onefn-fib.kl`, `spike-*.kl` — original spike repro/benchmark.
- Generated artifacts (`*.tmp`, `*_raw.go`, `_in.*`, `plugintmp*/`, `pluginso*/`) are gitignored.
