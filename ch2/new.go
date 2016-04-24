package main

import (
	"fmt"
)

func newInt() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}

func main() {
	i := newInt()
	fmt.Println(i)
	fmt.Println(*i)
	*i = 10
	fmt.Println(*i)

	i2 := newInt2()
	fmt.Println(i2)
	fmt.Println(*i2)
	*i2 = 10
	fmt.Println(*i2)
}
