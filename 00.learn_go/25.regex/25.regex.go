package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("h([a-z]+)e", "house")
	fmt.Println(match)

	r, _ := regexp.Compile("h([a-z]+)e")
	fmt.Println(r.MatchString("hose"))
}
