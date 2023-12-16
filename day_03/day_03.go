package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	a "github.com/et-codes/advent2023-go"
)

const (
	symbols = "!@#$%^&*()_+-=[]{}|;:',<>/?`~"
	numbers = "0123456789"
	test    = "day_03_test_data.txt"
	puzzle  = "day_03_data.txt"
)

func main() {
	fmt.Println(day03(puzzle))
}

func day03(path string) []int {
	lines := a.ReadLines(path)
	regex := regexp.MustCompile(`\d+`)

	partNumberSum := 0

	for i, line := range lines {
		indices := regex.FindAllStringIndex(line, -1)

		for _, index := range indices {
			if hasSymbolNeighbor(lines, i, index) {
				numstr := line[index[0]:index[1]]
				num, _ := strconv.Atoi(numstr)
				partNumberSum += num
			}
		}
	}
	return []int{partNumberSum, 0}
}

func hasSymbolNeighbor(lines []string, row int, index []int) bool {
	var startRow, endRow, startCol, endCol int
	if row > 0 {
		startRow = row - 1
	} else {
		startRow = row
	}

	if row < len(lines)-1 {
		endRow = row + 1
	} else {
		endRow = row
	}

	if index[0] > 0 {
		startCol = index[0] - 1
	} else {
		startCol = 0
	}

	if index[1] < len(lines[row])-1 {
		endCol = index[1] + 1
	} else {
		endCol = index[1]
	}

	for r := startRow; r <= endRow; r++ {
		if strings.ContainsAny(lines[r][startCol:endCol], symbols) {
			return true
		}
	}
	return false
}
