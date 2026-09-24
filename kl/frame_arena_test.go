package kl

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// Tests for the VM frame arena (issue #50). Before the arena, activation slabs
// came from a freelist capped at 128 entries, so every frame of a non-tail
// recursion deeper than 128 allocated a fresh slab and every return past the
// cap scanned the whole pool before dropping the slab to the GC.

// TestDeepRecursionDoesNotAllocateFrames: a deep non-tail recursion should
// allocate its N conses and (almost) nothing else once the arena is warm.
// On the bounded freelist (mk 2000) allocated 3873 objects.
func TestDeepRecursionDoesNotAllocateFrames(t *testing.T) {
	for _, tc := range []struct {
		depth int
		max   float64
	}{
		{2000, 2100},
		// Depth 20000 crosses several arena blocks and exercises the
		// retention path: the second descent must reuse the retained
		// blocks rather than reallocate them.
		{20000, 20100},
	} {
		var ctx ControlFlow
		mustDefun(t, &ctx, defMk)
		call := mustReadOneT(t, "(mk "+strconv.Itoa(tc.depth)+")")
		if res := Eval(&ctx, call); IsError(res) { // warm the arena
			t.Fatal(ObjString(res))
		}
		allocs := testing.AllocsPerRun(5, func() { Eval(&ctx, call) })
		t.Logf("(mk %d): %.0f allocs/run", tc.depth, allocs)
		if allocs > tc.max {
			t.Errorf("(mk %d) allocated %.0f objects; want about %d conses", tc.depth, allocs, tc.depth)
		}
	}
}

// TestShallowFrameHeadroomDoesNotAllocate: a frame whose operand stack gets
// deeper than 16 but shallower than 46 slots must still fit in its slab. The
// old freelist rounded every slab up to 48 slots; the arena must keep that
// headroom or `append` reallocates the operand stack off-slab on exactly the
// shallow kernel/typechecker workload that motivated frame recycling.
func TestShallowFrameHeadroomDoesNotAllocate(t *testing.T) {
	// (+ A (+ A ... (+ A A))) nested 30 deep: ~31 operands live at the
	// innermost point, Nlocals = 1.
	const depth = 30
	expr := "A"
	for i := 0; i < depth; i++ {
		expr = "(+ A " + expr + ")"
	}
	var ctx ControlFlow
	if res := evalString(&ctx, "(defun deep-expr (A) "+expr+")"); IsError(res) {
		t.Fatal(ObjString(res))
	}
	call := mustReadOneT(t, `(deep-expr 1)`)
	if res := Eval(&ctx, call); res != MakeInteger(depth+1) {
		t.Fatalf("(deep-expr 1) => %s, want %d", ObjString(res), depth+1)
	}
	if allocs := testing.AllocsPerRun(10, func() { Eval(&ctx, call) }); allocs != 0 {
		t.Fatalf("(deep-expr 1) allocated %.0f objects; want 0 (operand stack should fit the slab)", allocs)
	}
}

// TestFrameArenaUnwindsAfterCaughtError: frames abandoned by a raised error
// must not leak arena space, whether the error is caught by a compiled
// trap-error (Try) or a top-level interpreted one (evalTrapError). Depth 700
// spans ~34k slots, so the abandoned frames cross block boundaries; the block
// count must also stay constant, i.e. frameReset must not drop and re-grow
// blocks on every caught error.
func TestFrameArenaUnwindsAfterCaughtError(t *testing.T) {
	var ctx ControlFlow
	mustDefun(t, &ctx, defDeepBoom)
	mustDefun(t, &ctx, `(defun catch-deep (N) (trap-error (deep-boom N) (lambda E caught)))`)
	calls := []string{`(catch-deep 700)`, `(trap-error (deep-boom 700) (lambda E caught))`}
	// First round grows the arena; the count must then hold steady.
	if res := evalString(&ctx, calls[0]); ObjString(res) != "caught" {
		t.Fatalf("%s => %s", calls[0], ObjString(res))
	}
	nblocks := len(ctx.frameBlocks)
	if nblocks < 2 {
		t.Fatalf("expected depth 700 to span more than one arena block, got %d", nblocks)
	}
	for i := 0; i < 20; i++ {
		for _, call := range calls {
			if res := evalString(&ctx, call); ObjString(res) != "caught" {
				t.Fatalf("%s => %s", call, ObjString(res))
			}
			if ctx.frameCur != 0 || ctx.frameTop != 0 {
				t.Fatalf("after %s: arena not unwound: cur=%d top=%d", call, ctx.frameCur, ctx.frameTop)
			}
			if n := len(ctx.frameBlocks); n != nblocks {
				t.Fatalf("after %s (round %d): block count changed %d -> %d", call, i, nblocks, n)
			}
		}
	}
}

