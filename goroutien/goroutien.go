package main

import (
	"fmt"
	"time"
)

func main() {

	data := []rune{'a', 'b', 'c', 'd'}

	capitalize := []rune{}

	cap := func(l rune) {
		capitalize = append(capitalize, l)
		fmt.Println("Done")

	}

	for _, v := range data {
		go cap(v)
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Capitalize: %c\n", capitalize)
}
