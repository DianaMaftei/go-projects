package main

import "fmt"

func main() {
	nextInt := intSequence()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intSequence()
	fmt.Println(newInt())
	fmt.Println(newInt())
}

func intSequence() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}