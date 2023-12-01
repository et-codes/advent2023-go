package main

import (
	"fmt"
	"regexp"
	"strconv"

	a "github.com/et-codes/advent2023-go"
)

func main() {
	fmt.Println(day_1())
}

func day_1() int {
	input := a.ReadLines("day_1_data.txt")

	sum := 0

	for _, line := range input {
		var first, last int

		r := regexp.MustCompile(`\d`)
		matches := r.FindAllString(line, -1)

		first, _ = strconv.Atoi(matches[0])
		if len(matches) > 1 {
			last, _ = strconv.Atoi(matches[len(matches) - 1])
		} else {
			last = first
		}
		sum += first * 10 + last
	}

	return sum
}