package thinclienthandler

import (
	"encoding/json"
	"log/slog"
	"os"

	"github.com/VladimirRytov/advertisementer/internal/datastorage/server"
	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/encodedecoder"
	"github.com/VladimirRytov/advertisementer/internal/filestorage"
	"github.com/VladimirRytov/advertisementer/internal/front/httpsender"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func CreateControllerForTests() (*AdvertisementController, error) {
	CreateLogger()
	param := &datatransferobjects.ServerDSN{
		Source: "127.0.0.1", UserName: "admin", Password: "admin", Port: 8080}
	mar, err := json.Marshal(&param)
	if err != nil {
		return nil, err
	}
	db := server.NewServerStorage(httpsender.NewSender(), encodedecoder.NewBase64Encoder())

	err = db.ConnectToDatabase(mar)
	if err != nil {
		return nil, err
	}
	ac := NewAdvertisementController(&responcerDummy{}, &filestorage.Storage{})
	ac.InitAdvertisementController(db)

	return ac, nil
}
func CreateLogger() {
	logging.CreateLogger(".", 0, &slog.HandlerOptions{Level: slog.LevelInfo}, false, os.Stderr)
}
