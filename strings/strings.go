package main

import (
	"fmt"
	"strings"
)

func main() {

	// there are lot of string methods in Go. let us explore some of those.

	str1 := "Hello, World!"
	str2 := strings.Split(str1, ",")
	fmt.Println(str2)

	// trim spaces from both ends of the string
	str1 = "  Hello, World!  "
	str3 := strings.TrimSpace(str1)
	fmt.Println(str3)

	// join, joins multiple strings with a specified separator
	str4 := "hello"
	str5 := "Sup"
	str6 := "guys!"
	str7 := strings.Join([]string{str4, str5, str6}, " ")
	fmt.Println(str7)
}
