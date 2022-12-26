package main

import (
	"fmt"

	"example.com/Practise/pkg/display"

	"example.com/Practise/pkg/msg"
)

func main() {
	fmt.Println("is it working")
	display.Display("Welcome")
	msg.Msg("mesg")
	msg.Emote("Exited")
}
