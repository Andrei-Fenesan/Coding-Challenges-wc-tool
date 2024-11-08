package main

import (
	"fmt"
	"io"
	"os"
)

func countNumberOfBytes(file *os.File) int {
	bytesBuffer := make([]byte, 1000)
	bytesCount := 0
	bytesRead, err := file.Read(bytesBuffer)
	for {
		if err == io.EOF {
			break
		}
		bytesCount += bytesRead
		bytesRead, err = file.Read(bytesBuffer)
	}

	return bytesCount
}

func main() {
	cmdArgs := os.Args
	if len(cmdArgs) != 3 {
		panic("Invalid number of arguments")
	}
	flag := cmdArgs[1]
	filePath := cmdArgs[2]
	if flag == "-c" {
		file, err := os.Open(filePath)
		if err != nil {
			panic("No file found: " + filePath)
		}
		defer file.Close()
		fmt.Printf("%d", countNumberOfBytes(file))
	}
}
