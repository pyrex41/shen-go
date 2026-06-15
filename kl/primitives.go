package kl

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"
)

var home_directory Obj
var stinputEOF bool

// ResetStinputEOF clears the stdin-EOF flag. The REPL calls this before each
// read so a fresh prompt starts from a clean state.
func ResetStinputEOF() {
	stinputEOF = false
}

// StinputEOF reports whether the last read on *stinput* hit EOF. The CLI uses
// this to exit the top-level loop cleanly when piped input is exhausted,
// instead of looping forever on "error: empty stream".
func StinputEOF() bool {
	return stinputEOF
}

var klPrimitives = []struct {
	name  string
	arity int
	fn    interface{}
}{
	// &scmNative{scmHead: scmHeadNative, name: "type", require: 2, fn: PrimTypeFunc},
	{"get-time", 1, PrimGetTime},
	{"close", 1, PrimCloseStream},
	{"open", 2, PrimOpenStream},
	{"read-byte", 1, PrimReadByte},
	{"write-byte", 2, PrimWriteByte},
	{"absvector?", 1, PrimIsVector},
	{"<-address", 2, PrimVectorGet},
	{"address->", 3, PrimVectorSet},
	{"absvector", 1, PrimAbsvector},
	{"str", 1, PrimStr},
	{"<=", 2, PrimLessEqual},
	{">=", 2, PrimGreatEqual},
	{"<", 2, PrimLessThan},
	{">", 2, PrimGreatThan},
	{"error-to-string", 1, PrimErrorToString},
	{"simple-error", 1, PrimSimpleError},
	{"=", 2, PrimEqual},
	{"-", 2, PrimNumberSubtract},
	{"*", 2, PrimNumberMultiply},
	{"/", 2, PrimNumberDivide},
	{"+", 2, PrimNumberAdd},
	{"string->n", 1, PrimStringToNumber},
	{"n->string", 1, PrimNumberToString},
	{"number?", 1, PrimIsNumber},
	{"string?", 1, PrimIsString},
	{"pos", 2, PrimPos},
	{"tlstr", 1, PrimTailString},
	{"cn", 2, PrimStringConcat},
	{"intern", 1, PrimIntern},
	{"hd", 1, PrimHead},
	{"tl", 1, PrimTail},
	{"cons", 2, PrimCons},
	{"cons?", 1, PrimIsPair},
	{"value", 1, PrimValue},
	{"set", 2, PrimSet},
	{"not", 1, PrimNot},
	{"if", 3, PrimIf},
	{"symbol?", 1, PrimIsSymbol},
	{"read-file-as-bytelist", 1, PrimReadFileAsByteList},
	{"read-file-as-string", 1, PrimReadFileAsString},
	{"variable?", 1, PrimIsVariable},
	{"integer?", 1, PrimIsInteger},
	{"shen.char-stoutput?", 1, PrimCharStOutput},
	{"shen.char-stinput?", 1, PrimCharStInput},
}

func init() {
	for _, item := range klPrimitives {
		sym := MakeSymbol(item.name)
		prim := MakePrimitive(item.name, item.arity, item.fn)
		BindSymbolFunc(sym, prim)
	}

	// Overload for primitive set and value.
	BindSymbolFunc(MakeSymbol("eval-kl"), MakeNative(primEvalKL, 1))
	BindSymbolFunc(MakeSymbol("load-file"), MakeNative(primLoadFile, 1))

	PrimSet(MakeSymbol("*stinput*"), MakeStream(os.Stdin))
	PrimSet(MakeSymbol("*stoutput*"), MakeStream(os.Stdout))

	home_directory = MakeSymbol("*home-directory*")
	PrimSet(home_directory, MakeString(""))
	PrimSet(MakeSymbol("*language*"), MakeString("Go"))
	PrimSet(MakeSymbol("*implementation*"), MakeString("AOT+interpreter"))
	PrimSet(MakeSymbol("*release*"), MakeString(runtime.Version()))
	PrimSet(MakeSymbol("*os*"), MakeString(runtime.GOOS))
	PrimSet(MakeSymbol("*porters*"), MakeString("Arthur Mao"))
	PrimSet(MakeSymbol("*port*"), MakeString("1.0.0-rc1"))

	// Extended by shen-go implementation
	PrimSet(MakeSymbol("*package-path*"), MakeString(PackagePath()))

	BindSymbolFunc(MakeSymbol("defun"), MakePrimitive("defun", 2, primDefun))
	BindSymbolFunc(MakeSymbol("try-catch"), MakeNative(primTryCatch, 2))
}

