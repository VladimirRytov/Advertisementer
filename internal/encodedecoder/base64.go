package encodedecoder

import (
	"encoding/base64"
	"io"
)

func ToBase64URL(in []byte) []byte {
	enc := base64.RawURLEncoding.WithPadding(base64.NoPadding)
	encoded := make([]byte, enc.EncodedLen(len(in)))
	enc.Encode(encoded, in)
	return encoded
}

func ToBase64(in []byte) []byte {
	enc := base64.RawStdEncoding.WithPadding(base64.StdPadding)
	encoded := make([]byte, enc.EncodedLen(len(in)))
	enc.Encode(encoded, in)
	return encoded
}

func ToBase64String(in []byte) string {
	enc := base64.RawStdEncoding.WithPadding(base64.StdPadding)
	return enc.EncodeToString(in)
}

func ToBase64Stream(from io.Reader, to io.Writer) error {
	enc := base64.RawStdEncoding.WithPadding(base64.StdPadding)
	base64Stream := base64.NewEncoder(enc, to)
	defer base64Stream.Close()
	_, err := io.Copy(base64Stream, from)
	return err
}

func FromBase64URL(source []byte) ([]byte, error) {
	dec := base64.RawURLEncoding.WithPadding(base64.NoPadding)
	decoded := make([]byte, dec.DecodedLen(len(source)))
	i, err := dec.Decode(decoded, source)
	return decoded[:i], err
}

func FromBase64(source []byte) ([]byte, error) {
	dec := base64.RawStdEncoding.WithPadding(base64.StdPadding)
	decoded := make([]byte, dec.DecodedLen(len(source)))
	i, err := dec.Decode(decoded, source)
	return decoded[:i], err
}

func FromBase64String(source string) ([]byte, error) {
	dec := base64.RawStdEncoding.WithPadding(base64.StdPadding)
	return dec.DecodeString(source)
}
