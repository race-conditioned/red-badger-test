package parsing

import (
	"bufio"
	"strings"
	"testing"

	"red-badger-test/internal/domain"
)

func TestWorld(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantMaxX int
		wantMaxY int
		wantErr  bool
	}{
		{
			"valid bounds",
			"5 3\n",
			5, 3, false,
		},
		{
			"bounds with extra spaces",
			"   5     3  \n",
			5, 3, false,
		},
		{
			"invalid: not enough numbers",
			"5\n",
			0, 0, true,
		},
		{
			"invalid: non-numeric",
			"5 X\n",
			0, 0, true,
		},
		{
			"invalid: negative bounds",
			"-1 3\n",
			0, 0, true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.input))
			got, err := World(scanner)

			if (err != nil) != tt.wantErr {
				t.Errorf("World() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if got.MaxX != tt.wantMaxX || got.MaxY != tt.wantMaxY {
					t.Errorf("World() = %v, want MaxX=%d, MaxY=%d", got, tt.wantMaxX, tt.wantMaxY)
				}
			}
		})
	}
}

func TestParseRobot(t *testing.T) {
	tests := []struct {
		name             string
		input            string
		wantRobot        *domain.Robot
		wantInstructions string
		wantErr          bool
	}{
		{
			"valid robot",
			"1 1 E\nRFRFRFRF\n",
			domain.NewRobot(1, 1, domain.East),
			"RFRFRFRF",
			false,
		},
		{
			"valid robot with spaces",
			"   3   2   N   \nFRRFLLFFRRFLL\n",
			domain.NewRobot(3, 2, domain.North),
			"FRRFLLFFRRFLL",
			false,
		},
		{
			"invalid: missing instruction line",
			"1 1 E\n",
			nil, "", true,
		},
		{
			"invalid: bad orientation",
			"1 1 X\nRFRFRFRF\n",
			nil, "", true,
		},
		{
			"invalid: not enough position values",
			"1 1\nRFRFRFRF\n",
			nil, "", true,
		},
		{
			"skip empty lines before robot",
			"\n\n1 1 E\nRFRFRFRF\n",
			domain.NewRobot(1, 1, domain.East),
			"RFRFRFRF",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.input))
			gotRobot, gotInstructions, err := parseRobot(scanner)

			if (err != nil) != tt.wantErr {
				t.Errorf("parseRobot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if gotRobot.X != tt.wantRobot.X || gotRobot.Y != tt.wantRobot.Y ||
					gotRobot.Orientation != tt.wantRobot.Orientation {
					t.Errorf("parseRobot() robot = %v, want %v", gotRobot, tt.wantRobot)
				}
				if gotInstructions != tt.wantInstructions {
					t.Errorf("parseRobot() instructions = %q, want %q", gotInstructions, tt.wantInstructions)
				}
			}
		})
	}
}

func TestParseCommands(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantCount int
		wantTypes []string
	}{
		{
			"mixed commands",
			"RFRFRFRF",
			8,
			[]string{"R", "F", "R", "F", "R", "F", "R", "F"},
		},
		{
			"all left turns",
			"LLLL",
			4,
			[]string{"L", "L", "L", "L"},
		},
		{
			"empty string",
			"",
			0,
			[]string{},
		},
		{
			"with unknown command (should log but not crash)",
			"LXR",
			2,
			[]string{"L", "R"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseCommands(tt.input)

			if len(got) != tt.wantCount {
				t.Errorf("parseCommands() count = %d, want %d", len(got), tt.wantCount)
			}

			for i, cmd := range got {
				// execute on a dummy robot
				robot := domain.NewRobot(0, 0, domain.North)
				world := domain.NewWorld(5, 3)
				startOrientation := robot.Orientation

				cmd.Execute(robot, world)

				if tt.wantTypes[i] == "L" && robot.Orientation != startOrientation.Left() {
					t.Errorf("Command %d is not Left", i)
				} else if tt.wantTypes[i] == "R" && robot.Orientation != startOrientation.Right() {
					t.Errorf("Command %d is not Right", i)
				} else if tt.wantTypes[i] == "F" && (robot.X != 0 || robot.Y != 1) {
					t.Errorf("Command %d is not Forward", i)
				}
			}
		})
	}
}

func TestRobots(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantCount int
		wantErr   bool
	}{
		{
			"single robot",
			"1 1 E\nRFRFRFRF\n",
			1, false,
		},
		{
			"multiple robots",
			"1 1 E\nRFRFRFRF\n3 2 N\nFRRFLLFFRRFLL\n",
			2, false,
		},
		{
			"no robots",
			"",
			0, false,
		},
		{
			"incomplete robot (missing instructions)",
			"1 1 E\n",
			0, true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.input))

			got, err := Robots(scanner)

			if (err != nil) != tt.wantErr {
				t.Errorf("Robots() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && len(got) != tt.wantCount {
				t.Errorf("Robots() count = %d, want %d", len(got), tt.wantCount)
			}
		})
	}
}
