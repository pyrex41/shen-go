// Package codegen translates the bytecode IR emitted by the Shen-side
// compiler (src/compiler.shen, via compile-file) into Go source that runs on
// the kl runtime's trampoline. It was extracted from cmd/kl so that other
// build tools (cmd/kl's bc->go native, cmd/yggdrasil-build) can share it.
//
// The emitted shape is `var <Export> = MakeNative(func(__e *ControlFlow)
// {...}, 0)` — a 0-arity module thunk that, when Called, defun-binds every
// function in the compiled file and evaluates its toplevel expressions in
// source order. The generated code references two package-level vars the
// host package must define and set before running any thunk:
//
//	var ns2_1set Obj   // = PrimFunc(MakeSymbol("defun"))
//	var try_1catch Obj // = PrimFunc(MakeSymbol("try-catch"))
package codegen

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/pyrex41/shen-go/kl"
)

// CodeGenerator accumulates the set of symbols referenced by the code it has
// generated so far. Share one CodeGenerator across all files destined for the
// same Go package and emit the symbol declarations exactly once (HandleSymbol)
// after the last file.
//
// ScmHead must stay the first field: cmd/kl exposes a *CodeGenerator to KL via
// kl.MakeRaw(&cg.ScmHead) and casts back with unsafe.Pointer.
type CodeGenerator struct {
	ScmHead int
	declare map[kl.Obj]struct{}
}

func New() *CodeGenerator {
	return &CodeGenerator{
		declare: make(map[kl.Obj]struct{}),
	}
}

// HandleBody reads one bytecode IR form from f and writes a complete Go file
// (package main, dot-importing kl) defining `var <export>` as the module thunk.
func (cg *CodeGenerator) HandleBody(f io.Reader, export string, out io.Writer) error {
	fmt.Fprintf(out, "package main\n\n")
	fmt.Fprintf(out, "import . \"github.com/pyrex41/shen-go/kl\"\n\n")
	fmt.Fprintf(out, `var %s = MakeNative(func(__e *ControlFlow) {
`, export)
	r := kl.NewSexpReader(f, false)
	bc, err := r.Read()
	if err != nil {
		fmt.Println("read bytecode error", err)
		return err
	}
	if err := cg.generateExpr(out, bc); err != nil {
		return err
	}
	fmt.Fprintf(out, "\n\n")
	fmt.Fprintf(out, "}, 0)\n\n")
	return nil
}

// HandleBodyObj is HandleBody for an in-memory bytecode IR object: it skips
// the print-to-file/reparse round-trip (printing megabyte IR through the
// kernel's string printer dominates compile time for large files).
func (cg *CodeGenerator) HandleBodyObj(bc kl.Obj, export string, out io.Writer) error {
	fmt.Fprintf(out, "package main\n\n")
	fmt.Fprintf(out, "import . \"github.com/pyrex41/shen-go/kl\"\n\n")
	fmt.Fprintf(out, `var %s = MakeNative(func(__e *ControlFlow) {
`, export)
	if err := cg.generateExpr(out, bc); err != nil {
		return err
	}
	fmt.Fprintf(out, "\n\n")
	fmt.Fprintf(out, "}, 0)\n\n")
	return nil
}

// HandleSymbol writes `var symXXX = MakeSymbol(...)` declarations for every
// symbol referenced by all files generated so far.
func (cg *CodeGenerator) HandleSymbol(out io.Writer) {
	for sym := range cg.declare {
		symStr := kl.GetSymbol(sym)
		symVar := "sym" + symbolAsVar(sym)
		fmt.Fprintf(out, "var %s = MakeSymbol(\"%s\")\n", symVar, symStr)
	}
}

