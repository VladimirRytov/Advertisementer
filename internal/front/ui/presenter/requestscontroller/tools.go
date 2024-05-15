package requestscontroller

import (
	"bytes"
	"encoding/json"
)

func (r *RequestsHandler) encodeData(data any) ([]byte, error) {
	rawData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return rawData, nil
}

func (r *RequestsHandler) decodeData(data []byte, pattern any) error {
	var reader = bytes.NewReader(data)
	dec := json.NewDecoder(reader)
	dec.DisallowUnknownFields()
	err := dec.Decode(pattern)
	if err != nil {
		return err
	}
	return nil
}
