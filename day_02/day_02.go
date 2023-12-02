package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	a "github.com/et-codes/advent2023-go"
)

type bag map[string]int

type game struct {
	gameNum int
	rounds  []bag
}

const (
	test   = "day_2_test_data.txt"
	puzzle = "day_2_data.txt"
)

func main() {
	fmt.Println(day02(puzzle))
}

func day02(path string) []int {
	games := a.ReadLines(path)

	var result1, result2 int
	bag1 := bag{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for _, g := range games {
		game := parseGame(g)
		possible := true
		bag2 := bag{}

		for _, round := range game.rounds {
			for color, qty := range round {
				if qty > bag1[color] {
					possible = false
				}
				if qty > bag2[color] {
					bag2[color] = qty
				}
			}
		}
		if possible {
			result1 += game.gameNum
		}
		result2 += power(bag2)
	}
	return []int{result1, result2}
}

func parseGame(in string) game {
	gameMatch := regexp.MustCompile(`Game (\d+): `)
	roundMatch := regexp.MustCompile(`(\d+) (red|green|blue);?`)

	gameNum, _ := strconv.Atoi(gameMatch.FindStringSubmatch(in)[1])
	result := game{gameNum: gameNum}

	round := bag{}
	for roundMatch.MatchString(in) {
		indices := roundMatch.FindSubmatchIndex([]byte(in))
		match := roundMatch.FindStringSubmatch(in)
		round[match[2]], _ = strconv.Atoi(match[1])
		if strings.Contains(match[0], ";") {
			result.rounds = append(result.rounds, round)
			round = bag{}
		}
		in = in[indices[1]:]
	}
	result.rounds = append(result.rounds, round)
	return result
}

func power(bag bag) int {
	return bag["red"] * bag["blue"] * bag["green"]
}
