package main

import (
	"fmt"
)

func Print(value interface{}) {
	s, ok := value.(string)
	if ok {
		fmt.Printf("value is string: %s\n", s)
	} else {
		fmt.Println("value is not string")
	}
}

func Print2(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Printf("value is string: %s\n", v)
	case int:
		fmt.Printf("values is integer: %d\n", v)
	}
}

func main() {
	Print("abc")
	Print(123)
	Print(false)

	Print2("abc")
	Print2(123)
}
