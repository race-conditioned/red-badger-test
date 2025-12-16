package domain

import "fmt"

type RobotRun struct {
	Robot    *Robot
	Commands []Command
}

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

func (r *Robot) NextPosition() (int, int) {
	dx, dy := r.Orientation.ForwardDelta()
	return r.X + dx, r.Y + dy
}

func (r *Robot) MoveTo(x, y int) {
	r.X = x
	r.Y = y
}

func (r *Robot) MarkLost() {
	r.Lost = true
}

func (r *Robot) String() string {
	if r.Lost {
		return fmt.Sprintf("%d %d %s LOST", r.X, r.Y, r.Orientation.String())
	}
	return fmt.Sprintf("%d %d %s", r.X, r.Y, r.Orientation.String())
}
