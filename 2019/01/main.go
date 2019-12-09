package main

import (
	"log"
	"strconv"

	"github.com/mdelapenya/advent-of-code/io"
)

func main() {
	log.Println("Advent of code 2019: Day 1")

	output := "input"

	lines, err := io.ReadLines(output)
	if err != nil {
		panic(err)
	}

	fuel := 0
	for _, line := range lines {
		l, _ := strconv.Atoi(line)
		fuel += calculateRequiredFuel(l)
	}

	log.Printf("The sum of the fuel requirements for all of the modules on my spacecraft is: %d", fuel)
}

func calculateRequiredFuel(mass int) int {
	fuel := mass / 3

	fuel -= 2

	return fuel
}
