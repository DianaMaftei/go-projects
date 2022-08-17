package main

import (
	"encoding/json"
	"fmt"
)

type book struct {
	Title   string
	Authors []string
	Year    int
	price   float64 // private fields are not marshalled
}

func main() {
	bool1, _ := json.Marshal(true)
	fmt.Println(string(bool1))

	int1, _ := json.Marshal(1)
	fmt.Println(string(int1))

	str1, _ := json.Marshal("gopher")
	fmt.Println(string(str1))

	book1 := book{Title: "Good Omens", Authors: []string{"Terry Pratchett", "Neil Gaiman"}, Year: 1990, price: 10.99}
	book1M, _ := json.Marshal(book1)
	fmt.Println(book1)
	fmt.Println(string(book1M))

	book1U := book{}
	json.Unmarshal(book1M, &book1U)
	fmt.Println(book1U)

}
