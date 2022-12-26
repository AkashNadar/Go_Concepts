package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (chicken Chicken) PrepareDish() {
	fmt.Println("cooking chicken")
}

func (salad Salad) PrepareDish() {
	fmt.Println("salad preparing")
}

func multipleDish(dishes []Preparer) {
	for i := 0; i < len(dishes); i++ {
		dishes[i].PrepareDish()
	}
}

func main() {
	dishes := []Preparer{Chicken("chicken"), Salad("salad")}
	multipleDish(dishes)

}
