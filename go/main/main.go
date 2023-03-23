package main

import (
	"fmt"

	"github.com/Hurtsich/Gome-of-life/go/matrice"
)

func main() {
	m := matrice.NewGrid(100)
	fmt.Printf("%v", m)
}