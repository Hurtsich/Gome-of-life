package helpers

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