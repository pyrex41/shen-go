package codegen

import (
	"bytes"
	"strings"
	"testing"
)

func TestSealedPrimitiveCallHasNoDynamicFallback(t *testing.T) {
	form := readCodegenForm(t, "(call ($global cons) ($const 1) ($const 2))")
	var open, sealed bytes.Buffer
	if err := New().generateExpr(&open, form); err != nil {
		t.Fatal(err)
	}
	if err := NewSealed().generateExpr(&sealed, form); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(open.String(), "HasCanonicalPrimitiveBinding") {
		t.Fatalf("open codegen should guard cons:\n%s", open.String())
	}
	if !strings.Contains(open.String(), "Call(__e, PrimFunc") {
		t.Fatalf("open codegen should fall back to PrimFunc:\n%s", open.String())
	}
	s := sealed.String()
	if strings.Contains(s, "HasCanonicalPrimitiveBinding") || strings.Contains(s, "PrimFunc") {
		t.Fatalf("sealed cons should be a direct PrimCons with no guard:\n%s", s)
	}
	if !strings.Contains(s, "PrimCons(") {
		t.Fatalf("sealed cons missing PrimCons:\n%s", s)
	}
}

func TestSealedPlusIsDirectPrimNotCall(t *testing.T) {
	form := readCodegenForm(t, "(call ($global +) ($const 1) ($const 2))")
	var got bytes.Buffer
	if err := NewSealed().generateExpr(&got, form); err != nil {
		t.Fatal(err)
	}
	s := got.String()
	if strings.Contains(s, "HasCanonicalPrimitiveBinding") || strings.Contains(s, "PrimFunc") {
		t.Fatalf("sealed + still has dynamic dispatch:\n%s", s)
	}
	if !strings.Contains(s, "PrimNumberAdd(") {
		t.Fatalf("sealed + missing PrimNumberAdd:\n%s", s)
	}
}
