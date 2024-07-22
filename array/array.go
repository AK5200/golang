package main

import "fmt"

//Arrays size need to be declared while declaring and cannot be changed.
// in arrays uninitialized elements have their default datatype value, e.g for int 0, string "", etc.
func main() {

	var names [3]string

	names[0] = "Anupam"
	names[2] = "Anurag"

	fmt.Println(names)
	// [Anupam,  Anurag, ]
	fmt.Printf("names are %q\n", names)
	// %q -> quested string ["Anupam", "", "Anurag"]

	// declaring at initilization time.

	var names2 = [3]string{"Anupam"}
	names2[1] = "Anurag"

}
