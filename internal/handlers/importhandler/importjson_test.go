package importhandler

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"testing"

	"github.com/VladimirRytov/advertisementer/internal/datastorage/sql/orm"
	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/filestorage"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func CreateLogger() {
	logging.CreateLogger(".", 0, &slog.HandlerOptions{Level: slog.LevelInfo}, false, os.Stderr)
}

func TestImportJson(t *testing.T) {
	var err error
	CreateLogger()
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
	Import := CreateImporter(&responcerDummy{}, presenter.NewDataConverter(), &filestorage.Storage{}, db)
	p := datatransferobjects.ImportParams{
		AllBlocks:       true,
		AlllLines:       false,
		AllTags:         false,
		ActualClients:   false,
		AllExtraCharges: false,
		AllCostRates:    true,
		ThickMode:       true,
	}
	err = Import.ImportJson(context.Background(), "02.12.2022.json", p)
	if err != nil {
		t.Fatal(err)
	}
}
