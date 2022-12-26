package main

import "fmt"

func sum[T int | float32](arr []T) T {
	var total T
	for _, v := range arr {
		total += v
	}
	return total
}

type NumType interface {
	int | float32
}

func mul[T NumType](arr []T) T {
	var total T = 1
	for _, v := range arr {
		total *= v
	}
	return total
}

func main() {

	arr1 := []int{1, 2, 3, 4, 5, 6}
	arr2 := []float32{1.2, 3.3, 1.5, 0.5}

	fmt.Println("Sum of integers:- ", sum(arr1))
	fmt.Println("Mul of floats:- ", mul(arr2))
	fmt.Println("Sum of floats:- ", sum(arr2))
}
