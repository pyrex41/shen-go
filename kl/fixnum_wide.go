//go:build !windows && (amd64 || arm64)

package kl

// A 1 GiB array is the most that fits: the linker refuses a symbol over 2e9
// bytes. It is the fixnum range [-2^29, 2^29).
const fixnumBits = 30
