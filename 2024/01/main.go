package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/mdelapenya/advent-of-code/io"
)

type LocationID int

func main() {
	inputLeft, inputRight := mustReadInput("input.txt")

	distance := distance(inputLeft, inputRight)
	fmt.Println(distance)
}

func distance(left []LocationID, right []LocationID) int {
	if len(left) == 0 || len(right) == 0 {
		return 0
	}

	sortLocation(left)
	sortLocation(right)

	distance := 0
	for i := 0; i < len(left); i++ {
		distance += int(math.Abs(float64(left[i] - right[i])))
	}

	return distance
}

func mustParseLocationID(s string) LocationID {
	id, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return LocationID(id)
}

func mustReadInput(path string) ([]LocationID, []LocationID) {
	lines, err := io.ReadLines(path)
	if err != nil {
		panic(err)
	}

	var inputLeft, inputRight []LocationID

	for _, line := range lines {
		parts := strings.Fields(line)

		inputLeft = append(inputLeft, mustParseLocationID(parts[0]))
		inputRight = append(inputRight, mustParseLocationID(parts[1]))
	}

	return inputLeft, inputRight
}

func sortLocation(list []LocationID) {
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
}