func PrimCharStInput(x Obj) Obj {
	return False
}

func PrimCharStOutput(x Obj) Obj {
	return False
}

func PrimNumberAdd(x, y Obj) Obj {
	x1 := mustNumber(x)
	y1 := mustNumber(y)
	return MakeNumber(x1 + y1)
}

func PrimNumberSubtract(x, y Obj) Obj {
	x1 := mustNumber(x)
	y1 := mustNumber(y)
	return MakeNumber(x1 - y1)
}

func PrimNumberMultiply(x, y Obj) Obj {
	x1 := mustNumber(x)
	y1 := mustNumber(y)
	return MakeNumber(x1 * y1)
}

func PrimNumberDivide(x, y Obj) Obj {
	x1 := mustNumber(x)
	y1 := mustNumber(y)
	// Issue #10: guard divide-by-zero. Go's float64 division by 0 yields
	// +Inf/-Inf/NaN which, after truncation to int in the number->string
	// path, surfaced as a bogus maxint (9223372036854775807) and was not
	// catchable by trap-error. Raise the kernel's standard catchable error
	// instead, matching shen-cl's divide-by-zero behavior and the sibling
	// ports' simple-error mechanism.
	if y1 == 0 {
		panic(MakeError("division by zero"))
	}
	return MakeNumber(x1 / y1)
}

func PrimIntern(o Obj) Obj {
	str := mustString(o)
	switch str {
	case "true":
		return True
	case "false":
		return False
	}
	return MakeSymbol(str)
}

func PrimHead(o Obj) Obj {
	return car(o)
}

func PrimTail(o Obj) Obj {
	return cdr(o)
}

func PrimIsNumber(o Obj) Obj {
	if *o == scmHeadNumber {
		return True
	}
	return False
}

func PrimStringToNumber(o Obj) Obj {
	str := mustString(o)
	n := ([]rune(str))[0]
	return MakeInteger(int(n))
}

func PrimNumberToString(o Obj) Obj {
	n := mustInteger(o)
	return MakeString(string(rune(n)))
}

func PrimStr(o Obj) Obj {
	switch *o {
	case scmHeadPair:
		// Pair may contain recursive list.
		panic(MakeError("can't str pair object"))
	case scmHeadNull:
		return MakeString("()")
	case scmHeadSymbol:
		str := GetSymbol(o)
		return MakeString(str)
	case scmHeadNumber:
		if isFixnum(o) {
			return MakeString(fmt.Sprintf("%d", mustInteger(o)))
		}
		f := mustNumber(o)
		if !isPreciseInteger(f) {
			// Issue #11: render with Go's shortest round-trip form
			// (2.5, not 2.500000) to match shen-cl, shen-rust and
			// ShenScript. Integral-valued floats keep printing without
			// a decimal point below, as the kernel expects.
			return MakeString(strconv.FormatFloat(f, 'g', -1, 64))
		}
		return MakeString(fmt.Sprintf("%d", int(f)))
	case scmHeadString:
		return MakeString(fmt.Sprintf(`"%s"`, mustString(o)))
	case scmHeadProcedure:
		return MakeString("#<procedure >")
	case scmHeadBytecodeFunc:
		bf := mustBytecodeFunc(o)
		if bf.fn.Name != "" {
			return MakeString("#<compiled " + bf.fn.Name + ">")
		}
		return MakeString("#<compiled >")
	case scmHeadBoolean:
		switch o {
		case True:
			return MakeString("true")
		case False:
			return MakeString("false")
		}
	case scmHeadError:
		err := mustError(o)
		return MakeString("#<err " + err.err + ">")
	case scmHeadStream:
		return MakeString("#<stream >")
	case scmHeadRaw:
		return MakeString("#<raw >")
	case scmHeadNative:
		n := MustNative(o)
		if len(n.name) > 0 {
			return MakeString(fmt.Sprintf("#<primitive %s>", n.name))
		} else {
			return MakeString(fmt.Sprintf("#<native %v>", *n))
		}
	default:
		return MakeString("PrimStr unknown")

	}
	return MakeString("wrong input, the object is not atom ...")
}

