package matrice

import (
	"image"
	"image/color"
	"math/rand"
	"sync"

	"github.com/Hurtsich/Gome-of-life/go/cell"
)

type Matrice struct {
	grid [][]*cell.Cell
}

var (
	matrice Matrice
	logo    image.Image
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

func newNeighbor(column, row int, side Neighbors, membrane cell.Membrane) {
	if matrice.grid[column][row] == nil {
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
	return rand.Intn(100) <= 25
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func (m Matrice) Breath(deadzone image.Point) {
	wg := new(sync.WaitGroup)
	for y, cellColumn := range m.grid {
		for x, cell := range cellColumn {
			wg.Add(1)
			go cell.Live(wg, y <= deadzone.Y || x <= deadzone.X)
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

func (m Matrice) BigBang() {
	for _, cellColumn := range m.grid {
		for _, cell := range cellColumn {
			cell.Talk()
		}
	}
}

func (m Matrice) Photo() *image.Paletted {
	var palette = []color.Color{
		color.RGBA{uint8(255), uint8(255), uint8(255), uint8(255)},
		color.RGBA{uint8(0), uint8(0), uint8(0), uint8(255)},
	}
	topLeft := image.Point{0, 0}
	bottomRight := image.Point{len(m.grid[0]), len(m.grid)}
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
