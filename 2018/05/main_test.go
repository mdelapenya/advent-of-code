package main

import (
	"testing"

	"github.com/mdelapenya/advent-of-code/io"
	"github.com/stretchr/testify/assert"
)

func TestPolymer(t *testing.T) {
	assert := assert.New(t)

	polymers, _ := io.ReadLines("polymers")

	polymer := removeReactions(polymers[0])
	assert.Equal(10250, len(polymer))
}

func TestGetShorterPolymer(t *testing.T) {
	assert := assert.New(t)

	polymers, _ := io.ReadLines("polymers")

	polymer := removeReactions(polymers[0])
	key, shorterPolymer := getShorterReaction(polymer)
	assert.Equal(key, "l/L")
	assert.Equal(6188, len(shorterPolymer))
}

func TestGetShorterReaction(t *testing.T) {
	assert := assert.New(t)

	s, polymer := getShorterReaction("dabAcCaCBAcCcaDA")
	assert.Equal("daDA", polymer)
	assert.Equal("c/C", s)
}

func TestHasReactions(t *testing.T) {
	assert := assert.New(t)

	assert.True(hasReactions("aAa"))
	assert.True(hasReactions("aA"))
	assert.True(hasReactions("Aa"))

	assert.False(hasReactions("aa"))
	assert.False(hasReactions("ab"))
	assert.False(hasReactions("aB"))
	assert.False(hasReactions("Ab"))
}

func TestRemoveReactions(t *testing.T) {
	assert := assert.New(t)

	polymer := removeReactions("dabAcCaCBAcCcaDA")
	assert.Equal("dabCBAcaDA", polymer)
}
