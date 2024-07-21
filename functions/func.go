package main

import (
	"fmt"
)

// Function (a int) params shold must have datatypes and return type.

// return type
func add(a int, b int) int {
	return a + b
}

// if all params has same type then we can declare just one after all the parameters.
func sub(a, b int) int {
	return a - b
}

// here we are using mul var in return type. which is it returns into mul variable and hence we can directly return, and automatically returns the mul variable
func multiply(a, b int) (mul int) {
	mul = a * b
	return
}

func main() {
	fmt.Println("Sum:", add(2, 3))
	fmt.Println("Subtraction:", sub(5, 3))
	fmt.Println("Multiplication:", multiply(2, 3))
}
