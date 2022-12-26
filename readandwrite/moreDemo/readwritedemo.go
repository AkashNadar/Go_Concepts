package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func ReadStats(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Clean up code
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("file name: %s\n", stats.Name())
	fmt.Printf("Time Modified: %v\n", stats.ModTime().Format("15:04:03"))
}

func ReadWholeFile(filename string) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(contents))
}

func ReadLineByLine(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Clean up code
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func ReadByWord(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Clean up code
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func ReadByBytes(filename string, size uint8) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Clean up code
	defer file.Close()
	buf := make([]byte, size)

	for {
		totalRead, err := file.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err.Error())
			}
			break
		}
		fmt.Println(string(buf[:totalRead]))
	}
}

func ReadConfig(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Clean up code
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		raw := strings.Split(scanner.Text(), "=")
		fmt.Print("Key: ", raw[0], "  ")
		fmt.Println("value: ", raw[1])
	}

}

func main() {
	// filename := "file.dat"
	// ReadStats(filename)
	// ReadWholeFile(filename)
	// ReadLineByLine(filename)
	// ReadByWord(filename)
	// ReadByBytes(filename, 4)
	cfg_filename := "test.cfg"
	ReadConfig(cfg_filename)
}
