package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Printf("%T\n", 3)

	// v := reflect.ValueOf(3)
	v := reflect.ValueOf("hello")
	fmt.Println(v.String())
	fmt.Println(v.Type())
	fmt.Println(v.Interface())
	// fmt.Println(v.Int())
}
