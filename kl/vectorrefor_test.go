package kl

import "testing"

// TestVectorRefOrTreatsOnlyShenFailAsAbsent pins the presence check behind
// issue #46's fix, as issue #54 settled it. The VM compiler lowers
// (trap-error (<-vector V N) (lambda E D)) to the _kl.<-vector/or check. An
// unassigned slot holds what (fail) returned when the vector was built, and
// after #54 that is the symbol shen.fail! in every world: the native vector,
// the kernel's own KL vector (cmd/kl, the equivalence harness) and the
// compiled kernel. The symbol `...` is only how the printer shows the fail
// value; as a slot value it is present, as on shen-cl and shen-scheme. The
// check used to accept `...` as a second filler, which misread that value
// as absent.
func TestVectorRefOrTreatsOnlyShenFailAsAbsent(t *testing.T) {
	var e ControlFlow
	eval := func(src string) Obj {
		form, err := ReadForm(src)
		if err != nil {
			t.Fatal(err)
		}
		res := Eval(&e, form)
		if IsError(res) {
			t.Fatalf("%s: %s", src, GetString(PrimErrorToString(res)))
		}
		return res
	}
	// defun compiles the body, so the trap-error becomes _kl.<-vector/or.
	eval(`(defun t.slot-or (V N) (trap-error (<-vector V N) (lambda E absent)))`)
	eval(`(set t.v (absvector 4))`)
	eval(`(address-> (value t.v) 1 shen.fail!)`)
	eval(`(address-> (value t.v) 2 (intern "..."))`)
	eval(`(address-> (value t.v) 3 present)`)
	for _, tc := range []struct{ n, want string }{
		{"0", "absent"}, // slot 0 is the limit, never a value
		{"1", "absent"}, // the fail filler, in every world
		{"2", "..."},    // an ordinary symbol
		{"3", "present"},
		{"4", "absent"}, // out of range
	} {
		got := ObjString(eval("(t.slot-or (value t.v) " + tc.n + ")"))
		if got != tc.want {
			t.Errorf("(t.slot-or V %s) = %s, want %s", tc.n, got, tc.want)
		}
	}
}
