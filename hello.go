package main

import (
	"fmt"
)

func main() {
	// fmt.Println("hello world")

	// bye2()
	// dir.Print()
	// dir.Printtwo()

	// variable declaration 1. can be written without type as well as type
	var v1 = "hello"
	fmt.Println(v1)

	var v2 string = "world"
	fmt.Println(v2)

	var v3 bool = true
	fmt.Println(v3)

	const v4 float64 = 3.14
	//v4 = 10 error
	fmt.Println(v4)

	var v5 int = 10
	fmt.Println(v5)

	// vaiable shorthand declaration MOST POPULAR
	v6 := "hello"
	fmt.Println(v6)

}

// to run bye.go --> go run hello.go bye.go

// func bye2() {

// 	fmt.Println(`bye bye world`)
// }