func symbolAsVar(sym kl.Obj) string {
	str := kl.GetSymbol(sym)
	var buf bytes.Buffer
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '_':
			buf.WriteString("__")
		case '-':
			buf.WriteString("_1")
		case '?':
			buf.WriteString("_2")
		case '$':
			buf.WriteString("_3")
		case '.':
			buf.WriteString("_4")
		case '<':
			buf.WriteString("_5")
		case '>':
			buf.WriteString("_6")
		case '+':
			buf.WriteString("_7")
		case '@':
			buf.WriteString("_8")
		case '=':
			buf.WriteString("_a")
		case '!':
			buf.WriteString("_b")
		case '/':
			buf.WriteString("_c")
		case '*':
			buf.WriteString("_d")
		case '&':
			buf.WriteString("_e")
		case '%':
			buf.WriteString("_f")
		case '^':
			buf.WriteString("_g")
		case ':':
			buf.WriteString("_h")
		case '{':
			buf.WriteString("_i")
		case '}':
			buf.WriteString("_j")
		case ';':
			buf.WriteString("_k")
		case ',':
			buf.WriteString("_l")
		default:
			buf.WriteByte(str[i])
		}
	}
	return buf.String()
}

func (cg *CodeGenerator) generateExpr(w io.Writer, sexp kl.Obj) error {
	// fmt.Printf("handle %s ..\n", ObjString(sexp))
	if kl.IsSymbol(sexp) {
		fmt.Fprintf(w, "%s", symbolAsVar(sexp))
		return nil
	}
	kind := kl.GetSymbol(kl.Car(sexp))
	switch kind {
	case "$type":
		// Typed IR metadata is advisory.  Keep the expression itself in the
		// generated program so older artifacts and unproven annotations retain
		// exactly the legacy Obj semantics.
		// Source IR encodes [$type Annotation Expr].  Accept the transient
		// expression-first spelling too, so artifacts produced by older
		// compiler revisions remain consumable.
		annotation, expr := kl.Cadr(sexp), kl.Car(kl.Cdr(kl.Cdr(sexp)))
		if kl.IsSymbol(annotation) && (isTypeAnnotation(annotation) || !kl.IsSymbol(expr)) {
			return cg.generateExpr(w, expr)
		}
		return cg.generateExpr(w, annotation)
	case "block":
		return cg.generateBlock(w, sexp)
	case "<-":
		// (<- a b)
		a := kl.Car(kl.Cdr(sexp))
		b := kl.Car(kl.Cdr(kl.Cdr(sexp)))
		fmt.Fprintf(w, "%s = ", symbolAsVar(a))
		cg.generateExpr(w, b)
		fmt.Fprintln(w)
	case "<=":
		a := kl.Car(kl.Cdr(sexp))
		b := kl.Car(kl.Cdr(kl.Cdr(sexp)))
		fmt.Fprintf(w, "%s := ", symbolAsVar(a))
		cg.generateExpr(w, b)
		fmt.Fprintln(w)
	case "$global":
		sym := kl.Car(kl.Cdr(sexp))
		cg.declare[sym] = struct{}{}
		fmt.Fprintf(w, "PrimFunc(sym%s)", symbolAsVar(sym))
	case "$const":
		// (const Number)
		// (const ())
		// (const "xxx")
		c := kl.Cadr(sexp)
		if err := cg.generateConst(w, c); err != nil {
			return err
		}
	case "lambda":
		// (lambda (p1 p2 ...) ...)
		tmp := kl.Car(kl.Cdr(sexp))
		args := kl.ListToSlice(tmp)
		fmt.Fprintf(w, "MakeNative(func(__e *ControlFlow) {\n")
		for i, arg := range args {
			fmt.Fprintf(w, "%s := __e.Get(%d)\n", symbolAsVar(arg), i+1)
			fmt.Fprintf(w, "_ = %s\n", symbolAsVar(arg))
		}
		if err := cg.generateExpr(w, kl.Car(kl.Cdr(kl.Cdr(sexp)))); err != nil {
			return err
		}
		fmt.Fprintf(w, "}, %d)", len(args))
	case "if":
		// (if a b c)
		a := kl.Cadr(sexp)
		b := kl.Car(kl.Cdr(kl.Cdr(sexp)))
		c := kl.Car(kl.Cdr(kl.Cdr(kl.Cdr(sexp))))
		if err := cg.generateIfExpr(w, a, b, c); err != nil {
			return err
		}
	case "ignore": // (ignore xx)
		fmt.Fprintf(w, "_ = ")
		exp := kl.Car(kl.Cdr(sexp))
		cg.generateExpr(w, exp)
		fmt.Fprintln(w)
	case "var":
		// (var xxx)
		sym := kl.Car(kl.Cdr(sexp))
		fmt.Fprintf(w, "var %s Obj\n", symbolAsVar(sym))
	case "return":
		val := kl.Cadr(sexp)
		fmt.Fprintf(w, "__e.Return(")
		cg.generateExpr(w, val)
		fmt.Fprintf(w, ")\nreturn\n")
	case "call":
		// (call f a b c ...)
		ok, err := cg.primitiveCallOptimize(w, sexp, false)
		if err != nil {
			return err
		}
		if ok {
			return nil
		}

		fmt.Fprintf(w, "Call(__e, ")
		args := kl.ListToSlice(kl.Cdr(sexp))
		for i, arg := range args {
			if i != 0 {
				fmt.Fprintf(w, ", ")
			}
			if kl.IsSymbol(arg) {
				fmt.Fprintf(w, "%s", symbolAsVar(arg))
			} else {
				if err := cg.generateExpr(w, arg); err != nil {
					return err
				}
			}
		}
		fmt.Fprintf(w, ")\n")
	case "tailapply":
		// (tailapply f a b c ...)
		ok, err := cg.primitiveCallOptimize(w, sexp, true)
		if err != nil {
			return err
		}
		if ok {
			return nil
		}

		fmt.Fprintf(w, "__e.TailApply(")
		args := kl.ListToSlice(kl.Cdr(sexp))
		for i, arg := range args {
			if i != 0 {
				fmt.Fprintf(w, ", ")
			}
			if kl.IsSymbol(arg) {
				fmt.Fprintf(w, "%s", symbolAsVar(arg))
			} else {
				if err := cg.generateExpr(w, arg); err != nil {
					return err
				}
			}
		}
		fmt.Fprintf(w, ")\nreturn\n")
	default:
		return fmt.Errorf("unknown instruct: %s", kind)
	}
	return nil
}

