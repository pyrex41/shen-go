package main

// Parity tests for the Go-native Shen reader (kl.ShenReadSExprs) against the
// interpreted kernel reader. The kernel's read-from-string-unprocessed is
//
//	String -> (let Bytelist (str->bytes String)
//	               S-exprs (compile (/. X (<s-exprs> X)) Bytelist)
//	               S-exprs)
//
// i.e. exactly the tokenize/parse stage our native reader replaces, with no
// process-sexprs. It is therefore the perfect oracle: for any (ASCII) source
// text, kl.ShenReadSExprs must produce a form list deeply equal to what
// read-from-string-unprocessed produces — or decline (fall back), which is
// always safe.
//
// The interpreted reader is ~quadratic in input length (it is the slowness this
// work targets), so whole-file oracle comparison is bounded to the smaller
// human-written sources. Two cheaper-but-sharper tests cover the rest:
//   - every numeric literal token across the WHOLE corpus is oracle-compared
//     individually (float-arithmetic parity is the subtle risk, and each token
//     is tiny), and
//   - the Go reader is required to engage (not decline) and to be deterministic
//     on every corpus file, including the pathological 375 KB stlib.kl.

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/pyrex41/shen-go/kl"
)

// oracleMaxBytes bounds whole-file oracle comparison so the ~quadratic
// interpreted reader keeps the test fast. All kernel/sources/*.shen files are
// under this; the large compiled klambda/*.kl files are covered by the numeric
// and engagement tests instead.
const oracleMaxBytes = 23000

// interpretedParse runs the interpreted <s-exprs> on src via
// read-from-string-unprocessed, returning the raw form list. It recovers any
// kernel-raised error into a Go error so the test can compare error-vs-decline.
func interpretedParse(e *kl.ControlFlow, src string) (result kl.Obj, err error) {
	defer func() {
		if r := recover(); r != nil {
			if o, ok := r.(kl.Obj); ok && kl.IsError(o) {
				err = errFromObj(o)
				return
			}
			panic(r)
		}
	}()
	fn := kl.PrimFunc(kl.MakeSymbol("read-from-string-unprocessed"))
	res := kl.Call(e, fn, kl.MakeString(src))
	if kl.IsError(res) {
		return nil, errFromObj(res)
	}
	return res, nil
}

type objErr struct{ msg string }

func (o objErr) Error() string { return o.msg }

func errFromObj(o kl.Obj) error {
	return objErr{kl.GetString(kl.PrimErrorToString(o))}
}

// assertParity checks that the Go reader either matches the interpreted reader
// exactly, or declined (acceptable but counted). It fails only on a confident
// mismatch — the dangerous outcome.
func assertParity(t *testing.T, e *kl.ControlFlow, label, src string) (matched bool) {
	t.Helper()
	goForms, goErr := kl.ShenReadSExprs([]byte(src))
	if goErr != nil {
		return false // declined -> interpreted reader handles it; always safe
	}
	want, wantErr := interpretedParse(e, src)
	if wantErr != nil {
		t.Errorf("%s: Go reader accepted %q but interpreted reader errored: %v",
			label, src, wantErr)
		return false
	}
	if !kl.Equal(goForms, want) {
		t.Errorf("%s: parse mismatch for %q\n  go:   %s\n  want: %s",
			label, src, kl.ObjString(goForms), kl.ObjString(want))
		return false
	}
	return true
}

