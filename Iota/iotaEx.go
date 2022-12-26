package main

import "fmt"

const (
	Add = iota
	Sub
	Mul
	Div
)

type Operation int

func (op Operation) calculate(num1, num2 int) int {
	switch op {
	case Add:
		return num1 + num2
	case Sub:
		return num1 - num2
	case Mul:
		return num1 * num2
	case Div:
		return num1 / num2
	}
	panic("unhandled operation")
}

func main() {
	operate := Operation(Add)
	fmt.Println("Addition :-", operate.calculate(10, 20))

	fmt.Println("Subraction :-", Operation(Sub).calculate(40, 20))
	fmt.Println("Multiplication :-", Operation(Mul).calculate(40, 20))
	fmt.Println("Divisioon :-", Operation(Div).calculate(40, 20))
}