// generateBlock recognizes the flattened temporary chains emitted by the
// Shen compiler. A pure scalar assignment is inlined only when its temporary
// has exactly one later use; all other assignments retain their historical
// statement form. This keeps side effects, rebinding, and evaluation order
// unchanged while allowing scalarExpr to see across <= boundaries.
func (cg *CodeGenerator) generateBlock(w io.Writer, sexp kl.Obj) error {
	forms := kl.ListToSlice(kl.Cdr(sexp))
	defs := make(map[kl.Obj]kl.Obj)
	for i, p := range forms {
		if kl.PrimIsPair(p) != kl.False && kl.GetSymbol(kl.Car(p)) == "<=" {
			target, rhs := kl.Cadr(p), kl.Car(kl.Cdr(kl.Cdr(p)))
			if kl.IsSymbol(target) {
				// Expand only definitions already proven single-use.
				rhs = expandScalarRefs(rhs, defs)
				// Keep the assignment's evaluation point stable: only inline when
				// the temporary is consumed by the immediately following form.
				// This prevents moving even a pure-looking expression across an
				// intervening effectful/control-flow instruction.
				if _, ok := cg.scalarExpr(rhs); ok && len(forms) > i+1 && countSymbol([]kl.Obj{forms[i+1]}, target) == 1 && countSymbol(forms[i+1:], target) == 1 && !reassigned(forms[i+1:], target) {
					defs[target] = rhs
					continue
				}
			}
		}
		p = expandScalarCall(p, defs)
		if err := cg.generateExpr(w, p); err != nil {
			return err
		}
		fmt.Fprintln(w)
		// Definitions are consumed at their sole use. Never substitute a
		// scalar expression into a second expression, even if malformed IR
		// reports an inaccurate use count.
		for s := range defs {
			if countSymbol([]kl.Obj{p}, s) > 0 {
				delete(defs, s)
			}
		}
	}
	fmt.Fprintln(w)
	return nil
}

func countSymbol(forms []kl.Obj, target kl.Obj) int {
	n := 0
	var walk func(kl.Obj)
	walk = func(o kl.Obj) {
		if kl.IsSymbol(o) {
			if o == target {
				n++
			}
			return
		}
		if kl.PrimIsPair(o) == kl.False {
			return
		}
		for cur := o; cur != kl.Nil; cur = kl.Cdr(cur) {
			walk(kl.Car(cur))
		}
	}
	for _, f := range forms {
		walk(f)
	}
	return n
}

