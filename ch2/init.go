package main

import "fmt"

func init() {
	fmt.Println("hello init")
}

func main() {
	if false {
		fmt.Println("false")
	}
}
