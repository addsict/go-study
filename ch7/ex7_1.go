package main

import (
	"bufio"
	"fmt"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	idx := 0
	for adv, _, err := bufio.ScanWords(p[idx:], true); idx < len(p); idx += adv {
		if err != nil {
			return 0, err
		}
		*c += WordCounter(1)
	}
	return len(p), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	idx := 0
	for adv, _, err := bufio.ScanLines(p[idx:], true); idx < len(p); idx += adv {
		if err != nil {
			return 0, err
		}
		*c += LineCounter(1)
	}
	return len(p), nil
}

func main() {
	var c WordCounter
	fmt.Fprintf(&c, "We're so %s!", "sleepy")
	fmt.Println(c)

	var lc LineCounter
	fmt.Fprintf(&lc, "We're so %s!\nAnd you?\n", "sleepy")
	fmt.Println(lc)
}
