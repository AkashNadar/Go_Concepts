package main

import (
	"fmt"
)

type Product struct {
	name  string
	price int
}

func printList(list [4]Product) {
	var noItems int = 0
	totalCost := 0
	for i := 0; i < len(list); i++ {
		item := list[i]

		if item.name != "" {
			totalCost += item.price
			noItems++
		}
	}
	fmt.Println(list[noItems-1].name, "is the last item")
	fmt.Println(noItems, "Toal no of items")
	fmt.Println(totalCost, "Total cost")
}

func main() {
	fmt.Println("Ya its working.. ")
	list := [4]Product{
		{name: "carrot", price: 20},
		{name: "beet", price: 80},
		{name: "cheese", price: 100},
	}

	printList(list)

	list[3] = Product{name: "fish", price: 200}
	printList(list)
}
