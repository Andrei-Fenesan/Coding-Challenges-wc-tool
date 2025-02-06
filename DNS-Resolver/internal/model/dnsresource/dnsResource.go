package dnsresource

import (
	"dnsresolver/internal/model/utils"
	"encoding/binary"
	"fmt"
)

type DnsResource struct {
	Name         string
	ResourceType [2]byte
	Class        [2]byte
	Ttl          uint32
	RdLenght     uint16
	RData        []byte
}

// ParseReource extracts the resource described in RFC 1035 (https://datatracker.ietf.org/doc/html/rfc1035#section-4.1.3).
//
// response []byte 					 - The response received from DNS
//
// numberOfResourcesInSection uint16 - The number of resources (the value is found in the header section)
//
// startOffset uint16                - the start offset.
// The first byte parsed will be response[startOffset]
//
// Returns a slice of DnsResource and the position where the next section is found. The length of the slice is equal to numberOfResourcesInSection
func ParseReource(response []byte, numberOfResourcesInSection uint16, startOffset uint16) ([]DnsResource, uint16) {
	resources := make([]DnsResource, 0, numberOfResourcesInSection)
	for i := uint16(0); i < numberOfResourcesInSection; i++ {
		resource, nextPos := extreactResource(response, startOffset)
		resources = append(resources, resource)
		startOffset = nextPos
	}

	return resources, startOffset
}

func extreactResource(response []byte, startOffset uint16) (DnsResource, uint16) {
	name, nextPos := utils.ParseName(response, startOffset)
	resourceType := [2]byte(response[nextPos : nextPos+2])
	class := [2]byte(response[nextPos+2 : nextPos+4])
	ttl := binary.BigEndian.Uint32(response[nextPos+4 : nextPos+8])
	rdLength := binary.BigEndian.Uint16(response[nextPos+8 : nextPos+10])
	rdData := response[nextPos+10 : nextPos+10+rdLength]
	resource := DnsResource{
		Name:         name,
		ResourceType: resourceType,
		Class:        class,
		Ttl:          ttl,
		RdLenght:     rdLength,
		RData:        rdData,
	}
	return resource, nextPos + 10 + rdLength
}

// String method reuturns the string representation if the DnsResource
func (r DnsResource) String() string {
	return fmt.Sprintf("{\n Name: %s\n Type: %x\n Class: %x\n ttl: %d\n rdLength: %d\n rData: %x\n}",
		r.Name, r.ResourceType, r.Class, r.Ttl, r.RdLenght, r.RData)
}
