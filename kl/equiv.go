package kl

// Differential equivalence harness for the InstallKernelFast natives.
//
// InstallKernelFast rebinds a set of kernel functions (reverse, put, map and
// so on) to Go natives. Each rebinding is a claim that the native implements
// the KLambda definition of the same name. This file is the shared machinery
// that checks those claims one case at a time. Two consumers use it:
//
//   - kl/equiv_test.go generates the cases, runs them, and writes the result
//     out as kl/equiv.json. That file is the port's declared, verified
//     lowering table, and the Yggdrasil tree-shaker imports it.
//   - `kl equiv-check kl/equiv.json` (cmd/kl) re-runs the cases stored in the
//     JSON on any host and prints one line per row.
//
// The KL side of a row is the kernel's own defun for the function, loaded
// under the name equiv.NAME. Every kernel defun that body reaches through
// head-position calls is loaded the same way, transitively: shen.map-h under
// map, shen.app under the error paths, and so on down. The KL side therefore
// runs as KL bytecode all the way to the primitives and never enters a native
// under test. Two names are deliberately left alone:
//
//   - hash. BootKernel swaps in the native FNV-1a hash before the first
//     property vector is built, the way cmd/shen does, so every vector in the
//     process is keyed by it. A KL-side hash would key the same vector
//     differently, and the two sides would compare nothing.
//   - fail. The vendored sys.kl says (defun fail () shen.fail!), while the
//     port's natives and its compiled kernel (cmd/shen/sys.go) return the
//     symbol `...`. That divergence is the fail row's own finding. The other
//     KL copies call the port's fail, so both sides share one vector filler,
//     and the rows that depend on it (vector, <-vector, put, get, unput) are
//     judged on their own logic.
//
// The native side of a row is the rebound symbol itself. A case is a list of
// KL expressions for the arguments, plus optional setup forms run before the
// call and optional observe forms evaluated after it. The observe forms are
// how the harness compares property vectors and side-effect logs. Setup,
// argument and observe forms are evaluated identically on both sides, with the
// natives installed, so both sides see the same input objects. The two sides
// must produce equal values, or raise errors with the same text, and equal
// observations.

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// EquivCase is one differential case. Every string is a KL form in the syntax
// kl.NewSexpReader accepts. Setup forms run before the call on each side, so
// both sides start from identical state. Args are evaluated to give the call
// its arguments. Observe forms are evaluated after the call, and their values
// are compared alongside the result.
type EquivCase struct {
	Name    string   `json:"name"`
	Setup   []string `json:"setup,omitempty"`
	Args    []string `json:"args"`
	Observe []string `json:"observe,omitempty"`
}

// EquivRow is one rebound kernel function.
type EquivRow struct {
	KernelFn string `json:"kernel_fn"`
	Native   string `json:"native"`
	Arity    int    `json:"arity"`
	// Verified is true when every case agreed at the time the table was
	// generated: equal values, or equal error text, and equal observations.
	Verified bool `json:"verified"`
	// Cases is len(Inputs).
	Cases   int      `json:"cases"`
	Effects []string `json:"effects,omitempty"`
	// Reason is the first disagreeing case, "case=NAME kl=… native=…".
	Reason string `json:"reason,omitempty"`
	// Disagreement classifies Reason. See MismatchKind.
	Disagreement string `json:"disagreement,omitempty"`
	// Note is prose about an unverified row, saying where the divergence comes
	// from. It is kept apart from Reason so that Reason stays machine-readable.
	Note   string      `json:"note,omitempty"`
	Source string      `json:"source"`
	Inputs []EquivCase `json:"inputs"`
}

