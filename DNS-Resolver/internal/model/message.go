package model

import (
	"dnsresolver/internal/model/dnsheader"
	dnsquestion "dnsresolver/internal/model/dnsquestion"
	dnsresource "dnsresolver/internal/model/dnsresource"
	"dnsresolver/internal/model/utils"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Message struct {
	header     dnsheader.DnsHeader
	question   []dnsquestion.DnsQuestion
	answer     []dnsresource.DnsResource
	authority  []dnsresource.DnsResource
	additional []dnsresource.DnsResource
}

const DEFAULT_ROOT_SERVER = "198.41.0.4"
const DNS_PORT = ":53"

// NewQuestion creates a message that represents a question
//
// id uint16 - the id of the message
//
// domainName string - the domain name you are searching for
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
	header.SetRecursion(false)
	question.SetClass(1)
	question.SetType(1)
	msg := Message{
		header:   header,
		question: []dnsquestion.DnsQuestion{question},
	}
	return &msg
}

// resolveName queries the DSN and tries to find the IP of the domain name contained in the question section
//
// msg *Message - the message that contains the query
//
// returns:
// []string - the ips associated with the domain name
// error    -  the enocuntered error if any, nil otherwise
func (msg *Message) ResolveName() ([]string, error) {
	return msg.resolveName(DEFAULT_ROOT_SERVER + DNS_PORT)
}

func (msg *Message) resolveName(ipAddress string) ([]string, error) {
	response, responseSize, err := msg.sendMessage(ipAddress)
	if err != nil {
		return nil, err
	}

	responseMessage := ParseResponse(response[:responseSize])
	if errorCode := responseMessage.header.ErrorCode(); errorCode != 0 {
		return nil, errors.New("Error code received: " + string(errorCode))
	}
	if responseMessage.header.Id != msg.header.Id {
		return nil, errors.New("INVALID ID RECEIVED")
	}

	if len(responseMessage.answer) != 0 {
		return responseMessage.extractAllIpsFromAnswers(), nil
	}

	for _, authority := range responseMessage.authority {
		if authority.IsNsType() {
			domainNameAuthority, _ := utils.ParseName(response, authority.RDataStartOffset)
			domainNameAuthorityIp := responseMessage.findDomainIpInAdditionalSection(domainNameAuthority)
			if domainNameAuthorityIp != "" {
				fmt.Printf("Querying %s (%s)\n", domainNameAuthority, domainNameAuthorityIp)
				ipAddress, err := msg.resolveName(domainNameAuthorityIp + DNS_PORT)
				if ipAddress != nil && err == nil {
					return ipAddress, nil
				}
			} else {
				// we have to find the ip of the autority
				fmt.Printf("Will try to find the ips of the autority: %s\n", domainNameAuthority)
				authorityIps, err := NewQuestion(10, domainNameAuthority).ResolveName()
				fmt.Printf("Found ips of the the autority: %s\n", domainNameAuthority)
				if err == nil && authorityIps != nil {
					for _, ip := range authorityIps {
						searchedIp, errResolve := msg.resolveName(ip + DNS_PORT)
						if searchedIp != nil && errResolve == nil {
							return searchedIp, nil
						}
					}
				}
			}
		}
	}

	return nil, nil
}

// sendMessage sends the message to the ipAddress and returns the response
// msg *Message - the message that is going to be send. The message will be ecnoded
// returns:
// []byte - the response received
// int    - the length of the response
// error  - any error that can be encountered (ex: connection do the dns, reading/writing data)
func (msg *Message) sendMessage(ipAddress string) ([]byte, int, error) {
	conn, errDial := net.Dial("udp", ipAddress)
	if errDial != nil {
		fmt.Printf("Error establishing connection: %v", errDial)
		return nil, -1, errDial
	}
	defer conn.Close()

	_, errWrite := conn.Write(msg.Encode())
	if errWrite != nil {
		fmt.Printf("Error writing: %v", errDial)
		return nil, -1, errWrite
	}

	response := make([]byte, 512)
	read, errRead := conn.Read(response)
	if errRead != nil {
		return nil, -1, errRead
	}

	return response[:read], read, nil
}

func (msg *Message) findDomainIpInAdditionalSection(domainName string) string {
	for _, additional := range msg.additional {
		if additional.IsAType() && additional.Name == domainName {
			return toIp(additional.RData)
		}
	}
	return ""
}

func toIp(data []byte) string {
	labels := make([]string, 0, 4)
	for _, number := range data {
		labels = append(labels, strconv.Itoa(int(uint8(number))))
	}
	return strings.Join(labels, ".")
}

func (msg *Message) extractAllIpsFromAnswers() []string {
	ips := make([]string, 0, 4)
	for _, ans := range msg.answer {
		if ans.IsAType() {
			ips = append(ips, toIp(ans.RData))
		}
	}
	return ips
}

// ParseResponse parses the response received into a Message acoring to RFC 1035.
//
// response []byte - represnets the response received from the DNS
func ParseResponse(response []byte) *Message {
	header := *dnsheader.Decode([12]byte(response[:12]))
	questions, nextPos := dnsquestion.ParseQuestionSection(response, header.QdCount)
	answers, nexPosAns := dnsresource.ParseReource(response, header.AnCount, nextPos)
	authorities, nextPosAuthorities := dnsresource.ParseReource(response, header.NsCount, nexPosAns)
	additionals, _ := dnsresource.ParseReource(response, header.ArCount, nextPosAuthorities)

	return &Message{
		header:     header,
		question:   questions,
		answer:     answers,
		authority:  authorities,
		additional: additionals,
	}
}

func (msg Message) String() string {
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
