package server

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestAllClients(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		clients, err := database.AllClients(contex)
		if err != nil {
			t.Fatal(err)
		}
		if len(clients) < 1 {
			t.Fatalf("want 1 or more clients, but got %d clients\n%v", len(clients), clients)
		}
		database.Close()
	}
}

func TestAllOrders(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		orders, err := database.AllOrders(contex)
		if err != nil {
			t.Fatal(err)
		}
		if len(orders) < 1 {
			t.Fatalf("want 1 or more orders, but got %d orders\n%v", len(orders), orders)
		}
		database.Close()
	}
}

func TestOrdersByClientName(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		wantClientID := "a"
		result, err := database.OrdersByClientName(contex, wantClientID)
		if err != nil {
			t.Fatalf(err.Error())
		}

		for _, v := range result {
			if wantClientID != v.ClientName {
				t.Fatalf("want id = %v,got order = %v", wantClientID, v)
			}
		}
		database.Close()
	}
}
func TestAllLineAdvertisements(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		lineAdv, err := database.AllLineAdvertisements(contex)
		if err != nil {
			t.Fatal(err)
		}
		if len(lineAdv) < 1 {
			t.Fatalf("want 1 or more LineAdvertisements, but got %d LineAdvertisements\n%v", len(lineAdv), lineAdv)
		}
		for _, v := range lineAdv {
			if len(v.Tags) < 1 && len(v.ExtraCharges) < 1 && len(v.ReleaseDates) < 1 {
				t.Fatalf("want 1 or more tags,extraCharges and releaseDates, but got %d tags: %v, %d extraCharges: %v, %d releaseDates: %v",
					len(v.Tags), v.Tags, len(v.ExtraCharges), v.ExtraCharges, len(v.ReleaseDates), v.ReleaseDates)
			}
		}
		database.Close()
	}
}

func TestLineAdvertisementsByOrderID(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		wantOrderID := 1
		result, err := database.LineAdvertisementsByOrderID(contex, wantOrderID)
		if err != nil {
			t.Fatalf(err.Error())
		}
		for _, v := range result {
			if wantOrderID != v.OrderID {
				t.Fatalf("want id = %v,got order = %v", wantOrderID, v)
			}
		}
		for _, v := range result {
			if len(v.Tags) < 1 && len(v.ExtraCharges) < 1 && len(v.ReleaseDates) < 1 {
				t.Fatalf("want 1 or more tags,extraCharges and releaseDates, but got %d tags: %v, %d extraCharges: %v, %d releaseDates: %v",
					len(v.Tags), v.Tags, len(v.ExtraCharges), v.ExtraCharges, len(v.ReleaseDates), v.ReleaseDates)
			}
		}
		database.Close()
	}
}

func TestAllBlockAdvertisements(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		blockAdv, err := database.AllBlockAdvertisements(contex)
		if err != nil {
			t.Fatal(err)
		}
		if len(blockAdv) < 1 {
			t.Fatalf("want 1 or more LineAdvertisements, but got %d LineAdvertisements\n%v", len(blockAdv), blockAdv)
		}
		for _, v := range blockAdv {
			if len(v.Tags) < 1 && len(v.ExtraCharges) < 1 && len(v.ReleaseDates) < 1 {
				t.Fatalf("want 1 or more tags,extraCharges and releaseDates, but got %d tags: %v, %d extraCharges: %v, %d releaseDates: %v",
					len(v.Tags), v.Tags, len(v.ExtraCharges), v.ExtraCharges, len(v.ReleaseDates), v.ReleaseDates)
			}
		}
		database.Close()
	}
}

func TestBlockAdvertisementsByOrderID(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		wantOrderID := 1
		result, err := database.BlockAdvertisementsByOrderID(contex, wantOrderID)
		if err != nil {
			t.Fatalf(err.Error())
		}
		for _, v := range result {
			if wantOrderID != v.OrderID {
				t.Fatalf("want id = %v,got order = %v", wantOrderID, v)
			}
		}
		for _, v := range result {
			if len(v.Tags) < 1 && len(v.ExtraCharges) < 1 && len(v.ReleaseDates) < 1 {
				t.Fatalf("want 1 or more tags,extraCharges and releaseDates, but got %d tags: %v, %d extraCharges: %v, %d releaseDates: %v",
					len(v.Tags), v.Tags, len(v.ExtraCharges), v.ExtraCharges, len(v.ReleaseDates), v.ReleaseDates)
			}
		}
		database.Close()
	}
}

