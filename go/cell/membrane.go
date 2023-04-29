package cell

type Membrane struct {
	In  chan bool
	Out chan bool
}

func NewMembrane() Membrane {
	return Membrane{
		In:  make(chan bool, 1),
		Out: make(chan bool, 1),
	}
}
