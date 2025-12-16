package main

import (
	"bufio"
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
}
