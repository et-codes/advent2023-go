package main

import "testing"

func TestDay1(t *testing.T) {
	want := "Welcome to Advent of Code 2023!"
	got := day_1()
	if day_1() != want {
		t.Errorf("wanted %q, got %q", want, got)
	}
}