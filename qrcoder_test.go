package qrcoder

import "testing"

func TestGenerateQRCodeReturnsValue(t *testing.T) {
	result := GenerateQRCode("555-1234")

	if result == nil {
		t.Error("Generated QRCode is nil")
	}

	if len(result) == 0 {
		t.Error("Generated QRCode as no data")
	}
}
