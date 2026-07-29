// D4: narrowing a number to a Go int must not escape the Shen error system.
//
// mustInteger/GetInteger converted a float64 with a bare int(f). Go leaves an
// out-of-range conversion implementation-defined (arm64 saturates), so every
// primitive that takes an index or a byte received 9223372036854775807 for
// +Inf. The consequences ranged from bad to unrecoverable:
//
//   - (address-> (absvector 3) +inf 1) had no bounds check at all and died
//     with a Go runtime panic that trap-error could not catch -- it came back
//     as "#<err trap-error result is not Obj>".
//   - (n->string +inf) silently emitted a replacement byte instead of
//     erroring, against the reference contract in Tarver's
//     Primitives/CL/n-to-string.lsp ("~A is not a natural number").
//   - pos and <-address did raise catchable errors, but embedded the nonsense
//     saturated index in the message.
//
// Negative indices had the same hole from the other side: they reached a Go
// slice index and panicked outside the error system.
package kl

import (
	"strings"
	"testing"
)

// klInfExpr builds +Inf arithmetically, so the test does not depend on how the
// reader spells infinity. Division by zero raises here (issue #10), hence the
// multiplication of two large reciprocals.
const klInfExpr = `(* (/ 1.0 1e-300) (/ 1.0 1e-300))`

// trapped evaluates expr inside trap-error and returns the error string, or
// the ordinary result if nothing was raised.
func trapped(ctx *ControlFlow, expr string) string {
	return ObjString(evalString(ctx,
		`(trap-error `+expr+` (lambda E (error-to-string E)))`))
}

// Every primitive that narrows a number must raise a Shen error that
// trap-error can actually catch, and must not name a saturated index.
func TestUnnarrowableIndexRaisesCatchableError(t *testing.T) {
	var ctx ControlFlow

	for _, tc := range []struct{ name, expr, want string }{
		{
			"n->string on +inf",
			`(n->string ` + klInfExpr + `)`,
			`"inf is not a natural number"`,
		},
		{
			"pos on +inf",
			`(pos "abc" ` + klInfExpr + `)`,
			`"inf is not a valid integer"`,
		},
		{
			"write-byte on +inf",
			`(write-byte ` + klInfExpr + ` (value *stoutput*))`,
			`"inf is not a valid integer"`,
		},
		{
			"<-address on +inf",
			`(<-address (absvector 3) ` + klInfExpr + `)`,
			`"inf is not a valid integer"`,
		},
		{
			"address-> on +inf",
			`(address-> (absvector 3) ` + klInfExpr + ` 1)`,
			`"inf is not a valid integer"`,
		},
		{
			"address-> on a huge finite float",
			`(address-> (absvector 3) 1e19 1)`,
			`"1e+19 is not a valid integer"`,
		},
		{
			"n->string on a huge finite float",
			`(n->string 1e19)`,
			`"1e+19 is not a natural number"`,
		},
	} {
		got := trapped(&ctx, tc.expr)
		if got != tc.want {
			t.Errorf("%s: got %s, want %s", tc.name, got, tc.want)
		}
	}

	// The evaluator is still usable afterwards.
	if got := ObjString(evalString(&ctx, "(+ 40 2)")); got != "42" {
		t.Fatalf("post-error eval: got %q, want 42", got)
	}
}

// The saturated maxint must not appear in any message: it is not a value the
// caller ever supplied.
func TestNarrowingErrorsDoNotNameMaxint(t *testing.T) {
	var ctx ControlFlow

	for _, expr := range []string{
		`(n->string ` + klInfExpr + `)`,
		`(pos "abc" ` + klInfExpr + `)`,
		`(write-byte ` + klInfExpr + ` (value *stoutput*))`,
		`(<-address (absvector 3) ` + klInfExpr + `)`,
		`(address-> (absvector 3) ` + klInfExpr + ` 1)`,
	} {
		got := trapped(&ctx, expr)
		for _, bad := range []string{"9223372036854775807", "-9223372036854775808", "not Obj"} {
			if strings.Contains(got, bad) {
				t.Errorf("%s: message %s contains %q", expr, got, bad)
			}
		}
	}
}

// Negative indices used to reach a Go slice index directly. address-> had no
// bounds check whatsoever, in either direction.
func TestNegativeAndOutOfRangeIndicesAreCatchable(t *testing.T) {
	var ctx ControlFlow

	for _, tc := range []struct{ name, expr, want string }{
		{"pos with a negative index", `(pos "abc" -1)`, `"-1 is not valid index for abc"`},
		{"<-address with a negative index", `(<-address (absvector 3) -1)`, `"index -1 out of range 3"`},
		{"address-> with a negative index", `(address-> (absvector 3) -1 1)`, `"index -1 out of range 3"`},
		{"address-> past the end", `(address-> (absvector 3) 99 1)`, `"index 99 out of range 3"`},
		{"n->string with a negative code", `(n->string -1)`, `"-1 is not a natural number"`},
	} {
		if got := trapped(&ctx, tc.expr); got != tc.want {
			t.Errorf("%s: got %s, want %s", tc.name, got, tc.want)
		}
	}
}

// In-range uses are unaffected.
func TestInRangeIndexingStillWorks(t *testing.T) {
	var ctx ControlFlow

	for _, tc := range []struct{ expr, want string }{
		{`(n->string 65)`, `"A"`},
		{`(pos "abc" 1)`, `"b"`},
		{`(<-address (address-> (absvector 3) 2 7) 2)`, `7`},
		{`(<-address (address-> (absvector 3) 0 (cons 1 2)) 0)`, `(1 . 2)`},
	} {
		if got := ObjString(evalString(&ctx, tc.expr)); got != tc.want {
			t.Errorf("%s: got %s, want %s", tc.expr, got, tc.want)
		}
	}
}
