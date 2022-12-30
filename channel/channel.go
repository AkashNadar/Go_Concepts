package main

import (
	"fmt"
	"sync"
)

// Channels are the way for which go routines talk to each other.
func main() {

	chn := make(chan int, 2)

	// Deadlock error
	// chn <- 4
	// fmt.Println(<-chn)

	wg := &sync.WaitGroup{}

	wg.Add(2)
	// Receiver Goroutine
	go func(chn <-chan int, wg *sync.WaitGroup) {
		// close(chn) will be invalid because in the parameter we have declared it as
		fmt.Println(<-chn)
		fmt.Println(<-chn)
		wg.Done()
	}(chn, wg)

	// Sender Goroutine (Read only)
	go func(chn chan<- int, wg *sync.WaitGroup) {
		chn <- 5
		chn <- 6
		close(chn)
		wg.Done()
	}(chn, wg)

	wg.Wait()
}
