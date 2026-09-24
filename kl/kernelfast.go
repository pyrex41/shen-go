package kl

// Native fast paths for hot kernel functions.
//
// The portable Shen kernel (sys.kl and friends) implements many frequently
// called functions in KLambda so a new port only has to provide the small
// primitive set. Those implementations are correct, but they are also the
// ones that show up in every profile: `empty?` is a VM `cond` around `=`,
// `integer?` is a recursive doubling loop (`shen.magless`) that does not even
// terminate on infinities, `get`/`put` wrap every bucket access in
// trap-error, `shen.pvar?`/`tuple?`/`vector?` do the same, and `symbol?` /
// `variable?` explode the name and walk it character by character.
//
// A Shen-specific tripwire: this port's `(fail)` returns the interned symbol
// `...`, and so does its compiled kernel (cmd/shen/sys.go), while the kernel
// text (sys.kl, and (define fail -> fail!) in sys.shen) returns `shen.fail!`;
// `...` is how the Shen printer shows the fail value. The natives here fill
// vector slots with `...` and `<-vector` compares against it. A vector built
// by one world and read by the other must still agree on "unassigned", or
// `put` feeds a non-list to change-pointer-value and dies at boot; see
// nativeVectorRefOr and kernelArity, which accept either filler.
//
// Shen/Scheme's approach (src/overrides.shen, src/compiler.shen) is to leave
// the kernel sources alone and rebind those functions to natives after load,
// taking advantage of KLambda's uniform calling convention. The same idea
// applies here: BindSymbolFunc after the module Mains have run, matching
// kernel semantics, including error messages on the paths that raise.
//
// Restoring `not` and `integer?` to their original primitive objects has a
// second payoff unique to this port: the AOT kernel is compiled with
// `HasCanonicalPrimitiveBinding` guards, which currently fail because the
// kernel `defun` overwrote those bindings. Putting the canonical primitive
// back turns those guards into inlined typed paths (~62 `not` sites, ~7
// `integer?` sites) instead of trampolined Calls.
//
// New natives (`empty?`, `boolean?`, …) are registered as canonical
// primitives so subsequently compiled VM code can use OP_GUARDED_PRIM.
//
// Every rebinding is unconditional: it is installed whether or not the kernel
// defined the name (issue #49). kl/equiv.json is a table of replacements, and
// Yggdrasil's `lower` pass relies on that by deleting the KL (defun NAME …) of
// each declared native_override; a native that was only installed once its
// defun had run would leave such a slice with "variable vector not bound". On
// a sparse (eval-free) kernel a native that is called behaves at least as
// well as the unbound symbol it replaces: every native is self-contained
// except the two error paths that call shen.app (nativeFn's "fn: X is
// undefined" and nativeGetRaise's get failures), which fall back to a plain
// message when shen.app is unbound; TestNativeFnUndefinedWithoutShenApp
// pins the former.
// TestInstallHelpersToleratesSparseKernel pins only that nothing panics on an
// empty symbol table; TestInstallKernelFastBindsEveryOverride pins that every
// row is bound afterwards.

var (
	symArity          Obj
	symPropertyVector Obj
	symLambdaTable    Obj
	symShenApp        Obj
	symShenA          Obj
	symShenFailBang   Obj // "shen.fail!" — what the kernel text's (fail) returns; fills KL-built vectors
	symFailDots       Obj // "..." — what this port's native (fail) returns; fills native-built vectors
	symShenTuple      Obj
	symShenPvar       Obj
	symShenS          Obj
	fixnumZero        Obj
	fixnumMinusOne    Obj
	fixnumFortyEight  Obj
	fnCacheTable      Obj              // shen.*lambdatable* value the cache was built against
	fnCache           map[Obj]*scmPair // F -> alist entry (F . lambda-value)
	shenAlpha         [128]bool
	shenMisc          [128]bool
)