func reassigned(forms []kl.Obj, target kl.Obj) bool {
	for _, f := range forms {
		if kl.PrimIsPair(f) != kl.False && kl.GetSymbol(kl.Car(f)) == "<=" && kl.Cadr(f) == target {
			return true
		}
	}
	return false
}

func expandScalarRefs(form kl.Obj, defs map[kl.Obj]kl.Obj) kl.Obj {
	if kl.IsSymbol(form) {
		if x, ok := defs[form]; ok {
			return x
		}
		return form
	}
	if kl.PrimIsPair(form) == kl.False {
		return form
	}
	items := kl.ListToSlice(form)
	out := make([]kl.Obj, len(items))
	for i, x := range items {
		out[i] = expandScalarRefs(x, defs)
	}
	ret := kl.Nil
	for i := len(out) - 1; i >= 0; i-- {
		ret = kl.Cons(out[i], ret)
	}
	return ret
}

func expandScalarCall(form kl.Obj, defs map[kl.Obj]kl.Obj) kl.Obj {
	if len(defs) == 0 || kl.PrimIsPair(form) == kl.False {
		return form
	}
	kind := kl.GetSymbol(kl.Car(form))
	if kind == "return" {
		items := kl.ListToSlice(form)
		if len(items) > 1 {
			items[1] = expandScalarRefs(items[1], defs)
		}
		ret := kl.Nil
		for i := len(items) - 1; i >= 0; i-- {
			ret = kl.Cons(items[i], ret)
		}
		return ret
	}
	if kind != "call" && kind != "tailapply" {
		return form
	}
	items := kl.ListToSlice(form)
	for i := 2; i < len(items); i++ {
		items[i] = expandScalarRefs(items[i], defs)
	}
	ret := kl.Nil
	for i := len(items) - 1; i >= 0; i-- {
		ret = kl.Cons(items[i], ret)
	}
	return ret
}

func isTypeAnnotation(sym kl.Obj) bool {
	switch kl.GetSymbol(sym) {
	case "number", "fixnum", "boolean", "string", "symbol", "pair", "nil", "vector", "callable", "stream", "error", "raw", "unknown":
		return true
	default:
		return false
	}
}

type scalarKind uint8

const (
	scalarNumber scalarKind = iota
	scalarBoolean
	scalarString
)

type scalarLeaf struct {
	kind scalarKind
	src  string
}

// scalarExpr is a side-effect-free expression which can be evaluated with
// native Go values. Leaves are deliberately limited to symbols and constants,
// so evaluation is never duplicated when the guarded path falls back.
type scalarExpr struct {
	kind       scalarKind
	expr       string
	leaves     []scalarLeaf
	primitives []string
}

func mergeScalar(a, b []string) []string {
	out := append([]string(nil), a...)
	for _, x := range b {
		seen := false
		for _, y := range out {
			if x == y {
				seen = true
				break
			}
		}
		if !seen {
			out = append(out, x)
		}
	}
	return out
}

