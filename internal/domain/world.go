package domain

type Point struct {
	X, Y int
}

type World struct {
	MaxX, MaxY int
	scents     map[Point]bool
}

func NewWorld(maxX, maxY int) *World {
	return &World{
		MaxX:   maxX,
		MaxY:   maxY,
		scents: make(map[Point]bool),
	}
}
