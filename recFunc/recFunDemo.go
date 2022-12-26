package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

// Pointer Function
func occupiedSpace(parkingLot *ParkingLot, spaceNum int) {
	parkingLot.spaces[spaceNum].occupied = false
}

// Receiver Function
func (parkingLot *ParkingLot) occupied(spaceNum int) {
	parkingLot.spaces[spaceNum].occupied = false
}

func (parkingLot *ParkingLot) available(spaceNum int) {
	parkingLot.spaces[spaceNum].occupied = true
}

func main() {

	lot := ParkingLot{
		spaces: make([]Space, 10),
	}
	fmt.Println(lot)
	lot.available(0)
	lot.available(9)
	fmt.Println(lot)

}
