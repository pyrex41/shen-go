package kl

import (
	"os"
	"strings"
	"sync/atomic"
)

// KindSetForAnnotation translates the conventional Shen annotation names to
// their conservative runtime shape. Unknown names deliberately produce
// KindUnknown instead of rejecting source or enforcing a type at runtime.
func KindSetForAnnotation(name string) KindSet {
	switch strings.ToLower(strings.TrimSpace(name)) {
	case "number", "numeric":
		return Kinds(KindNumber)
	case "fixnum", "integer":
		return Kinds(KindFixnum)
	case "boolean", "bool":
		return Kinds(KindBoolean)
	case "string":
		return Kinds(KindString)
	case "symbol":
		return Kinds(KindSymbol)
	case "pair", "cons", "list":
		return Kinds(KindPair)
	case "nil":
		return Kinds(KindNil)
	case "vector", "absvector":
		return Kinds(KindVector)
	case "callable", "function", "procedure":
		return Kinds(KindCallable)
	case "stream":
		return Kinds(KindStream)
	case "error":
		return Kinds(KindError)
	case "raw":
		return Kinds(KindRaw)
	default:
		return Kinds(KindUnknown)
	}
}

// ValueKind describes the runtime shape a typed-IR value is known to have.
// Kinds are deliberately conservative: unknown values must remain on the Obj
// path, while a KindSet can represent the result of a control-flow join.
type ValueKind uint16

const (
	KindUnknown ValueKind = 1 << iota
	KindNumber
	KindFixnum
	KindBoolean
	KindString
	KindSymbol
	KindPair
	KindNil
	KindVector
	KindCallable
	KindStream
	KindError
	KindRaw
)

// KindSet is a compact union of possible ValueKinds.
type KindSet uint32

func Kinds(kinds ...ValueKind) KindSet {
	var set KindSet
	for _, k := range kinds {
		set |= KindSet(k)
	}
	return set
}

// KindSetOf is a descriptive alias useful at compiler call sites.
func KindSetOf(kinds ...ValueKind) KindSet { return Kinds(kinds...) }

func (s KindSet) Contains(k ValueKind) bool { return s&KindSet(k) != 0 }
func (s KindSet) Accepts(k ValueKind) bool {
	return s.Contains(k) || (k == KindFixnum && s.Contains(KindNumber))
}
func (s KindSet) Union(other KindSet) KindSet     { return s | other }
func (s KindSet) Intersect(other KindSet) KindSet { return s & other }
func (s KindSet) Empty() bool                     { return s == 0 }
func (s KindSet) IsSingle() bool                  { return s != 0 && s&(s-1) == 0 }

// PrimitiveID is the stable identifier used by compiler metadata. It is the
// Shen primitive name, not a pointer address, so generated code can use it too.
type PrimitiveID string

type PrimitiveEffects uint16

const (
	EffectPure PrimitiveEffects = 1 << iota
	EffectMayRaise
	EffectReadsGlobal
	EffectWritesGlobal
	EffectWritesHeap
)

// PrimitiveSpec describes the type/effect contract known to the optimizer.
// Empty argument/result sets mean that no specialization is available.
type PrimitiveSpec struct {
	ID      PrimitiveID
	Name    string
	Arity   int
	Args    []KindSet
	Result  KindSet
	Effects PrimitiveEffects
	Scalar  bool
}

