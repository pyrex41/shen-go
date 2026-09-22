package main

// `kl equiv-check kl/equiv.json`: re-run the differential cases recorded in
// the equivalence table (see kl/equiv.go) against the natives InstallKernelFast
// binds, on this host, and print one line per row:
//
//	equiv NAME ok cases=N
//	equiv NAME FAIL case=... kl=... native=...
//
// Exit status is non-zero when any row fails. This is the harness Yggdrasil's
// conformance report invokes to check the port's declared lowering table.

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

// runEquivCheck boots the kernel, replays every row's cases and reports. It
// returns the exit status (0 ok, 1 any FAIL) and the number of failing rows.
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
	var e kl.ControlFlow
	if err := kl.BootKernel(&e, kernelDir); err != nil {
		return 2, 0, err
	}
	names := make([]string, 0, len(table.Rows))
	for _, r := range table.Rows {
		names = append(names, r.KernelFn)
	}
	if err := kl.InstallEquivDefinitions(&e, kernelDir, names); err != nil {
		return 2, 0, err
	}
	failed, cases := 0, 0
	for _, r := range table.Rows {
		cases += len(r.Inputs)
		var first string
		for _, c := range r.Inputs {
			if msg := kl.RunEquivCase(&e, r.KernelFn, c); msg != "" {
				first = msg
				break
			}
		}
		if first == "" {
			fmt.Fprintf(out, "equiv %s ok cases=%d\n", r.KernelFn, len(r.Inputs))
		} else {
			failed++
			fmt.Fprintf(out, "equiv %s FAIL %s\n", r.KernelFn, first)
		}
	}
	fmt.Fprintf(out, "equiv-check: %s %s: %d rows, %d cases, %d FAIL, %s\n",
		table.Port, table.Kernel, len(table.Rows), cases, failed, time.Since(start).Round(time.Millisecond))
	if failed > 0 {
		return 1, failed, nil
	}
	return 0, 0, nil
}
