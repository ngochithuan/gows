package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
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
	//fmt.Println(images)

	fmt.Println("DONE!")
}
