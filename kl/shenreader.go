package kl

// Go-native implementation of the Shen kernel reader's tokenize/parse stage —
// the defcc grammar <s-exprs> in kernel/sources/reader.shen.
//
// The kernel's (read-file F) pipeline is:
//
//	(read-file-as-bytelist F)  ; raw bytes        — already native (cheap)
//	(compile (/. X (<s-exprs> X)) Bytes)          — tokenize + parse to raw
//	                                                 s-expr forms (EXPENSIVE,
//	                                                 interpreted per-byte)
//	(process-sexprs S-exprs)   ; macroexpand,      — kept interpreted (it has
//	                             arity side effects,  global side effects and
//	                             application currying  is comparatively cheap)
//
// ShenReadSExprs reimplements only the middle stage — bytes -> raw s-expr form
// list — directly in Go. Tokenizing one identifier in the interpreted reader is
// thousands of trampoline iterations through eval() and the parser-combinator
// machinery; here it is a byte scan. The result must be *byte-identical* to the
// interpreted <s-exprs> for every downstream stage to agree, so this is a
// faithful, rule-for-rule port of the defcc grammar, including:
//
//   - the ordered-choice / backtracking structure of defcc,
//   - the [ ] cons-form lowering and the | (bar!) dotted-tail handling,
//   - the { } | ; : := , structural tokens (interned as literal symbols),
//   - <> -> (vector 0),
//   - string literals with the c#NN; character-code escape, built byte-by-byte
//     via n->string (string(rune(b))) exactly like the kernel's cn chain,
//   - numbers (sign, integer, float, e-notation) computed with the kernel's own
//     expt/compute-integer/compute-fraction arithmetic so the resulting float64
//     bit pattern matches the interpreted reader (NOT strconv.ParseFloat, which
//     can round differently).
//
// Anything this reader is not certain it handles identically — the $ splice
// reader macro, a |-misapplication, malformed/unterminated input, or any
// leftover unconsumed bytes — raises errShenReaderFallback so the caller can
// delegate to the interpreted kernel reader, which produces the canonical
// result (or the canonical reader-error). Correctness never depends on this
// reader being complete, only on it being right when it answers.

import "errors"

// errShenReaderFallback signals that the Go-native reader declined to parse the
// input (an unhandled construct or malformed bytes); the caller should fall
// back to the interpreted kernel reader.
var errShenReaderFallback = errors.New("shen reader: fall back to interpreted reader")

// Symbols emitted as literal structural tokens by the grammar. Cached once so
// the hot path does no trie lookups for them.
var (
	symLcurly    = MakeSymbol("{")
	symRcurly    = MakeSymbol("}")
	symBar       = MakeSymbol("bar!")
	symSemicolon = MakeSymbol(";")
	symColonEq   = MakeSymbol(":=")
	symColon     = MakeSymbol(":")
	symComma     = MakeSymbol(",")
	symConsForm  = MakeSymbol("cons")
	symVectorTok = MakeSymbol("vector")
	symDollar    = MakeSymbol("$")
)

type shenReader struct {
	data []byte
}

// ShenReadSExprs parses raw .shen/.kl source bytes into the list of raw s-expr
// forms the kernel's <s-exprs> would produce. On any construct it does not
// handle identically, or on leftover unconsumed input, it returns
// errShenReaderFallback and the caller must use the interpreted reader.
func ShenReadSExprs(data []byte) (result Obj, err error) {
	defer func() {
		if r := recover(); r != nil {
			if r == errShenReaderFallback {
				result = Nil
				err = errShenReaderFallback
				return
			}
			panic(r)
		}
	}()

	r := &shenReader{data: data}
	forms, p := r.sexprs(0)
	// <s-exprs> always succeeds (its last clause is <e> := []), but for read-file
	// the parse must consume the whole input. Leftover bytes (a stray ) or ], an
	// unterminated string/comment) are exactly the inputs on which the kernel's
	// compile would surface a reader-error; defer to the interpreted reader so it
	// produces that error verbatim.
	if p != len(data) {
		return Nil, errShenReaderFallback
	}
	return forms, nil
}

