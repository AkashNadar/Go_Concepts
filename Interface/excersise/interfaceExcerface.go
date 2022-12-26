package main

import "fmt"

const (
	SmallLifts = iota
	MediumLifts
	XlLifts
)

type SelectLift int

type LiftPicker interface {
	Pick() SelectLift
}

type SmallCars string
type LargeCars string
type XLCars string

func (sc SmallCars) String() string {
	return fmt.Sprintf("small car : %v", string(sc))
}

func (sc SmallCars) Pick() SelectLift {
	return SmallLifts
}

func (lg LargeCars) String() string {
	return fmt.Sprintf("large car : %v", string(lg))
}

func (lg LargeCars) Pick() SelectLift {
	return MediumLifts
}

func (xl XLCars) String() string {
	return fmt.Sprintf("small car : %v", string(xl))
}

func (xl XLCars) Pick() SelectLift {
	return XlLifts
}

func sendToLift(lp LiftPicker) {
	switch lp.Pick() {
	case SmallLifts:
		fmt.Println("required small lift")
	case MediumLifts:
		fmt.Println("required medium lift")
	case XlLifts:
		fmt.Println("required xl size lift")
	}
}

func main() {
	nano := SmallCars("nano")
	suv := LargeCars("suv")
	truck := XLCars("BestBus")

	sendToLift(nano)
	sendToLift(suv)
	sendToLift(truck)

}
