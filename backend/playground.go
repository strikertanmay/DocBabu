package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/clsung/grcode"
)

func run() {
	// Create the barcode
	qrCode, _ := qr.Encode("5e3e4a90706e1ada9d8f42a1", qr.H, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 250, 250)

	// create the output file
	file, _ := os.Create("qrcode1.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)
	decodeQR("qrcode1.png")
}

//go build -ldflags "-linkmode external -extldflags -static"
func decodeQR(filePath string) {

	results, err := grcode.GetDataFromFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	if len(results) == 0 {
		log.Printf("No qrcode detected from file: %s", filePath)
	}
	for _, result := range results {
		fmt.Printf("%s\n", result)
	}
}
