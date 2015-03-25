package main

import (
	"fmt"
)

func printArr4(arr [4]int) {
	fmt.Println(arr)
}

func modifyArr4(arr [4]int) {
	for i, _ := range arr {
		arr[i] = arr[i] * 2
	}
}

func modifyArr4P(arr *[4]int) {
	for i, _ := range arr {
		arr[i] = arr[i] * 2
	}
}

func main() {
	arr4 := [4]int{1, 2, 3, 4}
	printArr4(arr4)

	modifyArr4(arr4)
	printArr4(arr4)

	modifyArr4P(&arr4)
	printArr4(arr4)
	// arr5 := [5]int{1, 2, 3, 4, 5}
	// printArr4(arr5)
}
