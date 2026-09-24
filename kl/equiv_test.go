package kl

// Differential test of every InstallKernelFast native against the kernel's own
// KLambda definition, and the generator of kl/equiv.json. equiv.go holds the
// harness and documents the JSON schema.
//
//	go test ./kl -run TestEquiv                    # check the natives and the table
//	EQUIV_WRITE=1 go test ./kl -run TestEquivTable # regenerate kl/equiv.json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"hash/fnv"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"
)

// ---- the rebound-function table, parsed from kernelfast.go -----------------

type fastBinding struct {
	kernelFn string
	native   string
	arity    int
	source   string // "kl/FILE.go:Ident"
	helper   string // the InstallKernelFast call that binds it: overridePrimitive, overrideNative, BindSymbolFunc, restoreCanonicalPrimitive
}

// parseInstallKernelFast extracts the kernel name, the native Go identifier
// and the arity of every rebinding InstallKernelFast performs, in source
// order, by walking its body with go/ast. It recognises these forms:
//
//	overridePrimitive("name", N, ident | func literal calling ident)
//	overrideNative("name", N, ident | func literal calling ident)
//	BindSymbolFunc(symArity | MakeSymbol("fn"), canonicalOrMake("name", N, ident) | MakeNative(ident, N))
//	restoreCanonicalPrimitive("name")   -> the canonical primitive registered in primitives.go
func parseInstallKernelFast(t *testing.T) []fastBinding {
	t.Helper()
	fset := token.NewFileSet()
	kf, err := parser.ParseFile(fset, "kernelfast.go", nil, 0)
	if err != nil {
		t.Fatal(err)
	}
	pf, err := parser.ParseFile(fset, "primitives.go", nil, 0)
	if err != nil {
		t.Fatal(err)
	}
	// Where each Go identifier is declared, and its parameter count.
	declFile := map[string]string{}
	declArity := map[string]int{}
	for file, f := range map[string]*ast.File{"kernelfast.go": kf, "primitives.go": pf} {
		for _, d := range f.Decls {
			fd, ok := d.(*ast.FuncDecl)
			if !ok || fd.Recv != nil {
				continue
			}
			declFile[fd.Name.Name] = file
			n := 0
			for _, p := range fd.Type.Params.List {
				n += len(p.Names)
			}
			declArity[fd.Name.Name] = n
		}
	}
	// primitiveRegistry.register("name", Ident) in primitives.go.
	canonical := map[string]string{}
	ast.Inspect(pf, func(n ast.Node) bool {
		c, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		sel, ok := c.Fun.(*ast.SelectorExpr)
		if !ok || sel.Sel.Name != "register" || len(c.Args) != 2 {
			return true
		}
		if x, ok := sel.X.(*ast.Ident); !ok || x.Name != "primitiveRegistry" {
			return true
		}
		name, ok := strLit(c.Args[0])
		if !ok {
			return true
		}
		if id, ok := c.Args[1].(*ast.Ident); ok {
			canonical[name] = id.Name
		}
		return true
	})

	var install *ast.FuncDecl
	for _, d := range kf.Decls {
		if fd, ok := d.(*ast.FuncDecl); ok && fd.Name.Name == "InstallKernelFast" {
			install = fd
		}
	}
	if install == nil {
		t.Fatal("InstallKernelFast not found in kernelfast.go")
	}
	src := func(ident string) string {
		file, ok := declFile[ident]
		if !ok {
			t.Fatalf("native %s: no FuncDecl in kernelfast.go or primitives.go", ident)
		}
		return "kl/" + file + ":" + ident
	}
	var rows []fastBinding
	for _, stmt := range install.Body.List {
		es, ok := stmt.(*ast.ExprStmt)
		if !ok {
			continue
		}
		c, ok := es.X.(*ast.CallExpr)
		if !ok {
			continue
		}
		fn, ok := c.Fun.(*ast.Ident)
		if !ok {
			continue
		}
		switch fn.Name {
		case "overridePrimitive", "overrideNative":
			name, ok1 := strLit(c.Args[0])
			arity, ok2 := intLit(c.Args[1])
			native := calleeIdent(c.Args[2])
			if !ok1 || !ok2 || native == "" {
				t.Fatalf("cannot decode %s call at %s", fn.Name, fset.Position(c.Pos()))
			}
			rows = append(rows, fastBinding{name, native, arity, src(native), fn.Name})
		case "BindSymbolFunc":
			var name string
			switch a := c.Args[0].(type) {
			case *ast.Ident:
				if a.Name != "symArity" {
					t.Fatalf("BindSymbolFunc with unknown symbol var %s at %s", a.Name, fset.Position(c.Pos()))
				}
				name = "arity"
			case *ast.CallExpr:
				name, _ = strLit(a.Args[0])
			}
			mk, ok := c.Args[1].(*ast.CallExpr)
			if !ok || name == "" {
				t.Fatalf("cannot decode BindSymbolFunc at %s", fset.Position(c.Pos()))
			}
			var native string
			var arity int
			switch mk.Fun.(*ast.Ident).Name {
			case "canonicalOrMake":
				arity, _ = intLit(mk.Args[1])
				native = calleeIdent(mk.Args[2])
			case "MakeNative":
				native = calleeIdent(mk.Args[0])
				arity, _ = intLit(mk.Args[1])
			default:
				t.Fatalf("cannot decode BindSymbolFunc value at %s", fset.Position(c.Pos()))
			}
			rows = append(rows, fastBinding{name, native, arity, src(native), fn.Name})
		case "restoreCanonicalPrimitive":
			name, _ := strLit(c.Args[0])
			native, ok := canonical[name]
			if !ok {
				t.Fatalf("restoreCanonicalPrimitive(%q): no primitiveRegistry.register in primitives.go", name)
			}
			rows = append(rows, fastBinding{name, native, declArity[native], src(native), fn.Name})
		}
	}
	return rows
}

