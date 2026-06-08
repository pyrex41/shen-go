package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/tiancaiamao/shen-go/kl"
)

var pprof bool
var precompiled precompiledFlag

func init() {
	flag.BoolVar(&pprof, "pprof", false, "enable pprof")
	flag.Var(&precompiled, "precompiled", "path to a precompiled .so (repeatable or comma-separated); functions in it run as compiled Go instead of the VM")
}

func regist(e *kl.ControlFlow) {
	for _, init := range []kl.Obj{
		TopLevelMain,
		CoreMain,
		SysMain,
		SequentMain,
		YaccMain,
		ReaderMain,
		PrologMain,
		TrackMain,
		LoadMain,
		WriterMain,
		MacrosMain,
		DeclarationsMain,
		TStarMain,
		TypesMain,
		DictMain,
		InitMain,
	} {
		res := kl.Call(e, init)
		if kl.IsError(res) {
			fmt.Println("load ...fail")
		}
	}
}

var ns2_1set kl.Obj
var try_1catch kl.Obj

// repl runs the Shen top-level loop in Go so it can exit cleanly when stdin
// reaches EOF (e.g. piped input). The Shen-level shen.repl/shen.loop tail-calls
// itself forever; on EOF lineread raises "error: empty stream", which the old
// loop caught and ignored, spinning indefinitely. Here we detect that EOF via
// the StinputEOF flag (set in kl.PrimReadByte) and return instead.
func repl(e *kl.ControlFlow) {
	kl.Call(e, kl.PrimFunc(symshen_4credits))
	for {
		kl.Call(e, kl.PrimFunc(symshen_4initialise__environment))
		kl.Call(e, kl.PrimFunc(symshen_4prompt))
		kl.ResetStinputEOF()

		body := kl.MakeNative(func(e *kl.ControlFlow) {
			e.TailApply(kl.PrimFunc(symshen_4read_1evaluate_1print))
		}, 0)
		handler := kl.MakeNative(func(e *kl.ControlFlow) {
			err := e.Get(1)
			if isStinputEOFError(err) {
				e.Return(kl.Nil)
				return
			}
			e.TailApply(kl.PrimFunc(symshen_4toplevel_1display_1exception), err)
		}, 1)

		kl.Try(e, body).Catch(handler)
		if kl.StinputEOF() {
			return
		}
	}
}

// isStinputEOFError reports whether err is the "empty stream" error raised by
// lineread when *stinput* hit EOF (as opposed to a genuine user error).
func isStinputEOFError(err kl.Obj) bool {
	return kl.StinputEOF() && kl.GetString(kl.PrimErrorToString(err)) == "error: empty stream"
}

func main() {
	flag.Parse()

	if pprof {
		go http.ListenAndServe(":8080", nil)
	}

	ns2_1set = kl.PrimFunc(kl.MakeSymbol("defun"))
	try_1catch = kl.PrimFunc(kl.MakeSymbol("try-catch"))

	kl.BindSymbolFunc(kl.MakeSymbol("load-native"), loadNative)

	var e kl.ControlFlow
	regist(&e)
	// Override the kernel's interpreted `hash` (sys.kl) with the native FNV-1a
	// version. Must come AFTER regist (which binds the kernel's hash) and BEFORE
	// shen.initialise (which builds dictionaries): every dict must be created and
	// queried with the same hash, so the swap has to precede the first dict op.
	kl.BindSymbolFunc(kl.MakeSymbol("hash"), kl.MakePrimitive("hash", 2, kl.PrimHash))
	kl.Eval(&e, kl.Cons(kl.MakeSymbol("shen.initialise"), kl.Nil))
	loadPrecompiled(&e, precompiled)
	repl(&e)
}
