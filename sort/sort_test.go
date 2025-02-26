package sort

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	input []int
	want  []int
}{
	{
		[]int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	{
		[]int{1, 5, 8, 7, 3, 6, 7, 4, 9},
		[]int{1, 3, 4, 5, 6, 7, 7, 8, 9},
	},
}

func TestInsertionSort(t *testing.T) {
	for _, test := range testCases {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			arr := make([]int, len(test.input))
			copy(arr, test.input)
			InsertionSort(arr)
			if !reflect.DeepEqual(arr, test.want) {
				t.Errorf("got %v want %v", arr, test.want)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	for _, test := range testCases {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			arr := make([]int, len(test.input))
			copy(arr, test.input)
			MergeSort(arr)
			if !reflect.DeepEqual(arr, test.want) {
				t.Errorf("got %v want %v", arr, test.want)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	for _, test := range testCases {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			arr := make([]int, len(test.input))
			copy(arr, test.input)
			QuickSort(arr)
			if !reflect.DeepEqual(arr, test.want) {
				t.Errorf("got %v want %v", arr, test.want)
			}
		})
	}
}

func TestBucketSort(t *testing.T) {
	for _, test := range testCases {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			arr := make([]int, len(test.input))
			copy(arr, test.input)
			BucketSort(arr, 1, 9)
			if !reflect.DeepEqual(arr, test.want) {
				t.Errorf("got %v want %v", arr, test.want)
			}
		})
	}
}
