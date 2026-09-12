package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type RequestLine struct {
	HttpVesion    string
	RequestTarget string
	Method        string
}

type Request struct {
	RequestLine RequestLine
}

var ERROR_BAD_START_LINE = fmt.Errorf("bad start line")

var INCOMPLETE_START_LINE = fmt.Errorf("Incomplete startlne")

var UNSUPPORTED_METHOD = fmt.Errorf("Unsupported http method")

var Separator = "\r\n"

var ERROR_Malformed_Request_Line = fmt.Errorf("malformed request line")

func ParseRequestLine(b string) (*RequestLine, string, error) {
	idx := strings.Index(b, Separator)
	if idx == -1 {
		return nil, b, nil
	}
	startline := b[:idx]
	restOFMSG := b[idx+len(Separator):]
	parts := strings.Split(startline, " ")
	if len(parts) != 3 {
		return nil, restOFMSG, ERROR_BAD_START_LINE
	}

	httpparts := strings.Split(parts[2], "/")
	if len(httpparts) != 2 || httpparts[0] != "HTTP" || httpparts[1] != "1.1" {
		return nil, restOFMSG, ERROR_Malformed_Request_Line
	}
	rl := &RequestLine{
		Method:        parts[0],
		RequestTarget: parts[1],
		HttpVesion:    parts[1],
	}

	return rl, restOFMSG, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, errors.Join(
			fmt.Errorf("Unable to io.ReadAll"),
			err,
		)
	}
	str := string(data)
	rl, _, err := ParseRequestLine(str)

	if err != nil {
		return nil, err
	}

	return &Request{
		RequestLine: *rl,
	}, err

}
