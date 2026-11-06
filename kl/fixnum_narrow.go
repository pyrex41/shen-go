//go:build windows || !(amd64 || arm64)

package kl

// Windows charges BSS against the commit limit whether or not it is touched, and
// a 32-bit address space has little to spare, so these keep 64 MiB of array: the
// fixnum range [-2^25, 2^25).
const fixnumBits = 26
