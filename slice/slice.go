package main

import "fmt"

// Slice is a dynamic array in Go.
// slices are arrays uder the hood
// slice have two key properties: length and capacity
// length is the number of elements in the slice
// capacity is the numeber of elements it can store, capacity starts with equal to length and double everytime capacity is exhausted.

func main() {

	// an empty slice
	var slice = []int{}
	fmt.Print(slice)

	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	// append to add new elements in the slice
	slice2 = append(slice2, 4, 5)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2)) // capacity 6 -> doubled from 3.

	// Slice also has an other way of declaration that is using make function.
	// in which we can specify length and capacity.
	// and capacity still dynamically grow if needed.

	slice3 := make([]int, 3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // len 3
	fmt.Println(cap(slice3)) // cap 5

	slice4 := make([]int, 3)
	fmt.Println(cap(slice4)) // cap 3

}
