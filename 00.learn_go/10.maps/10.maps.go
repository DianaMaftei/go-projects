package main

import "fmt"

// keys and values are statically typed
// all keys and all values must be of the same type
func main() {

	var fruit map[string]string

	vegetables := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00FF00",
		"white": "#ffffff",
		"blue":  "#0000FF",
	}

	fmt.Println(fruit)
	fmt.Println(vegetables)
	fmt.Println(colors)

	vegetables["apple"] = "red"
	fmt.Println(vegetables)

	delete(colors, "green")
	fmt.Println(colors)

	printMap(colors)

}

func printMap(input map[string]string) {
	for key, value := range input {
		fmt.Println("The value for", key, "is", value)

	}
}
