package matrice

import (
	"math/rand"

	"github.com/Hurtsich/Gome-of-life/go/cell"
)

type Matrice struct {
	grid [][]cell.Cell
}

var (
	matrice Matrice
)

type Neighbors int

const (
	Up Neighbors = iota
	UpLeft
	Left
	DownLeft
	Down
	DownRight
	Right
	UpRight
)

func NewGrid(length int) Matrice {
	matrice = Matrice{make([][]cell.Cell, length)}
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			var blob cell.Cell
			if matrice.grid[i][j] != (cell.Cell{}) {
				blob = cell.NewCell(RandomStatus())
				matrice.grid[i][j] = blob
			} else {
				blob = matrice.grid[i][j]
			}
			if i > 0 {
				newNeighbor(i-1, j, Up, &blob.Up)
				if j > 0 {
					newNeighbor(i-1, j-1, UpLeft, &blob.UpLeft)
				}
				if j < length {
					newNeighbor(i-1, j+1, UpRight, &blob.UpRight)
				}
			}
			if j > 0 {
				newNeighbor(i, j-1, Left, &blob.Left)
				if (i < length) {
					newNeighbor(i+1, j-1, DownLeft, &blob.DownLeft)
				}
			}
			if i < length {
				newNeighbor(i+1, j, Down, &blob.Down)
				if j < length {
					newNeighbor(i+1, j+1, DownRight, &blob.DownRight)
				}
			}
			if j < length {
				newNeighbor(i, j+1, Right, &blob.Right)
			}
		}
	}
	return matrice
}

func newNeighbor(column, row int, side Neighbors, membrane *cell.Membrane) {
	if matrice.grid[column][row] != (cell.Cell{}) {
		blob := cell.NewCell(RandomStatus())
		neighborsMembrane(blob, membrane, side)
	} else {
		neighborsMembrane(matrice.grid[column][row], membrane, side)
	}
}

func neighborsMembrane(cell cell.Cell, membrane *cell.Membrane, side Neighbors) {
	switch side {
	case Up:
		cell.Down = *membrane
	case UpLeft:
		cell.DownRight = *membrane
	case Left:
		cell.Right = *membrane
	case DownLeft:
		cell.UpRight = *membrane
	case Down:
		cell.Up = *membrane
	case DownRight:
		cell.UpLeft = *membrane
	case Right:
		cell.Left = *membrane
	case UpRight:
		cell.DownLeft = *membrane
	}
}

func RandomStatus() bool {
	return rand.Intn(100) < 40
}
