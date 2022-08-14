package main

import "fmt"

func main() {
	animal := newAnimal()

	fmt.Println(animal)

	title, pages := getBookInfo()

	fmt.Println(title)
	fmt.Println(pages)

}

func newAnimal() string {
	return "Dog"
}

func getBookInfo() (string, int) {
	return "War and Peace", 1000
}
