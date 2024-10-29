package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	var fruitlist [4]string
	fruitlist[0] = "Mango"
	fruitlist[1] = "Apple"
	fruitlist[3] = "Peach"

	fmt.Println("Fruit list is ", fruitlist)
	fmt.Println("Fruit list is ", len(fruitlist))

	var vegList = [5]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegies list is: ", len(vegList))
	fmt.Println("Vegies List is: ", vegList)
}
