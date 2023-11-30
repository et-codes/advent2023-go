package advent2023go_test

import (
	"testing"

	a "github.com/et-codes/advent2023-go"
)

func TestGetDataFromFile(t *testing.T) {
	got := a.ReadLines("README.md")
	want := "# Advent of Code 2023"
	if len(got) == 0 || got[0] != want {
		t.Errorf("wanted %v, got %v", want, got[2])
	}
}