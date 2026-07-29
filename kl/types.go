package kl

import (
	"bytes"
	"fmt"
	"io"
	"time"
	"unsafe"
)

// All kinds of Scheme object.
type Obj *scmHead

type scmHead int

const (
	scmHeadNumber    scmHead = 0
	scmHeadPair      scmHead = 1
	scmHeadVector    scmHead = 2
	scmHeadNull      scmHead = 3
	scmHeadString    scmHead = 4
	scmHeadSymbol    scmHead = 5
	scmHeadBoolean   scmHead = 6
	scmHeadProcedure scmHead = 14
	scmHeadStream    scmHead = 17
	scmHeadError     scmHead = 22
	scmHeadNative    scmHead = 23
	scmHeadRaw       scmHead = 42
)

type scmNumber struct {
	scmHead
	val float64
}

type scmSymbol struct {
	scmHead
	// The string of this symbol.
	str      string
	value    Obj
	function Obj
}

type scmPair struct {
	scmHead
	car Obj
	cdr Obj
}

// scmEnv is one binding frame in the tree-walking interpreter's lexical
// environment: a (sym -> val) binding plus a link to the enclosing env. The
// interpreter previously represented the env as an alist — cons(cons(sym,val),
// env) — which costs TWO heap allocations per binding. A scmEnv folds the
// binding and the link into a single node, so each let/lambda binding allocates
// once instead of twice. The env is fully encapsulated (built only by
// envExtend, walked only by envGet, otherwise threaded opaquely and never
// printed, compared, or exposed to Shen code), so this representation change is
// invisible everywhere else.
type scmEnv struct {
	scmHead
	sym  Obj
	val  Obj
	next Obj
}

const scmHeadEnv scmHead = 51

func envCons(sym, val, next Obj) Obj {
	tmp := &scmEnv{scmHeadEnv, sym, val, next}
	return &tmp.scmHead
}

func mustEnv(o Obj) *scmEnv {
	return (*scmEnv)(unsafe.Pointer(o))
}

type scmVector struct {
	scmHead
	vector []Obj
}

type scmString struct {
	scmHead
	str string
}

type scmStream struct {
	scmHead
	raw interface{}
}

type scmBoolean struct {
	scmHead
	bool
}

type scmProcedure struct {
	scmHead
	name  string
	arg   []Obj
	arity int
	body  Obj
	env   Obj
}

type scmNative struct {
	scmHead
	name     string
	fn       func(*ControlFlow)
	require  int
	captured []Obj
}

func MakeNative(fn func(*ControlFlow), require int, captured ...Obj) Obj {
	tmp := scmNative{
		scmHead:  scmHeadNative,
		fn:       fn,
		require:  require,
		captured: captured,
	}
	return &tmp.scmHead
}

func MustNative(o Obj) *scmNative {
	if *o != scmHeadNative {
		panic("mustNative")
	}
	return (*scmNative)(unsafe.Pointer(o))
}

type scmError struct {
	scmHead
	err string
}

// MakeRaw makes a struct into a raw object.
// Usage:
//
//	type T struct {
//	   scmHead int
//	   ... // xxx
//	}
//
// tmp := &T{}
// raw := MakeRaw(&tmp.scmHead)
func MakeRaw(scmHead *int) Obj {
	*scmHead = int(scmHeadRaw)
	return Obj(unsafe.Pointer(scmHead))
}

func MakeError(err string) Obj {
	tmp := scmError{scmHeadError, err}
	return &tmp.scmHead
}

func mustError(o Obj) *scmError {
	if *o != scmHeadError {
		panic(MakeError("mustError"))
	}
	return (*scmError)(unsafe.Pointer(o))
}

func IsError(o Obj) bool {
	return o != nil && *o == scmHeadError
}

func IsNumber(o Obj) bool {
	return *o == scmHeadNumber
}

func IsSymbol(o Obj) bool {
	return *o == scmHeadSymbol
}

func mustVector(o Obj) []Obj {
	if (*o) != scmHeadVector {
		panic(MakeError("mustVector"))
	}
	tmp := (*scmVector)(unsafe.Pointer(o))
	return tmp.vector
}

func mustProcedure(o Obj) *scmProcedure {
	if (*o) != scmHeadProcedure {
		panic(MakeError("mustProcedure"))
	}
	return (*scmProcedure)(unsafe.Pointer(o))
}

func mustString(o Obj) string {
	if (*o) != scmHeadString {
		panic(MakeError("mustString"))
	}
	return (*scmString)(unsafe.Pointer(o)).str
}

