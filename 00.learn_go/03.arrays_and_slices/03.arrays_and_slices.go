package main

import (
	"fmt"
	"strings"
)

func main() {
	// array - fixed length list of things
	// slice - an array that can grow or shrink
	fruit := [2]string{"apple", "pear"}
	vegetables := []string{"tomato", "potato", "carrot"}

	// fruit = append(fruit, "banana") -- does not work with arrays, that are fixed
	vegetables = append(vegetables, "cabbage") // doens't modify the original slice, creates a new one

	fmt.Println(fruit)
	fmt.Println(vegetables)

	fmt.Println(vegetables[:1])

	fmt.Println(strings.Join([]string(vegetables), "|"))

}
