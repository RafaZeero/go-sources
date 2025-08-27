package main

import (
	"fmt"
	"os"
)

func createFile(filename string, content []byte) *os.File {
	f, _ := os.Create(filename)
	n, _ := f.Write(content)
	fmt.Println("file written with bytes:", n)
	return f
}

func removeFile(filename string) error {
	return os.Remove(filename)
}

func readFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func main() {
	// setup
	filename := "foo.txt"
	// content := "asdjkaspokd"

	byteSize := 300

	var emptyContent []byte
	for range byteSize {
		emptyContent = append(emptyContent, 0)
	}

	// creating file
	// f := createFile(filename, []byte(content))
	f := createFile(filename, emptyContent)
	defer f.Close()
	fmt.Println("file created:", f.Name())

	info, _ := f.Stat()

	// reading file content
	b, err := readFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("file content:", string(b))
	fmt.Println("file size:", info.Size(), "bytes")

	// removing file created
	// if err := removeFile(filename); err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("file removed:", filename)
}
