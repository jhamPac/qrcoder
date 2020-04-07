package qrcoder

import (
	"image"
	"image/png"
	"io"
)

// Version represents module size for QR codes
type Version int8

// PatternSize calculates the pattern size for each module pattern in a QR code
func (v Version) PatternSize() int {
	return 4*int(v) + 17
}

// Generate generates a QR code with a given string
func Generate(w io.Writer, code string, version Version) error {
	size := version.PatternSize()
	img := image.NewNRGBA(image.Rect(0, 0, size, size))
	return png.Encode(w, img)
}
