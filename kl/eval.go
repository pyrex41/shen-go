package kl

import (
	"fmt"
	"os"
	"runtime"
)

type ControlFlowKind int

const (
	ControlFlowReturn ControlFlowKind = iota
	ControlFlowEval
	ControlFlowApply
)

type ControlFlow struct {
	// controlFlowReturn: result = data[0]
	// controlFlowApply: fn, args = data[0], data[1], data[2] ...
	// controlFlowEval: exp, env = data[0], data[1]
	kind ControlFlowKind

	// data[pos : len(data)] is the arguments to current function.
	// Why it needs to be a stack?
	// Most of the time, an array is sufficient. But when eval expression like
	// (f (g1 h) (g2 n) ...), the stack is used to store the temporary arguments
	data []Obj
	pos  int

	// stepLimit, when > 0, caps the number of evaluation steps (trampoline
	// iterations + VM instructions) a single Eval may take before it aborts
	// with a catchable MakeError("eval step limit exceeded"). The default
	// zero means *unlimited* — production never sets it, so the evaluator's
	// behaviour is unchanged.
	//
	// The motivation is fuzzing: a KLambda program may legally fail to
	// terminate (e.g. `(do (defun f (X) (f X)) (f 0))` tail-recurses
	// forever without growing the Go stack, so no panic ever fires). That
	// is the halting problem, not a memory-safety bug — but to a fuzzer it
	// looks like a stuck worker and hides the crashes we actually care
	// about. A step limit lets the harness turn "ran too long" into an
	// ordinary catchable error so the fuzz loop keeps making progress.
	stepLimit int64
	steps     int64

	// Frame arena: VM activation slabs (locals + operand stack) are carved
	// LIFO out of contiguous blocks. takeFrame bumps frameTop; putFrame resets
	// it to the frame's base. Blocks double as recursion deepens. Spare
	// blocks above the current one are kept until a trim point (a return to
	// an empty arena at a Go entry point, a recover site, or the periodic
	// idle check in frameCrossed), where they are cut back to at most
	// frameRetainSlots, so a one-shot deep recursion cannot pin an unbounded
	// amount of memory forever but a repeated one does not reallocate its
	// blocks on every descent. Replaces a freelist capped at 128 slabs which
	// allocated one slab per frame past the cap and scanned the whole pool
	// on every return (issue #50). Single-threaded per ControlFlow — no
	// locking.
	frameBlocks [][]vmSlot
	frameCur    int
	frameTop    int
	// frameHigh is the highest block index any frame has entered since the
	// last trim; frameCross counts block-boundary returns since then.
	frameHigh  int
	frameCross int
	// frameCap is how far into the current block frames may reach before
	// frameEnsure is called. It is below the block's end while the block is
	// not fully used, and so also bounds the part of the block that may hold
	// stale slots. frameDirty saves it per block while a later block is current.
	frameCap   int
	frameDirty []int
	// frameEpoch is the GC cycle stale slots were last cleared for (vmAttend).
	frameEpoch uint32

	// vmFrames holds the callers of the running bytecode activations, up to
	// vmDepth. len(vmFrames) bounds the records that may be stale.
	vmFrames []vmFrame
	vmDepth  int
}

// frameBlockMin is the size in slots of the first arena block: 512 slots =
// 8 KiB, i.e. room for ten minimum-size frames. It is chosen small on purpose:
// every ControlFlow that runs any bytecode pays for one zeroed block, and
// production constructs throwaway ControlFlows per call (equiv.go safeCall and
// evalQuiet, one per audit probe) as do many tests. A deep recursion reaches
// any depth in O(log depth) block transitions anyway since blocks double.
const frameBlockMin = 512

