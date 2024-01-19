package matrice

import (
	"fmt"
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
		blob := cell.NewCell(false)
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

func (m *Matrice) Breath(deadzone image.Point) {
	wg := new(sync.WaitGroup)
	for y, cellColumn := range m.grid {
		for x, cell := range cellColumn {
			wg.Add(1)
			go cell.Live(wg, y <= deadzone.Y || x <= deadzone.X)
		}
	}

	wg.Wait()
}

func (m *Matrice) Alive() bool {
	for _, cellColumn := range m.grid {
		for _, cell := range cellColumn {
			if cell.Status {
				return true
			}
		}
	}
	return false
}

func (m *Matrice) BigBang() {
	for _, cellColumn := range m.grid {
		for _, cell := range cellColumn {
			cell.Talk()
		}
	}
}

func (m *Matrice) Photo() *image.Paletted {
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

func (m *Matrice) GetMerger(target *Matrice) []Merger {
	var moves []Merger
	for x, rows := range m.grid {
		fmt.Println("Parsing SourceRows")
		for y, blob := range rows {
			if !blob.Status {
				continue
			}
			sourcePoint := image.Point{x, y}
			var nearestPoint image.Point
			var targetPoint image.Point
			fmt.Printf("Get nearest point for %d, %d \n", x, y)
			for tx, trows := range target.grid {
				for ty, tblob := range trows {
					if !tblob.Status {
						continue
					}
					targetPoint = image.Point{tx, ty}
					fmt.Printf("Parsing point %d, %d \n", tx, ty)
					nearestPoint = getNearestPoint(sourcePoint, nearestPoint, targetPoint)
				}
			}
			moves = append(moves, Merger{sourcePoint, nearestPoint})
		}
	}
	return moves
}

func (m *Matrice) Merge(merger []Merger) bool {
	mergLength := len(merger) - 1
	cpt := 0
	for _, merg := range merger {
		if merg.sourcePoint.X == merg.targetPoint.X &&
			merg.sourcePoint.Y == merg.targetPoint.Y {
			fmt.Println("Same point")
			cpt++
			continue
		}
		m.grid[merg.sourcePoint.X][merg.sourcePoint.Y].Status = false
		newX := moveX(merg.sourcePoint.X, merg.targetPoint.X)
		newY := moveY(merg.sourcePoint.Y, merg.targetPoint.Y)
		fmt.Printf("Moving point %d %d to %d %d \n", merg.sourcePoint.X, merg.sourcePoint.Y, newX, newY)
		m.grid[newX][newY].Status = true
		merg.sourcePoint = image.Point{newX, newY}
	}
	return cpt < mergLength
}

func moveX(sp, tp int) int {
	if sp == tp {
		return sp
	}

	if sp < tp {
		return sp + 1
	} else {
		return sp - 1
	}
}

func moveY(sp, tp int) int {
	if sp == tp {
		return sp
	}

	if sp < tp {
		return sp + 1
	} else {
		return sp - 1
	}
}

type Merger struct {
	sourcePoint image.Point
	targetPoint image.Point
}

func getNearestPoint(sp, np, tp image.Point) image.Point {
	if np.X == tp.X && np.Y == tp.Y {
		return np
	}

	distNP := getDistance(sp, np)
	distTP := getDistance(sp, tp)

	if distNP < distTP {
		return np
	} else {
		return tp
	}
}

func getDistance(sp, np image.Point) int {
	calc := (sp.X - np.X) + (sp.Y - np.Y)
	if calc < 0 {
		return -calc
	}
	return calc
}
