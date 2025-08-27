package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// reads from stdin -> it writes in a file and reads from it /dev/stdin
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		token := scanner.Text()

		if token == "" {
			continue
		}

		fmt.Println("echo:", token)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading std input:", err)
	}
}
