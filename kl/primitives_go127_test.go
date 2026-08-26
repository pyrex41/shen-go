package kl

import "testing"

func TestPrimitiveRegistrarInfersArity(t *testing.T) {
	cases := []struct {
		name string
		want int
		prim Obj
	}{
		{"one", 1, primitiveRegistry.register("one", func(x Obj) Obj { return x })},
		{"two", 2, primitiveRegistry.register("two", func(x, _ Obj) Obj { return x })},
		{"three", 3, primitiveRegistry.register("three", func(x, _, _ Obj) Obj { return x })},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := MustNative(tc.prim).require; got != tc.want {
				t.Fatalf("inferred arity = %d, want %d", got, tc.want)
			}
		})
	}
}