// EquivHarness documents, inside the table, what the rows' verdicts mean and
// the vocabulary the cases are written in, so that a consumer can read the
// file without this package.
type EquivHarness struct {
	KLSide     string            `json:"kl_side"`
	NativeSide string            `json:"native_side"`
	Inputs     string            `json:"inputs"`
	Verified   string            `json:"verified"`
	Cases      string            `json:"cases"`
	Source     string            `json:"source"`
	Effects    string            `json:"effects"`
	Vocabulary map[string]string `json:"vocabulary"`
	// KLHelpers are the kernel defuns that are not themselves rebound but are
	// loaded as equiv.* copies because a rebound function's body reaches them.
	KLHelpers []string `json:"kl_helpers"`
}

// EquivTable is the exported lowering table (kl/equiv.json).
type EquivTable struct {
	Port    string       `json:"port"`
	Kernel  string       `json:"kernel"`
	Harness EquivHarness `json:"harness"`
	Rows    []EquivRow   `json:"rows"`
}

// EquivPrefix is the namespace the KL-side copies of the kernel functions are
// loaded under. The kernel's (defun reverse ...) becomes equiv.reverse.
const EquivPrefix = "equiv."

// EquivName returns the KL-side name for a kernel function.
func EquivName(kernelFn string) string { return EquivPrefix + kernelFn }

// equivKeepNative are the kernel defuns the KL side keeps calling by their
// real, native name. The file comment says why.
var equivKeepNative = map[string]bool{"hash": true, "fail": true}

// EquivVocabulary describes the harness-only names a case may use.
var EquivVocabulary = map[string]string{
	"equiv.NAME":        "the kernel's own (defun NAME …) loaded as KL bytecode, with every kernel defun its body reaches renamed the same way; hash and fail are the two exceptions and stay native",
	"equiv.iota N":      "the list (1 2 … N)",
	"equiv.note X":      "conses X onto (value equiv.*calls*) and returns X; this is how a case counts and orders the calls made to a function argument",
	"equiv.native-of F": "the function object currently bound to the symbol F, which for a rebound name is the Go native",
	"equiv.*calls*":     "the global holding the equiv.note log; setup resets it and observe compares it",
	"equiv.*pv*":        "the global holding a fresh property vector for the put, get and unput cases; both sides start from an identical fresh vector, and observe compares it",
}

// ReadForms parses every top-level form in src.
func ReadForms(src io.Reader) ([]Obj, error) {
	r := NewSexpReader(src, false)
	var out []Obj
	for {
		form, err := r.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return out, nil
			}
			return nil, err
		}
		out = append(out, form)
	}
}

// ReadForm parses exactly one form from src.
func ReadForm(src string) (Obj, error) {
	// The reader only finishes an atom at a delimiter, so a bare "foo" at
	// end of input needs the trailing space.
	forms, err := ReadForms(strings.NewReader(src + " "))
	if err != nil {
		return Nil, err
	}
	if len(forms) != 1 {
		return Nil, fmt.Errorf("expected one form in %q, got %d", src, len(forms))
	}
	return forms[0], nil
}

// KernelDefuns returns every (defun NAME …) form in the kernel sources under
// dir, keyed by NAME, taking the last definition in load order when a name is
// defined more than once.
func KernelDefuns(dir string) (map[string]Obj, error) {
	symDefun := MakeSymbol("defun")
	out := make(map[string]Obj)
	for _, f := range KernelLoadOrder {
		fh, err := os.Open(filepath.Join(dir, f))
		if err != nil {
			return nil, err
		}
		forms, err := ReadForms(fh)
		fh.Close()
		if err != nil {
			return nil, fmt.Errorf("%s: %v", f, err)
		}
		for _, form := range forms {
			ok, p := isPair(form)
			if !ok || p.car != symDefun {
				continue
			}
			if nameObj := Cadr(form); IsSymbol(nameObj) {
				out[GetSymbol(nameObj)] = form
			}
		}
	}
	return out, nil
}

// defunBody returns the body of a (defun NAME PARAMS BODY) form.
func defunBody(defun Obj) Obj { return Car(Cdr(Cdr(Cdr(defun)))) }