func (cg *CodeGenerator) scalarExpr(form kl.Obj) (scalarExpr, bool) {
	// Type annotations are advisory metadata. Strip a well-formed wrapper for
	// optimization, while preserving the underlying expression for the guarded
	// runtime check and dynamic fallback.
	if kl.PrimIsPair(form) != kl.False && kl.GetSymbol(kl.Car(form)) == "$type" {
		annotation := kl.Cadr(form)
		expr := kl.Car(kl.Cdr(kl.Cdr(form)))
		if kl.IsSymbol(annotation) && (isTypeAnnotation(annotation) || !kl.IsSymbol(expr)) {
			form = expr
		} else {
			form = annotation
		}
	}
	if kl.PrimIsPair(form) == kl.False {
		return scalarExpr{}, false
	}
	kind := kl.GetSymbol(kl.Car(form))
	if kind != "call" && kind != "tailapply" {
		return scalarExpr{}, false
	}
	fn := kl.Cadr(form)
	if kl.PrimIsPair(fn) == kl.False || kl.Car(fn) != symGlobal {
		return scalarExpr{}, false
	}
	name := kl.GetSymbol(kl.Cadr(fn))
	args := kl.ListToSlice(kl.Cdr(kl.Cdr(form)))
	var want scalarKind
	switch name {
	case "+", "-", "*", "/", "<", "<=", ">", ">=":
		if len(args) != 2 {
			return scalarExpr{}, false
		}
		if name == "<" || name == "<=" || name == ">" || name == ">=" {
			want = scalarBoolean
		} else {
			want = scalarNumber
		}
	case "not":
		if len(args) != 1 {
			return scalarExpr{}, false
		}
		want = scalarBoolean
	case "cn":
		if len(args) != 2 {
			return scalarExpr{}, false
		}
		want = scalarString
	case "tlstr":
		if len(args) != 1 {
			return scalarExpr{}, false
		}
		want = scalarString
	default:
		return scalarExpr{}, false
	}
	parts := make([]string, len(args))
	var leaves []scalarLeaf
	var prims []string
	operandKind := scalarNumber
	if name == "cn" || name == "tlstr" {
		operandKind = scalarString
	}
	if name == "not" {
		operandKind = scalarBoolean
	}
	for i, a := range args {
		wasConst := false
		if kl.PrimIsPair(a) != kl.False && kl.GetSymbol(kl.Car(a)) == "$const" {
			wasConst = true
			a = kl.Cadr(a)
		}
		if kl.IsSymbol(a) && !wasConst {
			idx := len(leaves)
			leaves = append(leaves, scalarLeaf{kind: operandKind, src: symbolAsVar(a)})
			// A symbol's type is guarded at runtime. The operation-specific
			// extractor below determines the required kind.
			parts[i] = fmt.Sprintf("__typedV%d", idx)
			continue
		}
		if kl.IsNumber(a) || kl.IsString(a) || a == kl.True || a == kl.False {
			idx := len(leaves)
			k := scalarNumber
			src := ""
			switch {
			case kl.IsNumber(a):
				src = "MakeNumber(" + strconv.FormatFloat(kl.GetNumber(a), 'g', -1, 64) + ")"
			case kl.IsString(a):
				k, src = scalarString, fmt.Sprintf("MakeString(%#v)", kl.GetString(a))
			case a == kl.True:
				k, src = scalarBoolean, "True"
			case a == kl.False:
				k, src = scalarBoolean, "False"
			}
			if k != operandKind {
				return scalarExpr{}, false
			}
			leaves = append(leaves, scalarLeaf{kind: k, src: src})
			parts[i] = fmt.Sprintf("__typedV%d", idx)
			continue
		}
		x, ok := cg.scalarExpr(a)
		if !ok {
			return scalarExpr{}, false
		}
		// Nested expressions must produce the argument kind expected by this
		// operation; unknown/mixed kinds split back to the Obj path.
		if (name == "cn" || name == "tlstr") && x.kind != scalarString {
			return scalarExpr{}, false
		}
		if name == "not" && x.kind != scalarBoolean {
			return scalarExpr{}, false
		}
		if name != "cn" && name != "tlstr" && name != "not" && x.kind != scalarNumber {
			return scalarExpr{}, false
		}
		off := len(leaves)
		leaves = append(leaves, x.leaves...)
		p := x.expr
		// Rewrite placeholders in reverse order to avoid collisions when the
		// destination range overlaps the source (e.g. V0 -> V1, then V1 -> V2).
		for j := len(x.leaves) - 1; j >= 0; j-- {
			p = strings.ReplaceAll(p, fmt.Sprintf("__typedV%d", j), fmt.Sprintf("__typedV%d", off+j))
		}
		parts[i] = p
		prims = mergeScalar(prims, x.primitives)
	}
	prims = mergeScalar(prims, []string{name})
	var expr string
	switch name {
	case "cn":
		expr = "(" + parts[0] + " + " + parts[1] + ")"
	case "tlstr":
		expr = "TypedStringTailValue(" + parts[0] + ")"
	case "not":
		expr = "(!" + parts[0] + ")"
	case "/":
		expr = "TypedDivideValue(" + parts[0] + ", " + parts[1] + ")"
	default:
		expr = "(" + parts[0] + " " + name + " " + parts[1] + ")"
	}
	return scalarExpr{kind: want, expr: expr, leaves: leaves, primitives: prims}, true
}

