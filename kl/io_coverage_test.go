package kl

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// TestFileReadPrimitives covers read-file-as-string and read-file-as-bytelist
// against a real temp file. These primitives use the raw path (no
// *home-directory* prefix), so an absolute path works directly.
func TestFileReadPrimitives(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "data.txt")
	if err := os.WriteFile(p, []byte("AB"), 0644); err != nil {
		t.Fatal(err)
	}

	var ctx ControlFlow
	if got := ObjString(evalString(&ctx, fmt.Sprintf(`(read-file-as-string "%s")`, p))); got != `"AB"` {
		t.Errorf("read-file-as-string: got %q, want %q", got, `"AB"`)
	}
	// 'A'=65, 'B'=66
	if got := ObjString(evalString(&ctx, fmt.Sprintf(`(read-file-as-bytelist "%s")`, p))); got != "(65 66)" {
		t.Errorf("read-file-as-bytelist: got %q, want %q", got, "(65 66)")
	}
}

// TestStreamRoundTrip covers open (in/out), write-byte, read-byte and close by
// writing bytes to a file and reading them back. *home-directory* is "" in the
// kl package, so the absolute path passes through open unchanged.
func TestStreamRoundTrip(t *testing.T) {
	p := filepath.Join(t.TempDir(), "stream.bin")

	var ctl ControlFlow
	open := ctl.Global(MakeSymbol("open"))
	closeFn := ctl.Global(MakeSymbol("close"))
	writeByte := ctl.Global(MakeSymbol("write-byte"))
	readByte := ctl.Global(MakeSymbol("read-byte"))

	out := Call(&ctl, open, MakeString(p), MakeSymbol("out"))
	for _, b := range []int{72, 105} { // "Hi"
		Call(&ctl, writeByte, MakeInteger(b), out)
	}
	Call(&ctl, closeFn, out)

	in := Call(&ctl, open, MakeString(p), MakeSymbol("in"))
	for _, want := range []int{72, 105, -1} { // -1 = EOF
		got := Call(&ctl, readByte, in)
		if GetInteger(got) != want {
			t.Errorf("read-byte: got %d, want %d", GetInteger(got), want)
		}
	}
	Call(&ctl, closeFn, in)
}

// TestGetTime covers both arms of get-time (unix wall clock and run-time).
func TestGetTime(t *testing.T) {
	var ctl ControlFlow
	fn := ctl.Global(MakeSymbol("get-time"))
	for _, kind := range []string{"unix", "run"} {
		got := Call(&ctl, fn, MakeSymbol(kind))
		if !IsNumber(got) {
			t.Errorf("get-time %s: expected a number, got %q", kind, ObjString(got))
		}
	}
}

// TestEvalKL covers eval-kl: its argument is first evaluated to a KL form, which
// eval-kl then evaluates. Here (cons + (cons 3 (cons 4 ()))) builds (+ 3 4).
func TestEvalKL(t *testing.T) {
	var ctx ControlFlow
	got := ObjString(evalString(&ctx, "(eval-kl (cons + (cons 3 (cons 4 ()))))"))
	if got != "7" {
		t.Errorf("eval-kl: got %q, want 7", got)
	}
}

// TestLoadFile covers load-file (and loadFile): loading a .kl file evaluates each
// top-level form and returns the filename on success.
func TestLoadFile(t *testing.T) {
	p := filepath.Join(t.TempDir(), "snippet.kl")
	if err := os.WriteFile(p, []byte("(set loaded-marker 123)\n"), 0644); err != nil {
		t.Fatal(err)
	}

	var ctx ControlFlow
	got := evalString(&ctx, fmt.Sprintf(`(load-file "%s")`, p))
	if IsError(got) {
		t.Fatalf("load-file errored: %q", ObjString(got))
	}
	// The form in the file ran, so the side effect is visible.
	if v := ObjString(evalString(&ctx, "(value loaded-marker)")); v != "123" {
		t.Errorf("expected loaded side effect, got %q", v)
	}
}
