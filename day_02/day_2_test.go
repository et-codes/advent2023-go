package main

import "testing"

func TestDay2(t *testing.T) {
	tests := map[string]struct {
		source string
		want   int
	}{
		"with test data":   {"day_2_test_data.txt", 8},
		"with puzzle data": {"day_2_data.txt", 2776},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := day_2(test.source)
			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}

}
