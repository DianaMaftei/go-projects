package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var wg sync.WaitGroup

	var counter uint64

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			doTask(i)
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&counter, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

func doTask(id int) {
	fmt.Println("Starting task", id)
	time.Sleep(time.Second)
	fmt.Println("Finished task", id)
}
