package main

import (
	"fmt"

	searching "github.com/algorithm/searching"
)

func main() {
	arr := []int{10, 12, 13, 16, 18, 19, 20, 21,
		22, 23, 24, 33, 35, 42, 47}

	// fmt.Println("Linear Search Result :", searching.LinearSearch(arr, 12))
	// fmt.Println("Binary Search Result :", searching.BinarySearch(arr, 11))
	fmt.Println("Interpolation Search Result :", searching.InterpolationSearch(arr, 0, len(arr)-1, 19))
}
