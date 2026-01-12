package main

import (
	"errors"
	"fmt"
	_ "image"
	"log"
	"os"
	"os/exec"
	_ "path/filepath"
	"strings"

	//GUI
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"	
)

const (
	DIR string = "/home/thuan/Pictures/Wallpapers/"
	SCREEN_WIDTH int32 = 800
	SCREEN_HEIGHT int32 = 450
	MAX_IMG int = 6
)

var (
	IMG_EXT = [...]string{".jpg", ".jpeg", ".png", ".gif", ".svg", ".bmp", ".tiff", ".tif", ".webp", ".heic", ".heif", ".avif"}
	images = []string{}
	images_name = []string{}
)

func isImage(file_name string) bool{
	for _, ext := range(IMG_EXT) {
		if strings.HasSuffix(file_name, ext){
			return true
		}
	}
	return false
}

func switchWallpaper(file_path string) error{
	cmd := exec.Command("swww", "img", file_path, "--transition-type", "random")

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
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
		images_name = append(images_name, file_name)
		images = append(images, DIR + file_name)
	}


	file, err := os.Open(images[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 2. Decode the image
	// image.Decode detects the image format (e.g., "png", "jpeg") automatically
	//img, _, err := image.Decode(file)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//for _, img := range(images) {

	//}

	gtk.Init(&os.Args)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("GTK Go!")
	window.SetIconName("gtk-dialog-info")
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		fmt.Println("got destroy!", ctx.Data().(string))
		gtk.MainQuit()
	}, "foo")
	scrolledWindow := gtk.NewScrolledWindow(nil, nil)

	vbox := gtk.NewVBox(false, 1)

	menubar := gtk.NewMenuBar()
	vbox.PackStart(menubar, false, false, 0)

	vpaned := gtk.NewVPaned()
	vbox.Add(vpaned)

	frame := gtk.NewFrame("gows wallpaper switcher")
	framebox1 := gtk.NewVBox(false, 0)
	
	frame.Add(framebox1)

	vpaned.Pack1(frame, false, false)
	//framebox1.Add(image)
	//buttons := gtk.NewHBox(false, 1)

	row := gtk.NewHBox(false, 0)
	wallpapers := gtk.NewVBox(false, 0)
	
	count := 0
	for i, _ := range(images) {
		button := gtk.NewButtonWithLabel("")
	
		pixbuf, err := gdkpixbuf.NewPixbufFromFile(images[i])
		if err != nil {
			log.Fatal(err)
		}
		pixbuf = pixbuf.ScaleSimple(192, 180, gdkpixbuf.INTERP_BILINEAR)
		
		image := gtk.NewImage()
		image.SetFromPixbuf(pixbuf)
		
		button.SetImage(image)
		button.Clicked(func() {
			err := switchWallpaper(images[i])
			if err != nil {
				fmt.Println("blabla")
			}
		})
		row.Add(button)
		wallpapers.Add(row)

		if count == MAX_IMG {
				row = gtk.NewHBox(false, 0)
			count = 0
			continue
		}
		count++
	}

	framebox1.Add(wallpapers)

	//

	//framebox1.PackStart(buttons, false, false, 0)

	scrolledWindow.AddWithViewPort(vbox)
	//scrolledWindow.Add(framebox1)
	window.Add(scrolledWindow)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk.Main()
	fmt.Println("DONE")
}
