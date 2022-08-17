package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(2 * time.Second)
	ticker := time.NewTicker(500 * time.Millisecond)

	go func() {
		for {
			select {
			case <-timer.C:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	<-timer.C
}
