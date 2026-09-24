package main

import (
	"testing"
	"time"
)

// End-to-end checks that InstallKernelFast's natives agree with the kernel
// after a full boot. Complements kl/kernelfast_test.go, which cannot load
// sys.kl.
func TestKernelFastAfterBoot(t *testing.T) {
	bin := buildShen(t)
	for _, tc := range []struct{ expr, want string }{
		{`(not true)`, "false"},
		{`(not false)`, "true"},
		{`(empty? ())`, "true"},
		{`(empty? [1])`, "false"},
		{`(boolean? true)`, "true"},
		{`(boolean? 0)`, "false"},
		{`(integer? 7)`, "true"},
		{`(integer? 7.5)`, "false"},
		{`(symbol? foo)`, "true"},
		{`(symbol? {)`, "true"},
		{`(symbol? true)`, "false"},
		{`(variable? Foo)`, "true"},
		{`(variable? foo)`, "false"},
		{`(length [1 2 3])`, "3"},
		{`(reverse [1 2 3])`, "[3 2 1]"},
		{`(append [1 2] [3])`, "[1 2 3]"},
		{`(element? 2 [1 2 3])`, "true"},
		{`(fst (@p 1 2))`, "1"},
		{`(snd (@p a b))`, "b"},
		{`(tuple? (@p 1 2))`, "true"},
		{`(vector? (@p 1 2))`, "false"},
		{`(hdstr "xy")`, `x`},
		{`(fail)`, `...`}, // the printer's rendering of the fail value, on every port
		// Issue #54: the fail value is the symbol shen.fail!, as on shen-cl
		// and shen-scheme; `...` is only how it prints.
		{`(= (fail) shen.fail!)`, "true"},
		{`(str (fail))`, "shen.fail!"},
		{`(= (<-address (vector 1) 1) shen.fail!)`, "true"},
		{`(trap-error (<-vector (address-> (vector 3) 1 shen.fail!) 1) (/. E (error-to-string E)))`, "vector element not found"},
		{`(<-vector (address-> (vector 3) 1 (intern "...")) 1)`, "..."},
		{`(let V (vector 2) (do (vector-> V 1 9) (<-vector V 1)))`, "9"},
		{`(let D (vector 8) (do (put a b 1 D) (get a b D)))`, "1"},
		{`(let D (vector 8) (do (put a b 1 D) (put a b 2 D) (get a b D)))`, "2"},
		{`(trap-error (<-vector (vector 1) 1) (/. E true))`, "true"},
		{`(head [1 2 3])`, "1"},
		{`(tail [1 2])`, "[2]"},
		{`(nth 2 [a b c])`, "b"},
		{`(concat a b)`, "ab"},
		{`(== 1 1)`, "true"},
		{`(sum [1 2 3])`, "6"},
		{`(adjoin 1 [1 2])`, "[1 2]"},
		{`(remove 2 [1 2 3])`, "[1 3]"},
		{`(union [1 2] [2 3])`, "[1 2 3]"},
		{`(map (/. X (+ X 1)) [1 2 3])`, "[2 3 4]"},
		{`(bound? *os*)`, "true"},
		{`(bound? definitely-unbound-xyz)`, "false"},
		{`(shen.digit? 48)`, "true"},
		{`(shen.uppercase? 65)`, "true"},
		{`(shen.alphanums? "Foo-1")`, "true"},
		{`(let D (vector 8) (do (put a b 1 D) (do (unput a b D) (trap-error (get a b D) (/. E true)))))`, "true"},
		{`(trap-error (value missing-xyz-123) (/. E 42))`, "42"},
		// Error texts and edge behaviour the equivalence audit (kl/equiv.json)
		// found diverging from the kernel bodies; each is pinned to the KL.
		{`(trap-error (not 0) (/. E (error-to-string E)))`, "if requires a boolean"},
		{`(trap-error (vector -1) (/. E (error-to-string E)))`, "index 0 out of range 0"},
		{`(trap-error (vector 1.5) (/. E (error-to-string E)))`, "index 2 out of range 2"},
		{`(trap-error (nth 0 [a b]) (/. E (error-to-string E)))`, "nth applied to -2, []"},
		{`(trap-error (nth 3 [a b]) (/. E (error-to-string E)))`, "nth applied to 1, []"},
		{`(trap-error (nth a [a b]) (/. E (error-to-string E)))`, "mustNumber"},
		{`(trap-error (string->symbol "1") (/. E (error-to-string E)))`, `cannot intern "1" to a symbol`},
		{`(shen.misc? "a")`, "false"},
		{`(trap-error (<-vector (vector 3) -0.5) (/. E (error-to-string E)))`, "-0.5 is not a valid integer"},
		{`(trap-error (vector-> (vector 3) -0.5 x) (/. E (error-to-string E)))`, "-0.5 is not a valid integer"},
		{`(let V (vector 5) (do (unput zz yy V) (trap-error (<-vector V (hash zz 5)) (/. E none))))`, "[]"},
		{`(trap-error (unput a b (vector 0)) (/. E (error-to-string E)))`, "index 1 out of range 1"},
		// nth's error path runs shen.app in KL; nothing may leak to stdout.
		{`(trap-error (nth 0 [a b]) (/. E ok))`, "ok"},
	} {
		if got := evalBounded(t, bin, tc.expr, 30*time.Second); got != tc.want {
			t.Errorf("%s: got %q, want %q", tc.expr, got, tc.want)
		}
	}
}
