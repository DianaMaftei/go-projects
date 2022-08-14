package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	write("Hello there!", "greeting.txt")
	fmt.Println("Greeting is: ", read("greeting.txt"))
}

func write(message, filename string) {
	messageBytes := []byte(message)

	ioutil.WriteFile(filename, messageBytes, 0666)
}

func read(filename string) string {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return string(bs)
}
