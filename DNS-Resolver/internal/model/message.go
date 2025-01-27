package model

import (
	. "dnsresolver/internal/model/dnsheader"
	. "dnsresolver/internal/model/question"
	. "dnsresolver/internal/model/resource"
)

type Message struct {
	header     DnsHeader
	question   DnsQuestion
	answer     DnsResource
	authority  DnsResource
	additional DnsResource
}

func NewQuestion(id uint16, domainName string) *Message {
	header := DnsHeader{
		Id:      id,
		QdCount: 1,
		ArCount: 0,
		AnCount: 0,
		NsCount: 0,
	}
	question := DnsQuestion{
		Name: domainName,
	}
	header.SetQR(true)
	header.SetRecursion(true)
	question.SetClass(1)
	question.SetType(1)
	msg := Message{
		header:   header,
		question: question,
	}
	return &msg
}

func ParseResponse(response []byte) *Message {
	msg := Message{}
	msg.header = *Decode([12]byte(response[:12]))
	return &msg
}

func (msg *Message) Print() {
	msg.header.Print()
}

func (msg *Message) Encode() []byte {
	result := make([]byte, 0)
	result = append(result, msg.header.Encode()...)
	result = append(result, msg.question.Encode()...)
	return result
}
