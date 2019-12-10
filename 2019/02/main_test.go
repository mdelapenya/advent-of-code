package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIntcode(t *testing.T) {
	assert := assert.New(t)

	codes := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}

	inst := NewInstruction(0, codes)

	assert.Equal(1, inst.opcode)
	assert.Equal(9, inst.arguments[0])
	assert.Equal(10, inst.arguments[1])
	assert.Equal(3, inst.arguments[2])

	inst = NewInstruction(4, codes)

	assert.Equal(2, inst.opcode)
	assert.Equal(3, inst.arguments[0])
	assert.Equal(11, inst.arguments[1])
	assert.Equal(0, inst.arguments[2])
}

func TestExecuteIntcodeMultiply(t *testing.T) {
	assert := assert.New(t)

	codes := []int{2, 3, 0, 3, 99}

	inst := NewInstruction(0, codes)
	result := inst.execute(codes)

	assert.Equal(6, result[inst.arguments[len(inst.arguments)-1]])
}

func TestExecuteIntcodeMultiply2(t *testing.T) {
	assert := assert.New(t)

	codes := []int{2, 4, 4, 5, 99, 0}

	inst := NewInstruction(0, codes)
	result := inst.execute(codes)

	assert.Equal(9801, result[inst.arguments[len(inst.arguments)-1]])
}

func TestExecuteIntcodeSum(t *testing.T) {
	assert := assert.New(t)

	codes := []int{1, 0, 0, 0, 99}
	inst := NewInstruction(0, codes)
	result := inst.execute(codes)

	assert.Equal(2, result[inst.arguments[len(inst.arguments)-1]])
}

func TestLoop(t *testing.T) {
	assert := assert.New(t)

	codes := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}

	i := loop(1, 1, codes)

	assert.Equal(30, i)
}
