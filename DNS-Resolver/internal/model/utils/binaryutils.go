package utils

func SetBit(b *byte, pos uint8) {
	*b = (*b) | (1 << pos)
}

func ClearBit(b *byte, pos uint8) {
	*b = (*b) & ^(1 << pos)
}
