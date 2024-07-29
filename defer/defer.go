package main

import "fmt"

func add(a, b int) {
	fmt.Println(a + b)
}

func main() {
	// defer keywords runs at the end of function execution
	// and if multiple defer, stored in stack so executes in reverse order

	// fmt.Println("hello 1")
	// defer fmt.Println("hello 2")
	// fmt.Println("hello 3")

	// hello1 hello3 hello2

	defer add(2, 3)
	fmt.Println("hello 4")
	fmt.Println("hello 5")

	// usually used for time consuming and not urgent function calls

}
