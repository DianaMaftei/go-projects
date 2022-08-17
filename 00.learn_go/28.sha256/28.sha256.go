package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	h := sha256.New() // get new hash

	s := "super secret code"

	h.Write([]byte(s))

	bs := h.Sum(nil) // finalized hash result as a byte slice

	fmt.Println(s)
	fmt.Printf("%x\n", bs)

}
