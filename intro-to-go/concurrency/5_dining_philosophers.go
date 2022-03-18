/*
The classical Dining philosophers problem.

Implemented with forks (aka chopsticks) as mutexes.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Fork struct {
	sync.Mutex
	philosopher int // who's using the fork
}

type Table struct {
	*sync.Mutex
	forks []Fork
}

func NewTable(n int) *Table {
	return &Table{
		Mutex: &sync.Mutex{},
		forks: make([]Fork, n),
	}
}

func (t *Table) forkPositions(philosopher int) (int, int) {
	// philophers are in the range [1,n]
	right := philosopher % len(t.forks)
	left := philosopher - 1
	return left, right
}

func (t *Table) print() {
	for _, fork := range t.forks {
		fmt.Print(fork.philosopher, " ")
	}
	fmt.Println()
}

func (t *Table) StartDining(philosopher int) {
	left, right := t.forkPositions(philosopher)
	t.forks[left].Lock()
	t.forks[right].Lock()

	// updating the table
	t.Lock()
	defer t.Unlock()
	t.forks[left].philosopher = philosopher
	t.forks[right].philosopher = philosopher
	t.print()
}

func (t *Table) StopDining(philosopher int) {
	left, right := t.forkPositions(philosopher)
	t.forks[left].Unlock()
	t.forks[right].Unlock()

	// updating the table
	t.Lock()
	defer t.Unlock()
	t.forks[left].philosopher = 0
	t.forks[right].philosopher = 0
	t.print()
}

func philosopher(t *Table, p int) {
	for {
		t.StartDining(p)
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))

		t.StopDining(p)
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))
	}
}

func main() {
	n := 5
	table := NewTable(n)

	for p := 1; p <= n; p++ {
		go philosopher(table, p)
	}

	// wait indefinitely
	<-make(chan bool)
}