func TestNativeReaderParityLiterals(t *testing.T) {
	e := bootShen(t)

	cases := []string{
		// atoms / symbols
		`abc`, `a-b-c?`, `foo123`, `->`, `<-`, `:=x`, `*foo*`, `+`, `-`, `.`,
		`x.y`, `a#b`, `$money`, `'quoted`, "`back", `log10`, `2pi`,
		// the empty-vector token
		`<>`, `(a <> b)`, `[<>]`,
		// numbers: integers, floats, signs, e-notation (incl. real stlib constants)
		`0`, `42`, `-7`, `+9`, `3.14`, `-2.5`, `.5`, `1.6180339887498`,
		`2.7182818284590002`, `3.1415926535897`, `0.8660254037844001`,
		`0.70710678118651`, `1e3`, `1e-4`, `1e4`, `1e10`, `1.5e2`, `-1.5e-3`,
		`123abc`, `1.2.3`, `+1a`, `1e`, `1.`, `1E3`,
		// strings, including the c#NN; char-code escape
		`"hello"`, `""`, `"c#65;BC"`, `"line c#10; break"`, `"a c#16;b c#17;"`,
		`"tab	and spaces"`,
		// lists, brackets, dotted-tail bar
		`(a b c)`, `(a (b c) d)`, `()`, `[a b c]`, `[]`, `[a | b]`, `[a b | c]`,
		`[a [b c] | d]`, `(cons 1 (cons 2 ()))`,
		// structural tokens used by type/datatype/yacc syntax
		`(a : b)`, `(a := b)`, `(a ; b)`, `(a , b)`, `{ a }`, `(P | Q)`,
		// comments
		"\\\\ a comment\n(defun f () 1)",
		"(defun f () 1) \\\\ trailing\n(defun g () 2)",
		"\\* block comment *\\ (real form)",
		"\\* nested \\* inner *\\ still in *\\ x",
		// multi-form input, mixed whitespace
		"(a)\n(b)\n  (c)  ",
		`(define foo X -> (+ X 1))`,
		// kernel-style datatype/yacc snippet
		`(datatype foo X : bar; ____ [X|Y] : (list bar);)`,
	}

	declined := 0
	for _, src := range cases {
		if !assertParity(t, e, "literal", src) {
			if _, goErr := kl.ShenReadSExprs([]byte(src)); goErr != nil {
				declined++
			}
		}
	}
	t.Logf("literal cases: %d total, %d declined (fell back to interpreted)",
		len(cases), declined)
}

// numberToken matches a Shen numeric literal: optional sign, digits with an
// optional fraction or a leading dot, and an optional lowercase-e exponent.
var numberToken = regexp.MustCompile(`[-+]?(?:[0-9]+\.?[0-9]*|\.[0-9]+)(?:e[-+]?[0-9]+)?`)

// TestNativeReaderParityNumbersCorpus oracle-compares every numeric literal
// token found anywhere in the corpus — including stlib.kl's many float
// constants and e-notation — without paying the quadratic whole-file cost. This
// is the sharp test for float-arithmetic parity (compute-integer / fraction /
// expt), which is where strconv.ParseFloat would silently diverge.
func TestNativeReaderParityNumbersCorpus(t *testing.T) {
	e := bootShen(t)

	seen := map[string]bool{}
	for _, f := range corpusFiles(t) {
		data, err := os.ReadFile(f)
		if err != nil || !isASCII(data) {
			continue
		}
		for _, tok := range numberToken.FindAllString(string(data), -1) {
			seen[tok] = true
		}
	}
	if len(seen) == 0 {
		t.Fatal("no numeric tokens extracted from corpus")
	}

	var checked, declined int
	for tok := range seen {
		if assertParity(t, e, "number", tok) {
			checked++
		} else if _, goErr := kl.ShenReadSExprs([]byte(tok)); goErr != nil {
			declined++
		}
	}
	t.Logf("numeric tokens: %d unique, %d matched, %d declined", len(seen), checked, declined)
	if checked == 0 {
		t.Errorf("no numeric token matched — number parity unverified")
	}
}

