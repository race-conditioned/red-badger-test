package domain

type Orientation int

const (
	North Orientation = iota
	East
	South
	West
)

func (o Orientation) Left() Orientation {
	return (o + 3) % 4
}

func (o Orientation) Right() Orientation {
	return (o + 1) % 4
}

func (o Orientation) ForwardDelta() (dx, dy int) {
	switch o {
	case North:
		return 0, 1
	case East:
		return 1, 0
	case South:
		return 0, -1
	case West:
		return -1, 0
	}
	return 0, 0
}
