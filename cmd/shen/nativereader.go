package main

// Go-native read-file front end.
//
// The kernel's interpreted read-file (kernel/sources/reader.shen) is:
//
//	File -> (let Bytelist (read-file-as-bytelist File)
//	             S-exprs  (trap-error (compile (/. X (<s-exprs> X)) Bytelist)
//	                                  (/. E (reader-error (value *residue*))))
//	             Process  (process-sexprs S-exprs)
//	          Process)
//
// The (compile (/. X (<s-exprs> X)) Bytelist) middle stage — tokenizing and
// parsing the raw bytes into s-expr forms — is the expensive, fully interpreted
// half: thousands of trampoline iterations per identifier. installNativeReader
// replaces read-file with a native that does that stage in Go (kl.ShenReadSExprs)
// and then hands the resulting raw forms to the *interpreted* process-sexprs, so
// every observable side effect of reading (arity registration, macroexpansion,
// application currying, and the canonical reader-error on malformed input) is
// preserved exactly.
//
// Safety: kl.ShenReadSExprs answers only when it is certain it matches the
// interpreted <s-exprs> byte-for-byte; on any construct it does not handle, or
// on malformed/leftover input, it declines and we delegate the *whole* read-file
// to the captured interpreted definition, which produces the canonical result or
// error. Correctness therefore never depends on the Go reader being complete.
//
// Composition with the load cache: installNativeReader must run BEFORE
// installLoadCache so the cache wraps this native read-file (capturing it as the
// miss-path "orig"). A cold load then parses in Go; a cached reload skips even
// that. See main().

import (
	"os"

	"github.com/tiancaiamao/shen-go/kl"
)

// installNativeReader rebinds read-file to a native that parses with
// kl.ShenReadSExprs and lowers via the interpreted process-sexprs, falling back
// to the captured interpreted read-file on anything the Go reader declines.
// Must run after regist (which binds the interpreted read-file and
// shen.process-sexprs) and before installLoadCache.
func installNativeReader() {
	readFileSym := kl.MakeSymbol("read-file")
	interpreted := kl.PrimFunc(readFileSym) // canonical fallback
	processSexprs := kl.PrimFunc(kl.MakeSymbol("shen.process-sexprs"))

	native := kl.MakeNative(func(e *kl.ControlFlow) {
		arg := e.Get(1)
		path := kl.GetString(arg)

		data, err := os.ReadFile(path)
		if err != nil {
			// Missing/unreadable: let the interpreted read-file raise the
			// kernel's canonical error (it opens the stream the same way).
			e.TailApply(interpreted, arg)
			return
		}

		forms, perr := kl.ShenReadSExprs(data)
		if perr != nil {
			// The Go reader declined (unhandled construct or malformed input):
			// delegate the whole read-file so the kernel produces the canonical
			// forms or reader-error.
			e.TailApply(interpreted, arg)
			return
		}

		// process-sexprs runs unguarded in the kernel's read-file (only the
		// parse is under trap-error), so mirror that: any error it raises
		// propagates to the caller's own handler.
		e.TailApply(processSexprs, forms)
	}, 1)

	kl.BindSymbolFunc(readFileSym, native)
}
