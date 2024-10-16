package web

import (
	"errors"
	"net/http"
	"strconv"
)

const (
	UnexpectedType string = "unexpected content-type on client request"
)

func ReadJson(request *http.Request) ([]byte, error) {
	contentType := request.Header.Get("Content-Type")
	if contentType != "application/json" {
		return make([]byte, 0), errors.New(UnexpectedType)
	}
	sizeStr := request.Header.Get("Content-Length")
	reader := request.Body
	size, _ := strconv.Atoi(sizeStr)
	buf := make([]byte, size)
	reader.Read(buf)
	return buf, nil
}
