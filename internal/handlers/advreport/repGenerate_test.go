package advreport

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/advertisements"
	"github.com/VladimirRytov/advertisementer/internal/datastorage/dbmanager"
	"github.com/VladimirRytov/advertisementer/internal/datastorage/sql/orm"
	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/logging"
	"github.com/VladimirRytov/advertisementer/internal/mapper"
)

func CreateLogger() {
	logging.CreateLogger(".", 0, &slog.HandlerOptions{Level: slog.LevelDebug}, false, os.Stderr)
}

func TestCollectBlocks(t *testing.T) {
	CreateLogger()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db, err := createDatabaseForTests()
	if err != nil {
		t.Fatal(err)
	}

	par := datatransferobjects.ReportParams{
		ReportType:       "Exel",
		FromDate:         time.Now(),
		ToDate:           time.Now(),
		BlocksFolderPath: "BlockFiles",
		DeployPath:       "gotFiles",
	}
	r := NewReportGenerator()
	r.InitReportGenerator(&responcerDummy{}, db)
	r.params = par
	blocks, err := r.collectBlocks(ctx)
	if err != nil {
		t.Fatal(err)
	}
	var (
		paid    int
		notPaid int
	)
	for i := range blocks {
		switch {
		case blocks[i].Paid && blocks[i].Block.OrderID%2 == 0:
			paid++
		case !blocks[i].Paid && blocks[i].Block.OrderID%2 != 0:
			notPaid++
		}
	}
	if paid+notPaid != len(blocks) {
		t.Fatalf("paid and notPaid summ must be equal blocks length")
	}
}

type Database interface {
	OrderByID(context.Context, int) (datatransferobjects.OrderDTO, error)
	NewClient(context.Context, *datatransferobjects.ClientDTO) (string, error)
	NewAdvertisementsOrder(context.Context, *datatransferobjects.OrderDTO) (datatransferobjects.OrderDTO, error)
	NewLineAdvertisement(context.Context, *datatransferobjects.LineAdvertisementDTO) (int, error)
	NewBlockAdvertisement(context.Context, *datatransferobjects.BlockAdvertisementDTO) (int, error)
	BlockAdvertisementBetweenReleaseDates(ctx context.Context, from, to time.Time) ([]datatransferobjects.BlockAdvertisementDTO, error)
	LineAdvertisementBetweenReleaseDates(ctx context.Context, from, to time.Time) ([]datatransferobjects.LineAdvertisementDTO, error)
}

func createDatabaseForTests() (Database, error) {
	param := &datatransferobjects.LocalDSN{Name: ":memory:"}

	mar, err := json.Marshal(&param)
	if err != nil {
		return nil, err
	}
	dbm := dbmanager.NewDatabaseManager()
	dbm.Register("Sqlite", 0, dbmanager.LocalDB, orm.NewDataStorageOrm(orm.Sqlite))
	db, err := dbm.ConnectToDatabase(orm.Sqlite, mar)
	if err != nil {
		return nil, err
	}

	err = fillTestData(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func fillTestData(db Database) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client, err := advertisements.NewClient("Вася")
	if err != nil {
		return err
	}
	clientDto := mapper.ClientToDTO(&client)
	_, err = db.NewClient(ctx, &clientDto)
	if err != nil {
		return err
	}

	for i := 1; i < 11; i++ {
		order, err := advertisements.NewAdvertisementOrder("Вася")
		if err != nil {
			return err
		}
		order.SetCreaatedDate(time.Now().Add(time.Duration(i * 24 * int(time.Hour))))
		order.SetPaymentStatus(i%2 == 0)
		orderDto := mapper.OrderToDTO(&order)
		_, err = db.NewAdvertisementsOrder(ctx, &orderDto)
		if err != nil {
			return err
		}
	}

	for i := 1; i < 11; i++ {
		line := advertisements.NewAdvertisementLine()
		line.SetOrderId(i)
		line.SetReleaseDates([]time.Time{time.Now().Add(24 * time.Hour)})
		lineDto := mapper.LineAdvertisementToDTO(&line)

		_, err = db.NewLineAdvertisement(ctx, &lineDto)
		if err != nil {
			return err
		}
		block := advertisements.NewAdvertisementBlock()
		block.SetOrderId(i)
		block.SetSize(10)
		block.SetReleaseDates([]time.Time{time.Now().Add(24 * time.Hour)})

		blockDto := mapper.BlockAdvertisementToDTO(&block)
		_, err = db.NewBlockAdvertisement(ctx, &blockDto)
		if err != nil {
			return err
		}
	}
	return nil
}