// TestNativeReaderParitySources runs whole-file oracle parity over the bounded
// human-written sources, where every syntactic construct (datatypes, yacc
// defcc, type signatures, comments, strings) appears in real context.
func TestNativeReaderParitySources(t *testing.T) {
	e := bootShen(t)

	var matched, declined int
	for _, f := range corpusFiles(t) {
		data, err := os.ReadFile(f)
		if err != nil {
			t.Fatalf("read %s: %v", f, err)
		}
		if len(data) > oracleMaxBytes || !isASCII(data) {
			continue
		}
		goForms, goErr := kl.ShenReadSExprs(data)
		if goErr != nil {
			declined++
			t.Logf("declined (fallback) for %s", f)
			continue
		}
		want, wantErr := interpretedParse(e, string(data))
		if wantErr != nil {
			t.Errorf("%s: Go reader accepted but interpreted reader errored: %v", f, wantErr)
			continue
		}
		if !kl.Equal(goForms, want) {
			t.Errorf("%s: parse mismatch (Go reader diverged from interpreted)", f)
			continue
		}
		matched++
	}
	t.Logf("source files: %d matched, %d declined", matched, declined)
	if matched == 0 {
		t.Errorf("no source file matched — native reader never engaged")
	}
}

// TestNativeReaderEngagesCorpus asserts the Go reader engages (does not fall
// back) and is deterministic on EVERY corpus file — including the 375 KB
// stlib.kl with its pathological single mega-form and hundreds of float
// constants. This is the cheap whole-corpus guard that the fast path is
// actually taken and never crashes or yields a non-list.
func TestNativeReaderEngagesCorpus(t *testing.T) {
	for _, f := range corpusFiles(t) {
		data, err := os.ReadFile(f)
		if err != nil {
			t.Fatalf("read %s: %v", f, err)
		}
		if !isASCII(data) {
			continue
		}
		forms, perr := kl.ShenReadSExprs(data)
		if perr != nil {
			t.Errorf("%s: Go reader declined (fell back) — expected it to engage", f)
			continue
		}
		// Result must be a proper list of forms.
		if forms != nil && !isProperList(forms) {
			t.Errorf("%s: Go reader produced a non-list result", f)
		}
		// Deterministic: a second parse must be deeply equal.
		forms2, perr2 := kl.ShenReadSExprs(data)
		if perr2 != nil || !kl.Equal(forms, forms2) {
			t.Errorf("%s: Go reader not deterministic across two parses", f)
		}
	}
}

// BenchmarkReaderGoNative and BenchmarkReaderInterpreted measure the same parse
// — a real kernel source file (sys.shen, ~14 KB) — through the Go-native reader
// vs the interpreted <s-exprs>, quantifying the speedup the native reader gives
// the cold-load / one-shot path.
func BenchmarkReaderGoNative(b *testing.B) {
	data := readBenchSource(b)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := kl.ShenReadSExprs(data); err != nil {
			b.Fatalf("native reader declined: %v", err)
		}
	}
}

func BenchmarkReaderInterpreted(b *testing.B) {
	e := bootShen(b)
	data := readBenchSource(b)
	src := string(data)
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := interpretedParse(e, src); err != nil {
			b.Fatalf("interpreted reader errored: %v", err)
		}
	}
}

func readBenchSource(b *testing.B) []byte {
	b.Helper()
	data, err := os.ReadFile("../../kernel/sources/sys.shen")
	if err != nil {
		b.Fatalf("read bench source: %v", err)
	}
	return data
}

func corpusFiles(t testing.TB) []string {
	t.Helper()
	var files []string
	for _, root := range []string{"../../kernel/sources", "../../kernel/klambda"} {
		_ = filepath.Walk(root, func(p string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() {
				return nil
			}
			if strings.HasSuffix(p, ".kl") || strings.HasSuffix(p, ".shen") {
				files = append(files, p)
			}
			return nil
		})
	}
	if len(files) == 0 {
		t.Skip("no corpus files found")
	}
	return files
}

func isProperList(o kl.Obj) bool {
	for o != kl.Nil {
		ok, rest := kl.IsPair(o)
		if !ok {
			return false
		}
		o = rest
	}
	return true
}

func isASCII(b []byte) bool {
	for _, c := range b {
		if c >= 0x80 {
			return false
		}
	}
	return true
}
