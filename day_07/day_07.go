package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	a "github.com/et-codes/advent2023-go"
)

const (
	test   = "day_07_test_data.txt" // sample data from puzzle description
	puzzle = "day_07_data.txt"      // actual puzzle data
)

var cardRank = map[rune]int{
	'A': 12,
	'K': 11,
	'Q': 10,
	'J': 9,
	'T': 8,
	'9': 7,
	'8': 6,
	'7': 5,
	'6': 4,
	'5': 3,
	'4': 2,
	'3': 1,
	'2': 0,
}

// Hand types
const (
	highCard = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

type Hand struct {
	Cards    string
	Bid      int
	HandType int
}

var hands []Hand

func main() {
	fmt.Println(day_07(test))
}

func day_07(path string) []int {
	data := a.ReadLines(path)
	hands = parseHands(data)
	sort.SliceStable(hands, compareHands)

	partOne := scoreHands(hands)

	return []int{partOne, 0}
}

// scoreHands returns the total winnings of the hands
func scoreHands(hands []Hand) int {
	winnings := 0
	for i, hand := range hands {
		winnings += (i + 1) * hand.Bid
	}
	return winnings
}

// compareHands returns true if hands[i] < hands[j]
func compareHands(i, j int) bool {
	// If HandTypes are different, rank based on HandType
	if hands[i].HandType > hands[j].HandType {
		return false
	} else if hands[i].HandType < hands[j].HandType {
		return true
	}

	// If HandTypes are the same, rank based on highest card in order
	for k, card := range hands[i].Cards {
		if cardRank[card] > cardRank[rune(hands[j].Cards[k])] {
			return false
		} else if cardRank[card] < cardRank[rune(hands[j].Cards[k])] {
			return true
		}
	}

	return false
}

// parseHands returns a slice of Hand structs from the input data
func parseHands(data []string) []Hand {
	hands := []Hand{}
	for _, line := range data {
		parts := strings.Split(line, " ")
		bid, _ := strconv.Atoi(parts[1])
		hand := Hand{
			Cards: parts[0],
			Bid:   bid,
		}

		setHandType(&hand)

		hands = append(hands, hand)
	}

	return hands
}

// setHandType sets the best possible hand type from the given hand
func setHandType(hand *Hand) {
	cardCounts := make(map[rune]int)
	for _, card := range hand.Cards {
		cardCounts[card]++
	}

	hasThreeOfAKind := false
	pairs := 0
	for _, count := range cardCounts {
		switch count {
		case 5:
			hand.HandType = fiveOfAKind
			return
		case 4:
			hand.HandType = fourOfAKind
			return
		case 3:
			hasThreeOfAKind = true
		case 2:
			pairs++
		}
	}

	if hasThreeOfAKind && pairs == 1 {
		hand.HandType = fullHouse
	} else if hasThreeOfAKind {
		hand.HandType = threeOfAKind
	} else if pairs == 2 {
		hand.HandType = twoPair
	} else if pairs == 1 {
		hand.HandType = onePair
	} else {
		hand.HandType = highCard
	}
}
