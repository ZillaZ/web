package web

import "strconv"

type ServerResponse struct {
	statusCode int
	header     map[string]string
	body       []byte
}

func NewResponse() ServerResponse {
	return ServerResponse{
		header: make(map[string]string),
	}
}

func (serverResponse *ServerResponse) SetStatusCode(status int) {
	serverResponse.statusCode = status
}

func (serverResponse *ServerResponse) AddJson(content string) {
	serverResponse.header["Content-Type"] = "application/json"
	serverResponse.header["Content-Lenght"] = strconv.Itoa(len(content))
	serverResponse.body = []byte(content)
}
