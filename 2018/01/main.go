package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	log.Println("Advent of code 2018")

	output := "input"

	lines, err := readLines(output)
	if err != nil {
		panic(err)
	}

	log.Println(sum(lines))
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
