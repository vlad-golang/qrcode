package qrcode

import (
	"os"
	"testing"
)

const tmpDir = "tmp"

func TestGenerate(t *testing.T) {
	encoder := NewEncoder()
	qr, err := encoder.Encode("1", 20)
	if err != nil {
		t.Error(err)
		return
	}

	if len(qr) != 12144 {
		t.Errorf("qr size is %v", len(qr))
		return
	}
}

func TestGenerateFile(t *testing.T) {
	err := os.MkdirAll(tmpDir, os.ModePerm)
	if err != nil {
		t.Error(err)
		return
	}
	encoder := NewEncoder()
	qr, err := encoder.Encode("1", 20)
	if err != nil {
		t.Error(err)
		return
	}

	file, err := os.Create(tmpDir + "/qr.png")
	if err != nil {
		t.Error(err)
		return
	}

	_, err = file.Write(qr)
	if err != nil {
		t.Error(err)
		return
	}
}
