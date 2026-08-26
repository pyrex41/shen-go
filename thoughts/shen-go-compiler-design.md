# shen-go Compiler Design

**Date**: 2026-08-26
**Branch**: `feat/typed-ir`
**Goal**: Preserve Shen semantics while reducing allocation and dispatch costs in
the VM, and establish a guarded typed-region path shared with AOT codegen.

---

## Problem Statement

The current shen-go evaluates user-defined Shen functions entirely through a
KL tree-walking interpreter.  The profiler shows roughly 85% of CPU time in
interpreter plumbing (`eval`, `trampoline`, `apply`) and ~25% in Go GC.
Arithmetic and other "useful" work is sub-1%.

The three concrete bottlenecks are:

1. **`envGet` (O(depth) alist scan)** – every variable lookup walks a
   linked list of `(sym . val)` cons cells.
2. **`envExtend` (heap allocation)** – every function call allocates
   `2 × arity` cons cells to extend the environment.
3. **`eval` tree dispatch** – every node requires a tag switch and a
   9-way special-form dispatch, all resolvable at compile time.

A secondary correctness bug: `<`, `<=`, `>`, `>=` called `mustInteger`
(which truncates floats) instead of `mustNumber`.  Fixed in Phase 0.

---

## Architecture Choice: Bytecode VM (Option A)

**Chosen**: stack-based bytecode VM with per-call flat local-variable frames.

**Rejected alternatives**:
- *AOT-to-Go plugin*: Go plugin ABI is fragile; per-defun compile latency
  is hundreds of ms; no Windows; no clean redefinition.
- *CPS/direct-threaded*: harder to debug, limited advantage over a switch VM
  on Go (no tail-call guarantee in Go).
- *LLVM/native*: too heavy; no JIT in the Go toolchain.

**Why the bytecode VM**:
- Each `defun` is compiled once at define time to a flat `[]Instr` sequence.
- Variables become integer-indexed slots in a per-call `[]Obj` frame.
  `envGet` (O(depth)) → array read O(1).  `envExtend` cons allocation →
  eliminated entirely for the common case.
- Special-form dispatch happens at compile time; the VM loop is a single
  `switch` over `uint8` opcodes.
- Closures (lambda/freeze) carry `upvals []Obj` — a snapshot of captured
  values at creation time.  No alist chain.
- TCO: tail calls set `ControlFlowApply` and return to the existing
  trampoline.  Self-tail calls are trampolined without growing the Go stack.
- `trap-error` is lowered to `(try-catch (freeze body) handler)` at compile
  time, reusing the existing panic/recover primitive.

---

## Instruction Set

```
OP_LOAD_CONST  A   — push consts[A]
OP_LOAD_LOCAL  A   — push locals[A]
OP_STORE_LOCAL A   — locals[A] = pop()
OP_LOAD_GLOBAL A   — push mustSymbol(consts[A]).function  (runtime lookup)
OP_LOAD_UPVAL  A   — push upvals[A]
OP_CALL        A   — non-tail call: pop fn + A args, push result
OP_TAIL_CALL   A   — tail call: set ControlFlowApply, return to trampoline
OP_RETURN          — ctl.Return(pop())
OP_JUMP        A   — pc += A  (signed)
OP_JUMP_FALSE  A   — if pop()==False: pc += A
OP_MAKE_CLOSURE A B — consts[A] is a *BytecodeFunc; pop B upvals from stack;
                       push closure Obj
OP_POP             — discard top of stack
OP_ADD/SUB/MUL/DIV B — guarded numeric operation; B identifies the primitive
OP_LT/LE/GT/GE/EQ B — guarded comparison with dynamic fallback
OP_NOT B            — guarded boolean negation
OP_GUARDED_CONST    — folded literal result plus original operation/operands
OP_GUARDED_PRIM A B — checked typed helper for A arguments; B identifies the
                       primitive used by the dynamic fallback
```

`LOAD_GLOBAL` resolves the symbol's `.function` at call time so that
redefined functions are always picked up without re-compilation.

---

## Compilation Rules

