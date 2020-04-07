package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jhampac/qrcoder"
)

func main() {
	fmt.Println("Welcome to QR Coder 3000")

	file, err := os.Create("qrcode.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err := qrcoder.Generate(file, "1234567"); err != nil {
		log.Fatal(err)
	}
}
