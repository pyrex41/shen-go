// Command yggdrasil-build is the Go stage-2 builder for Yggdrasil (the Shen
// tree-shaker). It takes a Yggdrasil shaken directory (kernel.kl + user .kl
// files + yggdrasil.manifest.txt) and produces a self-contained Go module
// that `go build`s into a single static, cross-compilable binary.
//
//	yggdrasil-build [-shen-go DIR] <shaken-dir> <outdir>
//	cd <outdir> && go build -o prog .
//
// How it works (architecture "true codegen"): the builder boots the full
// shen-go kernel plus the Shen-side bytecode compiler (src/compiler.shen) on
// the kl VM in-process, runs compile-file over the shaken kernel and user
// files (split into chunks), and translates each bytecode-IR chunk to Go with
// the shared codegen package — the same pipeline that generated cmd/shen
// itself. The emitted module contains one 0-arity thunk per chunk plus a
// main.go that runs kernel chunks in order, calls the manifest's init
// function (shen.initialise), then runs the user chunks (whose toplevel
// non-defun forms execute in source order, after init, as the Yggdrasil
// builder contract requires).
//
// The generated module imports github.com/tiancaiamao/shen-go/kl via a
// `replace` directive pointing at this source tree, so the tree must remain
// available at `go build` time (but not at runtime — the binary is static).
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/tiancaiamao/shen-go/codegen"
	"github.com/tiancaiamao/shen-go/kl"
)

// kernelLoadOrder is the canonical load order of the full shen-go kernel,
// mirrored from compiled/precompile.kl. The builder needs the full kernel
// only to host the compiler; the program being built loads the *shaken*
// kernel from the manifest instead.
var kernelLoadOrder = []string{
	"toplevel.kl", "core.kl", "sys.kl", "sequent.kl", "yacc.kl",
	"reader.kl", "prolog.kl", "track.kl", "load.kl", "writer.kl",
	"macros.kl", "declarations.kl", "t-star.kl", "types.kl",
	"dict.kl", "init.kl",
}

// chunkTarget is the rough upper bound on the KL source size compiled into a
// single Go function. compile-file wraps each chunk in one nested-do thunk
// and bc->go emits it as one Go function; keeping chunks bounded keeps both
// the KL-side compile and the Go compile of the generated function tractable.
// The KL-side compile cost grows superlinearly with chunk size (the IR is
// printed through the kernel's string printer), so smaller chunks build
// faster overall. Tunable via -chunk-kb.
var chunkTarget = 24 * 1024

type manifest struct {
	kernel    string
	init      string
	users     []string
	needsEval bool
}

func parseManifest(dir string) (*manifest, error) {
	data, err := os.ReadFile(filepath.Join(dir, "yggdrasil.manifest.txt"))
	if err != nil {
		return nil, err
	}
	m := &manifest{}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, val, ok := strings.Cut(line, "=")
		if !ok {
			return nil, fmt.Errorf("malformed manifest line: %q", line)
		}
		switch key {
		case "kernel":
			m.kernel = val
		case "init":
			m.init = val
		case "user":
			m.users = append(m.users, val)
		case "needs-eval":
			m.needsEval = val == "true"
		case "manifest-version", "kernel-version", "primitive":
			// Informational. The kl runtime provides every KL primitive
			// natively (registered in the kl package's init()), so the
			// primitive= list needs no per-program action here.
		default:
			return nil, fmt.Errorf("unknown manifest key %q", key)
		}
	}
	if m.kernel == "" {
		return nil, fmt.Errorf("manifest has no kernel= entry")
	}
	if m.init == "" {
		return nil, fmt.Errorf("manifest has no init= entry")
	}
	return m, nil
}