// scalarTree is retained for package-local callers and compatibility tests.
func (cg *CodeGenerator) scalarTree(form kl.Obj) (string, []string, bool) {
	x, ok := cg.scalarExpr(form)
	if !ok || x.kind != scalarNumber {
		return "", nil, false
	}
	leaves := make([]string, len(x.leaves))
	for i, l := range x.leaves {
		leaves[i] = l.src
	}
	return x.expr, leaves, true
}

func (cg *CodeGenerator) generateConst(w io.Writer, c kl.Obj) error {
	switch {
	case kl.IsNumber(c):
		// Shen numbers are float64. Emitting via GetInteger truncated every
		// constant to a whole number, so a shaken artifact saw 1.5 as 1 --
		// silently, and only in compiled output. FormatFloat with precision
		// -1 gives the shortest form that parses back to the identical
		// float64, so the generated literal round-trips exactly.
		f := kl.GetNumber(c)
		if math.IsInf(f, 0) || math.IsNaN(f) {
			return fmt.Errorf("cannot emit non-finite number constant %v", f)
		}
		fmt.Fprintf(w, "MakeNumber(%s)", strconv.FormatFloat(f, 'g', -1, 64))
	case kl.IsString(c):
		str := kl.GetString(c)
		fmt.Fprintf(w, "MakeString(%#v)", str)
	case kl.IsSymbol(c):
		cg.declare[c] = struct{}{}
		fmt.Fprintf(w, "sym%s", symbolAsVar(c))
	case c == kl.Nil:
		fmt.Fprintf(w, "Nil")
	case c == kl.True:
		fmt.Fprintf(w, "True")
	case c == kl.False:
		fmt.Fprintf(w, "False")
	default:
		return errors.New("unknown $const instruct")
	}
	return nil
}

func (cg *CodeGenerator) generateIfExpr(w io.Writer, a, b, c kl.Obj) error {
	fmt.Fprintf(w, "if True == ")
	if err := cg.generateExpr(w, a); err != nil {
		return err
	}
	fmt.Fprintf(w, " {\n")
	if err := cg.generateExpr(w, b); err != nil {
		return err
	}
	fmt.Fprintf(w, "} else {\n")
	if err := cg.generateExpr(w, c); err != nil {
		return err
	}
	fmt.Fprintf(w, "}\n")
	return nil
}

var symGlobal = kl.MakeSymbol("$global")
var symConst = kl.MakeSymbol("$const")

