package main

import (
	"fmt"
	"unicode"
)

type CalculateRune func(line string)

func compute(list []string, callback CalculateRune) {
	for _, v := range list {
		callback(v)
	}
}

func main() {
	list := []string{
		"hello world",
		"this is unbeliabd 34",
		"This is canndf **(74",
		"Last word of the day",
	}

	letter := 0
	num := 0
	punc := 0
	space := 0

	compute(list, func(line string) {
		for _, v := range line {
			if unicode.IsLetter(v) {
				letter += 1
			}
			if unicode.IsNumber(v) {
				num += 1
			}
			if unicode.IsPunct(v) {
				punc += 1
			}
			if unicode.IsSpace(v) {
				space += 1
			}
		}
	})

	fmt.Println("Letter :", letter)
	fmt.Println("Number :", num)
	fmt.Println("Punctuation :", punc)
	fmt.Println("Space :", space)
}
