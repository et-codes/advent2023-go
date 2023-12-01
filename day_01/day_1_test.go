package main

import "testing"

func TestDay1(t *testing.T) {
	want := 54953
	got := day_1()
	if day_1() != want {
		t.Errorf("wanted %q, got %q", want, got)
	}
}