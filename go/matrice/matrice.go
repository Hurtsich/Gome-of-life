package matrice

import (
	"image"
	"image/color"
	"sync"

	"github.com/Hurtsich/Gome-of-life/go/cell"
)

type Matrice struct {
	grid [][]*cell.Cell
}

var (
	matrice       Matrice
	currentRow    int
	currentColumn int
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
	matrice = Matrice{grid: make([][]*cell.Cell, length)}
	for i := range matrice.grid {
		matrice.grid[i] = make([]*cell.Cell, length)
	}
	for i := 0; i <= length-1; i++ {
		for j := 0; j <= length-1; j++ {
			currentColumn = i
			currentRow = j
			var blob cell.Cell
			if matrice.grid[i][j] == nil {
				blob = cell.NewCell(randomStatus())
				matrice.grid[i][j] = &blob
			} else {
				blob = *matrice.grid[i][j]
				currentColumn = i
				currentRow = j
			}
			newNeighbor(mod((i-1), length), mod(j, length), Left, blob.Left)
			newNeighbor(mod((i-1), length), mod((j-1), length), UpLeft, blob.UpLeft)
			newNeighbor(mod((i-1), length), mod((j+1), length), DownLeft, blob.DownLeft)
			newNeighbor(mod(i, length), mod((j-1), length), Up, blob.Up)
			newNeighbor(mod((i+1), length), mod((j-1), length), UpRight, blob.UpRight)
			newNeighbor(mod((i+1), length), mod(j, length), Right, blob.Right)
			newNeighbor(mod((i+1), length), mod((j+1), length), DownRight, blob.DownRight)
			newNeighbor(mod(i, length), mod((j+1), length), Down, blob.Down)
		}
	}

	return matrice
}

func newNeighbor(column, row int, side Neighbors, membrane cell.Membrane) {
	if matrice.grid[column][row] == nil {
		currentColumn = column
		currentRow = row
		blob := cell.NewCell(randomStatus())
		matrice.grid[column][row] = &blob
		neighborsMembrane(&blob, membrane, side)
	} else {
		blob := matrice.grid[column][row]
		neighborsMembrane(blob, membrane, side)
	}
}

func neighborsMembrane(cell *cell.Cell, membrane cell.Membrane, side Neighbors) {
	switch side {
	case Up:
		cell.Down.In = membrane.Out
		cell.Down.Out = membrane.In
	case UpLeft:
		cell.DownRight.In = membrane.Out
		cell.DownRight.Out = membrane.In
	case Left:
		cell.Right.In = membrane.Out
		cell.Right.Out = membrane.In
	case DownLeft:
		cell.UpRight.In = membrane.Out
		cell.UpRight.Out = membrane.In
	case Down:
		cell.Up.In = membrane.Out
		cell.Up.Out = membrane.In
	case DownRight:
		cell.UpLeft.In = membrane.Out
		cell.UpLeft.Out = membrane.In
	case Right:
		cell.Left.In = membrane.Out
		cell.Left.Out = membrane.In
	case UpRight:
		cell.DownLeft.In = membrane.Out
		cell.DownLeft.Out = membrane.In
	}
}

func randomStatus() bool {
	if currentColumn == 2 && currentRow < 3 {
		return true
	}
	if currentColumn == 0 && currentRow == 1 {
		return true
	}
	if currentColumn == 1 && currentRow == 2 {
		return true
	}
	return false
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func (m Matrice) Breath() {
	wg := new(sync.WaitGroup)
	for _, cellColumn := range m.grid {
		for _, cell := range cellColumn {
			wg.Add(1)
			go cell.Live(wg)
		}
	}

	wg.Wait()
}

func (m Matrice) Alive() bool {
	for _, cellColumn := range m.grid {
		for _, cell := range cellColumn {
			if cell.Status {
				return true
			}
		}
	}
	return false
}

func (m Matrice) Photo() *image.Paletted {
	var palette = []color.Color{
		color.RGBA{uint8(255), uint8(255), uint8(255), uint8(255)},
		color.RGBA{uint8(0), uint8(0), uint8(0), uint8(255)},
	}
	topLeft := image.Point{0, 0}
	bottomRight := image.Point{len(m.grid), len(m.grid)}
	photo := image.NewPaletted(image.Rectangle{topLeft, bottomRight}, palette)
	for col, cellColumn := range m.grid {
		for row, cell := range cellColumn {
			if cell.Status {
				photo.Set(row, col, color.RGBA{uint8(255), uint8(255), uint8(255), uint8(255)})
			} else {
				photo.Set(row, col, color.RGBA{uint8(0), uint8(0), uint8(0), uint8(255)})
			}
		}
	}
	return photo
}
