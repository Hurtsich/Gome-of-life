package cell

type Cell struct {
	Up     chan bool
	Left   chan bool
	Right  chan bool
	Down   chan bool
	tick   chan bool
	Status bool
}

func NewCell(tick chan bool, status bool) Cell {
	return Cell{
		tick:   tick,
		Up:     make(chan bool),
		Left:   make(chan bool),
		Right:  make(chan bool),
		Down:   make(chan bool),
		Status: status,
	}
}

func (c *Cell) Live() {
	for <-c.tick {
		c.Status = false
		neighbors := 0
		neighbors += isAlive(<-c.Up)
		neighbors += isAlive(<-c.Left)
		neighbors += isAlive(<-c.Right)
		neighbors += isAlive(<-c.Down)
		if 1 < neighbors && neighbors < 4 {
			c.Status = true
		}
		c.Up <- c.Status
		c.Left <- c.Status
		c.Right <- c.Status
		c.Down <- c.Status
	}
	close(c.tick)
}

func isAlive(b bool) int {
	if b {
		return 1
	}
	return 0
}
