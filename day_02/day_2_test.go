package main

import (
	"reflect"
	"testing"
)

func TestDay2(t *testing.T) {
	tests := map[string]struct {
		source string
		want   []int
	}{
		"with test data":   {test, []int{8, 2286}},
		"with puzzle data": {puzzle, []int{2776, 68638}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := day_2(test.source)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}

}
