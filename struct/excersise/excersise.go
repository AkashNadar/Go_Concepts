package main

import (
	"fmt"
)

type Coordinate struct {
	x int
	y int
}

type Rectangle struct {
	a Coordinate //top left
	b Coordinate //bottom right
}

func width(data Rectangle) int {
	return (data.b.x - data.a.x)
}

func height(data Rectangle) int {
	return (data.a.y - data.b.y)
}

func area(len int, bre int) int {
	return len * bre
}

func peremeter(len int, bre int) int {
	return 2 * (len + bre)
}

func main() {
	recCord := Rectangle{
		a: Coordinate{0, 7},
		b: Coordinate{10, 0},
	}

	fmt.Printf("the area of rectangle is %v", area(height(recCord), width(recCord)))
	fmt.Printf("the peremeter of rectangle is %v", peremeter(height(recCord), width(recCord)))
}
