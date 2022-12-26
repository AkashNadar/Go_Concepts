package msg

import "fmt"

func Emote(emot string) {
	fmt.Println(emot, "From msg package emote.go")
}

func privateEmote(privEmote string) {
	fmt.Println(privEmote, "Private Emote ")
}