// frameRetainSlots bounds the spare arena blocks kept at a trim point
// (1<<19 slots = 8 MiB). Blocks double from frameBlockMin, so the budget keeps
// blocks 1..9 (523,264 slots); with block 0 the retained arena holds 523,776
// slots = 10,912 minimum-size (48-slot) frames, i.e. a non-tail recursion up
// to ~10,900 deep re-descends with no allocation after a trim, and a deeper
// one reallocates only the blocks above that, once, on its first descent
// after a trim. Trimming happens only where frameTrim is called: at a Go
// entry point returning to an empty arena (frameSettle), at a recover site
// (frameReset) and at the periodic idle check (frameCrossed). It never
// happens on an ordinary return that merely crosses a block boundary: that
// would drop and re-zero the top 8 MiB block on every return past depth
// ~10,900 of a repeated deep recursion.
const frameRetainSlots = 1 << 19

// frameTrimInterval is the number of block-boundary returns between the trim
// decisions taken inside a long-lived enclosing frame. A loop that runs
// inside one top-level form (a server loop, a long computation that once
// recursed deeply) never returns to an empty arena, so without these
// decisions one deep call would pin its peak for the rest of that form
// (measured: 284 MB live after a single 300,000-deep call, against 36 MB
// with them; CPU flat). Such a decision keeps every block some frame has
// entered since the previous one, so a repeated deep recursion never loses
// the blocks it uses; only blocks idle for a whole interval are released
// past the budget. 1024 is large against the ~10 crossings per descent of a
// 20,000-deep recursion and small against the time a shallow workload needs
// to make that many.
const frameTrimInterval = 1024

// frameHeadroom is the operand-stack capacity reserved per frame beyond
// Nlocals. If a frame needs more, append reallocates its operand stack away
// from the arena, which is rare and still correct.
const frameHeadroom = 16

// minFrameCap is the minimum slab capacity. Small frames are rounded up so a
// shallow frame keeps the same operand headroom the old freelist gave it
// (measured: without the floor the kernel test suite reallocated ~10x more
// operand stacks off-slab). Arena space is free, so the round-up only costs
// the clear on release.
const minFrameCap = 48

// frameMark is an arena position: where a Go entry point or recover site came in.
type frameMark struct{ cur, top, depth int }

// frameMarkNow returns the current arena position, for a recover site to
// frameReset to after a panic abandons the frames above it.
func (ctl *ControlFlow) frameMarkNow() frameMark {
	return frameMark{ctl.frameCur, ctl.frameTop, ctl.vmDepth}
}

// frameSoftChunk is how far past a frame frameEnsure moves frameCap. It sets how
// often frameEnsure runs while a block fills, and how many untouched slots a
// stale clear may touch.
const frameSoftChunk = 1024

// frameMaxSlots bounds the arena at 1 GiB. Bytecode recursion no longer uses
// the goroutine stack, whose 1 GB limit used to end a runaway recursion, so
// this has to, with the same exit status.
const frameMaxSlots = 1 << 26

// frameEnsure makes room for a frame of size slots at base, whose first args
// are already in place, and returns its base: a frame that does not fit in the
// current block starts a new one, with args copied there.
func (ctl *ControlFlow) frameEnsure(base, size int, args []vmSlot) int {
	need := base + size
	if ctl.frameCur < len(ctl.frameBlocks) && need <= len(ctl.frameBlocks[ctl.frameCur]) {
		ctl.frameCap = min(need+frameSoftChunk, len(ctl.frameBlocks[ctl.frameCur]))
		return base
	}
	if ctl.frameCur < len(ctl.frameDirty) {
		ctl.frameDirty[ctl.frameCur] = ctl.frameCap
	}
	ctl.advanceFrameBlock(size)
	for len(ctl.frameDirty) <= ctl.frameCur {
		ctl.frameDirty = append(ctl.frameDirty, 0)
	}
	blk := ctl.frameBlocks[ctl.frameCur]
	ctl.frameCap = min(max(size+frameSoftChunk, ctl.frameDirty[ctl.frameCur]), len(blk))
	copy(blk, args)
	return 0
}

