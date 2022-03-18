package main

import "fmt"

func main() {
	c := make(chan int, 1)

	c <- 1
	fmt.Println("sent 1")

	c <- 2
	fmt.Println("sent 2")
}
