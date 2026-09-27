//go:build !(linux || darwin || freebsd || netbsd || openbsd || dragonfly)

package kl

import "unsafe"

// reserveFixnumSpace has no reservation on this platform; the fixnum span is
// the static array.
func reserveFixnumSpace(int) unsafe.Pointer { return nil }
