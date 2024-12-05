package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	left  = []LocationID{3, 4, 2, 1, 3, 3}
	right = []LocationID{4, 3, 5, 3, 9, 3}
)

func TestDistance(t *testing.T) {
	distance := distance(left, right)
	require.Equal(t, 11, distance)
}

func TestMustReadInput(t *testing.T) {
	_, _ = mustReadInput("input.txt")
}

func TestSimilarityScore(t *testing.T) {
	score := similarityScore(left, right)
	require.Equal(t, 31, score)
}

func TestSortLocations(t *testing.T) {
	// copy left and right to avoid modifying the original arrays
	sortedLeft := append([]LocationID(nil), left...)
	sortedRight := append([]LocationID(nil), right...)

	sortLocation(sortedLeft)
	sortLocation(sortedRight)

	require.Equal(t, []LocationID{1, 2, 3, 3, 3, 4}, sortedLeft)
	require.Equal(t, []LocationID{3, 3, 3, 4, 5, 9}, sortedRight)
}
