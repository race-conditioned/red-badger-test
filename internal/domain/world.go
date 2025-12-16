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

func (w *World) IsWithinBounds(x, y int) bool {
	return x >= 0 && x <= w.MaxX && y >= 0 && y <= w.MaxY
}

func (w *World) HasScent(x, y int) bool {
	return w.scents[Point{x, y}]
}

func (w *World) AddScent(x, y int) {
	w.scents[Point{x, y}] = true
}

func (w *World) ProcessMove(fromX, fromY, toX, toY int) (lost bool) {
	if w.IsWithinBounds(toX, toY) {
		return false
	}

	if w.HasScent(fromX, fromY) {
		return false
	}

	w.AddScent(fromX, fromY)
	return true
}
