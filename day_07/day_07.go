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
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
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

	// Part One
	cardRank['J'] = 11 // Jacks have normal value
	hands = parseHands(data, false)
	sort.SliceStable(hands, compareHands)
	partOne := scoreHands(hands)

	// Part Two
	cardRank['J'] = 1 // Jokers are weakest-ranked card
	hands = parseHands(data, true)
	sort.SliceStable(hands, compareHands)
	partTwo := scoreHands(hands)

	return []int{partOne, partTwo}
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
func parseHands(data []string, useJokers bool) []Hand {
	hands := []Hand{}
	for _, line := range data {
		parts := strings.Split(line, " ")
		bid, _ := strconv.Atoi(parts[1])
		hand := Hand{
			Cards: parts[0],
			Bid:   bid,
		}
		hand.HandType = setHandType(hand, useJokers)

		hands = append(hands, hand)
	}

	return hands
}

// setHandType returns the best possible hand type from the given hand
func setHandType(hand Hand, useJokers bool) int {
	cardCounts := make(map[rune]int)
	for _, card := range hand.Cards {
		cardCounts[card]++
	}

	// Part Two uses Jokers
	if useJokers {
		jokerCount, found := cardCounts['J'] // check if J in hand
		if found {
			// Find the maximum card count
			max := 0
			for card, count := range cardCounts {
				if count >= max && card != 'J' {
					max = count
				}
			}

			// Find highest-ranking card with max card count
			highestCard := '2'
			for card, count := range cardCounts {
				if card != 'J' && count == max && cardRank[card] > cardRank[highestCard] {
					highestCard = card
				}
			}

			// Convert Jokers to highest-ranking max-count cards
			cardCounts[highestCard] += jokerCount
			delete(cardCounts, 'J')
		}
	}

	hasThreeOfAKind := false
	pairs := 0
	for _, count := range cardCounts {
		switch count {
		case 5:
			return fiveOfAKind
		case 4:
			return fourOfAKind
		case 3:
			hasThreeOfAKind = true
		case 2:
			pairs++
		}
	}

	if hasThreeOfAKind && pairs == 1 {
		return fullHouse
	} else if hasThreeOfAKind {
		return threeOfAKind
	} else if pairs == 2 {
		return twoPair
	} else if pairs == 1 {
		return onePair
	}

	return highCard
}
