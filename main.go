package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	BUFFER_SIZE = 4096
)

func main() {
	argsLen := len(os.Args)
	switch argsLen {
	case 0:
		runPrompt()
	case 2:
		path := os.Args[1]
		runFile(path)
	default:
		fmt.Println("Usage: fox [script]")
		os.Exit(1)
	}
}

func runFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := make([]byte, BUFFER_SIZE)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)

		}
		fmt.Println(n, "bytes read", string(buffer[:n]))
	}

	fmt.Println("bytes read: ", string(buffer))

}

func runPrompt() {

}
