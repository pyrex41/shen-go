package kl

import "testing"

// Booting the kernel's natives twice leaves factorising on: the second
// InstallKernelFast rebinds thaw and other factoriser dependencies, and those
// are its own, not a user's rebinding.
func TestInstallKernelFastTwiceKeepsFactorising(t *testing.T) {
	lowerMu.Lock()
	savedBroken, savedArmed := factorBroken, factorArmed
	factorBroken = false
	lowerMu.Unlock()
	defer func() {
		lowerMu.Lock()
		factorBroken, factorArmed = savedBroken, savedArmed
		lowerMu.Unlock()
	}()
	InstallKernelFast()
	InstallKernelFast()
	lowerMu.Lock()
	broken, armed := factorBroken, factorArmed
	lowerMu.Unlock()
	if broken || !armed {
		t.Fatalf("after two boots: factorBroken=%v factorArmed=%v, want false true", broken, armed)
	}
}