// headSymbols collects every symbol in head (operator) position anywhere in
// form.
func headSymbols(form Obj, into map[string]bool) {
	ok, p := isPair(form)
	if !ok {
		return
	}
	if IsSymbol(p.car) {
		into[GetSymbol(p.car)] = true
	} else {
		headSymbols(p.car, into)
	}
	for cur := p.cdr; ; {
		okc, cp := isPair(cur)
		if !okc {
			return
		}
		headSymbols(cp.car, into)
		cur = cp.cdr
	}
}

// EquivClosure returns, in a deterministic order, every kernel defun reachable
// from the bodies of roots through head-position calls, excluding the roots
// themselves and the names in equivKeepNative. These are the helpers the KL
// side must also run as KL so that the rows' KL copies stay native-free.
func EquivClosure(defuns map[string]Obj, roots []string) []string {
	isRoot := make(map[string]bool, len(roots))
	for _, r := range roots {
		isRoot[r] = true
	}
	seen := map[string]bool{}
	var helpers []string
	queue := append([]string(nil), roots...)
	for len(queue) > 0 {
		name := queue[0]
		queue = queue[1:]
		d, ok := defuns[name]
		if !ok {
			continue
		}
		heads := map[string]bool{}
		headSymbols(defunBody(d), heads)
		next := make([]string, 0, len(heads))
		for h := range heads {
			next = append(next, h)
		}
		sort.Strings(next)
		for _, h := range next {
			if seen[h] || isRoot[h] || equivKeepNative[h] {
				continue
			}
			if _, ok := defuns[h]; !ok {
				continue // a primitive, a special form or a variable
			}
			seen[h] = true
			helpers = append(helpers, h)
			queue = append(queue, h)
		}
	}
	return helpers
}

// renameHeads returns a copy of form in which every application whose head is
// a symbol in rename applies the renamed symbol instead. Only head positions
// change. A symbol appearing as data, such as an argument or a cond test
// value, is left alone. The body keeps its meaning, except that it calls the
// equiv.* copies rather than the natives under test.
func renameHeads(form Obj, rename map[string]string) Obj {
	ok, p := isPair(form)
	if !ok {
		return form
	}
	head := p.car
	if IsSymbol(head) {
		if to, ok := rename[GetSymbol(head)]; ok {
			head = MakeSymbol(to)
		}
	} else {
		head = renameHeads(head, rename)
	}
	return Cons(head, renameElems(p.cdr, rename))
}

// renameElems applies renameHeads to each element of the argument list l.
func renameElems(l Obj, rename map[string]string) Obj {
	ok, p := isPair(l)
	if !ok {
		return l
	}
	return Cons(renameHeads(p.car, rename), renameElems(p.cdr, rename))
}

// EquivDefun builds the equiv.NAME copy of a kernel defun: the defun's own
// name is prefixed, and so is every head-position call to a name in rename.
func EquivDefun(defun Obj, rename map[string]string) Obj {
	name := Cadr(defun)
	params := Car(Cdr(Cdr(defun)))
	return Cons(MakeSymbol("defun"),
		Cons(MakeSymbol(EquivName(GetSymbol(name))),
			Cons(params, Cons(renameHeads(defunBody(defun), rename), Nil))))
}

