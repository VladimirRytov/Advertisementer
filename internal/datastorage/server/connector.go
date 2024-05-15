package server

import (
	"bytes"
	"context"
	"io"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/encodedecoder"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

const (
	clients             = "clients"
	orders              = "orders"
	lineadvertisements  = "lineadvertisements"
	blockadvertisements = "blockadvertisements"
	tags                = "tags"
	extraCharges        = "extracharges"
	costRates           = "costrates"
	users               = "users"
	dbs                 = "databases"
	files               = "files"
)

type Base64EncodeDecoder interface {
	ToBase64([]byte) []byte
	FromBase64([]byte) ([]byte, error)
	ToBase64String(in []byte) string
	FromBase64String(source string) ([]byte, error)
	ToBase64Stream(from io.Reader, to io.Writer) error
}

type ServerStorage struct {
	s   Sender
	b64 Base64EncodeDecoder
}

func NewServerStorage(sender Sender, b64 Base64EncodeDecoder) *ServerStorage {
	return &ServerStorage{s: sender, b64: b64}
}

func (s *ServerStorage) ConnectToDatabase(p []byte) error {
	logging.Logger.Info("orm: Start Connecting to Postgresql")
	var dsn datatransferobjects.ServerDSN
	b := bytes.NewBuffer(p)
	err := encodedecoder.FromJSON(&dsn, b)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = s.s.Initialize(ctx, dsn)
	return err
}

func (ds *ServerStorage) Close() error {
	logging.Logger.Info("orm: Closing connection")
	return nil
}

func (ds *ServerStorage) ConnectionInfo() map[string]string {
	return ds.s.ConnectionInfo()
}
