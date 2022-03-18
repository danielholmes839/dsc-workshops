package main

import (
	"time"
)

func deadlock1() {
	c := make(chan int)
	<-c
}

func deadlock2() {
	c := make(chan int)
	c <- 1
}

func alive() {
	for {
		time.Sleep(time.Second)
	}
}

func main() {
	// go alive()
	deadlock2()
}
