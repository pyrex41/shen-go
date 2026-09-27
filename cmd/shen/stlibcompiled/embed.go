package stlibcompiled

import _ "embed"

// Macros contains only StLib's macro declarations (and any declarations that
// precede them in their package). The generator extracts these from the same
// vendored files as StlibMain. Runtime preprocessing registers the macros that
// bootstrap used at generation time but cannot itself emit into KL.
//go:embed macros.shen
var Macros string
