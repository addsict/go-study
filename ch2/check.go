package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	v, ok := m["test"]

	fmt.Println(v)
	if !ok {
		fmt.Println("not exist")
	}

	// inter, ok := m.(string)
	// if ok {
	// fmt.Println("type ok")
	// }
}
