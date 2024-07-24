package main

import "fmt"

func main() {

	// map is created using make keyword where 1st key datatype is in [] and value datatype directly without any brackets
	// go uses unordered_map -> it does not follows any order
	// using make we cannot initialize map

	example := make(map[string]int)

	example["Arun"] = 10
	example["Anupam"] = 20

	fmt.Println(example["Arun"])

	// here in map instead of index we get key and element value
	for index, element := range example {
		fmt.Printf("key is: %s and value is: %d\n", index, element)
	}

	delete(example, "Arun")
	fmt.Println(example["Arun"]) // 0
	// in go if the key val pair is not present in the map instead of throwing error or giving value not present, we get valueu = 0 for that key
	// this can be misleading as some key can have actual value as 0

	// for solving this we have a exist feature in map
	marks, exists := example["Arun"]
	fmt.Println("Marks of Arun: ", marks)
	fmt.Println("exists: ", exists)
	// gives false

	// we can also create map using literal and using loteral we can also initialize at the time of declaration.
	newExample := map[string]int{}

	newExample["Arun"] = 10

	// or
	newExampleTwo := map[int]int{
		1: 1,
		2: 4,
		3: 9,
	}

	for key, value := range newExampleTwo {
		fmt.Printf("Key is %d and value is %d\n", key, value)
	}

	// or

	fmt.Println(newExampleTwo)

}
