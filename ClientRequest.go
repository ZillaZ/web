package web

import (
	"encoding/json"
	"io"
	"net/http"
)

type LoginRequest struct {
	email    string
	password string
}

func FromJson(data []byte, instance any) {
	json.Unmarshal(data, instance)
}

func Login(request *http.Request) {
	jsonData := ReadJson(request.Body)
	instance := LoginRequest{}
	FromJson(jsonData, &instance)
}

func ReadJson(reader io.ReadCloser) []byte {
	buf := make([]byte, 10)
	bytes, _ := reader.Read(buf)
	for bytes != 0 {
		temp := make([]byte, 256)
		bytes, _ = reader.Read(temp)
		buf = append(buf, temp...)
	}
	return buf
}