func strLit(e ast.Expr) (string, bool) {
	bl, ok := e.(*ast.BasicLit)
	if !ok || bl.Kind != token.STRING {
		return "", false
	}
	s, err := strconv.Unquote(bl.Value)
	return s, err == nil
}

func intLit(e ast.Expr) (int, bool) {
	bl, ok := e.(*ast.BasicLit)
	if !ok || bl.Kind != token.INT {
		return 0, false
	}
	n, err := strconv.Atoi(bl.Value)
	return n, err == nil
}

// calleeIdent names the Go function an override argument denotes. That is the
// identifier itself, or, for a func literal wrapper, the first function the
// literal calls.
func calleeIdent(e ast.Expr) string {
	switch a := e.(type) {
	case *ast.Ident:
		return a.Name
	case *ast.FuncLit:
		name := ""
		ast.Inspect(a.Body, func(n ast.Node) bool {
			if name != "" {
				return false
			}
			if c, ok := n.(*ast.CallExpr); ok {
				if id, ok := c.Fun.(*ast.Ident); ok && id.Name != "e" {
					// Skip the e.Return wrapper: its argument is the real call.
					name = id.Name
					return false
				}
			}
			return true
		})
		return name
	}
	return ""
}

// ---- per-row metadata: argument kinds and global effects -------------------

type kind int

const (
	kAny kind = iota
	kList
	kString
	kNumber
	kSymbol
	kVector
	kThunk
	kFunc
	kAlist
	kNumList
	kPropKey // key/attr/value position of a property-vector function
)

// rowSignature is the argument kinds each rebound function accepts. A name
// missing here, or present here but no longer rebound, fails TestEquivTable.
// That is the drift check between this table and InstallKernelFast.
var rowSignature = map[string][]kind{
	"arity":                  {kSymbol},
	"fn":                     {kSymbol},
	"not":                    {kAny},
	"integer?":               {kAny},
	"empty?":                 {kAny},
	"boolean?":               {kAny},
	"vector?":                {kAny},
	"tuple?":                 {kAny},
	"shen.pvar?":             {kAny},
	"symbol?":                {kAny},
	"variable?":              {kAny},
	"shen.analyse-symbol?":   {kString},
	"shen.analyse-variable?": {kString},
	"shen.digit?":            {kNumber},
	"shen.lowercase?":        {kNumber},
	"shen.uppercase?":        {kNumber},
	"shen.misc?":             {kNumber},
	"shen.alpha?":            {kNumber},
	"shen.alphanums?":        {kString},
	"shen.+string?":          {kAny},
	"shen.hds=?":             {kList, kAny},
	"@p":                     {kAny, kAny},
	"vector":                 {kNumber},
	"<-vector":               {kVector, kNumber},
	"vector->":               {kVector, kNumber, kAny},
	"limit":                  {kVector},
	"fst":                    {kVector},
	"snd":                    {kVector},
	"hdstr":                  {kString},
	"shen.byte->digit":       {kNumber},
	"thaw":                   {kThunk},
	"fail":                   {},
	"length":                 {kList},
	"reverse":                {kList},
	"append":                 {kList, kList},
	"element?":               {kAny, kList},
	"assoc":                  {kAny, kAlist},
	"put":                    {kPropKey, kPropKey, kPropKey, kVector},
	"get":                    {kPropKey, kPropKey, kVector},
	"unput":                  {kPropKey, kPropKey, kVector},
	"head":                   {kList},
	"tail":                   {kList},
	"nth":                    {kNumber, kList},
	"bound?":                 {kSymbol},
	"concat":                 {kAny, kAny},
	"==":                     {kAny, kAny},
	"shen.abs":               {kNumber},
	"shen.posint?":           {kAny},
	"sum":                    {kNumList},
	"adjoin":                 {kAny, kList},
	"remove":                 {kAny, kList},
	"string->symbol":         {kString},
	"shen.string->bytes":     {kString},
	"protect":                {kAny},
	"union":                  {kList, kList},
	"intersection":           {kList, kList},
	"difference":             {kList, kList},
	"map":                    {kFunc, kList},
}

// rowEffects lists the globals a native reads or writes, by symbol name.
// put, get and unput mutate the vector they are handed, and the kernel only
// ever passes (value *property-vector*), which is why the table names it.
// bound? reads the value cell of whatever symbol it is given, written
// "(value X)".
var rowEffects = map[string][]string{
	"arity":  {"*property-vector*"},
	"fn":     {"*property-vector*", "shen.*lambdatable*"},
	"put":    {"*property-vector*"},
	"get":    {"*property-vector*"},
	"unput":  {"*property-vector*"},
	"bound?": {"(value X)"},
}

const effectsDoc = "The global symbols this native reads or writes. \"(value X)\" means the value cell of the symbol passed as an argument. A row with no effects is pure."

// propertyRows are the rows whose function takes the property vector as its
// last argument.
var propertyRows = map[string]bool{"put": true, "get": true, "unput": true}

// ---- case generation -------------------------------------------------------

const iota10k = "(equiv.iota 10000)"

