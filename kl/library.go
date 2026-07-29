package kl

import (
	"math"
	"os"
	"path"
	"strconv"
)

func PackagePath() string {
	gopath := os.Getenv("GOPATH")
	return path.Join(gopath, "src/github.com/tiancaiamao/shen-go")
}

func cadr(o Obj) Obj {
	return car(cdr(o))
}

func caddr(o Obj) Obj {
	return car(cdr(cdr(o)))
}

func cdddr(o Obj) Obj {
	return cdr(cdr(cdr(o)))
}

func cadddr(o Obj) Obj {
	return car(cdr(cdr(cdr(o))))
}

func reverse(o Obj) Obj {
	ret := Nil
	for o != Nil {
		ret = cons(car(o), ret)
		o = cdr(o)
	}
	return ret
}

func equal(x, y Obj) Obj {
	if x == y {
		return True
	}
	if *x != *y {
		return False
	}

	switch *x {
	case scmHeadNumber:
		if !isFixnum(x) && !isFixnum(y) {
			if mustNumber(x) == mustNumber(y) {
				return True
			}
		}
		// x == y is checked already
		return False
	case scmHeadString:
		if mustString(x) != mustString(y) {
			return False
		}
	case scmHeadBoolean:
		if x != y {
			return False
		}
	case scmHeadSymbol:
		if x != y {
			return False
		}
	case scmHeadNull:
		if *y != scmHeadNull {
			return False
		}
	case scmHeadPair:
		// TODO: maybe cycle exists!
		if x != y {
			if equal(car(x), car(y)) == False {
				return False
			}
			if equal(cdr(x), cdr(y)) == False {
				return False
			}
		}
	case scmHeadStream, scmHeadProcedure, scmHeadBytecodeFunc /* , scmHeadPrimitive */ :
		if x != y {
			return False
		}
	case scmHeadVector:
		v1 := mustVector(x)
		v2 := mustVector(y)
		if len(v1) != len(v2) {
			return False
		}
		for i := 0; i < len(v1); i++ {
			if v1[i] != nil || v2[i] != nil {
				if equal(v1[i], v2[i]) != True {
					return False
				}
			}
		}
	}

	return True
}

func listLength(l Obj) int {
	count := 0
	for *l == scmHeadPair {
		count++
		l = cdr(l)
	}
	return count
}

func ListToSlice(l Obj) []Obj {
	var ret []Obj
	for *l == scmHeadPair {
		ret = append(ret, car(l))
		l = cdr(l)
	}
	return ret
}

func Cadr(o Obj) Obj {
	return cadr(o)
}

func Car(o Obj) Obj {
	return car(o)
}

func Cdr(o Obj) Obj {
	return cdr(o)
}

func Cons(x, y Obj) Obj {
	return cons(x, y)
}

// isInteger determinate whether a float64 is actually a precise integer.
// Judge is according to IEEE754 standard.
func isPreciseInteger(f float64) bool {
	// math.Ilogb answers math.MaxInt32 for +-Inf and NaN, which would fall
	// straight into the `exp >= 52` shortcut below and report them as
	// integers. They are not integers under any reading, and callers narrow
	// on the strength of this answer.
	if math.IsInf(f, 0) || math.IsNaN(f) {
		return false
	}

	exp := math.Ilogb(f)
	if exp < 0 && exp != math.MinInt32 {
		return false
	}

	if exp >= 52 {
		return true
	}

	bits := math.Float64bits(f)
	return (bits << uint(12+exp)) == 0
}

// fitsInt reports whether f can be converted to Go's int without changing its
// value. Go leaves an out-of-range float64 -> int conversion
// implementation-defined (arm64 saturates to +-maxint64), so every narrowing
// of a Shen number has to be guarded by this. NaN compares false both ways and
// is correctly rejected.
func fitsInt(f float64) bool {
	return f >= minIntAsFloat && f < maxIntAsFloat
}

// formatNumber renders a Shen number the way both printers (scmHead.GoString
// and the `str` primitive) must agree on.
//
// Integral values that fit in an int print without a decimal point, as the
// kernel expects. Everything else uses Go's shortest round-trip form (issue
// #11: 2.5, not 2.500000), which is also what codegen emits for constants.
// That covers the values too large to narrow -- 1e19 is exactly representable
// and perfectly ordinary, but no int can hold it, so it prints as a float
// rather than as the saturated 9223372036854775807 it used to become.
//
// Non-finite values get the spelling shen-lua and shen-rust use. This is only
// about not printing a number the value is not; it says nothing about whether
// an overflowing operation should have signalled in the first place.
func formatNumber(f float64) string {
	switch {
	case math.IsNaN(f):
		return "nan"
	case math.IsInf(f, 1):
		return "inf"
	case math.IsInf(f, -1):
		return "-inf"
	case isPreciseInteger(f) && fitsInt(f):
		return strconv.FormatInt(int64(f), 10)
	default:
		return strconv.FormatFloat(f, 'g', -1, 64)
	}
}

func BindSymbolFunc(sym Obj, f Obj) {
	mustSymbol(sym).function = f
}

// Equal reports deep structural equality of two objects, comparing numbers by
// their full float64 value (not just printed form). Exposed for tests that
// prove the Go-native reader matches the interpreted reader bit-for-bit.
func Equal(x, y Obj) bool {
	return equal(x, y) == True
}

// IsPair reports whether o is a cons pair and, if so, returns its cdr. Exposed
// for tests that walk reader output without panicking on improper lists.
func IsPair(o Obj) (bool, Obj) {
	if ok, p := isPair(o); ok {
		return true, p.cdr
	}
	return false, Nil
}
