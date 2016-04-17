package main

import (
	"flag"
	"fmt"
	"time"
)

type MyFlag struct {
	num int
}

func (f *MyFlag) Set(s string) error {
	var num int
	fmt.Sscanf(s, "%d%", &num)
	f.num = num
	return nil
}

func (f *MyFlag) String() string {
	return fmt.Sprintf("%d", f.num)
}

var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	myflag := &MyFlag{}
	flag.CommandLine.Var(myflag, "myflag", "usage: xxxx")

	flag.Parse()
	fmt.Println(myflag)
	// fmt.Printf("Sleeping for %v...", *period)
	// time.Sleep(*period)
	// fmt.Println()
}