func PrimStringConcat(x, y Obj) Obj {
	s1 := mustString(x)
	s2 := mustString(y)
	return MakeString(s1 + s2)
}

func PrimTailString(o Obj) Obj {
	str := mustString(o)
	if len(str) == 0 {
		panic(MakeError("empty string"))
	}
	return MakeString(string([]rune(str)[1:]))
}

func PrimPos(x, y Obj) Obj {
	s := []rune(mustString(x))
	n := mustInteger(y)
	if n >= len(s) {
		panic(MakeError(fmt.Sprintf("%d is not valid index for %s", n, string(s))))
	}
	return MakeString(string([]rune(s)[n]))
}

func and(x, y Obj) Obj {
	if x == True && y == True {
		return True
	}
	return False
}

func or(x, y Obj) Obj {
	if x == False || y == False {
		return False
	}
	return True
}

func PrimSet(key, val Obj) Obj {
	sym := mustSymbol(key)
	sym.value = val
	return val
}

func PrimFunc(key Obj) Obj {
	sym := mustSymbol(key)
	if sym.function != nil {
		return sym.function
	}
	panic(MakeError(fmt.Sprintf("variable %s not bound", sym.str)))
}

func PrimValue(o Obj) Obj {
	key := o
	sym := mustSymbol(key)
	if sym.value != nil {
		return sym.value
	}
	panic(MakeError(fmt.Sprintf("variable %s not bound", sym.str)))
}

func PrimSimpleError(o Obj) Obj {
	str := mustString(o)
	panic(MakeError(str))
}

func PrimErrorToString(o Obj) Obj {
	err := mustError(o)
	return MakeString(err.err)
}

func PrimGreatThan(x, y Obj) Obj {
	x1 := mustNumber(x)
	y1 := mustNumber(y)
	if x1 > y1 {
		return True
	}
	return False
}

func PrimLessThan(a, b Obj) Obj {
	x := mustNumber(a)
	y := mustNumber(b)
	if x < y {
		return True
	}
	return False
}

func PrimLessEqual(a, b Obj) Obj {
	x := mustNumber(a)
	y := mustNumber(b)
	if x <= y {
		return True
	}
	return False
}

func PrimGreatEqual(a, b Obj) Obj {
	x := mustNumber(a)
	y := mustNumber(b)
	if x >= y {
		return True
	}
	return False
}

// maxAbsvectorSize bounds the element count `absvector` will allocate. A
// vector holds n pointers (n*8 bytes on 64-bit), so an attacker-controlled n
// turns `(absvector N)` into an arbitrarily large make() that aborts the
// whole process with an unrecoverable Go "out of memory" fatal — uncatchable
// by trap-error and fatal to any host embedding the evaluator (it was the
// last way a single fuzzer-reachable form could kill the worker). The ceiling
// is far larger than any realistic Shen vector (records, hash buckets) yet
// turns the absurd case into an ordinary catchable Shen error instead.
const maxAbsvectorSize = 1 << 28 // 268,435,456 elements

func PrimAbsvector(o Obj) Obj {
	n := mustInteger(o)
	if n < 0 {
		panic(MakeError("absvector wrong argument"))
	}
	if n > maxAbsvectorSize {
		panic(MakeError(fmt.Sprintf("absvector size %d exceeds maximum %d", n, maxAbsvectorSize)))
	}
	return MakeVector(n)
}

