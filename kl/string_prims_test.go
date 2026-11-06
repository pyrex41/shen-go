package kl

import (
	"fmt"
	"testing"
)

// The string primitives walk bytes instead of converting to []rune. These pin
// them to the conversions they replaced, on valid and invalid UTF-8 alike.
var stringPrimSamples = []string{
	"", "a", "ab", "hello world", "\x7f", "\x00z",
	"é", "éa", "aé", "日本語", "a日b本", "😀", "😀x", "x😀",
	"\xff", "a\xff", "\xffa", "\xc3", "a\xc3", "\xc3a", "\xe6\x97", "\xe6\x97a", "a\xe6\x97",
	"\xed\xa0\x80", "a\xed\xa0\x80b", "\xc0\x80", "\xf4\x90\x80\x80", "\xef\xbf\xbd", "a\xef\xbf\xbdb",
}

func oldTail(s string) string { return string([]rune(s)[1:]) }

func checkStringPrims(t *testing.T, s string) {
	t.Helper()
	runes := []rune(s)
	if len(s) > 0 {
		if got, want := stringTail(s), oldTail(s); got != want {
			t.Errorf("stringTail(%q) = %q, want %q", s, got, want)
		}
		if got, want := GetString(PrimTailString(MakeString(s))), oldTail(s); got != want {
			t.Errorf("tlstr(%q) = %q, want %q", s, got, want)
		}
		if got, want := GetInteger(PrimStringToNumber(MakeString(s))), int(runes[0]); got != want {
			t.Errorf("string->n(%q) = %d, want %d", s, got, want)
		}
	}
	for n := -1; n <= len(runes)+1; n++ {
		got, ok := stringRuneAt(s, n)
		inRange := n >= 0 && n < len(runes)
		if ok != inRange {
			t.Errorf("stringRuneAt(%q, %d) ok = %v, want %v", s, n, ok, inRange)
			continue
		}
		if inRange {
			if want := string(runes[n]); got != want {
				t.Errorf("stringRuneAt(%q, %d) = %q, want %q", s, n, got, want)
			}
			if pos := GetString(PrimPos(MakeString(s), MakeInteger(n))); pos != got {
				t.Errorf("pos(%q, %d) = %q, want %q", s, n, pos, got)
			}
		} else {
			want := fmt.Sprintf("%d is not valid index for %s", n, string(runes))
			if msg := primError(func() { PrimPos(MakeString(s), MakeInteger(n)) }); msg != want {
				t.Errorf("pos(%q, %d) error = %q, want %q", s, n, msg, want)
			}
		}
	}
}

func primError(f func()) (msg string) {
	defer func() {
		switch r := recover().(type) {
		case nil:
			msg = "<no error>"
		case Obj:
			msg = mustError(r).err
		case error:
			msg = "runtime: " + r.Error()
		default:
			msg = fmt.Sprint(r)
		}
	}()
	f()
	return
}

func TestStringPrimsMatchRuneConversion(t *testing.T) {
	for _, s := range stringPrimSamples {
		checkStringPrims(t, s)
	}
	for c := 0; c < 256; c++ {
		checkStringPrims(t, string([]byte{byte(c)}))
		checkStringPrims(t, string([]byte{'q', byte(c), 'r'}))
	}
}

func FuzzStringPrimsMatchRuneConversion(f *testing.F) {
	for _, s := range stringPrimSamples {
		f.Add(s)
	}
	f.Fuzz(checkStringPrims)
}

func TestStringPrimEmptyErrors(t *testing.T) {
	empty := MakeString("")
	if got := primError(func() { PrimTailString(empty) }); got != "empty string" {
		t.Errorf("tlstr of empty: %q", got)
	}
	if got := primError(func() { TypedStringTailValue("") }); got != "empty string" {
		t.Errorf("typed tail of empty: %q", got)
	}
	if got := primError(func() { PrimStringToNumber(empty) }); got != "runtime: runtime error: index out of range [0] with length 0" {
		t.Errorf("string->n of empty: %q", got)
	}
	if got := primError(func() { PrimPos(MakeInteger(1), MakeInteger(0)) }); got == "<no error>" {
		t.Errorf("pos of a non-string did not raise")
	}
	if got := primError(func() { TypedStringIndexValue("a\xffb", 7) }); got != "7 is not valid index for a\xffb" {
		t.Errorf("typed index out of range: %q", got)
	}
}

func TestSingleCharStringsAreShared(t *testing.T) {
	for c := 0; c < 128; c++ {
		o := PrimNumberToString(MakeInteger(c))
		if GetString(o) != string(rune(c)) || !IsString(o) {
			t.Fatalf("n->string %d = %q", c, GetString(o))
		}
		if o != MakeString(string(rune(c))) {
			t.Fatalf("one-character string %d not shared", c)
		}
	}
	if !allocFree(func() { PrimPos(MakeString("xyz"), MakeInteger(1)) }, 1) {
		t.Errorf("pos on ASCII allocated more than its argument")
	}
}

func allocFree(f func(), max float64) bool { return testing.AllocsPerRun(100, f) <= max }
