package costcalculationhandler

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"testing"

	"github.com/VladimirRytov/advertisementer/internal/datastorage/sql/orm"
	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func CreateLogger() {
	logging.CreateLogger(".", 0, &slog.HandlerOptions{Level: slog.LevelInfo}, false, os.Stderr)
}
func TestCreateDBForTests(t *testing.T) {
	CreateLogger()

	_, err := CreateDBForTests()
	if err != nil {
		t.Fatal(err)
	}
}
func TestSetActiveCostRate(t *testing.T) {
	var err error
	CreateLogger()

	db, err := CreateDBForTests()
	if err != nil {
		t.Fatal(err)
	}
	calc := NewCostRateCalculator()
	calc.InitCostRateCalculator(&responcerDummy{}, db)
	calc.SetActiveCostRate(costRateDto.Name)
	if ErrorHandler != nil {
		t.Fatal(ErrorHandler)
	}
	db.Close()
}

func TestCalculateBlockAdvertisementCost(t *testing.T) {
	var err error
	CreateLogger()

	db, err := CreateDBForTests()
	if err != nil {
		t.Fatal(err)
	}
	calc := NewCostRateCalculator()
	calc.InitCostRateCalculator(&responcerDummy{}, db)

	calc.SetActiveCostRate(costRateDto.Name)
	if ErrorHandler != nil {
		t.Fatal(ErrorHandler)
	}

	calc.CalculateBlockAdvertisementCost(blockDtoForTest)
	if ErrorHandler != nil {
		t.Fatal(ErrorHandler)
	}
	if Cost != 2664 {
		t.Fatal(Cost, blockDtoForTest)
	}
	db.Close()
}

func TestCalculateLineAdvertisementCostWordCount(t *testing.T) {
	var err error
	CreateLogger()

	db, err := CreateDBForTests()
	if err != nil {
		t.Fatal(err)
	}
	calc := NewCostRateCalculator()
	calc.InitCostRateCalculator(&responcerDummy{}, db)

	calc.SetActiveCostRate(costRateDto.Name)
	if ErrorHandler != nil {
		t.Fatal(ErrorHandler)
	}

	calc.CalculateLineAdvertisementCost(lineDtoForTest)
	if ErrorHandler != nil {
		t.Fatal(ErrorHandler)
	}
	if Cost != 1208 {
		t.Fatalf("want cost = 1208, got cost = %d", Cost)
	}
	db.Close()
}

func TestCalculateLineAdvertisementCostSymbolCount(t *testing.T) {
	var err error
	CreateLogger()

	db, err := CreateDBForTests()
	if err != nil {
		t.Fatal(err)
	}
	calc := NewCostRateCalculator()
	calc.InitCostRateCalculator(&responcerDummy{}, db)

	calc.SetActiveCostRate(costRateDto.Name)
	if ErrorHandler != nil {
		t.Fatal(ErrorHandler)
	}
	calc.activeCostRate.SetCalsForOneWord(false)

	calc.CalculateLineAdvertisementCost(lineDtoForTest)
	if ErrorHandler != nil {
		t.Fatal(ErrorHandler)
	}
	if Cost != 1232 {
		t.Fatalf("want cost = 1232, got cost = %d", Cost)
	}
	db.Close()
}

func TestCalculateNewOrder(t *testing.T) {
	var err error
	CreateLogger()

	db, err := CreateDBForTests()
	if err != nil {
		t.Fatal(err)
	}
	calc := NewCostRateCalculator()
	calc.InitCostRateCalculator(&responcerDummy{}, db)

	calc.SetActiveCostRate(costRateDto.Name)
	if ErrorHandler != nil {
		t.Fatal(ErrorHandler)
	}
	calc.activeCostRate.SetCalsForOneWord(false)
	newOrder := orderDtoForTest
	newOrder.LineAdvertisements = []datatransferobjects.LineAdvertisementDTO{lineDtoForTest}
	newOrder.LineAdvertisements[0].Cost = 100

	newOrder.BlockAdvertisements = []datatransferobjects.BlockAdvertisementDTO{blockDtoForTest}
	newOrder.BlockAdvertisements[0].Cost = 200
	newOrder.ID = 0
	calc.CalculateOrderCost(newOrder)
	if ErrorHandler != nil {
		t.Fatal(ErrorHandler)
	}
	if Cost != 300 {
		t.Fatalf("want cost = 300, got cost = %d", Cost)
	}
	db.Close()
}

func TestCalculateExistedOrder(t *testing.T) {
	var err error
	CreateLogger()
	db, err := CreateDBForTests()
	if err != nil {
		t.Fatal(err)
	}
	calc := NewCostRateCalculator()
	calc.InitCostRateCalculator(&responcerDummy{}, db)

	calc.SetActiveCostRate(costRateDto.Name)
	if ErrorHandler != nil {
		t.Fatal(ErrorHandler)
	}
	calc.activeCostRate.SetCalsForOneWord(false)
	calc.CalculateOrderCost(orderDtoForTest)
	if ErrorHandler != nil {
		t.Fatal(ErrorHandler)
	}
	if Cost != 46 {
		t.Fatalf("want cost = 46, got cost = %d", Cost)
	}
}

func CreateDBForTests() (DataBase, error) {
	context, cancel := context.WithCancel(context.Background())
	defer cancel()
	param := &datatransferobjects.LocalDSN{Name: ":memory:"}

	mar, err := json.Marshal(&param)
	if err != nil {
		return nil, err
	}
	db := orm.NewDataStorageOrm("Sqlite")
	err = db.ConnectToDatabase(mar)
	if err != nil {
		return nil, err
	}
	for _, v := range tagsDtoForTest {
		db.NewTag(context, &v)
	}
	for _, v := range extraChargesDtoForTest {
		db.NewExtraCharge(context, &v)
	}
	_, err = db.NewClient(context, &clientDtoForTest)
	if err != nil {
		return nil, err
	}
	_, err = db.NewAdvertisementsOrder(context, &orderDtoForTest)
	if err != nil {
		return nil, err
	}
	_, err = db.NewBlockAdvertisement(context, &blockDtoForTest)
	if err != nil {
		return nil, err
	}
	_, err = db.NewLineAdvertisement(context, &lineDtoForTest)
	if err != nil {
		return nil, err
	}
	_, err = db.NewCostRate(context, &costRateDto)
	if err != nil {
		return nil, err
	}
	return db, nil
}