func (cg *CodeGenerator) primitiveCallOptimize(w io.Writer, sexp kl.Obj, tail bool) (bool, error) {
	// (call ($global XX) ...)
	// (tailapply ($global XX) ...)
	fn := kl.Cadr(sexp)
	if kl.PrimIsPair(fn) == kl.False || kl.Car(fn) != symGlobal {
		return false, nil
	}
	global := kl.Cadr(fn)
	if scalarPrimitive[kl.GetSymbol(global)] {
		cg.declare[global] = struct{}{}
	}
	str := kl.GetSymbol(global)
	args := kl.ListToSlice(kl.Cdr(kl.Cdr(sexp)))
	prim, ok := shenPrimitive[str]
	if !ok || prim.Arity != len(args) {
		return false, nil
	}
	primName := prim.Name
	// Scalar lowering is valid only while the global still has its canonical
	// primitive binding.  Generated code must retain the dynamic fallback: a
	// Shen program may redefine a primitive before invoking compiled code.
	guarded := scalarPrimitive[str]
	if guarded {
		if tail {
			fmt.Fprint(w, "__e.Return(")
		}
		fmt.Fprint(w, "(func() Obj {\n")
		fmt.Fprintf(w, "if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym%s) {\n", symbolAsVar(global))
		// Fuse a pure nested numeric tree into one native expression. Every
		// leaf is extracted once and the result is boxed only at region exit.
		if sx, ok := cg.scalarExpr(sexp); ok && len(sx.leaves) > 0 {
			for _, p := range sx.primitives {
				cg.declare[kl.MakeSymbol(p)] = struct{}{}
			}
			for i, leaf := range sx.leaves {
				prefix := "N"
				if leaf.kind == scalarString {
					prefix = "S"
				} else if leaf.kind == scalarBoolean {
					prefix = "B"
				}
				v := fmt.Sprintf("__typed%s%d", prefix, i)
				switch leaf.kind {
				case scalarString:
					fmt.Fprintf(w, "%s, __typedOK%d := TypedString(%s)\n", v, i, leaf.src)
				case scalarBoolean:
					fmt.Fprintf(w, "%s, __typedOK%d := TypedBoolean(%s)\n", v, i, leaf.src)
				default:
					fmt.Fprintf(w, "%s, __typedOK%d := TypedFloat64(%s)\n", v, i, leaf.src)
				}
			}
			fmt.Fprint(w, "if ")
			for i := range sx.leaves {
				if i > 0 {
					fmt.Fprint(w, " && ")
				}
				fmt.Fprintf(w, "__typedOK%d", i)
			}
			for _, p := range sx.primitives {
				fmt.Fprintf(w, " && HasCanonicalPrimitiveBinding(sym%s)", symbolAsVar(kl.MakeSymbol(p)))
			}
			fmt.Fprintln(w, " {")
			expr := sx.expr
			for i, leaf := range sx.leaves {
				prefix := "N"
				if leaf.kind == scalarString {
					prefix = "S"
				} else if leaf.kind == scalarBoolean {
					prefix = "B"
				}
				expr = strings.ReplaceAll(expr, fmt.Sprintf("__typedV%d", i), fmt.Sprintf("__typed%s%d", prefix, i))
			}
			switch sx.kind {
			case scalarBoolean:
				fmt.Fprintf(w, "return TypedMaterializeBoolean(%s)\n}", expr)
			case scalarString:
				fmt.Fprintf(w, "return TypedMaterializeString(%s)\n}", expr)
			default:
				fmt.Fprintf(w, "return TypedMaterializeNumber(%s)\n}", expr)
			}
			fmt.Fprintln(w, "}")
		} else {
			// No pure scalar tree: leave the guarded region empty. Effects and
			// raising expressions must only execute on the dynamic path below.
			fmt.Fprintln(w, "_ = false")
			fmt.Fprintln(w, "}")
		}
		// Evaluate fallback arguments exactly once, after the fast guard.
		for i, arg := range args {
			fmt.Fprintf(w, "__typedArg%d := ", i)
			if kl.IsSymbol(arg) {
				fmt.Fprintf(w, "%s", symbolAsVar(arg))
			} else if err := cg.generateExpr(w, arg); err != nil {
				return true, err
			}
			fmt.Fprintln(w)
		}
		fmt.Fprintf(w, "return Call(__e, PrimFunc(sym%s)", symbolAsVar(global))
		for i := range args {
			fmt.Fprintf(w, ", __typedArg%d", i)
		}
		fmt.Fprint(w, ")\n})()")
		if tail {
			fmt.Fprintln(w, ")\nreturn")
		}
		return true, nil
	}
	// All known primitive lowerings retain a dynamic fallback. This matters for
	// Shen programs which redefine a primitive after generated code is loaded.
	// Keep the guard before evaluating arguments so effectful expressions run on
	// exactly one selected path.
	if !guarded {
		if tail {
			fmt.Fprint(w, "__e.Return(")
		}
		fmt.Fprint(w, "(func() Obj {\n")
		fmt.Fprintf(w, "if TypedIREnabled() && HasCanonicalPrimitiveBinding(sym%s) {\n", symbolAsVar(global))
		fmt.Fprintf(w, "return %s(", primName)
		for i, arg := range args {
			if i != 0 {
				fmt.Fprint(w, ", ")
			}
			if kl.IsSymbol(arg) {
				fmt.Fprint(w, symbolAsVar(arg))
			} else if err := cg.generateExpr(w, arg); err != nil {
				return true, err
			}
		}
		fmt.Fprint(w, ")\n}\n")
		for i, arg := range args {
			fmt.Fprintf(w, "__typedArg%d := ", i)
			if kl.IsSymbol(arg) {
				fmt.Fprint(w, symbolAsVar(arg))
			} else if err := cg.generateExpr(w, arg); err != nil {
				return true, err
			}
			fmt.Fprintln(w)
		}
		fmt.Fprintf(w, "return Call(__e, PrimFunc(sym%s)", symbolAsVar(global))
		for i := range args {
			fmt.Fprintf(w, ", __typedArg%d", i)
		}
		fmt.Fprint(w, ")\n})()")
		if tail {
			fmt.Fprint(w, ")\nreturn\n")
		}
		return true, nil
	}

	// ($prim f a b c ...)
	if tail {
		fmt.Fprintf(w, "__e.Return(")
	}

	fmt.Fprintf(w, "%s(", primName)
	for i, arg := range args {
		if i != 0 {
			fmt.Fprintf(w, ", ")
		}
		if kl.IsSymbol(arg) {
			fmt.Fprintf(w, "%s", symbolAsVar(arg))
		} else {
			if err := cg.generateExpr(w, arg); err != nil {
				return true, err
			}
		}
	}
	fmt.Fprintf(w, ")")

	if tail {
		fmt.Fprintf(w, ")\nreturn\n")
	}
	return true, nil
}