func init() {
	symArity = MakeSymbol("arity")
	symPropertyVector = MakeSymbol("*property-vector*")
	symLambdaTable = MakeSymbol("shen.*lambdatable*")
	symShenApp = MakeSymbol("shen.app")
	symShenA = MakeSymbol("shen.a")
	symShenFailBang = MakeSymbol("shen.fail!")
	symFailDots = MakeSymbol("...")
	symShenTuple = MakeSymbol("shen.tuple")
	symShenPvar = MakeSymbol("shen.pvar")
	symShenS = MakeSymbol("shen.s")
	fixnumZero = MakeInteger(0)
	fixnumMinusOne = MakeInteger(-1)
	fixnumFortyEight = MakeInteger(48)
	fnCache = make(map[Obj]*scmPair)
	BindSymbolFunc(MakeSymbol("_kl.value/or"), MakeNative(nativeValueOr, 2))
	BindSymbolFunc(MakeSymbol("_kl.<-vector/or"), MakeNative(nativeVectorRefOr, 3))

	for c := byte('A'); c <= 'Z'; c++ {
		shenAlpha[c] = true
	}
	for c := byte('a'); c <= 'z'; c++ {
		shenAlpha[c] = true
	}
	// shen.misc? in reader.kl — the non-letter characters that count as
	// "alpha" for analyse-symbol? / analyse-variable?.
	for _, c := range []byte{'=', '-', '*', '/', '+', '_', '?', '$', '!', '@', '~', '.', '>', '<', '&', '%', '\'', '#', '`'} {
		shenAlpha[c] = true
		shenMisc[c] = true
	}
}

// kernelArity replicates (trap-error (get F arity (value *property-vector*))
// (lambda _ -1)): any failure along the way — unbound *property-vector*, a
// non-vector value, an unassigned bucket (shen.fail!), a missing entry — is
// what the kernel's trap-error would turn into -1. On success it returns the
// stored attribute value unchanged (normally an integer, but `put` accepts
// anything, so no conversion is applied).
func kernelArity(f Obj) Obj {
	pv := mustSymbol(symPropertyVector)
	if pv.value == nil || *pv.value != scmHeadVector {
		return fixnumMinusOne
	}
	vec := mustVector(pv.value)
	// (limit V) is slot 0 of the underlying absvector.
	if len(vec) == 0 || vec[0] == nil || !IsNumber(vec[0]) {
		return fixnumMinusOne
	}
	h := GetInteger(PrimHash(f, vec[0]))
	if h <= 0 || h >= len(vec) {
		return fixnumMinusOne
	}
	bucket := vec[h]
	// (<-vector V h) raises for the never-written filler: `...` from the
	// native vector, shen.fail! from the kernel's KL vector.
	if bucket == nil || bucket == symFailDots || bucket == symShenFailBang {
		return fixnumMinusOne
	}
	// (assoc (cons F (cons arity ())) bucket): each entry is ((F attr) . val).
	for bucket != nil && *bucket == scmHeadPair {
		entry := car(bucket)
		if *entry == scmHeadPair {
			k := car(entry)
			if *k == scmHeadPair && equal(car(k), f) == True {
				krest := cdr(k)
				if *krest == scmHeadPair && car(krest) == symArity && cdr(krest) == Nil {
					return cdr(entry)
				}
			}
		}
		bucket = cdr(bucket)
	}
	// Not found (or bucket wasn't a proper list): kernel get raises, arity's
	// trap-error maps it to -1.
	return fixnumMinusOne
}

// primArityFast is the native binding for `arity`.
func primArityFast(f Obj) Obj {
	return kernelArity(f)
}

// nativeFn is the native binding for `fn`.
func nativeFn(e *ControlFlow) {
	f := e.Get(1)
	if equal(kernelArity(f), fixnumZero) == True {
		// Kernel: ((= (arity F) 0) (F)) — apply F itself with no arguments.
		e.TailApply(f)
		return
	}
	table := PrimValue(symLambdaTable) // unbound propagates, as in the kernel
	if table != fnCacheTable {
		fnCacheTable = table
		clear(fnCache)
	}
	if entry, ok := fnCache[f]; ok {
		e.Return(entry.cdr)
		return
	}
	// (assoc F (value shen.*lambdatable*)): entries are (F . lambda-value).
	for table != nil && *table == scmHeadPair {
		entry := car(table)
		if *entry == scmHeadPair && equal(f, car(entry)) == True {
			p := mustPair(entry)
			fnCache[f] = p
			e.Return(p.cdr)
			return
		}
		table = cdr(table)
	}
	if table != Nil {
		// assoc on a non-list tail raises (uncaught in fn).
		panic(MakeError("attempt to search a non-list with assoc\n"))
	}
	// (simple-error (cn "fn: " (shen.app F " is undefined\n" shen.a)))
	app := kernelBound("shen.app")
	if app == nil {
		// Sparse kernel (issue #49): fn is installed even when shen.app was
		// lowered away, so fall back to a plain message as nativeGetRaise does.
		panic(MakeError("fn: " + ObjString(f) + " is undefined\n"))
	}
	msg := Call(e, app, f, MakeString(" is undefined\n"), symShenA)
	panic(MakeError("fn: " + mustString(msg)))
}

