package qrcoder

import (
	"bytes"
	"errors"
	"fmt"
	"image/png"
	"testing"
)

type ErrorWriter struct{}

func (e *ErrorWriter) Write(b []byte) (int, error) {
	return 0, errors.New("Expected error")
}

func TestGenerateQRCodeReturnsValue(t *testing.T) {
	buffer := new(bytes.Buffer)

	if err := Generate(buffer, "555-1234", Version(1)); err != nil {
		t.Errorf("Oops something happened %v:", err)
	}
}

func TestGenerateQRCodeGeneratesPNG(t *testing.T) {
	buffer := new(bytes.Buffer)
	Generate(buffer, "1234567", Version(1))
	_, err := png.Decode(buffer)

	if err != nil {
		t.Errorf("Generated QRcode is not a PNG: %s", err)
	}
}

func TestGenerateQRCodePropagatesErrors(t *testing.T) {
	w := new(ErrorWriter)
	err := Generate(w, "898r434", Version(1))

	if err == nil || err.Error() != "Expected error" {
		t.Errorf("Error did not propagate, got %v:", err)
	}
}

func TestVersionDeterminesSize(t *testing.T) {
	cases := []struct {
		version  int
		expected int
	}{
		{1, 21},
		{2, 25},
		{6, 41},
		{7, 45},
		{14, 73},
		{40, 177},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("test case for version number %d", test.version), func(t *testing.T) {
			buffer := new(bytes.Buffer)
			_ = Generate(buffer, "some text", Version(test.version))
			img, _ := png.Decode(buffer)
			if width := img.Bounds().Dx(); width != test.expected {
				t.Errorf("Version %2d, expected %3d but got %3d", test.version, test.expected, width)
			}
		})
	}
}