func PrimVectorSet(x, y, z Obj) Obj {
	vec := mustVector(x)
	off := mustInteger(y)
	val := z
	vec[off] = val
	return x
}

func PrimVectorGet(x, y Obj) Obj {
	vec := mustVector(x)
	off := mustInteger(y)
	if off >= len(vec) {
		panic(MakeError(fmt.Sprintf("index %d out of range %d", off, len(vec))))
	}
	ret := vec[off]
	if ret == nil {
		return undefined
	}
	return ret
}

func PrimIsVector(x Obj) Obj {
	if *x == scmHeadVector {
		return True

	}
	return False
}

func PrimWriteByte(x, y Obj) Obj {
	n := mustInteger(x)
	s := mustStream(y)
	w, ok := s.raw.(io.Writer)
	if !ok {
		panic(MakeError("stream is not opened in out mode"))
	}
	var b [1]byte
	b[0] = byte(n)
	_, err := w.Write(b[:])
	if err != nil {
		panic(MakeError(err.Error()))
	}
	return x
}

func PrimReadByte(x Obj) Obj {
	s := mustStream(x)
	r, ok := s.raw.(io.Reader)
	if !ok {
		panic(MakeError("stream is closed of not opened in in mode"))
	}
	var buf [1]byte
	n, err := r.Read(buf[:])
	if n > 0 {
		return MakeInteger(int(buf[0]))
	}
	if err != nil {
		if err == io.EOF {
			if x == PrimValue(MakeSymbol("*stinput*")) {
				stinputEOF = true
			}
			return MakeInteger(-1)
		}
		panic(MakeError(err.Error()))
	}
	return MakeInteger(int(buf[0]))
}

func PrimOpenStream(x, y Obj) Obj {
	file := mustString(x)
	var flag int
	mode := GetSymbol(y)
	switch mode {
	case "in":
		flag |= os.O_RDONLY
	case "out":
		// O_TRUNC: `out` opens for writing from the start, replacing any
		// existing contents. Without it, writing fewer bytes than a file
		// already holds leaves the old tail behind (e.g. re-bootstrapping a
		// .kl over a longer one left trailing bytes -> malformed KL).
		flag |= os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	default:
		flag = os.O_RDWR | os.O_CREATE
	}

	file = ResolveHomePath(file)
	f, err := os.OpenFile(file, flag, 0666)
	if err != nil {
		panic(MakeError(err.Error()))
	}
	return MakeStream(f)
}

// ResolveHomePath resolves a file path the way the kernel `open` primitive does:
// by prepending the current *home-directory* value (literal string
// concatenation — *home-directory* is expected to already end with a
// separator). Native code that reads files WITHOUT going through `open` (the
// load cache and the Go-native read-file) MUST resolve through this so they
// touch exactly the file the interpreted reader would; otherwise a Shen program
// that changed *home-directory* (e.g. via cd) would make a relative (load "F")
// silently read F from the process CWD instead of *home-directory*/F.
func ResolveHomePath(file string) string {
	return GetString(PrimValue(home_directory)) + file
}

func PrimCloseStream(x Obj) Obj {
	s := mustStream(x)
	c := s.raw.(io.Closer)
	c.Close()
	return Nil
}

func PrimGetTime(x Obj) Obj {
	kind := GetSymbol(x)
	switch kind {
	case "unix":
		return MakeNumber(float64(time.Now().Unix()))
	case "run":
		return MakeNumber(time.Since(uptime).Seconds())

	}
	panic(MakeError(fmt.Sprintf("get-time does not understand the parameter %s", kind)))
}

func PrimIsString(x Obj) Obj {
	if *x == scmHeadString {
		return True

	}
	return False
}

func PrimIsPair(x Obj) Obj {
	if *x == scmHeadPair {
		return True
	}
	return False
}

func PrimNot(x Obj) Obj {
	switch x {
	case False:
		return True
	case True:
		return False
	}
	panic(MakeError("PrimNot"))
}

