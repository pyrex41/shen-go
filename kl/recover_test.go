package kl

import (
	"io"
	"os"
	"strings"
	"testing"
)

// captureFile swaps *fp (os.Stdout or os.Stderr) for the write end of a pipe
// while f runs, draining the read end in a goroutine so a writer that fills
// the pipe buffer cannot deadlock the test. It returns everything written.
func captureFile(t *testing.T, fp **os.File, f func()) string {
	t.Helper()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	saved := *fp
	*fp = w
	done := make(chan string)
	go func() {
		b, _ := io.ReadAll(r)
		r.Close()
		done <- string(b)
	}()
	func() {
		defer func() {
			*fp = saved
			w.Close()
		}()
		f()
	}()
	return <-done
}

// captureStdoutStderr runs f and returns what it wrote to os.Stdout and to
// os.Stderr. Tests using it must not run in parallel: it swaps package-level
// globals.
func captureStdoutStderr(t *testing.T, f func()) (stdout, stderr string) {
	t.Helper()
	stdout = captureFile(t, &os.Stdout, func() {
		stderr = captureFile(t, &os.Stderr, f)
	})
	return stdout, stderr
}

// bindGoPanicNative binds test-go-panic to a native that raises a bare Go
// string panic, the shape of a runtime bug inside a primitive rather than a
// KL error.
func bindGoPanicNative() {
	BindSymbolFunc(MakeSymbol("test-go-panic"), MakeNative(func(*ControlFlow) {
		panic("boom from test-go-panic")
	}, 0))
}

// TestRecoveredKLErrorsKeepStdoutClean is issue #52: a raised KL error is
// ordinary control flow, so neither the interpreter's recover paths nor
// mustPair may write anything to the program's stdout (or, without
// SHEN_DEBUG_RECOVER, to stderr). It also pins Eval's contract that an
// uncaught error comes back as the raised error object, message intact,
// rather than as an Error whose text is a goroutine dump.
func TestRecoveredKLErrorsKeepStdoutClean(t *testing.T) {
	var ctl ControlFlow
	var caught, uncaught, msg Obj
	stdout, stderr := captureStdoutStderr(t, func() {
		caught = evalString(&ctl, `(trap-error (tl 5) (lambda E ok))`)
		uncaught = evalString(&ctl, `(tl 5)`)
		msg = evalString(&ctl, `(simple-error "boom")`)
	})
	if stdout != "" {
		t.Errorf("recovered KL errors wrote to stdout:\n%s", stdout)
	}
	if os.Getenv("SHEN_DEBUG_RECOVER") == "" && stderr != "" {
		t.Errorf("recovered KL errors wrote to stderr without SHEN_DEBUG_RECOVER:\n%s", stderr)
	}
	if got := ObjString(caught); got != "ok" {
		t.Errorf("trap-error handler result: got %q, want ok", got)
	}
	if !IsError(uncaught) {
		t.Errorf("(tl 5) should evaluate to an error, got %s", ObjString(uncaught))
	} else if s := mustError(uncaught).err; strings.Contains(s, "goroutine") {
		t.Errorf("(tl 5) error text is a stack dump, not the raised message:\n%s", s)
	}
	if !IsError(msg) {
		t.Fatalf("(simple-error \"boom\") should evaluate to an error, got %s", ObjString(msg))
	}
	if got := mustError(msg).err; got != "boom" {
		t.Errorf("Eval discarded the raised message: got %q, want %q", got, "boom")
	}
}

// TestGoPanicUnderTrapErrorTracesToStderr covers the other recover branch: a
// non-Obj Go panic inside a native is a bug, not control flow. It must still
// reach the program as an Error (so trap-error handlers keep the process
// alive) but its trace goes to stderr, unconditionally, and never to stdout.
// Both trap-error implementations are exercised: the interpreter's
// evalTrapError at top level and Try/Catch inside a compiled defun.
func TestGoPanicUnderTrapErrorTracesToStderr(t *testing.T) {
	bindGoPanicNative()
	var ctl ControlFlow
	var tried, interp, compiled Obj
	stdout, stderr := captureStdoutStderr(t, func() {
		tried = Try(&ctl, PrimFunc(MakeSymbol("test-go-panic"))).Catch(evalString(&ctl, `(lambda E E)`))
		interp = evalString(&ctl, `(trap-error (test-go-panic) (lambda E E))`)
		evalString(&ctl, `(defun test-go-panic-vm () (trap-error (test-go-panic) (lambda E (error-to-string E))))`)
		compiled = evalString(&ctl, `(test-go-panic-vm)`)
	})
	if stdout != "" {
		t.Errorf("recovered Go panics wrote to stdout:\n%s", stdout)
	}
	for _, want := range []string{"recovered in Try", "recovered in trap-error", "boom from test-go-panic", "goroutine "} {
		if !strings.Contains(stderr, want) {
			t.Errorf("stderr trace lacks %q:\n%s", want, stderr)
		}
	}
	if !IsError(tried) || !strings.Contains(mustError(tried).err, "boom from test-go-panic") {
		t.Errorf("Try.Catch handler should receive the panic as an Error, got %s", ObjString(tried))
	}
	if !IsError(interp) {
		t.Errorf("interpreted trap-error should yield an Error, got %s", ObjString(interp))
	}
	if got := ObjString(compiled); !strings.Contains(got, "boom from test-go-panic") {
		t.Errorf("compiled trap-error handler should receive the panic message, got %s", got)
	}
}