// splitForms splits KL source text into its top-level forms. KL strings have
// no escape sequences, so a paren/string scanner is exact.
func splitForms(src []byte) ([][]byte, error) {
	var forms [][]byte
	depth, start := 0, -1
	inStr := false
	for i := 0; i < len(src); i++ {
		switch c := src[i]; {
		case inStr:
			if c == '"' {
				inStr = false
			}
		case c == '"':
			inStr = true
		case c == '(':
			if depth == 0 {
				start = i
			}
			depth++
		case c == ')':
			depth--
			if depth < 0 {
				return nil, fmt.Errorf("unbalanced ')' at byte %d", i)
			}
			if depth == 0 {
				forms = append(forms, src[start:i+1])
				start = -1
			}
		}
	}
	if depth != 0 || inStr {
		return nil, fmt.Errorf("unterminated form (depth=%d inString=%v)", depth, inStr)
	}
	return forms, nil
}

// chunkForms groups consecutive top-level forms into chunks of roughly
// chunkTarget bytes, preserving order.
func chunkForms(forms [][]byte) [][]byte {
	var chunks [][]byte
	var cur []byte
	for _, f := range forms {
		if len(cur) > 0 && len(cur)+len(f) > chunkTarget {
			chunks = append(chunks, cur)
			cur = nil
		}
		cur = append(cur, f...)
		cur = append(cur, '\n')
	}
	if len(cur) > 0 {
		chunks = append(chunks, cur)
	}
	return chunks
}

// evalKL parses src as KL forms and evaluates each on the VM, failing fast on
// the first error and returning the last form's value.
func evalKL(e *kl.ControlFlow, src string) error {
	_, err := evalKLValue(e, src)
	return err
}

func evalKLValue(e *kl.ControlFlow, src string) (kl.Obj, error) {
	r := kl.NewSexpReader(strings.NewReader(src), false)
	last := kl.Nil
	for {
		form, err := r.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return last, nil
			}
			return kl.Nil, fmt.Errorf("parse %q: %v", src, err)
		}
		last = kl.Eval(e, form)
		if kl.IsError(last) {
			return kl.Nil, fmt.Errorf("eval %s: %s", kl.ObjString(form), kl.GetString(kl.PrimErrorToString(last)))
		}
	}
}

