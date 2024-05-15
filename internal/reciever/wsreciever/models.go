package wsreciever

import "bytes"

type RecieveHandler interface {
	HandleClient(*bytes.Buffer, int) error
	HandleOrder(*bytes.Buffer, int) error
	HandleBlockAdvertisement(*bytes.Buffer, int) error
	HandleLineAdvertisement(*bytes.Buffer, int) error
	HandleTag(*bytes.Buffer, int) error
	HandleExtraCharge(*bytes.Buffer, int) error
	HandleCostRate(*bytes.Buffer, int) error
	ConnectionLost()
	ConnectionRestore()
}
type Encodedecoder interface {
	FromBase64([]byte) ([]byte, error)
}
