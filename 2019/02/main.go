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

	for i := 1; i < 100; i++ {
		for j := 1; j < 100; j++ {
			codes := make([]int, len(sCodes))
			for i, sCode := range sCodes {
				codes[i], _ = strconv.Atoi(sCode)
			}

			if loop(i, j, codes) == 19690720 {
				log.Printf("The input noun and verb that cause the program to produce the output 19690720 (100 * noun + verb) is: noun=%d verb=%d", i, j)

				return
			}
		}
	}

	log.Fatal("Sorry, not found :(")
}

func loop(noun int, verb int, codes []int) int {
	codes[1] = noun
	codes[2] = verb

	for i := 0; i < (len(codes) - 4); i += 4 {
		if codes[i] == exitCode {
			break
		}

		inst := NewInstruction(i, codes)
		codes = inst.execute(codes)
	}

	return codes[0]
}

// Instruction represented by the opcode and its arguments
type Instruction struct {
	opcode    int
	arguments []int
}

// NewInstruction creates an instruction from a Intcode
func NewInstruction(pos int, codes []int) Instruction {
	return Instruction{
		opcode:    codes[pos],
		arguments: []int{codes[pos+1], codes[pos+2], codes[pos+3]},
	}
}

// NewHaltInstruction creates an instruction for halting the program
func NewHaltInstruction() Instruction {
	return Instruction{
		opcode:    exitCode,
		arguments: []int{},
	}
}

func (i *Instruction) nextPointer() int {
	return 1 + len(i.arguments)
}

func (i *Instruction) execute(codes []int) []int {
	value := 0
	if i.opcode == sumCode {
		value = codes[i.arguments[0]] + codes[i.arguments[1]]
	} else if i.opcode == multiplyCode {
		value = codes[i.arguments[0]] * codes[i.arguments[1]]
	} else if i.opcode == exitCode {
		return codes
	} else {
		log.Panicf("Code not accepted: %d", i.opcode)
	}

	codes[i.arguments[2]] = value

	return codes
}
