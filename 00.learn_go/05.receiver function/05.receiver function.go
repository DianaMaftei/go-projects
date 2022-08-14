package main

import "fmt"

type app []string

// usually using 1 letter for the receiver
func (a app) print() {
	for i, app := range a {
		fmt.Println(i, app)
	}
}

func main() {
	apps := app{"Facebook", "Instagram", "WhatsApp"}
	apps.print()
}
