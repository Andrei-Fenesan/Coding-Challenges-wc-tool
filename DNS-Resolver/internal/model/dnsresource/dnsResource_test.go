package dnsresource

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDnsResource(t *testing.T) {
	assert := assert.New(t)

	header := "001681800001000200000000"
	questionSection := "03646e7306676f6f676c6503636f6d0000010001"
	answer1 := "c00c0001000100000384000408080404"
	answer2 := "c00c0001000100000384000408080808"
	answerSection := answer1 + answer2
	expectedAnswer1 := DnsResource{
		Name:             "dns.google.com",
		Class:            [2]byte{0x00, 0x01},
		ResourceType:     [2]byte{0x00, 0x01},
		Ttl:              900,
		RdLenght:         4,
		RData:            []byte{0x08, 0x08, 0x04, 0x04},
		RDataStartOffset: 44,
	}
	expectedAnswer2 := DnsResource{
		Name:             "dns.google.com",
		Class:            [2]byte{0x00, 0x01},
		ResourceType:     [2]byte{0x00, 0x01},
		Ttl:              900,
		RdLenght:         4,
		RData:            []byte{0x08, 0x08, 0x08, 0x08},
		RDataStartOffset: 60,
	}

	data := header + questionSection + answerSection
	byteData, _ := hex.DecodeString(data)

	resources, _ := ParseReource(byteData, 2, 12+20)

	assert.Equal(2, len(resources))
	assert.Equal(expectedAnswer1, resources[0])
	assert.Equal(expectedAnswer2, resources[1])
}
