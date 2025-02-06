package utils

import (
	"encoding/binary"
	"strings"
)

// ParseName parses the response starting from startOffset acording to RFC 1035 (https://datatracker.ietf.org/doc/html/rfc1035) taking compression into consideration.
//
// response []byte - the received response.
//
// startOffset     - the start offset.
// The first byte parsed will be response[startOffset].
//
// Returns the DNS name and the next position where the next section is found.
func ParseName(response []byte, startOffset uint16) (string, uint16) {
	nameSectionLength := uint16(0)
	labels := make([]string, 0, 1)
	currentPos := startOffset
	for {
		if currentPos >= uint16(len(response)) {
			break
		}
		if isPointer(response[currentPos]) {
			currentPos = binary.BigEndian.Uint16([]byte{ExtractTheLastSixBits(response[currentPos]), response[currentPos+1]})
			nameSectionLength += 2
		} else {
			length := uint16(response[currentPos])
			if currentPos >= startOffset {
				nameSectionLength += length + 1
			}
			if length == 0 {
				break
			}
			label := response[currentPos+1 : currentPos+1+length]
			labels = append(labels, string(label))
			currentPos += length + 1
		}
	}
	return strings.Join(labels, "."), startOffset + nameSectionLength
}

func isPointer(b byte) bool {
	return StartsWithTheFirstTwoBitsSet(b)
}
