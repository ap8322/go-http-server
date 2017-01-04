package main

import (
	"fmt"
)

func (f foo) Qux() {
	fmt.Println("foo")
}

func (b bar) Qux() {
	fmt.Println("bar")
}

func main() {
	fmt.Printf("Hello, world\n")
}

func Courge(b baz) {
	b.Qux()
}
