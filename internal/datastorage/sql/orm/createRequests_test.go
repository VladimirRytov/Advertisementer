package orm

import (
	"context"
	"testing"
	"time"

	"gorm.io/gorm"
)

var (
	client string
)

func TestCreateClient(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Error(err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		client, err = database.NewClient(contex, &clientDtoForTest)
		if err != nil {
			t.Error(err)
		}
		database.Close()
	}

}

func TestCreateOrder(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Error(err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		_, err = database.NewAdvertisementsOrder(contex, &orderDtoForTest)
		if err != nil {
			t.Error(err, v)
		}
		database.Close()
	}
}

func TestCreateTag(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Error(err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		for _, v := range tagsDtoForTest {
			tagName, err := database.NewTag(contex, &v)
			if err != nil && v.TagName != tagName {
				t.Error(err)
			}
			_, err = database.NewTag(contex, &v)
			if err == nil || err == gorm.ErrDuplicatedKey {
				t.Fatalf("want error")
			}
		}
		database.Close()
	}

}
func TestCreateExtraCharge(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Error(err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		for _, v := range extraChargesDtoForTest {
			chargeName, err := database.NewExtraCharge(contex, &v)
			if err != nil && v.ChargeName != chargeName {
				t.Error(err)
			}
			_, err = database.NewExtraCharge(contex, &v)
			if err == nil || err == gorm.ErrDuplicatedKey {
				t.Fatalf("want error")
			}
		}
		database.Close()
	}
}

func TestCreateCostRate(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Error(err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		_, err = database.NewCostRate(contex, &costRateDto)
		if err != nil {
			t.Fatal(err)
		}
		database.Close()
	}
}

func TestCreateBlockAdvertisement(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Error(err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		blockId, err := database.NewBlockAdvertisement(contex, &blockDtoForTest)
		if err != nil && blockId != blockDtoForTest.ID {
			t.Error(err)
		}
		database.Close()
	}
}
func TestCreateLineAdvertisement(t *testing.T) {
	CreateLogger()
	for _, v := range databases {
		database, err := connectToDBForTests(v)
		if err != nil {
			t.Error(err)
		}
		contex, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		_, err = database.NewLineAdvertisement(contex, &lineDtoForTest)
		if err != nil {
			t.Error(err)
		}
		database.Close()
	}
}
