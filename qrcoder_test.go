package qrcoder

import (
	"bytes"
	"image/png"
	"testing"
)

func TestGenerateQRCodeReturnsValue(t *testing.T) {
	result := GenerateQRCode("555-1234")

	if result == nil {
		t.Error("Generated QRCode is nil")
	}

	if len(result) == 0 {
		t.Error("Generated QRCode as no data")
	}
}

func TestGenerateQRCodeGeneratesPNG(t *testing.T) {
	result := GenerateQRCode("1234567")
	buffer := bytes.NewBuffer(result)
	_, err := png.Decode(buffer)

	if err != nil {
		t.Errorf("Generated QRCode is not a PNG: %s", err)
	}
}
