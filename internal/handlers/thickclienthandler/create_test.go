package thickclienthandler

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

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	ac.NewClient(client)
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

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	ac.NewAdvertisementsOrder(order)
	if OrdersError != nil {
		t.Fatal(OrdersError)
	}
}

func TestNewTag(t *testing.T) {
	TagsError = nil
	tagDTO := datatransferobjects.TagDTO{TagName: "tag A", TagCost: 12}

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	ac.NewTag(tagDTO)
	if TagsError != nil {
		t.Fatal(TagsError)
	}
}

func TestNewExtraCharge(t *testing.T) {
	ExtraChargesError = nil
	extraChargeDTO := datatransferobjects.ExtraChargeDTO{ChargeName: "charge A", Multiplier: 2}

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	ac.NewExtraCharge(extraChargeDTO)
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

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	ac.NewBlockAdvertisement(blockDtoForTest)
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

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	ac.NewLineAdvertisement(lineDtoForTest)
	if LineAdvertisementsError != nil {
		t.Fatal(LineAdvertisementsError)
	}
}
