package main

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <data>")
		return 
	}

	data := os.Args[1]

	qrcode, err := qrcode.Encode(data, qrcode.Medium, 256)
	if err != nil {
		fmt.Println("Error generating QR code:", err)
		return
	}
	
	file, err := os.Create("qrcode.png")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(qrcode)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("QR code saved to qrcode.png")
}
