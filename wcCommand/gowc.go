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

func numberOfLineBreaks(data []byte, size int) int {
	lineBreaks := 0
	for i := 0; i < size; i++ {
		if data[i] == '\n' {
			lineBreaks++
		}
	}
	return lineBreaks
}

func countNumberOfLines(file *os.File) int {
	reader := bufio.NewReader(file)
	bytesBuffer := make([]byte, 4096)
	lines := 0
	for {
		bytesRead, err := reader.Read(bytesBuffer)
		if err == io.EOF {
			break
		}
		lines += numberOfLineBreaks(bytesBuffer, bytesRead)
	}
	return lines
}

func isFlagValid(flag string) bool {
	switch flag {
	case "-c", "-l":
		return true
	}
	return false
}

func main() {
	cmdArgs := os.Args
	if len(cmdArgs) != 3 {
		fmt.Print("Invalid number of arguments")
		return
	}
	flag := cmdArgs[1]
	filePath := cmdArgs[2]
	if isFlagValid(flag) {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Print("No file found: " + filePath)
			return
		}
		defer file.Close()

		if flag == "-c" {
			fmt.Printf("%d %s", countNumberOfBytes(file), filePath)
		} else if flag == "-l" {
			fmt.Printf("%d %s\n", countNumberOfLines(file), filePath)
		}
	}
}
