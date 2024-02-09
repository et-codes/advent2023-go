package main

import (
	"fmt"
	"regexp"

	a "github.com/et-codes/advent2023-go"
)

const (
	test   = "day_08_test_data.txt" // sample data from puzzle description
	puzzle = "day_08_data.txt"      // actual puzzle data
)

type Node struct {
	Left  string
	Right string
}

type Network map[string]Node

func main() {
	fmt.Println(day_08(test))
}

func day_08(path string) []int {
	data := a.ReadLines(path)
	instructions, network := createNetwork(data)
	steps := traverseNetwork(instructions, network)

	return []int{steps}
}

func traverseNetwork(instructions string, network Network) int {
	var (
		zzzFound           = false
		steps              = 0
		pointer            = 0
		instructionsLength = len(instructions)
		node               = network["AAA"]
		next               string
	)

	for !zzzFound {
		switch instructions[pointer] {
		case 'L':
			next = node.Left
		case 'R':
			next = node.Right
		}

		if next == "ZZZ" {
			zzzFound = true
		}

		steps++
		pointer = (pointer + 1) % instructionsLength

		node = network[next]
	}

	return steps
}

func createNetwork(data []string) (string, Network) {
	instructions := data[0]
	network := make(Network)
	re := regexp.MustCompile(`(...) = \((...), (...)\)`)
	for _, line := range data[2:] {
		matches := re.FindAllStringSubmatch(line, -1)
		network[matches[0][1]] = Node{
			Left:  matches[0][2],
			Right: matches[0][3],
		}
	}

	return instructions, network
}
