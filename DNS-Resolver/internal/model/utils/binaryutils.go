package utils

// Sets to 1 the bit specified by the position parameter in the byte parameter.
//
// b byte - The byte that will have the bit set to 1.
//
// position uint8 - The position in the byte that will have the bit set to 1
func SetBit(b *byte, pos uint8) {
	*b = (*b) | (1 << pos)
}

// Clears (sets to 0) the bit specified by the position parameter in the byte parameter.
//
// b byte - The byte that will have the bit set to 0.
//
// position uint8 - The position in the byte that will have the bit set to 0
func ClearBit(b *byte, position uint8) {
	*b = (*b) & ^(1 << position)
}

// StartsWithTheFirstTwoBitsSet will return true if the byte starts with the bits 11 and false otherwiese.
//
// b byte - The byte from whivh the last six bits will be extracted.
//
// Examples:
//
//	StartsWithTheFirstTwoBitsSet(0b11001010) = true
//
//	StartsWithTheFirstTwoBitsSet(0b10001010) = false
func StartsWithTheFirstTwoBitsSet(b byte) bool {
	return b&(0b11000000) == (0b11000000)
}

// ExtractTheLastSixBits will return a byte that contains the last 6 bits from the byte parameter.
//
// b byte - The byte from which the last six bits will be extracted.
//
// Examples:
//
//	ExtractTheLastSixBits(0b11001010) = 0b00001010
func ExtractTheLastSixBits(b byte) byte {
	return b & (0b00111111)
}
