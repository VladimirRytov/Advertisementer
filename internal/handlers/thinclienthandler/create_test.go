package thinclienthandler

import (
	"testing"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
)

func TestNewClient(t *testing.T) {
	ClientsError = nil
	client := datatransferobjects.ClientDTO{
		Name:                  "Вася",
		Phones:                "8-800-555-35-35",
		Email:                 "mail@mail.com",
		AdditionalInformation: "asdasd",
	}

	req, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	req.NewClient(client)
	if ClientsError != nil {
		t.Fatal(ClientsError)
	}
}

func TestNewOrder(t *testing.T) {
	OrdersError = nil
	order := datatransferobjects.OrderDTO{
		ID:            1,
		ClientName:    "Вася",
		Cost:          10,
		PaymentType:   "asd",
		CreatedDate:   time.Now(),
		PaymentStatus: true,
	}

	req, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	req.NewAdvertisementsOrder(order)
	if OrdersError != nil {
		t.Fatal(OrdersError)
	}
}

func TestNewTag(t *testing.T) {
	TagsError = nil
	tagDTO := datatransferobjects.TagDTO{TagName: "tag A", TagCost: 12}

	req, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	req.NewTag(tagDTO)
	if TagsError != nil {
		t.Fatal(TagsError)
	}
}

func TestNewExtraCharge(t *testing.T) {
	ExtraChargesError = nil
	extraChargeDTO := datatransferobjects.ExtraChargeDTO{ChargeName: "charge A", Multiplier: 2}

	req, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	req.NewExtraCharge(extraChargeDTO)
	if ExtraChargesError != nil {
		t.Fatal(err)
	}
}

func TestNewBlockAdvertisement(t *testing.T) {
	BlockAdvertisementsError = nil
	blockDtoForTest := datatransferobjects.BlockAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           1,
			OrderID:      1,
			ReleaseCount: 1,
			Cost:         23,
			Text:         "asd",
			Tags:         []string{"tag A", "tag B", "tag C"},
			ExtraCharges: []string{"charge A", "charge B", "charge C"},
			ReleaseDates: []time.Time{time.Now().AddDate(0, 0, 1)},
		},
		Size: 10,
	}

	req, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	req.NewBlockAdvertisement(blockDtoForTest)
	if BlockAdvertisementsError != nil {
		t.Fatal(BlockAdvertisementsError)
	}
}

func TestNewLineAdvertisement(t *testing.T) {
	LineAdvertisementsError = nil
	lineDtoForTest := datatransferobjects.LineAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           1,
			OrderID:      1,
			ReleaseCount: 1,
			Cost:         23,
			Text:         "asd",
			Tags:         []string{"tag A", "tag B", "tag C"},
			ExtraCharges: []string{"charge A", "charge B", "charge C"},
			ReleaseDates: []time.Time{time.Now().AddDate(0, 0, 1)},
		},
	}

	req, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	req.NewLineAdvertisement(lineDtoForTest)
	if LineAdvertisementsError != nil {
		t.Fatal(LineAdvertisementsError)
	}
}
