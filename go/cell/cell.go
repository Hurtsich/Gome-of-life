package cell

import "context"

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
		Status: status,
		Up: NewMembrane(),
		UpLeft: NewMembrane(),
		Left: NewMembrane(),
		DownLeft: NewMembrane(),
		Down: NewMembrane(),
		DownRight: NewMembrane(),
		Right: NewMembrane(),
		UpRight: NewMembrane(),
	}
}

func (c *Cell) Live(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			neighbors := 0
			neighbors += isAlive(<-c.Up.In)
			neighbors += isAlive(<-c.UpLeft.In)
			neighbors += isAlive(<-c.Left.In)
			neighbors += isAlive(<-c.DownLeft.In)
			neighbors += isAlive(<-c.Down.In)
			neighbors += isAlive(<-c.DownRight.In)
			neighbors += isAlive(<-c.Right.In)
			neighbors += isAlive(<-c.UpRight.In)
			if 1 < neighbors && neighbors < 4 && c.Status {
				c.Status = true
			} else if neighbors > 4 {
				c.Status = true
			} else {
				c.Status = false
			}
			c.Up.Out <- c.Status
			c.UpLeft.Out <- c.Status
			c.Left.Out <- c.Status
			c.DownLeft.Out <- c.Status
			c.Down.Out <- c.Status
			c.DownRight.Out <- c.Status
			c.Right.Out <- c.Status
			c.UpRight.Out <- c.Status
		}
	}
}

func isAlive(b bool) int {
	if b {
		return 1
	}
	return 0
}
