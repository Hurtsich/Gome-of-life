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
	m := matrice.NewGrid(500)
	if _, err := os.Stat("../data/" + monde + ".gif"); err == nil {
		err := os.Remove("../data/" + monde + ".gif")
		if err != nil {
			fmt.Println(err)
		}
	}

	m.BigBang()

	createGIF(&m)
}

func createGIF(m *matrice.Matrice) {
	var images []*image.Paletted
	var delays []int

	for i := 0; i < 1000; i++ {
		delays = append(delays, 0)
		photo := m.Photo()
		images = append(images, photo)
		m.Breath()
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
