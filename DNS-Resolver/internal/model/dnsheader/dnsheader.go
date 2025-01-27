package dnsheader

import (
	. "dnsresolver/internal/model/utils"
	"encoding/binary"
	"fmt"
)

type DnsHeader struct {
	Id      uint16
	Flags   [2]byte
	QdCount uint16
	AnCount uint16
	NsCount uint16
	ArCount uint16
}

func (header *DnsHeader) Encode() []byte {
	result := make([]byte, 0)
	result = binary.BigEndian.AppendUint16(result, header.Id)
	result = append(result, header.Flags[:]...)
	result = binary.BigEndian.AppendUint16(result, header.QdCount)
	result = binary.BigEndian.AppendUint16(result, header.AnCount)
	result = binary.BigEndian.AppendUint16(result, header.NsCount)
	result = binary.BigEndian.AppendUint16(result, header.ArCount)
	return result
}

func Decode(header [12]byte) *DnsHeader {
	id := binary.BigEndian.Uint16(header[:2])
	flags := header[2:4]
	qdCount := header[4:6]
	anCount := header[6:8]
	nsCount := header[8:10]
	arCount := header[10:12]
	dnsHeaderResponse := DnsHeader{
		Id:      id,
		Flags:   [2]byte(flags),
		QdCount: binary.BigEndian.Uint16(qdCount),
		AnCount: binary.BigEndian.Uint16(anCount),
		NsCount: binary.BigEndian.Uint16(nsCount),
		ArCount: binary.BigEndian.Uint16(arCount)}
	return &dnsHeaderResponse
}

func (header *DnsHeader) Print() {
	fmt.Printf("Id: %d\nflags: %x\nQuestion count: %d\nAns count: %d\nName server resource count: %d\nAdditional Resource count: %d\n",
		header.Id,
		header.Flags,
		header.QdCount,
		header.AnCount,
		header.NsCount,
		header.ArCount)
}

func (header *DnsHeader) SetQR(question bool) {
	if question {
		ClearBit(&header.Flags[0], 7)
	} else {
		SetBit(&header.Flags[0], 7)
	}
}

func (header *DnsHeader) SetRecursion(recursion bool) {
	if recursion {
		SetBit(&header.Flags[0], 0)
	} else {
		ClearBit(&header.Flags[0], 0)
	}
}
