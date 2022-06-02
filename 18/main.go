package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
)

type counter struct {
	count int
	mutex sync.Mutex
	wg    sync.WaitGroup
}

func main() {
	tries := flag.Int("tries", 5, "tries for increasing counter")
	if *tries < 1 {
		log.Fatal("invalid tries")
	}

	c := newCounter()
	for i := 0; i < *tries; i++ {
		c.wg.Add(1)
		go c.increase()
	}

	c.wg.Wait()

	fmt.Println(c.count)
}

func newCounter() *counter {
	return &counter{}
}

func (c *counter) increase() {
	c.mutex.Lock()
	c.count++
	c.mutex.Unlock()
	c.wg.Done()
}
