package qrcoder

import (
	"bytes"
	"errors"
	"image/png"
	"testing"
)

type ErrorWriter struct{}

func (e *ErrorWriter) Write(b []byte) (int, error) {
	return 0, errors.New("Expected error")
}

func TestGenerateQRCodeReturnsValue(t *testing.T) {
	buffer := new(bytes.Buffer)

	if err := Generate(buffer, "555-1234"); err != nil {
		t.Errorf("Oops something happened %v:", err)
	}
}

func TestGenerateQRCodeGeneratesPNG(t *testing.T) {
	buffer := new(bytes.Buffer)
	Generate(buffer, "1234567")
	_, err := png.Decode(buffer)

	if err != nil {
		t.Errorf("Generated QRcode is not a PNG: %s", err)
	}
}

func TestGenerateQRCodePropagatesErrors(t *testing.T) {
	w := new(ErrorWriter)
	err := Generate(w, "898r434")

	if err == nil || err.Error() != "Expected error" {
		t.Errorf("Error did not propagate, got %v:", err)
	}
}