func klPath(p string) string {
	if strings.ContainsAny(p, `"\`) {
		fatal("path %q cannot be written as a KL string literal", p)
	}
	return `"` + p + `"`
}

func fatal(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "yggdrasil-build: "+format+"\n", args...)
	os.Exit(1)
}

// findShenGoRoot walks up from dir looking for the shen-go module root.
func findShenGoRoot(dir string) (string, error) {
	dir, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}
	for {
		data, err := os.ReadFile(filepath.Join(dir, "go.mod"))
		if err == nil && strings.Contains(string(data), "module github.com/tiancaiamao/shen-go") {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("shen-go source tree not found above %s (use -shen-go)", dir)
		}
		dir = parent
	}
}

type unit struct {
	name   string // Go identifier of the chunk thunk
	goFile string // file name within the outdir
	src    []byte // KL source of the chunk
}

type fnArity struct {
	name  string
	arity int
}

// scanDefunArities returns the (name, arity) of every top-level defun in a KL
// source, in order. Mirrors cmd/kl's emit-arities / the .fns manifest format.
func scanDefunArities(src []byte) ([]fnArity, error) {
	var out []fnArity
	r := kl.NewSexpReader(strings.NewReader(string(src)), false)
	symDefun := kl.MakeSymbol("defun")
	for {
		form, err := r.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return out, nil
			}
			return nil, err
		}
		if kl.PrimIsPair(form) == kl.False || kl.Car(form) != symDefun {
			continue
		}
		nameObj := kl.Cadr(form)
		if !kl.IsSymbol(nameObj) {
			continue
		}
		arity := 0
		for p := kl.Car(kl.Cdr(kl.Cdr(form))); p != kl.Nil; p = kl.Cdr(p) {
			arity++
		}
		out = append(out, fnArity{name: kl.GetSymbol(nameObj), arity: arity})
	}
}

func main() {
	shenGoFlag := flag.String("shen-go", "", "path to the shen-go source tree (default: walk up from cwd)")
	keepWork := flag.Bool("keep-work", false, "keep the _work directory with intermediate .kl/.bc chunks")
	chunkKB := flag.Int("chunk-kb", chunkTarget/1024, "target KL chunk size in KB per generated Go function")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: yggdrasil-build [-shen-go DIR] <shaken-dir> <outdir>\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(2)
	}
	if *chunkKB > 0 {
		chunkTarget = *chunkKB * 1024
	}

	shakenDir, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		fatal("%v", err)
	}
	outDir, err := filepath.Abs(flag.Arg(1))
	if err != nil {
		fatal("%v", err)
	}

	root := *shenGoFlag
	if root == "" {
		root, err = findShenGoRoot(".")
		if err != nil {
			fatal("%v", err)
		}
	}
	root, err = filepath.Abs(root)
	if err != nil {
		fatal("%v", err)
	}

	m, err := parseManifest(shakenDir)
	if err != nil {
		fatal("manifest: %v", err)
	}

	if err := os.MkdirAll(outDir, 0755); err != nil {
		fatal("%v", err)
	}
	workDir := filepath.Join(outDir, "_work")
	if err := os.MkdirAll(workDir, 0755); err != nil {
		fatal("%v", err)
	}

	// Boot the compiler image: full kernel + shen.initialise + compiler.shen.
	t0 := time.Now()
	var e kl.ControlFlow
	for _, f := range kernelLoadOrder {
		p := filepath.Join(root, "kernel", "klambda", f)
		if err := evalKL(&e, "(load-file "+klPath(p)+")"); err != nil {
			fatal("boot kernel: %v", err)
		}
	}
	// Mirror cmd/shen: native hash, swapped in after the kernel binds its
	// interpreted version and before shen.initialise builds dictionaries.
	kl.BindSymbolFunc(kl.MakeSymbol("hash"), kl.MakePrimitive("hash", 2, kl.PrimHash))
	if err := evalKL(&e, "(shen.initialise)"); err != nil {
		fatal("shen.initialise: %v", err)
	}
	if err := evalKL(&e, "(load "+klPath(filepath.Join(root, "src", "compiler.shen"))+")"); err != nil {
		fatal("load compiler.shen: %v", err)
	}
	if err := evalKL(&e, "(set *maximum-print-sequence-size* 100000)"); err != nil {
		fatal("%v", err)
	}
	bootDur := time.Since(t0)

	// Register every user defun's arity in the compiler image before
	// compiling: with arities known, saturated calls compile to direct calls
	// instead of runtime (fn F) lookups. The same pairs are replayed at
	// runtime by the generated main.go (the static-build mirror of the
	// plugin path's .fns manifest, cmd/shen/plugin.go).
	var userArities []fnArity
	for _, uf := range m.users {
		src, err := os.ReadFile(filepath.Join(shakenDir, uf))
		if err != nil {
			fatal("%v", err)
		}
		fas, err := scanDefunArities(src)
		if err != nil {
			fatal("scan %s: %v", uf, err)
		}
		userArities = append(userArities, fas...)
	}
	symStoreArity := kl.MakeSymbol("shen.store-arity")
	for _, fa := range userArities {
		form := kl.Cons(symStoreArity, kl.Cons(kl.MakeSymbol(fa.name), kl.Cons(kl.MakeInteger(fa.arity), kl.Nil)))
		if res := kl.Eval(&e, form); kl.IsError(res) {
			fmt.Fprintf(os.Stderr, "yggdrasil-build: warning: store-arity %s/%d: %s\n",
				fa.name, fa.arity, kl.GetString(kl.PrimErrorToString(res)))
		}
	}

	// Carve the inputs into compilation units.
	var kernelUnits, userUnits []unit
	makeUnits := func(klFile, prefix string) []unit {
		src, err := os.ReadFile(filepath.Join(shakenDir, klFile))
		if err != nil {
			fatal("%v", err)
		}
		forms, err := splitForms(src)
		if err != nil {
			fatal("%s: %v", klFile, err)
		}
		chunks := chunkForms(forms)
		units := make([]unit, len(chunks))
		for i, c := range chunks {
			units[i] = unit{
				name:   fmt.Sprintf("%sChunk%d", prefix, i),
				goFile: fmt.Sprintf("%s_%02d.go", strings.ToLower(prefix), i),
				src:    c,
			}
		}
		return units
	}
	kernelUnits = makeUnits(m.kernel, "Kernel")
	for ui, uf := range m.users {
		base := strings.TrimSuffix(filepath.Base(uf), ".kl")
		base = strings.Map(func(r rune) rune {
			if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
				return r
			}
			return -1
		}, base)
		userUnits = append(userUnits, makeUnits(uf, fmt.Sprintf("User%d%s", ui, strings.Title(base)))...)
	}

	// Compile every unit: KL chunk -> bytecode IR (compile-file on the VM)
	// -> Go (codegen). One shared CodeGenerator so symbol declarations are
	// emitted exactly once for the whole package.
	cg := codegen.New()
	t1 := time.Now()
	compileUnit := func(u unit) {
		klFile := filepath.Join(workDir, u.name+".kl")
		if err := os.WriteFile(klFile, u.src, 0644); err != nil {
			fatal("%v", err)
		}
		// Same pipeline as compile-file (src/compiler.shen) minus the
		// print-to-file/reparse round-trip: take the bytecode IR off the VM
		// as an object and hand it straight to the Go translator. The
		// trailing shen.skip keeps the do-wrap n-ary even for single-form
		// chunks: (do X) with one argument is not special-formed by the
		// compiler and would compile to a unary call of the do *function*
		// (arity 2), which traps at runtime.
		bc, err := evalKLValue(&e, "(codegen (macroexpand (cons do (append (read-file "+klPath(klFile)+") (cons shen.skip ())))))")
		if err != nil {
			fatal("compile %s: %v", u.name, err)
		}
		out, err := os.Create(filepath.Join(outDir, u.goFile))
		if err != nil {
			fatal("%v", err)
		}
		defer out.Close()
		if err := cg.HandleBodyObj(bc, u.name, out); err != nil {
			fatal("codegen %s: %v", u.name, err)
		}
		fmt.Printf("  %-24s %6.1f KB kl -> %s\n", u.name, float64(len(u.src))/1024, u.goFile)
	}
	fmt.Printf("compiling %d kernel + %d user chunks\n", len(kernelUnits), len(userUnits))
	for _, u := range kernelUnits {
		compileUnit(u)
	}
	for _, u := range userUnits {
		compileUnit(u)
	}
	compileDur := time.Since(t1)

	// symbols.go: the symXXX vars referenced by every chunk, sorted for
	// reproducible output.
	var symBuf strings.Builder
	cg.HandleSymbol(&symBuf)
	symLines := strings.Split(strings.TrimRight(symBuf.String(), "\n"), "\n")
	sort.Strings(symLines)
	symSrc := "// Code generated by yggdrasil-build. DO NOT EDIT.\npackage main\n\nimport . \"github.com/tiancaiamao/shen-go/kl\"\n\n" +
		strings.Join(symLines, "\n") + "\n"
	if err := os.WriteFile(filepath.Join(outDir, "symbols.go"), []byte(symSrc), 0644); err != nil {
		fatal("%v", err)
	}

	// main.go + go.mod for the generated module.
	if err := os.WriteFile(filepath.Join(outDir, "main.go"), []byte(genMain(m, kernelUnits, userUnits, userArities)), 0644); err != nil {
		fatal("%v", err)
	}
	gomod := fmt.Sprintf("module yggdrasil.local/%s\n\ngo 1.25\n\nrequire github.com/tiancaiamao/shen-go v0.0.0\n\nreplace github.com/tiancaiamao/shen-go => %s\n",
		moduleName(m), root)
	if err := os.WriteFile(filepath.Join(outDir, "go.mod"), []byte(gomod), 0644); err != nil {
		fatal("%v", err)
	}

	if !*keepWork {
		os.RemoveAll(workDir)
	}

	fmt.Printf("boot %.1fs, compile %.1fs\n", bootDur.Seconds(), compileDur.Seconds())
	fmt.Printf("wrote %s\nnext: cd %s && go build -o prog .\n", outDir, outDir)
}

func moduleName(m *manifest) string {
	if len(m.users) > 0 {
		base := strings.TrimSuffix(filepath.Base(m.users[0]), ".kl")
		if base != "" {
			return base
		}
	}
	return "prog"
}

func genMain(m *manifest, kernelUnits, userUnits []unit, userArities []fnArity) string {
	var b strings.Builder
	b.WriteString(`// Code generated by yggdrasil-build. DO NOT EDIT.
package main

import (
	"fmt"
	"os"

	. "github.com/tiancaiamao/shen-go/kl"
)

// Bound at startup; the generated chunks reference these for defun binding
// and trap-error (see codegen package docs).
var ns2_1set Obj
var try_1catch Obj

`)
	b.WriteString("var kernelChunks = []Obj{")
	for i, u := range kernelUnits {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(u.name)
	}
	b.WriteString("}\n\nvar userChunks = []Obj{")
	for i, u := range userUnits {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(u.name)
	}
	b.WriteString("}\n\n")
	b.WriteString(`// userArities mirrors the plugin path's .fns manifest: every user defun's
// arity, replayed via shen.store-arity after init so the kernel's (fn F)
// lookup and partial application resolve compiled user functions.
var userArities = []struct {
	name  string
	arity int
}{
`)
	for _, fa := range userArities {
		fmt.Fprintf(&b, "\t{%q, %d},\n", fa.name, fa.arity)
	}
	b.WriteString("}\n\n")
	b.WriteString(`func fail(name string, x Obj) {
	fmt.Fprintf(os.Stderr, "yggdrasil: %s failed: %s\n", name, GetString(PrimErrorToString(x)))
	os.Exit(1)
}

// run invokes a 0-arity thunk, turning both returned and raised KL errors
// into a clean nonzero exit.
func run(e *ControlFlow, name string, f Obj) {
	defer func() {
		if r := recover(); r != nil {
			if x, ok := r.(Obj); ok && IsError(x) {
				fail(name, x)
			}
			panic(r)
		}
	}()
	if res := Call(e, f); IsError(res) {
		fail(name, res)
	}
}

func main() {
	ns2_1set = PrimFunc(MakeSymbol("defun"))
	try_1catch = PrimFunc(MakeSymbol("try-catch"))

	var e ControlFlow
	for i, c := range kernelChunks {
		run(&e, fmt.Sprintf("kernel chunk %d", i), c)
	}
	// Mirror cmd/shen: swap in the native hash after the kernel binds its
	// interpreted version and before the init function builds dictionaries
	// (every dict must be created and queried with the same hash).
	BindSymbolFunc(MakeSymbol("hash"), MakePrimitive("hash", 2, PrimHash))
`)
	fmt.Fprintf(&b, "\trun(&e, %q, PrimFunc(MakeSymbol(%q)))\n", m.init, m.init)
	b.WriteString(`	storeArity := MakeSymbol("shen.store-arity")
	for _, fa := range userArities {
		if res := Eval(&e, Cons(storeArity, Cons(MakeSymbol(fa.name), Cons(MakeInteger(fa.arity), Nil)))); IsError(res) {
			fmt.Fprintf(os.Stderr, "yggdrasil: warning: store-arity %s/%d: %s\n",
				fa.name, fa.arity, GetString(PrimErrorToString(res)))
		}
	}
	for i, c := range userChunks {
		run(&e, fmt.Sprintf("user chunk %d", i), c)
	}
}
`)
	return b.String()
}
