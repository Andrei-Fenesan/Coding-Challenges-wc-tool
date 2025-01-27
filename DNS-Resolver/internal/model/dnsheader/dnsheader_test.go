package dnsheader

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeQuestionDnsHeader(t *testing.T) {
	assert := assert.New(t)
	header := DnsHeader{
		Id:      22,
		Flags:   [2]byte{0x00, 0x00},
		QdCount: 1,
		AnCount: 0,
		NsCount: 0,
		ArCount: 0,
	}
	header.SetQR(true)

	assert.Equal("001600000001000000000000", fmt.Sprintf("%x", header.Encode()))
}

func TestEncodeQuestionDnsHeaderWithRecursion(t *testing.T) {
	assert := assert.New(t)
	header := DnsHeader{
		Id:      22,
		Flags:   [2]byte{0x00, 0x00},
		QdCount: 1,
		AnCount: 0,
		NsCount: 0,
		ArCount: 0,
	}
	header.SetQR(true)
	header.SetRecursion(true)

	assert.Equal("001601000001000000000000", fmt.Sprintf("%x", header.Encode()))
}