func PrimIf(x, y, z Obj) Obj {
	switch x {
	case True:
		return y
	case False:
		return z
	}
	panic(MakeError("PrimIf"))
}

func PrimEqual(x, y Obj) Obj {
	return equal(x, y)
}

func PrimCons(x, y Obj) Obj {
	return cons(x, y)
}

func PrimIsSymbol(x Obj) Obj {
	if IsSymbol(x) {
		return True

	}
	return False
}

func PrimReadFileAsByteList(x Obj) Obj {
	fileName := mustString(x)
	buf, err := os.ReadFile(fileName)
	if err != nil {
		panic(MakeError(err.Error()))
	}

	ret := cons(Nil, Nil)
	curr := mustPair(ret)
	for _, b := range buf {
		tmp := cons(MakeInteger(int(b)), Nil)
		curr.cdr = tmp
		curr = mustPair(tmp)
	}
	return cdr(ret)
}

func PrimReadFileAsString(x Obj) Obj {
	fileName := mustString(x)
	buf, err := os.ReadFile(fileName)
	if err != nil {
		panic(MakeError(err.Error()))
	}
	return MakeString(string(buf))
}

// PrimHash is a native replacement for the kernel's `hash` (sys.kl), which
// builds a product of char codes (overflows past ~8 chars) and then takes a
// modulo by repeated subtraction of a doubling list of multiples — slow and
// lossy. We hash the key's string form with FNV-1a and take a true modulo.
//
// Hash values are never persisted across runs, so only determinism and
// distribution matter, not matching the KL value: equal keys still land in the
// same bucket, and unequal keys that collide are separated by assoc. The result
// is in [1, K-1] to honour the kernel's bucket-index contract (offset = 3+hash
// into an absvector of size 3+K), with 0 mapped to 1 exactly as the KL hash does.
func PrimHash(v, k Obj) Obj {
	limit := mustInteger(k)
	if limit <= 1 {
		return MakeInteger(1)
	}
	s := hashKeyString(v)
	var h uint64 = 14695981039346656037 // FNV-1a 64-bit offset basis
	for i := 0; i < len(s); i++ {
		h ^= uint64(s[i])
		h *= 1099511628211 // FNV-1a 64-bit prime
	}
	m := int(h % uint64(limit))
	if m == 0 {
		m = 1
	}
	return MakeInteger(m)
}

func hashKeyString(v Obj) string {
	switch {
	case IsString(v):
		return GetString(v)
	case IsSymbol(v):
		return GetSymbol(v)
	default:
		return ObjString(v)
	}
}

func PrimIsVariable(x Obj) Obj {
	if *x != scmHeadSymbol {
		return False
	}

	sym := GetSymbol(x)
	if len(sym) == 0 || sym[0] < 'A' || sym[0] > 'Z' {
		return False
	}
	return True
}

func PrimIsInteger(x Obj) Obj {
	if *x != scmHeadNumber {
		return False

	}
	if isFixnum(x) {
		return True

	}
	f := mustNumber(x)
	if isPreciseInteger(f) {
		return True

	}
	return False
}

func primEvalKL(e *ControlFlow) {
	res := evalExp(e, e.Get(1), Nil)
	e.Return(res)
}

func primLoadFile(e *ControlFlow) {
	path := mustString(e.Get(1))
	res := loadFile(e, path)
	e.Return(res)
}

func loadFile(e *ControlFlow, file string) Obj {
	var filePath string
	if _, err := os.Stat(file); err == nil {
		filePath = file
	} else {
		filePath = path.Join(PackagePath(), file)
		if _, err := os.Stat(filePath); err != nil {
			return MakeError(err.Error())
		}
	}

	f, err := os.Open(filePath)
	if err != nil {
		return MakeError(err.Error())
	}
	defer f.Close()

	r := NewSexpReader(f, false)
	for {
		exp, err := r.Read()
		if err != nil {
			if err != io.EOF {
				return MakeError(err.Error())
			}
			break
		}

		res := evalExp(e, exp, Nil)
		if *res == scmHeadError {
			return res
		}
	}
	return MakeString(file)
}

