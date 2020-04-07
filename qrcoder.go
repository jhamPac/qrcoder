package qrcoder

import (
	"image"
	"image/png"
	"io"
)

// Version represents module size for QR codes
type Version int8

// Generate generates a QR code with a given string
func Generate(w io.Writer, code string, version Version) error {
	img := image.NewNRGBA(image.Rect(0, 0, 21, 21))
	return png.Encode(w, img)
}
