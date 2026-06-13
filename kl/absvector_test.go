// Resource-exhaustion robustness for primitives.
//
// `(absvector N)` allocates a slice of N pointers. With N attacker-controlled
// (trivially fuzzer-reachable as `(absvector 10000000000)`), the bare make()
// aborted the whole process with an unrecoverable Go "out of memory" fatal —
// the most uncatchable failure possible, killing any host embedding the
// evaluator. PrimAbsvector now rejects implausible sizes with a catchable
// Shen error; these tests pin that, and that ordinary sizes still work.
package kl

import (
	"fmt"
	"testing"
)

func TestAbsvectorRejectsHugeSizeCatchably(t *testing.T) {
	// A size past the ceiling must surface as a trappable error, not OOM the
	// process. We go well past maxAbsvectorSize so the test is robust to the
	// exact threshold.
	huge := int64(maxAbsvectorSize) + 1
	var ctx ControlFlow
	expr := fmt.Sprintf(`(trap-error (absvector %d) (lambda E (error-to-string E)))`, huge)
	got := ObjString(evalString(&ctx, expr))
	want := fmt.Sprintf(`"absvector size %d exceeds maximum %d"`, huge, maxAbsvectorSize)
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	// Evaluator stays usable afterwards.
	if cont := ObjString(evalString(&ctx, "(+ 40 2)")); cont != "42" {
		t.Fatalf("post-error eval: got %q, want 42", cont)
	}
}

func TestAbsvectorAllowsReasonableSize(t *testing.T) {
	var ctx ControlFlow
	// A small vector still builds and is addressable.
	got := ObjString(evalString(&ctx,
		`(do (set v (absvector 3)) (address-> (value v) 0 here))`))
	if got != "#vector" {
		t.Fatalf("got %q, want #vector", got)
	}
}
