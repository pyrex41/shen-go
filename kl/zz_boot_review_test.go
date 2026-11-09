package kl

import "testing"

func TestZZBootKernelFactorState(t *testing.T) {
	dir, err := FindKernelDir(".")
	if err != nil {
		t.Skip(err)
	}
	var e ControlFlow
	if err := BootKernel(&e, dir); err != nil {
		t.Fatal(err)
	}
	t.Logf("after BootKernel: armed=%v broken=%v stdFactor=%v", factorArmed, factorBroken, stdFactor != nil)
	InstallKernelFast()
	t.Logf("after 2nd InstallKernelFast: armed=%v broken=%v", factorArmed, factorBroken)
	factorBroken = false
	BindSymbolFunc(MakeSymbol("shen.factor-recognisors"), MakeNative1(func(_ *ControlFlow, a Obj) Obj { return a }))
	t.Logf("after embedder rebinds shen.factor-recognisors: broken=%v", factorBroken)
	factorBroken = false
	BindSymbolFunc(MakeSymbol("shen.tuple"), MakeNative1(func(_ *ControlFlow, a Obj) Obj { return a }))
	t.Logf("after rebinding shen.tuple: broken=%v", factorBroken)
}
