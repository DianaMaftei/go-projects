package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("color", "purple")
	fmt.Println("color:", os.Getenv("color"))

	fmt.Println(os.Environ())
}
