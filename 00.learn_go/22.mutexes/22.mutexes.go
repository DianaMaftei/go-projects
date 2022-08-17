package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type SafeCounter struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *SafeCounter) inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.counters[key]++
}

func main() {
	c := SafeCounter{
		counters: map[string]int{"even": 0, "odd": 0},
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < 100; i++ {
		if r1.Int()%2 == 0 {
			go c.inc("even")
		} else {
			go c.inc("odd")
		}
	}

	time.Sleep(time.Second)
	fmt.Println(c.counters)
}
