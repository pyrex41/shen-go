package kl

import "testing"

// TestVectorRefOrAcceptsBothFailFillers pins the fix for issue #46's actual
// cause. The VM compiler lowers (trap-error (<-vector V N) (lambda E D)) to
// the _kl.<-vector/or presence check. An unassigned slot holds whatever (fail)
// returned when the vector was built: `...` from the native vector, shen.fail!
// from the kernel's own KL vector. The check used to know only `...`, so a
// kernel booted without InstallKernelFast (cmd/kl, the equivalence harness)
// got shen.fail! back as a present value from every fresh property-vector
// slot, and put died in shen.change-pointer-value.
func TestVectorRefOrAcceptsBothFailFillers(t *testing.T) {
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
		{"1", "absent"}, // KL kernel's filler
		{"2", "absent"}, // native filler
		{"3", "present"},
		{"4", "absent"}, // out of range
	} {
		got := ObjString(eval("(t.slot-or (value t.v) " + tc.n + ")"))
		if got != tc.want {
			t.Errorf("(t.slot-or V %s) = %s, want %s", tc.n, got, tc.want)
		}
	}
}
