package main

import (
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/mdelapenya/advent-of-code/io"
)

const regexpAction = "\\[([0-9]+)-([0-9]+)-([0-9]+) ([0-9]+):([0-9]+)\\] (.+)"
const regexpGuard = "\\[([0-9]+)-([0-9]+)-([0-9]+) ([0-9]+):([0-9]+)\\] Guard #([0-9]+)? (.+)"

type date struct {
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
}

type minute struct {
	Asleep bool
	Minute int
}

type minutes struct {
	GuardID int
	Minutes []minute
}

type guard struct {
	ID      int
	actions []guardAction
}

type guardAction struct {
	Text string
	Date date
}

type solution struct {
	MostSleepyGuardID           int
	MinuteMostSlept             int
	MostFrequentlySleepyGuardID int
	MinuteMostFrequentlySlept   int
}

func main() {
	log.Println("Advent of code 2018: Day 4")

	inputFile := "input"

	lines, err := io.ReadLines(inputFile)
	if err != nil {
		panic(err)
	}

	solution := chooseGuard(lines)
	guardID := solution.MostSleepyGuardID
	min := solution.MinuteMostSlept
	frequentGuardID := solution.MostFrequentlySleepyGuardID
	frequentMin := solution.MinuteMostFrequentlySlept

	log.Printf(
		"The selected Guard (%d), multiplied by the minute she is mostly sleeping (%d) is: %d",
		guardID, min, (guardID * min))
	log.Printf(
		"The most frequently sleepy Guard (%d), multiplied by the minute she is mostly sleeping (%d) is: %d",
		frequentGuardID, frequentMin, (frequentGuardID * frequentMin))
}

func checkFrequentMinutesAsleep(m minutes) [60]int {
	mins := [60]int{}

	for i, min := range m.Minutes {
		if min.Asleep {
			mins[i] = 1
		} else {
			mins[i] = 0
		}
	}

	return mins
}

func chooseGuard(lines []string) solution {
	guards := captureGuards(lines)

	timeline := getTimeline(guards)

	minutesAsleepByGuard := map[int]int{}
	frequestMinutesAsleepByGuard := map[int][60]int{}

	for _, days := range timeline {
		minutesAsleep := countMinutesAsleep(days)

		if val, ok := minutesAsleepByGuard[days.GuardID]; ok {
			minutesAsleepByGuard[days.GuardID] = minutesAsleep + val
		} else {
			minutesAsleepByGuard[days.GuardID] = minutesAsleep
		}

		frequentMinutesAsleep := checkFrequentMinutesAsleep(days)

		if val, ok := frequestMinutesAsleepByGuard[days.GuardID]; ok {
			frequestMinutesAsleepByGuard[days.GuardID] = sumArrays(frequentMinutesAsleep, val)
		} else {
			frequestMinutesAsleepByGuard[days.GuardID] = frequentMinutesAsleep
		}
	}

	mostSleepyGuard := 0
	maxMinutesSleep := 0

	for guardID, minutesAsleep := range minutesAsleepByGuard {
		if minutesAsleep > maxMinutesSleep {
			maxMinutesSleep = minutesAsleep
			mostSleepyGuard = guardID
		}
	}

	goldenMinute := getMostProbablyMinute(mostSleepyGuard, timeline)

	highestFrequency := 0
	mostFrequentlySleepyGuardID := 0
	mostFrequentMinuteSleep := 0

	for guardID, frequentMinutesAsleep := range frequestMinutesAsleepByGuard {
		frequentMinute, frequency := getMostFrequentMinuteAndFrequency(guardID, frequentMinutesAsleep)

		if frequency > highestFrequency {
			highestFrequency = frequency
			mostFrequentMinuteSleep = frequentMinute
			mostFrequentlySleepyGuardID = guardID
		}
	}

	return solution{
		MostSleepyGuardID:           mostSleepyGuard,
		MinuteMostSlept:             goldenMinute,
		MostFrequentlySleepyGuardID: mostFrequentlySleepyGuardID,
		MinuteMostFrequentlySlept:   mostFrequentMinuteSleep,
	}
}

func countMinutesAsleep(m minutes) int {
	count := 0

	for _, min := range m.Minutes {
		if min.Asleep {
			count++
		}
	}

	return count
}