var primitiveSpecs = map[string]PrimitiveSpec{
	"+":               {ID: "+", Name: "+", Arity: 2, Args: []KindSet{Kinds(KindNumber), Kinds(KindNumber)}, Result: Kinds(KindNumber), Effects: EffectPure | EffectMayRaise, Scalar: true},
	"-":               {ID: "-", Name: "-", Arity: 2, Args: []KindSet{Kinds(KindNumber), Kinds(KindNumber)}, Result: Kinds(KindNumber), Effects: EffectPure | EffectMayRaise, Scalar: true},
	"*":               {ID: "*", Name: "*", Arity: 2, Args: []KindSet{Kinds(KindNumber), Kinds(KindNumber)}, Result: Kinds(KindNumber), Effects: EffectPure | EffectMayRaise, Scalar: true},
	"/":               {ID: "/", Name: "/", Arity: 2, Args: []KindSet{Kinds(KindNumber), Kinds(KindNumber)}, Result: Kinds(KindNumber), Effects: EffectPure | EffectMayRaise, Scalar: true},
	"=":               {ID: "=", Name: "=", Arity: 2, Args: []KindSet{Kinds(KindUnknown), Kinds(KindUnknown)}, Result: Kinds(KindBoolean), Effects: EffectPure, Scalar: true},
	"<":               {ID: "<", Name: "<", Arity: 2, Args: []KindSet{Kinds(KindNumber), Kinds(KindNumber)}, Result: Kinds(KindBoolean), Effects: EffectPure | EffectMayRaise, Scalar: true},
	"<=":              {ID: "<=", Name: "<=", Arity: 2, Args: []KindSet{Kinds(KindNumber), Kinds(KindNumber)}, Result: Kinds(KindBoolean), Effects: EffectPure | EffectMayRaise, Scalar: true},
	">":               {ID: ">", Name: ">", Arity: 2, Args: []KindSet{Kinds(KindNumber), Kinds(KindNumber)}, Result: Kinds(KindBoolean), Effects: EffectPure | EffectMayRaise, Scalar: true},
	">=":              {ID: ">=", Name: ">=", Arity: 2, Args: []KindSet{Kinds(KindNumber), Kinds(KindNumber)}, Result: Kinds(KindBoolean), Effects: EffectPure | EffectMayRaise, Scalar: true},
	"number?":         {ID: "number?", Name: "number?", Arity: 1, Args: []KindSet{Kinds(KindUnknown)}, Result: Kinds(KindBoolean), Effects: EffectPure, Scalar: true},
	"integer?":        {ID: "integer?", Name: "integer?", Arity: 1, Args: []KindSet{Kinds(KindUnknown)}, Result: Kinds(KindBoolean), Effects: EffectPure, Scalar: true},
	"string?":         {ID: "string?", Name: "string?", Arity: 1, Args: []KindSet{Kinds(KindUnknown)}, Result: Kinds(KindBoolean), Effects: EffectPure, Scalar: true},
	"symbol?":         {ID: "symbol?", Name: "symbol?", Arity: 1, Args: []KindSet{Kinds(KindUnknown)}, Result: Kinds(KindBoolean), Effects: EffectPure, Scalar: true},
	"cons?":           {ID: "cons?", Name: "cons?", Arity: 1, Args: []KindSet{Kinds(KindUnknown)}, Result: Kinds(KindBoolean), Effects: EffectPure, Scalar: true},
	"absvector?":      {ID: "absvector?", Name: "absvector?", Arity: 1, Args: []KindSet{Kinds(KindUnknown)}, Result: Kinds(KindBoolean), Effects: EffectPure, Scalar: true},
	"variable?":       {ID: "variable?", Name: "variable?", Arity: 1, Args: []KindSet{Kinds(KindUnknown)}, Result: Kinds(KindBoolean), Effects: EffectPure, Scalar: true},
	"not":             {ID: "not", Name: "not", Arity: 1, Args: []KindSet{Kinds(KindBoolean)}, Result: Kinds(KindBoolean), Effects: EffectPure | EffectMayRaise, Scalar: true},
	"if":              {ID: "if", Name: "if", Arity: 3, Args: []KindSet{Kinds(KindBoolean), Kinds(KindUnknown), Kinds(KindUnknown)}, Result: Kinds(KindUnknown), Effects: EffectPure | EffectMayRaise},
	"cn":              {ID: "cn", Name: "cn", Arity: 2, Args: []KindSet{Kinds(KindString), Kinds(KindString)}, Result: Kinds(KindString), Effects: EffectPure | EffectMayRaise, Scalar: true},
	"tlstr":           {ID: "tlstr", Name: "tlstr", Arity: 1, Args: []KindSet{Kinds(KindString)}, Result: Kinds(KindString), Effects: EffectPure | EffectMayRaise, Scalar: true},
	"pos":             {ID: "pos", Name: "pos", Arity: 2, Args: []KindSet{Kinds(KindString), Kinds(KindNumber)}, Result: Kinds(KindString), Effects: EffectPure | EffectMayRaise, Scalar: true},
	"string->n":       {ID: "string->n", Name: "string->n", Arity: 1, Args: []KindSet{Kinds(KindString)}, Result: Kinds(KindNumber), Effects: EffectPure | EffectMayRaise, Scalar: true},
	"n->string":       {ID: "n->string", Name: "n->string", Arity: 1, Args: []KindSet{Kinds(KindNumber)}, Result: Kinds(KindString), Effects: EffectPure | EffectMayRaise, Scalar: true},
	"str":             {ID: "str", Name: "str", Arity: 1, Args: []KindSet{Kinds(KindUnknown)}, Result: Kinds(KindString), Effects: EffectPure | EffectMayRaise, Scalar: true},
	"cons":            {ID: "cons", Name: "cons", Arity: 2, Args: []KindSet{Kinds(KindUnknown), Kinds(KindUnknown)}, Result: Kinds(KindPair), Effects: EffectPure, Scalar: false},
	"hd":              {ID: "hd", Name: "hd", Arity: 1, Args: []KindSet{Kinds(KindPair)}, Result: Kinds(KindUnknown), Effects: EffectPure | EffectMayRaise},
	"tl":              {ID: "tl", Name: "tl", Arity: 1, Args: []KindSet{Kinds(KindPair)}, Result: Kinds(KindUnknown), Effects: EffectPure | EffectMayRaise},
	"intern":          {ID: "intern", Name: "intern", Arity: 1, Args: []KindSet{Kinds(KindString)}, Result: Kinds(KindSymbol, KindBoolean), Effects: EffectPure | EffectMayRaise},
	"absvector":       {ID: "absvector", Name: "absvector", Arity: 1, Args: []KindSet{Kinds(KindNumber)}, Result: Kinds(KindVector), Effects: EffectMayRaise},
	"<-address":       {ID: "<-address", Name: "<-address", Arity: 2, Args: []KindSet{Kinds(KindVector), Kinds(KindNumber)}, Result: Kinds(KindUnknown), Effects: EffectMayRaise},
	"address->":       {ID: "address->", Name: "address->", Arity: 3, Args: []KindSet{Kinds(KindVector), Kinds(KindNumber), Kinds(KindUnknown)}, Result: Kinds(KindVector), Effects: EffectMayRaise | EffectWritesHeap},
	"value":           {ID: "value", Name: "value", Arity: 1, Args: []KindSet{Kinds(KindSymbol)}, Result: Kinds(KindUnknown), Effects: EffectMayRaise | EffectReadsGlobal},
	"set":             {ID: "set", Name: "set", Arity: 2, Args: []KindSet{Kinds(KindSymbol), Kinds(KindUnknown)}, Result: Kinds(KindUnknown), Effects: EffectWritesGlobal},
	"simple-error":    {ID: "simple-error", Name: "simple-error", Arity: 1, Args: []KindSet{Kinds(KindString)}, Result: Kinds(KindError), Effects: EffectMayRaise},
	"error-to-string": {ID: "error-to-string", Name: "error-to-string", Arity: 1, Args: []KindSet{Kinds(KindError)}, Result: Kinds(KindString), Effects: EffectPure | EffectMayRaise, Scalar: true},
}

