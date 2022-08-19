package io

import (
	"fmt"
	"log"
	"os"
)

func ReadFile(fileName string) []byte {
	f, err := os.Open(fileName)
	handleError(err, fmt.Sprintf("Unable to open file %s", fileName))

	defer f.Close()

	data, err := os.ReadFile(fileName)
	handleError(err, fmt.Sprintf("Unable to read file %s", fileName))

	return data
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
		os.Exit(1)
	}
}