func kernelBound(name string) Obj {
	return mustSymbol(MakeSymbol(name)).function
}

// restoreCanonicalPrimitive puts the init-registered primitive object back
// on a name the kernel's defun overwrote (or never defined), so the AOT
// HasCanonicalPrimitiveBinding guards see the canonical object again.
func restoreCanonicalPrimitive(name string) {
	primitiveRegistry.mu.RLock()
	canonical, ok := primitiveRegistry.canonical[name]
	primitiveRegistry.mu.RUnlock()
	if ok {
		BindSymbolFunc(MakeSymbol(name), canonical)
	}
}

func canonicalOrMake(name string, arity int, fn interface{}) Obj {
	primitiveRegistry.mu.RLock()
	existing, ok := primitiveRegistry.canonical[name]
	primitiveRegistry.mu.RUnlock()
	if ok {
		return existing
	}
	return MakePrimitive(name, arity, fn)
}

// overridePrimitive rebinds a kernel function to a native primitive whether
// or not the kernel defined the name, so a kernel slice whose KL body was
// dropped (Yggdrasil lower) still gets the binding. Reuses the canonical
// primitive object so InstallKernelFast is idempotent under
// HasCanonicalPrimitiveBinding.
func overridePrimitive(name string, arity int, fn interface{}) {
	BindSymbolFunc(MakeSymbol(name), canonicalOrMake(name, arity, fn))
}

// overrideNative is overridePrimitive for natives that need the ControlFlow
// (symbol?, variable?, thaw, fail, put, get, unput, map). A fresh MakeNative
// object never satisfies HasCanonicalPrimitiveBinding, which is deliberate
// for symbol?/variable?: the canonical PrimIsSymbol/PrimIsVariable are weaker
// type-tag checks than the kernel's analyse-symbol?/analyse-variable?.
func overrideNative(name string, arity int, fn func(*ControlFlow)) {
	BindSymbolFunc(MakeSymbol(name), MakeNative(fn, arity))
}

// ---- predicates ----------------------------------------------------------

func primEmptyp(x Obj) Obj {
	if x == Nil {
		return True
	}
	return False
}

func primBooleanp(x Obj) Obj {
	if x == True || x == False {
		return True
	}
	return False
}

func primVectorp(x Obj) Obj {
	if PrimIsVector(x) != True {
		return False
	}
	slots := mustVector(x)
	if len(slots) == 0 || slots[0] == nil || !IsNumber(slots[0]) {
		return False
	}
	if GetNumber(slots[0]) >= 0 {
		return True
	}
	return False
}

func primTuplep(x Obj) Obj {
	if PrimIsVector(x) != True {
		return False
	}
	slots := mustVector(x)
	if len(slots) == 0 || slots[0] == nil {
		return False
	}
	if slots[0] == symShenTuple {
		return True
	}
	return False
}

func primPvarp(x Obj) Obj {
	if PrimIsVector(x) != True {
		return False
	}
	slots := mustVector(x)
	if len(slots) == 0 || slots[0] == nil {
		return False
	}
	if slots[0] == symShenPvar {
		return True
	}
	return False
}

func shenAlphaRune(r rune) bool {
	if r >= 0 && r < 128 {
		return shenAlpha[r]
	}
	return false
}

func shenDigitRune(r rune) bool {
	return r >= '0' && r <= '9'
}

func analyseSymbolName(s string) bool {
	if s == "" {
		return false
	}
	first := true
	for _, r := range s {
		if first {
			if !shenAlphaRune(r) {
				return false
			}
			first = false
			continue
		}
		if !shenAlphaRune(r) && !shenDigitRune(r) {
			return false
		}
	}
	return !first
}

func analyseVariableName(s string) bool {
	if s == "" {
		return false
	}
	first := true
	for _, r := range s {
		if first {
			if r < 'A' || r > 'Z' {
				return false
			}
			first = false
			continue
		}
		if !shenAlphaRune(r) && !shenDigitRune(r) {
			return false
		}
	}
	return !first
}

func primAnalysSymbolp(s Obj) Obj {
	if !IsString(s) || !analyseSymbolName(mustString(s)) {
		// Kernel raises on a non-+string? argument. A non-string is not a
		// +string, so match that with the same implementation-error text.
		if !IsString(s) || mustString(s) == "" {
			panic(MakeError("implementation error in shen.analyse-symbol?"))
		}
		return False
	}
	return True
}