// func readFileAsSexp(e Evaluator) {
// 	filePath := mustString(e.Get(1))
// 	f, err := os.Open(filePath)
// 	if err != nil {
// 		e.Return(MakeError(err.Error()))
// 		return
// 	}
// 	defer f.Close()

// 	ret := Nil
// 	r := NewSexpReader(f, true)
// 	for {
// 		exp, err := r.Read()
// 		if err != nil {
// 			if err != io.EOF {
// 				e.Return(MakeError(err.Error()))
// 				return
// 			}
// 			break
// 		}
// 		ret = cons(exp, ret)
// 	}
// 	e.Return(reverse(ret))
// 	return
// }

// func writeSexpToFile(e Evaluator) {
// 	filePath := mustString(e.Get(1))
// 	str := ObjString(e.Get(2))
// 	err := ioutil.WriteFile(filePath, []byte(str), 0644)
// 	if err != nil {
// 		e.Return(MakeError(err.Error()))
// 		return
// 	}
// 	e.Return(Nil)
// 	return
// }

func wrapf1(f func(Obj) Obj) func(*ControlFlow) {
	return func(e *ControlFlow) {
		tmp := e.Get(1)
		res := f(tmp)
		e.Return(res)
	}
}

func wrapf2(f func(x, y Obj) Obj) func(*ControlFlow) {
	return func(e *ControlFlow) {
		tmp1 := e.Get(1)
		tmp2 := e.Get(2)
		res := f(tmp1, tmp2)
		e.Return(res)
	}
}

func wrapf3(f func(x, y, z Obj) Obj) func(*ControlFlow) {
	return func(e *ControlFlow) {
		tmp1 := e.Get(1)
		tmp2 := e.Get(2)
		tmp3 := e.Get(3)
		res := f(tmp1, tmp2, tmp3)
		e.Return(res)
	}
}

func wrapf4(f func(x, y, z Obj) Obj) func(*ControlFlow) {
	return func(e *ControlFlow) {
		tmp1 := e.Get(1)
		tmp2 := e.Get(2)
		tmp3 := e.Get(3)
		res := f(tmp1, tmp2, tmp3)
		e.Return(res)
	}
}

func MakePrimitive(name string, arity int, f interface{}) Obj {
	var fn func(*ControlFlow)
	switch arity {
	case 1:
		raw, ok := f.(func(Obj) Obj)
		if !ok {
			panic(fmt.Sprintf("MakePrimitive %s failed: %#v", name, f))
		}
		fn = wrapf1(raw)
	case 2:
		raw, ok := f.(func(x, y Obj) Obj)
		if !ok {
			panic(fmt.Sprintf("MakePrimitive %s failed: %#v", name, f))
		}
		fn = wrapf2(raw)
	case 3:
		raw, ok := f.(func(x, y, z Obj) Obj)
		if !ok {
			panic(fmt.Sprintf("MakePrimitive %s failed: %#v", name, f))
		}
		fn = wrapf3(raw)
	default:
		panic(fmt.Sprintf("MakePrimitive %s failed: %#v", name, f))
	}
	tmp := &scmNative{
		scmHead: scmHeadNative,
		name:    name,
		require: arity,
		fn:      fn,
	}
	return Obj(&tmp.scmHead)
}

func primDefun(key, val Obj) Obj {
	sym := mustSymbol(key)
	// If given a raw scmProcedure (from the Shen-level compiler), compile it.
	if *val == scmHeadProcedure {
		proc := mustProcedure(val)
		compiled := CompileFunc(sym.str, proc.arg, proc.body)
		sym.function = compiled
		return key
	}
	sym.function = val
	return key
}

func primTryCatch(e *ControlFlow) {
	exp := e.Get(1)
	act := e.Get(2)
	res := Try(e, exp).Catch(act)
	e.Return(res)
}
