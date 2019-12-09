package main

import (
	"testing"

	"github.com/mdelapenya/advent-of-code/io"
	"github.com/stretchr/testify/assert"
)

func TestCaptureGuardActions(t *testing.T) {
	assert := assert.New(t)

	lines, _ := io.ReadLines("input")
	guards := captureGuards(lines)

	guard := guards[2179]
	actions := guard.actions

	assert.Equal(2179, guard.ID)
	assert.Equal(63, len(actions))

	action := actions[60]

	//[1518-11-11 00:04] Guard #2179 begins shift
	assert.Equal(1518, action.Date.Year)
	assert.Equal(11, action.Date.Month)
	assert.Equal(11, action.Date.Day)
	assert.Equal(0, action.Date.Hour)
	assert.Equal(4, action.Date.Minute)
	assert.Equal("begins shift", action.Text)
}

func TestChooseGuard(t *testing.T) {
	assert := assert.New(t)

	lines, _ := io.ReadLines("input")
	solution := chooseGuard(lines)

	guardID := solution.MostSleepyGuardID
	min := solution.MinuteMostSlept

	assert.Equal(1021, guardID)
	assert.Equal(30, min)

	frequentGuardID := solution.MostFrequentlySleepyGuardID
	frequentMin := solution.MinuteMostFrequentlySlept

	assert.Equal(3331, frequentGuardID)
	assert.Equal(41, frequentMin)
}

func TestCountMinutesAsleep(t *testing.T) {
	assert := assert.New(t)

	mins := minutes{}

	for i := 0; i <= 59; i++ {
		m := minute{
			Minute: i, Asleep: false,
		}

		mins.Minutes = append(mins.Minutes, m)
	}

	assert.Equal(0, countMinutesAsleep(mins))

	mins.Minutes[2].Asleep = true
	mins.Minutes[20].Asleep = true
	mins.Minutes[40].Asleep = true

	assert.Equal(3, countMinutesAsleep(mins))
}

func TestGetTimeline(t *testing.T) {
	assert := assert.New(t)

	lines, _ := io.ReadLines("input")
	guards := captureGuards(lines)
	timeline := getTimeline(guards)

	mins := timeline["02-14 #587"]

	assert.Equal(587, mins.GuardID)

	for i := 0; i < 36; i++ {
		assert.False(mins.Minutes[i].Asleep)
	}
	for i := 36; i < 54; i++ {
		assert.True(mins.Minutes[i].Asleep)
	}
	for i := 54; i < 60; i++ {
		assert.False(mins.Minutes[i].Asleep)
	}

	mins1 := timeline["02-11 #631"]

	assert.Equal(631, mins1.GuardID)

	for i := 0; i < 3; i++ {
		assert.False(mins1.Minutes[i].Asleep)
	}
	for i := 3; i < 19; i++ {
		assert.True(mins1.Minutes[i].Asleep)
	}
	for i := 19; i < 60; i++ {
		assert.False(mins1.Minutes[i].Asleep)
	}

	mins2 := timeline["09-01 #3331"]

	assert.Equal(3331, mins2.GuardID)

	for i := 0; i < 1; i++ {
		assert.False(mins2.Minutes[i].Asleep)
	}
	for i := 1; i < 12; i++ {
		assert.True(mins2.Minutes[i].Asleep)
	}
	for i := 12; i < 31; i++ {
		assert.False(mins2.Minutes[i].Asleep)
	}
	for i := 31; i < 42; i++ {
		assert.True(mins2.Minutes[i].Asleep)
	}
	for i := 42; i < 54; i++ {
		assert.False(mins2.Minutes[i].Asleep)
	}
	for i := 54; i < 58; i++ {
		assert.True(mins2.Minutes[i].Asleep)
	}
	for i := 58; i < 60; i++ {
		assert.False(mins2.Minutes[i].Asleep)
	}
}

func TestMostProbablyAsleep(t *testing.T) {
	assert := assert.New(t)

	lines, _ := io.ReadLines("input")
	guards := captureGuards(lines)
	timeline := getTimeline(guards)

	minute := getMostProbablyMinute(587, timeline)

	assert.Equal(48, minute)
}

func TestRegexpAction(t *testing.T) {
	assert := assert.New(t)

	action := parseAction("[1518-10-19 00:22] wakes up")

	assert.Equal(1518, action.Date.Year)
	assert.Equal(10, action.Date.Month)
	assert.Equal(19, action.Date.Day)
	assert.Equal(0, action.Date.Hour)
	assert.Equal(22, action.Date.Minute)
	assert.Equal("wakes up", action.Text)
}

func TestRegexpGuard(t *testing.T) {
	assert := assert.New(t)

	guard := parseGuard("[1518-10-19 00:22] Guard #123 begins shift")

	action := guard.actions[0]

	assert.Equal(123, guard.ID)
	assert.Equal(1518, action.Date.Year)
	assert.Equal(10, action.Date.Month)
	assert.Equal(19, action.Date.Day)
	assert.Equal(0, action.Date.Hour)
	assert.Equal(22, action.Date.Minute)
	assert.Equal("begins shift", action.Text)
}
