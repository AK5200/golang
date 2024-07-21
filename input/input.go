package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// fmt.Println(`Hello, What is your name?`)
	// var name string
	// fmt.Scan(&name)
	// // stops at space.
	// fmt.Printf(`Hello, %s!`, name)

	// for multi word input we need to use bufio-> buffered input/output.
	reader := bufio.NewReader(os.Stdin) // standin is input. i.e keyboard input.
	fmt.Print("What is your name?\n")   // Print can also be used.
	fullname, _ := reader.ReadString('\n')
	fmt.Printf(`Hello, %s`, fullname)

}
