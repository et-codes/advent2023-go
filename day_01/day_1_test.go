package main

import "testing"

func TestDay1(t *testing.T) {
	want := 54953
	got := day_1(rePart1)
	if got != want {
		t.Errorf("Part 1: wanted %d, got %d", want, got)
	}

	want = 53868
	got = day_1(rePart2)
	if got != want {
		t.Errorf("Part 2: wanted %d, got %d", want, got)
	}
}
