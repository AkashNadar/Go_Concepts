package main

import (
	"fmt"
)

func main() {
	shoppingList := make(map[string]int)
	shoppingList["egg"] = 11
	shoppingList["cake"] = 2
	shoppingList["milk"] += 1
	shoppingList["egg"] += 1

	fmt.Println(shoppingList)

	delete(shoppingList, "cake")

	fmt.Println("new List :- ", shoppingList)

	qty, found := shoppingList["almonds"]

	if !found {
		fmt.Println("Almonds not found")
	} else {
		fmt.Println("%v Almonds found", qty)
	}

	totalQty := 0
	for _, val := range shoppingList {
		totalQty += val
	}
	fmt.Println("total no of items :-", totalQty)

}
