package kl

import (
	"strings"
	"testing"
)

func TestLocalThawControlFlow(t *testing.T) {
	tests := []struct {
		name, src string
		want      int
	}{
		{"nested fallback", `(let Go (freeze 9) (if false 1 (if false 2 (thaw Go))))`, 9},
		{"two tail sites", `(let Go (freeze 9) (if true (thaw Go) (thaw Go)))`, 9},
		{"two tail sites false", `(let Go (freeze 9) (if false (thaw Go) (thaw Go)))`, 9},
		{"cond branches", `(let Go (freeze 9) (cond (false 1) (true (thaw Go))))`, 9},
		{"two executed sites", `(let Go (freeze 9) (+ (thaw Go) (thaw Go)))`, 18},
		{"lazy", `(let Go (freeze (simple-error "not evaluated")) (if true 7 (thaw Go)))`, 7},
		{"capture shadowed local", `(let X 9 (let Go (freeze X) (let X 2 (thaw Go))))`, 9},
		{"capture own name", `(let Go 9 (let Go (freeze Go) (thaw Go)))`, 9},
		{"shadowed thunk", `(let Go (freeze 9) (+ (thaw Go) (let Go (freeze 2) (thaw Go))))`, 11},
		{"thaw in shadowing initializer", `(let Go (freeze 9) (let Go (thaw Go) Go))`, 9},
		{"nested thunk scopes", `(let X 9 (let Go (freeze X) (let X 2 (let Next (freeze (+ X 1)) (+ (thaw Go) (thaw Next))))))`, 12},
		{"repeated effects", `(do (set thaw-inline-counter 0) (let Go (freeze (set thaw-inline-counter (+ (value thaw-inline-counter) 1))) (do (thaw Go) (thaw Go))))`, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form, err := NewSexpReader(strings.NewReader(tt.src), true).Read()
			if err != nil {
				t.Fatal(err)
			}
			obj := CompileFunc("thaw-test", nil, form)
			fn := mustBytecodeFunc(obj).fn
			for _, instr := range fn.Code {
				if instr.Op == OP_MAKE_CLOSURE {
					t.Fatalf("unexpected closure: %v", fn.Code)
				}
			}
			var e ControlFlow
			if got := Call(&e, obj); GetInteger(got) != tt.want {
				t.Fatalf("got %s, want %d", ObjString(got), tt.want)
			}
		})
	}
}

func TestLocalThawRetainsEscapingClosures(t *testing.T) {
	for _, src := range []string{
		`(let Go (freeze 9) Go)`,
		`(let thaw true (let Go (freeze 9) (cond (thaw Go))))`,
		`(let Go (freeze 9) (cons Go ()))`,
		`(let Go (freeze 9) (do (set thaw-inline-escape Go) (thaw Go)))`,
		`(let Go (freeze 9) (lambda X (thaw Go)))`,
		`(let Go (freeze 9) (freeze (thaw Go)))`,
		`(let Go (freeze 9) (defun thaw-inline-escape () (thaw Go)))`,
		`(let Go (freeze 9) (trap-error (thaw Go) (lambda E 0)))`,
		`(let Go (freeze 9) (thaw (if true Go Go)))`,
	} {
		t.Run(src, func(t *testing.T) {
			form, err := NewSexpReader(strings.NewReader(src), true).Read()
			if err != nil {
				t.Fatal(err)
			}
			fn := mustBytecodeFunc(CompileFunc("thaw-escape", nil, form)).fn
			for _, instr := range fn.Code {
				if instr.Op == OP_MAKE_CLOSURE && mustBytecodeFunc(fn.Consts[instr.A]).fn.Arity == 0 {
					return
				}
			}
			t.Fatalf("missing freeze closure: %v", fn.Code)
		})
	}
}

func TestLocalThawCaptureAcrossLambda(t *testing.T) {
	// thaw belongs to the kernel, which these package tests do not load.
	sym := mustSymbol(MakeSymbol("thaw"))
	old := sym.function
	defer func() { sym.function = old }()
	BindSymbolFunc(MakeSymbol("thaw"), MakeNative(nativeThaw, 1))
	for _, src := range []string{
		`(let X 9 (let Go (freeze X) (let Next (freeze (thaw Go)) (let X 2 (thaw Next)))))`,
		`((lambda X (let Go (freeze X) (let X 2 (thaw Go)))) 9)`,
		`(let X 9 ((lambda Y (let Go (freeze X) (let X 2 (thaw Go)))) 0))`,
		`(let X 9 (let Go (freeze X) ((let X 2 (lambda Y (thaw Go))) 0)))`,
		`(let Go (freeze 9) (trap-error (thaw Go) (lambda E 0)))`,
	} {
		t.Run(src, func(t *testing.T) {
			form, err := NewSexpReader(strings.NewReader(src), true).Read()
			if err != nil {
				t.Fatal(err)
			}
			var e ControlFlow
			got := Call(&e, CompileFunc("thaw-capture", nil, form))
			if GetInteger(got) != 9 {
				t.Fatalf("got %s, want 9", ObjString(got))
			}
		})
	}
}
