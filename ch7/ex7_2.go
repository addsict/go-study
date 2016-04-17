package main

import (
	"fmt"
	"io"
	"os"
)

type countingWriter struct {
	writer io.Writer
	count  *int64
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var count int64
	cw := countingWriter{
		writer: w,
		count:  &count,
	}
	return cw, cw.count
}

func (cw countingWriter) Write(p []byte) (int, error) {
	*cw.count += int64(len(p))
	return cw.writer.Write(p)
}

func main() {
	stdout := os.Stdout
	cw, count := CountingWriter(stdout)
	fmt.Fprintf(cw, "hello %s\n", "world")
	fmt.Printf("%d bytes written\n", *count)

	fmt.Fprintf(cw, "what happend\n")
	fmt.Printf("%d bytes written\n", *count)
}
