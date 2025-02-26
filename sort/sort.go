package sort

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

func QuickSort(arr []int) {
	quickSortR(arr, 0, len(arr)-1)
}

func quickSortR(arr []int, low int, high int) {
	if low >= high {
		return
	}

	pivot := arr[high]
	// index denoting where left side ends (exclusive)
	leftEnd := low

	// iterate through all elements except for the pivot which is at high
	for i := low; i < high; i++ {
		// will move all the values less than pivot to the left
		if arr[i] < pivot {
			tmp := arr[i]
			arr[i] = arr[leftEnd]
			arr[leftEnd] = tmp
			// if a swap was made then there is one more value
			// added to left side, so increment leftEnd by one
			leftEnd++
		}
	}

	// move pivot in between right and left side
	arr[high] = arr[leftEnd]
	arr[leftEnd] = pivot

	// sort left and right sides
	quickSortR(arr, low, leftEnd-1)
	quickSortR(arr, leftEnd+1, high)
}

func BucketSort(arr []int, min int, max int) {
	count := make([]int, max-min+1)
	// add all values into their buckets

	for _, v := range arr {
		// for cases where the min is not equal to zero
		// hence the first index in count should refer to min
		// instead of zero
		count[v-min]++
	}

	i := 0
	// iterate over each index in count
	for j := range count {
		// iterate equal to value at index of count
		for range count[j] {
			// each index in count counts up from min instead of zero
			// so add back min to the index
			arr[i] = j + min
			i++
		}
	}
}
