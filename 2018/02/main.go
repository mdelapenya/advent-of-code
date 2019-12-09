package main

import (
	"log"

	"github.com/mdelapenya/advent-of-code/io"
)

func main() {
	log.Println("Advent of code 2018: Day 2")

	inputFile := "input"

	lines, err := io.ReadLines(inputFile)
	if err != nil {
		panic(err)
	}

	log.Printf("Checksum: %d", findIDs(lines))
	log.Printf("Common IDs: %s", findNearIDs(lines))
}

func diffExactlyByOne(s1 string, s2 string) (bool, string) {
	if len(s1) != len(s2) {
		return false, s1
	}

	r1 := []rune(s1)
	r2 := []rune(s2)
	differences := 0
	common := []rune("")

	for i, c1 := range r1 {
		if c1 != r2[i] {
			differences++
			if differences == 2 {
				return false, s1
			}
		} else {
			common = append(common, c1)
		}
	}

	return (differences == 1), string(common)
}

func findNearIDs(ids []string) string {
	common := ""
	diff := false

	for i, idi := range ids {
		for j, idj := range ids {
			if i == j {
				continue
			}

			diff, common = diffExactlyByOne(idi, idj)

			if diff {
				return common
			}
		}
	}

	return common
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
