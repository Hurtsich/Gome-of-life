package cell

import (
	"fmt"
	"sync"
)

type Cell struct {
	Up        Membrane
	UpLeft    Membrane
	Left      Membrane
	DownLeft  Membrane
	Down      Membrane
	DownRight Membrane
	Right     Membrane
	UpRight   Membrane
	Status    bool
}

func NewCell(status bool) Cell {
	return Cell{
		Status:    status,
		Up:        NewMembrane(),
		UpLeft:    NewMembrane(),
		Left:      NewMembrane(),
		DownLeft:  NewMembrane(),
		Down:      NewMembrane(),
		DownRight: NewMembrane(),
		Right:     NewMembrane(),
		UpRight:   NewMembrane(),
	}
}

func (c *Cell) Live(wg *sync.WaitGroup) {
	defer wg.Done()
	c.Up.Out <- c.Status
	c.UpLeft.Out <- c.Status
	c.Left.Out <- c.Status
	c.DownLeft.Out <- c.Status
	c.Down.Out <- c.Status
	c.DownRight.Out <- c.Status
	c.Right.Out <- c.Status
	c.UpRight.Out <- c.Status
	fmt.Println("Status OUT")
	neighbors := 0
	neighbors += isAlive(<-c.Up.In)
	neighbors += isAlive(<-c.UpLeft.In)
	neighbors += isAlive(<-c.Left.In)
	neighbors += isAlive(<-c.DownLeft.In)
	neighbors += isAlive(<-c.Down.In)
	neighbors += isAlive(<-c.DownRight.In)
	neighbors += isAlive(<-c.Right.In)
	neighbors += isAlive(<-c.UpRight.In)
	fmt.Println("Calculating Status")
	// c.Status = randomStatus()
	if 1 < neighbors && neighbors < 4 && c.Status {
		c.Status = true
	} else if neighbors > 4 {
		c.Status = true
	} else {
		c.Status = false
	}
}

// func randomStatus() bool {
// 	return rand.Intn(100) < 10
// }

func isAlive(b bool) int {
	if b {
		return 1
	}
	return 0
}
