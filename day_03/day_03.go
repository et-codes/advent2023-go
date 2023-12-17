package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	a "github.com/et-codes/advent2023-go"
)

const (
	allSymbols = "!@#$%^&*()_+-=[]{}|;:',<>/?`~"
	test       = "day_03_test_data.txt"
	puzzle     = "day_03_data.txt"
)

func main() {
	fmt.Println(day03(puzzle))
}

type symbol struct {
	symbol rune
	row    int
	col    int
}

func day03(path string) []int {
	lines := a.ReadLines(path)
	gears := map[string][]int{}

	regex := regexp.MustCompile(`\d+`)

	partNumberSum := 0

	for i, line := range lines {
		indices := regex.FindAllStringIndex(line, -1)

		for _, index := range indices {
			neighbors := findSymbolNeighbors(lines, i, index)
			if len(neighbors) > 0 {
				numstr := line[index[0]:index[1]]
				num, _ := strconv.Atoi(numstr)
				partNumberSum += num

				for _, n := range neighbors {
					if n.symbol == '*' {
						key := fmt.Sprintf("%d,%d", n.row, n.col)
						gears[key] = append(gears[key], num)
					}
				}
			}
		}
	}

	// Part 2
	gearRatioSum := 0
	for _, parts := range gears {
		if len(parts) == 2 {
			gearRatioSum += parts[0] * parts[1]
		}
	}

	return []int{partNumberSum, gearRatioSum}
}

func findSymbolNeighbors(lines []string, row int, index []int) []symbol {
	symbols := []symbol{}
	startRow := max(row-1, 0)
	endRow := min(len(lines)-1, row+1)

	startCol := max(index[0]-1, 0)
	endCol := min(len(lines[row])-1, index[1])

	for r := startRow; r <= endRow; r++ {
		for c := startCol; c <= endCol; c++ {
			if strings.ContainsAny(string(lines[r][c]), allSymbols) {
				symbols = append(symbols, symbol{
					symbol: rune(lines[r][c]),
					row:    r,
					col:    c,
				})
			}
		}
	}
	return symbols
}
