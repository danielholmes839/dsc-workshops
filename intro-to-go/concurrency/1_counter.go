package main

import (
	"fmt"
	"time"
)

func counter(name string, offset, interval time.Duration) {
	// initial sleep
	time.Sleep(offset)

	// counter
	for i := 0; i < 100; i++ {
		time.Sleep(interval)
		fmt.Printf("%s: %d\n", name, i)
	}
}

func main() {
	go counter("A", 0, time.Second)
	counter("B", time.Second/2, time.Second)
	
	// go counter("A", 0, time.Second)
	// go counter("B", time.Second/3, time.Second)
	// counter("C", time.Second*2/3, time.Second)
}