var boundary = map[kind][]string{
	kAny: {"0", "-1", "1", "1000000000000", "1.5", "-2.5", `""`, `"a"`, `"λ"`, "foo", "shen.foo", "*foo*", "Abc",
		"true", "false", "()", "(cons 1 ())", "(cons (cons 1 ()) (cons 2 ()))", "(vector 0)", "(vector 1)", "(@p 1 2)", "(absvector 0)",
		`(intern "1")`, "(lambda X X)"},
	kList: {"()", "(cons 1 ())", "(cons 1 (cons 2 (cons 3 ())))", "(cons (cons 1 (cons 2 ())) (cons () (cons (cons 3 ()) ())))",
		`(cons "" (cons "a" (cons "λ" ())))`, "(cons a (cons shen.b (cons *c* ())))", "(cons 1 (cons 1.0 (cons -1 ())))",
		"(cons true (cons false ()))", "(cons (vector 0) (cons (vector 1) ()))", iota10k, "5", `"notalist"`, "foo", "(cons 1 2)", "(cons 1 (cons 2 3))"},
	kString: {`""`, `"a"`, `"λ"`, `"aλb"`, `"*foo*"`, `"1"`, `"A1"`, `"Abc"`, `"a-b"`, `"hello world"`, `"shen.x"`, `"{"`, `"_"`, `"a1_b?"`, `"1a"`,
		"5", "foo", "()", "true"},
	kNumber: {"0", "-1", "1", "47", "48", "57", "58", "64", "65", "90", "91", "96", "97", "122", "123", "127", "128", "42", "61", "1.5", "48.0", "-0.5",
		"1000000000000", "100000000000000000000.0", `"a"`, "foo", "()"},
	kSymbol: {"foo", "shen.foo", "*foo*", "Abc", "X", `(intern "1")`, `(intern "")`, `(intern "a b")`, `(intern "*foo*")`, "{", `(intern ":")`,
		`(intern "-")`, "reverse", "fail", "equiv.undefined-fn", "5", `"str"`, "()", "true"},
	kVector: {"(vector 0)", "(vector 1)", "(vector 3)", "(absvector 0)", "(absvector 1)", "(absvector 3)", "(@p 1 2)",
		"(let V (vector 2) (do (vector-> V 1 a) V))", "5", "()", `"s"`},
	kThunk: {"(freeze 1)", "(freeze (equiv.note 1))", `(freeze (simple-error "boom"))`, "(equiv.native-of fail)", "(lambda X X)", "5", "foo"},
	kFunc: {"(lambda X X)", "(lambda X (equiv.note X))", "(lambda X (cons X X))", "(equiv.native-of equiv.note)", "(equiv.native-of reverse)",
		`(lambda X (simple-error "boom"))`, "(fn reverse)", "5", "foo"},
	kAlist: {"()", "(cons (cons a 1) (cons (cons b 2) ()))", "(cons (cons 1 x) (cons 5 ()))", "(cons a ())", "(cons (cons a 1) (cons (cons a 2) ()))",
		"(cons (cons () 1) ())", "5"},
	kNumList: {"()", "(cons 1 ())", "(cons 1 (cons 2.5 (cons -3 ())))", `(cons 1 (cons "a" ()))`, "(cons 1000000000000 (cons 1000000000000 ()))", "5", iota10k},
	kPropKey: {"a", "b", "shen.x", `"str"`, "1", "1.0", "()", "(cons a (cons b ()))", "true", "arity"},
}

// typical is the value the other parameters take while one parameter walks its
// boundary list.
var typical = map[kind]string{
	kAny: "1", kList: "(cons 1 (cons 2 (cons 3 ())))", kString: `"abc"`, kNumber: "48", kSymbol: "foo", kVector: "(vector 3)",
	kThunk: "(freeze 1)", kFunc: "(lambda X X)", kAlist: "(cons (cons a 1) (cons (cons b 2) ()))", kNumList: "(cons 1 (cons 2 ()))", kPropKey: "a",
}