// scalarPrimitive identifies operations whose generated native implementation
// is safe to select only behind the canonical-binding guard above.  Other
// primitives continue to use the historical direct lowering unchanged.
var scalarPrimitive = map[string]bool{
	"+": true, "-": true, "*": true, "/": true,
	"<": true, "<=": true, ">": true, ">=": true,
	"not": true, "cn": true, "tlstr": true,
}

var shenPrimitive = map[string]struct {
	Arity int
	Name  string
}{
	"get-time":              {1, "PrimGetTime"},
	"close":                 {1, "PrimCloseStream"},
	"open":                  {2, "PrimOpenStream"},
	"read-byte":             {1, "PrimReadByte"},
	"write-byte":            {2, "PrimWriteByte"},
	"absvector?":            {1, "PrimIsVector"},
	"<-address":             {2, "PrimVectorGet"},
	"address->":             {3, "PrimVectorSet"},
	"absvector":             {1, "PrimAbsvector"},
	"str":                   {1, "PrimStr"},
	"<=":                    {2, "PrimLessEqual"},
	">=":                    {2, "PrimGreatEqual"},
	"<":                     {2, "PrimLessThan"},
	">":                     {2, "PrimGreatThan"},
	"error-to-string":       {1, "PrimErrorToString"},
	"simple-error":          {1, "PrimSimpleError"},
	"=":                     {2, "PrimEqual"},
	"-":                     {2, "PrimNumberSubtract"},
	"*":                     {2, "PrimNumberMultiply"},
	"/":                     {2, "PrimNumberDivide"},
	"+":                     {2, "PrimNumberAdd"},
	"string->n":             {1, "PrimStringToNumber"},
	"n->string":             {1, "PrimNumberToString"},
	"number?":               {1, "PrimIsNumber"},
	"string?":               {1, "PrimIsString"},
	"pos":                   {2, "PrimPos"},
	"tlstr":                 {1, "PrimTailString"},
	"cn":                    {2, "PrimStringConcat"},
	"intern":                {1, "PrimIntern"},
	"hd":                    {1, "PrimHead"},
	"tl":                    {1, "PrimTail"},
	"cons":                  {2, "PrimCons"},
	"cons?":                 {1, "PrimIsPair"},
	"value":                 {1, "PrimValue"},
	"set":                   {2, "PrimSet"},
	"not":                   {1, "PrimNot"},
	"if":                    {3, "PrimIf"},
	"symbol?":               {1, "PrimIsSymbol"},
	"read-file-as-bytelist": {1, "PrimReadFileAsByteList"},
	"read-file-as-string":   {1, "PrimReadFileAsString"},
	"variable?":             {1, "PrimIsVariable"},
	"integer?":              {1, "PrimIsInteger"},
}
