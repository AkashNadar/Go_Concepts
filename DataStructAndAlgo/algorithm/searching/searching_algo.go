package searching_algo

import (
	"sort"
)

func LinearSearch(arr []int, no int) bool { //n
	for _, val := range arr {
		if val == no {
			return true
		}
	}

	return false
}

func BinarySearch(arr []int, no int) bool { //log n
	low := 0
	high := len(arr)
	sort.Ints(arr)
	for low <= high {
		mid := low + (high-1)/2
		if arr[mid] == no {
			return true
		}

		if no > arr[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}

func InterpolationSearch(arr []int, low int, high int, no int) int { // O(log2(log2 n))
	var pos int

	if low <= high && no >= arr[low] && no <= arr[high] {

		pos = low + ((high-low)/(arr[high]-arr[low]))*(no-arr[low])

		if arr[pos] == no {
			return pos
		}

		if arr[pos] < no {
			return InterpolationSearch(arr, pos+1, high, no)
		}

		if arr[pos] > no {
			return InterpolationSearch(arr, pos-1, high, no)
		}
	}

	return -1
}