func fixnum(o Obj) int {
	return int(uintptr(unsafe.Pointer(o))-uintptr(fixnumBaseAddr)) + fixnumMin
}

// narrowToInt converts a Shen number's float64 to a Go int, raising an
// ordinary catchable Shen error rather than letting the conversion go
// out of range. Go leaves an out-of-range float64 -> int conversion
// implementation-defined; on arm64 it saturates, so +Inf used to reach the
// index-taking primitives as 9223372036854775807 and, past any bounds check
// they happened to have, as an uncatchable Go runtime panic.
func narrowToInt(f float64) int {
	if !fitsInt(f) {
		panic(MakeError(fmt.Sprintf("%s is not a valid integer", formatNumber(f))))
	}
	return int(f)
}

// mustByte narrows a Shen number to a byte for write-byte, raising an ordinary
// catchable Shen error unless it really is a whole number in 0..255.
//
// PrimWriteByte used to take mustInteger's answer and do byte(n), so
// (write-byte 321 S) silently wrote 0x41, (write-byte -1 S) wrote 0xFF, and
// (write-byte 65.5 S) wrote 'A' -- mustInteger checks range but not
// integrality. Tarver's Primitives/write-byte.lsp is (WRITE-BYTE Byte S), and
// CL's WRITE-BYTE signals a type error for anything outside (UNSIGNED-BYTE 8),
// so an out-of-range byte has to raise here too.
//
// A non-integer or unnarrowable value keeps D4's "is not a valid integer"
// wording; only the range check is new.
func mustByte(o Obj) int {
	if (*o) != scmHeadNumber {
		panic(MakeError("mustNumber"))
	}
	f := GetNumber(o)
	if !isPreciseInteger(f) || !fitsInt(f) {
		panic(MakeError(fmt.Sprintf("%s is not a valid integer", formatNumber(f))))
	}
	if f < 0 || f > 255 {
		panic(MakeError(fmt.Sprintf("%s is not a byte", formatNumber(f))))
	}
	return int(f)
}

func mustInteger(o Obj) int {
	if (*o) != scmHeadNumber {
		panic(MakeError("mustNumber"))
	}
	if isFixnum(o) {
		return fixnum(o)
	}

	return narrowToInt((*scmNumber)(unsafe.Pointer(o)).val)
}

func GetInteger(o Obj) int {
	if isFixnum(o) {
		return fixnum(o)
	}
	return narrowToInt((*scmNumber)(unsafe.Pointer(o)).val)
}

// GetNumber returns o's numeric value without truncating it. Shen numbers are
// float64; GetInteger narrows to int, which silently discards the fractional
// part, so anything that must round-trip a number (code generation, for one)
// has to use this instead.
func GetNumber(o Obj) float64 {
	if isFixnum(o) {
		return float64(fixnum(o))
	}
	return (*scmNumber)(unsafe.Pointer(o)).val
}

func mustNumber(o Obj) float64 {
	if (*o) != scmHeadNumber {
		panic(MakeError("mustNumber"))
	}
	if isFixnum(o) {
		return float64(fixnum(o))
	}
	x := (*scmNumber)(unsafe.Pointer(o))
	return x.val
}

func mustSymbol(o Obj) *scmSymbol {
	if (*o) != scmHeadSymbol {
		panic(MakeError("mustSymbol"))
	}
	return (*scmSymbol)(unsafe.Pointer(o))
}

func isSymbol(o Obj) (bool, *scmSymbol) {
	if *o == scmHeadSymbol {
		return true, (*scmSymbol)(unsafe.Pointer(o))
	}
	return false, nil
}

func mustStream(o Obj) *scmStream {
	if (*o) != scmHeadStream {
		panic(MakeError("mustStream"))
	}
	return (*scmStream)(unsafe.Pointer(o))
}

func mustPair(o Obj) *scmPair {
	if (*o) != scmHeadPair {
		fmt.Println(ObjString(o))
		panic(MakeError("mustPair"))
	}
	return (*scmPair)(unsafe.Pointer(o))
}

func isPair(o Obj) (bool, *scmPair) {
	if (*o) == scmHeadPair {
		return true, (*scmPair)(unsafe.Pointer(o))
	}
	return false, nil
}

var True, False, Nil, undefined Obj
var uptime time.Time
var symQuote, symDefun, symLambda, symFreeze, symLet, symAnd Obj
var symOr, symIf, symCond, symTrapError, symDo, symMacroExpand Obj
var symType Obj

