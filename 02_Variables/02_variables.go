package main

import (
	"fmt"
)

const LoginToken string = "hsdjbu" // public variable can  be Access by any method & function.

func main() {
	var username string = "Yunus"
	fmt.Printf("User Name: %T \n", username)
	fmt.Println("User Name:", username)

	var isloggedin bool = true
	fmt.Println(isloggedin)
	fmt.Printf("Type: %T \n", isloggedin)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Type: %T \n", smallValue)

	var num1 uint = 678
	var num2 uint = 7923
	var TwoSUm uint = num1 + num2
	fmt.Println("Sum:", TwoSUm)

	var smallFloat float32 = 255.68732476
	fmt.Println(smallFloat)
	fmt.Printf("%T \n", smallFloat)

	//default values & some aliases
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Type: %T \n", anotherVar)

	// implicit type
	var website = "yunus.eu.org"
	fmt.Println(website)
	website = "yunuscloud.eu.org"
	fmt.Println(website)

	// no var style

	numberOfUser := 76234.89
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
}
