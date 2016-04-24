package main

import (
	"fmt"
)

func main() {
	x, y := 10, 20
	fmt.Printf("x = %d, y = %d\n", x, y)
	x, y = y, x
	fmt.Printf("x = %d, y = %d\n", x, y)
	x = x ^ y
	y = x ^ y
	x = x ^ y
	fmt.Printf("x = %d, y = %d\n", x, y)
}
