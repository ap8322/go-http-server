package main

import (
	"fmt"
)


type foo struct {
	Bar int
	Baz string
}

type qux struct {
	foo
}

func (f foo) Courge() {
	fmt.Println(f.Baz)
}

func (f foo) Qux() {
	fmt.Println(f.Bar)
}

func main() {
	q := qux{}
	q.Baz = "world"
	q.Courge()
}
