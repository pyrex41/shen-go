package kl

import (
	"crypto/sha256"
	"os"
)

// InstallShenX binds optional shen.x host extensions.
// SHA-256: crypto/sha256 (stdlib). Disable with SHEN_X_SHA256=pure.
func InstallShenX() {
	installShenXZmq()
	if os.Getenv("SHEN_X_SHA256") == "pure" {
		return
	}
	BindSymbolFunc(MakeSymbol("shen.x.sha256-octets-host"),
		MakePrimitive("shen.x.sha256-octets-host", 1, primSha256OctetsHost))
	PrimSet(MakeSymbol("shen.x.*sha256-backend*"), MakeSymbol("host"))
}

func primSha256OctetsHost(xs Obj) Obj {
	var buf []byte
	for xs != Nil {
		if *xs != scmHeadPair {
			panic(MakeError("shen.x.sha256-octets-host: expected list of bytes 0..255"))
		}
		p := mustPair(xs)
		if !IsNumber(p.car) {
			panic(MakeError("shen.x.sha256-octets-host: expected list of bytes 0..255"))
		}
		n := GetInteger(p.car)
		if n < 0 || n > 255 {
			panic(MakeError("shen.x.sha256-octets-host: expected list of bytes 0..255"))
		}
		buf = append(buf, byte(n))
		xs = p.cdr
	}
	sum := sha256.Sum256(buf)
	out := Nil
	for i := 31; i >= 0; i-- {
		out = cons(MakeInteger(int(sum[i])), out)
	}
	return out
}