// frameLeaveBlock returns to block cur from a later one. frameTop is the
// caller's to set. It counts as a crossing for the periodic idle check.
func (ctl *ControlFlow) frameLeaveBlock(cur int) {
	ctl.frameDirty[ctl.frameCur] = ctl.frameCap
	ctl.frameCur = cur
	ctl.frameCap = ctl.frameDirty[cur]
	ctl.frameCrossed()
}

// frameClearStale zeroes every slot above the top that may hold a value, so that
// a dead value stays reachable for at most one GC cycle after its frame returned.
func (ctl *ControlFlow) frameClearStale() {
	if ctl.frameCur < len(ctl.frameBlocks) {
		blk := ctl.frameBlocks[ctl.frameCur]
		if ctl.frameTop < ctl.frameCap {
			clear(blk[ctl.frameTop:ctl.frameCap])
			ctl.frameCap = max(min(ctl.frameTop+frameSoftChunk, len(blk)), ctl.frameTop)
		}
		for k := ctl.frameCur + 1; k < len(ctl.frameBlocks) && k < len(ctl.frameDirty); k++ {
			clear(ctl.frameBlocks[k][:ctl.frameDirty[k]])
			ctl.frameDirty[k] = 0
		}
	}
	clear(ctl.vmFrames[ctl.vmDepth:])
	ctl.vmFrames = ctl.vmFrames[:ctl.vmDepth]
}

func (ctl *ControlFlow) vmFramesGrow() {
	d := len(ctl.vmFrames)
	if d < cap(ctl.vmFrames) {
		ctl.vmFrames = ctl.vmFrames[:min(d+64, cap(ctl.vmFrames))]
		return
	}
	frames := make([]vmFrame, max(2*d, 64))
	copy(frames, ctl.vmFrames)
	ctl.vmFrames = frames
}

func vmStackOverflow() {
	fmt.Fprintln(os.Stderr, "fatal error: stack overflow: bytecode recursion exceeds the 1 GiB frame arena")
	os.Exit(2)
}

// advanceFrameBlock moves to the next block, allocating (or replacing a spare
// that is too small) so it holds at least need slots.
func (ctl *ControlFlow) advanceFrameBlock(need int) {
	if len(ctl.frameBlocks) > 0 {
		ctl.frameCur++
	}
	if ctl.frameCur > ctl.frameHigh {
		ctl.frameHigh = ctl.frameCur
	}
	ctl.frameTop = 0
	if ctl.frameCur < len(ctl.frameBlocks) && len(ctl.frameBlocks[ctl.frameCur]) >= need {
		return
	}
	total := need
	for _, b := range ctl.frameBlocks[:min(ctl.frameCur, len(ctl.frameBlocks))] {
		total += len(b)
	}
	if total > frameMaxSlots {
		vmStackOverflow()
	}
	size := frameBlockMin
	if ctl.frameCur > 0 {
		size = 2 * len(ctl.frameBlocks[ctl.frameCur-1])
	}
	for size < need {
		size *= 2
	}
	blk := make([]vmSlot, size)
	if ctl.frameCur < len(ctl.frameBlocks) {
		ctl.frameBlocks[ctl.frameCur] = blk
	} else {
		ctl.frameBlocks = append(ctl.frameBlocks, blk)
	}
}

// frameUnwind resets the arena to mark. Slots above it are left for the next
// stale clear; the frame records go now, since a crossed flag must not survive.
func (ctl *ControlFlow) frameUnwind(mark frameMark) {
	if ctl.vmDepth > mark.depth {
		clear(ctl.vmFrames[mark.depth:ctl.vmDepth])
	}
	if mark.cur != ctl.frameCur {
		ctl.frameDirty[ctl.frameCur] = ctl.frameCap
		ctl.frameCap = ctl.frameDirty[mark.cur]
	}
	ctl.frameCur, ctl.frameTop, ctl.vmDepth = mark.cur, mark.top, mark.depth
}