// InstallEquivDefinitions loads the equiv.* copy of every rebound kernel
// function in names, and of every kernel defun their bodies reach, from the
// kernel sources in dir. EquivClosure computes that set of helpers. It also
// binds the harness helpers the cases rely on, which EquivVocabulary
// describes. BootKernel must have run first. It returns the helper names
// loaded, or an error naming any function in names that has no kernel defun.
func InstallEquivDefinitions(e *ControlFlow, dir string, names []string) ([]string, error) {
	defuns, err := KernelDefuns(dir)
	if err != nil {
		return nil, err
	}
	var missing []string
	for _, n := range names {
		if _, ok := defuns[n]; !ok {
			missing = append(missing, n)
		}
	}
	if len(missing) > 0 {
		return nil, fmt.Errorf("no kernel defun for: %s", strings.Join(missing, ", "))
	}
	helpers := EquivClosure(defuns, names)
	rename := map[string]string{}
	for _, n := range append(append([]string(nil), names...), helpers...) {
		if !equivKeepNative[n] {
			rename[n] = EquivName(n)
		}
	}
	// Helpers first. A root's body may be evaluated by defun before the helper
	// exists, which is fine for KL, but keeping this order makes the
	// definitions independent of how defun resolves names.
	for _, n := range append(append([]string(nil), helpers...), names...) {
		form := EquivDefun(defuns[n], rename)
		if res := Eval(e, form); IsError(res) {
			return nil, fmt.Errorf("defining %s: %s", EquivName(n), GetString(PrimErrorToString(res)))
		}
	}
	BindSymbolFunc(MakeSymbol("equiv.native-of"), MakePrimitive("equiv.native-of", 1, func(f Obj) Obj {
		return PrimFunc(f)
	}))
	BindSymbolFunc(MakeSymbol("equiv.iota"), MakePrimitive("equiv.iota", 1, func(n Obj) Obj {
		acc := Nil
		for i := mustInteger(n); i >= 1; i-- {
			acc = Cons(MakeInteger(i), acc)
		}
		return acc
	}))
	symCalls := MakeSymbol("equiv.*calls*")
	BindSymbolFunc(MakeSymbol("equiv.note"), MakePrimitive("equiv.note", 1, func(x Obj) Obj {
		prev := mustSymbol(symCalls).value
		if prev == nil {
			prev = Nil
		}
		PrimSet(symCalls, Cons(x, prev))
		return x
	}))
	return helpers, nil
}

// equivOutcome is what one side of a case produced.
type equivOutcome struct {
	failed   string // non-empty: the case could not even be set up
	isErr    bool
	err      string
	val      Obj
	observed []Obj
}

func (o equivOutcome) String() string {
	if o.failed != "" {
		return "harness: " + o.failed
	}
	var b strings.Builder
	if o.isErr {
		fmt.Fprintf(&b, "error(%q)", o.err)
	} else {
		b.WriteString(equivString(o.val))
	}
	for i, x := range o.observed {
		fmt.Fprintf(&b, " observe[%d]=%s", i, equivString(x))
	}
	return b.String()
}

// equivString prints a value for a mismatch report. It matches ObjString
// except for vectors, which it shows as "#vector[limit slot1 …]" with
// never-written slots as #unset, because the property-vector rows disagree in
// the slots. Lists are printed element by element, so a vector nested in one
// prints the same way.
func equivString(o Obj) string {
	if o == nil {
		return "#unset"
	}
	if ok, p := isPair(o); ok {
		var b strings.Builder
		b.WriteString("(")
		b.WriteString(equivString(p.car))
		for cur := p.cdr; ; {
			okc, cp := isPair(cur)
			if !okc {
				if cur != Nil {
					b.WriteString(" . ")
					b.WriteString(equivString(cur))
				}
				break
			}
			b.WriteString(" ")
			b.WriteString(equivString(cp.car))
			cur = cp.cdr
		}
		b.WriteString(")")
		return b.String()
	}
	if !isFixnum(o) && *o == scmHeadVector {
		parts := make([]string, 0, len(mustVector(o)))
		for _, x := range mustVector(o) {
			parts = append(parts, equivString(x))
		}
		return "#vector[" + strings.Join(parts, " ") + "]"
	}
	return ObjString(o)
}