| KL form | Compiled to |
|---|---|
| number/string/bool/nil | `OP_LOAD_CONST` |
| local symbol | `OP_LOAD_LOCAL idx` |
| upvalue symbol | `OP_LOAD_UPVAL idx` |
| global symbol (value pos) | `OP_LOAD_CONST sym` (returns symbol itself) |
| `(defun f (x…) body)` | compile body; bind `scmBytecodeFunc` to `f.function` |
| `(lambda x body)` | inner compile; `OP_MAKE_CLOSURE` with captured upvals |
| `(freeze body)` | same as lambda with 0 params |
| `(let x val body)` | compile val, `OP_STORE_LOCAL`, compile body |
| `(if a b c)` | compile a, `OP_JUMP_FALSE`, compile b, `OP_JUMP`, compile c |
| `(and a b)` | → `(if a b false)` |
| `(or a b)` | → `(if a true b)` |
| `(cond …)` | → nested ifs |
| `(do a b)` | compile a (non-tail), `OP_POP`, compile b (inherits tail) |
| `(type x T)` | compile x unchanged and retain advisory `TypeHint` metadata |
| `(trap-error b h)` | → `(try-catch (freeze b) h)` |
| `(f arg…)` | compile f + args, `OP_CALL`/`OP_TAIL_CALL` |

---

## Integration with Existing Code

- `eval()` special-form handler for `defun` calls `CompileFunc` and stores
  the resulting `scmBytecodeFunc` in the symbol's `.function` slot.
- `apply()` gains a `scmHeadBytecodeFunc` case that calls `vmExec`.
- The existing `scmProcedure` / tree-walker path remains as a fallback for
  any form that fails compilation (should not happen in practice).
- `primDefun` (called by Shen-level compiler) also compiles its argument if
  it receives a `scmProcedure`.
- `ObjString`, `equal`, and the `apply` partial-application logic all handle
  `scmBytecodeFunc` alongside the existing types.

---

## Upvalue Capture

Upvalues are discovered during compilation of a nested lambda/freeze.
When the inner compiler cannot resolve a symbol as a local or already-known
upvalue, it queries the outer compiler's `resolveVar`.  The outer compiler
may itself recurse upward for multi-level closures.

At closure creation (in the outer function's bytecode), the outer compiler
emits `OP_LOAD_LOCAL` or `OP_LOAD_UPVAL` for each captured variable before
`OP_MAKE_CLOSURE`.  The inner function's bytecode accesses them via
`OP_LOAD_UPVAL`.

---

## Phased Plan and status

| Phase | Description | Expected gain |
|---|---|---|
| 0 | Float fix; design doc; benchmark baseline | complete |
| 2 | Bytecode VM (params → slots; compile defun/lambda/let) | complete |
| 3 | Numeric specialization, finite `vmSlot` values, and self-tail loop | complete |
| 4 | Guarded typed IR for VM and AOT scalar regions | implemented; continuing coverage/tests |
| 5 | Pattern-match decision trees and broader allocation work | future |

---

## Things Tried and Reverted

*(fill in as work proceeds)*

- Phase 1 (annotated-AST indexed slots) was folded directly into Phase 2.
  Keeping the tree-walker with slot annotations would have required threading
  a `frame []Obj` through all eval/apply paths — essentially a partial VM —
so it was cleaner to build the full VM directly.

## Go 1.27 typed-region implementation

Go 1.27 generic methods are now used for compile-time checked primitive
registration and arity inference. They improve the host-side API, but they do
not make Shen values homogeneous: the public evaluator, closures, globals,
dynamic calls, and plugin ABI still exchange `Obj`.

The VM uses a conservative tagged-slot representation. A `vmSlot` carries either
an `Obj` or a finite `float64`; fixnums remain directly represented as their
existing tagged `Obj`. Arithmetic, comparisons, boolean operations, selected
string operations, pair/vector access, and conversion predicates use guarded
specialized opcodes where their operands can be checked. Values materialize at
dynamic boundaries, returns, closure capture, and identity-sensitive operations.
Non-finite numbers stay boxed so NaN/Inf identity and float64 rounding semantics
remain unchanged.

The AOT compiler emits the same guarded shape for scalar number, boolean, and
string regions. It fuses nested, side-effect-free primitive call trees,
including single-use temporary chains in generated blocks, and materializes
once at region exit. It does not claim whole-program type proofs. Type
annotations are hints, never proofs: unknown calls, mutable globals, closure
capture, and incompatible joins retain the `Obj` path. Every region checks
both typed mode and canonical primitive bindings before entering the fast
path; the fallback evaluates arguments once through ordinary dynamic lookup,
so redefinition stays observable.

Pairs, vectors, symbols, streams, errors, and callables retain `Obj` payloads
even when their kind is known. Generated plugins preserve `Install(*ControlFlow)`
and dynamic lookup; integration tests compile and invoke that ABI, while AOT
performance parity remains future work. The benchmark matrix in
`bench/RESULTS.md` is the acceptance harness. Keep the typed path enabled by
default only when ten-sample `benchstat` runs show no more than 5% median
regression on key workloads and continue to preserve the large-sum allocation
reduction and tak improvement.
