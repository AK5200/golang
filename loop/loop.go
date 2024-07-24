package main

import "fmt"

func main() {

	// go just has for loop :) but have variations of it

	number1 := []int{1, 2, 3, 4, 5}

	// 1st the basic cpp type for loop
	for i := 0; i < len(number1); i++ {
		fmt.Println(number1[i])
	}

	// 2nd python js type inbuilt params index and element, can write anything instead of element and index keyword
	for index, element := range number1 {
		fmt.Printf("index is: %d, and element is %d\n", index, element)
	}

	// similar to while loop
	// count := 0
	// for {

	// 	fmt.Println("loop")
	// 	count++
	// }

	count := 0
	for {

		if count < 3 {
			fmt.Println("hello")
		}
		count++
	}

}
