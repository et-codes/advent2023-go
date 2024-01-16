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
	Edges []int
	Seeds []int
	Maps  [][]Map
}

func main() {
	fmt.Println(day05(test))
}

func day05(path string) []int {
	data := load(path)
	partOne := convert(data, 1)
	partTwo := convert(data, 2)

	return []int{partOne, partTwo}
}

func convert(data *Data, part int) int {
	var locations []int

	if part == 1 {
		locations = make([]int, len(data.Seeds))
		copy(locations, data.Seeds)
	} else {
		locations = make([]int, len(data.Edges))
		copy(locations, data.Edges)
	}

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

	return min(locations)
}

func load(path string) *Data {
	data := &Data{
		Seeds: []int{},
		Maps:  make([][]Map, 7),
	}
	lines := a.ReadLines(path)
	edges := []int{}

	// Get seed numbers
	seeds := strings.Split(lines[0], " ")
	for _, seed := range seeds[1:] {
		s, _ := strconv.Atoi(seed)
		data.Seeds = append(data.Seeds, s)
	}

	// Add edges for seeds
	for i := 0; i < len(data.Seeds); i += 2 {
		edges = append(edges, data.Seeds[i])
		edges = append(edges, data.Seeds[i]+data.Seeds[i+1])
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

		// Add edges if they are in range of seeds
		if inSeedRange(data, dest) {
			edges = append(edges, dest)
		}
		if inSeedRange(data, dest+rng) {
			edges = append(edges, dest+rng)
		}
		if inSeedRange(data, src) {
			edges = append(edges, src)
		}
		if inSeedRange(data, src+rng) {
			edges = append(edges, src+rng)
		}
	}

	data.Edges = edges
	return data
}

func inSeedRange(data *Data, num int) bool {
	for i := 0; i < len(data.Seeds); i += 2 {
		start := data.Seeds[i]
		end := start + data.Seeds[i+1]
		if num >= start && num <= end {
			return true
		}
	}
	return false
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

func max(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
