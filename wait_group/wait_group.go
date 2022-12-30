package main

import (
	"fmt"
	"net/http"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	endpoints := []string{
		"https://github.com",
		"https://google.com",
		"https://yahoo.com",
		"https://fb.com",
	}

	for _, val := range endpoints {
		go checkStatus(val)
		waitGroup.Add(1)
	}
	waitGroup.Wait()
}

func checkStatus(endpoint string) {

	res, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%v status: %d\n", endpoint, res.StatusCode)
	}

	defer waitGroup.Done()
}
