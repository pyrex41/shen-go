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

func TestRegisterShenXFeatureMembership(t *testing.T) {
	const name = "shen.x/test-feature-membership"
	registerShenXFeature(name)
	registerShenXFeature(name)
	if n := countShenXFeature(name); n != 1 {
		t.Fatalf("want %s registered once, found %d in %s", name, n, ObjString(currentShenXFeatures()))
	}
}

func TestInstallShenXSHA256PureOmitsFeature(t *testing.T) {
	const sha256Host = "shen.x/sha256-host"
	if countShenXFeature(sha256Host) != 0 {
		t.Skip("sha256-host already registered in this process")
	}
	t.Setenv("SHEN_X_SHA256", "pure")
	InstallShenX()
	if countShenXFeature(sha256Host) != 0 {
		t.Fatalf("SHEN_X_SHA256=pure must not advertise %s, got %s", sha256Host, ObjString(currentShenXFeatures()))
	}
	if countShenXFeature("shen.x/zmq-host") != 1 {
		t.Fatalf("expected shen.x/zmq-host after InstallShenX, got %s", ObjString(currentShenXFeatures()))
	}
}

func countShenXFeature(name string) int {
	n := 0
	for _, feature := range ListToSlice(currentShenXFeatures()) {
		if GetSymbol(feature) == name {
			n++
		}
	}
	return n
}
