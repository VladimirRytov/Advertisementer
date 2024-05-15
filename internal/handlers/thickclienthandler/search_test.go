package thickclienthandler

import (
	"testing"
	"time"
)

func TestAllClients(t *testing.T) {
	Clients = nil
	ClientsError = nil

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	ac.AllClients()
	if ClientsError != nil {
		t.Fatal(ClientsError)
	}
	if len(Clients) == 0 {
		t.Fatal("clients length = 0")
	}
}

func TestOrdersByClientID(t *testing.T) {
	Orders = nil
	OrdersError = nil

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	ac.OrdersByClientName(Clients[0].Name)
	if OrdersError != nil {
		t.Fatal(OrdersError)
	}

	if len(Orders) == 0 {
		t.Fatal("orders length = 0")
	}
}

func TestBlockAdvertisementsByOrderID(t *testing.T) {
	BlockAdvertisements = nil
	BlockAdvertisementsError = nil

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	ac.BlockAdvertisementsByOrderID(Orders[0].ID)
	if BlockAdvertisementsError != nil {
		t.Fatal(BlockAdvertisementsError)
	}

	if len(BlockAdvertisements) == 0 {
		t.Fatal("blockAdvertisements length = 0")
	}
}

func TestBlockAdvertisementsBetweenReleaseDates(t *testing.T) {
	BlockAdvertisements = nil
	BlockAdvertisementsError = nil

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	ac.BlockAdvertisementsBetweenReleaseDates(time.Date(2000, 12, 12, 12, 0, 0, 0, time.UTC), time.Now().AddDate(22, 2, 2))
	if BlockAdvertisementsError != nil {
		t.Fatal(BlockAdvertisementsError)
	}

	if len(BlockAdvertisements) == 0 {
		t.Fatal("blockAdvertisements length = 0")
	}
}

func TestBlockAdvertisementsActualReleaseDate(t *testing.T) {
	BlockAdvertisements = nil
	BlockAdvertisementsError = nil

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	ac.BlockAdvertisementsActualReleaseDate()
	if BlockAdvertisementsError != nil {
		t.Fatal(BlockAdvertisementsError)
	}

	if len(BlockAdvertisements) == 0 {
		t.Fatal("blockAdvertisements length = 0")
	}
}

func TestBlockAdvertisementsFromDate(t *testing.T) {
	BlockAdvertisements = nil
	BlockAdvertisementsError = nil

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	ac.BlockAdvertisementsFromReleaseDate(time.Date(2000, 12, 12, 12, 0, 0, 0, time.UTC))
	if BlockAdvertisementsError != nil {
		t.Fatal(BlockAdvertisementsError)
	}

	if len(BlockAdvertisements) == 0 {
		t.Fatal("blockAdvertisements length = 0")
	}
}

func TestLineAdvertisementsByOrderID(t *testing.T) {
	LineAdvertisements = nil
	LineAdvertisementsError = nil

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	ac.LineAdvertisementsByOrderID(Orders[0].ID)
	if LineAdvertisementsError != nil {
		t.Fatal(LineAdvertisementsError)
	}

	if len(LineAdvertisements) == 0 {
		t.Fatal("lineAdvertisements length = 0")
	}
}

func TestLineAdvertisementsBetweenReleaseDates(t *testing.T) {
	LineAdvertisements = nil
	LineAdvertisementsError = nil

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	ac.LineAdvertisementsBetweenReleaseDates(time.Date(2000, 12, 12, 12, 0, 0, 0, time.UTC), time.Now().AddDate(2, 2, 2))
	if LineAdvertisementsError != nil {
		t.Fatal(LineAdvertisementsError)
	}

	if len(LineAdvertisements) == 0 {
		t.Fatal("lineAdvertisements length = 0")
	}
}

func TestLineAdvertisementsActualReleaseDate(t *testing.T) {
	LineAdvertisements = nil
	LineAdvertisementsError = nil

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	ac.LineAdvertisementsActualReleaseDate()
	if LineAdvertisementsError != nil {
		t.Fatal(LineAdvertisementsError)
	}

	if len(LineAdvertisements) == 0 {
		t.Fatal("lineAdvertisements length = 0")
	}
}

func TestLineAdvertisementsFromDate(t *testing.T) {
	LineAdvertisements = nil
	LineAdvertisementsError = nil

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}

	ac.LineAdvertisementsFromReleaseDate(time.Date(2000, 12, 12, 12, 0, 0, 0, time.UTC))
	if LineAdvertisementsError != nil {
		t.Fatal(LineAdvertisementsError)
	}

	if len(LineAdvertisements) == 0 {
		t.Fatal("lineAdvertisements length = 0")
	}
}

func TestAllTags(t *testing.T) {
	Tags = nil
	TagsError = nil

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}
	ac.AllTags()
	if TagsError != nil {
		t.Fatal(TagsError)
	}
	if len(Tags) == 0 {
		t.Fatal("tags length = 0")
	}
}

func TestAllExtraCharges(t *testing.T) {
	ExtraCharges = nil
	ExtraChargesError = nil

	ac, err := CreateControllerForTests()
	if err != nil {
		t.Fatal(err)
	}
	ac.AllExtraCharges()
	if ExtraChargesError != nil {
		t.Fatal(err)
	}
	if len(ExtraCharges) == 0 {
		t.Fatal("extraCharges length = 0")
	}
}
