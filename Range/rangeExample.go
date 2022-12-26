package main

import (
	"fmt"
)

func main() {
	names := []string{"Akash", "Rahul", "Ganesh", "Mohan"}

	for _, name := range names {
		for _, ch := range name {
			fmt.Printf("%q  ", ch)
		}
		fmt.Println()
		fmt.Println("-------------")
	}
}
