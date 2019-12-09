package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/mdelapenya/advent-of-code/io"
)

const exitCode = 99
const multiplyCode = 2
const sumCode = 1

func main() {
	log.Println("Advent of code 2019: Day 2")

	output := "input"

	lines, err := io.ReadLines(output)
	if err != nil {
		panic(err)
	}

	sCodes := strings.Split(lines[0], ",")
	codes := make([]int, len(sCodes))
	for i, sCode := range sCodes {
		codes[i], _ = strconv.Atoi(sCode)
	}

	codes[1] = 12
	codes[2] = 2

	zero := loop(codes)

	log.Printf("The value that is left at position 0 after the program halts is %d", zero)
}

func loop(codes []int) int {
	for i := 0; i < (len(codes) - 4); i += 4 {
		if codes[i] == exitCode {
			break
		}

		intCode := NewIntcode(i, codes)
		codes = intCode.execute(codes)
	}

	return codes[0]
}

// Intcode represents an int code program
type Intcode struct {
	operator int // 1 = sum; 2 = multiplication
	first    int // position of the first operator
	second   int // position of the second operator
	result   int // position of the result
}

// NewIntcode returns an Intcode from a position in the array
func NewIntcode(pos int, codes []int) Intcode {
	return Intcode{
		operator: codes[pos],
		first:    codes[pos+1],
		second:   codes[pos+2],
		result:   codes[pos+3],
	}
}

func (ic *Intcode) execute(codes []int) []int {
	value := 0
	if ic.operator == sumCode {
		value = codes[ic.first] + codes[ic.second]
	} else if ic.operator == multiplyCode {
		value = codes[ic.first] * codes[ic.second]
	} else if ic.operator == exitCode {
		return codes
	} else {
		log.Panicf("Code not accepted: %d", ic.operator)
	}

	codes[ic.result] = value

	return codes
}
