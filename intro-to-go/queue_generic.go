package main

// Go 1.18 released March 15th, 2022 has generics!!!
import (
	"fmt"
)

type Queue[T any] struct {
	items []T
}

func (s *Queue[T]) Enqueue(t T) {
	s.items = append(s.items, t)
}

func (s *Queue[T]) Dequeue() T {
	item := s.items[0]
	s.items = s.items[1:len(s.items)]
	return item
}

func (s *Queue[T]) Front() T {
	return s.items[0]
}

func (s *Queue[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Queue[T]) Size() int {
	return len(s.items)
}

// ---------------------------------------------------------
func main() {
	// queue with ints
	intQueue := &Queue[int]{
		items: make([]int, 0),
	}

	intQueue.Enqueue(1)
	intQueue.Enqueue(2)
	intQueue.Enqueue(3)

	for !intQueue.IsEmpty() {
		fmt.Println(intQueue.Dequeue())
	}

	// queue with strings
	stringQueue := &Queue[string]{
		items: make([]string, 0),
	}

	stringQueue.Enqueue("A")
	stringQueue.Enqueue("B")
	stringQueue.Enqueue("C")

	for !stringQueue.IsEmpty() {
		fmt.Println(stringQueue.Dequeue())
	}
}
