package main

import (
	"fmt"
	"strconv"
	"strings"

	a "github.com/et-codes/advent2023-go"
)

const (
	test   = "day_06_test_data.txt" // sample data from puzzle description
	puzzle = "day_06_data.txt"      // actual puzzle data
)

type Race struct {
	Time   int // race duration in ms
	Record int // record distance in mm
}

func main() {
	fmt.Println(day_06(puzzle))
}

func day_06(path string) []int {
	data := a.ReadLines(path)
	races := parseRaces(data)

	return []int{partOne(races), 0}
}

func partOne(races []Race) int {
	margin := 1

	for _, race := range races {
		raceMargin := 0
		for speed := 1; speed < race.Time; speed++ {
			duration := race.Time - speed
			distance := duration * speed

			if distance > race.Record {
				raceMargin++
			}
		}
		margin *= raceMargin
	}

	return margin
}

func parseRaces(data []string) []Race {
	races := []Race{}

	timeStrings := strings.Split(data[0], " ")
	for _, time := range timeStrings[1:] {
		t, _ := strconv.Atoi(time)
		if t != 0 {
			races = append(races, Race{Time: t})
		}
	}

	distStrings := strings.Split(data[1], " ")
	i := 0
	for _, dist := range distStrings[1:] {
		d, _ := strconv.Atoi(dist)
		if d != 0 {
			races[i].Record = d
			i++
		}
	}

	return races
}
