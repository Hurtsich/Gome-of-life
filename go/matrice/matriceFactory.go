package matrice

import (
	"fmt"
	"image"

	"github.com/Hurtsich/Gome-of-life/go/cell"
	gimg "github.com/Hurtsich/Gome-of-life/go/image"
)

func NewGrid(height, width int) Matrice {
	matrice = Matrice{grid: make([][]*cell.Cell, height)}
	for i := range matrice.grid {
		matrice.grid[i] = make([]*cell.Cell, width)
	}
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			var blob cell.Cell
			if matrice.grid[i][j] == nil {
				blob = cell.NewCell(randomStatus())
				matrice.grid[i][j] = &blob
			} else {
				blob = *matrice.grid[i][j]
			}
			newNeighbor(mod((i-1), height), mod(j, width), Left, blob.Left)
			newNeighbor(mod((i-1), height), mod((j-1), width), UpLeft, blob.UpLeft)
			newNeighbor(mod((i-1), height), mod((j+1), width), DownLeft, blob.DownLeft)
			newNeighbor(mod(i, height), mod((j-1), width), Up, blob.Up)
			newNeighbor(mod((i+1), height), mod((j-1), width), UpRight, blob.UpRight)
			newNeighbor(mod((i+1), height), mod(j, width), Right, blob.Right)
			newNeighbor(mod((i+1), height), mod((j+1), width), DownRight, blob.DownRight)
			newNeighbor(mod(i, height), mod((j+1), width), Down, blob.Down)
		}
		fmt.Println("...")
	}
	return matrice
}

func NewGridFromImage(img image.Image) Matrice {
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y
	matrice = NewGrid(height, width)

	addImageAt(&matrice, image.Point{X: 0, Y: 0}, img)

	return matrice
}

func addImageAt(m *Matrice, start image.Point, img image.Image) {
	width := len(m.grid[0])
	height := len(m.grid)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i >= start.Y && j >= start.X {
				m.grid[i][j].Status = gimg.StatusByColor(img.At(j-start.X, i-start.Y))
			}
		}
	}
}

func addGosperGliderGun(matrix *Matrice, start image.Point) {
	matrix.grid[start.X+5][start.Y+1].Status, matrix.grid[start.X+5][start.Y+2].Status = true, true
	matrix.grid[start.X+6][start.Y+1].Status, matrix.grid[start.X+6][start.Y+2].Status = true, true
	matrix.grid[start.X+3][start.Y+13].Status, matrix.grid[start.X+3][start.Y+14].Status = true, true
	matrix.grid[start.X+4][start.Y+12].Status, matrix.grid[start.X+4][start.Y+16].Status = true, true
	matrix.grid[start.X+5][start.Y+11].Status, matrix.grid[start.X+5][start.Y+17].Status = true, true
	matrix.grid[start.X+6][start.Y+11].Status, matrix.grid[start.X+6][start.Y+15].Status = true, true
	matrix.grid[start.X+6][start.Y+17].Status, matrix.grid[start.X+6][start.Y+18].Status = true, true
	matrix.grid[start.X+7][start.Y+11].Status, matrix.grid[start.X+7][start.Y+17].Status = true, true
	matrix.grid[start.X+8][start.Y+12].Status, matrix.grid[start.X+8][start.Y+16].Status = true, true
	matrix.grid[start.X+9][start.Y+13].Status, matrix.grid[start.X+9][start.Y+14].Status = true, true
	matrix.grid[start.X+1][start.Y+25].Status = true
	matrix.grid[start.X+2][start.Y+23].Status, matrix.grid[start.X+2][start.Y+25].Status = true, true
	matrix.grid[start.X+3][start.Y+21].Status, matrix.grid[start.X+3][start.Y+22].Status = true, true
	matrix.grid[start.X+4][start.Y+21].Status, matrix.grid[start.X+4][start.Y+22].Status = true, true
	matrix.grid[start.X+5][start.Y+21].Status, matrix.grid[start.X+5][start.Y+22].Status = true, true
	matrix.grid[start.X+6][start.Y+23].Status, matrix.grid[start.X+6][start.Y+25].Status = true, true
	matrix.grid[start.X+7][start.Y+25].Status = true
	matrix.grid[start.X+3][start.Y+35].Status = true
	matrix.grid[start.X+3][start.Y+36].Status = true
	matrix.grid[start.X+4][start.Y+35].Status = true
	matrix.grid[start.X+4][start.Y+36].Status = true
}
