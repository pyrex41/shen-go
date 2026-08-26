package kl

import (
	"math"
	"strings"
	"testing"
)

func TestTypedObjectKindAndScalars(t *testing.T) {
	cases := []struct {
		o    Obj
		kind ValueKind
	}{{MakeInteger(2), KindFixnum}, {MakeNumber(2.5), KindNumber}, {True, KindBoolean}, {Nil, KindNil}, {MakeString("x"), KindString}, {cons(True, Nil), KindPair}, {MakeVector(1), KindVector}, {nil, KindUnknown}}
	for _, tc := range cases {
		if got := TypedObjectKind(tc.o); got != tc.kind {
			t.Fatalf("kind = %v, want %v", got, tc.kind)
		}
	}
	if n, ok := TypedFloat64(MakeNumber(2.5)); !ok || n != 2.5 {
		t.Fatalf("number extraction = %v, %v", n, ok)
	}
	if _, ok := TypedFloat64(True); ok {
		t.Fatal("boolean extracted as number")
	}
	if b, ok := TypedBoolean(True); !ok || !b {
		t.Fatalf("boolean extraction = %v, %v", b, ok)
	}
	if s, ok := TypedString(MakeString("x")); !ok || s != "x" {
		t.Fatalf("string extraction = %q, %v", s, ok)
	}
}

func TestTypedDividePreservesIEEEAndGuardsZero(t *testing.T) {
	if got := GetNumber(TypedDivide(3, 2)); got != 1.5 {
		t.Fatalf("divide = %v", got)
	}
	if !math.IsInf(GetNumber(TypedDivide(math.Inf(1), 2)), 1) {
		t.Fatal("infinity was not preserved")
	}
	assertTypedPanic(t, "division by zero", func() { TypedDivide(1, 0) })
}

func TestTypedUnicodeStringOperations(t *testing.T) {
	s := MakeString("é🙂x")
	if got := GetString(TypedStringIndex(s, 1)); got != "🙂" {
		t.Fatalf("index = %q", got)
	}
	if got := GetString(TypedStringTail(s)); got != "🙂x" {
		t.Fatalf("tail = %q", got)
	}
	assertTypedPanic(t, "empty string", func() { TypedStringTail(MakeString("")) })
	assertTypedPanic(t, "valid index", func() { TypedStringIndex(s, -1) })
}

func TestTypedPairVectorAccessPreservesAlias(t *testing.T) {
	p := cons(True, Nil)
	if h, ok := TypedPairHead(p); !ok || h != True {
		t.Fatal("pair head mismatch")
	}
	if _, ok := TypedPairHead(True); ok {
		t.Fatal("non-pair passed pair guard")
	}
	v := MakeVector(1)
	TypedVectorSet(v, 0, p)
	if got := TypedVectorGet(v, 0); got != p {
		t.Fatal("vector alias was not preserved")
	}
	if got := TypedVectorGet(v, 0); got != p {
		t.Fatal("vector read changed value")
	}
	assertTypedPanic(t, "out of range", func() { TypedVectorGet(v, 2) })
}

func assertTypedPanic(t *testing.T, want string, fn func()) {
	t.Helper()
	defer func() {
		r := recover()
		if r == nil {
			t.Fatalf("expected panic containing %q", want)
		}
		o, ok := r.(Obj)
		if !ok || !IsError(o) || !containsError(o, want) {
			t.Fatalf("panic = %#v, want Shen error containing %q", r, want)
		}
	}()
	fn()
}

func containsError(o Obj, want string) bool {
	return len(want) == 0 || strings.Contains(mustError(o).err, want)
}
