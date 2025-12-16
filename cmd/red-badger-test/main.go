package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"red-badger-test/internal/domain"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		log.Fatal("failed to read world bounds")
	}

	bounds := scanner.Text()
	parts := strings.Fields(bounds)
	if len(parts) != 2 {
		log.Fatal("invalid world bounds")
	}

	maxX, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal("Invalid maxX")
	}

	maxY, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal("Invalid maxY")
	}

	world := domain.NewWorld(maxX, maxY)
	log.Println(world)

	for {
		robot, instructions, err := parseRobot(scanner)
		if err != nil {
			log.Fatal(err)
		}
		if robot == nil {
			break
		}

		log.Printf("Parsed robot: %+v, instructions: %s\n", robot, instructions)
	}
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
			return nil, "", fmt.Errorf("invalid robot position: %q", line)
		}

		x, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, "", fmt.Errorf("invalid robot x: %w", err)
		}

		y, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, "", fmt.Errorf("invalid robot y: %w", err)
		}

		orientation, err := domain.OrientationFromString(parts[2])
		if err != nil {
			return nil, "", fmt.Errorf("invalid orientation: %w", err)
		}

		robot := domain.NewRobot(x, y, orientation)

		if !scanner.Scan() {
			return nil, "", fmt.Errorf("missing instruction line for robot")
		}

		instructions := strings.TrimSpace(scanner.Text())
		if instructions == "" {
			return nil, "", fmt.Errorf("empty instruction line")
		}

		return robot, instructions, nil
	}
}
