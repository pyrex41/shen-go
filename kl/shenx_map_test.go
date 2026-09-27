package kl

import (
	"math"
	"testing"
)

// TestShenValueHashAgreesWithEqual: the map finds keys by hash, then =, so
// any two values that are = must hash alike, including numbers written
// differently and structures built separately.
func TestShenValueHashAgreesWithEqual(t *testing.T) {
	var ctx ControlFlow
	pairs := [][2]string{
		{`1`, `1.0`},
		{`0`, `-0.0`},
		{`(cons 1 (cons "a" (cons b ())))`, `(cons 1.0 (cons "a" (cons b ())))`},
		{`(cons (cons 1 ()) (cons (cons 2 ()) ()))`, `(cons (cons 1 ()) (cons (cons 2 ()) ()))`},
		{`(let V (absvector 2) (do (address-> V 0 x) (do (address-> V 1 (cons 3 ())) V)))`,
			`(let V (absvector 2) (do (address-> V 0 x) (do (address-> V 1 (cons 3 ())) V)))`},
		{`(* 1000000 1000000)`, `1000000000000`},
		{`"λ"`, `(cn "" "λ")`},
	}
	for _, p := range pairs {
		a, b := evalString(&ctx, p[0]), evalString(&ctx, p[1])
		if equal(a, b) != True {
			t.Fatalf("test premise: %s and %s are not =", p[0], p[1])
		}
		if shenValueHash(a) != shenValueHash(b) {
			t.Errorf("%s and %s are = but hash differently", p[0], p[1])
		}
	}
	distinct := []string{`1`, `2`, `"1"`, `x`, `()`, `(cons 1 (cons 2 ()))`, `(cons 2 (cons 1 ()))`, `(cons (cons 1 (cons 2 ())) ())`, `true`}
	seen := map[uint64]string{}
	for _, s := range distinct {
		h := shenValueHash(evalString(&ctx, s))
		if prev, ok := seen[h]; ok {
			t.Errorf("%s and %s collide", prev, s)
		}
		seen[h] = s
	}
	if shenValueHash(MakeNumber(math.NaN())) == 0 {
		t.Log("NaN hashes") // must not crash; NaN is never = anything
	}
}

func TestShenXMapOperations(t *testing.T) {
	var ctx ControlFlow
	InstallShenX()
	prog := `(let M (shen.x.map-new)
	  (do (shen.x.map-put M (cons 1 (cons 2 ())) listval)
	  (do (shen.x.map-put M 3 three)
	  (do (shen.x.map-put M 3.0 three-again)
	  (do (shen.x.map-put M "s" str)
	  (do (shen.x.map-remove M "s")
	  (do (shen.x.map-remove M "absent")
	    (cons (shen.x.map-get M (cons 1 (cons 2 ())) none)
	    (cons (shen.x.map-get M 3 none)
	    (cons (shen.x.map-get M "s" none)
	    (cons (shen.x.map-has? M 3)
	    (cons (shen.x.map-count M) ())))))))))))))`
	got := ObjString(evalString(&ctx, prog))
	want := ObjString(evalString(&ctx, `(cons listval (cons three-again (cons none (cons true (cons 2 ())))))`))
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if r := evalString(&ctx, `(shen.x.map-get 5 k d)`); !IsError(r) {
		t.Fatalf("map-get on a non-map returned %s, want an error", ObjString(r))
	}
}

// TestShenXMapManyKeys drives enough keys through the map to exercise
// buckets with more than one entry and removal from the middle of one.
func TestShenXMapManyKeys(t *testing.T) {
	m := makeMap()
	key := func(i int) Obj { return cons(MakeInteger(i%97), cons(MakeInteger(i/97), Nil)) }
	const n = 20000
	for i := 0; i < n; i++ {
		primMapPut(m, key(i), MakeInteger(i))
	}
	for i := 0; i < n; i += 2 {
		primMapRemove(m, key(i))
	}
	if c := primMapCount(m); c != MakeInteger(n/2) {
		t.Fatalf("count %s, want %d", ObjString(c), n/2)
	}
	for i := 0; i < n; i++ {
		got := primMapGet(m, key(i), False)
		if i%2 == 0 && got != False || i%2 == 1 && got != MakeInteger(i) {
			t.Fatalf("key %d => %s", i, ObjString(got))
		}
	}
	if l := listLength(primMapKeys(m)); l != n/2 {
		t.Fatalf("keys: %d, want %d", l, n/2)
	}
}
