package main

import "fmt"

func add(lhs, rhs int) int {
	return lhs + rhs
}

func compute(lhs, rhs int, op func(lhs, rhs int) int) int {
	return op(lhs, rhs)
}

func main() {
	fmt.Printf("Addition by using func-literals: %v\n", compute(2, 2, add))

	fmt.Printf("Subtraction Inline anonymous function: %v\n", compute(2, 2, func(lhs, rhs int) int {
		return lhs - rhs
	}))

	mul := func(lhs, rhs int) int {
		return lhs * rhs
	}

	fmt.Println("Multiplication : ", compute(2, 2, mul))
}
