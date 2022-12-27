package main

import (
	"fmt"
	"time"
)

func greeting(msg string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(msg)
	}
}

func main() {
	go greeting("hello")
	greeting("world")
}
