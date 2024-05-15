package server

import (
	"encoding/json"
	"errors"
	"log/slog"
	"os"
	"testing"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/encodedecoder"
	"github.com/VladimirRytov/advertisementer/internal/front/httpsender"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func CreateLogger() {
	logging.CreateLogger(".", 0, &slog.HandlerOptions{Level: slog.LevelInfo}, false, os.Stderr)
}

func connectToDBForTests(dbName string) (*ServerStorage, error) {
	CreateLogger()
	switch dbName {
	case "Server":
		param := &datatransferobjects.ServerDSN{
			Source: "127.0.0.1", UserName: "admin", Password: "admin", Port: 8080}
		mar, err := json.Marshal(&param)
		if err != nil {
			return nil, err
		}
		srv := NewServerStorage(httpsender.NewSender(), encodedecoder.NewBase64Encoder())
		err = srv.ConnectToDatabase(mar)
		if err != nil {
			return nil, err
		}
		return srv, nil
	}
	return nil, errors.New("unexpected error")
}

func TestConnectToServer(t *testing.T) {
	CreateLogger()
	param := &datatransferobjects.ServerDSN{
		Source: "127.0.0.1", UserName: "admin", Password: "admin", Port: 8080}

	mar, err := json.Marshal(&param)
	if err != nil {
		t.Error(err)
	}
	srv := NewServerStorage(httpsender.NewSender(), encodedecoder.NewBase64Encoder())
	err = srv.ConnectToDatabase(mar)
	if err != nil {
		t.Error(err)
	}
	srv.Close()
	err = srv.ConnectToDatabase([]byte("asdas"))
	if err != nil {
		switch err.(type) {
		case *json.SyntaxError:
			t.Logf("got syntax error")
		default:
			t.Fatalf("unexpected error - %T", err)
		}
	} else {
		t.Fatalf("want error")
	}

	param = &datatransferobjects.ServerDSN{
		Source: "127.0.0.1", UserName: "admin", Password: "admin", Port: 8081}
	mar, err = json.Marshal(&param)
	if err != nil {
		t.Error(err)
	}
	err = srv.ConnectToDatabase(mar)
	if err == nil {
		t.Fatal("want error")
	}
}
