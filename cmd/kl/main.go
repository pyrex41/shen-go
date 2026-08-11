package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/pyrex41/shen-go/kl"
)

var pprof bool
var sha256File string

func init() {
	flag.BoolVar(&pprof, "pprof", false, "enable pprof")
	flag.StringVar(&sha256File, "sha256", "", "print the sha256 of FILE and exit (used by `make precompile` to stamp a plugin's source provenance)")
}

func main() {
	flag.Parse()

	// -sha256 is a tiny utility mode, not part of the KL interpreter: it
	// exists so `make precompile` can stamp the .fns manifest with the source
	// hash using exactly the digest cmd/shen checks it against, without
	// depending on a shasum/sha256sum being present and identically spelled on
	// every platform that can build this repo.
	if sha256File != "" {
		data, err := os.ReadFile(sha256File)
		if err != nil {
			fmt.Fprintln(os.Stderr, "sha256:", err)
			os.Exit(1)
		}
		fmt.Printf("%x\n", sha256.Sum256(data))
		return
	}

	if pprof {
		go http.ListenAndServe(":8080", nil)
	}

	var ctl kl.ControlFlow
	kl.BindSymbolFunc(kl.MakeSymbol("bc->go"), bcToGo)
	kl.BindSymbolFunc(kl.MakeSymbol("make-code-generator"), makeCodeGenerator)
	kl.BindSymbolFunc(kl.MakeSymbol("go-build-and-load"), goBuildAndLoad)
	kl.BindSymbolFunc(kl.MakeSymbol("go-build-plugin"), goBuildPlugin)
	kl.BindSymbolFunc(kl.MakeSymbol("emit-arities"), emitArities)

	r := kl.NewSexpReader(os.Stdin, false)
	for i := 0; ; i++ {
		fmt.Printf("%d #> ", i)
		sexp, err := r.Read()
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}
		res := kl.Eval(&ctl, sexp)
		fmt.Println(kl.ObjString(res))
	}
}
