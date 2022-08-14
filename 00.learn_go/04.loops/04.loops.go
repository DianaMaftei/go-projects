package main

import "fmt"

func main() {
	animals := []string{"cat", "dog", "rabbit"}

	// itterate over a slice
	for i, animal := range animals {
		fmt.Println(i, animal)
	}

}
