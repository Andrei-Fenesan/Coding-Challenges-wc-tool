package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func registerHandlers() map[string]func(io.Reader, string) {
	handlers := make(map[string]func(io.Reader, string))
	handlers["-c"] = func(reader io.Reader, fileName string) {
		fmt.Printf("%d %s", CountNumberOfBytes(reader), fileName)
	}

	handlers["-l"] = func(reader io.Reader, fileName string) {
		fmt.Printf("%d %s\n", CountNumberOfLines(reader), fileName)
	}

	handlers["-w"] = func(reader io.Reader, fileName string) {
		fmt.Printf("%d %s\n", CountNumberOfWords(reader), fileName)
	}

	handlers["-m"] = func(reader io.Reader, fileName string) {
		fmt.Printf("%d %s\n", CountNumberOfRunes(reader), fileName)
	}
	handlers[""] = func(reader io.Reader, fileName string) {
		lines, words, bytes := CountLinesWordsAndBytes(reader)
		fmt.Printf("%d %d %d %s\n", lines, words, bytes, fileName)
	}
	return handlers
}

func extractFlagAndFilePath(cmdArgs []string) (string, string) {
	flag := ""
	filePath := ""
	if len(cmdArgs) == 2 {
		if strings.HasPrefix(cmdArgs[1], "-") {
			flag = cmdArgs[1]
		} else {
			filePath = cmdArgs[1]
		}
	} else if len(cmdArgs) == 3 {
		flag = cmdArgs[1]
		filePath = cmdArgs[2]
	}
	return flag, filePath
}

func main() {

	flag, filePath := extractFlagAndFilePath(os.Args)
	handlers := registerHandlers()

	reader := os.Stdin
	if filePath != "" {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Print("No file found: " + filePath)
			return
		}
		defer file.Close()
		reader = file
	}
	handleReader := handlers[flag]
	if handleReader != nil {
		handleReader(reader, filePath)
	}
}
