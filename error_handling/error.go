package main

import "fmt"

// error is a type in go
// if error print that error else return nil.
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("error: division by zero")
	}
	return a / b, nil
}

func main() {

	// _ is used for a storing a function's return value, which we will not use or not need.
	// here _ is the error value from the function divide()
	result, _ := divide(10, 0)
	fmt.Println("Result:", result)

	// we can alse use a name instead of _ but then we will have to use it.
	result, err := divide(10, 2)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

}