// explicitCases are the hand-written cases for a row, including the
// property-vector sequences. For a property row, finishCase appends the vector
// argument and adds the observe of the vector.
var explicitCases = map[string][]EquivCase{
	"reverse": {
		{Name: "empty", Args: []string{"()"}},
		{Name: "one", Args: []string{"(cons 1 ())"}},
		{Name: "10k", Args: []string{iota10k}},
	},
	"hdstr":    {{Name: "empty-string-raises", Args: []string{`""`}}},
	"element?": {{Name: "int-vs-float", Args: []string{"1", "(cons 1.0 ())"}}, {Name: "float-vs-int", Args: []string{"1.0", "(cons 1 ())"}}},
	"nth": {
		{Name: "index-0", Args: []string{"0", "(cons a (cons b ()))"}},
		{Name: "index-1", Args: []string{"1", "(cons a (cons b ()))"}},
		{Name: "index-2", Args: []string{"2", "(cons a (cons b ()))"}},
		{Name: "index-3-out-of-range", Args: []string{"3", "(cons a (cons b ()))"}},
		{Name: "negative", Args: []string{"-1", "(cons a (cons b ()))"}},
		{Name: "float-index", Args: []string{"1.0", "(cons a (cons b ()))"}},
		{Name: "non-number-index", Args: []string{"a", "(cons a ())"}},
	},
	"map": {
		{Name: "never-called", Args: []string{"(lambda X (equiv.note X))", "()"}},
		{Name: "call-order", Args: []string{"(lambda X (equiv.note X))", "(cons 1 (cons 2 (cons 3 ())))"}},
		{Name: "native-fn", Args: []string{"(equiv.native-of equiv.note)", "(cons 1 (cons 2 (cons 3 ())))"}},
		{Name: "raising-fn-midway", Args: []string{`(lambda X (if (= X 2) (simple-error "boom") (equiv.note X)))`, "(cons 1 (cons 2 (cons 3 ())))"}},
	},
	"thaw": {
		{Name: "freeze-effect-once", Args: []string{"(freeze (equiv.note 1))"}},
		{Name: "native-thunk", Args: []string{"(equiv.native-of fail)"}},
	},
	"vector": {{Name: "zero", Args: []string{"0"}}, {Name: "negative", Args: []string{"-1"}}, {Name: "fractional", Args: []string{"1.5"}}},
	"vector->": {
		{Name: "zero-size-slot-0", Args: []string{"(vector 0)", "0", "x"}},
		{Name: "zero-size-slot-1", Args: []string{"(vector 0)", "1", "x"}},
		{Name: "slot-1", Args: []string{"(vector 1)", "1", "x"}},
		{Name: "fractional-index", Args: []string{"(vector 3)", "-0.5", "x"}},
	},
	"<-vector": {
		{Name: "zero-size-slot-0", Args: []string{"(vector 0)", "0"}},
		{Name: "unset-slot-raises", Args: []string{"(vector 1)", "1"}},
		{Name: "set-slot", Args: []string{"(let V (vector 1) (do (vector-> V 1 x) V))", "1"}},
		{Name: "fractional-index", Args: []string{"(vector 3)", "-0.5"}},
	},
	// On this port string->n yields the code point, so both the KL body and
	// the native return (955) for "λ" rather than its UTF-8 bytes 206 187.
	// The case checks that the two sides agree on that, not that the result is
	// bytes.
	"shen.string->bytes": {{Name: "lambda-codepoint", Args: []string{`"λ"`}}, {Name: "mixed", Args: []string{`"aλ€"`}}},
	"string->symbol":     {{Name: "stars", Args: []string{`"*foo*"`}}, {Name: "digit-raises", Args: []string{`"1"`}}},
	"shen.misc?":         {{Name: "non-number", Args: []string{`"a"`}}},
	"not":                {{Name: "non-boolean", Args: []string{"0"}}},
	"symbol?":            {{Name: "intern-digit", Args: []string{`(intern "1")`}}, {Name: "intern-empty", Args: []string{`(intern "")`}}},
	// arity and fn read the kernel's own tables, so these cases seed the real
	// *property-vector*, identically for both sides. The last observe form is
	// teardown: it unputs the seed so that nothing leaks into later rows. Its
	// value, the key, is compared like any other observe.
	"arity": {
		{Name: "kernel-fn", Args: []string{"reverse"}},
		{Name: "unknown", Args: []string{"equiv.undefined-fn"}},
		{Name: "custom-arity", Setup: []string{"(put equiv.zero arity 0 (value *property-vector*))"}, Args: []string{"equiv.zero"},
			Observe: []string{"(unput equiv.zero arity (value *property-vector*))"}},
	},
	"fn": {
		{Name: "kernel-fn", Args: []string{"reverse"}},
		{Name: "kl-fn", Args: []string{"shen.app"}},
		{Name: "undefined-raises", Args: []string{"equiv.undefined-fn"}},
		{Name: "zero-arity-applies", Setup: []string{"(defun equiv.zero () 42)", "(put equiv.zero arity 0 (value *property-vector*))"}, Args: []string{"equiv.zero"},
			Observe: []string{"(unput equiv.zero arity (value *property-vector*))"}},
		{Name: "fail-applies", Args: []string{"fail"}},
	},
	"bound?": {
		{Name: "set-symbol", Setup: []string{"(set equiv.bound-x 1)"}, Args: []string{"equiv.bound-x"}},
		{Name: "unset-symbol", Args: []string{"equiv.never-set"}},
		{Name: "set-to-marker", Setup: []string{"(set equiv.bound-m shen.this-symbol-is-unbound)"}, Args: []string{"equiv.bound-m"}},
	},
	"put": {
		{Name: "fresh", Args: []string{"a", "b", "c"}},
		// A put, then a get, an unput, and a get that raises, with the vector
		// observed after every step. Only the put is the function under test.
		// The later steps run the kernel's own get and unput on both sides, so
		// the observed vectors show what the put left behind.
		{Name: "put-get-unput-get", Args: []string{"a", "b", "c"}, Observe: []string{
			"(value equiv.*pv*)",
			"(get a b (value equiv.*pv*))",
			"(value equiv.*pv*)",
			"(unput a b (value equiv.*pv*))",
			"(value equiv.*pv*)",
			"(trap-error (get a b (value equiv.*pv*)) (lambda E (error-to-string E)))",
			"(value equiv.*pv*)",
		}},
		{Name: "overwrite", Setup: []string{"(put a b c (value equiv.*pv*))"}, Args: []string{"a", "b", "d"}},
		{Name: "second-attr", Setup: []string{"(put a b c (value equiv.*pv*))"}, Args: []string{"a", "e", "f"}},
		{Name: "collision", Setup: []string{"(set equiv.*pv* (vector 1))", "(put a b c (value equiv.*pv*))", "(put x y z (value equiv.*pv*))"}, Args: []string{"q", "r", "s"}},
		{Name: "zero-size-vector", Setup: []string{"(set equiv.*pv* (vector 0))"}, Args: []string{"a", "b", "c"}},
	},
	"get": {
		{Name: "after-put", Setup: []string{"(put a b c (value equiv.*pv*))"}, Args: []string{"a", "b"}},
		{Name: "missing-attr", Setup: []string{"(put a b c (value equiv.*pv*))"}, Args: []string{"a", "z"}},
		{Name: "missing-key", Args: []string{"a", "b"}},
		{Name: "after-unput-raises", Setup: []string{"(put a b c (value equiv.*pv*))", "(unput a b (value equiv.*pv*))"}, Args: []string{"a", "b"}},
		{Name: "zero-size-vector", Setup: []string{"(set equiv.*pv* (vector 0))"}, Args: []string{"a", "b"}},
	},
	"unput": {
		{Name: "after-put", Setup: []string{"(put a b c (value equiv.*pv*))"}, Args: []string{"a", "b"}},
		{Name: "missing", Args: []string{"a", "b"}},
		{Name: "collision", Setup: []string{"(set equiv.*pv* (vector 1))", "(put a b c (value equiv.*pv*))", "(put x y z (value equiv.*pv*))"}, Args: []string{"x", "y"}},
		{Name: "zero-size-vector", Setup: []string{"(set equiv.*pv* (vector 0))"}, Args: []string{"a", "b"}},
	},
}

// finishCase adds the bookkeeping every case of a row needs. A property row
// gets a fresh property vector, unless the case set its own, the vector as its
// last argument, and the vector observed. Any case that uses equiv.note gets
// the call log reset and observed.
func finishCase(row string, c EquivCase) EquivCase {
	if propertyRows[row] {
		hasPV := false
		for _, s := range c.Setup {
			if strings.HasPrefix(s, "(set equiv.*pv* ") {
				hasPV = true
			}
		}
		if !hasPV {
			c.Setup = append([]string{"(set equiv.*pv* (vector 20))"}, c.Setup...)
		}
		c.Args = append(append([]string(nil), c.Args...), "(value equiv.*pv*)")
		c.Observe = append(c.Observe, "(value equiv.*pv*)")
	}
	all := strings.Join(c.Args, " ") + strings.Join(c.Setup, " ")
	if strings.Contains(all, "equiv.note") {
		c.Setup = append([]string{"(set equiv.*calls* ())"}, c.Setup...)
		c.Observe = append(c.Observe, "(value equiv.*calls*)")
	}
	return c
}

