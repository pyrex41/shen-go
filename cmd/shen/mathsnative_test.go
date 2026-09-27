package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"
)

// TestMathsNativesMatchStdlib is the equivalence audit for mathsnative.go.
// Every case runs twice through the same binary, once with the native
// floor/ceiling/round/mod/power/gcd/lcd/isqrt and once with SHEN_NO_MATHS_NATIVE=1 (StLib's own
// Shen definitions), and the printed results, error text included, must be
// byte-identical. Numbers print shortest-round-trip, so equal text is equal
// float64 values.
//
// Inputs are confined to where the Shen definitions terminate in reasonable
// time: finite values below 2^53 in magnitude, power exponents 0..100000,
// small gcd/lcd arguments, isqrt up to 10^12. Outside their domains the
// natives defer to the Shen definitions (mathsnative.go), which the
// non-number, non-integer, zero and division-by-zero cases exercise.
func TestMathsNativesMatchStdlib(t *testing.T) {
	if testing.Short() {
		t.Skip("runs the stdlib's Shen maths loops on ~4000 cases")
	}
	exprs := mathsAuditExprs()
	var script strings.Builder
	for _, e := range exprs {
		fmt.Fprintf(&script, "(output \"~S~%%\" (trap-error %s (/. E (cn \"error: \" (error-to-string E)))))\n", e)
	}
	path := filepath.Join(t.TempDir(), "audit.shen")
	if err := os.WriteFile(path, []byte(script.String()), 0o600); err != nil {
		t.Fatal(err)
	}
	bin := buildShen(t)
	run := func(env ...string) []string {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()
		cmd := exec.CommandContext(ctx, bin, "script", path)
		cmd.Env = append(os.Environ(), env...)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("script failed (%v): %v\n%s", env, err, out)
		}
		return strings.Split(strings.TrimRight(string(out), "\n"), "\n")
	}
	native := run()
	shen := run("SHEN_NO_MATHS_NATIVE=1")
	if len(native) != len(exprs) || len(shen) != len(exprs) {
		t.Fatalf("got %d native and %d shen lines for %d cases", len(native), len(shen), len(exprs))
	}
	bad := 0
	for i, e := range exprs {
		if native[i] != shen[i] {
			bad++
			if bad <= 20 {
				t.Errorf("%s: native %s, stdlib %s", e, native[i], shen[i])
			}
		}
	}
	if bad > 0 {
		t.Fatalf("%d of %d cases differ", bad, len(exprs))
	}
	t.Logf("%d cases agree", len(exprs))
}

