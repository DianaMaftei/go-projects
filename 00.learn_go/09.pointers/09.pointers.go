package main

import "fmt"

type animal struct {
	name string
	age  int
}

// & - get the memory address
// * - get the value at this pointer

// value types: int, float, string, bool, structs
// reference types: slices, maps, channels, pointers, functions
func main() {
	cat := animal{
		name: "Dorian",
		age:  6,
	}

	fmt.Println(cat)

	catPointer := &cat
	catPointer.updateAge(7)

	fmt.Println(cat)
	fmt.Println(catPointer)

	dog := animal{
		name: "Spyke",
		age:  3,
	}

	fmt.Println(dog)
	dog.updateAge(4)
	fmt.Println(dog)

	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)
	fmt.Println(mySlice)

}

func updateSlice(s []string) {
	s[0] = "Hello"
}

func (pointerToAnimal *animal) updateAge(newAge int) {
	(*pointerToAnimal).age = newAge
}
