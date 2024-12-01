package main

import (
	"bufio"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func CountNumberOfBytes(file *os.File) int {
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

func CountNumberOfLines(file *os.File) int {
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

func CountNumberOfWords(file *os.File) int {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}
	return wordCount
}

func CountNumberOfRunes(file *os.File) int {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	runeCount := 0
	for scanner.Scan() {
		runeCount++
	}
	return runeCount
}

func numberOFWords(byteBuffer []byte, nonSpaceBefore bool) (int, bool, []byte) {
	i := 0
	wordCount := 0
	hadNonSpaceRuneBefore := nonSpaceBefore
	for {
		if i >= len(byteBuffer) {
			break
		}
		r, width := utf8.DecodeRune(byteBuffer[i:])
		if r == utf8.RuneError {
			leftover := make([]byte, len(byteBuffer[i:]))
			copy(leftover, byteBuffer[i:])
			return wordCount, hadNonSpaceRuneBefore, leftover
		}
		if unicode.IsSpace(r) {
			if hadNonSpaceRuneBefore {
				wordCount++
			}
			hadNonSpaceRuneBefore = false
		} else {
			hadNonSpaceRuneBefore = true
		}
		i += width
	}
	return wordCount, hadNonSpaceRuneBefore, nil
}

// returns the number of (lines, words, bytes) in a file
func CountLinesWordsAndBytes(file *os.File) (int, int, int) {
	reader := bufio.NewReader(file)
	bytesBuffer := make([]byte, 4096)
	buff := make([]byte, 0, 10)
	var leftover []byte = nil
	byteCount, lineCount, wordCount := 0, 0, 0
	wordC := 0
	nonSpaceBefore := false
	for {
		readBytes, err := reader.Read(bytesBuffer)
		if err == io.EOF {
			if nonSpaceBefore {
				wordCount++
			}
			break
		}
		leftoverLength := 0
		if leftover != nil {
			leftoverLength = len(leftover)
			buff = append(buff[0:], leftover...)
		}
		buff = append(buff, bytesBuffer...)
		wordC, nonSpaceBefore, leftover = numberOFWords(buff[:readBytes+leftoverLength], nonSpaceBefore)
		buff = nil

		wordCount += wordC
		lineCount += numberOfLineBreaks(bytesBuffer, readBytes)
		byteCount += readBytes
	}

	return lineCount, wordCount, byteCount
}