// boundaryCases walks each parameter's boundary list while the other
// parameters hold their typical value, then crosses the first six boundary
// values of every parameter.
func boundaryCases(row string, sig []kind) []EquivCase {
	var out []EquivCase
	seen := map[string]bool{}
	add := func(args []string) {
		key := strings.Join(args, "\x00")
		if seen[key] {
			return
		}
		seen[key] = true
		out = append(out, EquivCase{Name: fmt.Sprintf("boundary:%d", len(out)), Args: args})
	}
	if len(sig) == 0 {
		add(nil)
		return out
	}
	for i, k := range sig {
		for _, b := range boundary[k] {
			args := make([]string, len(sig))
			for j, kj := range sig {
				args[j] = typical[kj]
			}
			args[i] = b
			add(args)
		}
	}
	if len(sig) >= 2 {
		const n = 6
		var cross func(i int, acc []string)
		cross = func(i int, acc []string) {
			if i == len(sig) {
				add(append([]string(nil), acc...))
				return
			}
			bs := boundary[sig[i]]
			if len(bs) > n {
				bs = bs[:n]
			}
			for _, b := range bs {
				cross(i+1, append(acc, b))
			}
		}
		cross(0, nil)
	}
	return out
}

// kernelUsageCases extracts, mechanically, every call of row inside
// kernel/klambda/*.kl whose arguments are all literals the reader can
// round-trip. Those are numbers, strings without quotes or backslashes,
// booleans, (), non-variable symbols, and cons or intern of such literals. For
// the property-vector rows the vector argument is never a literal but always
// (value *property-vector*); those call sites are taken with the vector
// dropped, and finishCase supplies the fresh vector.
func kernelUsageCases(t *testing.T, dir, row string, arity int) []EquivCase {
	t.Helper()
	var out []EquivCase
	seen := map[string]bool{}
	sym := MakeSymbol(row)
	pvArg, _ := ReadForm("(value *property-vector*)")
	for _, f := range KernelLoadOrder {
		fh, err := os.Open(filepath.Join(dir, f))
		if err != nil {
			t.Fatal(err)
		}
		forms, err := ReadForms(fh)
		fh.Close()
		if err != nil {
			t.Fatal(err)
		}
		var walk func(o Obj)
		walk = func(o Obj) {
			ok, p := isPair(o)
			if !ok {
				return
			}
			if p.car == sym {
				args := ListToSlice(p.cdr)
				want := arity
				if propertyRows[row] {
					// Only call sites on the kernel's own vector count. The
					// vector argument is dropped and finishCase supplies one.
					if len(args) != arity || equal(args[arity-1], pvArg) != True {
						args, want = nil, -1
					} else {
						args, want = args[:arity-1], arity-1
					}
				}
				if len(args) == want && allLiteral(args) {
					strs := make([]string, len(args))
					for i, a := range args {
						strs[i] = klExprString(a)
					}
					key := strings.Join(strs, "\x00")
					if !seen[key] {
						seen[key] = true
						out = append(out, EquivCase{Name: fmt.Sprintf("kernel:%s:%d", f, len(out)), Args: strs})
					}
				}
			}
			for cur := o; ; {
				ok, cp := isPair(cur)
				if !ok {
					return
				}
				walk(cp.car)
				cur = cp.cdr
			}
		}
		for _, form := range forms {
			walk(form)
		}
	}
	return out
}

func allLiteral(args []Obj) bool {
	for _, a := range args {
		if !isLiteral(a) {
			return false
		}
	}
	return true
}

func isLiteral(o Obj) bool {
	switch {
	case o == Nil, o == True, o == False, IsNumber(o):
		return true
	case IsString(o):
		return !strings.ContainsAny(GetString(o), "\"\\")
	case IsSymbol(o):
		name := GetSymbol(o)
		// Uppercase-initial symbols are KL variables. Anything else in
		// argument position is symbol data.
		return name != "" && !(name[0] >= 'A' && name[0] <= 'Z') && !strings.ContainsAny(name, "\"()|; ")
	}
	ok, p := isPair(o)
	if !ok {
		return false
	}
	if p.car == MakeSymbol("cons") {
		args := ListToSlice(p.cdr)
		return len(args) == 2 && allLiteral(args)
	}
	if p.car == MakeSymbol("intern") {
		args := ListToSlice(p.cdr)
		return len(args) == 1 && IsString(args[0]) && isLiteral(args[0])
	}
	return false
}

// klExprString prints a literal expression back out as KL source.
func klExprString(o Obj) string {
	switch {
	case o == Nil:
		return "()"
	case IsString(o):
		return `"` + GetString(o) + `"`
	}
	if ok, p := isPair(o); ok {
		parts := []string{klExprString(p.car)}
		for cur := p.cdr; cur != Nil; {
			okc, cp := isPair(cur)
			if !okc {
				break
			}
			parts = append(parts, klExprString(cp.car))
			cur = cp.cdr
		}
		return "(" + strings.Join(parts, " ") + ")"
	}
	return ObjString(o)
}

// ---- seeded random draws ----------------------------------------------------

const randomPerRow = 24

func rowSeed(row string) int64 {
	h := fnv.New64a()
	h.Write([]byte("equiv:" + row))
	return int64(h.Sum64() & 0x7fffffffffffffff)
}

type gen struct{ r *rand.Rand }

func (g *gen) pick(xs ...string) string { return xs[g.r.Intn(len(xs))] }

