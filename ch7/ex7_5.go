package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type limitReader struct {
	reader io.Reader
	n      int64
}

func (r *limitReader) Read(p []byte) (int, error) {
	if r.n < 0 {
		return 0, io.EOF
	}

	bytes := make([]byte, r.n)
	n, err := r.reader.Read(bytes)
	copy(p, bytes)
	if err != nil {
		return n, err
	}
	return n, io.EOF
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{
		reader: r,
		n:      n,
	}
}

func main() {
	r := strings.NewReader("hello world")
	lr := LimitReader(r, 0)

	buf := make([]byte, 128)
	n, err := lr.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Read %d bytes: %s", n, buf[0:n])
}
