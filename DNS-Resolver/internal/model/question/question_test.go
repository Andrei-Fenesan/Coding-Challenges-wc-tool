package question

import (
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
