package encodedecoder

import (
	"encoding/json"
	"io"
)

func ToJSON(out io.Writer, in any, readble bool) error {
	enc := json.NewEncoder(out)
	if readble {
		enc.SetIndent("", "\t")
	}
	err := enc.Encode(in)
	if err != nil {
		return err
	}
	return nil
}

func FromJSON(out any, source io.Reader) error {
	dec := json.NewDecoder(source)
	dec.DisallowUnknownFields()
	err := dec.Decode(out)
	if err != nil {
		return err
	}
	return nil
}
