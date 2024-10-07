package matrice

import (
	"bufio"
	"fmt"
	"github.com/Hurtsich/Gome-of-life/go/cell"
	gimg "github.com/Hurtsich/Gome-of-life/go/image"
	"image"
	"os"
	"regexp"
	"strconv"
	"strings"
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
				blob = cell.NewCell(false)
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

func NewGridWithGliderGun(width, height int) Matrice {
	matrice = NewGrid(height, width)

	addGosperGliderGun(&matrice, image.Point{10, 10})

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

func NewMatriceFromRLEFile(filename string) (Matrice, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Matrice{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	width, height, err := parseFirstLine(line)
	if err != nil {
		return Matrice{}, err
	}

	fmt.Printf("Width = %d, Height = %d\n", width, height)
	matrice = NewGrid(height, width+300)

	re := regexp.MustCompile(`(\d*)([bo$])`)

	col := 300
	row := 0
	length := 0

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if match[1] == "" {
				length = 1
			} else {
				length, err = strconv.Atoi(match[1])
				if err != nil {
					return Matrice{}, err
				}
			}

			if match[2] == "$" {
				fmt.Printf("Getting $ : skipping %d lines\n", length)
				col = 300
				row += length
			} else if match[2] == "o" {
				fmt.Printf("Getting o\n")
				for i := 0; i < length; i++ {
					fmt.Printf("Printing this cell : %d, %d\n", col, row)
					matrice.grid[row][col].Status = true
					col++
				}
			} else {
				fmt.Printf("Getting b, skipping %d columns\n", length)
				col += length
			}
		}
	}

	return matrice, nil
}

func parseFirstLine(line string) (width int, height int, err error) {
	x := 0
	y := 0
	lineParts := strings.Split(line, ",")
	fmt.Println(lineParts)
	for _, part := range lineParts {
		coordinates := strings.Split(part, "=")
		trimedCoordinate := strings.TrimSpace(coordinates[0])
		trimedNumber := strings.TrimSpace(coordinates[1])
		if trimedCoordinate == "x" {
			x, err = strconv.Atoi(trimedNumber)
			if err != nil {
				return
			}
			width = x
		} else if trimedCoordinate == "y" {
			y, err = strconv.Atoi(trimedNumber)
			if err != nil {
				return
			}
			height = y
		}
	}
	return
}
