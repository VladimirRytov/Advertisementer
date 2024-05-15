package server

import (
	"context"
	"io"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
)

type Sender interface {
	Initialize(ctx context.Context, dsn datatransferobjects.ServerDSN) error
	CreateRequest(ctx context.Context, params *datatransferobjects.RequestDTO, data io.Reader) ([]byte, error)
	DeleteRequest(ctx context.Context, params *datatransferobjects.RequestDTO) ([]byte, error)
	GetRequest(ctx context.Context, params *datatransferobjects.RequestDTO, data io.Reader) ([]byte, error)
	UpdateRequest(ctx context.Context, params *datatransferobjects.RequestDTO, data io.Reader) ([]byte, error)
	UploadMultipart(ctx context.Context, params *datatransferobjects.RequestDTO, data io.Reader, contentType string) ([]byte, error)
	ConnectionInfo() map[string]string
}