func getMostFrequentMinuteAndFrequency(ID int, minutes [60]int) (int, int) {
	mostFrequentMinute := 0
	highestFrequency := 0

	for i, m := range minutes {
		if m > highestFrequency {
			mostFrequentMinute = i
			highestFrequency = m
		}
	}

	return mostFrequentMinute, highestFrequency
}

func getMostProbablyMinute(guardID int, timeline map[string]minutes) int {
	guardMinutes := [][]minute{}

	for _, line := range timeline {
		if line.GuardID == guardID {
			guardMinutes = append(guardMinutes, line.Minutes)
		}
	}

	counts := [60]int{}
	for i := 0; i < len(guardMinutes); i++ {
		for j := 0; j < len(guardMinutes[i]); j++ {
			if guardMinutes[i][j].Asleep {
				counts[j] = counts[j] + 1
			}
		}
	}

	maxMinutesSleep := 0
	mostProbablyMinute := 0

	for i, n := range counts {
		if n > maxMinutesSleep {
			mostProbablyMinute = i
			maxMinutesSleep = n
		}
	}

	return mostProbablyMinute
}

func getTimeline(guards map[int]guard) map[string]minutes {
	timeline := map[string]minutes{}

	for _, g := range guards {
		startMinute := 0
		endMinute := 59

		// actions are sorted chronologically
		for _, action := range g.actions {
			date := action.Date

			key := toString(date.Month) + "-" + toString(date.Day) + " #" + toString(g.ID)

			mins := minutes{GuardID: g.ID}

			if val, ok := timeline[key]; ok {
				mins = val
			} else {
				for i := startMinute; i <= endMinute; i++ {
					m := minute{
						Minute: i, Asleep: false,
					}

					mins.Minutes = append(mins.Minutes, m)
				}
			}

			switch action.Text {
			case "begins shift":
				for j := action.Date.Minute; j <= 59; j++ {
					mins.Minutes[j].Asleep = false
				}
			case "falls asleep":
				for k := action.Date.Minute; k <= 59; k++ {
					mins.Minutes[k].Asleep = true
				}
			case "wakes up":
				for m := action.Date.Minute; m <= 59; m++ {
					mins.Minutes[m].Asleep = false
				}
			}

			timeline[key] = mins
		}
	}

	return timeline
}

// captureGuards lines will be sorted so that the guard actions belongs to the same guard
func captureGuards(lines []string) map[int]guard {
	sort.Strings(lines)

	guards := map[int]guard{}

	for i := 0; i < len(lines); i++ {
		var currentGuard *guard

		line := lines[i]

		isGuard := strings.Contains(line, "Guard #")

		if isGuard {
			g := parseGuard(line)

			if val, ok := guards[g.ID]; ok {
				currentGuard = &val
				currentGuard.actions = append(currentGuard.actions, g.actions[0])
			} else {
				currentGuard = g
			}

			i++

			for i < len(lines) {
				line = lines[i]

				isGuard := strings.Contains(line, "Guard #")

				if !isGuard {
					action := parseAction(line)

					currentGuard.actions = append(currentGuard.actions, *action)

					i++
				} else {
					guards[currentGuard.ID] = *currentGuard

					i--

					break
				}
			}
		}
	}

	return guards
}

func sumArrays(a1 [60]int, a2 [60]int) [60]int {
	sum := [60]int{}

	for i := 0; i < 60; i++ {
		sum[i] = a1[i] + a2[i]
	}

	return sum
}

func parseAction(line string) *guardAction {
	re := regexp.MustCompile(regexpAction)

	groups := re.FindStringSubmatch(line)

	return &guardAction{
		Text: groups[6],
		Date: date{
			Year:   toInt(groups[1]),
			Month:  toInt(groups[2]),
			Day:    toInt(groups[3]),
			Hour:   toInt(groups[4]),
			Minute: toInt(groups[5]),
		},
	}
}

func parseGuard(line string) *guard {
	re := regexp.MustCompile(regexpGuard)

	groups := re.FindStringSubmatch(line)

	guard := guard{
		ID: toInt(groups[6]),
	}

	action := guardAction{
		Text: groups[7],
		Date: date{
			Year:   toInt(groups[1]),
			Month:  toInt(groups[2]),
			Day:    toInt(groups[3]),
			Hour:   toInt(groups[4]),
			Minute: toInt(groups[5]),
		},
	}

	guard.actions = append(guard.actions, action)

	return &guard
}

func toString(i int) string {
	if i < 10 {
		return "0" + strconv.Itoa(i)
	}

	return strconv.Itoa(i)
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)

	return i
}
