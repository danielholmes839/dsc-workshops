package main

import "fmt"

type Greeter interface {
	Greet(person string) string
}

// Implementation #1
type BasicGreeter struct {
}

func (g *BasicGreeter) Greet(person string) string {
	return fmt.Sprintf("Hello %s", person)
}

// Implementation #2
type NamedGreeter struct {
	name string
}

func (g *NamedGreeter) Greet(person string) string {
	return fmt.Sprintf("Hello %s my name is %s", person, g.name)
}

// Function that uses a Greeter
func greet(g Greeter, person string) {
	fmt.Println(g.Greet(person))
}

func main() {
	basic := &BasicGreeter{}
	named := &NamedGreeter{name: "John Doe"}

	greet(basic, "Jane Doe")
	greet(named, "Jane Doe")
}
