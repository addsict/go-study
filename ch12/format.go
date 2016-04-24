package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Bool:
		if v.Bool() {
			return "true"
		} else {
			return "false"
		}
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.String:
		return v.String()
	case reflect.Map:
		return fmt.Sprintf("%d", v.Int()) // panic
	default:
		return v.Type().String() + " value"
	}
}

func main() {
	fmt.Println(Any(true))
	fmt.Println(Any(-100))
	fmt.Println(Any(100))
	fmt.Println(Any(struct{ a int }{10}))
	fmt.Println(Any([3]int{1, 2, 3}))
	fmt.Println(Any("hello"))
	fmt.Println(Any([]time.Duration{time.Since(time.Now())}))
	fmt.Println(Any(map[string]string{"hello": "world"}))
	// fmt.Println(Any(interface {
	// MyMethod() string
	// }))
}
