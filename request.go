package main

import (
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
	rl := &RequestLine{
		Method:        parts[0],
		RequestTarget: parts[1],
		HttpVesion:    parts[2],
	}
	if !rl.Validation() {
		return nil, restOFMSG, UNSUPPORTED_METHOD
	}
	return rl, restOFMSG, nil
}
func (r *RequestLine) Validation() bool {
	return r.HttpVesion == "HTTP/1.1"

}
func RequestFromReader(reader io.Reader) (*Request, error) {
	return &Request{}, nil

}
