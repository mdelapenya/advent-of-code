package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/mdelapenya/advent-of-code/2018/io"
)

func main() {
	log.Println("Advent of code 2018: Day 3")

	inputFile := "input"

	lines, err := io.ReadLines(inputFile)
	if err != nil {
		panic(err)
	}

	log.Printf("Inches within two or more claims: %v", countClaims(lines))
}

type square struct {
	ID     int
	Left   int
	Top    int
	Width  int
	Height int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (s square) checkOverlap(plotter [][]int) bool {
	for i := s.Left; i < s.totalWidth(); i++ {
		for j := s.Top; j < s.totalHeight(); j++ {
			if plotter[i][j] != 1 {
				return false
			}
		}
	}

	return true
}

func (s square) plot(plotter [][]int) {
	for i := s.Left; i < s.totalWidth(); i++ {
		for j := s.Top; j < s.totalHeight(); j++ {
			if plotter[i][j] == 0 {
				plotter[i][j] = 1
			} else {
				current := plotter[i][j]
				plotter[i][j] = current + 1
			}
		}
	}
}

func (s square) totalHeight() int {
	return s.Top + s.Height
}

func (s square) totalWidth() int {
	return s.Left + s.Width
}

func countClaims(lines []string) int {
	maxWidth := 0
	maxHeight := 0
	squares := []square{}

	for _, linei := range lines {
		square := parse(linei)

		squares = append(squares, square)

		if square.totalHeight() >= maxHeight {
			maxHeight = square.totalHeight()
		}
		if square.totalWidth() >= maxWidth {
			maxWidth = square.totalWidth()
		}
	}

	var plotter = initPlotter(maxWidth, maxHeight)

	for _, square := range squares {
		square.plot(plotter)
	}

	collitions := 0
	claims := 0
	zeros := 0
	for i, row := range plotter {
		for j := range row {
			value := plotter[i][j]
			if value > 1 {
				collitions++
			} else if value == 1 {
				claims++
			} else {
				zeros++
			}
		}
	}

	log.Printf("Using Matrix of (%v, %v): %d.", maxWidth, maxHeight, maxHeight*maxHeight)
	log.Printf("There were %d claims.", claims)

	for _, square := range squares {
		clean := square.checkOverlap(plotter)
		if clean {
			log.Printf("The only one square with no overlaps is: #%d", square.ID)
		}
	}

	return collitions
}

func initPlotter(w int, h int) [][]int {
	var plotter = make([][]int, w)
	for i := range plotter {
		plotter[i] = make([]int, h)
		for j := range plotter[i] {
			plotter[i][j] = 0
		}
	}

	return plotter
}

func parse(line string) square {
	// capture ID
	s := strings.Split(line, "@")

	id := strings.TrimSuffix(s[0], " ")
	id = strings.TrimPrefix(id, "#")

	i, _ := strconv.Atoi(id)

	// capture paddings
	paddings := strings.Split(s[1], ":")

	padding := strings.TrimSpace(paddings[0])

	pad := strings.Split(padding, ",")

	left, _ := strconv.Atoi(pad[0])
	top, _ := strconv.Atoi(pad[1])

	// capture dimensions
	dimensions := strings.TrimSpace(paddings[1])

	dim := strings.Split(dimensions, "x")

	width, _ := strconv.Atoi(dim[0])
	height, _ := strconv.Atoi(dim[1])

	return square{
		ID:     i,
		Left:   left,
		Top:    top,
		Width:  width,
		Height: height,
	}
}
