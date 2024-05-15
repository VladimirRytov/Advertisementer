package encodedecoder

import (
	"encoding/gob"
	"io"
)

func ToGob(out io.Writer, in any) error {
	enc := gob.NewEncoder(out)
	err := enc.Encode(in)
	if err != nil {
		return err
	}
	return nil
}

func FromGob(out any, source io.Reader) error {
	dec := gob.NewDecoder(source)

	err := dec.Decode(out)
	if err != nil {
		return err
	}
	return nil
}
