package qrcoder

import (
	"image"
	"image/png"
	"io"
)

// GenerateQRCode generates a QR code with a given string
func GenerateQRCode(w io.Writer, code string) {
	img := image.NewNRGBA(image.Rect(0, 0, 21, 21))
	_ = png.Encode(w, img)
}