func primSymbolp(x Obj) Obj {
	switch {
	case x == True || x == False, IsNumber(x), IsString(x), x == Nil:
		return False
	}
	if ok, _ := isPair(x); ok {
		return False
	}
	if PrimIsVector(x) == True {
		return False
	}
	if x == MakeSymbol("{") || x == MakeSymbol("}") || x == MakeSymbol(":") || x == MakeSymbol(";") || x == MakeSymbol(",") {
		return True
	}
	// Kernel then does (trap-error (analyse-symbol? (str V)) false). For a
	// real symbol `str` is the name; anything else stringifies to a form
	// analyse-symbol? rejects (#procedure, #vector, …), so a type-tag check
	// is equivalent and does not need trap-error.
	if IsSymbol(x) && analyseSymbolName(GetSymbol(x)) {
		return True
	}
	return False
}

func numberInRange(n Obj, lo, hi float64) Obj {
	f := mustNumber(n)
	if f >= lo && f <= hi {
		return True
	}
	return False
}

func primDigitp(n Obj) Obj     { return numberInRange(n, 48, 57) }
func primLowercasep(n Obj) Obj { return numberInRange(n, 97, 122) }
func primUppercasep(n Obj) Obj { return numberInRange(n, 65, 90) }

func primMiscp(n Obj) Obj {
	f := mustNumber(n)
	if f >= 0 && f < 128 && shenMisc[byte(f)] && float64(byte(f)) == f {
		return True
	}
	return False
}

func primAlphap(n Obj) Obj {
	if primLowercasep(n) == True || primUppercasep(n) == True || primMiscp(n) == True {
		return True
	}
	return False
}

func primPlusStringp(s Obj) Obj {
	if !IsString(s) || mustString(s) == "" {
		return False
	}
	return True
}

func primAlphanumsp(s Obj) Obj {
	if !IsString(s) {
		panic(MakeError("implementation error in shen.alphanums?"))
	}
	str := mustString(s)
	if str == "" {
		return True
	}
	for _, r := range str {
		if !shenAlphaRune(r) && !shenDigitRune(r) {
			return False
		}
	}
	return True
}

func primAnalyseVariablep(s Obj) Obj {
	if !IsString(s) || mustString(s) == "" {
		panic(MakeError("implementation error in shen.analyse-variable?"))
	}
	if analyseVariableName(mustString(s)) {
		return True
	}
	return False
}

func primHdseq(l, x Obj) Obj {
	ok, p := isPair(l)
	if !ok {
		return False
	}
	if equal(p.car, x) == True {
		return True
	}
	return False
}

func primVariablep(x Obj) Obj {
	if x == True || x == False || IsNumber(x) || IsString(x) {
		return False
	}
	if IsSymbol(x) && analyseVariableName(GetSymbol(x)) {
		return True
	}
	return False
}

// ---- constructors / accessors --------------------------------------------

func primAtp(x, y Obj) Obj {
	v := MakeVector(3)
	slots := mustVector(v)
	slots[0] = symShenTuple
	slots[1] = x
	slots[2] = y
	return v
}

func primShenVector(n Obj) Obj {
	size := mustInteger(n)
	if size < 0 {
		panic(MakeError("absvector wrong argument"))
	}
	if size+1 > maxAbsvectorSize {
		panic(MakeError("absvector wrong argument"))
	}
	v := MakeVector(size + 1)
	slots := mustVector(v)
	// Kernel stores the original N at slot 0 (the limit).
	slots[0] = n
	// (fail) returns the interned symbol "...", not "shen.fail!".
	for i := 1; i <= size; i++ {
		slots[i] = symFailDots
	}
	return v
}

func primShenVectorRef(v, n Obj) Obj {
	off := mustInteger(n)
	if off == 0 {
		panic(MakeError("cannot access 0th element of a vector\n"))
	}
	ret := PrimVectorGet(v, n)
	// Kernel: (if (= W (fail)) (simple-error "vector element not found\n") W)
	if ret == symFailDots {
		panic(MakeError("vector element not found\n"))
	}
	return ret
}

func primShenVectorSet(v, n, x Obj) Obj {
	off := mustInteger(n)
	if off == 0 {
		panic(MakeError("cannot access 0th element of a vector\n"))
	}
	return PrimVectorSet(v, n, x)
}

func primLimit(v Obj) Obj {
	return PrimVectorGet(v, fixnumZero)
}

