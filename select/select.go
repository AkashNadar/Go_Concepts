package main

import (
	"fmt"
)

func main() {

	myChannel := make(chan string)
	anotherChannel := make(chan string)

	// child
	go func() {
		myChannel <- "data of music"
	}()

	// child
	go func() {
		anotherChannel <- "data of programming"
	}()

	// joining child
	select {
	case music := <-myChannel:
		fmt.Println(music)
	case programming := <-anotherChannel:
		fmt.Println(programming)
	}
}
