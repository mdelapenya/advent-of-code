package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const testInput = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestIsSafe(t *testing.T) {
	var reports []Report

	for _, line := range strings.Split(testInput, "\n") {
		reports = append(reports, mustParseReport(line))
	}

	require.True(t, isSafe(reports[0]))
	require.False(t, isSafe(reports[1]))
	require.False(t, isSafe(reports[2]))
	require.False(t, isSafe(reports[3]))
	require.False(t, isSafe(reports[4]))
	require.True(t, isSafe(reports[5]))
}