func nativeFail(e *ControlFlow) {
	e.Return(symFailDots)
}

func primFst(v Obj) Obj {
	return PrimVectorGet(v, MakeInteger(1))
}

func primSnd(v Obj) Obj {
	return PrimVectorGet(v, MakeInteger(2))
}

func primHdstr(s Obj) Obj {
	return PrimPos(s, fixnumZero)
}

func primByteToDigit(n Obj) Obj {
	return PrimNumberSubtract(n, fixnumFortyEight)
}

func nativeThaw(e *ControlFlow) {
	e.TailApply(e.Get(1))
}

func nativeValueOr(e *ControlFlow) {
	sym := e.Get(1)
	if IsSymbol(sym) {
		if v := mustSymbol(sym).value; v != nil {
			e.Return(v)
			return
		}
	}
	e.TailApply(e.Get(2))
}

func nativeVectorRefOr(e *ControlFlow) {
	v, n, thunk := e.Get(1), e.Get(2), e.Get(3)
	if PrimIsVector(v) != True || !IsNumber(n) {
		e.TailApply(thunk)
		return
	}
	f := GetNumber(n)
	if !isPreciseInteger(f) || !fitsInt(f) {
		e.TailApply(thunk)
		return
	}
	off := int(f)
	slots := mustVector(v)
	if off == 0 || off < 0 || off >= len(slots) {
		e.TailApply(thunk)
		return
	}
	ret := slots[off]
	if ret == nil {
		ret = undefined
	}
	// An unassigned slot holds whatever (fail) returned when the vector was
	// built: `...` from the native vector, shen.fail! from the kernel's own
	// KL vector (sys.kl's fail). Compiled KL runs in both worlds -- cmd/kl and
	// the equivalence harness boot the KL kernel without the natives -- so
	// both fillers mean absent here, as they do in kernelArity. Before this,
	// a KL-built property vector handed shen.fail! back as a present value and
	// the kernel's put died in shen.change-pointer-value (issue #46).
	if ret == symFailDots || ret == symShenFailBang {
		e.TailApply(thunk)
		return
	}
	e.Return(ret)
}

// ---- lists ---------------------------------------------------------------

func primLength(x Obj) Obj {
	n := 0
	for x != Nil {
		ok, p := isPair(x)
		if !ok {
			// Kernel length-h does (tl V) which is mustPair.
			panic(MakeError("mustPair"))
		}
		x = p.cdr
		n++
	}
	return MakeInteger(n)
}

func primReverse(x Obj) Obj {
	acc := Nil
	for x != Nil {
		ok, p := isPair(x)
		if !ok {
			panic(MakeError("attempt to reverse a non-list\n"))
		}
		acc = cons(p.car, acc)
		x = p.cdr
	}
	return acc
}

func primAppend(x, y Obj) Obj {
	if x == Nil {
		return y
	}
	var elems []Obj
	for x != Nil {
		ok, p := isPair(x)
		if !ok {
			panic(MakeError("attempt to append a non-list"))
		}
		elems = append(elems, p.car)
		x = p.cdr
	}
	acc := y
	for i := len(elems) - 1; i >= 0; i-- {
		acc = cons(elems[i], acc)
	}
	return acc
}

func primElementp(x, l Obj) Obj {
	for l != Nil {
		ok, p := isPair(l)
		if !ok {
			panic(MakeError("attempt to find an element in a non-list\n"))
		}
		if equal(x, p.car) == True {
			return True
		}
		l = p.cdr
	}
	return False
}

func primAssoc(x, l Obj) Obj {
	for l != Nil {
		ok, p := isPair(l)
		if !ok {
			panic(MakeError("attempt to search a non-list with assoc\n"))
		}
		entry := p.car
		if ok, ep := isPair(entry); ok && equal(x, ep.car) == True {
			return entry
		}
		l = p.cdr
	}
	return Nil
}

// ---- property-vector get/put ---------------------------------------------

func changePointerValue(key, attr, val, bucket Obj) Obj {
	if bucket == Nil {
		return cons(cons(cons(key, cons(attr, Nil)), val), Nil)
	}
	ok, p := isPair(bucket)
	if !ok {
		panic(MakeError("implementation error in shen.change-pointer-value"))
	}
	entry := p.car
	if eok, ep := isPair(entry); eok {
		if kok, kp := isPair(ep.car); kok {
			krest := kp.cdr
			if rok, rp := isPair(krest); rok && rp.cdr == Nil {
				if equal(kp.car, key) == True && equal(rp.car, attr) == True {
					return cons(cons(ep.car, val), p.cdr)
				}
			}
		}
	}
	return cons(entry, changePointerValue(key, attr, val, p.cdr))
}