// sexprs is the <s-exprs> rule: a flat list of forms. It returns the parsed
// form list and the position one past the last byte it consumed. It never
// fails — when no clause matches the current byte it returns via <e> (the empty
// list) WITHOUT consuming, which is how nested ( [ { contexts detect their
// closing delimiter and how the top level detects end of input.
func (r *shenReader) sexprs(p int) (Obj, int) {
	if p >= len(r.data) {
		return Nil, p
	}

	switch r.data[p] {
	case '[': // <lsb> <s-exprs1> <rsb> <s-exprs2> := [(cons-form <s-exprs1>) | <s-exprs2>]
		inner, p2 := r.sexprs(p + 1)
		if p2 >= len(r.data) || r.data[p2] != ']' {
			panic(errShenReaderFallback)
		}
		rest, p3 := r.sexprs(p2 + 1)
		return cons(r.consForm(inner), rest), p3
	case '(': // <lrb> <s-exprs1> <rrb> <s-exprs2> := (add-sexpr <s-exprs1> <s-exprs2>)
		inner, p2 := r.sexprs(p + 1)
		if p2 >= len(r.data) || r.data[p2] != ')' {
			panic(errShenReaderFallback)
		}
		rest, p3 := r.sexprs(p2 + 1)
		return r.addSexpr(inner, rest), p3
	case '{': // <lcurly> <s-exprs> := [{ | <s-exprs>]
		rest, p2 := r.sexprs(p + 1)
		return cons(symLcurly, rest), p2
	case '}': // <rcurly> <s-exprs> := [} | <s-exprs>]
		rest, p2 := r.sexprs(p + 1)
		return cons(symRcurly, rest), p2
	case '|': // <bar> <s-exprs> := [bar! | <s-exprs>]
		rest, p2 := r.sexprs(p + 1)
		return cons(symBar, rest), p2
	case ';': // <semicolon> <s-exprs> := [(intern ";") | <s-exprs>]
		rest, p2 := r.sexprs(p + 1)
		return cons(symSemicolon, rest), p2
	case ':': // <colon> <equal> ... := :=  OR  <colon> ... := :
		if p+1 < len(r.data) && r.data[p+1] == '=' {
			rest, p2 := r.sexprs(p + 2)
			return cons(symColonEq, rest), p2
		}
		rest, p2 := r.sexprs(p + 1)
		return cons(symColon, rest), p2
	case ',': // <comma> <s-exprs> := [(intern ",") | <s-exprs>]
		rest, p2 := r.sexprs(p + 1)
		return cons(symComma, rest), p2
	}

	// <comment> <s-exprs> := <s-exprs>
	if np, ok := r.comment(p); ok {
		return r.sexprs(np)
	}

	// <atom> <s-exprs> := [<atom> | <s-exprs>]
	if obj, np, ok := r.atom(p); ok {
		rest, p2 := r.sexprs(np)
		return cons(obj, rest), p2
	}

	// <whitespaces> <s-exprs> := <s-exprs>
	if isWhitespace(r.data[p]) {
		np := p + 1
		for np < len(r.data) && isWhitespace(r.data[np]) {
			np++
		}
		return r.sexprs(np)
	}

	// <e> := [] — no clause matched; succeed matching nothing, leave p.
	return Nil, p
}

// consForm lowers a [ ... ] bracket list (whose elements may include the bar!
// symbol) into the equivalent (cons ... ) construction form, mirroring the
// kernel's cons-form. It panics with errShenReaderFallback on the
// |-misapplication case ([X bar! Y Z ...]) so the interpreted reader raises the
// kernel's "misapplication of |" error.
func (r *shenReader) consForm(list Obj) Obj {
	if list == Nil { // cons-form [] -> []
		return Nil
	}
	x := car(list)
	rest := cdr(list)
	if rest != Nil && car(rest) == symBar {
		afterBar := cdr(rest)
		if afterBar != Nil {
			if cdr(afterBar) == Nil { // [X bar! Y] -> [cons X Y]
				return cons(symConsForm, cons(x, cons(car(afterBar), Nil)))
			}
			panic(errShenReaderFallback) // [X bar! Y Z | _] -> error
		}
		// [X bar!] (nothing after bar!) falls through to the general clause,
		// where bar! is treated as an ordinary trailing element.
	}
	// [X | Y] -> [cons X (cons-form Y)]
	return cons(symConsForm, cons(x, cons(r.consForm(rest), Nil)))
}

// addSexpr implements the kernel's add-sexpr for a ( ... ) group: normally the
// inner list becomes a single element, except the [$ X] splice macro, which we
// do not implement — fall back so the interpreted reader's explode runs.
func (r *shenReader) addSexpr(inner, rest Obj) Obj {
	if inner != Nil && car(inner) == symDollar {
		tail := cdr(inner)
		if tail != Nil && cdr(tail) == Nil { // inner is exactly [$ X]
			panic(errShenReaderFallback)
		}
	}
	return cons(inner, rest)
}