// mathsAuditExprs builds the audit's Shen expressions.
func mathsAuditExprs() []string {
	num := func(f float64) string {
		s := strconv.FormatFloat(f, 'f', -1, 64)
		if f == float64(int64(f)) && f > -1e15 && f < 1e15 {
			return s
		}
		return strconv.FormatFloat(f, 'g', -1, 64)
	}
	var vals []string
	for _, v := range []string{
		"0", "1", "-1", "7", "-7", "2147483648", "4294967295", "4294967296",
		"4503599627370496", "9007199254740991", "-9007199254740991",
		"0.1", "-0.1", "0.5", "-0.5", "0.49999999999999994", "0.5000000000000001",
		"1e-300", "-1e-300", "0.00001", "123.456", "-123.456",
		"1000000000000000.25", "2251799813685248.5", "-2251799813685248.5",
		"(/ 1 3)", "(/ -2 3)", "(/ 10 4)", "(* 3 0.1)", "(- 0 0.0)",
	} {
		vals = append(vals, v)
	}
	for k := -6; k <= 6; k++ {
		vals = append(vals, num(float64(k)+0.5), num(float64(k)+0.25), num(float64(k)-0.75))
	}
	r := rand.New(rand.NewSource(20260927))
	for i := 0; i < 300; i++ {
		mag := r.Intn(19) - 3 // 10^-3 .. 10^15
		f := r.Float64() * mathPow10(mag)
		if r.Intn(2) == 0 {
			f = -f
		}
		if i%5 == 0 { // exact halves
			f = float64(int64(f)) + 0.5
		}
		vals = append(vals, num(f))
	}

	var exprs []string
	for _, fn := range []string{"floor", "ceiling", "round"} {
		for _, v := range vals {
			exprs = append(exprs, "("+fn+" "+v+")")
		}
		exprs = append(exprs, "("+fn+" a)", "("+fn+" \"x\")", "("+fn+" [])")
		exprs = append(exprs, "(map "+fn+" [1.5 -1.5 2.5 -2.5 0.5 -0.5])")
	}

	xs := []string{"0", "7", "-7", "13", "-13", "20550170533348", "4294967301",
		"4503599627370495", "-4503599627370495", "7.5", "-7.5", "0.3", "(/ 1 3)", "1e-10"}
	ys := []string{"1", "3", "-3", "2", "4294967296", "0.5", "-2.5", "10", "7", "(/ 2 3)", "1e15"}
	for _, x := range xs {
		for _, y := range ys {
			exprs = append(exprs, "(mod "+x+" "+y+")")
		}
		exprs = append(exprs, "(mod "+x+" 0)")
	}
	for i := 0; i < 300; i++ {
		x := r.Int63n(1<<50) - 1<<49
		y := r.Int63n(1<<32) + 1
		if r.Intn(3) == 0 {
			y = -y
		}
		exprs = append(exprs, fmt.Sprintf("(mod %d %d)", x, y))
	}
	for i := 0; i < 150; i++ {
		x := (r.Float64() - 0.5) * mathPow10(r.Intn(12))
		y := (r.Float64() - 0.5) * mathPow10(r.Intn(6))
		exprs = append(exprs, "(mod "+num(x)+" "+num(y)+")")
	}
	exprs = append(exprs,
		"(mod a 1)", "(mod 1 a)", "(mod \"x\" 2)",
		"(map (mod 17) [2 3 5 -4])",
		"(div 17 5)", "(div -17 5)", "(div 17.5 2)",
		"(modf 3.75)", "(modf -3.75)", "(maths.float->pair 12.5)", "(maths.float->pair -12.5)",
		"(do (set maths.*seed* 95795) [(random 1 100) (random 1 100) (random -50 50)])",
	)

	// power: only integer exponents 0..100000; the original's recursion
	// overflows the Go stack on anything else.
	for _, n := range []string{"0", "1", "-1", "2", "-2", "3", "10", "1.5", "-0.5", "0.1", "(/ 1 3)", "1e10", "1e300", "-1e300"} {
		for _, m := range []string{"0", "1", "2", "3", "10", "31", "52", "53", "64", "100", "1023", "1100"} {
			exprs = append(exprs, "(power "+n+" "+m+")")
		}
	}
	exprs = append(exprs, "(power 1 100000)", "(power 1.0000001 100000)", "(power -1 99999)",
		"(power a 0)", "(power a 2)", "(power 2 a)", "(map (power 2) [0 1 8])")

	// gcd, lcd: the originals loop up to the smaller argument, so random
	// pairs stay small; the large-magnitude cases keep one argument small.
	pairs := [][2]string{
		{"12", "18"}, {"-12", "18"}, {"12", "-18"}, {"-12", "-18"}, {"5", "5"}, {"-5", "5"},
		{"0", "7"}, {"7", "0"}, {"0", "0"}, {"1", "1"}, {"1", "9"}, {"9", "1"}, {"2", "4"},
		{"4", "6"}, {"9", "15"}, {"-9", "15"}, {"-4", "6"}, {"7", "11"}, {"21", "35"}, {"49", "77"},
		{"4503599627370495", "15"}, {"15", "4503599627370495"}, {"4503599627370495", "4503599627370490"},
		{"9007199254740991", "7"}, {"-9007199254740991", "3"}, {"4294967296", "96"}, {"96", "4294967296"},
		{"7.5", "3"}, {"3", "7.5"}, {"9.5", "3"}, {"a", "3"}, {"3", "a"}, {"\"x\"", "2"},
	}
	for i := 0; i < 150; i++ {
		a, b := r.Int63n(20000)-2000, r.Int63n(20000)-2000
		if i%3 == 0 { // share a factor
			f := r.Int63n(60) + 2
			a, b = a/f*f, b/f*f
		}
		pairs = append(pairs, [2]string{strconv.FormatInt(a, 10), strconv.FormatInt(b, 10)})
	}
	for _, fn := range []string{"gcd", "lcd"} {
		for _, p := range pairs {
			exprs = append(exprs, "("+fn+" "+p[0]+" "+p[1]+")")
		}
	}

	// isqrt counts up to the root, so N stays at or below 10^12.
	for _, n := range []string{"0", "1", "2", "3", "4", "15", "16", "17", "-1", "-5", "-0.5", "2.5", "0.99",
		"1e-300", "999999999999", "1000000000000", "999998000001", "999998000000", "123456.789",
		"(/ 1 3)", "a", "\"x\"", "[]"} {
		exprs = append(exprs, "(isqrt "+n+")")
	}
	for i := 0; i < 100; i++ {
		v := r.Int63n(1_000_000_000)
		if i%4 == 0 {
			w := r.Int63n(30000)
			v = w * w
		}
		if i%4 == 1 {
			w := r.Int63n(30000) + 1
			v = w*w - 1
		}
		exprs = append(exprs, "(isqrt "+strconv.FormatInt(v, 10)+")")
		if i%5 == 0 {
			exprs = append(exprs, "(isqrt "+num(float64(v)+0.5)+")")
		}
	}
	return exprs
}

func mathPow10(n int) float64 {
	f := 1.0
	for ; n > 0; n-- {
		f *= 10
	}
	for ; n < 0; n++ {
		f /= 10
	}
	return f
}
