package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
