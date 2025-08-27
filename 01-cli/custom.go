package main

import (
	"bufio"
	"fmt"
	"os"
)

type customReader struct{}

func (cr *customReader) Read(p []byte) (n int, err error) {
	return customStdin.Read(p)
}

func NewCustomReader() *customReader {
	return &customReader{}
}

var (
	customStdin  = os.NewFile(uintptr(0), "./stdin")
	customStdout = os.NewFile(uintptr(1), "./stdout")
	customStderr = os.NewFile(uintptr(2), "./stderr")
)

func main() {
	reader := NewCustomReader()

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		token := scanner.Text()

		if token == "" {
			continue
		}

		fmt.Println("echo:", token)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(customStderr, "reading std input:", err)
	}
}
