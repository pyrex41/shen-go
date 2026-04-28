package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/tiancaiamao/shen-go/kl"
)

var pprof bool

func init() {
	flag.BoolVar(&pprof, "pprof", false, "enable pprof")
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
	installLoadCache()
}

var ns2_1set kl.Obj
var try_1catch kl.Obj

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

	var e kl.ControlFlow
	regist(&e)
	kl.Eval(&e, kl.Cons(kl.MakeSymbol("shen.initialise"), kl.Nil))
	repl(&e)
}
