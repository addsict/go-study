package main

import "fmt"

const (
	Foo int = iota
	Bar
	Baz
)

const (
	One int = 1 << (iota + 1)
	Two
	Three
)

func main() {
	fmt.Println(Foo, Bar, Baz)
	fmt.Println(One, Two, Three)
}
