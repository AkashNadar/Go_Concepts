package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func writeFile(content string, fileName string) {
	file, err := os.Create(fileName) //For creation we need os operation
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("Length :", length)
	defer file.Close()
}

func readFile(filePath string) {
	dataByte, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data :", string(dataByte))
}

func main() {

	readFile("./demo.txt")

}
