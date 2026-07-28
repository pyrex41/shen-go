package kl

// Native fast paths for the kernel's `arity` and `fn` (declarations.kl /
// reader.kl). Both are on the hot path of every runtime higher-order function
// reference: Shen compiles `(fn f)` for function values and partial
// applications, and the kernel implements it as
//
//	(defun fn (F)
//	  (cond ((= (arity F) 0) (F))
//	        (true (let W (assoc F (value shen.*lambdatable*))
//	                (if (empty? W) (simple-error ...) (tl W))))))
//
// where `arity` is (trap-error (get F arity (value *property-vector*))
// (lambda _ -1)). Interpreted, one (fn f) costs: two trap-error closure
// allocations, a hash-dict lookup (hash + bucket assoc) for the arity, and a
// linear assoc over shen.*lambdatable* — an alist with one entry per defined
// function (hundreds of entries in a loaded application), scanned per call.
//
// The natives below reproduce those semantics exactly against the same kernel
// data structures (the *property-vector* hash dict and the shen.*lambdatable*
// alist), including error messages and the arity-0 behaviour, but without the
// trap-error machinery, and with a lookaside cache that memoizes lambdatable
// hits keyed on the table's identity: the table is an immutable alist only
// ever replaced wholesale via (set shen.*lambdatable* ...), so a simple
// pointer comparison detects any change and empties the cache.

var (
	symArity           Obj
	symPropertyVector  Obj
	symLambdaTable     Obj
	symShenApp         Obj
	symShenA           Obj
	symShenFailBang    Obj
	fixnumZero         Obj
	fixnumMinusOne     Obj
	fnCacheTable       Obj             // shen.*lambdatable* value the cache was built against
	fnCache            map[Obj]*scmPair // F -> alist entry (F . lambda-value)
)

func init() {
	symArity = MakeSymbol("arity")
	symPropertyVector = MakeSymbol("*property-vector*")
	symLambdaTable = MakeSymbol("shen.*lambdatable*")
	symShenApp = MakeSymbol("shen.app")
	symShenA = MakeSymbol("shen.a")
	symShenFailBang = MakeSymbol("shen.fail!")
	fixnumZero = MakeInteger(0)
	fixnumMinusOne = MakeInteger(-1)
	fnCache = make(map[Obj]*scmPair)
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
	if len(vec) == 0 || vec[0] == nil || *vec[0] != scmHeadNumber {
		return fixnumMinusOne
	}
	h := GetInteger(PrimHash(f, vec[0]))
	if h <= 0 || h >= len(vec) {
		return fixnumMinusOne
	}
	bucket := vec[h]
	// (<-vector V h) raises an error for the never-written filler shen.fail!
	// (and address-> can technically store anything, so stay defensive).
	if bucket == nil || bucket == symShenFailBang {
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
	msg := Call(e, PrimFunc(symShenApp), f, MakeString(" is undefined\n"), symShenA)
	panic(MakeError("fn: " + mustString(msg)))
}

// InstallKernelFast rebinds `arity` and `fn` to the natives above. It must be
// called after the kernel modules have run (they define the interpreted
// versions and build the *property-vector* / shen.*lambdatable* structures the
// natives read).
func InstallKernelFast() {
	BindSymbolFunc(symArity, MakePrimitive("arity", 1, primArityFast))
	BindSymbolFunc(MakeSymbol("fn"), MakeNative(nativeFn, 1))
}
