package mapper

import (
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/advertisements"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func CreateLogger() {
	logging.CreateLogger(".", 0, &slog.HandlerOptions{Level: slog.LevelInfo}, false, os.Stderr)
}

func TestClientToDTO(t *testing.T) {
	CreateLogger()
	client, err := advertisements.NewClient("asd")
	if err != nil {
		t.Fatal(err)
	}
	err = client.SetContactInformation("88005553535", "mail@mail.com")
	if err != nil {
		t.Fatal(err)
	}
	client.SetAdditionalInformation("asdasdads")

	got := ClientToDTO(&client)
	if got.Email != client.Email() || got.Phones != client.ContactNumber() ||
		got.Name != client.Name() || got.AdditionalInformation != client.AdditionalInformation() {
		t.Fatalf("values not equal, \n%v\n%v", client, got)
	}
}

func TestOrderToDTO(t *testing.T) {
	CreateLogger()
	order, err := advertisements.NewAdvertisementOrder("123")
	if err != nil {
		t.Fatal(err)
	}
	order.SetCreaatedDate(time.Now())
	order.SetOrderId(1)
	order.SetPaymentStatus(true)
	got := OrderToDTO(&order)
	if got.ClientName != order.ClientId() || got.CreatedDate != order.CreaatedDate() {
		t.Fatalf("values not equal, \n%v\n%v", order, got)
	}
}

func TestBlockAdvertisementToDTO(t *testing.T) {
	CreateLogger()
	blockAdv := advertisements.NewAdvertisementBlock()
	blockAdv.SetComment("Вася")
	blockAdv.SetCost(100)
	blockAdv.SetId(1)
	blockAdv.SetOrderId(1)
	blockAdv.SetSize(50)
	blockAdv.AppendTag("Tag A")
	blockAdv.SetFileName("asdasd.er")
	got := BlockAdvertisementToDTO(&blockAdv)
	if got.Cost != blockAdv.Cost() || got.ID != blockAdv.Id() ||
		got.OrderID != blockAdv.OrderId() || got.Text != blockAdv.Comment() || blockAdv.FileName() != got.FileName {
		t.Fatalf("values not equal, \n%v\n%v", got, blockAdv)
	}
}
func TestLineAdvertisementToDTO(t *testing.T) {
	CreateLogger()
	lineAdv := advertisements.NewAdvertisementLine()
	lineAdv.SetContent("Вася")
	lineAdv.SetCost(100)
	lineAdv.SetId(1)
	lineAdv.SetOrderId(1)
	got := LineAdvertisementToDTO(&lineAdv)
	if got.Cost != lineAdv.Cost() || got.ID != lineAdv.Id() ||
		got.OrderID != lineAdv.OrderId() || got.Text != lineAdv.Content() {
		t.Fatalf("values not equal, \n%v\n%v", got, lineAdv)
	}
}
func TestTagToDTO(t *testing.T) {
	CreateLogger()
	tag, err := advertisements.NewTag("asd", 123)
	if err != nil {
		t.Fatal(err)
	}
	got := TagToDTO(&tag)
	if got.TagCost != tag.Cost() || got.TagName != tag.Name() {
		t.Fatalf("values not equal, \n%v\n%v", tag, got)

	}
}
func TestExtraChargeToDTO(t *testing.T) {
	CreateLogger()
	charge, err := advertisements.NewExtraCharge("asd", 132)
	if err != nil {
		t.Fatal(err)
	}
	got := ExtraChargeToDTO(&charge)
	if got.ChargeName != charge.Name() || got.Multiplier != charge.Multiplier() {
		t.Fatalf("values not equal, \n%v\n%v", charge, got)
	}
}

func TestCostRateToDTO(t *testing.T) {
	CreateLogger()
	costRate, err := advertisements.NewCostRate("asd", 1, 2, true)
	if err != nil {
		t.Fatal(err)
	}
	got := CostRateToDTO(&costRate)
	if got.Name != costRate.Name() || got.CalcForOneWord != costRate.CalsForOneWord() ||
		got.ForOneWordSymbol != costRate.CostForWordOrSymbol() || got.ForOneSquare != costRate.CostForOnecm2() {
		t.Fatalf("values not equal, \n%v\n%v", costRate, got)
	}
}
