package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByteStartsWithFirstTwoBitsSetShouldReturnTrueWhenTheFirstTwoBitsAreSet(t *testing.T) {
	assert := assert.New(t)

	assert.True(StartsWithTheFirstTwoBitsSet(0b11001010))
}

func TestByteStartsWithFirstTwoBitsSetShouldReturnFalseWhenTheFirstTwoBitsAreNotSet(t *testing.T) {
	assert := assert.New(t)

	assert.False(StartsWithTheFirstTwoBitsSet(0b01001010))
}

func TestExtractTheLastSixBits(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(byte(0b00011110), byte(ExtractTheLastSixBits(0b11011110)))
}
