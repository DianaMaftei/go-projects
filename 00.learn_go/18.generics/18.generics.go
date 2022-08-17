package main

import "fmt"

// constraints: any, comparable, different types separated by pipe
func main() {
	display(2)
	display("two")

	fmt.Println(multiplyByTwo(5))
	fmt.Println(multiplyByTwo(3.14))
	//fmt.Println(multiplyByTwo("cat"))
}

func display[T any](item T) {
	fmt.Println(item)
}

func multiplyByTwo[T int | float64](number T) T {
	return number * 2
}
