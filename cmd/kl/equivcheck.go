package main

// `kl equiv-check kl/equiv.json`: re-run the differential cases recorded in
// the equivalence table (see kl/equiv.go) against the natives InstallKernelFast
// binds, on this host, and print one line per row:
//
//	equiv NAME ok cases=N
//	equiv NAME FAIL case=... kl=... native=...
//
// followed by one summary line, "equiv-check: PORT KERNEL: R rows, C cases,
// F FAIL, DURATION". Nothing else reaches stdout: the runtime's own traces on
// recovered errors are discarded while the cases run. Exit status is 1 when
// any row fails, 2 when the table or the kernel could not be loaded. This is
// the harness Yggdrasil's conformance report invokes to check the port's
// declared lowering table.

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/pyrex41/shen-go/kl"
)

func equivCheckMain(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("kl equiv-check", flag.ContinueOnError)
	fs.SetOutput(stderr)
	kernelDir := fs.String("kernel-dir", "", "kernel/klambda directory (default: found by walking up from the current directory)")
	fs.Usage = func() {
		fmt.Fprintf(stderr, "usage: kl equiv-check [-kernel-dir DIR] path/to/equiv.json\n")
		fs.PrintDefaults()
	}
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if fs.NArg() != 1 {
		fs.Usage()
		return 2
	}
	rc, _, err := runEquivCheck(fs.Arg(0), *kernelDir, stdout)
	if err != nil {
		fmt.Fprintln(stderr, "equiv-check:", err)
		return 2
	}
	return rc
}

// runEquivCheck boots the kernel, replays every row's cases and reports to
// out. It returns the exit status (0 ok, 1 any FAIL) and the number of
// failing rows.
func runEquivCheck(jsonPath, kernelDir string, out io.Writer) (int, int, error) {
	start := time.Now()
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return 2, 0, err
	}
	var table kl.EquivTable
	if err := json.Unmarshal(data, &table); err != nil {
		return 2, 0, fmt.Errorf("%s: %v", jsonPath, err)
	}
	if kernelDir == "" {
		kernelDir, err = kl.FindKernelDir(".")
		if err != nil {
			return 2, 0, err
		}
	}
	names := make([]string, 0, len(table.Rows))
	for _, r := range table.Rows {
		names = append(names, r.KernelFn)
	}
	var e kl.ControlFlow
	failed, cases := 0, 0
	// Booting the kernel and running the cases both hit runtime error paths
	// that trace to os.Stdout; out was captured by the caller before this
	// point, so the report itself is unaffected by the redirect.
	var bootErr error
	quietErr := kl.WithQuietStdout(func() {
		if bootErr = kl.BootKernel(&e, kernelDir); bootErr != nil {
			return
		}
		if _, bootErr = kl.InstallEquivDefinitions(&e, kernelDir, names); bootErr != nil {
			return
		}
		for _, r := range table.Rows {
			cases += len(r.Inputs)
			var first *kl.EquivMismatch
			for _, c := range r.Inputs {
				if first = kl.RunEquivCase(&e, r.KernelFn, c); first != nil {
					break
				}
			}
			if first == nil {
				fmt.Fprintf(out, "equiv %s ok cases=%d\n", r.KernelFn, len(r.Inputs))
			} else {
				failed++
				fmt.Fprintf(out, "equiv %s FAIL %s\n", r.KernelFn, first)
			}
		}
	})
	if quietErr != nil {
		return 2, 0, quietErr
	}
	if bootErr != nil {
		return 2, 0, bootErr
	}
	fmt.Fprintf(out, "equiv-check: %s %s: %d rows, %d cases, %d FAIL, %s\n",
		table.Port, table.Kernel, len(table.Rows), cases, failed, time.Since(start).Round(time.Millisecond))
	if failed > 0 {
		return 1, failed, nil
	}
	return 0, 0, nil
}
