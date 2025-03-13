package main

import (
	"github.com/skip2/go-qrcode"
	"os"
)

func main() {
	data := os.Args[1]

	qrcode, _ := qrcode.Encode(data, qrcode.Medium, 256)

	file, _ := os.Create("qrcode.png")

	defer file.Close()
	file.Write(qrcode)
}
