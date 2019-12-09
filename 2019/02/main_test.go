package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIntcode(t *testing.T) {
	assert := assert.New(t)

	codes := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}

	i := NewIntcode(0, codes)

	assert.Equal(1, i.operator)
	assert.Equal(9, i.first)
	assert.Equal(10, i.second)
	assert.Equal(3, i.result)

	i = NewIntcode(4, codes)

	assert.Equal(2, i.operator)
	assert.Equal(3, i.first)
	assert.Equal(11, i.second)
	assert.Equal(0, i.result)
}

func TestExecuteIntcodeMultiply(t *testing.T) {
	assert := assert.New(t)

	codes := []int{2, 3, 0, 3, 99}

	i := NewIntcode(0, codes)

	result := i.execute(codes)

	assert.Equal(6, result[i.result])
}

func TestExecuteIntcodeMultiply2(t *testing.T) {
	assert := assert.New(t)

	codes := []int{2, 4, 4, 5, 99, 0}

	i := NewIntcode(0, codes)

	result := i.execute(codes)

	assert.Equal(9801, result[i.result])
}

func TestExecuteIntcodeSum(t *testing.T) {
	assert := assert.New(t)

	codes := []int{1, 0, 0, 0, 99}
	i := NewIntcode(0, codes)
	result := i.execute(codes)

	assert.Equal(2, result[i.result])
}

func TestLoop(t *testing.T) {
	assert := assert.New(t)

	codes := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}

	i := loop(codes)

	assert.Equal(30, i)
}
