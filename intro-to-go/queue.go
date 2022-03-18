// queue
package main

import (
	"fmt"
)

type Queue struct {
	items []interface{}
}

func (s *Queue) Enqueue(t interface{}) {
	s.items = append(s.items, t)
}

func (s *Queue) Dequeue() interface{} {
	item := s.items[0]
	s.items = s.items[1:len(s.items)]
	return item
}

func (s *Queue) Front() *interface{} {
	item := s.items[0]
	return &item
}

func (s *Queue) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Queue) Size() int {
	return len(s.items)
}

// ---------------------------------------------------------
func main() {

	var queue Queue

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	// NOT TYPE SAFE!!!
	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("c")

	for !queue.IsEmpty() {
		t := queue.Dequeue()
		fmt.Println(t)
	}

	// x := t.(int) + 5
	// fmt.Println(x)
}
