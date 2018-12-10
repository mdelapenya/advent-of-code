package main

import (
	"log"
	"strings"

	"github.com/mdelapenya/advent-of-code/2018/io"
)

const lowerCase = "abcdefghijklmnñopqrstuvwxyz"
const upperCase = "ABCDEFGHIJKLMNÑOPQRSTUVWXYZ"

var runesLower = []rune(lowerCase)
var runesUpper = []rune(upperCase)

func concat(r1 rune, r2 rune) string {
	return string(r1) + string(r2)
}

func hasReactions(polymer string) bool {
	for i := 0; i < 27; i++ {
		reaction1 := concat(runesUpper[i], runesLower[i])
		reaction2 := concat(runesLower[i], runesUpper[i])

		if strings.Contains(polymer, reaction1) || strings.Contains(polymer, reaction2) {
			return true
		}
	}

	return false
}

func main() {
	log.Println("Advent of code 2018: Day 5")

	inputFile := "polymers"

	lines, err := io.ReadLines(inputFile)
	if err != nil {
		panic(err)
	}

	if len(lines) > 1 {
		log.Fatalln("The input polymer is wrong")
	}

	polymer := removeReactions(lines[0])
	log.Printf("The resultant polymer is %s, with length of %d", polymer, len(polymer))
}

func removeReactions(polymer string) string {
	result := polymer

	for {
		for i := 0; i < 27; i++ {
			reaction1 := concat(runesUpper[i], runesLower[i])
			reaction2 := concat(runesLower[i], runesUpper[i])

			result = strings.Replace(result, reaction1, "", -1)
			result = strings.Replace(result, reaction2, "", -1)
		}

		if !hasReactions(result) {
			break
		}
	}

	return result
}
