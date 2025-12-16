package domain

import (
	"testing"
)

func TestWorld_HasScent(t *testing.T) {
	world := NewWorld(5, 3)

	if world.HasScent(2, 2) {
		t.Error("HasScent() = true, want false for initial world")
	}

	world.AddScent(2, 2)
	if !world.HasScent(2, 2) {
		t.Error("HasScent() = false, want true after AddScent")
	}

	if world.HasScent(2, 3) {
		t.Error("HasScent() = true for different coordinate, want false")
	}
}

func TestWorld_ProcessMove(t *testing.T) {
	world := NewWorld(5, 3)

	tests := []struct {
		name           string
		fromX, fromY   int
		toX, toY       int
		wantLost       bool
		shouldAddScent bool
	}{
		{"move within bounds", 1, 1, 1, 2, false, false},
		{"move out of bounds no scent", 5, 3, 5, 4, true, true},
		{"move out of bounds with scent", 5, 3, 5, 4, false, false},
		{"move out of bounds different direction", 5, 3, 6, 3, false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "move out of bounds no scent" {
				world = NewWorld(5, 3)
			}

			gotLost := world.ProcessMove(tt.fromX, tt.fromY, tt.toX, tt.toY)

			if gotLost != tt.wantLost {
				t.Errorf("ProcessMove() = %v, want %v", gotLost, tt.wantLost)
			}

			if tt.shouldAddScent && !world.HasScent(tt.fromX, tt.fromY) {
				t.Error("scent not added at expected location")
			}
		})
	}
}

func TestWorld_ProcessMove_MultipleRobotsSameEdge(t *testing.T) {
	world := NewWorld(5, 3)

	// first robot gets lost
	lost := world.ProcessMove(5, 3, 5, 4)
	if !lost {
		t.Error("first robot should be lost")
	}
	if !world.HasScent(5, 3) {
		t.Error("scent should be added at (5, 3)")
	}

	// second robot from same point should not get lost
	lost = world.ProcessMove(5, 3, 5, 4)
	if lost {
		t.Error("second robot should not be lost due to scent")
	}

	// third robot from same point but different direction, also shouldn't be lost
	lost = world.ProcessMove(5, 3, 6, 3)
	if lost {
		t.Error("Third robot should not be lost due to scent (any direction)")
	}
}
