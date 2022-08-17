package main

import "fmt"

func main() {
	animals := []string{"cat", "dog", "rabbit"}

	// itterate over a slice
	for i, animal := range animals {
		fmt.Println(i, animal)
	}

	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

}