// frameReset unwinds the arena to mark and trims the spare blocks to the
// retention budget. Recover sites call it so frames abandoned by a panic
// neither leak arena space nor retain objects, and so retention is bounded
// after every caught error.
func (ctl *ControlFlow) frameReset(mark frameMark) {
	ctl.frameUnwind(mark)
	ctl.frameTrim(0)
}

// frameSettle is the normal-return counterpart of frameReset at a Go entry
// point (Eval, Try): when the entry point took the arena empty and every frame
// since has been released, this is a top-level return and the spare blocks
// are trimmed to the retention budget.
func (ctl *ControlFlow) frameSettle(mark frameMark) {
	if mark == (frameMark{}) && ctl.frameCur == 0 && ctl.frameTop == 0 {
		ctl.frameTrim(0)
	}
}

// frameCrossed accounts for a return that crossed a block boundary and, every
// frameTrimInterval of them, trims the spare blocks beyond the budget that no
// frame has entered since the previous decision.
func (ctl *ControlFlow) frameCrossed() {
	ctl.frameCross++
	if ctl.frameCross >= frameTrimInterval {
		ctl.frameTrim(ctl.frameHigh + 1)
	}
}

// frameTrim releases the spare blocks above the current one that exceed
// frameRetainSlots, keeping at least blocks [0:keep), and starts a new trim
// epoch.
func (ctl *ControlFlow) frameTrim(keep int) {
	k, total := ctl.frameCur+1, 0
	for k < len(ctl.frameBlocks) && total+len(ctl.frameBlocks[k]) <= frameRetainSlots {
		total += len(ctl.frameBlocks[k])
		k++
	}
	if k < keep {
		k = keep
	}
	if k < len(ctl.frameBlocks) {
		clear(ctl.frameBlocks[k:])
		ctl.frameBlocks = ctl.frameBlocks[:k]
		if len(ctl.frameDirty) > k {
			ctl.frameDirty = ctl.frameDirty[:k]
		}
	}
	ctl.frameHigh = ctl.frameCur
	ctl.frameCross = 0
}

// stepLimited reports whether this ControlFlow carries a step budget.
//
// stepLimit is write-once: it is set when the ControlFlow is constructed (only
// the fuzz harness and its tests ever do so — production leaves it zero) and is
// never mutated while evaluation is running. The hot loops therefore hoist this
// call *out* of their loop into a local and test that local per step, instead of
// re-loading ctl.stepLimit through the pointer on every instruction. The load
// was not free: the compiler cannot cache it across the eval/apply/vmExec calls
// in the loop body (they all take ctl), so every VM instruction paid a memory
// load + compare for a counter that is off in every non-test process.
//
// Anyone adding a way to change the budget mid-run must revisit that hoist.
func (ctl *ControlFlow) stepLimited() bool {
	return ctl.stepLimit != 0
}

// tick accounts one evaluation step against the step limit. It must only be
// called when ctl.stepLimited() is true — the callers hoist that test out of
// their loop (see stepLimited). The abort lives in a separate function so tick
// itself stays inlinable.
func (ctl *ControlFlow) tick() {
	ctl.steps++
	if ctl.steps > ctl.stepLimit {
		ctl.tripStepLimit()
	}
}

// SetStepLimit sets the evaluation budget (see stepLimited); zero means
// unlimited. It is for harnesses outside this package, such as tests of
// generated code.
func (ctl *ControlFlow) SetStepLimit(n int64) {
	ctl.stepLimit = n
}

// Tick charges one step when a step limit is set. Generated code calls it
// on each iteration of a loop that does not go through the trampoline (a
// compiled self-tail-call), so the budget still bounds it.
func (ctl *ControlFlow) Tick() {
	if ctl.stepLimit != 0 {
		ctl.tick()
	}
}

func (ctl *ControlFlow) tripStepLimit() {
	panic(MakeError("eval step limit exceeded"))
}

