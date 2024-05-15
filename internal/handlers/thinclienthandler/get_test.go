package thinclienthandler

import (
	"testing"
)

func TestGetClient(t *testing.T) {
	ClientsError = nil
	Clients = nil

	req, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}
	clientName := "Вася"
	req.ClientByName(clientName)
	if ClientsError != nil {
		t.Fatal(ClientsError)
	}
	if Clients[0].Name != clientName {
		t.Fatalf("want Client name %s, got %s", clientName, Clients[0].Name)
	}
}

func TestGetOrder(t *testing.T) {
	Orders = nil
	OrdersError = nil

	req, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	orderId := 1
	req.OrderByID(orderId)
	if OrdersError != nil {
		t.Fatal(OrdersError)
	}

	if Orders[0].ID != orderId {
		t.Fatalf("want OrderID %d, got %d", orderId, Orders[0].ID)
	}
}

func TestGetBlockAdvertisement(t *testing.T) {
	BlockAdvertisements = nil
	BlockAdvertisementsError = nil

	req, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	blockAdvID := 1
	req.BlockAdvertisementByID(blockAdvID)
	if BlockAdvertisementsError != nil {
		t.Fatal(BlockAdvertisementsError)
	}

	if blockAdvID != BlockAdvertisements[0].ID {
		t.Fatalf("want BlockAdvertisementID %d, got %d", blockAdvID, BlockAdvertisements[0].ID)
	}
}

func TestGetLineAdvertisement(t *testing.T) {
	LineAdvertisements = nil
	LineAdvertisementsError = nil

	req, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	lineAdvID := 1
	req.LineAdvertisementByID(lineAdvID)
	if LineAdvertisementsError != nil {
		t.Fatal(LineAdvertisementsError)
	}

	if lineAdvID != LineAdvertisements[0].ID {
		t.Fatalf("want LineAdvertisementID %d, got %d", lineAdvID, BlockAdvertisements[0].ID)
	}
}

func TestGetTag(t *testing.T) {
	Tags = nil
	TagsError = nil

	req, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}
	tagName := "tag A"
	req.TagByName(tagName)
	if TagsError != nil {
		t.Fatal(TagsError)
	}

	if tagName != Tags[0].TagName {
		t.Fatalf("want tag name = %s, got %s", tagName, Tags[0].TagName)
	}
}

func TestGetExtraCharge(t *testing.T) {
	ExtraCharges = nil
	ExtraChargesError = nil

	req, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}
	chargeName := "charge A"
	req.ExtraChargeByName(chargeName)
	if ExtraChargesError != nil {
		t.Fatal(err)
	}

	if chargeName != ExtraCharges[0].ChargeName {
		t.Fatalf("want charge name = %s, got %s", chargeName, ExtraCharges[0].ChargeName)
	}
}
