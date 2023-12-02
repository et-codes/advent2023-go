package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	a "github.com/et-codes/advent2023-go"
)

type bag map[string]int

const (
	test   = "day_2_test_data.txt"
	puzzle = "day_2_data.txt"
)

var (
	bag1 = bag{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
)

func main() {
	fmt.Println(day_2(puzzle))
}

func day_2(path string) []int {
	games := a.ReadLines(path)

	var result1, result2 int

	for _, game := range games {
		re := regexp.MustCompile(`\d+`)
		gameNumber, _ := strconv.Atoi(re.FindString(game))

		rounds := strings.Split(game, ": ")
		rounds = strings.Split(rounds[1], "; ")

		bag2 := bag{}
		possible := true
		for _, round := range rounds {
			cubes := strings.Split(round, ", ")
			for _, cube := range cubes {
				re = regexp.MustCompile(`(\d+|red|green|blue)`)
				matches := re.FindAllString(cube, -1)
				qty, _ := strconv.Atoi(matches[0])
				color := matches[1]

				if qty > bag2[color] {
					bag2[color] = qty
				}

				if qty > bag1[color] {
					possible = false
				}
			}
		}

		if possible {
			result1 += gameNumber
		}

		result2 += power(bag2)
	}
	return []int{result1, result2}
}

func power(bag bag) int {
	return bag["red"] * bag["blue"] * bag["green"]
}
