package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectFirstFrequency(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, detectFrequencyChangeList([]string{"+1", "-1"}))
	assert.Equal(10, detectFrequencyChangeList([]string{"+3", "+3", "+4", "-2", "-4"}))
	assert.Equal(5, detectFrequencyChangeList([]string{"-6", "+3", "+8", "+5", "-6"}))
	assert.Equal(14, detectFrequencyChangeList([]string{"+7", "+7", "-2", "-7", "-4"}))
}

func TestDetectFirstFrequencyFromInput(t *testing.T) {
	assert := assert.New(t)

	lines, _ := readLines("input")
	assert.Equal(241, detectFrequencyChangeList(lines))
}

func TestSum(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, sum([]string{"+1", "+1", "+1"}))
	assert.Equal(0, sum([]string{"+1", "+1", "-2"}))
	assert.Equal(-6, sum([]string{"-1", "-2", "-3"}))
}

func TestSumFromInput(t *testing.T) {
	assert := assert.New(t)

	lines, _ := readLines("input")
	assert.Equal(592, sum(lines))
}

func TestSumFromEmpty(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, sum([]string{}))
}
