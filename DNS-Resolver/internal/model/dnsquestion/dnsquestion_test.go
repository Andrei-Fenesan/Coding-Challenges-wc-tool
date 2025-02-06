package dnsquestion

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeQuestion(t *testing.T) {
	assert := assert.New(t)
	question := DnsQuestion{
		Name:       "emag.ro",
		QueryType:  [2]byte{0x00, 0x00},
		QueryClass: [2]byte{0x00, 0x00},
	}

	encodedQuestion := question.Encode()

	assert.Equal("04656d616702726f0000000000", fmt.Sprintf("%x", encodedQuestion))
}

func TestParseQuestionSectionOneQuestion(t *testing.T) {
	assert := assert.New(t)
	headerData := "123434343434343434343434"
	questionData := "03646e7306676f6f676c6503636f6d0000010001" //1 question: dns.google.com
	otherData := "01aa23"
	response, err := hex.DecodeString(headerData + questionData + otherData)
	if err != nil {
		t.FailNow()
	}

	questions, nextPos := ParseQuestionSection(response, 1)

	assert.Equal(1, len(questions))
	assert.Equal(uint16(32), nextPos)
	assert.Equal("dns.google.com", questions[0].Name)
	assert.Equal([2]byte{0x00, 0x01}, questions[0].QueryClass)
}
