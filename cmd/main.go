package main

import (
	"fmt"
	"os"

	"github.com/jhampac/qrcoder"
)

func main() {
	fmt.Println("Welcome to QR Coder 3000")

	file, _ := os.Create("qrcode.png")
	defer file.Close()

	qrcoder.Generate(file, "1234567")
}
