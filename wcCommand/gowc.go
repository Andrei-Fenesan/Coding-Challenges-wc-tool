package main

import (
	"fmt"
	"os"
)

func isSupportedFlag(flag string) bool {
	switch flag {
	case "-c", "-l", "-w", "-m", "":
		return true
	}
	return false
}

func registerHandlers() map[string]func(*os.File) {
	handlers := make(map[string]func(*os.File))
	handlers["-c"] = func(file *os.File) {
		fmt.Printf("%d %s", CountNumberOfBytes(file), file.Name())
	}

	handlers["-l"] = func(file *os.File) {
		fmt.Printf("%d %s\n", CountNumberOfLines(file), file.Name())
	}

	handlers["-w"] = func(file *os.File) {
		fmt.Printf("%d %s\n", CountNumberOfWords(file), file.Name())
	}

	handlers["-m"] = func(file *os.File) {
		fmt.Printf("%d %s\n", CountNumberOfRunes(file), file.Name())
	}
	handlers[""] = func(file *os.File) {
		lines, words, bytes := CountLinesWordsAndBytes(file)
		fmt.Printf("%d %d %d %s\n", lines, words, bytes, file.Name())
	}
	return handlers
}

func main() {
	cmdArgs := os.Args

	flag := ""
	filePath := ""
	if len(cmdArgs) < 2 {
		panic("Invalid number of args")
	}
	if len(cmdArgs) == 2 {
		//default
		filePath = cmdArgs[1]
	}
	if len(cmdArgs) == 3 {
		flag = cmdArgs[1]
		filePath = cmdArgs[2]
	}

	handlers := registerHandlers()

	if isSupportedFlag(flag) {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Print("No file found: " + filePath)
			return
		}
		defer file.Close()

		handlers[flag](file)
	}
}
