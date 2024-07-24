package main

import "fmt"

// if-else
func main() {

	// x := 11

	// if x == 10 {
	// 	fmt.Println("X is equal to 10")
	// } else if x > 10 {
	// 	fmt.Println("X is greater than 10")
	// } else {
	// 	fmt.Println("X is less than 10")
	// }

	//switch
	day := 40

	switch day {
	case 1, 2, 3:
		fmt.Println("Monday")
	case 4:
		fmt.Println("Tuesday")
	default:
		fmt.Println("great day")
	}

	// expreessions
	switch {
	// here go dosen't remove brackets as here it is needed
	case (day > 10 && day > 20) || (day > 30):
		fmt.Println("Nice day")
	default:
		fmt.Println("Okay-okay")
	}

}