func (ctl *ControlFlow) TailEval(exp Obj, env Obj) {
	ctl.data = ctl.data[:ctl.pos]
	ctl.data = append(ctl.data, exp)
	ctl.data = append(ctl.data, env)
	ctl.kind = ControlFlowEval
}

func (ctl *ControlFlow) TailApply(f Obj, args ...Obj) {
	ctl.tailApplySlice(f, args)
}

// tailApplySlice is TailApply taking an explicit slice, letting hot callers
// (the VM's OP_CALL/OP_TAIL_CALL) hand over a view of their operand stack
// without first materialising a temporary []Obj for varargs. args must not
// alias ctl.data (all VM callers pass frame-local slices; the append below
// copies the values into ctl.data immediately).
func (ctl *ControlFlow) tailApplySlice(f Obj, args []Obj) {
	ctl.data = ctl.data[:ctl.pos]
	ctl.data = append(ctl.data, f)
	ctl.data = append(ctl.data, args...)
	ctl.kind = ControlFlowApply
}

// tailApplySlots is the VM-slot equivalent of tailApplySlice. It materializes
// values directly into ctl.data, avoiding a temporary []Obj at dynamic call
// boundaries while retaining unboxed numbers inside bytecode-to-bytecode calls.
func (ctl *ControlFlow) tailApplySlots(f Obj, args []vmSlot) {
	ctl.data = ctl.data[:ctl.pos]
	ctl.data = append(ctl.data, f)
	for _, arg := range args {
		ctl.data = append(ctl.data, arg.objValue())
	}
	ctl.kind = ControlFlowApply
}

func (ctl *ControlFlow) Return(result Obj) {
	ctl.data = ctl.data[:ctl.pos]
	ctl.data = append(ctl.data, result)
	ctl.kind = ControlFlowReturn
}

func (ctl *ControlFlow) Get(n int) Obj {
	return ctl.data[ctl.pos+n]
}

// trampoline is introduced for tail call optimization.
func trampoline(ctl *ControlFlow) Obj {
	limited := ctl.stepLimited()
	for ctl.kind != ControlFlowReturn {
		if limited {
			ctl.tick()
		}
		switch ctl.kind {
		case ControlFlowEval:
			eval(ctl)
		case ControlFlowApply:
			apply(ctl)
		}
	}
	ret := ctl.data[ctl.pos]
	ctl.data = ctl.data[:ctl.pos]
	return ret
}

func evalExp(e *ControlFlow, exp Obj, env Obj) Obj {
	e.TailEval(exp, env)
	return trampoline(e)
}

func evalIf(e *ControlFlow, a, b, c Obj, env Obj) {
	t := evalExp(e, a, env)
	switch t {
	case True:
		e.TailEval(b, env)
		return
	case False:
		e.TailEval(c, env)
		return
	}
	// Match the compiled VM path (OP_JUMP_FALSE), which raises a catchable
	// Shen error so trap-error can intercept a non-boolean condition.
	panic(MakeError("if requires a boolean"))
}

func evalAnd(e *ControlFlow, a, b Obj, env Obj) {
	if evalExp(e, a, env) == False {
		e.Return(False)
		return
	}
	e.TailEval(b, env)
}

func evalOr(e *ControlFlow, a, b Obj, env Obj) {
	if evalExp(e, a, env) == True {
		e.Return(True)
		return
	}
	e.TailEval(b, env)
}

// partialApply works when Required > providArgs
func partialApply(required int, providArgs []Obj, env Obj, proc Obj) Obj {
	// Partial apply...
	// (f x y z) => (lambda (z) (f x y z)) with x y in env
	symbols := makeTempSymbols(required)
	env1 := envExtend(env, symbols[:len(providArgs)], providArgs)

	args := Nil
	for i, count := len(symbols)-1, required-len(providArgs); count > 0; count-- {
		args = cons(symbols[i], args)
		i--
	}

	body := Nil
	for i := len(symbols) - 1; i >= 0; i-- {
		body = cons(symbols[i], body)
	}
	body = cons(proc, body)

	return makeProcedure(args, body, env1)
}

