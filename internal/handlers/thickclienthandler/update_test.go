package thickclienthandler

import (
	"reflect"
	"testing"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
)

func TestUpdateClient(t *testing.T) {
	Clients = nil
	ClientsError = nil
	client := datatransferobjects.ClientDTO{
		Name:                  "Вася",
		Phones:                "8-800-555-35-35",
		Email:                 "mail@mail.com",
		AdditionalInformation: "asdasd",
	}
	newClient := datatransferobjects.ClientDTO{
		Name:                  "Вася",
		Phones:                "8-800-555-35-35",
		Email:                 "mail@mail.com",
		AdditionalInformation: "zzxczx",
	}

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	ac.UpdateClient(newClient)
	if ClientsError != nil {
		t.Fatal(ClientsError)
	}

	if reflect.DeepEqual(client, Clients[0]) {
		t.Fatalf("old and new client are equal")
	}
}

func TestUpdateOrder(t *testing.T) {
	Orders = nil
	OrdersError = nil
	order := datatransferobjects.OrderDTO{
		ID:            1,
		ClientName:    "Вася",
		Cost:          10,
		PaymentType:   "asd",
		CreatedDate:   time.Now(),
		PaymentStatus: true,
	}

	newOrder := datatransferobjects.OrderDTO{
		ID:            1,
		ClientName:    "Вася",
		Cost:          10,
		PaymentType:   "иит",
		CreatedDate:   time.Now(),
		PaymentStatus: false,
	}

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	ac.UpdateOrder(newOrder)
	if OrdersError != nil {
		t.Fatal(OrdersError)
	}
	if reflect.DeepEqual(order, Orders[0]) {
		t.Fatalf("old and new orders are equal")
	}
}

func TestUpdateBlockAdvertisement(t *testing.T) {
	BlockAdvertisements = nil
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
			ReleaseDates: []time.Time{time.Now()},
		},
		Size: 10,
	}
	newBlockDtoForTest := datatransferobjects.BlockAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           1,
			OrderID:      1,
			ReleaseCount: 1,
			Cost:         23,
			Text:         "asd",
			Tags:         []string{"tag A", "tag B", "tag C"},
			ExtraCharges: []string{"charge A", "charge B", "charge C"},
			ReleaseDates: []time.Time{time.Now()},
		},
		Size: 210,
	}

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}
	ac.UpdateBlockAdvertisement(newBlockDtoForTest)
	if BlockAdvertisementsError != nil {
		t.Fatal(BlockAdvertisementsError)
	}
	if reflect.DeepEqual(blockDtoForTest, BlockAdvertisements[0]) {
		t.Fatalf("old and new blockAdv are equal")
	}
}

func TestUpdateLienAdvertisement(t *testing.T) {
	LineAdvertisements = nil
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
			ReleaseDates: []time.Time{time.Now()},
		},
	}
	newLineDtoForTest := datatransferobjects.LineAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           1,
			OrderID:      1,
			ReleaseCount: 1,
			Cost:         55,
			Text:         "cxzx",
			Tags:         []string{"tag A", "tag B", "tag C"},
			ExtraCharges: []string{"charge A", "charge B", "charge C"},
			ReleaseDates: []time.Time{time.Now()},
		},
	}

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	ac.UpdateLineAdvertisement(newLineDtoForTest)
	if LineAdvertisementsError != nil {
		t.Fatal(LineAdvertisementsError)
	}
	if reflect.DeepEqual(lineDtoForTest, LineAdvertisements[0]) {
		t.Fatalf("old and new lineAdv are equal")
	}
}

func TestUpdateTag(t *testing.T) {
	Tags = nil
	TagsError = nil
	tagDTO := datatransferobjects.TagDTO{TagName: "tag A", TagCost: 12}
	newTagDTO := datatransferobjects.TagDTO{TagName: "tag A", TagCost: 22}

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	ac.UpdateTag(newTagDTO)
	if TagsError != nil {
		t.Fatal(TagsError)
	}

	if reflect.DeepEqual(tagDTO, Tags[0]) {
		t.Fatalf("old and new tag are equal")
	}
}

func TestUpdateExtraCharge(t *testing.T) {
	ExtraCharges = nil
	ExtraChargesError = nil
	extraChargeDTO := datatransferobjects.ExtraChargeDTO{ChargeName: "charge A", Multiplier: 2}
	newExtraChargeDTO := datatransferobjects.ExtraChargeDTO{ChargeName: "charge A", Multiplier: 4}

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	ac.UpdateExtraCharge(newExtraChargeDTO)
	if ExtraChargesError != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(extraChargeDTO, ExtraCharges[0]) {
		t.Fatalf("old and new extraCharge are equal")
	}
}
