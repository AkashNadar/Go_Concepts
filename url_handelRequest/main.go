package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func PerformGetRequest(myUrl string) {

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content-Length :", response.ContentLength)

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(content))
	var responseContent strings.Builder
	byteCount, err := responseContent.Write(content)
	fmt.Println("ByteCount:- ", byteCount)
	fmt.Println("Content :", responseContent.String())
}

func PerformPostRequest(myUrl string) {

	requestBody := strings.NewReader(`
		{
			"name": "GoLang",
			"price": 2500,
			"site": "udemy"
		}
	`)
	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	dataByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var responseContent strings.Builder
	byteLen, err := responseContent.Write(dataByte)
	if err != nil {
		panic(err)
	}
	fmt.Println("ByteLength :", byteLen)
	fmt.Println("Content :", responseContent.String())
}

func PerformPostFormRequest(myUrl string) {
	// Form Data
	data := url.Values{}
	data.Add("firstname", "Akash")
	data.Add("lastname", "Nadar")
	data.Add("email", "Akash@gmail.com")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	dataByte, err := ioutil.ReadAll(response.Body)

	fmt.Println("Data :", string(dataByte))
}

func main() {
	fmt.Println("Welcome")
	// PerformGetRequest("http://localhost:8000/get")
	// PerformPostRequest("http://localhost:8000/post")
	PerformPostFormRequest("http://localhost:8000/postform")

}
