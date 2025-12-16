package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"red-badger-test/internal/parsing"
	"red-badger-test/internal/simulator"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	world, err := parsing.World(scanner)
	if err != nil {
		log.Fatal(fmt.Errorf("parsing world: %w", err))
	}

	runs, err := parsing.Robots(scanner)
	if err != nil {
		log.Fatal(err)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sim := &simulator.Simulator{}

	for _, robotRun := range runs {
		fmt.Println("running robot")
		sim.RunRobot(robotRun.Robot, world, robotRun.Commands)
	}
}
