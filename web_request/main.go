package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://akash.dev"

func main() {
	fmt.Println("Web Request")
	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of the response %T", resp)
	defer resp.Body.Close() //caller responsibility to close the connection

	dataByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	content := string(dataByte)
	fmt.Println("data :", content)
}
