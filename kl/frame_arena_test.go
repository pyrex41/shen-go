package kl

import (
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
		evalString(&ctx, `(defun mk (N) (if (= N 0) () (cons N (mk (- N 1)))))`)
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
	evalString(&ctx, `(defun deep-boom (N) (if (= N 0) (simple-error "boom") (+ 1 (deep-boom (- N 1)))))`)
	evalString(&ctx, `(defun catch-deep (N) (trap-error (deep-boom N) (lambda E caught)))`)
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
	evalString(&ctx, `(defun mk (N) (if (= N 0) () (cons N (mk (- N 1)))))`)
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
