package kl

import "testing"

// Keys that are `=` must share one entry; keys that are not must not.
func TestShenXMapKeysFollowEqual(t *testing.T) {
	list := func(xs ...Obj) Obj {
		out := Nil
		for i := len(xs) - 1; i >= 0; i-- {
			out = Cons(xs[i], out)
		}
		return out
	}
	same := [][2]Obj{
		{MakeInteger(3), MakeNumber(3.0)},
		{MakeInteger(0), MakeNumber(-0.0)},
		{list(MakeInteger(1), MakeSymbol("a")), list(MakeNumber(1), MakeSymbol("a"))},
		{MakeString("x"), MakeString("x")},
	}
	for i, p := range same {
		if equal(p[0], p[1]) != True {
			t.Fatalf("case %d: fixture keys are not =", i)
		}
		h := primShenXMapNew(MakeInteger(4))
		primShenXMapPut(h, p[0], MakeInteger(7))
		if got := primShenXMapGet(h, p[1]); got != MakeInteger(7) {
			t.Fatalf("case %d: = keys missed: %s", i, ObjString(got))
		}
	}
	differ := [][2]Obj{
		{MakeNumber(1.5), MakeInteger(1)},
		{MakeString("a"), MakeSymbol("a")},
		{list(MakeInteger(1), MakeInteger(2)), list(MakeInteger(12))},
		{list(), MakeString("")},
	}
	for i, p := range differ {
		if equal(p[0], p[1]) == True {
			t.Fatalf("case %d: fixture keys are =", i)
		}
		h := primShenXMapNew(MakeInteger(4))
		primShenXMapPut(h, p[0], MakeInteger(7))
		if got := primShenXMapGet(h, p[1]); got != shenxNone {
			t.Fatalf("case %d: distinct keys collided", i)
		}
	}
}
