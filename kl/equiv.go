package kl

// Differential equivalence harness for the InstallKernelFast natives.
//
// InstallKernelFast rebinds a set of kernel functions (reverse, put, map, …)
// to Go natives. Each rebinding is a claim that the native implements the
// KLambda definition of the same name. This file is the shared machinery that
// checks those claims one by one, used by two consumers:
//
//   - kl/equiv_test.go generates the cases, runs them, and writes the result
//     out as kl/equiv.json (the port's declared, verified lowering table that
//     the Yggdrasil tree-shaker imports);
//   - `kl equiv-check kl/equiv.json` (cmd/kl) re-runs the cases stored in the
//     JSON on any host and prints one line per row.
//
// The KL side of a row is the kernel's own defun for the function, loaded
// under the name equiv.NAME with every call to a rebound function inside its
// body renamed the same way (so the whole chain runs as KL bytecode, not as
// the natives under test). The native side is the rebound symbol itself. A
// case is a list of KL expressions for the arguments, optional setup forms
// run before the call, and optional observe forms evaluated after it (used to
// compare property vectors and side-effect logs). Both sides must produce
// equal values, or raise errors with the same text, and equal observations.

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// KernelLoadOrder is the canonical load order of the KLambda kernel under
// kernel/klambda, mirrored from compiled/precompile.kl and cmd/shen's regist().
// The Tarver S42 kernel has no shen.initialise: each module runs its own
// top-level init forms as it loads, so this order must not change.
var KernelLoadOrder = []string{
	"sys.kl", "writer.kl", "core.kl", "reader.kl", "declarations.kl",
	"toplevel.kl", "macros.kl", "load.kl", "prolog.kl", "sequent.kl",
	"track.kl", "t-star.kl", "yacc.kl", "types.kl",
	"extension-launcher.kl",
}

// BootKernel loads the KLambda kernel from dir (a kernel/klambda directory)
// into the process the same way cmd/shen boots its compiled kernel:
//
//   - sys.kl defines the interpreted `hash`; the native FNV-1a hash is swapped
//     in right after, before declarations.kl builds the first property
//     dictionary, so every dictionary is written and read with one hash;
//   - InstallKernelFast runs after every module. In cmd/shen it runs once, after
//     all modules, because the compiled modules never execute the interpreted
//     put/get. Here the modules are interpreted, and since 5edf47e the
//     interpreted `put` (a tail call inside trap-error) escapes the
//     interpreter's recover during declarations.kl (issue #46), so the natives
//     have to be in place as soon as sys.kl has defined the names they replace.
//     InstallKernelFast only rebinds names the kernel has defined, so calling
//     it per module is the way to catch each module's definitions;
//   - InstallPr binds shen.native-pr, as cmd/shen's fixPrHush does.
func BootKernel(e *ControlFlow, dir string) error {
	for i, f := range KernelLoadOrder {
		p := filepath.Join(dir, f)
		if _, err := os.Stat(p); err != nil {
			return fmt.Errorf("boot kernel: %v", err)
		}
		form := Cons(MakeSymbol("load-file"), Cons(MakeString(p), Nil))
		if res := Eval(e, form); IsError(res) {
			return fmt.Errorf("boot kernel: load-file %s: %s", p, GetString(PrimErrorToString(res)))
		}
		if i == 0 {
			BindSymbolFunc(MakeSymbol("hash"), MakePrimitive("hash", 2, PrimHash))
		}
		InstallKernelFast()
	}
	InstallPr()
	return nil
}

// FindKernelDir walks up from start looking for kernel/klambda/sys.kl and
// returns that kernel/klambda directory.
func FindKernelDir(start string) (string, error) {
	dir, err := filepath.Abs(start)
	if err != nil {
		return "", err
	}
	for {
		cand := filepath.Join(dir, "kernel", "klambda")
		if _, err := os.Stat(filepath.Join(cand, "sys.kl")); err == nil {
			return cand, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("kernel/klambda not found above %s", start)
		}
		dir = parent
	}
}

// EquivCase is one differential input. Every string is a KL expression (or
// form) in the syntax kl.NewSexpReader accepts. Setup forms run before the
// call on each side (so both sides start from identical state), Args are
// evaluated to the call's arguments, and Observe forms are evaluated after the
// call; their values are compared alongside the result.
type EquivCase struct {
	Name    string   `json:"name"`
	Setup   []string `json:"setup,omitempty"`
	Args    []string `json:"args"`
	Observe []string `json:"observe,omitempty"`
}

// EquivRow is one rebound kernel function.
type EquivRow struct {
	KernelFn string      `json:"kernel_fn"`
	Native   string      `json:"native"`
	Arity    int         `json:"arity"`
	Verified bool        `json:"verified"`
	Cases    int         `json:"cases"`
	Effects  []string    `json:"effects,omitempty"`
	Reason   string      `json:"reason,omitempty"`
	Source   string      `json:"source"`
	Inputs   []EquivCase `json:"inputs"`
}

// EquivTable is the exported lowering table (kl/equiv.json).
type EquivTable struct {
	Port   string     `json:"port"`
	Kernel string     `json:"kernel"`
	Rows   []EquivRow `json:"rows"`
}

// EquivPrefix is the namespace the KL-side copies of the rebound functions
// are loaded under: the kernel's (defun reverse ...) becomes equiv.reverse.
const EquivPrefix = "equiv."

// EquivName returns the KL-side name for a rebound kernel function.
func EquivName(kernelFn string) string { return EquivPrefix + kernelFn }

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

