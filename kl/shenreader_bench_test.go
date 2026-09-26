package kl

import (
	"os"
	"path/filepath"
	"testing"
)

// BenchmarkShenReadKernelSources times the native reader over the kernel .shen
// sources (the files it accepts without falling back).
func BenchmarkShenReadKernelSources(b *testing.B) {
	files, _ := filepath.Glob("../kernel/sources/*.shen")
	var datas [][]byte
	n := 0
	for _, f := range files {
		d, _ := os.ReadFile(f)
		if _, err := ShenReadSExprs(d); err == nil {
			datas = append(datas, d)
			n += len(d)
		}
	}
	b.SetBytes(int64(n))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, d := range datas {
			ShenReadSExprs(d)
		}
	}
}
