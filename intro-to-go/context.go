package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

func doSomething(ctx context.Context, message string) {
	run := true
	ticker := time.NewTicker(time.Second)

	for run {
		select {
		case <-ticker.C:
			fmt.Println(message)
		case <-ctx.Done():
			fmt.Printf("doSomething %s cancelled\n", message)
			ticker.Stop()
			run = false
		}
	}

	fmt.Printf("doSomething %s return\n", message)
}

func main() {
	// https://pkg.go.dev/context#Context
	// https://go.dev/blog/pipelines

	// ctx := context.Background()
	// ctx, cancel := context.WithCancel(ctx)
	// ctx, _ = context.WithTimeout(ctx, time.Second*5)
	// ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	go doSomething(ctx, "A")
	go doSomething(ctx, "B")

	// hit enter to cancel early
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	cancel()

	time.Sleep(time.Second)
}
