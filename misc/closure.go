package main

import (
	"fmt"
)

func makeCounter() func() int {
	var count int
	return func() int {
		count++
		return count
	}
}

func main() {
	counter := makeCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	counter2 := makeCounter()
	fmt.Println(counter2())
	fmt.Println(counter2())
	fmt.Println(counter2())
}
