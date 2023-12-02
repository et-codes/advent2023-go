package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	a "github.com/et-codes/advent2023-go"
)

var bag = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	fmt.Println(day_2("day_2_data.txt"))
}

func day_2(path string) int {
	games := a.ReadLines(path)
	result := 0

	for _, game := range games {
		re := regexp.MustCompile(`\d+`)
		gameNumber, _ := strconv.Atoi(re.FindString(game))

		rounds := strings.Split(game, ": ")
		rounds = strings.Split(rounds[1], "; ")

		possible := true
		for _, round := range rounds {
			cubes := strings.Split(round, ", ")

			for _, cube := range cubes {
				re = regexp.MustCompile(`(\d+|red|green|blue)`)
				matches := re.FindAllString(cube, -1)
				qty, _ := strconv.Atoi(matches[0])
				color := matches[1]
				if qty > bag[color] {
					possible = false
					break
				}
			}
		}

		if possible {
			result += gameNumber
		}
	}
	return result
}
