package main

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"testing"
	"time"
)

// Mutual tail calls through every tail position, 5M deep, on the VM: the result
// is right and the footprint stays flat. cmd/kl's TestAOTLetAndSelfTailSemantics
// is the same gate for code compiled to Go. A lost tail call can pass every
// certification suite, whose recursions are shallow.
func TestDeepMutualTailCalls(t *testing.T) {
	const defs = `
(define ev? 0 -> true N -> (and (number? N) (od? (- N 1))))
(define od? 0 -> false N -> (and (number? N) (ev? (- N 1))))
(define ev-or? 0 -> true N -> (or (not (number? N)) (od-or? (- N 1))))
(define od-or? 0 -> false N -> (or (not (number? N)) (ev-or? (- N 1))))
(define ev-if? N -> (if (= N 0) true (od-if? (- N 1))))
(define od-if? N -> (if (= N 0) false (ev-if? (- N 1))))
(define ev-else? N -> (if (not (= N 0)) (od-else? (- N 1)) true))
(define od-else? N -> (if (not (= N 0)) (ev-else? (- N 1)) false))
(define ev-cond? N -> (cond ((= N 0) true) ((= N 1) false) (true (od-cond? (- N 1)))))
(define od-cond? N -> (cond ((= N 0) false) ((= N 1) true) (true (ev-cond? (- N 1)))))
(define ev-let? 0 -> true N -> (let M (- N 1) (od-let? M)))
(define od-let? 0 -> false N -> (let M (- N 1) (ev-let? M)))
(define ev-do? 0 -> true N -> (do (number? N) (od-do? (- N 1))))
(define od-do? 0 -> false N -> (do (number? N) (ev-do? (- N 1))))
(define ev-var? _ 0 -> true F N -> (F (fn ev-var?) (- N 1)))
(define od-var? _ 0 -> false F N -> (F (fn od-var?) (- N 1)))
(define ev-thaw? 0 -> true N -> (thaw (freeze (od-thaw? (- N 1)))))
(define od-thaw? 0 -> false N -> (thaw (freeze (ev-thaw? (- N 1)))))
`
	bin := buildShen(t)
	dir := t.TempDir()
	for _, c := range []string{
		"(ev? 5000000)", "(ev-or? 5000000)", "(ev-if? 5000000)", "(ev-else? 5000000)", "(ev-cond? 5000000)",
		"(ev-let? 5000000)", "(ev-do? 5000000)", "(ev-var? (fn od-var?) 5000000)", "(ev-thaw? 5000000)",
	} {
		path := filepath.Join(dir, "t.shen")
		src := defs + `(pr (make-string "RESULT ~A~%" ` + c + `) (stoutput))` + "\n"
		if err := os.WriteFile(path, []byte(src), 0o644); err != nil {
			t.Fatal(err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
		cmd := exec.CommandContext(ctx, bin, "script", path)
		out, err := cmd.CombinedOutput()
		cancel()
		if err != nil {
			t.Fatalf("%s: %v\n%.2000s", c, err, out)
		}
		if !strings.Contains(string(out), "RESULT true") {
			t.Errorf("%s: output %.500s", c, out)
		}
		if ru, ok := cmd.ProcessState.SysUsage().(*syscall.Rusage); ok {
			rss := int64(ru.Maxrss)
			if runtime.GOOS != "darwin" {
				rss *= 1024
			}
			// A frame per call would be gigabytes. Garbage from var and thaw is
			// bounded by the collector, at about 80MB here.
			if rss > 300<<20 {
				t.Errorf("%s: peak RSS %dMB", c, rss>>20)
			}
		}
	}
}
