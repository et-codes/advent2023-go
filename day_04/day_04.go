package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	a "github.com/et-codes/advent2023-go"
)

const (
	test   = "day_04_test_data.txt"
	puzzle = "day_04_data.txt"
)

type card struct {
	numbers []int
	winning []int
}

func day04(path string) []int {
	totalScore := 0
	cards := getNumbers(path)
	for _, card := range cards {
		totalScore += calcScore(card)
	}

	return []int{totalScore, 0}
}

func calcScore(card card) int {
	winningNums := 0
	for _, num := range card.numbers {
		if contains(card.winning, num) {
			winningNums++
		}
	}
	return int(math.Pow(float64(2), float64(winningNums-1)))
}

func contains(nums []int, num int) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
	}
	return false
}

func getNumbers(path string) []card {
	lines := a.ReadLines(path)
	cards := []card{}

	re := regexp.MustCompile(`.*: *(\d*) *| *(\d*)`)
	for _, line := range lines {
		line = strings.Split(line, ":")[1]
		matches := re.FindAllString(line, -1)
		winning := false
		card := card{}
		for _, match := range matches {
			match := strings.TrimSpace(match)
			if match == "" {
				winning = true
				continue
			}
			num, _ := strconv.Atoi(match)
			if !winning {
				card.numbers = append(card.numbers, num)
			} else {
				card.winning = append(card.winning, num)
			}
		}
		cards = append(cards, card)
	}
	return cards
}

func main() {
	fmt.Println(day04(puzzle))
}
