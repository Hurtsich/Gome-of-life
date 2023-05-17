package main

import (
	"bufio"
	"fmt"
	"image"
	"image/gif"
	"os"

	"github.com/Hurtsich/Gome-of-life/go/matrice"
)

var monde = "Test"

func main() {
	m := matrice.NewGrid(200)
	//i, err := getImageFromFilePath("../data/logo.png")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//m := matrice.NewGridFromImage(i)
	if _, err := os.Stat("../data/" + monde + ".gif"); err == nil {
		err := os.Remove("../data/" + monde + ".gif")
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Big BANG !!")
	m.BigBang()

	createGIF(&m)
}

func createGIF(m *matrice.Matrice) {
	var images []*image.Paletted
	var delays []int

	for i := 0; i < 450; i++ {
		fmt.Printf("Year: %v", i)
		delays = append(delays, 0)
		photo := m.Photo()
		images = append(images, photo)
		if i < 320 {
			m.Breath(image.Point{X: 50, Y: 90})
		} else {
			m.Breath(image.Point{X: 200, Y: 200})
		}
		fmt.Println()
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
