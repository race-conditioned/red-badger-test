package domain

import (
	"testing"
)

func TestRobot_NextPosition(t *testing.T) {
	tests := []struct {
		name         string
		x, y         int
		orientation  Orientation
		wantX, wantY int
	}{
		{"North from (1,1)", 1, 1, North, 1, 2},
		{"East from (1,1)", 1, 1, East, 2, 1},
		{"South from (1,1)", 1, 1, South, 1, 0},
		{"West from (1,1)", 1, 1, West, 0, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			robot := NewRobot(tt.x, tt.y, tt.orientation)
			gotX, gotY := robot.NextPosition()
			if gotX != tt.wantX || gotY != tt.wantY {
				t.Errorf("NextPosition() = (%d, %d), want (%d, %d)", gotX, gotY, tt.wantX, tt.wantY)
			}
		})
	}
}
