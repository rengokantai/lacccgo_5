package main
import (
"fmt"
"os"
"path/filepath"
	"log"
	//"image/jpeg"
	"github.com/rwcarlsen/goexif/exif"
	"io/ioutil"
)

func main() {

filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
if info.IsDir() {
return nil
}

// ext is empty if no file extension
ext := filepath.Ext(path)

switch ext {
case ".jpg", ".jpeg":
	fmt.Println(ext)

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

//	img, err := jpeg.Decode(f)
//	r, g, b, a := img.At(0, 0).RGBA()
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("%d %d %d %d \n", r, g, b, a)


	xi(f)
}
return nil
})
}

func xi(f *os.File) {
	x, err := exif.Decode(f)
	if err != nil {
		log.SetOutput(ioutil.Discard)
		log.Println(err)
	}
	if x != nil {
		tm, _ := x.DateTime()
		fmt.Println("Taken: ", tm)

		lat, long, _ := x.LatLong()
		fmt.Println("lat, long: ", lat, ", ", long)
		fmt.Println(x.String())
	}
}