func makeTempSymbols(n int) []Obj {
	ret := make([]Obj, n)
	for i := 0; i < n; i++ {
		ret[i] = MakeSymbol(fmt.Sprintf("tmp%d", i))
	}
	return ret
}

func Call(e *ControlFlow, f Obj, args ...Obj) Obj {
	// The compiled kernel calls natives almost exclusively: run an
	// exact-arity one without captured arguments directly, and enter the
	// trampoline only if it tail-calls. This is the same shortcut as the
	// VM's OP_CALL, and like the trampoline it charges the step budget.
	if f != nil && !isFixnum(f) && *f == scmHeadNative {
		if nf := MustNative(f); len(nf.captured) == 0 && nf.require == len(args) {
			if e.stepLimit != 0 {
				e.tick()
			}
			e.tailApplySlice(f, args)
			nf.fn(e)
			if e.kind == ControlFlowReturn {
				ret := e.data[e.pos]
				e.data = e.data[:e.pos]
				return ret
			}
			return trampoline(e)
		}
	}
	e.tailApplySlice(f, args)
	return trampoline(e)
}

// callSlice is Call taking an explicit args slice (see tailApplySlice for the
// aliasing contract).
func (ctl *ControlFlow) callSlice(f Obj, args []Obj) Obj {
	ctl.tailApplySlice(f, args)
	return trampoline(ctl)
}

// debugRecover enables a stderr trace (expression, message, Go stack) each
// time Eval recovers a raised KL error. KL error paths are ordinary control
// flow -- (trap-error (tl 5) H) is a normal program -- so the runtime never
// writes to the program's stdout there, and by default not to stderr either;
// set SHEN_DEBUG_RECOVER=1 when debugging the interpreter itself. A non-Obj Go
// panic (a runtime.Error or a bare string from a native) is a bug rather than
// control flow, so traceRecover always reports those on stderr.
var debugRecover = os.Getenv("SHEN_DEBUG_RECOVER") != ""

// traceRecover writes the trace for a panic r recovered in where while
// evaluating exp. A raised KL error is reported only when debugRecover is
// set; anything else is reported unconditionally. Nothing goes to stdout.
func traceRecover(where string, exp Obj, r interface{}) {
	if x, ok := r.(Obj); ok && IsError(x) {
		if !debugRecover {
			return
		}
		fmt.Fprintf(os.Stderr, "kl: recovered in %s: %s\n  raised: %s\n", where, ObjString(exp), mustError(x).err)
	} else {
		fmt.Fprintf(os.Stderr, "kl: recovered in %s: %s\n  panic: %v\n", where, ObjString(exp), r)
	}
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stderr.Write(buf[:n])
}

// Eval evaluates exp at top level. A KL error raised and not caught inside
// exp is returned as the raised error object itself, so callers see the
// kernel's message; any other Go panic is returned as an Error carrying the
// panic value's text.
func Eval(e *ControlFlow, exp Obj) (res Obj) {
	fmark := e.frameMarkNow()
	defer func() {
		if r := recover(); r != nil {
			e.frameReset(fmark)
			traceRecover("Eval", exp, r)
			if x, ok := r.(Obj); ok && IsError(x) {
				res = x
				return
			}
			res = MakeError(fmt.Sprint(r))
		}
	}()
	res = evalExp(e, exp, Nil)
	e.frameSettle(fmark)
	return
}

type tryResult struct {
	e    *ControlFlow
	data Obj
}

