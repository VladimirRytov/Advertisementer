package presenter

import (
	"log/slog"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func CreateLogger() {
	logging.CreateLogger(".", 0, &slog.HandlerOptions{Level: slog.LevelInfo}, false, os.Stderr)
}

func TestClientToDto(t *testing.T) {
	dc := NewDataConverter()
	CreateLogger()
	want := datatransferobjects.ClientDTO{
		Name:                  "Вася",
		Phones:                "8-800-555-35-35",
		Email:                 "mail@mail.com",
		AdditionalInformation: "asdasd",
	}
	testCase := &ClientDTO{
		Name:                  "Вася",
		Phones:                "8-800-555-35-35",
		Email:                 "mail@mail.com",
		AdditionalInformation: "asdasd",
	}
	got, err := dc.ClientToDto(testCase)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestOrderToDto(t *testing.T) {
	CreateLogger()
	dc := NewDataConverter()
	want := datatransferobjects.OrderDTO{
		ID:            1,
		ClientName:    "Вася",
		Cost:          123455,
		PaymentType:   "zxc",
		CreatedDate:   time.Date(2023, 12, 12, 0, 0, 0, 0, time.UTC),
		PaymentStatus: true,
	}
	testCase := &OrderDTO{
		ID:            1,
		ClientName:    "Вася",
		Cost:          "1234.55",
		PaymentType:   "zxc",
		CreatedDate:   "12.12.2023",
		PaymentStatus: true,
	}
	got, err := dc.OrderToDto(testCase)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestBlockAdvertisementToDto(t *testing.T) {
	dc := NewDataConverter()
	CreateLogger()
	want := datatransferobjects.BlockAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           1,
			OrderID:      1,
			ReleaseCount: 2,
			Cost:         2223,
			Text:         "asd",
			Tags:         []string{"asd", "dsa"},
			ExtraCharges: []string{"asd", "dsa"},
			ReleaseDates: []time.Time{
				time.Date(2023, 12, 15, 0, 0, 0, 0, time.UTC),
				time.Date(2007, 12, 12, 0, 0, 0, 0, time.UTC),
				time.Date(2000, 5, 22, 0, 0, 0, 0, time.UTC),
				time.Date(2000, 5, 2, 0, 0, 0, 0, time.UTC)},
		},
		Size:     10,
		FileName: "asd.jpg",
	}
	testCase := &BlockAdvertisementDTO{
		Advertisement: Advertisement{
			ID:           1,
			OrderID:      1,
			ReleaseCount: 2,
			ReleaseDates: string("15.12.2023, 12.12.2007, 22.05.2000, 02.05.2000"),
			Cost:         "22.23",
			Text:         "asd",
			Tags:         string("asd, dsa"),
			ExtraCharge:  string("asd, dsa"),
		},
		Size:     10,
		FileName: "asd.jpg",
	}
	got, err := dc.BlockAdvertisementToDto(testCase)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestLineAdvertisementToDto(t *testing.T) {
	CreateLogger()
	dc := NewDataConverter()
	want := datatransferobjects.LineAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           1,
			OrderID:      1,
			ReleaseCount: 2,
			Cost:         2223,
			Text:         "asd",
			Tags:         []string{"asd", "dsa"},
			ExtraCharges: []string{"asd", "dsa"},
			ReleaseDates: []time.Time{
				time.Date(2023, 12, 15, 0, 0, 0, 0, time.UTC),
				time.Date(2007, 12, 12, 0, 0, 0, 0, time.UTC),
				time.Date(2000, 5, 22, 0, 0, 0, 0, time.UTC),
				time.Date(2000, 5, 2, 0, 0, 0, 0, time.UTC)},
		},
	}
	testCase := &LineAdvertisementDTO{
		Advertisement: Advertisement{
			ID:           1,
			OrderID:      1,
			ReleaseCount: 2,
			ReleaseDates: string("15.12.2023, 12.12.2007, 22.05.2000, 02.05.2000"),
			Cost:         "22.23",
			Text:         "asd",
			Tags:         string("asd, dsa"),
			ExtraCharge:  string("asd, dsa"),
		},
	}
	got, err := dc.LineAdvertisementToDto(testCase)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, got %v", want, got)
	}
}
func TestTagToDto(t *testing.T) {
	dc := NewDataConverter()
	CreateLogger()
	want := datatransferobjects.TagDTO{
		TagName: "asd",
		TagCost: 2223,
	}
	testCase := &TagDTO{
		TagName: "asd",
		TagCost: "22.23",
	}
	got, err := dc.TagToDto(testCase)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestExtraChargeToDto(t *testing.T) {
	dc := NewDataConverter()
	CreateLogger()
	want := datatransferobjects.ExtraChargeDTO{
		ChargeName: "asd",
		Multiplier: 200,
	}
	testCase := &ExtraChargeDTO{
		ChargeName: "asd",
		Multiplier: "200",
	}
	got, err := dc.ExtraChargeToDto(testCase)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestLocalDsnToDto(t *testing.T) {
	CreateLogger()
}

func TestNetworkDsnToDto(t *testing.T) {
	CreateLogger()
}

func TestCostRateToDto(t *testing.T) {
	dc := NewDataConverter()
	CreateLogger()
	want := datatransferobjects.CostRateDTO{
		CalcForOneWord:   true,
		Name:             "asd",
		ForOneWordSymbol: 5000,
		ForOneSquare:     1000,
	}
	testCase := CostRateDTO{Name: "asd", OneWordOrSymbol: "50,00", Onecm2: "10,00", CalcForOneWord: true}
	got, err := dc.CostRateToDTO(&testCase)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v,got %v", want, got)
	}
}
