package main

import (
	"fmt"
	"time"
)

// used to wait for multiple channels concurrently
// blocks until one of its cases can run, then it executes that case
// it chooses one at random if multiple are ready.
func main() {
	ch1 := make(chan string)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 2
	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch3 <- 0
	}()

	for i := 0; i < 4; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("received", msg1)
		case msg2 := <-ch2:
			fmt.Println("received", msg2)
		case <-ch3:
			fmt.Println("received on ch3")
		case <-time.After(3 * time.Second):
			fmt.Println("timeout 1")
			//default:
			//	fmt.Println("no message received")
		}
	}
}
