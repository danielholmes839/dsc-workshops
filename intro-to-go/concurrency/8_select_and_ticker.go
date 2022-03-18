package main

import (
	"fmt"
	"time"
)

func background(done <-chan bool) {
	ticker := time.NewTicker(time.Millisecond * 500)
	i := 0

	for {
		select {
		case <-done:
			ticker.Stop()
		case <-ticker.C:
			fmt.Println("tick", i)
			i++
		}
	}
}

func main() {
	done := make(chan bool)
	go background(done)

	time.Sleep(time.Second * 5)
	done <- true
	fmt.Println("done")
}
