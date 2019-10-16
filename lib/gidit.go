package gidit

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
	"strconv"
)

func Gicat(file1 string, file2 string) {
	// open "test.jpg"
	//var ratio uint

	imgFile1, err := os.Open(file1)
	imgFile2, err := os.Open(file2)
	if err != nil {
		fmt.Println(err)
	}
	img1, _, err := image.Decode(imgFile1)
	img2, _, err := image.Decode(imgFile2)
	if err != nil {
		fmt.Println(err)
	}
	sp2 := image.Point{img1.Bounds().Dx(), 0}
	r2 := image.Rectangle{sp2, sp2.Add(img2.Bounds().Size())}
	newPoint := img1.Bounds().Dx() + img2.Bounds().Dx()
	twoPoint := img1.Bounds().Dy() + img2.Bounds().Dy()
	r := image.Rectangle{image.Point{0, 0}, image.Point{newPoint, twoPoint}}
	rgba := image.NewRGBA(r)
	draw.Draw(rgba, img1.Bounds(), img1, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, r2, img2, image.Point{0, 0}, draw.Src)
	out, err := os.Create("new" + os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	png.Encode(out, rgba)
}

func Giappend(file1 string, file2 string) {
	// open "test.jpg"
	//var ratio uint

	imgFile1, err := os.Open(file1)
	imgFile2, err := os.Open(file2)
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
	r := image.Rectangle{image.Point{0, 0}, r2.Max}
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

func Gisize(file1 string, wide string, hide string, args string) {

	var width uint
	var height uint
	var ofile string
	if args == "" {
		ofile = "new" + os.Args[len(os.Args)-1]
	} else {
		ofile = args
	}
	file, err := os.Open(file1)
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
	if w, err := strconv.Atoi(wide); err == nil {
		width = uint(w)
	} else {
		panic("No resize given")
	}
	if h, err := strconv.Atoi(hide); err == nil {
		height = uint(h)
	} else {
		var tmp float64
		tmp = float64(b.Max.Y) / float64(b.Max.X)
		height = uint(tmp) * width
	}

	m := resize.Resize(width, height, img, resize.Lanczos3)

	out, err := os.Create(ofile)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	png.Encode(out, m)
}
