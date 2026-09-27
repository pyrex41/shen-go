//go:build linux || darwin || freebsd || netbsd || openbsd || dragonfly

package kl

import (
	"syscall"
	"unsafe"
)

// reserveFixnumSpace maps n bytes of zero-filled, read-only, private anonymous
// memory for the fixnum span, or returns nil. Read-only private memory is not
// committed, so this reserves address space only.
func reserveFixnumSpace(n int) unsafe.Pointer {
	b, err := syscall.Mmap(-1, 0, n, syscall.PROT_READ, syscall.MAP_PRIVATE|syscall.MAP_ANON)
	if err != nil || len(b) != n {
		return nil
	}
	return unsafe.Pointer(&b[0])
}
