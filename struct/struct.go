package main

import "fmt"

type Person struct {
	firstname string
	lastname  string
	age       int
}

type Designation struct {
	position string
	person   Person
}

func main() {

	// ERROR
	// var person1 Person = {
	// 	firstname: "Anupam",
	// 	lastname:  "Kumar",
	// 	age:       22,
	// }

	var person1 Person
	person1.firstname = "Anupam"
	person1.lastname = "Kumar"
	person1.age = 22

	fmt.Println(person1)

	// 2nd way
	person2 := Person{
		firstname: "Anurag",
		lastname:  "Kumar",
		age:       26,
	}
	fmt.Println(person2)

	// 3rd way using new keyword -> returns a pointer to the datatype
	person3 := new(Person)
	person3.firstname = "Gita"

	fmt.Println(person3)
	// &{Gita 0}
	fmt.Println(person3.firstname)

	example := Designation{
		position: "CEO",
		person: Person{
			firstname: "Anil",
			age:       50,
		},
	}

	var example2 Designation
	example2.position = "CA"
	example2.person.age = 45

	fmt.Println(example)
	fmt.Println(example2)

}
