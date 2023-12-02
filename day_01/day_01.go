package main

import (
	"fmt"
	"regexp"
	"strconv"

	a "github.com/et-codes/advent2023-go"
)

var (
	rePart1 = regexp.MustCompile(`\d`)
	rePart2 = regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)
	numbers = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
)

func main() {
	fmt.Println(day01(rePart1), day01(rePart2)) // 54953, 53868
}

func day01(r *regexp.Regexp) int {
	input := a.ReadLines("day_1_data.txt")

	sum := 0

	for _, line := range input {
		var first, last int

		match := r.FindString(line)

		first, _ = strconv.Atoi(match)
		if first == 0 {
			first = numbers[match]
		}

		for i := len(line) - 1; i >= 0; i-- {
			match := r.FindString(line[i:])
			if match == "" {
				continue
			}

			last, _ = strconv.Atoi(match)
			if last == 0 {
				last = numbers[match]
			}
			break
		}

		sum += first*10 + last
	}

	return sum
}
