package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Finished A")
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("Finished B")
		wg.Done()
	}()

	fmt.Println("Waiting...")
	wg.Wait()
	fmt.Println("Finished A and B")

}
