package domain

import (
	"testing"
)

func TestForwardCommand_Execute(t *testing.T) {
	tests := []struct {
		name                 string
		startX, startY       int
		startOrientation     Orientation
		worldMaxX, worldMaxY int
		hasScentAtStart      bool
		wantX, wantY         int
		wantLost             bool
		wantScentAdded       bool
	}{
		{
			"move within bounds",
			1, 1, North, 5, 3,
			false,
			1, 2, false, false,
		},
		{
			"move out of bounds no scent",
			5, 3, North, 5, 3,
			false,
			5, 3, true, true,
		},
		{
			"move out of bounds with scent",
			5, 3, North, 5, 3,
			true,
			5, 3, false, false,
		},
		{
			"lost robot doesn't move",
			1, 1, North, 5, 3,
			false,
			1, 1, true, false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			world := NewWorld(tt.worldMaxX, tt.worldMaxY)
			if tt.hasScentAtStart {
				world.AddScent(tt.startX, tt.startY)
			}

			robot := NewRobot(tt.startX, tt.startY, tt.startOrientation)
			if tt.name == "lost robot doesn't move" {
				robot.MarkLost()
			}

			cmd := ForwardCommand{}
			cmd.Execute(robot, world)

			if robot.X != tt.wantX || robot.Y != tt.wantY {
				t.Errorf("position = (%d, %d), want (%d, %d)", robot.X, robot.Y, tt.wantX, tt.wantY)
			}
			if robot.Lost != tt.wantLost {
				t.Errorf("lost = %v, want %v", robot.Lost, tt.wantLost)
			}
			if tt.wantScentAdded && !world.HasScent(tt.startX, tt.startY) {
				t.Error("scent not added when expected")
			}
		})
	}
}
