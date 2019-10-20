package proto

import (
	"compress/gzip"
	"encoding/json"
	"io"
)

type HuoBiProto struct {
}

func (proto *HuoBiProto) Decode(r io.Reader) (interface{}, error) {
	gzip, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	var resp map[string]interface{}

	decoder := json.NewDecoder(gzip)
	err = decoder.Decode(&resp)

	return resp, err
}
