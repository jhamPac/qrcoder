package qrcoder

import (
	"bytes"
	"image/png"
	"testing"
)

func TestGenerateQRCodeReturnsValue(t *testing.T) {
	buffer := new(bytes.Buffer)
	Generate(buffer, "555-1234")

	if buffer.Len() == 0 {
		t.Error("No QRCode generated")
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
