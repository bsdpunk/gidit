package main

import (
	"github.com/nfnt/resize"
	//	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strconv"
)

func main() {
	// open "test.jpg"
	var width uint
	var height uint
	//var ratio uint
	file, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		panic("No file selected")
	}

	// decode jpeg into image.Image
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	file.Close()
	b := img.Bounds()
	if w, err := strconv.Atoi(os.Args[1]); err == nil {
		width = uint(w)
	} else {
		panic("No resize given")
	}
	if h, err := strconv.Atoi(os.Args[2]); err == nil {
		height = uint(h)
	} else {
		var tmp float64
		//height = int(float64(b.Dy)/float64(b.Dx)) * width
		//ratio := uint(0)
		tmp = float64(b.Max.Y) / float64(b.Max.X)
		height = uint(tmp) * width
	}
	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(width, height, img, resize.Lanczos3)
	//fmt.Printf("%+v\n", b)
	//fmt.Println("Original", b)

	//fmt.Println(m.Bounds())
	out, err := os.Create("new" + os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	png.Encode(out, m)
}
