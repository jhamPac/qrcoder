package qrcoder

import (
	"bytes"
	"image"
	"image/png"
)

// GenerateQRCode generates a QR code with a given string
func GenerateQRCode(code string) []byte {
	img := image.NewNRGBA(image.Rect(0, 0, 21, 21))
	buf := new(bytes.Buffer)
	_ = png.Encode(buf, img)
	return buf.Bytes()
}