func Try(e *ControlFlow, f Obj) (res tryResult) {
	fmark := e.frameMarkNow()
	defer func() {
		if err := recover(); err != nil {
			e.frameReset(fmark)
			if val, ok := err.(Obj); ok {
				if IsError(val) {
					res = tryResult{e: e, data: val}
					return
				}
			}
			// Unexpected panic?
			traceRecover("Try", f, err)
			res = tryResult{e: e, data: MakeError(fmt.Sprintf("%v", err))}
		}
	}()
	// f must be a 0-arity callable (native thunk or bytecode closure).
	val := Call(e, f)
	e.frameSettle(fmark)
	res = tryResult{e: e, data: val}
	return
}

func (t tryResult) Catch(f Obj) Obj {
	if IsError(t.data) {
		return Call(t.e, f, t.data)
	}
	return t.data
}

func apply(ctl *ControlFlow) {
	f := ctl.data[ctl.pos]
	args := ctl.data[ctl.pos+1:]
	switch *f {
	case scmHeadBytecodeFunc:
		bf := mustBytecodeFunc(f)
		if len(args) == bf.fn.Arity {
			// Exact arity (the common case): no defensive copy needed. args
			// aliases ctl.data, but vmExec copies it into a fresh locals slice
			// before anything can mutate ctl.data.
			vmExec(ctl, bf, args)
			return
		}
		// Partial/over-application: the over-application path reads args after
		// an inner Call has clobbered ctl.data, so it needs its own copy.
		argsCopy := make([]Obj, len(args))
		copy(argsCopy, args)
		vmApply(ctl, f, argsCopy)
		return
	case scmHeadProcedure:
		args := ctl.data[ctl.pos+1:]
		proc := mustProcedure(f)
		switch {
		case len(args) < proc.arity:
			ctl.Return(partialApply(proc.arity, args, proc.env, f))
			return
		case len(args) == proc.arity:
			newEnv := envExtend(proc.env, proc.arg, args)
			ctl.TailEval(proc.body, newEnv)
			return
		case len(args) > proc.arity:
			newEnv := envExtend(proc.env, proc.arg, args[:proc.arity])
			res := evalExp(ctl, proc.body, newEnv)
			ctl.TailApply(res, args[proc.arity:]...)
			return
		}
	case scmHeadNative:
		fn := MustNative(f)
		provided := len(fn.captured) + len(args)
		required := fn.require
		if provided != required {
			panic(MakeError("partial apply unsupported"))
		}
		if len(fn.captured) > 0 {
			// Captured data should come fisrt, then the arguments.
			ctl.data = append(ctl.data, make([]Obj, len(fn.captured))...)
			dst := ctl.data[len(ctl.data)-len(args):]
			copy(dst, args)
			dst = ctl.data[ctl.pos+1 : ctl.pos+1+len(fn.captured)]
			copy(dst, fn.captured)
		}
		fn.fn(ctl)
		return
	}
	panic(MakeError(fmt.Sprintf("can't apply non function: %s", ObjString(f))))
}

func envGet(env Obj, sym Obj) (Obj, bool) {
	for env != Nil {
		fr := mustEnv(env)
		if fr.sym == sym {
			return fr.val, true
		}
		env = fr.next
	}
	return nil, false
}

func envExtend(env Obj, symbols, values []Obj) Obj {
	for i := 0; i < len(symbols); i++ {
		env = envCons(symbols[i], values[i], env)
	}
	return env
}

