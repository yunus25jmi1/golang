package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent
	yunus := User{"Yunus", "yunus@go.dev", true, 20}
	fmt.Println(yunus)
	fmt.Printf("Yunus details are: %+v\n", yunus)
	fmt.Printf("Name is %v and email is %v.", yunus.Name, yunus.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