// TestFrameArenaRetentionBounded: after a one-shot deep recursion the arena
// keeps at most frameRetainSlots spare slots (plus the block in use). This is
// the property the old 128-entry cap existed for.
func TestFrameArenaRetentionBounded(t *testing.T) {
	var ctx ControlFlow
	mustDefun(t, &ctx, defMk)
	if res := evalString(&ctx, `(mk 60000)`); IsError(res) {
		t.Fatal(ObjString(res))
	}
	if ctx.frameCur != 0 || ctx.frameTop != 0 {
		t.Fatalf("arena not unwound: cur=%d top=%d", ctx.frameCur, ctx.frameTop)
	}
	total := 0
	for _, b := range ctx.frameBlocks[1:] {
		total += len(b)
	}
	if total > frameRetainSlots {
		t.Fatalf("retained %d spare slots > %d", total, frameRetainSlots)
	}
}

const (
	defMk       = `(defun mk (N) (if (= N 0) () (cons N (mk (- N 1)))))`
	defDeepBoom = `(defun deep-boom (N) (if (= N 0) (simple-error "boom") (+ 1 (deep-boom (- N 1)))))`
)

// TestRepeatedDeepRecursionKeepsBlocks: a loop of (mk 11000) inside one
// enclosing bytecode frame must allocate its conses and nothing else per
// iteration. Depth 11000 reaches block 10 (8 MiB), which on its own exceeds
// frameRetainSlots; before the fix every return past depth ~10,900 dropped
// that block and the next descent re-made and re-zeroed it (measured: 8.7 MB
// and one extra alloc per (mk 11000), a 2x cliff versus (mk 10000)). Spare
// blocks are now trimmed only at a top-level return, a recover site or the
// periodic idle check (frameTrimInterval), never on an ordinary return, so
// the only permitted extra allocation per run is block 10 coming back once
// after the trim at the end of the previous Eval. The iteration count makes
// the loop cross more than frameTrimInterval block boundaries, so the
// periodic idle trim runs at least once and must keep the block the loop
// enters on every descent.
func TestRepeatedDeepRecursionKeepsBlocks(t *testing.T) {
	const depth, iters = 11000, 120
	var ctx ControlFlow
	mustDefun(t, &ctx, defMk)
	mustDefun(t, &ctx, `(defun mk-loop (K N) (if (= K 0) done (do (mk N) (mk-loop (- K 1) N))))`)
	call := mustReadOneT(t, fmt.Sprintf("(mk-loop %d %d)", iters, depth))
	if res := Eval(&ctx, call); ObjString(res) != "done" { // warm the arena
		t.Fatalf("(mk-loop %d %d) => %s", iters, depth, ObjString(res))
	}
	allocs := testing.AllocsPerRun(2, func() { Eval(&ctx, call) })
	conses := iters * depth
	t.Logf("%d x (mk %d): %.0f allocs/run (%d conses)", iters, depth, allocs, conses)
	if allocs > float64(conses)+2 {
		t.Errorf("%d x (mk %d) allocated %.0f objects; want at most %d conses + 2, not one arena block per iteration",
			iters, depth, allocs, conses)
	}
}

