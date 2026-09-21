package kl

import (
	"bytes"
	"errors"
	"testing"
)

func TestInstallPrBuffer(t *testing.T) {
	InstallPr()
	var ctl ControlFlow
	for _, text := range []string{"", "HELLO", "A\x00éÿ"} {
		var buf bytes.Buffer
		value := MakeString(text)
		got := Call(&ctl, ctl.Global(MakeSymbol("shen.native-pr")), value, MakeStream(&buf))
		if got != value {
			t.Fatal("pr did not return its input")
		}
		var want []byte
		for _, r := range text {
			want = append(want, byte(r))
		}
		if !bytes.Equal(buf.Bytes(), want) {
			t.Fatalf("%q: got %x, want %x", text, buf.Bytes(), want)
		}
	}
}

func TestPrByteErrorKeepsPrefix(t *testing.T) {
	for _, text := range []string{"AĀ", "A\xff"} {
		t.Run(text, func(t *testing.T) {
			var buf bytes.Buffer
			defer func() {
				if recover() == nil {
					t.Error("expected byte-range error")
				}
				if buf.String() != "A" {
					t.Errorf("written prefix = %q", buf.String())
				}
			}()
			primPr(MakeString(text), MakeStream(&buf))
		})
	}
}

type prFailWriter struct{}

func (prFailWriter) Write([]byte) (int, error) { return 0, errors.New("write failed") }

func TestPrPropagatesWriteError(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("expected write error")
		}
	}()
	primPr(MakeString("x"), MakeStream(prFailWriter{}))
}
