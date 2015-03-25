package main

import (
	"fmt"
)

func printSlice(s []int) {
	fmt.Println(s)
}

func modifySlice(s []int) {
	for i, _ := range s {
		s[i] = s[i] * 2
	}
}

func main() {
	s := []int{1, 2, 3, 4}
	printSlice(s)

	modifySlice(s)
	printSlice(s)

	fmt.Println(s, len(s), cap(s))

	arr := make([]int, 4)
	fmt.Println(arr, len(arr), cap(arr))
	arr = append(arr, 5)
	fmt.Println(arr, len(arr), cap(arr))
	arr[5] = 6
	// arr[6] = 7
	// arr[7] = 8
	fmt.Println(arr, len(arr), cap(arr))
}
