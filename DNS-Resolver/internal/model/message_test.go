package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeDnsMessage(t *testing.T) {
	assert := assert.New(t)
	encodedMessage := NewQuestion(22, "dns.google.com")

	assert.Equal("00160100000100000000000003646e7306676f6f676c6503636f6d0000010001", fmt.Sprintf("%x", encodedMessage.Encode()))
}
