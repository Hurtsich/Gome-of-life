package matrice

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
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

func NewGrid(length int) Matrice {
	matrice = Matrice{grid: make([][]*cell.Cell, length)}
	for i := range matrice.grid {
		matrice.grid[i] = make([]*cell.Cell, length)
	}
	for i := 0; i <= length-1; i++ {
		for j := 0; j <= length-1; j++ {
			var blob cell.Cell
			if matrice.grid[i][j] == nil {
				blob = cell.NewCell(randomStatus())
				matrice.grid[i][j] = &blob
			} else {
				blob = *matrice.grid[i][j]
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
		fmt.Println("...")
	}
	addGosperGliderGun(&matrice, image.Point{
		X: 10,
		Y: 10,
	})
	addImageAt(&matrice, image.Point{
		X: 49,
		Y: 49,
	})
	return matrice
}

func NewGridFromImage(image image.Image) Matrice {
	logo = image
	width := image.Bounds().Max.X
	height := image.Bounds().Max.Y
	fmt.Println("World creation...")
	matrice = Matrice{grid: make([][]*cell.Cell, height)}
	for i := range matrice.grid {
		matrice.grid[i] = make([]*cell.Cell, width)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			var blob cell.Cell
			if matrice.grid[i][j] == nil {
				blob = cell.NewCell(statusByColor(image.At(j, i)))
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

func addImageAt(m *Matrice, start image.Point) {
	image, err := getImageFromFilePath("../data/Slide.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	width := len(m.grid)
	height := len(m.grid)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i >= start.Y && j >= start.X {
				m.grid[i][j].Status = statusByColor(image.At(j-start.X, i-start.Y))
			}
		}
	}
}

func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return png.Decode(f)
}

func statusByColor(pixel color.Color) bool {
	blanc := color.NRGBA{uint8(255), uint8(255), uint8(255), uint8(255)}
	result := pixel != blanc
	return result
}

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
	// if currentColumn == 2 && currentRow < 3 {
	// 	return true
	// }
	// if currentColumn == 0 && currentRow == 1 {
	// 	return true
	// }
	// if currentColumn == 1 && currentRow == 2 {
	// 	return true
	// }
	// rand.Intn(100) <= 25
	return false
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
