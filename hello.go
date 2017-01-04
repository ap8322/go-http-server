package main

import (
	"fmt"
)

func foo() {
	fmt.Println("foo")
}

func main() {
	fmt.Println()
}

func bar(msg string) {
	fmt.Println(msg)
}

func tuple() (int, string) {
	return 1,"a"
}
