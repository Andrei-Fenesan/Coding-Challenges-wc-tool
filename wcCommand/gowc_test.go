package main

import (
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

	assert.Equal(336747, countNumberOfBytes(file))
}

func TestFileLineCount(t *testing.T) {
	assert := assert.New(t)

	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("Cannot open file")
	}
	defer file.Close()

	assert.Equal(8850, countNumberOfLines(file))
}

func TestFileWordCount(t *testing.T) {
	assert := assert.New(t)

	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatalf("Cannot open file")
	}
	defer file.Close()

	assert.Equal(58164, countNumberOfWords(file))
}
