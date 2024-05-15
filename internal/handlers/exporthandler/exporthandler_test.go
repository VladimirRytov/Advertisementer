package exporthandler

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datastorage/sql/orm"
	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/filestorage"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func CreateLogger() {
	logging.CreateLogger(".", 0, &slog.HandlerOptions{Level: slog.LevelInfo}, false, os.Stderr)
}

func TestExportJsonToFile(t *testing.T) {
	CreateLogger()
	context, cancel := context.WithCancel(context.Background())
	defer cancel()
	var err error
	param := &datatransferobjects.LocalDSN{Path: "", Name: ":memory:", Type: orm.Sqlite}
	mar, err := json.Marshal(&param)
	if err != nil {
		t.Fatal(err)
	}
	db := orm.NewDataStorageOrm("Sqlite")
	err = db.ConnectToDatabase(mar)
	if err != nil {
		t.Fatal(err)
	}
	exh := NewExportHandler(&responcerDummy{}, db, &filestorage.Storage{})

	err = exh.ExportJsonToFile(context, "export.json")
	if err != nil {
		t.Fatal(err)
	}
}

func TestExportJsonToFileCancel(t *testing.T) {
	CreateLogger()
	var err error
	param := &datatransferobjects.LocalDSN{Path: "", Name: ":memory:", Type: orm.Sqlite}
	mar, err := json.Marshal(&param)
	if err != nil {
		t.Fatal(err)
	}
	db := orm.NewDataStorageOrm("Sqlite")
	err = db.ConnectToDatabase(mar)
	if err != nil {
		t.Fatal(err)
	}
	exh := NewExportHandler(&responcerDummy{}, db, &filestorage.Storage{})
	context, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = exh.ExportJsonToFile(context, "export.json")
	if err == nil {
		t.Fatal(err)
	}
}
