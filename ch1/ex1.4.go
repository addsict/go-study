package main

import (
	"bufio"
	"fmt"
	"os"
)

type Line struct {
	count int
	files map[string]int
}

func main() {
	counts := make(map[string]*Line)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, f := range files {
			fd, err := os.Open(f)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(fd, counts)
			fd.Close()
		}
	}

	for key, c := range counts {
		if c.count > 1 {
			fmt.Printf("%d\t%s\n", c.count, key)
			for file, _ := range c.files {
				fmt.Println(file)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]*Line) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		if counts[text] == nil {
			counts[text] = &Line{
				count: 0,
				files: map[string]int{},
			}
		}
		counts[text].files[f.Name()] = 1
		counts[text].count++
	}
}
