package main

import (
	"fmt"
	"net/url"
)

const Url = "https://akash.dev:3000/learn?coursename=reactjs&paymentid=Ujfdk323"

func main() {
	fmt.Println("Handling Urls in Golang")
	fmt.Println(Url)

	// Parsing
	result, err := url.Parse(Url)
	if err != nil {
		panic(err)
	}
	// fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("type of query params are %T\n", qparams)

	fmt.Println(qparams["coursename"])
	for key, val := range qparams {
		fmt.Println("key: ", key, "& val: ", val)
	}

	// Essential thing pass the reference not the copy
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "akash.dev",
		Path:    "/tuttorial",
		RawPath: "user=rahul",
	}

	fmt.Println("Another url :", partsOfUrl.String())
}
