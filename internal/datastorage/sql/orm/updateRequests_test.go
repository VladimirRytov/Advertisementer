package orm

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestUpdateClient(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		newClient := clientDtoForTest
		newClient.Email = "asdasdaszxc@mail.xz"
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		result := database.UpdateClient(contex, &newClient)
		if result != nil {
			t.Fatalf("error %v", result)
		}
		gotCLient, err := database.ClientByName(contex, newClient.Name)
		if err != nil {
			t.Fatalf("error %v", err)
		}
		if gotCLient.Email != newClient.Email {
			t.Fatalf("want client name: %s, got client name: %s", newClient.Email, gotCLient.Email)
		}
		database.Close()
	}
}
func TestUpdateOrder(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		newOrder := orderDtoForTest
		newOrder.Cost = 9000
		newOrder.PaymentStatus = false
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		result := database.UpdateOrder(contex, &newOrder)
		if result != nil {
			t.Fatalf("error %v", result)
		}
		gotOrder, err := database.OrderByID(contex, newOrder.ID)
		if err != nil {
			t.Fatalf("error %v", err)
		}
		if gotOrder.Cost != newOrder.Cost {
			t.Fatalf("want order cost: %d, got %d", newOrder.Cost, gotOrder.Cost)
		}
		if gotOrder.PaymentStatus != newOrder.PaymentStatus {
			t.Fatalf("want order payment status: %v, got %v", newOrder.PaymentStatus, gotOrder.PaymentStatus)
		}
		database.Close()
	}
}

func TestUpdateAdvertisementBlock(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		newBlockAdv := blockDtoForTest
		newBlockAdv.Tags = []string{"Tag A", "Tag B"}
		newBlockAdv.ExtraCharges = []string{"Charge A", "Charge B"}
		newBlockAdv.Cost = 999999
		newBlockAdv.ReleaseDates = []time.Time{time.Date(2023, 12, 12, 0, 0, 0, 0, time.UTC)}
		result := database.UpdateBlockAdvertisement(contex, &newBlockAdv)
		if result != nil {
			t.Fatalf("error %v", result)
		}
		gotBlock, err := database.BlockAdvertisementByID(contex, newBlockAdv.ID)
		if err != nil {
			t.Fatal(err)
		}
		if newBlockAdv.Cost != gotBlock.Cost && !reflect.DeepEqual(newBlockAdv.ExtraCharges, gotBlock.ExtraCharges) &&
			!reflect.DeepEqual(newBlockAdv.Tags, gotBlock.Tags) {
			t.Fatalf("want %v,got %v,database %s", newBlockAdv, gotBlock, v)
		}
		for i, v := range gotBlock.ReleaseDates {
			if !v.Equal(newBlockAdv.ReleaseDates[i]) {
				t.Fatalf("want %v,got %v,database %s", newBlockAdv, gotBlock, v)
			}
		}
		database.Close()
	}
}
func TestUpdateAdvertisementLine(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		newAdvLine := lineDtoForTest
		newAdvLine.Tags = []string{"Tag A", "Tag B"}
		newAdvLine.ExtraCharges = []string{"Charge A", "Charge B"}
		newAdvLine.Cost = 999999
		newAdvLine.ReleaseDates = []time.Time{time.Date(2023, 12, 12, 0, 0, 0, 0, time.UTC)}
		result := database.UpdateLineAdvertisement(contex, &newAdvLine)
		if result != nil {
			t.Fatalf("error %v", result)
		}
		gotLine, err := database.LineAdvertisementByID(contex, newAdvLine.ID)
		if err != nil {
			t.Fatal(err)
		}
		if newAdvLine.Cost != gotLine.Cost && !reflect.DeepEqual(newAdvLine.ExtraCharges, gotLine.ExtraCharges) &&
			!reflect.DeepEqual(newAdvLine.Tags, gotLine.Tags) {
			t.Fatalf("want %v,got %v,database %s", newAdvLine, gotLine, v)
		}
		for i, v := range gotLine.ReleaseDates {
			if !v.Equal(newAdvLine.ReleaseDates[i]) {
				t.Fatalf("want %v,got %v,database %s", newAdvLine, gotLine, v)
			}
		}
		database.Close()
	}
}
func TestUpdateTag(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		newTag := tagsDtoForTest[0]
		newTag.TagCost = 9999999
		result := database.UpdateTag(contex, &newTag)
		if result != nil {
			t.Fatalf("error %v", result)
		}

		gotTag, err := database.TagByName(contex, newTag.TagName)
		if err != nil {
			t.Fatalf("error %v", err)
		}

		if gotTag.TagCost != newTag.TagCost {
			t.Fatalf("want %v,got %v", gotTag, newTag)
		}
		database.Close()
	}
}
func TestUpdateExtraCharge(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		newExtraCharge := extraChargesDtoForTest[0]
		newExtraCharge.Multiplier = 9999999
		result := database.UpdateExtraCharge(contex, &newExtraCharge)
		if result != nil {
			t.Fatalf("error %v", result)
		}

		gotExtraCharge, err := database.ExtraChargeByName(contex, newExtraCharge.ChargeName)
		if err != nil {
			t.Fatalf("error %v", err)
		}

		if gotExtraCharge.Multiplier != newExtraCharge.Multiplier {
			t.Fatalf("want %v,got %v", gotExtraCharge, newExtraCharge)
		}
		database.Close()
	}
}

func TestUpdateCostRate(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		costRate := costRateDto
		costRate.ForOneWordSymbol = 1000
		result := database.UpdateCostRate(contex, &costRate)
		if result != nil {
			t.Fatalf("error %v", result)
		}

		gotCostRate, err := database.CostRateByName(contex, costRate.Name)
		if err != nil {
			t.Fatalf("error %v", err)
		}

		if gotCostRate.ForOneWordSymbol != costRate.ForOneWordSymbol {
			t.Fatalf("want %v,got %v", costRate, gotCostRate)
		}
		database.Close()
	}
}
