package orm

import (
	"context"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestRemoveTagByName(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		testTags := tagsDtoForTest
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		for _, tag := range testTags {
			err = database.RemoveTagByName(contex, tag.TagName)
			if err != nil {
				t.Fatal(err, v)
			}
			_, err = database.TagByName(contex, tag.TagName)
			if err != gorm.ErrRecordNotFound {
				t.Fatal(err, v)
			}
		}
		database.Close()
	}
}

func TestRemoveExtraChargeByName(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		testExtraCharges := extraChargesDtoForTest
		for _, charge := range testExtraCharges {
			err = database.RemoveExtraChargeByName(contex, charge.ChargeName)
			if err != nil {
				t.Fatal(err, v)
			}
			_, err = database.ExtraChargeByName(contex, charge.ChargeName)
			if err != gorm.ErrRecordNotFound {
				t.Fatal(err, v)
			}
		}
		database.Close()
	}
}

func TestRemoveLineAdvertisementByID(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		err = database.RemoveLineAdvertisementByID(contex, 1)
		if err != nil {
			t.Fatal(err, v)
		}
		_, err = database.LineAdvertisementByID(contex, 1)
		if err != gorm.ErrRecordNotFound {
			t.Fatal(err, v)
		}
		database.Close()
	}
}

func TestRemoveBlockAdvertisementByID(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		err = database.RemoveBlockAdvertisementByID(contex, 1)
		if err != nil {
			t.Fatal(err, v)
		}
		_, err = database.BlockAdvertisementByID(contex, 1)
		if err != gorm.ErrRecordNotFound {
			t.Fatal(err, v)
		}
		database.Close()
	}
}

func TestRemoveOrderByID(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		orderTest := orderDtoForTest
		err = database.RemoveOrderByID(contex, orderTest.ID)
		if err != nil {
			t.Fatal(err, v)
		}
		_, err = database.OrderByID(contex, orderTest.ID)
		if err != gorm.ErrRecordNotFound {
			t.Fatal(err, v)
		}
		database.Close()
	}
}

func TestRemoveClientByName(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		clientTest := clientDtoForTest
		err = database.RemoveClientByName(contex, clientTest.Name)
		if err != nil {
			t.Fatal(err, v)
		}
		_, err = database.ClientByName(contex, clientTest.Name)
		if err != gorm.ErrRecordNotFound {
			t.Fatal(err, v)
		}
		database.Close()
	}
}

func TestRemoveCostRateByName(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Fatalf("error: coud`n connect to database. got error: %v", err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		err = database.RemoveCostRateByName(contex, costRateDto.Name)
		if err != nil {
			t.Fatal(err, v)
		}
		_, err = database.CostRateByName(contex, costRateDto.Name)
		if err != gorm.ErrRecordNotFound {
			t.Fatal(err, v)
		}
		database.Close()
	}
}
