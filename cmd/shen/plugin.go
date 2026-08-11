package main

// Load side of the compile-to-Go path: open a .so produced by `make precompile`
// and run its exported Install hook, which binds every function in the file
// (rebinding those symbols' .function to plugin-compiled natives). Everything
// not in a loaded .so stays on the bytecode VM — including anything defined at
// the REPL. This is the only codegen-related code the production binary needs;
// the compiler toolchain lives in cmd/kl.

import (
	"bufio"
	"fmt"
	"os"
	"plugin"
	"strconv"
	"strings"

	"github.com/pyrex41/shen-go/kl"
)

// loadPlugin opens soPath, runs its exported Install hook (which binds every
// function in the file), then registers each function's arity at the Shen level
// from the companion "<soPath>.fns" manifest. The arity step is required: binding
// the KL function alone leaves it "undefined" to the Shen reader, which resolves
// applications via registered arity (see shen.record-and-evaluate in core.kl).
//
// The plugin must be built from the same module + Go toolchain as this binary
// (see `make precompile`); a mismatch surfaces here as a plugin.Open error.
func loadPlugin(e *kl.ControlFlow, soPath string) error {
	p, err := plugin.Open(soPath)
	if err != nil {
		return fmt.Errorf("plugin.Open(%s): %w", soPath, err)
	}
	sym, err := p.Lookup("Install")
	if err != nil {
		return fmt.Errorf("lookup Install in %s: %w", soPath, err)
	}
	install, ok := sym.(func(*kl.ControlFlow))
	if !ok {
		return fmt.Errorf("%s: Install has unexpected type %T", soPath, sym)
	}
	install(e)
	registerArities(e, soPath+".fns")
	return nil
}

// registerArities replays the "NAME ARITY" manifest via shen.store-arity so the
// plugin's functions are callable at the Shen REPL. A missing manifest is not
// fatal — the functions are still callable from compiled/KL code.
func registerArities(e *kl.ControlFlow, fnsPath string) {
	f, err := os.Open(fnsPath)
	if err != nil {
		return
	}
	defer f.Close()
	storeArity := kl.PrimFunc(kl.MakeSymbol("shen.store-arity"))
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		fields := strings.Fields(sc.Text())
		if len(fields) != 2 {
			continue
		}
		arity, err := strconv.Atoi(fields[1])
		if err != nil {
			continue
		}
		kl.Call(e, storeArity, kl.MakeSymbol(fields[0]), kl.MakeInteger(arity))
	}
}

// registerLoadNativeArity publishes load-native's arity to the Shen level, the
// same way registerArities does for the functions a plugin installs. Binding the
// KL function is necessary but not sufficient: shen.record-and-evaluate resolves
// an application through the registered arity, so a KL-bound-but-arity-less
// symbol is reported as undefined the moment it is called from Shen source.
//
// Must be called after regist (which binds shen.store-arity).
func registerLoadNativeArity(e *kl.ControlFlow) {
	storeArity := kl.PrimFunc(kl.MakeSymbol("shen.store-arity"))
	kl.Call(e, storeArity, kl.MakeSymbol("load-native"), kl.MakeInteger(1))
}

// loadNative is the REPL/file-callable form: (load-native "foo.so").
var loadNative = kl.MakeNative(func(e *kl.ControlFlow) {
	soPath := kl.GetString(e.Get(1))
	if err := loadPlugin(e, soPath); err != nil {
		e.Return(kl.MakeError(err.Error()))
		return
	}
	e.Return(kl.Nil)
}, 1)

// precompiledFlag is a repeatable -precompiled flag (also accepts comma-separated).
type precompiledFlag []string

func (p *precompiledFlag) String() string { return fmt.Sprint([]string(*p)) }
func (p *precompiledFlag) Set(v string) error {
	for _, s := range splitComma(v) {
		if s != "" {
			*p = append(*p, s)
		}
	}
	return nil
}

func splitComma(s string) []string {
	var out []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ',' {
			out = append(out, s[start:i])
			start = i + 1
		}
	}
	return append(out, s[start:])
}

// loadPrecompiled loads each .so given via -precompiled, after shen.initialise
// and before the REPL. A failure warns and continues on the VM (never aborts).
func loadPrecompiled(e *kl.ControlFlow, paths []string) {
	for _, soPath := range paths {
		if err := loadPlugin(e, soPath); err != nil {
			fmt.Fprintf(os.Stderr, "warning: -precompiled %s: %v (continuing on the VM)\n", soPath, err)
		}
	}
}
