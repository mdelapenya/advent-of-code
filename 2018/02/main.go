package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	log.Println("Advent of code 2018: Day 2")

	inputFile := "input"

	lines, err := readLines(inputFile)
	if err != nil {
		panic(err)
	}

	log.Printf("Checksum: %d", findIDs(lines))
}

func findIDs(ids []string) int {
	two := 0   // detect IDs with two duplicate letters
	three := 0 // detect IDs with three duplicate letters

	for _, id := range ids {
		if hasExactlyTwo(id) == 1 {
			two++
		}

		if hasExactlyThree(id) == 1 {
			three++
		}

	}

	return two * three
}

func hasExactly(id string, x int) int {
	chars := []rune(id)

	dict := make(map[rune]int)

	for _, c := range chars {
		dict[c]++
	}

	items := 0
	for _, c := range chars {
		if dict[c] == x {
			items++

			return items
		}
	}

	return 0
}

func hasExactlyThree(id string) int {
	return hasExactly(id, 3)
}

func hasExactlyTwo(id string) int {
	return hasExactly(id, 2)
}

// readLines reads a whole file into memory and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
