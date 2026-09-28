package kl

import (
	"encoding/binary"
	"math"
	"os"
	"sync"
	"unsafe"
)

// Native structural hash maps for Shen data (shen.x/map-host).
//
//	(shen.x.map-new Capacity)   -> handle (a number)
//	(shen.x.map-get Handle Key) -> the stored value, or the symbol shen.x.none
//	(shen.x.map-put Handle Key Value) -> Handle
//	(shen.x.map-size Handle)    -> number of keys
//
// Keys are compared exactly as `=` compares them (library.go equal): lists,
// absvectors (so tuples) and strings structurally; symbols, booleans,
// streams and procedures by identity; numbers by value (1 and 1.0 alike).
// The key is flattened once into a canonical byte string and looked up in a
// Go map, so a lookup costs one walk of the key instead of a Shen-level hash
// and bucket scan. Maps live for the whole process (model checking is a
// batch job); there is no delete. Disable with SHEN_X_MAP=off.

type shenxMapTable struct {
	mu   sync.Mutex
	maps []map[string]Obj
}

var shenxMaps shenxMapTable

var shenxNone = MakeSymbol("shen.x.none")

func installShenXMap() {
	if os.Getenv("SHEN_X_MAP") == "off" {
		return
	}
	BindSymbolFunc(MakeSymbol("shen.x.map-new"), MakePrimitive("shen.x.map-new", 1, primShenXMapNew))
	BindSymbolFunc(MakeSymbol("shen.x.map-get"), MakePrimitive("shen.x.map-get", 2, primShenXMapGet))
	BindSymbolFunc(MakeSymbol("shen.x.map-put"), MakePrimitive("shen.x.map-put", 3, primShenXMapPut))
	BindSymbolFunc(MakeSymbol("shen.x.map-size"), MakePrimitive("shen.x.map-size", 1, primShenXMapSize))
	registerShenXFeature("shen.x/map-host")
}

func primShenXMapNew(capacity Obj) Obj {
	n := 0
	if IsNumber(capacity) {
		n = GetInteger(capacity)
	}
	if n < 0 {
		n = 0
	}
	shenxMaps.mu.Lock()
	defer shenxMaps.mu.Unlock()
	shenxMaps.maps = append(shenxMaps.maps, make(map[string]Obj, n))
	return MakeInteger(len(shenxMaps.maps) - 1)
}

func shenxMapOf(handle Obj) map[string]Obj {
	if !IsNumber(handle) {
		panic(MakeError("shen.x.map: handle is not a number"))
	}
	i := GetInteger(handle)
	shenxMaps.mu.Lock()
	defer shenxMaps.mu.Unlock()
	if i < 0 || i >= len(shenxMaps.maps) {
		panic(MakeError("shen.x.map: unknown handle"))
	}
	return shenxMaps.maps[i]
}

func primShenXMapGet(handle, key Obj) Obj {
	m := shenxMapOf(handle)
	if v, ok := m[string(shenxKey(nil, key))]; ok {
		return v
	}
	return shenxNone
}

func primShenXMapPut(handle, key, value Obj) Obj {
	m := shenxMapOf(handle)
	m[string(shenxKey(nil, key))] = value
	return handle
}

func primShenXMapSize(handle Obj) Obj {
	return MakeInteger(len(shenxMapOf(handle)))
}

// shenxKey appends the canonical encoding of o: one tag byte, then a fixed
// or length-prefixed payload, so distinct `=`-classes never collide.
func shenxKey(buf []byte, o Obj) []byte {
	if o == nil {
		return append(buf, 'z')
	}
	if isFixnum(o) {
		buf = append(buf, 'i')
		return binary.AppendVarint(buf, int64(fixnum(o)))
	}
	switch *o {
	case scmHeadNumber:
		// `=` compares numbers by value: (= 3 (+ 1.5 1.5)) and (= 0 -0.0)
		// hold, so an integral float is keyed exactly like the fixnum.
		f := mustNumber(o)
		if f == math.Trunc(f) && f >= -(1<<62) && f < (1<<62) {
			buf = append(buf, 'i')
			return binary.AppendVarint(buf, int64(f))
		}
		buf = append(buf, 'f')
		return binary.LittleEndian.AppendUint64(buf, math.Float64bits(f))
	case scmHeadString:
		s := mustString(o)
		buf = append(buf, 's')
		buf = binary.AppendUvarint(buf, uint64(len(s)))
		return append(buf, s...)
	case scmHeadNull:
		return append(buf, 'n')
	case scmHeadPair:
		// A proper or improper list: 'p' per cell, then the tail.
		for !isFixnum(o) && *o == scmHeadPair {
			buf = append(buf, 'p')
			buf = shenxKey(buf, car(o))
			o = cdr(o)
		}
		return shenxKey(buf, o)
	case scmHeadVector:
		v := mustVector(o)
		buf = append(buf, 'v')
		buf = binary.AppendUvarint(buf, uint64(len(v)))
		for _, x := range v {
			buf = shenxKey(buf, x)
		}
		return buf
	default:
		// Symbols, booleans, streams, procedures, natives, errors: `=` is
		// identity for all of them, and they are never moved by the Go GC.
		buf = append(buf, 'a')
		return binary.LittleEndian.AppendUint64(buf, uint64(uintptr(unsafe.Pointer(o))))
	}
}
