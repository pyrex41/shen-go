package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/tiancaiamao/shen-go/kl"
)

var (
	symReaderCons    = kl.MakeSymbol("cons")
	symReaderVector  = kl.MakeSymbol("vector")
	symReaderBarBang = kl.MakeSymbol("bar!")
	symReaderLcurly  = kl.MakeSymbol("{")
	symReaderRcurly  = kl.MakeSymbol("}")
	symReaderSemi    = kl.MakeSymbol(";")
	symReaderColon   = kl.MakeSymbol(":")
	symReaderColonEq = kl.MakeSymbol(":=")
	symReaderComma   = kl.MakeSymbol(",")
)

var errReaderUnsupported = errors.New("native reader: unsupported syntax, falling back")

type shenReader struct {
	data []byte
	pos  int
}

// parseShenFile reads a .shen source file and returns the equivalent of the
// Shen `<s-exprs>` parser output (a Shen list of top-level forms), without
// running `process-sexprs`. Returns errReaderUnsupported for syntax outside the
// supported subset (e.g. the `($ ...)` splice macro), so callers can fall back
// to the Shen reader.
func parseShenFile(path string) (kl.Obj, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	r := &shenReader{data: data}
	forms, err := r.readForms(0)
	if err != nil {
		return nil, err
	}
	return forms, nil
}

// readForms reads forms until the given closer byte (0 = EOF, ')' or ']').
// Returns a Shen list of the forms read.
func (r *shenReader) readForms(closer byte) (kl.Obj, error) {
	var forms []kl.Obj
	for {
		if err := r.skipSpaceAndComments(); err != nil {
			return nil, err
		}
		if r.pos >= len(r.data) {
			if closer != 0 {
				return nil, fmt.Errorf("native reader: unexpected EOF, expected %q", closer)
			}
			return sliceToShenList(forms), nil
		}
		b := r.data[r.pos]
		if closer != 0 && b == closer {
			r.pos++
			return sliceToShenList(forms), nil
		}
		if b == ')' || b == ']' {
			return nil, fmt.Errorf("native reader: unexpected %q at pos %d", b, r.pos)
		}
		form, err := r.readToken()
		if err != nil {
			return nil, err
		}
		forms = append(forms, form)
	}
}

func sliceToShenList(forms []kl.Obj) kl.Obj {
	out := kl.Nil
	for i := len(forms) - 1; i >= 0; i-- {
		out = kl.Cons(forms[i], out)
	}
	return out
}

func (r *shenReader) readToken() (kl.Obj, error) {
	b := r.data[r.pos]
	switch b {
	case '(':
		r.pos++
		list, err := r.readForms(')')
		if err != nil {
			return nil, err
		}
		// Reject `($ X)` splice form: not supported by the native reader.
		if isDollarSplice(list) {
			return nil, errReaderUnsupported
		}
		return list, nil
	case '[':
		r.pos++
		list, err := r.readForms(']')
		if err != nil {
			return nil, err
		}
		return consForm(list)
	case '"':
		return r.readString()
	case '{':
		r.pos++
		return symReaderLcurly, nil
	case '}':
		r.pos++
		return symReaderRcurly, nil
	case '|':
		r.pos++
		return symReaderBarBang, nil
	case ';':
		r.pos++
		return symReaderSemi, nil
	case ':':
		r.pos++
		if r.pos < len(r.data) && r.data[r.pos] == '=' {
			r.pos++
			return symReaderColonEq, nil
		}
		return symReaderColon, nil
	case ',':
		r.pos++
		return symReaderComma, nil
	}
	return r.readAtom()
}

func (r *shenReader) readAtom() (kl.Obj, error) {
	start := r.pos
	for r.pos < len(r.data) && !isShenDelim(r.data[r.pos]) {
		r.pos++
	}
	if r.pos == start {
		return nil, fmt.Errorf("native reader: unexpected character %q at pos %d", r.data[start], start)
	}
	s := string(r.data[start:r.pos])
	if s == "<>" {
		return kl.Cons(symReaderVector, kl.Cons(kl.MakeInteger(0), kl.Nil)), nil
	}
	if i, err := strconv.Atoi(s); err == nil {
		return kl.MakeInteger(i), nil
	}
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return kl.MakeNumber(f), nil
	}
	return kl.MakeSymbol(s), nil
}

