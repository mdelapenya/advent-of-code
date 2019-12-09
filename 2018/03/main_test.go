package main

import (
	"testing"

	"github.com/mdelapenya/advent-of-code/io"
	"github.com/stretchr/testify/assert"
)

func TestCountClaims(t *testing.T) {
	assert := assert.New(t)

	lines, _ := io.ReadLines("input")
	assert.Equal(105231, countClaims(lines))
}

func TestParse(t *testing.T) {
	assert := assert.New(t)

	expected := square{
		ID:     1,
		Left:   808,
		Top:    550,
		Width:  12,
		Height: 22,
	}
	actual := parse("#1 @ 808,550: 12x22")

	assert.Equal(expected.ID, actual.ID)
	assert.Equal(expected.Left, actual.Left)
	assert.Equal(expected.Top, actual.Top)
	assert.Equal(expected.Width, actual.Width)
	assert.Equal(expected.Height, actual.Height)
}

func TestParseInput(t *testing.T) {
	assert := assert.New(t)

	expected := square{
		ID:     195,
		Left:   604,
		Top:    177,
		Width:  16,
		Height: 11,
	}
	actual := parse("#195 @ 604,177: 16x11")

	assert.Equal(expected.ID, actual.ID)
	assert.Equal(expected.Left, actual.Left)
	assert.Equal(expected.Top, actual.Top)
	assert.Equal(expected.Width, actual.Width)
	assert.Equal(expected.Height, actual.Height)
}
