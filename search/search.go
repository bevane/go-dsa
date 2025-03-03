package search

import "fmt"

func BinarySearch(arr []int, value int) (int, error) {
	low := 0
	high := len(arr) - 1

	// LEQ instead of LE is important to find the value when value
	// is at a high or low index
	for low <= high {
		mid := (low + high) / 2
		if value > arr[mid] {
			low = mid + 1
		}
		if value < arr[mid] {
			high = mid - 1
		}
		if value == arr[mid] {
			return mid, nil
		}
	}
	return 0, fmt.Errorf("%v was not found", value)
}
