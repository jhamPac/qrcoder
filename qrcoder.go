package qrcoder

import (
	"image"
	"image/png"
	"io"
)

// Generate generates a QR code with a given string
func Generate(w io.Writer, code string) error {
	img := image.NewNRGBA(image.Rect(0, 0, 21, 21))
	return png.Encode(w, img)
}
