package main

import (
	"testing"

	"github.com/mdelapenya/advent-of-code/2018/io"
	"github.com/stretchr/testify/assert"
)

func TestDistance(t *testing.T) {
	p1 := point{ID: 23, X: 12, Y: -4}
	p2 := point{ID: 45, X: 26, Y: -3}

	assert := assert.New(t)

	distance := p1.getManhattanDistance(p2)
	assert.Equal(23, distance.From.ID)
	assert.Equal(1, len(distance.PointIds))
	assert.Equal(45, distance.PointIds[0])
	assert.Equal(15, distance.Distance)
}

func TestDistanceSamePoint(t *testing.T) {
	p1 := point{ID: 1, X: 0, Y: 0}

	assert := assert.New(t)

	distance := p1.getManhattanDistance(p1)
	assert.Equal(1, distance.From.ID)
	assert.Equal(1, len(distance.PointIds))
	assert.Equal(1, distance.PointIds[0])
	assert.Equal(0, distance.Distance)
}

func TestDistanceSameAbscises(t *testing.T) {
	p1 := point{ID: 1, X: 0, Y: 0}
	p2 := point{ID: 2, X: 1, Y: 0}

	assert := assert.New(t)

	distance := p1.getManhattanDistance(p2)
	assert.Equal(1, distance.From.ID)
	assert.Equal(1, len(distance.PointIds))
	assert.Equal(2, distance.PointIds[0])
	assert.Equal(1, distance.Distance)

	p2.X = 49
	distance = p1.getManhattanDistance(p2)
	assert.Equal(1, distance.From.ID)
	assert.Equal(1, len(distance.PointIds))
	assert.Equal(2, distance.PointIds[0])
	assert.Equal(49, distance.Distance)
}

func TestDistanceSameOrdered(t *testing.T) {
	p1 := point{ID: 1, X: 0, Y: 0}
	p2 := point{ID: 2, X: 0, Y: 1}

	assert := assert.New(t)

	distance := p1.getManhattanDistance(p2)
	assert.Equal(1, distance.From.ID)
	assert.Equal(1, len(distance.PointIds))
	assert.Equal(2, distance.PointIds[0])
	assert.Equal(1, distance.Distance)

	p2.Y = 49
	distance = p1.getManhattanDistance(p2)
	assert.Equal(1, distance.From.ID)
	assert.Equal(1, len(distance.PointIds))
	assert.Equal(2, distance.PointIds[0])
	assert.Equal(49, distance.Distance)
}

func TestGetGreaterArea(t *testing.T) {
	assert := assert.New(t)

	lines, _ := io.ReadLines("input")

	area := getGreaterArea(lines)

	assert.Equal(3882, area)
}

func TestGetMinDistances(t *testing.T) {
	assert := assert.New(t)

	lines, _ := io.ReadLines("input")

	points := getPoints(lines)

	min, max := detectLimits(points)

	minDistances := calculateMinimumDistances(min, max, points)

	assert.Equal(8, len(minDistances))
}

func TestGetPointIds(t *testing.T) {
	assert := assert.New(t)

	d := distance{}

	assert.Equal("", d.getPointIDs())

	d = distance{PointIds: []int{1}}

	assert.Equal("1", d.getPointIDs())

	d = distance{PointIds: []int{1, 2, 3}}

	assert.Equal("1-2-3", d.getPointIDs())
}

func TestDetectLimits(t *testing.T) {
	assert := assert.New(t)

	points := []point{
		point{X: 0, Y: 0},
		point{X: 10, Y: 0},
		point{X: 0, Y: 10},
		point{X: 10, Y: 10},
		point{X: 5, Y: 5},
	}

	min, max := detectLimits(points)

	assert.Equal(0, min.X)
	assert.Equal(0, min.Y)
	assert.Equal(10, max.X)
	assert.Equal(10, max.Y)
}

func TestParsePoint(t *testing.T) {
	assert := assert.New(t)

	p := parseCoordinate("983, 2341")
	assert.Equal(983, p.X)
	assert.Equal(2341, p.Y)
}
