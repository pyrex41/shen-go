package kl

import "testing"

func TestShenXFeaturesCurrentPrimitive(t *testing.T) {
	InstallShenXFeatures()
	var ctl ControlFlow
	got := Call(&ctl, PrimFunc(MakeSymbol("shen.x.features.current")))
	if !Equal(got, currentShenXFeatures()) {
		t.Fatalf("feature primitive returned %s, registry has %s", ObjString(got), ObjString(currentShenXFeatures()))
	}
	for _, feature := range ListToSlice(got) {
		if *feature != scmHeadSymbol {
			t.Fatalf("feature list contains non-symbol %s", ObjString(feature))
		}
	}
}
