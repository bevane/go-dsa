package sort

import "fmt"

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i - 1
		for j >= 0 && arr[j+1] < arr[j] {
			tmp := arr[j+1]
			arr[j+1] = arr[j]
			arr[j] = tmp
			j--
		}

	}
	fmt.Println("Sorted with insertion sort")
}

func MergeSort(arr []int) {
	low := 0
	// high is the last index of the arr
	high := len(arr) - 1
	mergeSortR(arr, low, high)
}

func mergeSortR(arr []int, low int, high int) {
	if low >= high {
		return
	}

	mid := (low + high) / 2

	mergeSortR(arr, low, mid)
	mergeSortR(arr, mid+1, high)

	merge(arr, low, mid, high)
}

func merge(arr []int, low int, mid int, high int) {
	left := make([]int, mid-low+1)
	right := make([]int, high-mid)

	copy(left, arr[low:mid+1])
	copy(right, arr[mid+1:high+1])

	i := 0
	j := 0
	k := low

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
			k++
		} else {
			arr[k] = right[j]
			j++
			k++
		}
	}

	// if there are remaining elements in left or right
	// copy them over to arr

	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}
