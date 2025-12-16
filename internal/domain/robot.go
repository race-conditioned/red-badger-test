package domain

type Robot struct {
	X, Y int
	Orientation
	Lost bool
}

func NewRobot(x, y int, orientation Orientation) *Robot {
	return &Robot{
		X:           x,
		Y:           y,
		Orientation: orientation,
		Lost:        false,
	}
}
