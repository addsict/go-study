package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}

func main() {
	c := Celsius(40)
	f := CtoF(c)

	fmt.Println(c)
	fmt.Println(f)
}