func eval(e *ControlFlow) {
	exp := e.data[e.pos]
	env := e.data[e.pos+1]

	switch *exp {
	// handle constant
	case scmHeadNumber, scmHeadString, scmHeadVector, scmHeadBoolean, scmHeadNull,
		scmHeadProcedure, scmHeadBytecodeFunc /* , scmHeadPrimitive */ :
		e.Return(exp)
		return
	case scmHeadSymbol:
		if val, ok := envGet(env, exp); ok {
			exp = val
		}
		e.Return(exp)
		return
	}

	pair := mustPair(exp)
	if IsSymbol(pair.car) {
		exp = pair.cdr // handle special form
		switch pair.car {
		case symDefun: // (defun f (x y) z)
			funName := car(exp)
			params := ListToSlice(cadr(exp))
			body := caddr(exp)
			compiled := CompileFunc(mustSymbol(funName).str, params, body)
			BindSymbolFunc(funName, compiled)
			e.Return(funName)
			return
		case symLambda: // (lambda x x)
			e.Return(makeProcedure(car(exp), cadr(exp), env))
			return
		case symFreeze: // (freeze body)
			e.Return(makeProcedure(Nil, car(exp), env))
			return
		case symLet: // (let x y z)
			args := evalExp(e, cadr(exp), env)
			newEnv := envExtend(env, []Obj{car(exp)}, []Obj{args})
			e.TailEval(caddr(exp), newEnv)
			return
		case symAnd:
			evalAnd(e, car(exp), cadr(exp), env)
			return
		case symOr:
			evalOr(e, car(exp), cadr(exp), env)
			return
		case symIf: // (if a b c)
			if listLength(pair.cdr) == 3 {
				evalIf(e, car(exp), cadr(exp), caddr(exp), env)
				return
			} // if may also be a function for partial apply
		case symCond: // (cond (false 1) (true 2))
			evalCond(e, exp, env)
			return
		case symType:
			// (type XXX (list Symbol)) => XXX, just ignore the second one
			e.TailEval(car(exp), env)
			return
		case symTrapError: // (trap-error ~body ~handler)
			evalTrapError(e, exp, env)
			return
		case symDo: // (do A A)
			evalExp(e, car(exp), env)
			e.TailEval(cadr(exp), env)
			return
		}
	}

	savePOS := e.pos
	fn := evalFunction(e, pair.car, env)
	e.data = e.data[:e.pos]
	e.data = append(e.data, fn)
	e.pos++

	args := pair.cdr
	for *args == scmHeadPair {
		v := evalExp(e, car(args), env)
		if *v == scmHeadError {
			e.pos = savePOS
			e.TailApply(fn, v)
			return
		}
		e.data = append(e.data, v)
		e.pos++
		args = cdr(args)
	}
	e.pos = savePOS
	e.kind = ControlFlowApply
}

func evalCond(e *ControlFlow, l Obj, env Obj) {
	for *l == scmHeadPair {
		curr := car(l)
		if evalExp(e, car(curr), env) == True {
			e.TailEval(cadr(curr), env)
			return
		}
		l = cdr(l)
	}
	e.Return(Nil)
}

func evalTrapError(e *ControlFlow, exp Obj, env Obj) {
	savePOS := e.pos
	fmark := e.frameMarkNow()
	defer func() {
		if err := recover(); err != nil {
			e.frameReset(fmark)
			if val, ok := err.(Obj); ok {
				if IsError(val) {
					e.pos = savePOS
					e.data = e.data[:e.pos]
					handle := evalFunction(e, cadr(exp), env)
					e.TailApply(handle, val)
					return
				}
			}
			// Unexpected panic?
			traceRecover("trap-error", exp, err)
			e.Return(MakeError("trap-error result is not Obj"))
			return
		}
	}()
	v := evalExp(e, car(exp), env)
	e.Return(v)
}

func evalFunction(e *ControlFlow, fn Obj, env Obj) Obj {
	if ok, sym := isSymbol(fn); ok {
		if sym.function != nil {
			return sym.function
		}

		if proc, ok := envGet(env, fn); ok {
			return proc
		}
	}

	switch *fn {
	case /* scmHeadPrimitive, */ scmHeadProcedure, scmHeadNative, scmHeadBytecodeFunc:
		return fn
	case scmHeadPair:
		return evalExp(e, fn, env)
	}
	panic(MakeError(fmt.Sprintf("can't apply non function: %#v", (*scmHead)(fn))))
}

func (e *ControlFlow) Global(key Obj) Obj {
	sym := mustSymbol(key)
	if sym.function != nil {
		return sym.function
	}
	errMsg := fmt.Sprintf("variable %s not bound", sym.str)
	panic(MakeError(errMsg))
}
