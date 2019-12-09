package main

import (
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/mdelapenya/advent-of-code/2018/io"
)

const regexpCoordinates = "([0-9]+), ([0-9]+)"

func main() {
	log.Println("Advent of code 2018: Day 6")

	inputFile := "input"

	lines, err := io.ReadLines(inputFile)
	if err != nil {
		panic(err)
	}

	area := getGreaterArea(lines)
	log.Printf("The greater area is %d", area)
}

type point struct {
	ID int
	X  int
	Y  int
}

func (p point) getManhattanDistance(p1 point) distance {
	xDistance := abs(p1.X - p.X)
	yDistance := abs(p1.Y - p.Y)

	return distance{
		From:     p,
		PointIds: []int{p1.ID},
		Distance: (xDistance + yDistance),
	}
}

func (p point) isEdge(points []point) bool {
	greaterX := false
	greaterY := false
	smallerX := false
	smallerY := false

	for _, p1 := range points {
		if p.X == p1.X && p.Y == p1.Y {
			continue
		}

		if p.X < p1.X {
			greaterX = false
		} else if p.X > p1.X {
			smallerX = false
		} else {
			greaterX = true
			smallerX = true
		}

		if p.Y < p1.Y {
			greaterY = false
		} else if p.Y > p1.Y {
			smallerY = false
		} else {
			greaterY = true
			smallerY = true
		}
	}

	if greaterX || greaterY || smallerX || smallerY {
		return true
	}

	return false
}

type distance struct {
	From     point
	Distance int
	PointIds []int
}

func (d distance) getPointIDs() string {
	ids := ""

	for _, p := range d.PointIds {
		ids += strconv.Itoa(p) + "-"
	}

	ids = strings.TrimRight(ids, "-")
	return ids
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calculateMinimumDistances(min point, max point, points []point) [][]string {
	log.Printf("Min/Max: (%d, %d) - (%d, %d)", min.X, min.Y, max.X, max.Y)
	horizontalRange := abs(max.X - min.X)
	verticalRange := abs(max.Y - min.Y)

	distances := make([][]string, horizontalRange+1)

	h := 0

	for i := min.X; i <= max.X; i++ {
		distances[h] = make([]string, verticalRange+1)
		v := 0

		for j := min.Y; j <= max.Y; j++ {
			closestPoint := getClosest(point{X: i, Y: j}, points)

			distances[h][v] = closestPoint.getPointIDs()

			v++
		}

		h++
	}

	return distances
}

func countDistances(minDistances [][]string) map[string]int {
	distanceCounts := map[string]int{}

	for i := 0; i < len(minDistances); i++ {
		for j := 0; j < len(minDistances[0]); j++ {
			distance := minDistances[i][j]

			if strings.Contains(distance, "-") {
				continue
			}

			if val, ok := distanceCounts[distance]; ok {
				distanceCounts[distance] = val + 1
			} else {
				distanceCounts[distance] = 1
			}
		}
	}

	return distanceCounts
}

func detectLimits(points []point) (point, point) {
	minX := math.MaxInt32
	maxX := 0

	minY := math.MaxInt32
	maxY := 0

	for _, p := range points {
		if p.X > maxX {
			maxX = p.X
		}
		if p.X < minX {
			minX = p.X
		}

		if p.Y > maxY {
			maxY = p.Y
		}
		if p.Y < minY {
			minY = p.Y
		}
	}

	return point{X: minX, Y: minY}, point{X: maxX, Y: maxY}
}

func getClosest(p point, points []point) distance {
	minDistance := math.MaxInt32
	closest := distance{}

	for _, point := range points {
		distance := p.getManhattanDistance(point)

		if distance.Distance < minDistance {
			minDistance = distance.Distance
			closest.Distance = distance.Distance
			closest.From = p
			closest.PointIds = []int{point.ID}
		} else if distance.Distance == minDistance {
			closest.PointIds = append(closest.PointIds, point.ID)
		}
	}

	return closest
}

func getGreaterArea(lines []string) int {
	points := getPoints(lines)

	min, max := detectLimits(points)

	minDistances := calculateMinimumDistances(min, max, points)

	distanceCounts := countDistances(minDistances)

	maxArea := 0
	greatestID := ""

	for k, v := range distanceCounts {
		if v > maxArea {
			maxArea = v
			greatestID = k
		}
	}

	log.Printf("The greatest are is for %s, resulting in %d", greatestID, maxArea)

	return maxArea
}

func getPoints(lines []string) []point {
	points := []point{}

	for i, line := range lines {
		p := parseCoordinate(line)
		p.ID = i

		points = append(points, p)
	}

	return points
}

func parseCoordinate(line string) point {
	re := regexp.MustCompile(regexpCoordinates)

	groups := re.FindStringSubmatch(line)

	return point{
		X: toInt(groups[1]),
		Y: toInt(groups[2]),
	}
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)

	return i
}
