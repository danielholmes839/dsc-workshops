package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 10; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d\n", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'm'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c\n", i)
	}
}
func main() {
	/* from: 
	https://golangbot.com/goroutines/#:~:text=Goroutines%20are%20functions%20or%20methods,thousands%20of%20Goroutines%20running%20concurrently.
	*/
	go numbers()
	go alphabets()

	// waiting a minimum of 3 seconds
	time.Sleep(3 * time.Second)
	fmt.Println("done")
}
