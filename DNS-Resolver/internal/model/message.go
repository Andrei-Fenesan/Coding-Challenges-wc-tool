package model

import (
	"dnsresolver/internal/model/dnsheader"
	dnsquestion "dnsresolver/internal/model/dnsquestion"
	dnsresource "dnsresolver/internal/model/dnsresource"
	"fmt"
)

type Message struct {
	header     dnsheader.DnsHeader
	question   []dnsquestion.DnsQuestion
	answer     []dnsresource.DnsResource
	authority  []dnsresource.DnsResource
	additional []dnsresource.DnsResource
}

func NewQuestion(id uint16, domainName string) *Message {
	header := dnsheader.DnsHeader{
		Id:      id,
		QdCount: 1,
		ArCount: 0,
		AnCount: 0,
		NsCount: 0,
	}
	question := dnsquestion.DnsQuestion{
		Name: domainName,
	}
	header.SetQR(true)
	header.SetRecursion(true)
	question.SetClass(1)
	question.SetType(1)
	msg := Message{
		header:   header,
		question: []dnsquestion.DnsQuestion{question},
	}
	return &msg
}

// ParseResponse parses the response received into a Message acoring to RFC 1035.
//
// response []byte - represnets the response received from the DNS
func ParseResponse(response []byte) *Message {
	header := *dnsheader.Decode([12]byte(response[:12]))
	questions, nextPos := dnsquestion.ParseQuestionSection(response, header.QdCount)
	answers, nexPosAns := dnsresource.ParseReource(response, 2, nextPos)
	authorities, nextPosAuthorities := dnsresource.ParseReource(response, header.ArCount, nexPosAns)
	additionals, _ := dnsresource.ParseReource(response, header.NsCount, nextPosAuthorities)

	return &Message{
		header:     header,
		question:   questions,
		answer:     answers,
		authority:  authorities,
		additional: additionals,
	}
}

func (msg *Message) Print() string {
	return fmt.Sprintln(msg.header, msg.question, msg.answer, msg.authority, msg.additional)
}

// Encode encodes the Message into a byte[] that can be send to the DNS.
// Only the questions are encoded and part of the returned byte[]
//
// msg *Message - The message that is going to be encoded
func (msg *Message) Encode() []byte {
	result := make([]byte, 0)
	result = append(result, msg.header.Encode()...)
	for _, question := range msg.question {
		result = append(result, question.Encode()...)
	}
	return result
}
