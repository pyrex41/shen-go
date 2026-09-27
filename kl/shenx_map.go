package kl

import (
	"math"
	"strconv"
	"unsafe"
)

// shen.x.map: a mutable hash map keyed by Shen values (issue #57).
//
// Keys are compared with the kernel's = (equal), and bucketed by
// shenValueHash, which is consistent with it: two values that are = hash
// alike. Lists, strings, numbers, symbols, booleans and absvectors are
// hashed by content, so a list-of-numbers state or a tuple can be a key
// directly; closures and streams are = only to themselves and hash by
// identity. As with any hash map, a key that is mutated after insertion
// (an absvector written with address->) is not found again.
//
// A library that ships a portable map (for example tla.shen's
// tla.map-new / tla.map-get / tla.map-put over an assoc list or put/get)
// can test for the feature and rebind to these:
//
//	(if (element? shen.x/map (shen.x.features.current)) ... ...)
//
//	(shen.x.map-new)              -> a new, empty map
//	(shen.x.map-put M K V)        -> M, with K bound to V
//	(shen.x.map-get M K Default)  -> the value bound to K, or Default
//	(shen.x.map-has? M K)         -> true when K is bound
//	(shen.x.map-remove M K)       -> M, without K
//	(shen.x.map-count M)          -> the number of bound keys
//	(shen.x.map-keys M)           -> the bound keys, as a list
type scmMap struct {
	scmHead
	buckets map[uint64][]mapEntry
	n       int
}

type mapEntry struct{ key, val Obj }

const scmHeadMap scmHead = 43

func makeMap() Obj {
	m := &scmMap{scmHead: scmHeadMap, buckets: make(map[uint64][]mapEntry)}
	return &m.scmHead
}

func mustMap(o Obj, who string) *scmMap {
	if o == nil || isFixnum(o) || *o != scmHeadMap {
		panic(MakeError(who + ": not a shen.x map: " + ObjString(o)))
	}
	return (*scmMap)(unsafe.Pointer(o))
}

func (m *scmMap) find(k Obj) (uint64, int) {
	h := shenValueHash(k)
	for i, e := range m.buckets[h] {
		if equal(e.key, k) == True {
			return h, i
		}
	}
	return h, -1
}

const (
	fnvOffset = 14695981039346656037
	fnvPrime  = 1099511628211
)

func mix(h, x uint64) uint64 {
	h ^= x
	h *= fnvPrime
	h ^= h >> 29
	return h
}

func mixString(h uint64, s string) uint64 {
	for i := 0; i < len(s); i++ {
		h ^= uint64(s[i])
		h *= fnvPrime
	}
	return mix(h, uint64(len(s)))
}

// shenValueHash hashes o so that equal(a, b) == True implies equal hashes.
// Every kind equal treats as "always equal within the kind" (natives, raw
// objects, errors: equal compares only their heads) hashes to its head.
func shenValueHash(o Obj) uint64 {
	return valueHash(fnvOffset, o)
}

func valueHash(h uint64, o Obj) uint64 {
	for {
		if o == nil {
			return mix(h, 0x9e3779b97f4a7c15)
		}
		if isFixnum(o) {
			return mixNumber(h, float64(fixnum(o)))
		}
		switch *o {
		case scmHeadNumber:
			return mixNumber(h, mustNumber(o))
		case scmHeadString:
			return mixString(mix(h, uint64(scmHeadString)), mustString(o))
		case scmHeadSymbol:
			return mixString(mix(h, uint64(scmHeadSymbol)), mustSymbol(o).str)
		case scmHeadPair:
			// Iterate down the spine so a long list does not recurse.
			h = mix(h, uint64(scmHeadPair))
			p := mustPair(o)
			h = valueHash(h, p.car)
			o = p.cdr
			continue
		case scmHeadVector:
			v := mustVector(o)
			h = mix(mix(h, uint64(scmHeadVector)), uint64(len(v)))
			for _, x := range v {
				h = valueHash(h, x)
			}
			return h
		case scmHeadStream, scmHeadProcedure, scmHeadBytecodeFunc, scmHeadMap:
			// = on these is identity.
			return mix(h, uint64(uintptr(unsafe.Pointer(o))))
		default:
			return mix(h, uint64(*o)+0x51)
		}
	}
}

func mixNumber(h uint64, f float64) uint64 {
	if f == 0 {
		f = 0 // fold -0 into +0: they are =
	}
	return mix(mix(h, uint64(scmHeadNumber)), math.Float64bits(f))
}

func primMapPut(mo, k, v Obj) Obj {
	m := mustMap(mo, "shen.x.map-put")
	h, i := m.find(k)
	if i >= 0 {
		m.buckets[h][i].val = v
		return mo
	}
	m.buckets[h] = append(m.buckets[h], mapEntry{k, v})
	m.n++
	return mo
}

func primMapGet(mo, k, dflt Obj) Obj {
	m := mustMap(mo, "shen.x.map-get")
	h, i := m.find(k)
	if i < 0 {
		return dflt
	}
	return m.buckets[h][i].val
}

func primMapHas(mo, k Obj) Obj {
	m := mustMap(mo, "shen.x.map-has?")
	if _, i := m.find(k); i >= 0 {
		return True
	}
	return False
}

func primMapRemove(mo, k Obj) Obj {
	m := mustMap(mo, "shen.x.map-remove")
	h, i := m.find(k)
	if i < 0 {
		return mo
	}
	b := m.buckets[h]
	b[i] = b[len(b)-1]
	b[len(b)-1] = mapEntry{}
	if len(b) == 1 {
		delete(m.buckets, h)
	} else {
		m.buckets[h] = b[:len(b)-1]
	}
	m.n--
	return mo
}

func primMapCount(mo Obj) Obj {
	return MakeInteger(mustMap(mo, "shen.x.map-count").n)
}

func primMapKeys(mo Obj) Obj {
	m := mustMap(mo, "shen.x.map-keys")
	ret := Nil
	for _, b := range m.buckets {
		for _, e := range b {
			ret = cons(e.key, ret)
		}
	}
	return ret
}

func mapString(o Obj) string {
	return "<shen.x.map " + strconv.Itoa((*scmMap)(unsafe.Pointer(o)).n) + ">"
}

func installShenXMap() {
	bind := func(name string, arity int, fn func(e *ControlFlow) Obj) {
		BindSymbolFunc(MakeSymbol(name), MakeNative(func(e *ControlFlow) { e.Return(fn(e)) }, arity))
	}
	bind("shen.x.map-new", 0, func(e *ControlFlow) Obj { return makeMap() })
	bind("shen.x.map-put", 3, func(e *ControlFlow) Obj { return primMapPut(e.Get(1), e.Get(2), e.Get(3)) })
	bind("shen.x.map-get", 3, func(e *ControlFlow) Obj { return primMapGet(e.Get(1), e.Get(2), e.Get(3)) })
	bind("shen.x.map-has?", 2, func(e *ControlFlow) Obj { return primMapHas(e.Get(1), e.Get(2)) })
	bind("shen.x.map-remove", 2, func(e *ControlFlow) Obj { return primMapRemove(e.Get(1), e.Get(2)) })
	bind("shen.x.map-count", 1, func(e *ControlFlow) Obj { return primMapCount(e.Get(1)) })
	bind("shen.x.map-keys", 1, func(e *ControlFlow) Obj { return primMapKeys(e.Get(1)) })
	registerShenXFeature("shen.x/map")
}
