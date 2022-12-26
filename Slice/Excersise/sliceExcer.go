package main

import (
	"fmt"
)

type Part string

func printParts(line []Part) {
	for i := 0; i < len(line); i++ {
		parts := line[i]
		fmt.Println(parts)
	}
	fmt.Println("---------------------")
}

func main() {
	assLine := make([]Part, 3)
	assLine[0] = "pipe"
	assLine[1] = "rod"
	assLine[2] = "wire"
	printParts(assLine)
	assLine = append(assLine, "screw", "tools")
	printParts(assLine)
	newTools := assLine[3:]
	printParts(newTools)

}