// TestTryResetsFrameArena pins the frameReset in Try. Called directly from
// Go there is no enclosing bytecode frame whose putFrame could self-heal the
// arena, so a thunk that recurses 700 deep and raises must leave the arena
// fully unwound the moment Try returns. Without the reset in Try this fails
// with the abandoned frames still claimed (cur=6, top>0).
func TestTryResetsFrameArena(t *testing.T) {
	var ctx ControlFlow
	mustDefun(t, &ctx, defDeepBoom)
	mustDefun(t, &ctx, `(defun boom-thunk () (deep-boom 700))`)
	res := Try(&ctx, PrimFunc(MakeSymbol("boom-thunk")))
	if !IsError(res.data) || !strings.Contains(ObjString(res.data), "boom") {
		t.Fatalf("Try(boom-thunk) => %s, want the raised boom error", ObjString(res.data))
	}
	if ctx.frameCur != 0 || ctx.frameTop != 0 {
		t.Fatalf("after Try: arena not unwound: cur=%d top=%d", ctx.frameCur, ctx.frameTop)
	}
}

// TestFrameArenaIdleTrimInsideLongLivedFrame: a frame that never returns (a
// script's load loop, a server loop) must not pin the peak of a one-shot
// deep recursion for its whole lifetime. Inside one enclosing frame: (mk
// 60000) once (blocks up to 12, 4.2M slots), then a shallow loop whose
// (mk 100) calls cross three block boundaries each, more than
// frameTrimInterval crossings in all. A Go probe called from inside the
// frame records the spare slots right after the deep call (still all
// retained: no trim point has been crossed) and after the shallow loop (cut
// back to the budget by the periodic idle trim, while the blocks the shallow
// loop uses are of course kept). The loop spans more than two intervals: the
// first decision after the deep call still counts its blocks as entered
// since the previous trim and keeps them; the second releases them.
func TestFrameArenaIdleTrimInsideLongLivedFrame(t *testing.T) {
	const shallowIters = 1000 // 3 crossings each: 3000 > 2*frameTrimInterval
	var ctx ControlFlow
	var readings []int
	probe := MakeSymbol("frame-arena-probe")
	BindSymbolFunc(probe, MakeNative(func(e *ControlFlow) {
		spare := 0
		for _, b := range e.frameBlocks[e.frameCur+1:] {
			spare += len(b)
		}
		readings = append(readings, spare)
		e.Return(Nil)
	}, 0))
	defer BindSymbolFunc(probe, nil)
	mustDefun(t, &ctx, defMk)
	mustDefun(t, &ctx, `(defun shallow-loop (K) (if (= K 0) done (do (mk 100) (shallow-loop (- K 1)))))`)
	mustDefun(t, &ctx, fmt.Sprintf(`(defun long-lived () (do (mk 60000) (do (frame-arena-probe) (do (shallow-loop %d) (frame-arena-probe)))))`, shallowIters))
	if res := evalString(&ctx, `(long-lived)`); IsError(res) {
		t.Fatal(ObjString(res))
	}
	if len(readings) != 2 {
		t.Fatalf("probe ran %d times, want 2", len(readings))
	}
	t.Logf("spare slots after (mk 60000): %d; after %d x (mk 100): %d (budget %d)", readings[0], shallowIters, readings[1], frameRetainSlots)
	if readings[0] <= frameRetainSlots {
		t.Errorf("after (mk 60000) inside the frame: %d spare slots, expected the peak to still be retained (> %d)", readings[0], frameRetainSlots)
	}
	if readings[1] > frameRetainSlots {
		t.Errorf("after the shallow loop: %d spare slots retained > %d; the idle trim did not run", readings[1], frameRetainSlots)
	}
}

func mustDefun(t *testing.T, ctx *ControlFlow, src string) {
	t.Helper()
	if res := evalString(ctx, src); IsError(res) {
		t.Fatalf("%s: %s", src, ObjString(res))
	}
}

func mustReadOneT(t *testing.T, src string) Obj {
	t.Helper()
	o, err := NewSexpReader(strings.NewReader(src), true).Read()
	if err != nil {
		t.Fatalf("read %q: %v", src, err)
	}
	return o
}

// BenchmarkVMDeepRecursion: 2000-deep non-tail recursion (issue #50).
func BenchmarkVMDeepRecursion(b *testing.B) {
	benchEval(b, `(defun mk (N) (if (= N 0) () (cons N (mk (- N 1)))))`, `(mk 2000)`)
}
