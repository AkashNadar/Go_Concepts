package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	mtx := &sync.Mutex{}

	score := []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup, mtx *sync.Mutex) {
		fmt.Println("Inserting 1")
		mtx.Lock()
		score = append(score, 1) //Write operation is carried out we need to lock it while one thread using it
		mtx.Unlock()
		wg.Done()
	}(wg, mtx)

	go func(wg *sync.WaitGroup, mtx *sync.Mutex) {
		fmt.Println("Inserting 2")
		mtx.Lock()
		score = append(score, 2)
		mtx.Unlock()
		wg.Done()
	}(wg, mtx)

	go func(wg *sync.WaitGroup, mtx *sync.Mutex) {
		fmt.Println("Inserting 3")
		mtx.Lock()
		score = append(score, 3)
		mtx.Unlock()
		wg.Done()
	}(wg, mtx)

	wg.Wait()
	fmt.Println(score)
}
