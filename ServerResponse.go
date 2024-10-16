package web

import (
	"net/http"
	"strconv"
)

type ServerResponse struct {
	statusCode int
	header     http.Header
	body       []byte
}

func NewResponse() ServerResponse {
	return ServerResponse{
		header: make(map[string][]string),
	}
}

func (serverResponse *ServerResponse) SetStatusCode(status int) {
	serverResponse.statusCode = status
}

func (serverResponse *ServerResponse) AddContent(content []byte, format string) {
	serverResponse.header.Add(HEADER_CONTENT_TYPE, format)
	serverResponse.header.Add(HEADER_CONTENT_LENGTH, strconv.Itoa(len(content)))
	serverResponse.body = content
}
