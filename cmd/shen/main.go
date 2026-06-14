package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"

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
		LauncherMain,
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

// splitArgs separates the leading Go-level flags (-pprof, -precompiled) from
// the launcher command line. Everything from the first argument that isn't one
// of our Go flags onward is handed verbatim to the kernel launcher, so
// launcher-protocol arguments like `--version`, `--help` or `eval -e ...`
// never reach Go's flag parser (which would reject them as unknown flags).
func splitArgs(args []string) (goFlags, launcherArgs []string) {
	i := 0
	for i < len(args) {
		arg := args[i]
		name := strings.TrimLeft(arg, "-")
		switch {
		case len(arg) > len(name) && (name == "pprof" || name == "precompiled" ||
			strings.HasPrefix(name, "pprof=") || strings.HasPrefix(name, "precompiled=")):
			goFlags = append(goFlags, arg)
			// -precompiled takes its value as the following argument.
			if name == "precompiled" && i+1 < len(args) {
				i++
				goFlags = append(goFlags, args[i])
			}
			i++
		default:
			return goFlags, args[i:]
		}
	}
	return goFlags, nil
}

func main() {
	goFlags, launcherArgs := splitArgs(os.Args[1:])
	flag.CommandLine.Parse(goFlags)

	if pprof {
		go http.ListenAndServe(":8080", nil)
	}

	ns2_1set = kl.PrimFunc(kl.MakeSymbol("defun"))
	try_1catch = kl.PrimFunc(kl.MakeSymbol("try-catch"))

	kl.BindSymbolFunc(kl.MakeSymbol("load-native"), loadNative)

	var e kl.ControlFlow
	regist(&e)
	// Memoize read-file (tokenize/parse/lower) per unchanged .shen file, so a
	// repeated (load ...) in the same process skips re-parsing. Must come after
	// regist (which binds the kernel read-file we wrap) and before any load.
	installLoadCache()
	// Override the kernel's interpreted `hash` (sys.kl) with the native FNV-1a
	// version. Must come AFTER regist (which binds the kernel's hash) and BEFORE
	// shen.initialise (which builds dictionaries): every dict must be created and
	// queried with the same hash, so the swap has to precede the first dict op.
	kl.BindSymbolFunc(kl.MakeSymbol("hash"), kl.MakePrimitive("hash", 2, kl.PrimHash))
	kl.Eval(&e, kl.Cons(kl.MakeSymbol("shen.initialise"), kl.Nil))
	fixPrHush(&e)
	loadPrecompiled(&e, precompiled)
	if len(launcherArgs) > 0 {
		runLauncher(&e, launcherArgs)
		return
	}
	repl(&e)
}

// fixPrHush rebinds the kernel's `pr` so the *hush* check only applies to the
// char-stoutput fast path. The stock definition (writer.kl) returns without
// writing whenever *hush* is true — for EVERY stream, so quiet mode
// (`shen eval -q ...`, which sets *hush*) silently truncates all file output
// written via pr/open (e.g. ratatoskr stage-1 emits its shaken kernel.kl that
// way). Since shen-go's PrimCharStOutput is stubbed to false, pr here
// effectively ignores *hush* altogether — the same behavior as shen-cl, whose
// native pr writes unconditionally. With *hush* false (the default: REPL and
// certification) behavior is identical to stock.
func fixPrHush(e *kl.ControlFlow) {
	const def = `(defun pr (V S)
	(if (shen.char-stoutput? S)
		(if (value *hush*) V (shen.write-string V S))
		(shen.write-chars V S (shen.string->byte V 0) 1)))`
	exp, err := kl.NewSexpReader(strings.NewReader(def), false).Read()
	if err != nil {
		panic(err)
	}
	kl.Eval(e, exp)
}

// runLauncher drives the kernel's launcher extension
// (kernel/klambda/extension-launcher.kl, compiled into launcher.go) so the CLI
// speaks the standard Shen launcher protocol shared with shen-cl and
// ShenScript: `shen repl`, `shen script FILE ARGS...`, and
// `shen eval [-e EXPR] [-l FILE] [-q] [-s KEY VALUE] [-r]`, plus --help and
// --version. It is only entered when non-flag arguments remain after Go flag
// parsing, so bare `shen` (and `shen -pprof`/`-precompiled ...`) still goes
// straight to the REPL.
//
// Result handling mirrors shen.x.launcher.default-handle-result, with one
// substitution: a launch-repl result runs the Go repl loop above (which exits
// cleanly on stdin EOF) instead of the kernel's shen.repl, which spins forever
// on a closed stdin.
func runLauncher(e *kl.ControlFlow, args []string) {
	argv := kl.Nil
	for i := len(args) - 1; i >= 0; i-- {
		argv = kl.Cons(kl.MakeString(args[i]), argv)
	}
	argv = kl.Cons(kl.MakeString(os.Args[0]), argv)

	// An uncaught Shen exception (e.g. `-l` on a missing file) panics through
	// kl.Call, so run the launcher under kl.Try and turn any exception into the
	// launcher protocol's own (error Msg) result.
	body := kl.MakeNative(func(e *kl.ControlFlow) {
		e.TailApply(kl.PrimFunc(kl.MakeSymbol("shen.x.launcher.launch-shen")), argv)
	}, 0)
	handler := kl.MakeNative(func(e *kl.ControlFlow) {
		err := e.Get(1)
		e.Return(kl.Cons(kl.MakeSymbol("error"), kl.Cons(kl.PrimErrorToString(err), kl.Nil)))
	}, 1)
	res := kl.Try(e, body).Catch(handler)
	tag := kl.Car(res)
	switch tag {
	case kl.MakeSymbol("success"):
		if rest := kl.Cdr(res); rest != kl.Nil {
			fmt.Println(kl.GetString(kl.Car(rest)))
		}
	case kl.MakeSymbol("launch-repl"):
		repl(e)
	case kl.MakeSymbol("show-help"):
		fmt.Println(kl.GetString(kl.Car(kl.Cdr(res))))
	case kl.MakeSymbol("error"):
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", kl.GetString(kl.Car(kl.Cdr(res))))
		os.Exit(1)
	case kl.MakeSymbol("unknown-arguments"):
		prog := kl.GetString(kl.Car(kl.Cdr(res)))
		bad := kl.GetString(kl.Car(kl.Cdr(kl.Cdr(res))))
		fmt.Fprintf(os.Stderr, "ERROR: Invalid argument: %s\nTry `%s --help' for more information.\n", bad, prog)
		os.Exit(1)
	default:
		fmt.Fprintf(os.Stderr, "ERROR: unexpected launcher result: %s\n", kl.ObjString(res))
		os.Exit(1)
	}
}
