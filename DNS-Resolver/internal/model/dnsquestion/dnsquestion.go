package dnsquestion

import (
	"dnsresolver/internal/model/utils"
	"encoding/binary"
	"fmt"
	"strings"
)

type DnsQuestion struct {
	Name       string
	QueryType  [2]byte
	QueryClass [2]byte
}

// Encode encodes the question into a []byte that can be transmited on netowrk.
//
// question *DnsQuestion - The question that is going to be encoded.
//
// Returns the question that is encoded acording to https://datatracker.ietf.org/doc/html/rfc1035#section-4.1.2
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

// ParseQuestionSection extracts the question described in RFC 1035 (https://datatracker.ietf.org/doc/html/rfc1035#section-4.1.2).
//
// response []byte 			- The response received from DNS
//
// numberOfQuestions uint16 - The number of questions (the value is found in the header section)
//
// Returns a slice of DnsQuestion and the position where the next section is found. The length of the slice is equal to numberOfQuestions
func ParseQuestionSection(response []byte, numberOfQuestions uint16) ([]DnsQuestion, uint16) {
	startPos := uint16(12) // question always starts at position 12 because the previous section ends at position 11
	questions := make([]DnsQuestion, 0, 1)
	for i := uint16(0); i < numberOfQuestions; i++ {
		question, nextPos := parseQuestionFromResponse(response, startPos)
		questions = append(questions, question)
		startPos = nextPos
	}
	return questions, startPos
}

func parseQuestionFromResponse(response []byte, startPos uint16) (DnsQuestion, uint16) {
	name, nextPos := utils.ParseName(response, startPos)
	qType := response[nextPos : nextPos+2]
	qClass := response[nextPos+2 : nextPos+4]
	return DnsQuestion{
		Name:       name,
		QueryType:  [2]byte(qType),
		QueryClass: [2]byte(qClass),
	}, nextPos + 4
}

func (q DnsQuestion) String() string {
	return fmt.Sprintf("{\n Name: %s\n QueryType: %x\n QueryClass: %x\n}", q.Name, q.QueryType, q.QueryClass)
}
