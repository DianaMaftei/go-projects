package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const world = "世界"
	fmt.Println("Len:", len(world))

	fmt.Println("Rune count:", utf8.RuneCountInString(world))

	for _, r := range world {
		examineRune(r)
	}
}

func examineRune(r rune) {
	if r == '世' {
		fmt.Println("found generation")
	} else if r == '猫' {
		fmt.Println("found cat")
	}
}
