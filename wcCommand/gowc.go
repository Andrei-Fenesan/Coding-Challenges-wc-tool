package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func countNumberOfBytes(file *os.File) int {
	reader := bufio.NewReader(file)
	bytesBuffer := make([]byte, 4096)
	bytesCount := 0
	for {
		bytesRead, err := reader.Read(bytesBuffer)
		if err == io.EOF {
			break
		}
		bytesCount += bytesRead
	}
	return bytesCount
}

func main() {
	cmdArgs := os.Args
	if len(cmdArgs) != 3 {
		fmt.Print("Invalid number of arguments")
		return
	}
	flag := cmdArgs[1]
	filePath := cmdArgs[2]
	if flag == "-c" {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Print("No file found: " + filePath)
			return
		}
		defer file.Close()
		fmt.Printf("%d", countNumberOfBytes(file))
	}
}
