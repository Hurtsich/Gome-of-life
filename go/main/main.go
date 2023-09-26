package main

import (
	"bufio"
	"fmt"
	"image"
	"image/gif"
	"os"

	"github.com/Hurtsich/Gome-of-life/go/matrice"
	"github.com/ichinaski/pxl"
)

var monde = "Test"

func main() {
	fmt.Println("World creation...")
	m := matrice.NewGrid(500, 500)
	// i, err := gimg.GetImageFromFilePath("../data/slide.png")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// m := matrice.NewGridFromImage(i)
	// if _, err := os.Stat("../data/" + monde + ".gif"); err == nil {
	// 	err := os.Remove("../data/" + monde + ".gif")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
	fmt.Println("Big BANG !!")
	m.BigBang()
	createGIF(&m)
}

func createGIF(m *matrice.Matrice) {
	var images []*image.Paletted
	var delays []int

	for i := 0; i < 150; i++ {
		fmt.Printf("Year: %v", i)
		delays = append(delays, 0)
		photo := m.Photo()
		images = append(images, photo)
		m.Breath(image.Point{X: 1000, Y: 1000})

		// if i < 320 {
		// 	m.Breath(image.Point{X: 50, Y: 90})
		// } else {
		// 	m.Breath(image.Point{X: 200, Y: 200})
		// }
		fmt.Println()

		pxl.
	}

	f, err := os.Create("../data/" + monde + ".gif")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	err = gif.EncodeAll(w, &gif.GIF{
		Image: images,
		Delay: delays,
	})
	if err != nil {
		fmt.Println(err)
	}
}