func (g *gen) number() string {
	switch g.r.Intn(6) {
	case 0:
		return strconv.Itoa(g.r.Intn(140) - 10)
	case 1:
		return strconv.Itoa(g.r.Intn(5))
	case 2:
		return strconv.FormatInt(int64(g.r.Int63n(1<<40))-(1<<39), 10)
	case 3:
		return floatLit(float64(g.r.Intn(130)) + 0.5)
	default:
		// Drawn as an integer number of thousandths so the literal is the
		// same on every architecture: a Float64()*200-100 draw was fused
		// into an FMA on arm64 and printed one ulp away from amd64's value.
		return floatLit(float64(g.r.Intn(200001)-100000) / 1000)
	}
}

func floatLit(f float64) string {
	s := strconv.FormatFloat(f, 'f', -1, 64)
	if !strings.ContainsAny(s, ".e") {
		s += ".0"
	}
	return s
}

const symAlpha = "abcxyzABC*-._?"

func (g *gen) rawName(n int, first string) string {
	var b strings.Builder
	for i := 0; i < n; i++ {
		src := symAlpha + "0123456789"
		if i == 0 && first != "" {
			src = first
		}
		b.WriteByte(src[g.r.Intn(len(src))])
	}
	return b.String()
}

func (g *gen) str() string {
	switch g.r.Intn(5) {
	case 0:
		return `""`
	case 1:
		return `"` + g.pick("λ", "aλ", "λ1", "€x", "日本") + `"`
	default:
		return `"` + g.rawName(g.r.Intn(6)+1, "") + `"`
	}
}

func (g *gen) symbol() string {
	switch g.r.Intn(4) {
	case 0:
		return g.pick("reverse", "map", "fail", "put", "shen.app", "equiv.undefined-fn", "arity", "fn", "not", "thaw", "vector", "*property-vector*")
	case 1:
		return `(intern "` + g.rawName(g.r.Intn(4)+1, "") + `")`
	default:
		return g.rawName(g.r.Intn(5)+1, symAlpha)
	}
}

func (g *gen) atom() string {
	switch g.r.Intn(6) {
	case 0, 1:
		return g.number()
	case 2:
		return g.str()
	case 3:
		return g.symbol()
	case 4:
		return g.pick("true", "false", "()")
	default:
		return g.pick("(vector 0)", "(vector 1)", "(@p 1 2)", "(absvector 2)")
	}
}

func (g *gen) list(elem func() string, maxLen int) string {
	n := g.r.Intn(maxLen + 1)
	s := "()"
	for i := 0; i < n; i++ {
		s = "(cons " + elem() + " " + s + ")"
	}
	return s
}

func (g *gen) any(depth int) string {
	if depth > 0 && g.r.Intn(3) == 0 {
		return g.list(func() string { return g.any(depth - 1) }, 3)
	}
	return g.atom()
}

func (g *gen) value(k kind) string {
	// One draw in five ignores the declared kind, so that the error paths get
	// cases too.
	if k != kAny && k != kPropKey && g.r.Intn(5) == 0 {
		return g.any(1)
	}
	switch k {
	case kList:
		return g.list(func() string { return g.any(1) }, 6)
	case kString:
		return g.str()
	case kNumber:
		return g.number()
	case kSymbol:
		return g.symbol()
	case kVector:
		switch g.r.Intn(4) {
		case 0:
			return "(vector " + strconv.Itoa(g.r.Intn(4)) + ")"
		case 1:
			return "(absvector " + strconv.Itoa(g.r.Intn(4)) + ")"
		case 2:
			return "(@p " + g.atom() + " " + g.atom() + ")"
		default:
			return "(let V (vector 3) (do (vector-> V " + strconv.Itoa(g.r.Intn(3)+1) + " " + g.atom() + ") V))"
		}
	case kThunk:
		return g.pick("(freeze "+g.any(1)+")", "(freeze (equiv.note "+g.atom()+"))", "(freeze (equiv.note "+g.atom()+"))")
	case kFunc:
		return g.pick("(lambda X (cons "+g.atom()+" X))", "(lambda X (equiv.note X))", "(equiv.native-of equiv.note)", "(lambda X (equiv.note (cons X "+g.atom()+")))")
	case kAlist:
		return g.list(func() string {
			if g.r.Intn(6) == 0 {
				return g.atom()
			}
			return "(cons " + g.atom() + " " + g.any(1) + ")"
		}, 5)
	case kNumList:
		return g.list(g.number, 6)
	case kPropKey:
		return g.pick(g.atom(), g.rawName(1, "abcxyz"), g.rawName(1, "abcxyz"))
	}
	return g.any(2)
}

func randomCases(row string, sig []kind) []EquivCase {
	g := &gen{r: rand.New(rand.NewSource(rowSeed(row)))}
	var out []EquivCase
	for i := 0; i < randomPerRow; i++ {
		c := EquivCase{Name: fmt.Sprintf("random:%d", i)}
		if propertyRows[row] {
			size := g.pick("1", "3", "7")
			c.Setup = append(c.Setup, "(set equiv.*pv* (vector "+size+"))")
			for n := g.r.Intn(4); n > 0; n-- {
				c.Setup = append(c.Setup, "(put "+g.value(kPropKey)+" "+g.value(kPropKey)+" "+g.any(1)+" (value equiv.*pv*))")
			}
		}
		for _, k := range sig {
			c.Args = append(c.Args, g.value(k))
		}
		out = append(out, c)
	}
	return out
}

// ---- shared kernel boot -----------------------------------------------------

var (
	equivBootOnce sync.Once
	equivBootErr  error
	equivKernel   string
	equivHelpers  []string
	equivCF       ControlFlow
)

// bootEquivKernel boots the KLambda kernel into this test process once, and
// installs the equiv.* copies of the rebound functions and of their helpers.
func bootEquivKernel(t *testing.T, names []string) *ControlFlow {
	t.Helper()
	equivBootOnce.Do(func() {
		equivKernel, equivBootErr = FindKernelDir(".")
		if equivBootErr != nil {
			return
		}
		if equivBootErr = BootKernel(&equivCF, equivKernel); equivBootErr != nil {
			return
		}
		equivHelpers, equivBootErr = InstallEquivDefinitions(&equivCF, equivKernel, names)
	})
	if equivBootErr != nil {
		t.Fatal(equivBootErr)
	}
	return &equivCF
}

