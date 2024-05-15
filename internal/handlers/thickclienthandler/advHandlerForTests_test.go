package thickclienthandler

import (
	"encoding/json"
	"log/slog"
	"os"

	"github.com/VladimirRytov/advertisementer/internal/datastorage/sql/orm"
	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func CreateControllerForTests() (*AdvertisementController, error) {
	CreateLogger()
	param := &datatransferobjects.NetworkDataBaseDSN{
		Source: "127.0.0.1", DataBase: "gorm_test", UserName: "gorm_test", Password: "gorm_test", SSLMode: false, Port: 5432}
	mar, err := json.Marshal(&param)
	if err != nil {
		return nil, err
	}
	db := orm.NewDataStorageOrm("Postgres")
	err = db.ConnectToDatabase(mar)
	if err != nil {
		return nil, err
	}
	ac := NewAdvertisementController(&responcerDummy{})
	ac.InitAdvertisementController(db)
	return ac, nil
}
func CreateLogger() {
	logging.CreateLogger(".", 0, &slog.HandlerOptions{Level: slog.LevelInfo}, false, os.Stderr)
}
