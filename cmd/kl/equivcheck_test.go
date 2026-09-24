package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/pyrex41/shen-go/kl"
)

// TestEquivCheckReplaysTable runs `kl equiv-check` over the committed table.
// It checks that every row's verdict matches what the JSON records, that the
// output has the documented one-line-per-row shape, that the exit status
// is non-zero exactly when a row fails, and that nothing but the report is
// written: the cases exercise the runtime's error paths, which once traced
// to os.Stdout (issue #52).
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
	var rc, failed int
	stray := captureStdout(t, func() {
		rc, failed, err = runEquivCheck(path, "", &out)
	})
	if err != nil {
		t.Fatal(err)
	}
	if stray != "" {
		t.Errorf("equiv-check wrote %d bytes to os.Stdout besides its report:\n%s", len(stray), stray)
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

// captureStdout runs f with os.Stdout redirected to a pipe and returns what
// was written. The read end is drained in a goroutine so a chatty f cannot
// block on a full pipe buffer.
func captureStdout(t *testing.T, f func()) string {
	t.Helper()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	saved := os.Stdout
	os.Stdout = w
	done := make(chan string)
	go func() {
		b, _ := io.ReadAll(r)
		r.Close()
		done <- string(b)
	}()
	func() {
		defer func() {
			os.Stdout = saved
			w.Close()
		}()
		f()
	}()
	return <-done
}