// rowNotes is prose attached to rows whose disagreement has a known origin.
// It is data for the reader of equiv.json, not an excuse: the row stays
// unverified.
var rowNotes = map[string]string{
	"fail": "The KL side returns shen.fail!, as the vendored kernel/klambda/sys.kl says (defun fail () shen.fail!) " +
		"and kernel/sources/sys.shen says (define fail -> fail!). " +
		"The native (nativeFail in kl/kernelfast.go) and the port's compiled kernel (cmd/shen/sys.go) return the symbol ... instead. " +
		"The harness leaves fail unrenamed inside the other KL copies, so every other row is judged with the port's filler " +
		"and this row alone carries the divergence.",
}

func defunArity(t *testing.T, dir, name string) int {
	t.Helper()
	defuns, err := KernelDefuns(dir)
	if err != nil {
		t.Fatal(err)
	}
	d, ok := defuns[name]
	if !ok {
		t.Fatalf("%s: no kernel defun", name)
	}
	return len(ListToSlice(Car(Cdr(Cdr(d)))))
}

func equivHarnessDoc() EquivHarness {
	return EquivHarness{
		KLSide: "equiv.NAME is the kernel's own (defun NAME …) from kernel/klambda, run as KL bytecode. " +
			"Every kernel defun its body reaches is loaded the same way and is listed in kl_helpers, so no native under test runs on this side. " +
			"Two names stay native: hash, the port's native hash, which is installed at boot before any property vector exists, " +
			"and fail, for the reason the fail row gives.",
		NativeSide: "NAME as bound after kl.InstallKernelFast. That is the Go function named in the native field.",
		Inputs: "The cases for this row. Each case is evaluated once per side, with the natives installed, in the order setup, args, the call, observe. " +
			"Both sides must return equal values, or raise errors with identical text, and every observe value must be equal. " +
			"Values are compared with the kernel's =, extended to functions: two functions are equal when they are the same object, " +
			"or when both are KL closures of the same arity.",
		Verified: "True when every case agreed at the time the table was generated on this port. " +
			"A false row records the first disagreeing case in reason, as \"case=NAME kl=OUT native=OUT\". " +
			"OUT is the value printed as KL source, with a vector as #vector[limit slot…] and never-written slots as #unset, " +
			"or error(\"text\") with the text Go-quoted, followed by one observe[i]=OUT per observe form. " +
			"The kind of disagreement is in disagreement: raise means one side raised and the other returned; " +
			"value means both returned and the values differ; error-text means both raised and the messages differ; " +
			"effect means an observed value differs; harness means a form of the case itself failed. " +
			"When the origin of the divergence is known, note explains it in prose.",
		Cases:      "The number of cases in inputs.",
		Source:     "Provenance of the native, as kl/FILE.go:GoIdentifier naming the Go function bound to kernel_fn. It is not a line location.",
		Effects:    effectsDoc,
		Vocabulary: EquivVocabulary,
		KLHelpers:  equivHelpers,
	}
}

// ---- the test ----------------------------------------------------------------

const equivJSONPath = "equiv.json"

func buildEquivTable(t *testing.T) EquivTable {
	t.Helper()
	bindings := parseInstallKernelFast(t)
	if len(bindings) == 0 {
		t.Fatal("no rebindings found in InstallKernelFast")
	}
	names := make([]string, 0, len(bindings))
	seen := map[string]bool{}
	for _, b := range bindings {
		if seen[b.kernelFn] {
			t.Fatalf("%s is rebound twice in InstallKernelFast", b.kernelFn)
		}
		seen[b.kernelFn] = true
		names = append(names, b.kernelFn)
		if _, ok := rowSignature[b.kernelFn]; !ok {
			t.Errorf("InstallKernelFast rebinds %q but rowSignature has no entry for it: add its argument kinds (and effects) to equiv_test.go", b.kernelFn)
		}
	}
	for n := range rowSignature {
		if !seen[n] {
			t.Errorf("rowSignature lists %q but InstallKernelFast no longer rebinds it", n)
		}
	}
	for n := range rowEffects {
		if !seen[n] {
			t.Errorf("rowEffects lists %q but InstallKernelFast no longer rebinds it", n)
		}
	}
	for n := range explicitCases {
		if !seen[n] {
			t.Errorf("explicitCases lists %q but InstallKernelFast no longer rebinds it", n)
		}
	}
	if t.Failed() {
		t.FailNow()
	}

	e := bootEquivKernel(t, names)
	table := EquivTable{Port: "shen-go", Kernel: "S42", Harness: equivHarnessDoc()}
	for _, b := range bindings {
		sig := rowSignature[b.kernelFn]
		if len(sig) != b.arity {
			t.Fatalf("%s: rowSignature has %d kinds, native arity is %d", b.kernelFn, len(sig), b.arity)
		}
		// The kernel's arity table covers the user-facing functions. Internal
		// shen.* helpers are absent from it, which KernelArity reports as -1,
		// and are checked against their defun's parameter list instead.
		if ka := KernelArity(e, b.kernelFn); ka != -1 && ka != b.arity {
			t.Errorf("%s: native arity %d, kernel arity table says %d", b.kernelFn, b.arity, ka)
		} else if ka == -1 {
			if da := defunArity(t, equivKernel, b.kernelFn); da != b.arity {
				t.Errorf("%s: native arity %d, kernel defun has %d parameters", b.kernelFn, b.arity, da)
			}
		}
		row := EquivRow{KernelFn: b.kernelFn, Native: b.native, Arity: b.arity, Source: b.source, Effects: rowEffects[b.kernelFn]}
		// For a property row the signature covers key, attr and value. The
		// vector argument is appended by finishCase.
		genSig := sig
		if propertyRows[b.kernelFn] {
			genSig = sig[:len(sig)-1]
		}
		var cases []EquivCase
		for _, c := range explicitCases[b.kernelFn] {
			c.Name = "explicit:" + c.Name
			cases = append(cases, c)
		}
		cases = append(cases, boundaryCases(b.kernelFn, genSig)...)
		cases = append(cases, kernelUsageCases(t, equivKernel, b.kernelFn, b.arity)...)
		cases = append(cases, randomCases(b.kernelFn, genSig)...)
		row.Verified = true
		if os.Getenv("EQUIV_TRACE") != "" {
			fmt.Println("equiv-row:", b.kernelFn)
		}
		for _, c := range cases {
			c = finishCase(b.kernelFn, c)
			row.Inputs = append(row.Inputs, c)
			if m := RunEquivCase(e, b.kernelFn, c); m != nil {
				if m.Kind == MismatchHarness {
					t.Errorf("equiv %s: harness cannot run %s", b.kernelFn, m)
				}
				if row.Verified {
					row.Verified = false
					row.Reason = m.String()
					row.Disagreement = string(m.Kind)
					row.Note = rowNotes[b.kernelFn]
				}
				t.Logf("equiv %s FAIL %s", b.kernelFn, m)
			}
		}
		row.Cases = len(row.Inputs)
		table.Rows = append(table.Rows, row)
	}
	return table
}

