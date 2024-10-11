package web

import (
	"encoding/json"
	"io"
)

func FromJson(data []byte, instance any) {
	json.Unmarshal(data, instance)
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
