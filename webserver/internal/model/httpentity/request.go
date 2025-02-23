package httpentity

import (
	"errors"
	"fmt"
	"strings"
)

type Request struct {
	httpVersion string
	httpMethod  HttpMethod
	Path        string
}

func ParseRequest(data []byte) (*Request, error) {
	stringData := string(data)
	firstLF := strings.IndexByte(stringData, '\n')
	if firstLF == -1 {
		return nil, errors.New("invalid request. no line feed")
	}

	firstLine := stringData[0:firstLF]
	parts := strings.Split(firstLine, " ")
	if len(parts) < 3 {
		return nil, errors.New("invalid request")
	}
	method := parts[0]
	path := parts[1]
	version := parts[2]
	httpMethod, err := methodFromString(method)
	if err != nil || !isValidPath(path) || len(version) == 0 {
		return nil, errors.New("invalid request parts")
	}

	return &Request{
		version,
		httpMethod,
		path,
	}, nil
}

func (r *Request) String() string {
	return fmt.Sprintf("%s %s %s\n", r.httpMethod, r.Path, r.httpVersion)
}

func isValidPath(path string) bool {
	return len(path) > 0
}
