package main

import (
	"fmt"
)

type Passenger struct {
	name     string
	ticketNo int
	boarded  bool
}

func main() {

	pass1 := Passenger{
		name:     "akash",
		ticketNo: 34,
	}

	pass2 := Passenger{
		name:     "prashant",
		ticketNo: 11,
		boarded:  true,
	}
	fmt.Println(pass2)

	// Anonymous function
	var pass3 Passenger
	pass3.name = "Rahul"
	pass3.ticketNo = 67
	pass3.boarded = true
	fmt.Println(pass3)

	if pass1.boarded {
		fmt.Println("Akash has boarded")
	}

	if pass2.boarded {
		fmt.Println("Prashant has boarded")
	}

}
