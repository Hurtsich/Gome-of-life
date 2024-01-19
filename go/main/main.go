package main

import (
	"bufio"
	"fmt"
	gimg "github.com/Hurtsich/Gome-of-life/go/image"
	"github.com/Hurtsich/Gome-of-life/go/matrice"
	"image"
	"image/gif"
	"image/png"
	"os"
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
	//m := matrice.NewGridFromImage(i)
	//if _, err := os.Stat("../data/" + monde + ".gif"); err == nil {
	//	err := os.Remove("../data/" + monde + ".gif")
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}
	createGIFWithTransition("test")
}

func createGIFWithTransition(name string) {
	img, err := gimg.GetImageFromFilePath("../data/b.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	m := matrice.NewGridFromImage(img)
	fmt.Println("Big BANG !!")
	m.BigBang()
	gif1, delays1 := createGIFWithLastImage(&m)
	generateImage(m.Photo(), "last")

	i2, err := gimg.GetImageFromFilePath("../data/cellule0.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	m2 := matrice.NewGridFromImage(i2)
	fmt.Println("Big BANG !!")
	m2.BigBang()
	gif2, delays2 := createGIFWithLastImage(&m2)
	gif2 = reverseSort(gif2)

	merger := m.GetMerger(&m2)

	for m.Merge(merger) {
		fmt.Println("Generating...")
		delays1 = append(delays1, 7)
		photo := m.Photo()
		pi := gimg.Upscale(photo, 3)
		gif1 = append(gif1, pi)
	}

	fmt.Printf("Merger Length :%d", len(merger))

	gif1 = append(gif1, gif2...)
	delays1 = append(delays1, delays2...)

	f, err := os.Create("../data/" + name + ".gif")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	err = gif.EncodeAll(w, &gif.GIF{
		Image: gif1,
		Delay: delays1,
	})
	if err != nil {
		fmt.Println(err)
	}
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

func createGIFFromName(imgName string) {
	i, err := gimg.GetImageFromFilePath("../data/" + imgName + ".png")
	if err != nil {
		fmt.Println(err)
		return
	}
	m := matrice.NewGridFromImage(i)
	fmt.Println("Big BANG !!")
	m.BigBang()
	createGIF(&m, imgName)
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

func generateImage(img image.Image, name string) {
	m := matrice.NewGridFromImage(img)
	p := m.Photo()

	f, err := os.Create("../data/generated-" + name + ".png")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	err = png.Encode(w, p)
	if err != nil {
		fmt.Println(err)
	}
}

func reverseSort(s []*image.Paletted) []*image.Paletted {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func createGIFWithLastImage(m *matrice.Matrice) ([]*image.Paletted, []int) {
	var images []*image.Paletted
	var delays []int

	for i := 0; i < 100; i++ {
		fmt.Printf("Year: %v", i)
		delays = append(delays, 7)
		photo := m.Photo()
		pi := gimg.Upscale(photo, 3)
		images = append(images, pi)
		m.Breath(image.Point{X: 3000, Y: 2000})
		fmt.Println()
	}
	return images, delays
}

func imagineGIF(m *matrice.Matrice) ([]*image.Paletted, []int) {
	var images []*image.Paletted
	var delays []int

	for i := 0; i < 200; i++ {
		fmt.Printf("Year: %v", i)
		delays = append(delays, 7)
		photo := m.Photo()
		pi := gimg.Upscale(photo, 3)
		images = append(images, pi)
		m.Breath(image.Point{X: 3000, Y: 2000})

		fmt.Println()
	}

	return images, delays
}

func createGIF(m *matrice.Matrice, imageName string) {
	var images []*image.Paletted
	var delays []int

	for i := 0; i < 100; i++ {
		fmt.Printf("Year: %v", i)
		delays = append(delays, 7)
		photo := m.Photo()
		pi := gimg.Upscale(photo, 3)
		images = append(images, pi)
		m.Breath(image.Point{X: 3000, Y: 2000})

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