func shenVectorBucket(vec, key Obj) (slots []Obj, h int, bucket Obj, missing bool) {
	slots = mustVector(vec)
	if len(slots) == 0 {
		return slots, 0, Nil, true
	}
	h = GetInteger(PrimHash(key, slots[0]))
	if h <= 0 || h >= len(slots) {
		return slots, h, Nil, true
	}
	bucket = slots[h]
	if bucket == nil {
		bucket = undefined
	}
	// Either fail filler means the bucket was never written; see the file
	// comment and nativeVectorRefOr.
	if bucket == symFailDots || bucket == symShenFailBang {
		return slots, h, Nil, true
	}
	return slots, h, bucket, false
}

func nativePut(e *ControlFlow) {
	key, attr, val, vec := e.Get(1), e.Get(2), e.Get(3), e.Get(4)
	_, h, bucket, missing := shenVectorBucket(vec, key)
	if missing {
		bucket = Nil
	}
	newBucket := changePointerValue(key, attr, val, bucket)
	// vector-> (not a raw slot write) so out-of-range is a catchable Shen error.
	primShenVectorSet(vec, MakeInteger(h), newBucket)
	e.Return(val)
}

func removePointer(key, attr, bucket Obj) Obj {
	if bucket == Nil {
		return Nil
	}
	ok, p := isPair(bucket)
	if !ok {
		panic(MakeError("implementation error in shen.remove-pointer"))
	}
	entry := p.car
	if eok, ep := isPair(entry); eok {
		if kok, kp := isPair(ep.car); kok {
			krest := kp.cdr
			if rok, rp := isPair(krest); rok && rp.cdr == Nil {
				if equal(kp.car, key) == True && equal(rp.car, attr) == True {
					return p.cdr
				}
			}
		}
	}
	return cons(entry, removePointer(key, attr, p.cdr))
}

func nativeUnput(e *ControlFlow) {
	key, attr, vec := e.Get(1), e.Get(2), e.Get(3)
	_, h, bucket, missing := shenVectorBucket(vec, key)
	if missing {
		e.Return(key)
		return
	}
	primShenVectorSet(vec, MakeInteger(h), removePointer(key, attr, bucket))
	e.Return(key)
}

func nativeGet(e *ControlFlow) {
	key, attr, vec := e.Get(1), e.Get(2), e.Get(3)
	_, _, bucket, missing := shenVectorBucket(vec, key)
	if missing {
		nativeGetRaise(e, key, attr, true)
		return
	}
	entry := primAssoc(cons(key, cons(attr, Nil)), bucket)
	if entry == Nil {
		nativeGetRaise(e, key, attr, false)
		return
	}
	ok, p := isPair(entry)
	if !ok {
		nativeGetRaise(e, key, attr, false)
		return
	}
	e.Return(p.cdr)
}

func primHeadShen(x Obj) Obj {
	ok, p := isPair(x)
	if !ok {
		panic(MakeError("head expects a non-empty list\n"))
	}
	return p.car
}

func primTailShen(x Obj) Obj {
	ok, p := isPair(x)
	if !ok {
		panic(MakeError("tail expects a non-empty list\n"))
	}
	return p.cdr
}

func primNth(n, l Obj) Obj {
	if !IsNumber(n) {
		panic(MakeError("nth applied to " + ObjString(n) + ", " + ObjString(l) + "\n"))
	}
	k := GetNumber(n)
	cur := l
	for {
		ok, p := isPair(cur)
		if !ok {
			panic(MakeError("nth applied to " + ObjString(n) + ", " + ObjString(l) + "\n"))
		}
		if k == 1 {
			return p.car
		}
		k--
		cur = p.cdr
	}
}

func primBoundp(x Obj) Obj {
	if primSymbolp(x) != True {
		return False
	}
	v := mustSymbol(x).value
	if v == nil || v == MakeSymbol("shen.this-symbol-is-unbound") {
		return False
	}
	return True
}

func primConcat(x, y Obj) Obj {
	return PrimIntern(PrimStringConcat(PrimStr(x), PrimStr(y)))
}

func primAbs(x Obj) Obj {
	f := mustNumber(x)
	if f > 0 {
		return x
	}
	return MakeNumber(-f)
}

