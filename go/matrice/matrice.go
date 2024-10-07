package matrice

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"math/rand"
	"reflect"
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

func (m *Matrice) Breath() {
	wg := new(sync.WaitGroup)
	for _, cellColumn := range m.grid {
		for _, cell := range cellColumn {
			wg.Add(1)
			go cell.Live(wg)
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

func (m *Matrice) GetMergerV2(target *Matrice) []*Merger {
	var moves []*Merger
	for x, col := range target.grid {
		for y, blob := range col {
			if blob.Status {
				moves = append(moves, &Merger{targetPoint: image.Point{x, y}})
			}
		}
	}
	nearestPoint := image.Point{}
	targetPoint := image.Point{}
	nearPoints := make([]image.Point, 0)
	for _, move := range moves {
		for x, col := range m.grid {
			for y, blob := range col {
				if !blob.Status {
					continue
				}
				targetPoint = image.Point{x, y}
				nearestPoint = getNearestPoint(targetPoint, nearestPoint, move.sourcePoint)
				nearPoints = append([]image.Point{nearestPoint}, nearPoints...)
			}
		}
		index := 0
		for i, point := range nearPoints {
			index = i
			if isInMergerSource(moves, point) {
				continue
			} else {
				break
			}
		}
		move.sourcePoint = nearPoints[index]
	}
	return moves
}

func (m *Matrice) GetMerger(target *Matrice) []*Merger {
	var moves []*Merger
	for x, rows := range m.grid {
		fmt.Println("Parsing SourceRows")
		for y, blob := range rows {
			sourcePoint := image.Point{x, y}
			nearPoints := make([]image.Point, 0)
			if !blob.Status {
				continue
			}
			nearestPoint := image.Point{}
			targetPoint := image.Point{}
			fmt.Printf("Get nearest point for %d, %d \n", x, y)
			for tx, trows := range target.grid {
				for ty, tblob := range trows {
					if !tblob.Status {
						continue
					}
					targetPoint = image.Point{tx, ty}
					fmt.Printf("Parsing point %d, %d \n", tx, ty)
					nearestPoint = getNearestPoint(sourcePoint, nearestPoint, targetPoint)
					nearPoints = append([]image.Point{nearestPoint}, nearPoints...)
					if nearestPoint.X == targetPoint.X && nearestPoint.Y == targetPoint.Y {
						nearestPoint = image.Point{}
					}
				}
			}
			index := 0
			for i, point := range nearPoints {
				index = i
				if isInMerger(moves, point) {
					continue
				} else {
					break
				}
			}
			moves = append(moves, &Merger{sourcePoint, nearPoints[index]})
		}
	}
	return moves
}

func isInMerger(merg []*Merger, point image.Point) bool {
	for _, m := range merg {
		if point.X == m.targetPoint.X && point.Y == m.targetPoint.Y {
			return true
		}
	}
	return false
}

func isInMergerSource(merg []*Merger, point image.Point) bool {
	for _, m := range merg {
		if point.X == m.sourcePoint.X && point.Y == m.sourcePoint.Y {
			return true
		}
	}
	return false
}

func (m *Matrice) Merge(merger []*Merger) bool {
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
		newX := moveOneByOne(merg.sourcePoint.X, merg.targetPoint.X)
		newY := moveOneByOne(merg.sourcePoint.Y, merg.targetPoint.Y)
		fmt.Printf("Moving point %d %d to %d %d \n", merg.sourcePoint.X, merg.sourcePoint.Y, newX, newY)
		m.grid[newX][newY].Status = true
		merg.sourcePoint = image.Point{newX, newY}
	}
	return cpt < mergLength
}

func moveOneByOne(sp, tp int) int {
	if sp == tp {
		return sp
	}

	gap := sp - tp
	stride := float64(gap) / 4
	jump := int(math.Ceil(math.Abs(stride)))

	result := 0

	if sp < tp {
		result = sp + jump
	} else {
		result = sp - jump
	}
	randum := rand.Intn(100)
	if math.Abs(float64(gap)) > 4 && randum < 25 && result < len(matrice.grid)-3 {
		return result + 3
	} else if math.Abs(float64(gap)) > 4 && randum > 75 && result > 2 {
		return result - 3
	}
	return result
}

type Merger struct {
	sourcePoint image.Point
	targetPoint image.Point
}

func getNearestPoint(sp, np, tp image.Point) image.Point {
	if reflect.ValueOf(np).IsZero() {
		return tp
	}
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
