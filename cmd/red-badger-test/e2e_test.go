package main

import (
	"bufio"
	"strings"
	"testing"

	"red-badger-test/internal/parsing"
	"red-badger-test/internal/simulator"
)

func TestEndToEnd(t *testing.T) {
	input := `5 3
1 1 E
RFRFRFRF
3 2 N
FRRFLLFFRRFLL
0 3 W
LLFFFLFLFL
`

	expectedOutput := []string{
		"1 1 E",
		"3 3 N LOST",
		"2 3 S",
	}

	scanner := bufio.NewScanner(strings.NewReader(input))

	world, err := parsing.World(scanner)
	if err != nil {
		t.Fatalf("parsing world: %v", err)
	}

	runs, err := parsing.Robots(scanner)
	if err != nil {
		t.Fatalf("parsing robots: %v", err)
	}

	sim := &simulator.Simulator{}

	for _, robotRun := range runs {
		sim.RunRobot(robotRun.Robot, world, robotRun.Commands)
	}

	for i, robotRun := range runs {
		output := robotRun.Robot.String()
		if output != expectedOutput[i] {
			t.Errorf("Robot %d: got %q, want %q", i, output, expectedOutput[i])
		}
	}
}

func TestEndToEnd_EmptyInput(t *testing.T) {
	input := ""
	scanner := bufio.NewScanner(strings.NewReader(input))

	_, err := parsing.World(scanner)
	if err == nil {
		t.Error("expected error for empty input")
	}
}

func TestWithTestData(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name: "Sample input",
			input: `5 3
1 1 E
RFRFRFRF
3 2 N
FRRFLLFFRRFLL
0 3 W
LLFFFLFLFL
`,
			expected: []string{"1 1 E", "3 3 N LOST", "2 3 S"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tc.input))

			world, err := parsing.World(scanner)
			if err != nil {
				t.Fatalf("parsing world: %v", err)
			}

			runs, err := parsing.Robots(scanner)
			if err != nil {
				t.Fatalf("parsing robots: %v", err)
			}

			sim := &simulator.Simulator{}
			for _, robotRun := range runs {
				sim.RunRobot(robotRun.Robot, world, robotRun.Commands)
			}

			if len(runs) != len(tc.expected) {
				t.Fatalf("got %d robots, expected %d", len(runs), len(tc.expected))
			}

			for i, robotRun := range runs {
				output := robotRun.Robot.String()
				if output != tc.expected[i] {
					t.Errorf("Robot %d: got %q, want %q", i, output, tc.expected[i])
				}
			}
		})
	}
}
