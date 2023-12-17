package main

import (
	"reflect"
	"testing"
)

func TestDay(t *testing.T) {
	tests := map[string]struct {
		source string
		want   []int
	}{
		"with test data":   {test, []int{13, 0}},
		"with puzzle data": {puzzle, []int{21821, 0}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := day04(test.source)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}

}
