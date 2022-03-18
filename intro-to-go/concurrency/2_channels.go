package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	go func() {
		// send 5 messages through the channel
		for i := 1; i <= 5; i++ {
			messages <- fmt.Sprintf("Message %d", i)
			time.Sleep(time.Second / 2)
		}
		close(messages)
	}()
	
	// read from a channel
	x := <-messages
	fmt.Println(x)

	// read from a channel (without assigning it to a variable)
	fmt.Println(<-messages)
	
	// read from a channel by iterating
	for message := range messages {
		fmt.Println(message)
	}

	// for loop stops because the channel is closed
	fmt.Println("done")
}
