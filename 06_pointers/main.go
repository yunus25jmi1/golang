package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the class on pointers.")

	var ptr *int
	fmt.Println("Value ", ptr)

	myNumber := 23
	var ptr1 = &myNumber
	fmt.Println("Value assigned: ", ptr1)
	fmt.Println("Value assigned: ", *ptr1)
	fmt.Println("Value assigned: ", ptr1)

	*ptr1 = *ptr1 * 2
	fmt.Println("New Value is: ", myNumber)

}
