package main

import (
	"fmt"
	"strings"

	"github.com/mdelapenya/advent-of-code/io"
	aocstrings "github.com/mdelapenya/advent-of-code/strings"
)

type Report []int

func main() {
	reports := mustReadInput("input.txt")

	safeReports := 0
	for _, report := range reports {
		if isSafe(report) {
			safeReports++
		}
	}

	fmt.Println("safe reports:", safeReports)

	safeTolerantReports := 0
	for _, report := range reports {
		if isSafeWithDampener(report) {
			safeTolerantReports++
		}
	}
	fmt.Println("safe tolerant reports:", safeTolerantReports)
}

func isSafeWithDampener(report Report) bool {
	if isSafe(report) {
		return true
	}

	// try removing each level one by one
	for i := 0; i < len(report); i++ {
		newReport := make([]int, 0, len(report)-1)
		newReport = append(newReport, report[:i]...)
		newReport = append(newReport, report[i+1:]...)

		if isSafe(newReport) {
			return true
		}
	}

	return false
}

func isSafe(report Report) bool {
	var directionFn func(i int, j int) bool

	if report[0] >= report[1] {
		directionFn = checkDecreasing
	} else if report[0] <= report[1] {
		directionFn = checkIncreasing
	} else {
		directionFn = func(_ int, _ int) bool {
			return false // no change means not safe
		}
	}

	for i := 0; i < len(report)-1; i++ {
		if !directionFn(report[i], report[i+1]) {
			return false
		}
	}

	return true
}

func checkDecreasing(i int, j int) bool {
	return i >= j && i-j <= 3 && i-j >= 1
}

func checkIncreasing(i int, j int) bool {
	return i <= j && j-i <= 3 && j-i >= 1
}

func mustParseReport(line string) Report {
	parts := strings.Fields(line)

	report := make(Report, len(parts))
	for i, part := range parts {
		report[i] = aocstrings.MustParseInt(part)
	}

	return report
}

func mustReadInput(path string) []Report {
	lines, err := io.ReadLines(path)
	if err != nil {
		panic(err)
	}

	reports := make([]Report, len(lines))
	for i, line := range lines {
		reports[i] = mustParseReport(line)
	}

	return reports
}
