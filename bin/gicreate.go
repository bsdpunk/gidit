package main

import (
	"image"
	"image/png"
	"os"
	"strconv"
)

func Gicreate(w string, h string) {
	upLeft := image.Point{0, 0}
	width, _ := strconv.Atoi(w)
	height, _ := strconv.Atoi(h)
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	//trans := color.RGBA{0, 0, 0, 0x0}
	// Set color for each pixel.

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)

}
func main() {
	Gicreate(os.Args[len(os.Args)-2], os.Args[len(os.Args)-1])
}
