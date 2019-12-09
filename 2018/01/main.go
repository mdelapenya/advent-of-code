package main

import (
	"log"
	"strconv"

	"github.com/mdelapenya/advent-of-code/io"
)

func main() {
	log.Println("Advent of code 2018: Day 1")

	output := "input"

	lines, err := io.ReadLines(output)
	if err != nil {
		panic(err)
	}

	log.Printf("Sum of all frequencies: %d", sum(lines))
	log.Printf("First duplicate frequency: %d", detectFrequencyChangeList(lines))
}

func detectFrequencyChangeList(lines []string) int {
	frequency := 0
	frequencies := map[int]bool{0: true}

	for { // we have to loop over the entire list until we find the repeated frequency
		for _, line := range lines {
			delta, _ := strconv.Atoi(line)
			frequency += delta

			if _, exists := frequencies[frequency]; exists {
				return frequency
			}

			frequencies[frequency] = true
		}
	}
}

func sum(lines []string) int {
	result := 0

	for _, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		result += number
	}

	return result
}
