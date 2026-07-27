# Urdr pure-Shen SHA workload performance

Tracking notes for [issue #20](https://github.com/pyrex41/shen-go/issues/20):
closing the gap between shen-go and shen-cl on Urdr's pure-Shen SHA-256 /
PRNG / world suites (no digest semantics changes).

## Workloads

From the urdr tree (`/Users/reuben/projects/urdr`):

```bash
export SHEN_GO=/Users/reuben/projects/shen-go/.bin/shen-go
/usr/bin/time -lp $SHEN_GO script shen/tests/prng/run-tests.shen
/usr/bin/time -lp $SHEN_GO script shen/tests/world/run-tests.shen
```

Stdout must stay byte-identical aside from the suite's own printed run times.

## Measured series (macOS arm64, interleaved A/B, user CPU)

| Binary | prng user (typ.) | world user (typ.) | Notes |
|--------|------------------|-------------------|-------|
| baseline `8775fc3` (pre-branch) | ~5–8s | ~3–5s | master before this branch |
| after call-path alloc cuts `f5e1899` | ~3.3s wall claimed | — | slab merge + skip arg copies |
| after arity/fn native `e61c724` | ~4.3–6.1s | ~3.5–5.5s | ~25% vs prior on loaded box |
| after VM frame pool (this branch tip) | **~2.5–3.3s** | **~2.6–2.8s** | freelist + OP_CALL fast paths |
| shen-cl (SBCL, same machine) | **~0.55–0.7s** | **~0.75–0.8s** | target |

Absolute numbers swing with machine load; ratios from interleaved runs are more
reliable. Frame-pool tip vs `e61c724` is roughly **~25–35% less user CPU** on
prng and **~20–30%** on world in quiet interleaved runs.

## What landed on this branch

1. **`035b6af`** — `-cpuprofile` / `-memprofile` on `cmd/shen`.
2. **`f5e1899`** — cut per-call slice copies on the VM call path; one slab for
   locals+stack; exact-arity bytecode skips defensive args copy.
3. **`e61c724`** — native `arity` / `fn` with lambdatable memoization
   (`kl/kernelfast.go`).
4. **VM frame freelist + OP_CALL fast paths** — recycle activation slabs on
   `ControlFlow` (min cap 48, pool ≤128, larger slabs displace smaller when
   full); exact-arity bytecode/native OP_CALL avoids a trampoline hop when the
   callee returns.

## Profiling takeaways (prng)

- Before the freelist, `vmExec`'s `make([]Obj, …)` was ~72% of alloc_space
  (~1.2 GB / run) and a large share of GC CPU.
- After the freelist, that allocation disappears from the alloc profile;
  remaining heap is mostly semantic `cons` / `MakeInteger` / startup natives.
- Residual CPU is interpreter overhead: `vmExec` dispatch, `tick`, primitive
  wrappers, and still ~5× behind SBCL on this workload.

## Remaining gap vs shen-cl

Still about **4–5×** on prng user CPU. High-leverage next candidates (not done):

- More primitive inlining / direct native dispatch without `ctl.data` setup.
- Optional precompile of hot urdr SHA files to Go (see `SPIKE-compile-to-go.md`
  and `make precompile`) — low risk for a bench-only path if stdout stays
  identical.
- Broader compiler improvements (better self-tail recognition, specialized
  fixnum ops already partially present).

Do **not** change Urdr digest / bit semantics while chasing port speed.
