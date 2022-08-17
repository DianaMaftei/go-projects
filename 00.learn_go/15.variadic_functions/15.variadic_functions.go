package main

import "fmt"

func main() {
	order("John", "Cola", "Steak", "Sweet Potatoes", "Apple Pie")
	order("Jack", "Becks")
	order("Tina", "Orange Juice", "Eggs Sunny Side Up", "Bread")

}

func order(patron string, products ...string) {
	fmt.Println(patron, "ordered:")
	for _, product := range products {
		fmt.Println(product)
	}
}