// Arithmetic intrinsic symbols for the compiler fast-path detection.
var symAdd, symSub, symMul, symLT, symLE, symGT, symGE, symNumEq, symNot Obj

// Fixnum representation: small integers are encoded as pointers into a dedicated
// sentinel byte array, costing zero heap allocation. The array's contents are
// never read or written — only the *addresses* of its bytes are used — so the
// array is pure virtual address space (BSS): it is never paged in and costs ~0
// RSS regardless of size, and the GC never scans it (a [N]byte has no pointers).
//
// The range is signed and centered, so the byte at offset 0 represents
// fixnumMin. Widening from the original unsigned [0, 2^20) to a signed
// [-2^25, 2^25) removes heap boxing for two large classes of common integers
// that previously always allocated a scmNumber: every negative integer, and
// every integer between 2^20 and 2^25. makeInteger (the boxed fallback) was the
// single largest allocation source in the VM (per the alloc profile of integer
// arithmetic), so widening its fixnum fast path directly cuts GC pressure for
// real integer workloads.
const (
	fixnumBits  = 26                       // sentinel array is 2^26 bytes = 64 MiB of pure address space
	fixnumCount = 1 << fixnumBits          // number of distinct fixnum values
	fixnumMin   = -(1 << (fixnumBits - 1)) // smallest fixnum, i.e. -2^25
	fixnumMax   = 1 << (fixnumBits - 1)    // one past the largest fixnum, i.e. 2^25
)

var addrForFixnum [fixnumCount]byte

// fixnumBaseAddr is the address representing the integer fixnumMin.
var fixnumBaseAddr = unsafe.Pointer(&addrForFixnum[0])
var fixnumEndAddr = unsafe.Add(fixnumBaseAddr, fixnumCount)

type trieNode struct {
	children [256]*trieNode
	value    scmSymbol
}

var symbolRoot trieNode

func trieFindOrInsert(str string) *trieNode {
	p := &symbolRoot
	for i := 0; i < len(str); i++ {
		v := str[i]
		if p.children[v] == nil {
			p.children[v] = &trieNode{
				value: scmSymbol{
					scmHead: scmHeadSymbol,
					str:     str[:i+1],
				},
			}
		}
		p = p.children[v]
	}
	return p
}

func init() {
	uptime = time.Now()
	tmp1 := &scmBoolean{scmHeadBoolean, false}
	False = Obj(&tmp1.scmHead)

	tmp2 := &scmBoolean{scmHeadBoolean, true}
	True = Obj(&tmp2.scmHead)

	tmp3 := &scmPair{scmHeadNull, nil, nil}
	Nil = Obj(&tmp3.scmHead)

	var tmp4 int
	undefined = MakeRaw(&tmp4)

	symQuote = MakeSymbol("quote")
	symDefun = MakeSymbol("defun")
	symLambda = MakeSymbol("lambda")
	symFreeze = MakeSymbol("freeze")
	symLet = MakeSymbol("let")
	symAnd = MakeSymbol("and")
	symOr = MakeSymbol("or")
	symIf = MakeSymbol("if")
	symCond = MakeSymbol("cond")
	symTrapError = MakeSymbol("trap-error")
	symDo = MakeSymbol("do")
	symMacroExpand = MakeSymbol("macroexpand")
	_ = symMacroExpand // mark as used to satisfy staticcheck
	symType = MakeSymbol("type")
	symAdd = MakeSymbol("+")
	symSub = MakeSymbol("-")
	symMul = MakeSymbol("*")
	symLT = MakeSymbol("<")
	symLE = MakeSymbol("<=")
	symGT = MakeSymbol(">")
	symGE = MakeSymbol(">=")
	symNumEq = MakeSymbol("=")
	symNot = MakeSymbol("not")
}

func MakeInteger(v int) Obj {
	if v >= fixnumMin && v < fixnumMax {
		return Obj(unsafe.Pointer(unsafe.Add(fixnumBaseAddr, v-fixnumMin)))
	}
	return makeInteger(v)
}

func isFixnum(o Obj) bool {
	v := uintptr(unsafe.Pointer(o))
	if v >= uintptr(fixnumBaseAddr) && v < uintptr(fixnumEndAddr) {
		return true
	}
	return false
}

func makeInteger(v int) Obj {
	tmp := scmNumber{scmHeadNumber, float64(v)}
	return &tmp.scmHead
}