func primPosint(x Obj) Obj {
	if PrimIsInteger(x) == True && GetNumber(x) >= 0 {
		return True
	}
	return False
}

func primSum(l Obj) Obj {
	if l == Nil {
		return MakeInteger(0)
	}
	total := 0.0
	for l != Nil {
		ok, p := isPair(l)
		if !ok {
			panic(MakeError("attempt to sum a non-list\n"))
		}
		total += mustNumber(p.car)
		l = p.cdr
	}
	return MakeNumber(total)
}

func primAdjoin(x, l Obj) Obj {
	if primElementp(x, l) == True {
		return l
	}
	return cons(x, l)
}

func primRemove(x, l Obj) Obj {
	acc := Nil
	for l != Nil {
		ok, p := isPair(l)
		if !ok {
			panic(MakeError("implementation error in shen.remove-h"))
		}
		if equal(x, p.car) != True {
			acc = cons(p.car, acc)
		}
		l = p.cdr
	}
	return primReverse(acc)
}

func primStringToSymbol(s Obj) Obj {
	w := PrimIntern(s)
	if primSymbolp(w) == True {
		return w
	}
	panic(MakeError("cannot intern " + mustString(s) + " to a symbol"))
}

func primStringToBytes(s Obj) Obj {
	str := mustString(s)
	if str == "" {
		return Nil
	}
	runes := []rune(str)
	acc := Nil
	for i := len(runes) - 1; i >= 0; i-- {
		acc = cons(MakeInteger(int(runes[i])), acc)
	}
	return acc
}

func primProtect(x Obj) Obj { return x }

func nativeMap(e *ControlFlow) {
	f, l := e.Get(1), e.Get(2)
	var out []Obj
	for l != Nil {
		ok, p := isPair(l)
		if !ok {
			panic(MakeError("partial function shen.map-h"))
		}
		out = append(out, Call(e, f, p.car))
		l = p.cdr
	}
	acc := Nil
	for i := len(out) - 1; i >= 0; i-- {
		acc = cons(out[i], acc)
	}
	e.Return(acc)
}

func primUnion(a, b Obj) Obj {
	var extra []Obj
	for a != Nil {
		ok, p := isPair(a)
		if !ok {
			panic(MakeError("attempt to find the union with a non-list\n"))
		}
		if primElementp(p.car, b) != True {
			extra = append(extra, p.car)
		}
		a = p.cdr
	}
	acc := b
	for i := len(extra) - 1; i >= 0; i-- {
		acc = cons(extra[i], acc)
	}
	return acc
}

func primIntersection(a, b Obj) Obj {
	acc := Nil
	for a != Nil {
		ok, p := isPair(a)
		if !ok {
			panic(MakeError("attempt to find the intersection with a non-list\n"))
		}
		if primElementp(p.car, b) == True {
			acc = cons(p.car, acc)
		}
		a = p.cdr
	}
	return primReverse(acc)
}

func primDifference(a, b Obj) Obj {
	acc := Nil
	for a != Nil {
		ok, p := isPair(a)
		if !ok {
			panic(MakeError("attempt to find the difference with a non-list\n"))
		}
		if primElementp(p.car, b) != True {
			acc = cons(p.car, acc)
		}
		a = p.cdr
	}
	return primReverse(acc)
}

func nativeGetRaise(e *ControlFlow, key, attr Obj, noAttrs bool) {
	app := kernelBound("shen.app")
	if app == nil {
		if noAttrs {
			panic(MakeError("has no attributes"))
		}
		panic(MakeError("attribute not found"))
	}
	if noAttrs {
		inner := Call(e, app, attr, MakeString("\n"), symShenS)
		msg := Call(e, app, key, PrimStringConcat(MakeString(" has no attributes: "), inner), symShenA)
		panic(MakeError(mustString(msg)))
	}
	inner := Call(e, app, key, MakeString("\n"), symShenS)
	mid := Call(e, app, attr, PrimStringConcat(MakeString(" not found for "), inner), symShenS)
	panic(MakeError("attribute " + mustString(mid)))
}

