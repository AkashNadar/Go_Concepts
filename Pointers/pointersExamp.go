package main

import "fmt"

type Counter struct {
	hit int
}

func incrementCounter(counter *Counter) {
	counter.hit += 1
	fmt.Println("counter hit: ", counter)
	fmt.Println("address of counter ", &counter)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	incrementCounter(counter)
}

func main() {
	counter := Counter{}

	first := "hello"
	second := "world"
	fmt.Println(first, second)

	replace(&first, "hii", &counter)
	fmt.Println(first, second)

	slice := []string{first, second}
	fmt.Println(slice)

	replace(&slice[0], "Go", &counter)
	fmt.Println(slice)
}
