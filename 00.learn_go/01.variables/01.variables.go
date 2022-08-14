package main

import "fmt"

// color := "green" -- not valid outside of func
var color string = "blue"

var amount float64

func main() {

	fmt.Printf("Color: %s \n", color)

	amount = 25.67
	fmt.Printf("Amount: %f \n", amount)

	var animal string = "cat"
	fmt.Printf("Animal: %s \n", animal)

	var fruit = "peach"
	fmt.Printf("Fruit: %s \n", fruit)

	vegetable := "tomato"
	fmt.Printf("Vegetable: %s \n", vegetable)

	goIsCool := true
	fmt.Printf("Go is cool: %t \n", goIsCool)

	favoriteNumber := 42
	fmt.Printf("Favorite number: %v \n", favoriteNumber)

	var age int
	age = 99
	fmt.Printf("Age: %v \n", age)

}