// LookupPrimitiveSpec returns a copy of the metadata for name.
func LookupPrimitiveSpec(name string) (PrimitiveSpec, bool) {
	spec, ok := primitiveSpecs[name]
	if !ok {
		return PrimitiveSpec{}, false
	}
	spec.Args = append([]KindSet(nil), spec.Args...)
	return spec, true
}

// PrimitiveSpecFor is an alias retained for callers that prefer noun-first
// lookup naming.
func PrimitiveSpecFor(name string) (PrimitiveSpec, bool) { return LookupPrimitiveSpec(name) }

var typedIRMode atomic.Uint32 // 0 unset, 1 enabled, 2 disabled

// HasCanonicalPrimitiveBinding reports whether sym still points at the first
// registered implementation for its name. This is the guard required before
// executing a specialized primitive region.
func HasCanonicalPrimitiveBinding(sym Obj) bool {
	if sym == nil || isFixnum(sym) || *sym != scmHeadSymbol {
		return false
	}
	name := GetSymbol(sym)
	primitiveRegistry.mu.RLock()
	canonical, ok := primitiveRegistry.canonical[name]
	primitiveRegistry.mu.RUnlock()
	return ok && mustSymbol(sym).function == canonical
}

// TypedIRModeEnabled controls specialization. The default is enabled; only an
// explicit SHEN_GO_TYPED_IR=off disables it. Unknown values retain the safe
// default so a typo cannot silently select the slower compatibility mode.
func TypedIRModeEnabled() bool {
	if mode := typedIRMode.Load(); mode != 0 {
		return mode == 1
	}
	mode := uint32(1)
	if strings.EqualFold(strings.TrimSpace(os.Getenv("SHEN_GO_TYPED_IR")), "off") {
		mode = 2
	}
	typedIRMode.CompareAndSwap(0, mode)
	return mode == 1
}

// TypedIREnabled is a concise alias for callers in compiler and VM packages.
func TypedIREnabled() bool { return TypedIRModeEnabled() }

// ResetTypedIRModeForTest clears the cached environment decision. Production
// callers should set SHEN_GO_TYPED_IR before process startup and never reset it.
func ResetTypedIRModeForTest() { typedIRMode.Store(0) }
