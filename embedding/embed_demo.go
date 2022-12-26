package main

import "fmt"

const (
	Small = iota
	Medium
	Large
)

const (
	Ground = iota
	Air
)

type BeltSize int
type Shipping int

func (b BeltSize) String() string {
	return []string{"Small", "Medium", "Large"}[b]
}

func (s Shipping) String() string {
	return []string{"Ground", "Air"}[s]
}

type Conveyor interface {
	Convey() BeltSize
}

type Shipper interface {
	Ship() Shipping
}

type WarehouseAutomator interface {
	Conveyor
	Shipper
}

func automate(item WarehouseAutomator) {
	fmt.Println("Shipping :", item.Ship())
	fmt.Println("Conveying :", item.Convey())
}

type SpamMail struct {
	amount int
}

func (s SpamMail) String() string {
	return "spam mail"
}

func (s *SpamMail) Ship() Shipping {
	return Air
}

func (s *SpamMail) Convey() BeltSize {
	return Small
}

func main() {
	mail := SpamMail{
		amount: 1000,
	}
	automate(&mail)
}
