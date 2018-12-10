package main

import (
	"log"
	"strings"
	"unicode"

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

func getShorterReaction(polymer string) (string, string) {
	shortestPolymers := map[string]string{}

	for _, lower := range lowerCase {
		upper := unicode.ToUpper(lower)
		result := strings.Replace(polymer, string(lower), "", -1)
		result = strings.Replace(result, string(upper), "", -1)

		r := removeReactions(result)

		shortestPolymers[string(lower)+"/"+string(upper)] = r
	}

	var k string
	shortestPolymer := polymer

	for _, lower := range lowerCase {
		upper := unicode.ToUpper(lower)
		key := string(lower) + "/" + string(upper)

		val := shortestPolymers[key]

		if len(val) < len(shortestPolymer) {
			shortestPolymer = val
			k = key
		}
	}

	return k, shortestPolymer
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

	r, p := getShorterReaction(polymer)
	log.Printf("The shorter polymer is caused by the removal of %s, with length %d", r, len(p))
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
