package main

import (
	"bufio"
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"os"

	gimg "github.com/Hurtsich/Gome-of-life/go/image"
	"github.com/Hurtsich/Gome-of-life/go/matrice"
)

var monde = "Test"

func main() {
	fmt.Println("World creation...")
	// m := matrice.NewGrid(500, 500)
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
	GenerateSlides()
}

func GenerateSlides() {
	generateSlides("cellule0")
	generateSlides("cellule1")
	generateSlides("cellule2")
	generateSlides("cellule3")

	i, err := gimg.GetImageFromFilePath("../data/cellule3.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	m := matrice.NewGridFromImage(i)
	fmt.Println("Big BANG !!")
	m.BigBang()
	createGIF(&m, "cellule3")

	generateSlides("gol0")
	generateSlides("gol1")
	generateSlides("gol2")
	generateSlides("gol3")
	generateSlides("gol4")
	generateSlides("gol5")
	generateSlides("gol6")
	generateSlides("gol7")
	generateSlides("gol8")
	generateSlides("gol9")
	generateSlides("gol10")

	i, err = gimg.GetImageFromFilePath("../data/gol10.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	m = matrice.NewGridFromImage(i)
	fmt.Println("Big BANG !!")
	m.BigBang()
	createGIF(&m, "gol10")
}

func generateSlides(slideName string) {
	i, err := gimg.GetImageFromFilePath("../data/" + slideName + ".png")
	if err != nil {
		fmt.Println(err)
		return
	}
	m := matrice.NewGridFromImage(i)
	p := m.Photo()
	pi := gimg.Upscale(p, 10)

	f, err := os.Create("../data/generated-" + slideName + ".png")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	err = png.Encode(w, pi)
	if err != nil {
		fmt.Println(err)
	}
}

func createGIF(m *matrice.Matrice, imageName string) {
	var images []*image.Paletted
	var delays []int

	for i := 0; i < 100; i++ {
		fmt.Printf("Year: %v", i)
		delays = append(delays, 0)
		photo := m.Photo()
		upPhoto := gimg.Upscale(photo, 10)
		images = append(images, upPhoto)
		m.Breath(image.Point{X: 2000, Y: 2000})

		// if i < 320 {
		// 	m.Breath(image.Point{X: 50, Y: 90})
		// } else {
		// 	m.Breath(image.Point{X: 200, Y: 200})
		// }
		fmt.Println()
	}

	f, err := os.Create("../data/" + imageName + ".gif")
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
