package main

import (
	"fmt"
)

type Room struct {
	name    string
	cleaned bool
}

func checkRoom(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		if rooms[i].cleaned {
			fmt.Printf("%v is cleaned \n", rooms[i].name)
		} else {
			fmt.Printf("%v is not cleaned \n", rooms[i].name)
		}
	}
}

func main() {
	rooms := [...]Room{
		{name: "kitchen"},
		{name: "hall"},
		{name: "bedRoom"},
		{name: "bathroom"},
	}

	checkRoom(rooms)
	rooms[1].cleaned = true
	rooms[3].cleaned = true

	fmt.Println("Cleaning some more rooms...")
	checkRoom(rooms)
}
