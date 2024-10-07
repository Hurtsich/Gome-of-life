package cell

import (
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
	neighbors := c.Listen()
	if (neighbors == 2 || neighbors == 3) && c.Status {
		c.Status = true
	} else if neighbors == 3 && !c.Status {
		c.Status = true
	} else {
		c.Status = false
	}
	c.Talk()
}

func (c *Cell) Listen() int {
	neighbors := 0
	neighbors += isAlive(c.Up.In)
	neighbors += isAlive(c.UpLeft.In)
	neighbors += isAlive(c.Left.In)
	neighbors += isAlive(c.DownLeft.In)
	neighbors += isAlive(c.Down.In)
	neighbors += isAlive(c.DownRight.In)
	neighbors += isAlive(c.Right.In)
	neighbors += isAlive(c.UpRight.In)
	return neighbors
}

func (c *Cell) Talk() {
	c.Up.Out <- c.Status
	c.UpLeft.Out <- c.Status
	c.Left.Out <- c.Status
	c.DownLeft.Out <- c.Status
	c.Down.Out <- c.Status
	c.DownRight.Out <- c.Status
	c.Right.Out <- c.Status
	c.UpRight.Out <- c.Status
}

func isAlive(b chan bool) int {
	select {
	case alive := <-b:
		if alive {
			return 1
		} else {
			return 0
		}
	default:
		return 0
	}
}
