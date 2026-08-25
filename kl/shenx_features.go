package kl

import "sync"

// shen.x.features.current is the port hook used by Shen Batteries' module
// loader.  Keep the registry private: feature names are a Shen-level contract,
// while registration is done by the host extensions as they are installed.
var shenXFeatures = struct {
	sync.RWMutex
	names []string
}{}

func registerShenXFeature(name string) {
	shenXFeatures.Lock()
	defer shenXFeatures.Unlock()
	for _, existing := range shenXFeatures.names {
		if existing == name {
			return
		}
	}
	shenXFeatures.names = append(shenXFeatures.names, name)
}

func currentShenXFeatures() Obj {
	shenXFeatures.RLock()
	names := append([]string(nil), shenXFeatures.names...)
	shenXFeatures.RUnlock()
	result := Nil
	for i := len(names) - 1; i >= 0; i-- {
		result = cons(MakeSymbol(names[i]), result)
	}
	return result
}

// InstallShenXFeatures binds the stable feature-discovery primitive expected
// by shen-batteries/library.shen. It is safe to call repeatedly.
func InstallShenXFeatures() {
	BindSymbolFunc(MakeSymbol("shen.x.features.current"),
		MakeNative(func(e *ControlFlow) { e.Return(currentShenXFeatures()) }, 0))
}
