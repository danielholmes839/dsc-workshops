package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	sync.Mutex
	value int
}

func (c *Counter) Increment(name string) int {
	c.Lock()
	fmt.Printf("%s acquired lock\n", name)

	c.value++

	c.Unlock()
	fmt.Printf("%s released lock\n", name)

	return c.value
}

func (c *Counter) Increment2(name string) int {
	c.Lock()
	defer c.Unlock()

	c.value++
	return c.value
}

func main() {
	counter := &Counter{value: 0, Mutex: sync.Mutex{}}

	go counter.Increment("A")
	go counter.Increment("B")

	time.Sleep(time.Second)
}
