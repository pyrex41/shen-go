# SIMD evaluation (Go 1.27 `simd` / `simd/archsimd`)

Question: can Go 1.27's SIMD API (`GOEXPERIMENT=simd`) speed up shen-go?

Short answer: not measurably. Nothing on a hot path is shaped like data-parallel work over flat numeric arrays.

Measured on linux/amd64 (Xeon, AVX2 + AVX-512), go1.27.0, 2026-09-26.

## Where the time goes

**Kernel certification** (`kernel/tests/runme.shen`, 6.6 s, `-cpuprofile`):

- About 65% of the time is interpretation: `vmExecSlots`, `apply`, `trampoline`, and tail-call handling.
- About 30% is GC mark work (`gcDrain`, `scanObject`, `scanSpan`) and allocation.
- Candidates for vectorisation total under 2%: `equal`, `MakeVector`/`memclr`, and the vector get and set primitives.

**Startup** (`eval -e '(+ 1 2)'`, ~170 ms; ~45 ms with `SHEN_NO_STDLIB=1`): the time goes to evaluating the stdlib and kernel init, not to reading.

**Native reader** (`BenchmarkShenReadKernelSources`, 157 KB of kernel source in ~3.1 ms):

- Symbol interning (`trieFindOrInsert`) and `cons`/malloc account for most of the time.
- The character-class loops account for about 10%.

## Why SIMD doesn't fit

- `Obj` is a pointer (`*scmHead`), and every Shen value is boxed. Vectors are `[]Obj`, and SIMD cannot load or store pointer slices under the GC. There are no unboxed numeric arrays to vectorise.
- The interpreter's cost is dispatch, pointer chasing and allocation. All of it is scalar and branchy.
- The byte-level work that already benefits from SIMD gets it for free from the Go runtime: `memmove`, `memclr`, string and byte comparison, and `IndexByte`. Go 1.26+'s Green Tea GC also uses vector instructions for span scanning on AVX-512 machines without opting in.
- Reader token runs are too short. Whitespace runs in `kernel/sources/*.shen` average 3.25 bytes, and only 398 of 18,692 runs are 32 bytes or longer. There are 32 line comments in total.

## Prototype results

I tried AVX2 `archsimd.Uint8x32` scanning for whitespace runs and `\\` line comments in `shenreader.go`, behind `//go:build goexperiment.simd && amd64`. `BenchmarkShenReadKernelSources` gave these results (ms/op):

| variant | ms/op |
|---|---|
| baseline (no experiment) | 3.0–3.3 |
| baseline built with `GOEXPERIMENT=simd` | 3.2–3.4 (the flag itself is neutral) |
| pure SIMD scan | **5.4–5.9 (~75% slower)**: broadcast and setup cost per call outweighs 3-byte runs |
| scalar first 16 bytes, then SIMD | 2.9–3.4 (noise) |

The VM micro-benchmarks (`BenchmarkVM*`) are unchanged by the experiment flag.

The prototype was not kept. It adds an experimental build tag and a second code path for no gain.

## What would actually move the needle

These are the levers the profile points at, none of them SIMD:

- Allocation and GC pressure: boxing, cons cells, frames.
- Dispatch overhead in `vmExecSlots`/`apply`.
- AOT via `make precompile` for hot files.

SIMD becomes relevant only if shen-go grows an unboxed numeric vector type, for example a typed-IR `float64`/`int64` array. Bulk primitives over such a type (map, fold, dot product, comparison) could then use the portable `simd` package.
