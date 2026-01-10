package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"image"
	_ "path/filepath"
	"strings"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	_ "fyne.io/fyne/v2/widget"
)

const (
	DIR string = "/home/thuan/Pictures/Wallpapers/"
)

var (
	IMG_EXT = [...]string{".jpg", ".jpeg", ".png", ".gif", ".svg", ".bmp", ".tiff", ".tif", ".webp", ".heic", ".heif", ".avif"}
	images = []string{}
)

func isImage(file_name string) bool{
	for _, ext := range(IMG_EXT) {
		if strings.HasSuffix(file_name, ext){
			return true
		}
	}
	return false
}

func main() {

	files, err := os.ReadDir(DIR)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range(files) {
		file_name := file.Name()
		if !isImage(file_name) {
			log.Fatal(errors.New("contains file(s) is not image"))
		}
		image_path := DIR + file_name
		images = append(images, image_path)

	}


	file, err := os.Open(images[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 2. Decode the image
	// image.Decode detects the image format (e.g., "png", "jpeg") automatically
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}


	//fmt.Println(images)

	
	myApp := app.New()
	w := myApp.NewWindow("Image")
	image := canvas.NewImageFromImage(img)
	image.FillMode = canvas.ImageFillOriginal
	w.SetContent(image)

	w.ShowAndRun()
	

	fmt.Println("DONE")
}