// InstallKernelFast rebinds hot kernel functions to the natives above. Every
// rebinding happens on any kernel, including one whose defuns were lowered
// away (issue #49) and an empty symbol table. It is called after each kernel
// module (kl.BootKernel) or chunk (cmd/yggdrasil-build), because a later
// module's own (defun …) of a rebound name overwrites the native until the
// next call restores it; the natives read *property-vector* and
// shen.*lambdatable* at call time, not at install time.
//
// Every rebinding here is audited against the kernel's KL definition by
// TestEquivTable (equiv_test.go), which parses this function's body and writes
// the verdicts to kl/equiv.json; adding or removing a rebinding fails that test
// until the table is regenerated.
func InstallKernelFast() {
	BindSymbolFunc(symArity, canonicalOrMake("arity", 1, primArityFast))
	BindSymbolFunc(MakeSymbol("fn"), MakeNative(nativeFn, 1))

	// Re-enable AOT HasCanonicalPrimitiveBinding fast paths. PrimIsInteger
	// also terminates on +-Inf/NaN, which the kernel magless loop does not.
	restoreCanonicalPrimitive("not")
	restoreCanonicalPrimitive("integer?")

	overridePrimitive("empty?", 1, primEmptyp)
	overridePrimitive("boolean?", 1, primBooleanp)
	overridePrimitive("vector?", 1, primVectorp)
	overridePrimitive("tuple?", 1, primTuplep)
	overridePrimitive("shen.pvar?", 1, primPvarp)
	// symbol? and variable? already have canonical primitives (PrimIsSymbol /
	// PrimIsVariable) registered at init. Those are type-tag checks and do
	// not match the kernel (analyse-symbol?/analyse-variable?). Bind a new
	// native rather than restoreCanonical / overridePrimitive, which would
	// reinstall the weaker primitive and trip HasCanonicalPrimitiveBinding.
	overrideNative("symbol?", 1, func(e *ControlFlow) { e.Return(primSymbolp(e.Get(1))) })
	overrideNative("variable?", 1, func(e *ControlFlow) { e.Return(primVariablep(e.Get(1))) })
	overridePrimitive("shen.analyse-symbol?", 1, primAnalysSymbolp)
	overridePrimitive("shen.analyse-variable?", 1, primAnalyseVariablep)
	overridePrimitive("shen.digit?", 1, primDigitp)
	overridePrimitive("shen.lowercase?", 1, primLowercasep)
	overridePrimitive("shen.uppercase?", 1, primUppercasep)
	overridePrimitive("shen.misc?", 1, primMiscp)
	overridePrimitive("shen.alpha?", 1, primAlphap)
	overridePrimitive("shen.alphanums?", 1, primAlphanumsp)
	overridePrimitive("shen.+string?", 1, primPlusStringp)
	overridePrimitive("shen.hds=?", 2, primHdseq)

	overridePrimitive("@p", 2, primAtp)
	overridePrimitive("vector", 1, primShenVector)
	overridePrimitive("<-vector", 2, primShenVectorRef)
	overridePrimitive("vector->", 3, primShenVectorSet)
	overridePrimitive("limit", 1, primLimit)
	overridePrimitive("fst", 1, primFst)
	overridePrimitive("snd", 1, primSnd)
	overridePrimitive("hdstr", 1, primHdstr)
	overridePrimitive("shen.byte->digit", 1, primByteToDigit)

	overrideNative("thaw", 1, nativeThaw)
	overrideNative("fail", 0, nativeFail)

	overridePrimitive("length", 1, primLength)
	overridePrimitive("reverse", 1, primReverse)
	overridePrimitive("append", 2, primAppend)
	overridePrimitive("element?", 2, primElementp)
	overridePrimitive("assoc", 2, primAssoc)

	overrideNative("put", 4, nativePut)
	overrideNative("get", 3, nativeGet)
	overrideNative("unput", 3, nativeUnput)

	overridePrimitive("head", 1, primHeadShen)
	overridePrimitive("tail", 1, primTailShen)
	overridePrimitive("nth", 2, primNth)
	overridePrimitive("bound?", 1, primBoundp)
	overridePrimitive("concat", 2, primConcat)
	overridePrimitive("==", 2, PrimEqual)
	overridePrimitive("shen.abs", 1, primAbs)
	overridePrimitive("shen.posint?", 1, primPosint)
	overridePrimitive("sum", 1, primSum)
	overridePrimitive("adjoin", 2, primAdjoin)
	overridePrimitive("remove", 2, primRemove)
	overridePrimitive("string->symbol", 1, primStringToSymbol)
	overridePrimitive("shen.string->bytes", 1, primStringToBytes)
	overridePrimitive("protect", 1, primProtect)
	overridePrimitive("union", 2, primUnion)
	overridePrimitive("intersection", 2, primIntersection)
	overridePrimitive("difference", 2, primDifference)
	overrideNative("map", 2, nativeMap)
}
