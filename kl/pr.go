package kl

// primPr is pr's byte-stream path. Range decodes code points just like
// pos/string->n, but avoids repeatedly converting the whole string to runes.
// Keep PrimWriteByte for its exact sink, byte validation and error behavior:
// in particular U+00FF writes FF, and U+0100 fails after any preceding bytes.
func primPr(value, stream Obj) Obj {
	// shen.string->byte traps a failed pos on non-strings as end-of-string.
	if !IsString(value) {
		return value
	}
	for _, r := range GetString(value) {
		PrimWriteByte(MakeInteger(int(r)), stream)
	}
	return value
}

// InstallPr installs only the byte-writing path; the launcher retains the
// Shen wrapper that checks char-stoutput? and *hush* before selecting it.
func InstallPr() {
	BindSymbolFunc(MakeSymbol("shen.native-pr"), MakePrimitive("shen.native-pr", 2, primPr))
}
