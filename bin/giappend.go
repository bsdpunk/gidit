package main

import (
	//	"image/jpeg"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	// open "test.jpg"
	//var ratio uint

	imgFile1, err := os.Open(os.Args[1])
	imgFile2, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Println(err)
	}
	img1, _, err := image.Decode(imgFile1)
	img2, _, err := image.Decode(imgFile2)
	if err != nil {
		fmt.Println(err)
	}
	sp2 := image.Point{0, img1.Bounds().Dy()}
	r2 := image.Rectangle{sp2, sp2.Add(img2.Bounds().Size())}
	newPoint := img1.Bounds().Dx() + img2.Bounds().Dx()
	twoPoint := img1.Bounds().Dy() + img2.Bounds().Dy()
	r := image.Rectangle{image.Point{0, 0}, image.Point{newPoint, twoPoint}}
	rgba := image.NewRGBA(r)
	draw.Draw(rgba, img1.Bounds(), img1, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, r2, img2, image.Point{0, 0}, draw.Src)
	// decode jpeg into image.Image
	out, err := os.Create("new" + os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	png.Encode(out, rgba)
}
