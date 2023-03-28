package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/Hurtsich/Gome-of-life/go/matrice"
)

func main() {
	m := matrice.NewGrid(100)
	photo := m.Breath()

	err := os.Remove("../data/Begin.png")
	if err != nil {
		fmt.Println(err)
	}

	f, err := os.Create("../data/Begin.png")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	err = png.Encode(w, photo.SubImage(photo.Rect))
	if err != nil {
		fmt.Println(err)
	}

}