// Bounds of the int representation, as float64. maxIntAsFloat is 2^63 --
// one past math.MaxInt64, which is not itself representable as a float64.
const (
	minIntAsFloat = -9223372036854775808.0
	maxIntAsFloat = 9223372036854775808.0
)

func MakeNumber(f float64) Obj {
	// A float beyond the int range is still mathematically integral, but
	// narrowing it overflows -- int(1e300) saturates to maxint64, turning the
	// value into a different one. Keep those as float64 instead.
	if isPreciseInteger(f) && f >= minIntAsFloat && f < maxIntAsFloat {
		return MakeInteger(int(f))
	}

	tmp := scmNumber{scmHeadNumber, f}
	return &tmp.scmHead
}

func MakeStream(raw interface{}) Obj {
	tmp := scmStream{
		scmHeadStream,
		raw,
	}
	return &tmp.scmHead
}

func IsString(o Obj) bool {
	return *o == scmHeadString
}

func GetString(o Obj) string {
	return mustString(o)
}

func GetSymbol(o Obj) string {
	return mustSymbol(o).str
}

func cons(x, y Obj) Obj {
	tmp := scmPair{
		scmHead: scmHeadPair,
		car:     x,
		cdr:     y,
	}
	return &tmp.scmHead
}

func car(x Obj) Obj {
	return mustPair(x).car
}

func cdr(x Obj) Obj {
	return mustPair(x).cdr
}

func MakeVector(n int) Obj {
	tmp := scmVector{
		scmHeadVector,
		make([]Obj, n),
	}
	return &tmp.scmHead
}

func MakeString(s string) Obj {
	tmp := scmString{scmHeadString, s}
	return &tmp.scmHead
}

func MakeSymbol(s string) Obj {
	p := trieFindOrInsert(s)
	return &p.value.scmHead
}

func makeProcedure(arg Obj, body Obj, env Obj) Obj {
	tmp := scmProcedure{
		scmHead: scmHeadProcedure,
		body:    body,
		env:     env,
	}
	if *arg == scmHeadSymbol {
		tmp.arg = []Obj{arg}
		tmp.arity = 1
	} else {
		tmp.arg = ListToSlice(arg)
		tmp.arity = len(tmp.arg)
	}
	return &tmp.scmHead
}

func ObjString(o Obj) string {
	return (*scmHead)(o).GoString()
}

func (o *scmPair) fmt(buf io.Writer, start bool) {
	if start {
		fmt.Fprintf(buf, "(%s", ObjString(o.car))
	} else {
		fmt.Fprintf(buf, " %s", ObjString(o.car))
	}
	switch *o.cdr {
	case scmHeadNull:
		fmt.Fprintf(buf, ")")
	case scmHeadPair:
		mustPair(o.cdr).fmt(buf, false)
	default:
		fmt.Fprintf(buf, " . %s)", ObjString(o.cdr))
	}
}

func (o *scmHead) GoString() string {
	switch *o {
	case scmHeadNumber:
		// formatNumber keeps the integral-fits-in-int case printing without
		// a decimal point and everything else (2.5, 1e19, inf) in a form
		// that does not silently become a different number.
		return formatNumber(mustNumber(o))
	case scmHeadPair:
		var buf bytes.Buffer
		mustPair(o).fmt(&buf, true)
		return buf.String()
	case scmHeadVector:
		return "#vector"
	case scmHeadNull:
		return "()"
	case scmHeadString:
		return fmt.Sprintf(`"%s"`, mustString(o))
	case scmHeadSymbol:
		return GetSymbol(o)
	case scmHeadBoolean:
		switch o {
		case True:
			return "true"
		case False:
			return "false"
		default:
			return "Boolean(something wrong)"
		}

	case scmHeadError:
		return fmt.Sprintf("Error(%s)", mustError(o).err)
	case scmHeadProcedure:
		return "#procedure"
	case scmHeadStream:
		return "#stream"
	case scmHeadRaw:
		return "#raw"
	case scmHeadNative:
		prim := MustNative(o)
		if len(prim.name) > 0 {
			return fmt.Sprintf("#primitive(%s)", prim.name)
		}
		return "#native"
	case scmHeadBytecodeFunc:
		bf := mustBytecodeFunc(o)
		if bf.fn.Name != "" {
			return fmt.Sprintf("#compiled(%s)", bf.fn.Name)
		}
		return "#compiled"
	}
	return fmt.Sprintf("unknown type %d", *o)
}
