package kl

import (
	"fmt"
	"unsafe"
)

// TypedObjectKind classifies an object without raising a Shen error.
func TypedObjectKind(o Obj) ValueKind {
	if o == nil {
		return KindUnknown
	}
	if isFixnum(o) {
		return KindFixnum
	}
	switch *o {
	case scmHeadNumber:
		return KindNumber
	case scmHeadPair:
		return KindPair
	case scmHeadVector:
		return KindVector
	case scmHeadNull:
		return KindNil
	case scmHeadString:
		return KindString
	case scmHeadSymbol:
		return KindSymbol
	case scmHeadBoolean:
		return KindBoolean
	case scmHeadProcedure, scmHeadNative, scmHeadBytecodeFunc:
		return KindCallable
	case scmHeadStream:
		return KindStream
	case scmHeadError:
		return KindError
	case scmHeadRaw:
		return KindRaw
	default:
		return KindUnknown
	}
}

// TypedFloat64 extracts a Shen number without raising. Fixnums and boxed
// numbers have exactly the same float64 semantics as GetNumber.
func TypedFloat64(o Obj) (float64, bool) {
	kind := TypedObjectKind(o)
	if kind != KindNumber && kind != KindFixnum {
		return 0, false
	}
	return GetNumber(o), true
}

// TypedMaterializeNumber boxes a scalar while preserving MakeNumber's fixnum,
// NaN/Inf, and float64 rounding behavior.
func TypedMaterializeNumber(v float64) Obj { return MakeNumber(v) }

// TypedDivide performs Shen's guarded float64 division. A zero denominator is
// a catchable Shen error; otherwise IEEE-754 NaN and infinity are preserved.
func TypedDivide(x, y float64) Obj {
	return TypedMaterializeNumber(TypedDivideValue(x, y))
}

// TypedDivideValue is the unboxed form used inside generated scalar regions.
func TypedDivideValue(x, y float64) float64 {
	if y == 0 {
		panic(MakeError("division by zero"))
	}
	return x / y
}

// TypedBoolean extracts the canonical Shen boolean without raising.
func TypedBoolean(o Obj) (bool, bool) {
	switch o {
	case True:
		return true, true
	case False:
		return false, true
	default:
		return false, false
	}
}

func TypedMaterializeBoolean(v bool) Obj {
	if v {
		return True
	}
	return False
}

// TypedString extracts a Shen string without raising.
func TypedString(o Obj) (string, bool) {
	if TypedObjectKind(o) != KindString {
		return "", false
	}
	return mustString(o), true
}

func TypedMaterializeString(v string) Obj { return MakeString(v) }

// TypedStringIndexValue returns the Unicode code point at index without
// materializing the result string.
func TypedStringIndexValue(value string, index int) string {
	runes := []rune(value)
	if index < 0 || index >= len(runes) {
		panic(MakeError(fmt.Sprintf("%d is not valid index for %s", index, value)))
	}
	return string(runes[index])
}

// TypedStringTailValue removes the first Unicode code point without
// materializing the result string.
func TypedStringTailValue(value string) string {
	runes := []rune(value)
	if len(runes) == 0 {
		panic(MakeError("empty string"))
	}
	return string(runes[1:])
}

// TypedStringIndex returns the Unicode code point at index. It raises the
// same catchable errors as pos for invalid type, non-negative bounds, and
// malformed numeric conversion.
func TypedStringIndex(o Obj, index int) Obj {
	value, ok := TypedString(o)
	if !ok {
		panic(MakeError("mustString"))
	}
	return TypedMaterializeString(TypedStringIndexValue(value, index))
}

// TypedStringTail returns all Unicode code points after the first one.
func TypedStringTail(o Obj) Obj {
	value, ok := TypedString(o)
	if !ok {
		panic(MakeError("mustString"))
	}
	return TypedMaterializeString(TypedStringTailValue(value))
}

// TypedPairHead and TypedPairTail are non-throwing accessors for proven pair
// values. The boolean reports whether the type guard succeeded.
func TypedPairHead(o Obj) (Obj, bool) {
	if TypedObjectKind(o) != KindPair {
		return nil, false
	}
	return (*scmPair)(unsafe.Pointer(o)).car, true
}

func TypedPairTail(o Obj) (Obj, bool) {
	if TypedObjectKind(o) != KindPair {
		return nil, false
	}
	return (*scmPair)(unsafe.Pointer(o)).cdr, true
}

// TypedVectorGet performs a checked read. An uninitialized slot has the same
// observable value as PrimVectorGet (undefined), while failures are catchable.
func TypedVectorGet(o Obj, index int) Obj {
	if TypedObjectKind(o) != KindVector {
		panic(MakeError("mustVector"))
	}
	vec := (*scmVector)(unsafe.Pointer(o)).vector
	if index < 0 || index >= len(vec) {
		panic(MakeError(fmt.Sprintf("index %d out of range %d", index, len(vec))))
	}
	if vec[index] == nil {
		return undefined
	}
	return vec[index]
}

func TypedVectorSet(o Obj, index int, value Obj) Obj {
	if TypedObjectKind(o) != KindVector {
		panic(MakeError("mustVector"))
	}
	vec := (*scmVector)(unsafe.Pointer(o)).vector
	if index < 0 || index >= len(vec) {
		panic(MakeError(fmt.Sprintf("index %d out of range %d", index, len(vec))))
	}
	vec[index] = value
	return o
}
