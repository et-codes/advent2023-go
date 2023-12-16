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
		"with test data":   {test, []int{4361, 467835}},
		"with puzzle data": {puzzle, []int{533784, 0}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := day03(test.source)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}

}
