package main

import (
	"fmt"
	"strconv"
	"strings"

	a "github.com/et-codes/advent2023-go"
)

const (
	test   = "day_05_test_data.txt"
	puzzle = "day_05_data.txt"
)

type Map struct {
	DestRangeStart int
	SrcRangeStart  int
	RangeLength    int
}

type Data struct {
	Seeds []int
	Maps  [][]Map
}

func main() {
	fmt.Println(day05(puzzle))
}

func day05(path string) []int {
	data := load(path)
	locations := convert(data)
	return []int{min(locations), 0}
}

func min(nums []int) int {
	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	return min
}

func convert(data *Data) []int {
	locations := make([]int, len(data.Seeds))
	copy(locations, data.Seeds)

	for i := range locations {
		for _, maps := range data.Maps {
			for _, m := range maps {
				if locations[i] >= m.SrcRangeStart && locations[i] <= m.SrcRangeStart+m.RangeLength {
					locations[i] = locations[i] - m.SrcRangeStart + m.DestRangeStart
					break
				}
			}
		}
	}

	return locations
}

func load(path string) *Data {
	data := &Data{
		Seeds: []int{},
		Maps:  make([][]Map, 7),
	}
	lines := a.ReadLines(path)

	// Get seed numbers
	seeds := strings.Split(lines[0], " ")
	for _, seed := range seeds[1:] {
		s, _ := strconv.Atoi(seed)
		data.Seeds = append(data.Seeds, s)
	}

	// Load maps
	i := 0
	for _, line := range lines[3:] {
		if line == "" {
			continue
		} else if strings.HasSuffix(line, "map:") {
			i++
			continue
		}
		vals := strings.Split(line, " ")
		dest, _ := strconv.Atoi(vals[0])
		src, _ := strconv.Atoi(vals[1])
		rng, _ := strconv.Atoi(vals[2])
		m := Map{dest, src, rng}
		data.Maps[i] = append(data.Maps[i], m)
	}

	return data
}
