package dnsresource

type DnsResource struct {
	name         []byte
	resourceType [2]byte
	class        [2]byte
	ttl          uint32
	rdLenght     uint16
	rData        []byte
}

func parseReource(resource []byte) *DnsResource {
	dnsResource := DnsResource{}

	return &dnsResource
}
