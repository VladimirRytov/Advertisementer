package thickclienthandler

import (
	"testing"
)

func TestRemoveClient(t *testing.T) {
	ClientsError = nil
	clientName := "Вася"

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	ac.RemoveClientByName(clientName)
	if ClientsError != nil {
		t.Fatal(ClientsError)
	}
}

func TestRemoveOrder(t *testing.T) {
	OrdersError = nil
	orderId := 1

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	ac.RemoveOrderByID(orderId)
	if OrdersError != nil {
		t.Fatal(OrdersError)
	}
}

func TestRemoveBlockAdvertisement(t *testing.T) {
	BlockAdvertisementsError = nil

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	blockAdvID := 1

	ac.RemoveBlockAdvertisementByID(blockAdvID)
	if BlockAdvertisementsError != nil {
		t.Fatal(BlockAdvertisementsError)
	}
}

func TestRemoveLineAdvertisement(t *testing.T) {
	LineAdvertisementsError = nil

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	lineAdvID := 1
	ac.RemoveLineAdvertisementByID(lineAdvID)
	if LineAdvertisementsError != nil {
		t.Fatal(LineAdvertisementsError)
	}
}

func TestRemoveTag(t *testing.T) {
	TagsError = nil

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	tagName := "tag A"
	ac.RemoveTagByName(tagName)
	if TagsError != nil {
		t.Fatal(TagsError)
	}
}

func TestRemoveExtraCharge(t *testing.T) {
	ExtraChargesError = nil

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	chargeName := "charge A"
	ac.RemoveExtraChargeByName(chargeName)
	if ExtraChargesError != nil {
		t.Fatal(err)
	}
}
