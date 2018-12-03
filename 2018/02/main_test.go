package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func extractB(s1 string, s2 string) bool {
	b, _ := diffExactlyByOne(s1, s2)

	return b
}

func extractS(s1 string, s2 string) string {
	_, s := diffExactlyByOne(s1, s2)

	return s
}
func TestDiffFindNear(t *testing.T) {
	assert := assert.New(t)

	lines, _ := readLines("input")
	assert.Equal("rteotyxzbodglnpkudawhijsc", findNearIDs(lines))
}

func TestDiffExactlyByOne(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(false, extractB("abcde", "abcde"))
	assert.Equal(false, extractB("abcde", "fghij"))
	assert.Equal(false, extractB("abcde", "axcye"))
	assert.Equal(true, extractB("fghij", "fguij"))
	assert.Equal("fgij", extractS("fghij", "fguij"))
}

func TestFindIDs(t *testing.T) {
	assert := assert.New(t)

	lines, _ := readLines("input")
	assert.Equal(6150, findIDs(lines))
}

func TestHasThree(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, hasExactlyThree("abcdef"))
	assert.Equal(1, hasExactlyThree("bababc"))
	assert.Equal(0, hasExactlyThree("abbcde"))
	assert.Equal(1, hasExactlyThree("abcccd"))
	assert.Equal(0, hasExactlyThree("aabcdd"))
	assert.Equal(0, hasExactlyThree("abcdee"))
	assert.Equal(1, hasExactlyThree("ababab"))
}

func TestHasTwo(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, hasExactlyTwo("abcdef"))
	assert.Equal(1, hasExactlyTwo("bababc"))
	assert.Equal(1, hasExactlyTwo("abbcde"))
	assert.Equal(0, hasExactlyTwo("abcccd"))
	assert.Equal(1, hasExactlyTwo("aabcdd"))
	assert.Equal(1, hasExactlyTwo("abcdee"))
	assert.Equal(0, hasExactlyTwo("ababab"))
}
