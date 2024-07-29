package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42
	fmt.Printf("type of number is :%T\n", num)

	//data conversion
	var data float64 = float64(num)
	fmt.Printf("type of data is :%T\n", data)

	//string conversions
	var data1 string = "hell"
	data2, _ := strconv.Atoi(data1)           //Atoi --> Alpha to i (integer)
	fmt.Printf("data2 value is :%d\n", data2) // prints 0 as not a valid integer
	fmt.Printf("type of data2 is :%T\n", data2)

	// integer to string conversion
	data3 := 123
	data4 := strconv.Itoa(data3)       // Itoa --> Integer to a (alphabet)
	fmt.Println("data 4 is : ", data4) // prints "123"
	fmt.Printf("Type of data4 : %T\n", data4)

	// string to float conversion
	data5 := "3.14"
	data6, _ := strconv.ParseFloat(data5, 64) // data and then 64 (float64 or float32)
	fmt.Printf("data6 value is %f\n and type is: %T\n", data6, data6)

	// there are a lot of other conversions as well, as go is developed by Google and google extensively uses strings.
}