// comment matches <comment> := <singleline> | <multiline>, returning the
// position after the comment and whether one matched.
func (r *shenReader) comment(p int) (int, bool) {
	if p+1 >= len(r.data) || r.data[p] != '\\' {
		return p, false
	}
	switch r.data[p+1] {
	case '\\': // <singleline> := \\ <shortnatters> <returns>
		q := p + 2
		for q < len(r.data) && !isReturn(r.data[q]) {
			q++
		}
		if q >= len(r.data) { // no terminating return char -> singleline fails
			return p, false
		}
		for q < len(r.data) && isReturn(r.data[q]) {
			q++
		}
		return q, true
	case '*': // <multiline> := \* <longnatter>
		if q, ok := r.longnatter(p + 2); ok {
			return q, true
		}
		return p, false
	}
	return p, false
}

// longnatter scans the body of a \* ... *\ multiline comment, honoring nested
// comments, exactly like the defcc <longnatter> ordered choice.
func (r *shenReader) longnatter(p int) (int, bool) {
	for {
		if np, ok := r.comment(p); ok { // <comment> <longnatter>
			p = np
			continue
		}
		if p+1 < len(r.data) && r.data[p] == '*' && r.data[p+1] == '\\' { // <times> <backslash>
			return p + 2, true
		}
		if p < len(r.data) { // _ <longnatter>
			p++
			continue
		}
		return p, false // unterminated
	}
}

// atom matches <atom> := <str> | <number> | <sym>.
func (r *shenReader) atom(p int) (Obj, int, bool) {
	if p >= len(r.data) {
		return nil, p, false
	}
	if r.data[p] == '"' {
		return r.str(p)
	}
	if obj, np, ok := r.number(p); ok {
		return obj, np, true
	}
	return r.sym(p)
}

// str matches <str> := <dbq> <strcontents> <dbq>, building the string content
// byte-by-byte through n->string (string(rune(b))) like the kernel's cn chain,
// and honoring the c#NN; character-code escape.
func (r *shenReader) str(p int) (Obj, int, bool) {
	if r.data[p] != '"' {
		return nil, p, false
	}
	var sb []rune
	q := p + 1
	for q < len(r.data) {
		c := r.data[q]
		if c == '"' { // closing <dbq>
			return MakeString(string(sb)), q + 1, true
		}
		// <control> := c # <integer> ; := (n->string <integer>)
		if c == 'c' && q+1 < len(r.data) && r.data[q+1] == '#' {
			j := q + 2
			start := j
			for j < len(r.data) && isDigit(r.data[j]) {
				j++
			}
			if j > start && j < len(r.data) && r.data[j] == ';' {
				n := int(computeInteger(digitsOf(r.data[start:j])))
				sb = append(sb, rune(n))
				q = j + 1
				continue
			}
			// not a valid escape: fall through, treat 'c' as <notdbq>
		}
		// <notdbq> := Byte where Byte != " := (n->string Byte)
		sb = append(sb, rune(c))
		q++
	}
	return nil, p, false // unterminated string
}

// sym matches <sym> := <alpha> <alphanums>, with the <> -> (vector 0) special
// case and intern's true/false -> booleans.
func (r *shenReader) sym(p int) (Obj, int, bool) {
	if p >= len(r.data) || !isAlpha(r.data[p]) {
		return nil, p, false
	}
	start := p
	p++
	for p < len(r.data) && isAlphanum(r.data[p]) {
		p++
	}
	name := string(r.data[start:p]) // symbol chars are all ASCII
	if name == "<>" {
		return cons(symVectorTok, cons(MakeInteger(0), Nil)), p, true
	}
	switch name { // intern: true/false fold to the booleans
	case "true":
		return True, p, true
	case "false":
		return False, p, true
	}
	return MakeSymbol(name), p, true
}

// number matches <number>, producing a Shen number whose float64 value is
// computed with the kernel's own arithmetic so it is bit-identical to the
// interpreted reader's result.
func (r *shenReader) number(p int) (Obj, int, bool) {
	f, np, ok := r.numberVal(p)
	if !ok {
		return nil, p, false
	}
	return MakeNumber(f), np, true
}

func (r *shenReader) numberVal(p int) (float64, int, bool) {
	if p >= len(r.data) {
		return 0, p, false
	}
	switch r.data[p] {
	case '-': // <minus> <number> := (- 0 <number>)
		if v, np, ok := r.numberVal(p + 1); ok {
			return 0 - v, np, true
		}
		return 0, p, false
	case '+': // <plus> <number> := <number>
		if v, np, ok := r.numberVal(p + 1); ok {
			return v, np, true
		}
		return 0, p, false
	}
	if v, np, ok := r.enumber(p); ok { // <e-number>
		return v, np, true
	}
	if v, np, ok := r.floatVal(p); ok { // <float>
		return v, np, true
	}
	return r.integerVal(p) // <integer>
}