// TestEquivTable runs every differential case and checks that kl/equiv.json is
// the table those runs produce. It fails when the committed table drifts from
// InstallKernelFast, that is when a rebinding is added or removed or a verdict
// changes. Regenerate with EQUIV_WRITE=1.
func TestEquivTable(t *testing.T) {
	table := buildEquivTable(t)
	got, err := json.MarshalIndent(table, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	got = append(got, '\n')

	verified, total := 0, 0
	for _, r := range table.Rows {
		total += r.Cases
		if r.Verified {
			verified++
		} else {
			t.Logf("unverified: %s: %s", r.KernelFn, r.Reason)
		}
	}
	t.Logf("equiv: %d rows, %d verified, %d cases", len(table.Rows), verified, total)

	if os.Getenv("EQUIV_WRITE") != "" {
		if err := os.WriteFile(equivJSONPath, got, 0o644); err != nil {
			t.Fatal(err)
		}
		t.Logf("wrote %s", equivJSONPath)
		return
	}
	want, err := os.ReadFile(equivJSONPath)
	if err != nil {
		t.Fatalf("%v (run with EQUIV_WRITE=1 to generate)", err)
	}
	if !bytes.Equal(want, got) {
		var old EquivTable
		if json.Unmarshal(want, &old) == nil {
			t.Errorf("%s drifted from InstallKernelFast: %s", equivJSONPath, describeDrift(old, table))
		}
		t.Fatalf("%s is stale; regenerate with: EQUIV_WRITE=1 go test ./kl -run TestEquivTable", equivJSONPath)
	}
}

func describeDrift(old, cur EquivTable) string {
	index := func(t EquivTable) map[string]EquivRow {
		m := map[string]EquivRow{}
		for _, r := range t.Rows {
			m[r.KernelFn] = r
		}
		return m
	}
	o, c := index(old), index(cur)
	var msgs []string
	for name, r := range c {
		or, ok := o[name]
		switch {
		case !ok:
			msgs = append(msgs, name+": new row")
		case or.Verified != r.Verified:
			msgs = append(msgs, fmt.Sprintf("%s: verified %v -> %v (%s)", name, or.Verified, r.Verified, r.Reason))
		case or.Cases != r.Cases:
			msgs = append(msgs, fmt.Sprintf("%s: cases %d -> %d", name, or.Cases, r.Cases))
		case or.Native != r.Native || or.Arity != r.Arity:
			msgs = append(msgs, name+": native/arity changed")
		case or.Reason != r.Reason || or.Note != r.Note || or.Disagreement != r.Disagreement:
			msgs = append(msgs, name+": reason/note changed")
		case !reflect.DeepEqual(or.Effects, r.Effects):
			msgs = append(msgs, name+": effects changed")
		default:
			for i := range r.Inputs {
				if !reflect.DeepEqual(or.Inputs[i], r.Inputs[i]) {
					msgs = append(msgs, fmt.Sprintf("%s: input %s differs: committed %v, generated %v", name, r.Inputs[i].Name, or.Inputs[i], r.Inputs[i]))
					break
				}
			}
		}
	}
	for name := range o {
		if _, ok := c[name]; !ok {
			msgs = append(msgs, name+": row removed")
		}
	}
	sort.Strings(msgs)
	if len(msgs) == 0 {
		if !reflect.DeepEqual(old.Harness, cur.Harness) || old.Port != cur.Port || old.Kernel != cur.Kernel {
			return "(the harness/port/kernel header differs)"
		}
		return "(files differ but no row does: formatting?)"
	}
	return strings.Join(msgs, "; ")
}

// TestEquivJSONReplays checks that the committed table replays. Running the
// cases stored in kl/equiv.json must give each row the verdict the JSON
// records. This is what `kl equiv-check` does on another host.
func TestEquivJSONReplays(t *testing.T) {
	data, err := os.ReadFile(equivJSONPath)
	if err != nil {
		t.Fatalf("%v (run with EQUIV_WRITE=1 to generate)", err)
	}
	var table EquivTable
	if err := json.Unmarshal(data, &table); err != nil {
		t.Fatal(err)
	}
	names := make([]string, 0, len(table.Rows))
	for _, r := range table.Rows {
		names = append(names, r.KernelFn)
	}
	e := bootEquivKernel(t, names)
	for _, r := range table.Rows {
		var first *EquivMismatch
		for _, c := range r.Inputs {
			if m := RunEquivCase(e, r.KernelFn, c); m != nil && first == nil {
				first = m
			}
		}
		if (first == nil) != r.Verified {
			t.Errorf("%s: JSON says verified=%v, replay says %v", r.KernelFn, r.Verified, first)
		}
	}
}
