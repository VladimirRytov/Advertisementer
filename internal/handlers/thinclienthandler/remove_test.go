package thinclienthandler

import (
	"testing"
)

func TestRemoveClient(t *testing.T) {
	ClientsError = nil
	clientName := "Вася"

	req, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	req.RemoveClientByName(clientName)
	if ClientsError != nil {
		t.Fatal(ClientsError)
	}
}

func TestRemoveOrder(t *testing.T) {
	OrdersError = nil
	orderId := 1

	req, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	req.RemoveOrderByID(orderId)
	if OrdersError != nil {
		t.Fatal(OrdersError)
	}
}

func TestRemoveBlockAdvertisement(t *testing.T) {
	BlockAdvertisementsError = nil

	req, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	blockAdvID := 1

	req.RemoveBlockAdvertisementByID(blockAdvID)
	if BlockAdvertisementsError != nil {
		t.Fatal(BlockAdvertisementsError)
	}
}

func TestRemoveLineAdvertisement(t *testing.T) {
	LineAdvertisementsError = nil

	req, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	lineAdvID := 1
	req.RemoveLineAdvertisementByID(lineAdvID)
	if LineAdvertisementsError != nil {
		t.Fatal(LineAdvertisementsError)
	}
}

func TestRemoveTag(t *testing.T) {
	TagsError = nil

	req, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	tagName := "tag A"
	req.RemoveTagByName(tagName)
	if TagsError != nil {
		t.Fatal(TagsError)
	}
}

func TestRemoveExtraCharge(t *testing.T) {
	ExtraChargesError = nil

	req, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	chargeName := "charge A"
	req.RemoveExtraChargeByName(chargeName)
	if ExtraChargesError != nil {
		t.Fatal(err)
	}
}
