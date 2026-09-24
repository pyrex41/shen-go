package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pyrex41/shen-go/kl"
)

// TestCompileFileKeepsFailSymbol is issue #54's regeneration guard.
// compiled/script.kl regenerates cmd/shen/*.go by running src/compiler.shen's
// compile-file in this bare interpreter, where (fail) is the kernel's own
// (defun fail () shen.fail!). compile-file used to serialize the IR with
// (make-string "~R" BC), and ~R renders anything equal to (fail) as `...`, so
// the compiled fail returned the wrong symbol. The IR printer pr-ir that
// replaced it prints symbols with str and everything else through ~R, so its
// output must be byte-identical to ~R except where a symbol equals (fail).
//
// Three checks, all in one interpreter run so the kernel boots once:
//  1. a fixed IR sample with no shen.fail! prints identically both ways; it
//     covers the symbols { } ; , | and (intern "..."), strings with ~ and ",
//     negative and float numbers, booleans, () and an improper list tail;
//  2. the fail defun's IR prints with shen.fail! where ~R writes `...`;
//  3. compile-file of (defun fail () shen.fail!) writes ($const shen.fail!).
func TestCompileFileKeepsFailSymbol(t *testing.T) {
	kernelDir, err := kl.FindKernelDir(".")
	if err != nil {
		t.Fatal(err)
	}
	compiler, err := filepath.Abs("../../src/compiler.shen")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(compiler); err != nil {
		t.Fatal(err)
	}
	bin := buildKL(t)
	dir := t.TempDir()
	path := func(name string) string { return filepath.Join(dir, name) }
	if err := os.WriteFile(path("fail.kl"), []byte("(defun fail () shen.fail!)\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	var in strings.Builder
	for _, f := range kl.KernelLoadOrder {
		in.WriteString("(load-file \"" + filepath.Join(kernelDir, f) + "\")\n")
	}
	in.WriteString("(load \"" + compiler + "\")\n")
	// script.kl sets this so ~R does not truncate long IR lists; the samples
	// here are short, but keep the two runs comparable. make-string is a
	// Shen macro, so the ~R side is spelled as its expansion, shen.app with
	// the shen.r mode.
	in.WriteString("(set *maximum-print-sequence-size* 100000)\n")
	// Check 1. The sample is built with cons/intern so the reader has no say
	// in it; the quote character comes from n->string.
	in.WriteString(`(set t.sample1 (cons (intern "{") (cons (intern "}") (cons (intern ";") (cons (intern ",") (cons (intern "|") (cons (intern "...") (cons "a~b" (cons (cn "q" (cn (n->string 34) "x")) (cons -3 (cons 2.5 (cons true (cons false (cons () (cons (cons (intern "$const") (cons 1 (cons "s" ()))) (cons (cons (intern "$global") (cons shen.foo? ())) (cons (cons 1 (cons 2 3)) ())))))))))))))))))` + "\n")
	in.WriteString(`(let S (open "` + path("sample1.pr-ir") + `" out) (do (pr-ir (value t.sample1) S) (close S)))` + "\n")
	in.WriteString(`(write-to-file "` + path("sample1.R") + `" (shen.app (value t.sample1) "" shen.r))` + "\n")
	// Check 2. The IR compile-file emits for (defun fail () shen.fail!).
	in.WriteString(`(set t.sample2 (cons (intern "<=") (cons tmp1 (cons (cons lambda (cons () (cons (cons return (cons (cons (intern "$const") (cons shen.fail! ())) ())) ()))) ()))))` + "\n")
	in.WriteString(`(let S (open "` + path("sample2.pr-ir") + `" out) (do (pr-ir (value t.sample2) S) (close S)))` + "\n")
	in.WriteString(`(write-to-file "` + path("sample2.R") + `" (shen.app (value t.sample2) "" shen.r))` + "\n")
	// Check 3.
	in.WriteString(`(compile-file "` + path("fail.kl") + `" "` + path("fail.tmp") + `")` + "\n")

	stdout, stderr := runKL(t, bin, in.String())
	for _, bad := range []string{"Error(", "Panic:", "Recovered in"} {
		if strings.Contains(stdout, bad) || strings.Contains(stderr, bad) {
			t.Fatalf("kl reported %q:\n%s%s", bad, stdout, stderr)
		}
	}
	read := func(name string) string {
		data, err := os.ReadFile(path(name))
		if err != nil {
			t.Fatalf("%v\nkl output:\n%s%s", err, stdout, stderr)
		}
		return string(data)
	}

	// 1. Byte-identical to ~R away from the fail symbol.
	prIR, viaR := read("sample1.pr-ir"), read("sample1.R")
	if prIR != viaR {
		t.Errorf("pr-ir and ~R differ on the sample without shen.fail!:\npr-ir: %s\n~R:    %s", prIR, viaR)
	}
	const wantSample1 = `({ } ; , | ... "a~b" "q"x" -3 2.5 true false () ($const 1 "s") ($global shen.foo?) (1 2 | 3))`
	if prIR != wantSample1 {
		t.Errorf("pr-ir sample1:\n got %s\nwant %s", prIR, wantSample1)
	}

	// 2. The fail symbol survives, and this is exactly where ~R loses it.
	prIR, viaR = read("sample2.pr-ir"), read("sample2.R")
	const wantFailIR = `(<= tmp1 (lambda () (return ($const shen.fail!))))`
	if prIR != wantFailIR {
		t.Errorf("pr-ir fail IR:\n got %s\nwant %s", prIR, wantFailIR)
	}
	if viaR != strings.Replace(wantFailIR, "shen.fail!", "...", 1) {
		t.Errorf("~R fail IR = %s; expected it to render shen.fail! as ... (the reason pr-ir exists)", viaR)
	}

	// 3. compile-file itself.
	tmp := read("fail.tmp")
	if !strings.Contains(tmp, "($const shen.fail!)") || strings.Contains(tmp, "($const ...)") {
		t.Errorf("compile-file of (defun fail () shen.fail!) wrote:\n%s\nwant ($const shen.fail!) and no ($const ...)", tmp)
	}
}
