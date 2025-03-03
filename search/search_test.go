package search

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	inputArray []int
	inputValue int
	found      bool
	index      int
}{
	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		3,
		true,
		2,
	},
	{
		[]int{99, 105, 120, 150, 200, 300, 450, 600, 900},
		900,
		true,
		8,
	},
	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		18,
		false,
		0,
	},
}

func TestBinarySearch(t *testing.T) {
	for _, test := range testCases {
		t.Run(fmt.Sprintf("%v", test.inputArray), func(t *testing.T) {
			index, err := BinarySearch(test.inputArray, test.inputValue)
			if test.found && err != nil {
				t.Errorf("got error, want index %v", test.index)
			}
			if !test.found && err == nil {
				t.Errorf("got index %v, want error", index)
			}
			if index != test.index {
				t.Errorf("got index %v, want index %v", index, test.index)
			}
		})
	}
}
