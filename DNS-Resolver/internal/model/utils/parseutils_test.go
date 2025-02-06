package utils

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseName(t *testing.T) {
	assert := assert.New(t)
	response, err := hex.DecodeString("03646e7306676f6f676c6503636f6d")
	if err != nil {
		t.FailNow()
	}
	name, nextPos := ParseName(response, 0)
	assert.Equal(uint16(15), nextPos)
	assert.Equal("dns.google.com", name)
}
