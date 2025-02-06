package model

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeDnsMessage(t *testing.T) {
	assert := assert.New(t)
	encodedMessage := NewQuestion(22, "dns.google.com")

	assert.Equal("00160100000100000000000003646e7306676f6f676c6503636f6d0000010001", fmt.Sprintf("%x", encodedMessage.Encode()))
}

func TestParseMessageResponse(t *testing.T) {
	assert := assert.New(t)
	response, err := hex.DecodeString("00168180000100020000000003646e7306676f6f676c6503636f6d0000010001c00c0001000100000081000408080404c00c000100010000008100040808080800")
	if err != nil {
		t.FailNow()
	}
	message := ParseResponse(response)

	assert.NotNil(message.header, message.question, message.answer, message.authority, message.additional)
	assert.Equal(1, len(message.question))
	assert.Equal(2, len(message.answer))
	assert.Empty(message.authority, message.additional)
}
