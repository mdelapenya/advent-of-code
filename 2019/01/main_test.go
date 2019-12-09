package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateRequiredFuel(t *testing.T) {
	assert := assert.New(t)

	fuel := calculateRequiredFuel(1969)
	assert.Equal(966, fuel)

	fuel = calculateRequiredFuel(100756)
	assert.Equal(50346, fuel)

	fuel = calculateRequiredFuel(14)
	assert.Equal(2, fuel)
}
