package domain

import "errors"

var orientationToString = map[Orientation]string{
	North: "N",
	East:  "E",
	South: "S",
	West:  "W",
}

var stringToOrientation = map[string]Orientation{
	"N": North,
	"E": East,
	"S": South,
	"W": West,
}

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

func (o Orientation) String() string {
	return orientationToString[o]
}

func OrientationFromString(s string) (Orientation, error) {
	v, ok := stringToOrientation[s]

	if !ok {
		return 0, errors.New("invalid orientation")
	}
	return v, nil
}
