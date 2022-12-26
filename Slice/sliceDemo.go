package main

import (
	"fmt"
)

func printSlice(routeName string, slice []string) {
	fmt.Println("---", routeName, "---")
	for i := 0; i < len(slice); i++ {
		place := slice[i]
		fmt.Println(place)
	}
}

func main() {
	slice := []string{"shop", "hospital", "market"}
	printSlice("route1", slice)

	fmt.Println(slice[0], "visited")
	fmt.Println(slice[1], "visited")
	slice = append(slice, "station", "office")
	newSlice := slice[2:]

	printSlice("route2", newSlice)
}
