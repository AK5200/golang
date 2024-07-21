package main

import (
	"fmt"
)

func main() {
	age := 25
	name := "John Doe"
	height := 1.75

	fmt.Println("Age:", age, "Name:", name, "Height:", height)
	// automatically gives spaces between elements.
	// Age: 25 Name: John Doe Height: 1.75 and at end move control to new line.

	// printf -> uses format specifier,  more control , not auto spaceing and newline, similar to C.
	fmt.Printf("Age:%d Name:%s Height:%.2f\n", age, name, height)
	fmt.Printf("Type pf height: %T\n", height)
}
