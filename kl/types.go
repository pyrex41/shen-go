package kl

import (
	"bytes"
	"fmt"
	"io"
	"os"
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
	// canonical is the first primitive registered under this name (see
	// primitiveRegistrar.register); nil means no primitive was ever
	// registered for it. HasCanonicalPrimitiveBinding compares function
	// against it with two loads instead of a locked string-keyed map lookup.
	// Written once, under primitiveRegistry.mu, by register() during
	// registration (package init / InstallKernelFast, single-threaded boot);
	// read unlocked like function.
	canonical Obj
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
	return o != nil && !isFixnum(o) && *o == scmHeadError
}

func IsNumber(o Obj) bool {
	return o != nil && (isFixnum(o) || *o == scmHeadNumber)
}

func IsSymbol(o Obj) bool {
	return o != nil && !isFixnum(o) && *o == scmHeadSymbol
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
	if !IsNumber(o) {
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
	if !IsNumber(o) {
		panic(MakeError("mustNumber"))
	}
	if isFixnum(o) {
		return fixnum(o)
	}

	return narrowToInt((*scmNumber)(unsafe.Pointer(o)).val)
}

// mustIndex narrows a vector index for <-address and address->. mustInteger
// truncates a fractional number, so (<-address V -0.5) used to read slot 0
// (the limit) and (address-> V 0.5 X) used to overwrite it; the kernel's
// <-vector and vector-> only guard against an exact 0, so both leaked
// through. A fractional index is an ordinary catchable error instead, with
// the wording kl/narrowing_test.go pins for inf and out-of-range indices.
func mustIndex(o Obj) int {
	if isFixnum(o) {
		return fixnum(o)
	}
	f := mustNumber(o)
	if !isPreciseInteger(f) {
		panic(MakeError(fmt.Sprintf("%s is not a valid integer", formatNumber(f))))
	}
	return narrowToInt(f)
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
	if o == nil {
		panic(MakeError("mustNumber"))
	}
	if isFixnum(o) {
		return float64(fixnum(o))
	}
	if (*o) != scmHeadNumber {
		panic(MakeError("mustNumber"))
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
		panic(MakeError("mustPair"))
	}
	return (*scmPair)(unsafe.Pointer(o))
}

func isPair(o Obj) (bool, *scmPair) {
	if o != nil && !isFixnum(o) && (*o) == scmHeadPair {
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

// Fixnum representation: small integers are encoded as pointers into a
// dedicated span of address space, costing zero heap allocation. Only the
// *addresses* are used; the span is zero-filled read-only memory, so code that
// peeks at *o on a fixnum reads scmHeadNumber (0), which is what the rest of the
// package relies on. The GC never scans it and never follows pointers into it.
//
// The range is signed and centered, so the byte at offset 0 represents
// fixnumMin. On 64-bit unix the span is an anonymous PROT_READ mapping reserved
// at startup (reserveFixnumSpace): 2^36 bytes, so fixnums cover [-2^35, 2^35),
// which holds every 32-bit word and sums of several of them. That is the range
// hashing and PRNG code in pure Shen (SHA-256 words, LCG state) lives in, and
// where every intermediate used to be a heap-boxed scmNumber. A read-only
// private mapping is not committed memory and is only paged in as the shared
// zero page where read, so the reservation costs address space, not RSS. If
// the reservation fails (32-bit, other OSes, a tight ulimit -v) the span falls
// back to the static 2^26-byte array below, i.e. [-2^25, 2^25) as before.
//
// fixnumMin and fixnumMax are therefore variables fixed at package init.
// Anything that must hold a fixnum's value in a narrower type (the compiler's
// folded int32 operands) or multiply two fixnums (the product of two 35-bit
// values overflows int64) has to check the range itself; see fixnumMulFits.
const (
	fixnumStaticBits = 26 // the fallback span: 2^26 bytes = 64 MiB of BSS
	// fixnumSpanTail is readable slack past the last fixnum address, so a
	// stray scmNumber-sized read at the top of the range stays inside mapped
	// memory as it always did with the static array (which BSS follows).
	fixnumSpanTail = 1 << 16
)

var addrForFixnum [1<<fixnumStaticBits + fixnumSpanTail]byte

var (
	// fixnumBaseAddr is the address representing the integer fixnumMin.
	fixnumBaseAddr, fixnumCount = initFixnumSpace()
	fixnumEndAddr               = unsafe.Add(fixnumBaseAddr, fixnumCount)
	fixnumMin                   = -(fixnumCount / 2) // smallest fixnum
	fixnumMax                   = fixnumCount / 2    // one past the largest fixnum
	fixnumMinFloat              = float64(fixnumMin)
	fixnumMaxFloat              = float64(fixnumMax)
)

// initFixnumSpace reserves the widest fixnum span it can, trying 2^36 then
// 2^34 bytes, and falls back to the static array.
func initFixnumSpace() (unsafe.Pointer, int) {
	if unsafe.Sizeof(uintptr(0)) == 8 && os.Getenv("SHEN_GO_NARROW_FIXNUM") == "" {
		for _, bits := range []uint{36, 34} {
			n := 1 << bits
			if base := reserveFixnumSpace(n + fixnumSpanTail); base != nil {
				return base, n
			}
		}
	}
	return unsafe.Pointer(&addrForFixnum[0]), 1 << fixnumStaticBits
}

// fixnumMulFits reports whether a*b is computed exactly in int: true when both
// magnitudes are below 2^31. Otherwise the product is taken in float64, which
// is Shen's number semantics anyway (the exact int product rounded once equals
// the IEEE product of the two exact operands).
func fixnumMulFits(a, b int) bool {
	const lim = 1 << 31
	return a > -lim && a < lim && b > -lim && b < lim
}

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
	if uint(v-fixnumMin) < uint(fixnumCount) {
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
	// Fast path: an integral value in the fixnum range. The range test also
	// rejects NaN and keeps the int conversion defined; -0 becomes fixnum 0,
	// as it always has.
	if f >= fixnumMinFloat && f < fixnumMaxFloat {
		if i := int(f); float64(i) == f {
			return MakeInteger(i)
		}
		tmp := scmNumber{scmHeadNumber, f}
		return &tmp.scmHead
	}
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
	return o != nil && !isFixnum(o) && *o == scmHeadString
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
	if isFixnum(o) {
		return formatNumber(float64(fixnum(o)))
	}
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
