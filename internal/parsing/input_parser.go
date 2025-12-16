package parsing

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"red-badger-test/internal/domain"
)

func World(scanner *bufio.Scanner) (*domain.World, error) {
	if !scanner.Scan() {
		return nil, fmt.Errorf("missing world bounds")
	}

	bounds := scanner.Text()
	parts := strings.Fields(bounds)

	if len(parts) != 2 {
		return nil, fmt.Errorf("parsing world bounds: %q", bounds)
	}

	maxX, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("parsing maxX: %w", err)
	}

	maxY, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("parsing maxY: %w", err)
	}

	if maxX < 0 {
		return nil, fmt.Errorf("negative maxX")
	}

	if maxY < 0 {
		return nil, fmt.Errorf("negative maxY")
	}

	world := domain.NewWorld(maxX, maxY)
	return world, nil
}

func Robots(scanner *bufio.Scanner) ([]domain.RobotRun, error) {
	var runs []domain.RobotRun
	for {
		robot, instructions, err := parseRobot(scanner)
		if err != nil {
			return nil, fmt.Errorf("parsing robots: %w", err)
		}
		if robot == nil {
			break
		}

		commands := parseCommands(instructions)

		runs = append(runs, domain.RobotRun{
			Robot:    robot,
			Commands: commands,
		})
	}
	return runs, nil
}

func parseRobot(scanner *bufio.Scanner) (*domain.Robot, string, error) {
	for {
		if !scanner.Scan() {
			return nil, "", nil
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 3 {
			return nil, "", fmt.Errorf("parsing robot position: %q", line)
		}

		x, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, "", fmt.Errorf("parsing robot x: %w", err)
		}

		y, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, "", fmt.Errorf("parsing robot y: %w", err)
		}

		orientation, err := domain.OrientationFromString(parts[2])
		if err != nil {
			return nil, "", fmt.Errorf("parsing orientation: %w", err)
		}

		robot := domain.NewRobot(x, y, orientation)

		if !scanner.Scan() {
			return nil, "", fmt.Errorf("missing instruction line for robot")
		}

		instructions := strings.TrimSpace(scanner.Text())

		return robot, instructions, nil
	}
}

func parseCommands(instructions string) []domain.Command {
	var commands []domain.Command
	for _, c := range instructions {
		switch c {
		case 'L':
			commands = append(commands, domain.LeftCommand{})
		case 'R':
			commands = append(commands, domain.RightCommand{})
		case 'F':
			commands = append(commands, domain.ForwardCommand{})
		default:
			log.Printf("Unknown command: %c", c)
		}
	}
	return commands
}