// KernelDefuns returns, for each name in names, the kernel's (defun NAME …)
// form, taking the last definition in load order when a name is defined more
// than once. Names without a defun are absent from the result.
func KernelDefuns(dir string, names []string) (map[string]Obj, error) {
	want := make(map[string]bool, len(names))
	for _, n := range names {
		want[n] = true
	}
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
			nameObj := Cadr(form)
			if IsSymbol(nameObj) && want[GetSymbol(nameObj)] {
				out[GetSymbol(nameObj)] = form
			}
		}
	}
	return out, nil
}

// renameHeads returns a copy of form in which every application whose head is
// a symbol in rename is rewritten to apply the renamed symbol instead. Only
// head positions change: a symbol appearing as data (an argument, a cond test
// value, …) is left alone, so the body's meaning is preserved except that it
// calls the equiv.* copies rather than the natives under test.
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

// EquivDefun builds the equiv.NAME copy of a kernel defun: the defun's name
// and every head-position call to a name in rebound are prefixed.
func EquivDefun(defun Obj, rebound []string) Obj {
	rename := make(map[string]string, len(rebound))
	for _, n := range rebound {
		rename[n] = EquivName(n)
	}
	// (defun NAME PARAMS BODY): rename NAME directly, then the body's heads.
	name := Cadr(defun)
	params := Car(Cdr(Cdr(defun)))
	body := Car(Cdr(Cdr(Cdr(defun))))
	return Cons(MakeSymbol("defun"),
		Cons(MakeSymbol(EquivName(GetSymbol(name))),
			Cons(params, Cons(renameHeads(body, rename), Nil))))
}

// InstallEquivDefinitions loads the equiv.* copy of every rebound kernel
// function in names from the kernel sources in dir, and binds the harness
// helpers the cases rely on:
//
//	(equiv.native-of F)  the function object currently bound to symbol F
//	(equiv.iota N)       the list (1 2 … N)
//	(equiv.note X)       conses X onto equiv.*calls* and returns X
//
// BootKernel must have run first. Returns an error naming any function that
// has no kernel defun.
func InstallEquivDefinitions(e *ControlFlow, dir string, names []string) error {
	defuns, err := KernelDefuns(dir, names)
	if err != nil {
		return err
	}
	var missing []string
	for _, n := range names {
		d, ok := defuns[n]
		if !ok {
			missing = append(missing, n)
			continue
		}
		form := EquivDefun(d, names)
		if res := Eval(e, form); IsError(res) {
			return fmt.Errorf("defining %s: %s", EquivName(n), GetString(PrimErrorToString(res)))
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("no kernel defun for: %s", strings.Join(missing, ", "))
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
	return nil
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
		b.WriteString(ObjString(o.val))
	}
	for i, x := range o.observed {
		fmt.Fprintf(&b, " observe[%d]=%s", i, ObjString(x))
	}
	return b.String()
}

// safeCall applies f to args on a fresh control flow, turning a raised Shen
// error into (nil, msg, true). A non-Shen Go panic is reported as an error
// too, prefixed "go panic:", so both sides can still be compared.
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

// evalQuiet evaluates one KL source expression, returning the raised error
// (if any) without Eval's stdout panic trace.
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

// equivEqual is structural equality for comparing the two sides. It is
// kernel `=` (equal) except that two function objects count as equal: a case
// whose argument is (lambda X X) evaluates it once per side, so the sides hold
// distinct closures that no structural test could relate, and the natives
// under test never construct functions of their own.
func equivEqual(x, y Obj) bool {
	if isCallable(x) && isCallable(y) {
		return true
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

func sameOutcome(a, b equivOutcome) bool {
	if a.failed != "" || b.failed != "" {
		return false
	}
	if a.isErr != b.isErr {
		return false
	}
	if a.isErr {
		if a.err != b.err {
			return false
		}
	} else if !equivEqual(a.val, b.val) {
		return false
	}
	if len(a.observed) != len(b.observed) {
		return false
	}
	for i := range a.observed {
		if !equivEqual(a.observed[i], b.observed[i]) {
			return false
		}
	}
	return true
}

// RunEquivCase runs c against the KL copy (equiv.NAME) and the native (NAME)
// of kernelFn. It returns "" when both sides agree, otherwise a description
// of the form "case=NAME kl=… native=…".
func RunEquivCase(e *ControlFlow, kernelFn string, c EquivCase) string {
	klFn := mustSymbol(MakeSymbol(EquivName(kernelFn))).function
	if klFn == nil {
		return fmt.Sprintf("case=%s kl=<%s not defined> native=<not run>", c.Name, EquivName(kernelFn))
	}
	nat := mustSymbol(MakeSymbol(kernelFn)).function
	if nat == nil {
		return fmt.Sprintf("case=%s kl=<not run> native=<%s not bound>", c.Name, kernelFn)
	}
	kl := runEquivSide(e, klFn, c)
	native := runEquivSide(e, nat, c)
	if sameOutcome(kl, native) {
		return ""
	}
	return fmt.Sprintf("case=%s kl=%s native=%s", c.Name, kl, native)
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

// IsNativeBinding reports whether f is a Go-implemented function object
// (MakePrimitive and MakeNative both build one) rather than KL bytecode or an
// interpreted procedure.
func IsNativeBinding(f Obj) bool {
	if f == nil || isFixnum(f) {
		return false
	}
	return *f == scmHeadNative
}