// safeCall applies f to args on a fresh control flow, turning a raised Shen
// error into (nil, msg, true). A Go panic that is not a Shen error is reported
// as an error too, prefixed "go panic:", so that both sides can still be
// compared.
func safeCall(f Obj, args []Obj) (val Obj, errMsg string, isErr bool) {
	defer func() {
		if r := recover(); r != nil {
			isErr = true
			if x, ok := r.(Obj); ok && IsError(x) {
				errMsg = mustError(x).err
				return
			}
			errMsg = fmt.Sprintf("go panic: %v", r)
		}
	}()
	var e ControlFlow
	val = Call(&e, f, args...)
	return val, "", false
}

// evalQuiet evaluates one KL source expression. It returns any raised error
// without Eval's stdout panic trace.
func evalQuiet(e *ControlFlow, src string) (val Obj, err error) {
	form, perr := ReadForm(src)
	if perr != nil {
		return Nil, perr
	}
	defer func() {
		if r := recover(); r != nil {
			if x, ok := r.(Obj); ok && IsError(x) {
				err = fmt.Errorf("%s raised %q", src, mustError(x).err)
				return
			}
			err = fmt.Errorf("%s panicked: %v", src, r)
		}
	}()
	var fresh ControlFlow
	val = evalExp(&fresh, form, Nil)
	return val, nil
}

func runEquivSide(e *ControlFlow, fn Obj, c EquivCase) equivOutcome {
	var out equivOutcome
	for _, s := range c.Setup {
		if _, err := evalQuiet(e, s); err != nil {
			out.failed = "setup " + err.Error()
			return out
		}
	}
	args := make([]Obj, 0, len(c.Args))
	for _, a := range c.Args {
		v, err := evalQuiet(e, a)
		if err != nil {
			out.failed = "arg " + err.Error()
			return out
		}
		args = append(args, v)
	}
	out.val, out.err, out.isErr = safeCall(fn, args)
	for _, s := range c.Observe {
		v, err := evalQuiet(e, s)
		if err != nil {
			out.failed = "observe " + err.Error()
			return out
		}
		out.observed = append(out.observed, v)
	}
	return out
}

// equivEqual is structural equality for comparing the two sides. It is the
// kernel's `=`, extended to function objects. A function is equal to itself,
// meaning the same pointer, as (fn reverse) gives on both sides. Two distinct
// closures built from KL source are equal when their arities agree: that
// covers a (lambda X X) argument evaluated once per side, and the lambda that
// a KL fn or protect wraps. No structural test could relate them further, and
// the natives under test never construct functions of their own. A native is
// equal only to itself.
func equivEqual(x, y Obj) bool {
	if isCallable(x) && isCallable(y) {
		if x == y {
			return true
		}
		if IsNativeBinding(x) || IsNativeBinding(y) {
			return false
		}
		return closureArity(x) == closureArity(y)
	}
	if ok, px := isPair(x); ok {
		oky, py := isPair(y)
		return oky && equivEqual(px.car, py.car) && equivEqual(px.cdr, py.cdr)
	}
	if PrimIsVector(x) == True && PrimIsVector(y) == True {
		vx, vy := mustVector(x), mustVector(y)
		if len(vx) != len(vy) {
			return false
		}
		for i := range vx {
			a, b := vx[i], vy[i]
			if a == nil {
				a = undefined
			}
			if b == nil {
				b = undefined
			}
			if !equivEqual(a, b) {
				return false
			}
		}
		return true
	}
	return equal(x, y) == True
}

func isCallable(o Obj) bool {
	if o == nil || isFixnum(o) {
		return false
	}
	switch *o {
	case scmHeadProcedure, scmHeadNative, scmHeadBytecodeFunc:
		return true
	}
	return false
}

func closureArity(o Obj) int {
	switch *o {
	case scmHeadProcedure:
		return mustProcedure(o).arity
	case scmHeadBytecodeFunc:
		return mustBytecodeFunc(o).fn.Arity
	}
	return -1
}

// MismatchKind classifies how the two sides of a case disagreed.
type MismatchKind string

