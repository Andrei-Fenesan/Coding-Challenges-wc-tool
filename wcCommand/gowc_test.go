package main

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileByteCount(t *testing.T) {
	assert := assert.New(t)

	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("Cannot open file")
	}
	defer file.Close()

	assert.Equal(336747, CountNumberOfBytes(file))
}

func TestFileLineCount(t *testing.T) {
	assert := assert.New(t)

	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("Cannot open file")
	}
	defer file.Close()

	assert.Equal(8850, CountNumberOfLines(file))
}

func TestFileWordCount(t *testing.T) {
	assert := assert.New(t)

	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("Cannot open file")
	}
	defer file.Close()

	assert.Equal(58164, CountNumberOfWords(file))
}

func TestFileRuneCount(t *testing.T) {
	assert := assert.New(t)

	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("Cannot open file")
	}
	defer file.Close()

	assert.Equal(333851, CountNumberOfRunes(file))
}

func TestNumberOfLinesWordsAndBytes(t *testing.T) {
	assert := assert.New(t)

	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("Cannot open file")
	}
	defer file.Close()

	lines, words, bytes := CountLinesWordsAndBytes(file)
	assert.Equal(8850, lines)
	assert.Equal(58164, words)
	assert.Equal(336747, bytes)
}

func TestNumberOfWordsTheTheLastCharacterIsNotEndLine(t *testing.T) {
	assert := assert.New(t)
	file := createAndWriteToTempFile("asd\n")
	defer os.Remove(file.Name())
	fileRead, err := os.Open(file.Name())
	if err != nil {
		log.Fatal(err)
	}
	defer fileRead.Close()

	_, words, _ := CountLinesWordsAndBytes(fileRead)
	assert.Equal(1, words)
}

func createAndWriteToTempFile(s string) *os.File {
	file, err := os.CreateTemp(".", "testfile")
	if err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(file)
	w.WriteString(s)
	w.Flush()
	return file
}
