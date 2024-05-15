package encodedecoder

import (
	"encoding/base64"
	"io"
)

type Base64Encoder struct {
	b64URL *base64.Encoding
	b64    *base64.Encoding
}

func NewBase64Encoder() *Base64Encoder {
	return &Base64Encoder{
		b64URL: base64URLEnc(),
		b64:    base64Enc(),
	}
}

func base64Enc() *base64.Encoding {
	return base64.RawStdEncoding.WithPadding(base64.StdPadding)
}

func base64URLEnc() *base64.Encoding {
	return base64.RawURLEncoding.WithPadding(base64.NoPadding)
}

func (be *Base64Encoder) ToBase64URL(in []byte) []byte {
	encoded := make([]byte, be.b64URL.EncodedLen(len(in)))
	be.b64URL.Encode(encoded, in)
	return encoded
}

func (be *Base64Encoder) ToBase64(in []byte) []byte {
	encoded := make([]byte, be.b64.EncodedLen(len(in)))
	be.b64.Encode(encoded, in)
	return encoded
}

func (be *Base64Encoder) ToBase64Stream(from io.Reader, to io.Writer) error {
	base64Stream := base64.NewEncoder(be.b64, to)
	defer base64Stream.Close()
	_, err := io.Copy(base64Stream, from)
	return err
}

func (be *Base64Encoder) ToBase64URLString(in []byte) string {
	return be.b64URL.EncodeToString(in)
}

func (be *Base64Encoder) ToBase64String(in []byte) string {
	return be.b64.EncodeToString(in)
}

func (be *Base64Encoder) FromBase64URL(source []byte) ([]byte, error) {
	decoded := make([]byte, be.b64URL.DecodedLen(len(source)))
	i, err := be.b64URL.Decode(decoded, source)
	return decoded[:i], err
}

func (be *Base64Encoder) FromBase64(source []byte) ([]byte, error) {
	decoded := make([]byte, be.b64.DecodedLen(len(source)))
	i, err := be.b64.Decode(decoded, source)
	return decoded[:i], err
}

func (be *Base64Encoder) FromBase64String(source string) ([]byte, error) {
	return be.b64.DecodeString(source)
}

func (be *Base64Encoder) FromBase64URLString(source string) ([]byte, error) {
	return be.b64URL.DecodeString(source)
}