// enumber matches <e-number> := <float> e <log10> | <integer> e <log10>.
func (r *shenReader) enumber(p int) (float64, int, bool) {
	if f, np, ok := r.floatVal(p); ok && np < len(r.data) && r.data[np] == 'e' {
		if l, np2, ok2 := r.log10(np + 1); ok2 {
			return computeE(f, l), np2, true
		}
	}
	if i, np, ok := r.integerVal(p); ok && np < len(r.data) && r.data[np] == 'e' {
		if l, np2, ok2 := r.log10(np + 1); ok2 {
			return computeE(i, l), np2, true
		}
	}
	return 0, p, false
}

// log10 matches the exponent: optional sign then an integer.
func (r *shenReader) log10(p int) (int, int, bool) {
	if p >= len(r.data) {
		return 0, p, false
	}
	switch r.data[p] {
	case '+':
		return r.log10(p + 1)
	case '-':
		if v, np, ok := r.log10(p + 1); ok {
			return -v, np, true
		}
		return 0, p, false
	}
	digits, np, ok := r.collectDigits(p)
	if !ok {
		return 0, p, false
	}
	return int(computeInteger(digits)), np, true
}

// floatVal matches <float> := <integer> . <fraction> | . <fraction>.
func (r *shenReader) floatVal(p int) (float64, int, bool) {
	if intDigits, np, ok := r.collectDigits(p); ok && np < len(r.data) && r.data[np] == '.' {
		if fracDigits, np2, ok2 := r.collectDigits(np + 1); ok2 {
			return computeInteger(intDigits) + computeFraction(fracDigits), np2, true
		}
	}
	if p < len(r.data) && r.data[p] == '.' {
		if fracDigits, np, ok := r.collectDigits(p + 1); ok {
			return computeFraction(fracDigits), np, true
		}
	}
	return 0, p, false
}

func (r *shenReader) integerVal(p int) (float64, int, bool) {
	digits, np, ok := r.collectDigits(p)
	if !ok {
		return 0, p, false
	}
	return computeInteger(digits), np, true
}

func (r *shenReader) collectDigits(p int) ([]int, int, bool) {
	start := p
	for p < len(r.data) && isDigit(r.data[p]) {
		p++
	}
	if p == start {
		return nil, start, false
	}
	return digitsOf(r.data[start:p]), p, true
}

func digitsOf(bs []byte) []int {
	digits := make([]int, len(bs))
	for i, b := range bs {
		digits[i] = int(b - '0')
	}
	return digits
}

// computeInteger, computeFraction and computeE replicate the kernel's
// arithmetic (reader.shen) operation-for-operation so the float64 results match
// the interpreted <s-exprs> bit-for-bit. They must therefore use the same
// (expt 10 N) the interpreted reader uses -- which is pow10 below, installed
// over shen.expt by InstallExactPow10 (see readerarith.go). Doing a one-shot
// strconv.ParseFloat of the whole token here instead would round differently
// from the interpreted reader and break that parity.

// computeInteger mirrors compute-integer: sum digit*10^pos over reversed digits.
func computeInteger(digits []int) float64 {
	return computeIntegerH(reverseInts(digits), 0)
}

func computeIntegerH(ds []int, exp int) float64 {
	if len(ds) == 0 {
		return 0
	}
	return pow10(exp)*float64(ds[0]) + computeIntegerH(ds[1:], exp+1)
}

// computeFraction mirrors compute-fraction: sum digit*10^(-1-i) in source order.
func computeFraction(digits []int) float64 {
	return computeFractionH(digits, -1)
}

func computeFractionH(ds []int, exp int) float64 {
	if len(ds) == 0 {
		return 0
	}
	return pow10(exp)*float64(ds[0]) + computeFractionH(ds[1:], exp-1)
}

// computeE mirrors (compute-E N Log10) := N * (expt 10 Log10).
func computeE(n float64, log10 int) float64 {
	return n * pow10(log10)
}

func reverseInts(in []int) []int {
	out := make([]int, len(in))
	for i, v := range in {
		out[len(in)-1-i] = v
	}
	return out
}

// Character classes from reader.shen.

func isDigit(b byte) bool { return b >= '0' && b <= '9' } // digit?

func isReturn(b byte) bool { return b == 9 || b == 10 || b == 13 } // return?

func isWhitespace(b byte) bool { // whitespace?
	return b == 32 || b == 13 || b == 10 || b == 9
}

// isMisc is the misc? set: = - * / + _ ? $ ! @ ~ . > < & % ' # `
func isMisc(b byte) bool {
	switch b {
	case 61, 45, 42, 47, 43, 95, 63, 36, 33, 64, 126, 46, 62, 60, 38, 37, 39, 35, 96:
		return true
	}
	return false
}

// isAlpha is alpha?: lowercase, uppercase or misc.
func isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || isMisc(b)
}

// isAlphanum is <alphanum>: alpha or numeral.
func isAlphanum(b byte) bool { return isAlpha(b) || isDigit(b) }