func (r *shenReader) readString() (kl.Obj, error) {
	r.pos++ // skip opening "
	var buf []byte
	for r.pos < len(r.data) {
		c := r.data[r.pos]
		if c == '"' {
			r.pos++
			return kl.MakeString(string(buf)), nil
		}
		// Shen control-char escape: c#N;
		if c == 'c' && r.pos+2 < len(r.data) && r.data[r.pos+1] == '#' {
			j := r.pos + 2
			for j < len(r.data) && r.data[j] >= '0' && r.data[j] <= '9' {
				j++
			}
			if j > r.pos+2 && j < len(r.data) && r.data[j] == ';' {
				if n, err := strconv.Atoi(string(r.data[r.pos+2 : j])); err == nil && n >= 0 && n <= 255 {
					buf = append(buf, byte(n))
					r.pos = j + 1
					continue
				}
			}
		}
		buf = append(buf, c)
		r.pos++
	}
	return nil, fmt.Errorf("native reader: unterminated string at pos %d", r.pos)
}

func (r *shenReader) skipSpaceAndComments() error {
	for r.pos < len(r.data) {
		b := r.data[r.pos]
		if b == ' ' || b == '\t' || b == '\n' || b == '\r' {
			r.pos++
			continue
		}
		if b == '\\' && r.pos+1 < len(r.data) {
			n := r.data[r.pos+1]
			if n == '\\' {
				r.pos += 2
				for r.pos < len(r.data) && r.data[r.pos] != '\n' && r.data[r.pos] != '\r' {
					r.pos++
				}
				continue
			}
			if n == '*' {
				r.pos += 2
				if err := r.skipBlockComment(); err != nil {
					return err
				}
				continue
			}
		}
		return nil
	}
	return nil
}

func (r *shenReader) skipBlockComment() error {
	depth := 1
	for r.pos < len(r.data) {
		if r.pos+1 < len(r.data) {
			a, b := r.data[r.pos], r.data[r.pos+1]
			if a == '\\' && b == '*' {
				depth++
				r.pos += 2
				continue
			}
			if a == '*' && b == '\\' {
				r.pos += 2
				depth--
				if depth == 0 {
					return nil
				}
				continue
			}
		}
		r.pos++
	}
	return fmt.Errorf("native reader: unterminated block comment")
}

func isShenDelim(b byte) bool {
	switch b {
	case ' ', '\t', '\n', '\r',
		'(', ')', '[', ']', '"',
		'{', '}', '|', ':', ';', ',', '\\':
		return true
	}
	return false
}

// consForm matches the kernel `cons-form` definition: it folds a list of forms
// from `[ ... ]` into the corresponding (cons X (cons Y ...)) chain, honoring
// the `bar!` separator for explicit tails.
func consForm(list kl.Obj) (kl.Obj, error) {
	if list == kl.Nil {
		return kl.Nil, nil
	}
	elts := kl.ListToSlice(list)
	return consFormSlice(elts)
}

func consFormSlice(elts []kl.Obj) (kl.Obj, error) {
	if len(elts) == 0 {
		return kl.Nil, nil
	}
	if len(elts) >= 2 && elts[1] == symReaderBarBang {
		if len(elts) != 3 {
			return nil, fmt.Errorf("native reader: misapplication of |")
		}
		return kl.Cons(symReaderCons, kl.Cons(elts[0], kl.Cons(elts[2], kl.Nil))), nil
	}
	rest, err := consFormSlice(elts[1:])
	if err != nil {
		return nil, err
	}
	return kl.Cons(symReaderCons, kl.Cons(elts[0], kl.Cons(rest, kl.Nil))), nil
}

// isDollarSplice detects the `($ X)` reader macro that splices the exploded
// chars of X. It is rare and not supported by the native reader.
func isDollarSplice(list kl.Obj) bool {
	if list == kl.Nil {
		return false
	}
	head := kl.Car(list)
	if !kl.IsSymbol(head) {
		return false
	}
	return kl.GetSymbol(head) == "$"
}