func TestBlockAdvertisementBetweenReleaseDates(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		result, err := database.BlockAdvertisementBetweenReleaseDates(
			contex, time.Date(2021, 11, 11, 0, 0, 0, 0, time.UTC),
			time.Now().AddDate(2, 0, 0))
		if err != nil || len(result) < 1 {
			t.Fatalf("error: %v\ngot result: %v", err, result)
		}
		for _, v := range result {
			if len(v.Tags) < 1 && len(v.ExtraCharges) < 1 && len(v.ReleaseDates) < 1 {
				t.Fatalf("want 1 or more tags,extraCharges and releaseDates, but got %d tags: %v, %d extraCharges: %v, %d releaseDates: %v",
					len(v.Tags), v.Tags, len(v.ExtraCharges), v.ExtraCharges, len(v.ReleaseDates), v.ReleaseDates)
			}
		}

		result, err = database.BlockAdvertisementBetweenReleaseDates(
			contex, time.Now().AddDate(2, 0, 0), time.Now().AddDate(5, 0, 0))
		if err != nil || len(result) != 0 {
			t.Fatalf("error: %v\ngot result: %v", err, result)
		}
		database.Close()
	}
}

func TestBlockAdvertisementFromReleaseDates(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		result, err := database.BlockAdvertisementFromReleaseDates(
			contex, time.Date(2021, 11, 11, 0, 0, 0, 0, time.UTC))
		if err != nil || len(result) < 1 {
			t.Fatalf("error: %v\ngot result: %v", err, result)
		}
		for _, v := range result {
			if len(v.Tags) < 1 && len(v.ExtraCharges) < 1 && len(v.ReleaseDates) < 1 {
				t.Fatalf("want 1 or more tags,extraCharges and releaseDates, but got %d tags: %v, %d extraCharges: %v, %d releaseDates: %v",
					len(v.Tags), v.Tags, len(v.ExtraCharges), v.ExtraCharges, len(v.ReleaseDates), v.ReleaseDates)
			}
		}

		result, err = database.BlockAdvertisementFromReleaseDates(
			contex, time.Now().AddDate(2, 0, 0))
		if err != nil || len(result) != 0 {
			t.Fatalf("error: %v\ngot result: %v", err, result)
		}
		database.Close()
	}
}
func TestLineAdvertisementBetweenReleaseDates(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		result, err := database.LineAdvertisementBetweenReleaseDates(
			contex, time.Date(2021, 11, 11, 0, 0, 0, 0, time.UTC),
			time.Now().AddDate(2, 0, 0))
		if err != nil || len(result) < 1 {
			t.Fatalf("error: %v\ngot result: %v", err, result)
		}

		for _, v := range result {
			if len(v.Tags) < 1 && len(v.ExtraCharges) < 1 && len(v.ReleaseDates) < 1 {
				t.Fatalf("want 1 or more tags,extraCharges and releaseDates, but got %d tags: %v, %d extraCharges: %v, %d releaseDates: %v",
					len(v.Tags), v.Tags, len(v.ExtraCharges), v.ExtraCharges, len(v.ReleaseDates), v.ReleaseDates)
			}
		}

		result, err = database.LineAdvertisementBetweenReleaseDates(
			contex, time.Now().AddDate(2, 0, 0), time.Now().AddDate(5, 0, 0))
		if err != nil || len(result) != 0 {
			t.Fatalf("error: %v\ngot result: %v", err, result)
		}
		database.Close()
	}
}

func TestLineAdvertisementFromReleaseDates(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		result, err := database.LineAdvertisementFromReleaseDates(
			contex, time.Date(2021, 11, 11, 0, 0, 0, 0, time.UTC))
		if err != nil || len(result) < 1 {
			t.Fatalf("error: %v\ngot result: %v", err, result)
		}

		for _, v := range result {
			if len(v.Tags) < 1 && len(v.ExtraCharges) < 1 && len(v.ReleaseDates) < 1 {
				t.Fatalf("want 1 or more tags,extraCharges and releaseDates, but got %d tags: %v, %d extraCharges: %v, %d releaseDates: %v",
					len(v.Tags), v.Tags, len(v.ExtraCharges), v.ExtraCharges, len(v.ReleaseDates), v.ReleaseDates)
			}
		}

		result, err = database.LineAdvertisementFromReleaseDates(
			contex, time.Now().AddDate(2, 0, 0))
		if err != nil || len(result) != 0 {
			t.Fatalf("error: %v\ngot result: %v", err, result)
		}
		database.Close()
	}
}

func TestAllTags(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		result, err := database.AllTags(contex)
		if err != nil || len(result) < 1 {
			t.Fatalf("error: %v\ngot result: %v", err, result)
		}
		if !reflect.DeepEqual(result, tagsDtoForTest) {
			t.Fatalf("want tags = %v,\ngot tags = %v", tagsDtoForTest, result)
		}
		database.Close()
	}
}

func TestAllExtraChargess(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		result, err := database.AllExtraCharges(contex)
		if err != nil || len(result) < 1 {
			t.Fatalf("error: %v\ngot result: %v", err, result)
		}
		if !reflect.DeepEqual(result, extraChargesDtoForTest) {
			t.Fatalf("want extraCharges = %v,\ngot extraCharges = %v", extraChargesDtoForTest, result)
		}
		database.Close()
	}
}

func TestAllCostRates(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		_, err = database.AllCostRates(contex)
		if err != nil {
			t.Fatal(err)
		}
		database.Close()
	}
}
