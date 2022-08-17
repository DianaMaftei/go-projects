package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	handleErr(err)
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	// basic
	// body := make([]byte, 99999)
	// resp.Body.Read(body)
	// fmt.Println(string(body))

	// better
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func handleErr(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
