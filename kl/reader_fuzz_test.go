// Fuzz the reader+evaluator pipeline.
//
// The contract we're fuzzing: for ANY input string, the pipeline
//
//   NewSexpReader.Read -> Eval -> ObjString
//
// must terminate without an uncaught Go panic. Errors are fine — they
// just have to be Shen-level (a panic(Obj) caught by Eval's recover and
// returned as MakeError) or surfaced as a SexpReader read error. The
// failure modes we're hunting are the ones that put `goroutine ...
// [running]` on a user's stdout: bare-string panics that fall through
// the recovery layers (the original bug), or nil-deref segfaults in the
// must*/Is* family.
//
// Seeds are intentionally biased toward malformed Shen — they're the
// shape of input that surfaced the original crash (`(/. _ false)` from
// the Witness layout-proofs work) and the shape any future port glitch
// is most likely to take.
//
// Run the corpus seeds as ordinary unit tests:
//
//	go test ./kl/ -run FuzzReaderEval
//
// Run the fuzzer (CPU-bound, opt-in):
//
//	go test ./kl/ -run FuzzReaderEval -fuzz=FuzzReaderEval -fuzztime=30s
package kl

import (
	"strings"
	"testing"
)

func FuzzReaderEval(f *testing.F) {
	seeds := []string{
		// Golden path.
		"(+ 1 2)",
		"(do (defun id (X) X) (id 42))",
		`(let X 1 (+ X 1))`,
		"(if true 1 2)",

		// Error paths that MUST surface as catchable Shen errors.
		"(trap-error (overflow->str) (lambda E (error-to-string E)))",
		"(trap-error (value never-bound) (lambda E (error-to-string E)))",
		"(trap-error (if 42 1 2) (lambda E (error-to-string E)))",
		`(trap-error (simple-error "x") (lambda E (error-to-string E)))`,
		"(trap-error (42 1) (lambda E (error-to-string E)))",

		// Malformed-but-parseable inputs in the shape of the bug that
		// triggered this whole investigation (Witness proofs.shen):
		// `_` as a lambda parameter is illegal Shen, but the kernel must
		// not segfault on it.
		"(/. _ false)",
		"(lambda _ false)",

		// Reader edge cases that should at minimum not panic.
		"",
		" ",
		"()",
		"(",
		")",
		`"unterminated`,
		"#\\",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, input string) {
		// Eval has its own recover() — but we add a belt-and-braces
		// recover here so that *any* uncaught panic out of the reader,
		// Eval, or ObjString turns into a failed test rather than a
		// crashed worker. ObjString is the most likely deref site if
		// Eval ever returns nil (which it currently never does, but
		// future regressions might).
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("uncaught panic %T(%v) for input %q", r, r, input)
			}
		}()

		r := NewSexpReader(strings.NewReader(input), true)
		sexp, err := r.Read()
		if err != nil {
			// Read errors are an OK terminal state for malformed input.
			return
		}
		var ctx ControlFlow
		res := Eval(&ctx, sexp)
		if res == nil {
			t.Fatalf("Eval returned nil Obj for input %q (must be a real Obj or an error)", input)
		}
		// Renderability is part of the contract — ObjString must not
		// crash on whatever Eval handed back.
		_ = ObjString(res)
	})
}
