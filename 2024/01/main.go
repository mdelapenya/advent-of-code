package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/mdelapenya/advent-of-code/io"
	aocstrings "github.com/mdelapenya/advent-of-code/strings"
)

type LocationID int

func main() {
	inputLeft, inputRight := mustReadInput("input.txt")

	distance := distance(inputLeft, inputRight)
	fmt.Println("distance:", distance)

	score := similarityScore(inputLeft, inputRight)
	fmt.Println("similarity score:", score)
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

func mustReadInput(path string) ([]LocationID, []LocationID) {
	lines, err := io.ReadLines(path)
	if err != nil {
		panic(err)
	}

	var inputLeft, inputRight []LocationID

	for _, line := range lines {
		parts := strings.Fields(line)

		inputLeft = append(inputLeft, LocationID(aocstrings.MustParseInt(parts[0])))
		inputRight = append(inputRight, LocationID(aocstrings.MustParseInt(parts[1])))
	}

	return inputLeft, inputRight
}

func similarityScore(left []LocationID, right []LocationID) int {
	score := 0

	for _, leftID := range left {
		rightCount := 0
		for _, rightID := range right {
			if leftID == rightID {
				rightCount++
			}
		}

		score += int(leftID) * rightCount
	}

	return score
}

func sortLocation(list []LocationID) {
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
}
