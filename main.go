package main

import (
	"fmt"

	"github.com/bevane/go-dsa/sort"
)

func main() {
	original := []int{3, 4, 6, 2, 1}
	arr := make([]int, 5)
	copy(arr, original)
	fmt.Println(arr)

	sort.InsertionSort(arr)
	fmt.Println(arr)
}
