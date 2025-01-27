package question

import (
	"encoding/binary"
	"strings"
)

type DnsQuestion struct {
	Name       string
	QueryType  [2]byte
	QueryClass [2]byte
}

func (question *DnsQuestion) Encode() []byte {
	result := make([]byte, 0)
	result = append(result, encodeName(question.Name)...)
	result = append(result, question.QueryType[:]...)
	result = append(result, question.QueryClass[:]...)
	return result
}

func (dnsq *DnsQuestion) SetType(queryType uint16) {
	binary.BigEndian.PutUint16(dnsq.QueryType[:], queryType)
}

func (dnsq *DnsQuestion) SetClass(queryCLass uint16) {
	binary.BigEndian.PutUint16(dnsq.QueryClass[:], queryCLass)
}

func encodeName(name string) []byte {
	labels := strings.Split(name, ".")
	result := make([]byte, 0)
	for _, label := range labels {
		labelSize := len(label)
		result = append(append(result, uint8(labelSize)), []byte(label)...)
	}
	result = append(result, byte(0))
	return result
}
