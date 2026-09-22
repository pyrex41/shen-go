package main

import (
	"bytes"
	"encoding/json"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/pyrex41/shen-go/kl"
)

// TestEquivCheckReplaysTable runs `kl equiv-check` over the committed table.
// It checks that every row's verdict matches what the JSON records, that the
// output has the documented one-line-per-row shape, and that the exit status
// is non-zero exactly when a row fails.
func TestEquivCheckReplaysTable(t *testing.T) {
	const path = "../../kl/equiv.json"
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var table kl.EquivTable
	if err := json.Unmarshal(data, &table); err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	rc, failed, err := runEquivCheck(path, "", &out)
	if err != nil {
		t.Fatal(err)
	}
	wantFailed := 0
	for _, r := range table.Rows {
		if !r.Verified {
			wantFailed++
		}
	}
	if failed != wantFailed {
		t.Errorf("equiv-check reports %d failing rows, JSON records %d unverified\n%s", failed, wantFailed, out.String())
	}
	if (rc != 0) != (wantFailed > 0) {
		t.Errorf("exit status %d with %d unverified rows", rc, wantFailed)
	}
	okLine := regexp.MustCompile(`^equiv (\S+) ok cases=(\d+)$`)
	failLine := regexp.MustCompile(`^equiv (\S+) FAIL case=\S+ kl=.* native=.*$`)
	verified := map[string]bool{}
	for _, r := range table.Rows {
		verified[r.KernelFn] = r.Verified
	}
	lines := strings.Split(strings.TrimRight(out.String(), "\n"), "\n")
	if len(lines) != len(table.Rows)+1 {
		t.Fatalf("got %d output lines for %d rows:\n%s", len(lines), len(table.Rows), out.String())
	}
	for _, line := range lines[:len(lines)-1] {
		if m := okLine.FindStringSubmatch(line); m != nil {
			if !verified[m[1]] {
				t.Errorf("%q: JSON has this row unverified", line)
			}
			continue
		}
		if m := failLine.FindStringSubmatch(line); m != nil {
			if verified[m[1]] {
				t.Errorf("%q: JSON has this row verified", line)
			}
			continue
		}
		t.Errorf("malformed line %q", line)
	}
	if !strings.HasPrefix(lines[len(lines)-1], "equiv-check: shen-go S42: ") {
		t.Errorf("missing summary line, got %q", lines[len(lines)-1])
	}
}