const (
	// MismatchRaise means one side raised an error and the other returned a
	// value.
	MismatchRaise MismatchKind = "raise"
	// MismatchValue means both sides returned and the values differ.
	MismatchValue MismatchKind = "value"
	// MismatchErrorText means both sides raised and the error messages differ.
	MismatchErrorText MismatchKind = "error-text"
	// MismatchEffect means the results agree but an observed value differs,
	// such as a property vector or a call log.
	MismatchEffect MismatchKind = "effect"
	// MismatchHarness means a setup, argument or observe form itself failed,
	// so the case could not be run. The harness is at fault, not the native.
	MismatchHarness MismatchKind = "harness"
)

// EquivMismatch describes a case on which the two sides disagreed.
type EquivMismatch struct {
	Case   string
	KL     string
	Native string
	Kind   MismatchKind
}

func (m *EquivMismatch) String() string {
	return fmt.Sprintf("case=%s kl=%s native=%s", m.Case, m.KL, m.Native)
}

// classify returns the kind of disagreement between the two sides, or "" when
// they agree.
func classify(kl, native equivOutcome) MismatchKind {
	if kl.failed != "" || native.failed != "" {
		return MismatchHarness
	}
	if kl.isErr != native.isErr {
		return MismatchRaise
	}
	if kl.isErr {
		if kl.err != native.err {
			return MismatchErrorText
		}
	} else if !equivEqual(kl.val, native.val) {
		return MismatchValue
	}
	if len(kl.observed) != len(native.observed) {
		return MismatchEffect
	}
	for i := range kl.observed {
		if !equivEqual(kl.observed[i], native.observed[i]) {
			return MismatchEffect
		}
	}
	return ""
}

// RunEquivCase runs the case c against both sides of kernelFn: the KL copy
// equiv.NAME and the native bound to NAME. It returns nil when the two sides
// agree.
func RunEquivCase(e *ControlFlow, kernelFn string, c EquivCase) *EquivMismatch {
	klFn := mustSymbol(MakeSymbol(EquivName(kernelFn))).function
	if klFn == nil {
		return &EquivMismatch{c.Name, "<" + EquivName(kernelFn) + " not defined>", "<not run>", MismatchHarness}
	}
	nat := mustSymbol(MakeSymbol(kernelFn)).function
	if nat == nil {
		return &EquivMismatch{c.Name, "<not run>", "<" + kernelFn + " not bound>", MismatchHarness}
	}
	kl := runEquivSide(e, klFn, c)
	native := runEquivSide(e, nat, c)
	kind := classify(kl, native)
	if kind == "" {
		return nil
	}
	return &EquivMismatch{c.Name, kl.String(), native.String(), kind}
}

// WithQuietStdout runs f with os.Stdout pointed at the null device, so that
// the runtime's stdout traces on recovered errors do not interleave with a
// report. Those traces come from kl/eval.go and from mustPair in kl/types.go.
// A writer captured before the call, such as the real os.Stdout, is
// unaffected.
func WithQuietStdout(f func()) error {
	devnull, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	if err != nil {
		return err
	}
	saved := os.Stdout
	os.Stdout = devnull
	defer func() {
		os.Stdout = saved
		devnull.Close()
	}()
	f()
	return nil
}

// KernelArity returns the kernel's registered arity for name via (arity NAME),
// or -1 when the kernel has none.
func KernelArity(e *ControlFlow, name string) int {
	v, _, isErr := safeCall(PrimFunc(MakeSymbol("arity")), []Obj{MakeSymbol(name)})
	if isErr || !IsNumber(v) {
		return -1
	}
	return GetInteger(v)
}

// IsNativeBinding reports whether f is a Go-implemented function object,
// rather than KL bytecode or an interpreted procedure. MakePrimitive and
// MakeNative both build one.
func IsNativeBinding(f Obj) bool {
	if f == nil || isFixnum(f) {
		return false
	}
	return *f == scmHeadNative
}
