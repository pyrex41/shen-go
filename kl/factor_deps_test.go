package kl

import (
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"
)

// factorDeps must be exactly what shen.factor can reach in the vendored kernel,
// derived the way factor_deps.go says: from shen.factor, every symbol in call
// position of a reachable defun that is not a special form, a variable or a
// number, recursing into those with a defun; plus the four vector print
// handlers added by hand. A kernel upgrade that changes the set fails here.
func TestFactorDepsMatchKernel(t *testing.T) {
	dir, err := FindKernelDir(".")
	if err != nil {
		t.Fatal(err)
	}
	files, _ := filepath.Glob(filepath.Join(dir, "*.kl"))
	sort.Strings(files)
	var src strings.Builder
	for _, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		src.Write(b)
	}
	s := src.String()
	defs := map[string]string{}
	for _, m := range regexp.MustCompile(`\(defun (\S+) `).FindAllStringSubmatchIndex(s, -1) {
		name := s[m[2]:m[3]]
		depth, end := 0, m[0]
		for i := m[0]; i < len(s); i++ {
			if s[i] == '(' {
				depth++
			} else if s[i] == ')' {
				if depth--; depth == 0 {
					end = i + 1
					break
				}
			}
		}
		if _, ok := defs[name]; !ok {
			defs[name] = s[m[0]:end]
		}
	}
	special := map[string]bool{"let": true, "if": true, "cond": true, "and": true, "or": true, "lambda": true,
		"freeze": true, "trap-error": true, "do": true, "defun": true, "type": true}
	callRE := regexp.MustCompile(`\(([^\s()"]+)`)
	skip := regexp.MustCompile(`^([A-Z]|-?[0-9])`)
	want := map[string]bool{"shen.printF": true, "shen.tuple": true, "shen.pvar": true, "shen.dictionary": true}
	seen := map[string]bool{}
	todo := []string{"shen.factor"}
	for len(todo) > 0 {
		f := todo[len(todo)-1]
		todo = todo[:len(todo)-1]
		body, ok := defs[f]
		if seen[f] || !ok {
			continue
		}
		seen[f] = true
		want[f] = true
		for _, m := range callRE.FindAllStringSubmatch(body[7:], -1) {
			tok := m[1]
			if special[tok] || skip.MatchString(tok) {
				continue
			}
			want[tok] = true
			if _, ok := defs[tok]; ok && !seen[tok] {
				todo = append(todo, tok)
			}
		}
	}
	for n := range want {
		if !factorDeps[n] {
			t.Errorf("factorDeps lacks %s, which shen.factor can reach", n)
		}
	}
	for n := range factorDeps {
		if !want[n] {
			t.Errorf("factorDeps has %s, which shen.factor cannot reach", n)
		}
	}
	if len(want) != len(factorDeps) {
		t.Errorf("%d derived, %d listed", len(want), len(factorDeps))
	}
}
