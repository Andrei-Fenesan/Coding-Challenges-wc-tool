## Description

The resolver takes a domain name as input and returns its corresponding IP address by querying DNS servers directly, instead of relying on system or library calls.

## Objectives

- Implement a recursive DNS resolver from scratch.
- Send queries to root name servers, top-level domain (TLD) name servers, and authoritative name servers.
- Parse DNS responses and extract relevant data.
- Handle common DNS record types (A, AAAA, etc.).
- Optimize for efficiency and robustness.

## Requirements

- Your implementation must support recursive resolution of domain names.
- It should correctly follow the DNS resolution hierarchy.
- The program should return an IPv4 or IPv6 address for a given domain.
- You may use UDP sockets to communicate with DNS servers.
- No reliance on built-in DNS resolution libraries.

## Getting Started

1. Clone this repository:
2. Install dependencies:
3. Go inside the project folder
3. Run the resolver with:
   ```sh
   go run . example.com
   ```

## Example Usage

```sh
$ go run . google.com
Querying l.gtld-servers.net (192.41.162.30)
Querying ns2.google.com (216.239.34.10)
Ips of the: google.com
[0]: 142.251.39.78
```

## Implemented Features

- Querying root, TLD, and authoritative DNS servers.
- Handling different record types (A, AAAA, etc.).
- Error handling for invalid domains or unreachable DNS servers.
- Logging and debugging tools.

## Possible Enhancements

- Support for CNAME record type.
- Implementing DNSSEC validation.
- Supporting both IPv4 and IPv6 transport.
Caching responses to optimize performance.

## References

- [RFC 1034: Domain Names - Concepts and Facilities](https://tools.ietf.org/html/rfc1034)
- [RFC 1035: Domain Names - Implementation and Specification](https://tools.ietf.org/html/rfc1035)