package codegen

import (
	"bytes"
	"strconv"
	"strings"
	"testing"

	"github.com/tiancaiamao/shen-go/kl"
)

// Shen numbers are float64. Emitting them through GetInteger truncated every
// constant to a whole number, so compiled output saw 1.5 as 1 while the
// interpreter saw 1.5 -- a silent divergence visible only in generated code.
func TestGenerateConstPreservesFractionalNumbers(t *testing.T) {
	for _, tc := range []struct {
		in   float64
		want string
	}{
		{1.5, "MakeNumber(1.5)"},
		{0.5, "MakeNumber(0.5)"},
		{-2.25, "MakeNumber(-2.25)"},
		{1e-9, "MakeNumber(1e-09)"},
		// Integral values must keep emitting cleanly, not as 3.0 or 3e+00.
		{3, "MakeNumber(3)"},
		{0, "MakeNumber(0)"},
		{-7, "MakeNumber(-7)"},
	} {
		var buf bytes.Buffer
		if err := New().generateConst(&buf, kl.MakeNumber(tc.in)); err != nil {
			t.Fatalf("%v: %v", tc.in, err)
		}
		if got := buf.String(); got != tc.want {
			t.Errorf("const %v emitted %q, want %q", tc.in, got, tc.want)
		}
	}
}

// The emitted literal is Go source; it has to parse back to the identical
// float64 or the artifact still diverges, just less obviously.
func TestGenerateConstRoundTripsExactly(t *testing.T) {
	for _, f := range []float64{
		1.5, 0.5, -2.25, 0.1, 1.0 / 3.0, 1e300, 1e-300, 12345.6789, 3, 0,
	} {
		var buf bytes.Buffer
		if err := New().generateConst(&buf, kl.MakeNumber(f)); err != nil {
			t.Fatalf("%v: %v", f, err)
		}
		lit := strings.TrimSuffix(strings.TrimPrefix(buf.String(), "MakeNumber("), ")")
		back, err := strconv.ParseFloat(lit, 64)
		if err != nil {
			t.Fatalf("%v emitted %q, which is not a parseable literal: %v", f, lit, err)
		}
		if back != f {
			t.Errorf("%v round-tripped to %v via %q", f, back, lit)
		}
	}
}

func TestGetNumberDoesNotTruncate(t *testing.T) {
	for _, f := range []float64{1.5, -0.25, 3, 0} {
		if got := kl.GetNumber(kl.MakeNumber(f)); got != f {
			t.Errorf("GetNumber(%v) = %v", f, got)
		}
	}
}
