package mapper

import (
	"errors"
	"testing"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/advertisements"
	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
)

func TestDtoToClient(t *testing.T) {
	CreateLogger()
	dto := datatransferobjects.ClientDTO{
		Name:                  "Вася",
		Phones:                "8-800-555-35-35",
		Email:                 "mail.email@mail.com",
		AdditionalInformation: "asdasd",
	}
	got, err := DtoToClient(&dto)
	if err != nil {
		t.Fatalf("got error %v", err)
	}
	if dto.Name != got.Name() || dto.Email != got.Email() || dto.Phones != got.ContactNumber() || dto.AdditionalInformation != got.AdditionalInformation() {
		t.Error("error")
	}

	dto.Name = "a"
	_, err = DtoToClient(&dto)
	if err != nil && !errors.Is(err, advertisements.ErrClientName()) {
		t.Fatalf("want error %v, got %v", advertisements.ErrClientName(), err)
	}
	dto.Name = "asd"
	dto.Phones = "836  34534 768"
	_, err = DtoToClient(&dto)
	if err != nil && !errors.Is(err, advertisements.ErrPhonePattern()) {
		t.Fatalf("want error %v, got %v", advertisements.ErrPhonePattern(), err)
	}
	dto.Phones = ""
	dto.Email = "asdzxc@.asd"
	_, err = DtoToClient(&dto)
	if err != nil && !errors.Is(err, advertisements.ErrEmail()) {
		t.Fatalf("want error %v, got %v", advertisements.ErrEmail(), err)
	}
}

func TestDtoToOrder(t *testing.T) {
	CreateLogger()
	dto := datatransferobjects.OrderDTO{
		ID:            1,
		ClientName:    "Вася",
		Cost:          123,
		CreatedDate:   time.Now(),
		PaymentStatus: true,
	}
	got, err := DtoToOrder(&dto)
	if got.ClientId() != dto.ClientName || got.Cost() != dto.Cost || got.CreaatedDate() != dto.CreatedDate || err != nil {
		t.Error()
	}
	dto.ClientName = "a"
	_, err = DtoToOrder(&dto)
	if err != nil && !errors.Is(err, advertisements.ErrClientName()) {
		t.Fatalf("want error %v, got %v", advertisements.ErrClientName(), err)
	}
}

func TestDtoToBlockAdvertisement(t *testing.T) {
	CreateLogger()
	dto := datatransferobjects.BlockAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           1,
			OrderID:      1,
			ReleaseCount: 1,
			Cost:         12,
			Text:         "asd",
			Tags:         []string{"b", "c", "a"},
			ExtraCharges: []string{"b", "c", "a"},
			ReleaseDates: []time.Time{time.Now()},
		},
		Size: 10,
	}
	got, err := DtoToAdvertisementBlock(&dto)
	if dto.ID != got.Id() || dto.Cost != got.Cost() || dto.OrderID != got.OrderId() ||
		dto.Size != got.Size() || dto.Text != got.Comment() || err != nil {
		t.Fatalf("values not equal. error: %v", err)
	}
	for i := range dto.ExtraCharges {
		if dto.ExtraCharges[i] != got.ExtraCharge()[i] {
			t.Fatal("values not equal")
		}
	}
	for i := range dto.Tags {
		if dto.Tags[i] != got.Tags()[i] {
			t.Fatal("values not equal")
		}
	}
	dto.ID = 1
	dto.OrderID = 2
	dto.Size = 0
	_, err = DtoToAdvertisementBlock(&dto)
	if err != nil && !errors.Is(err, advertisements.ErrSize()) {
		t.Fatalf("want err %v, got %v", advertisements.ErrSize(), err)
	}
}
func TestDtoToLineAdvertisement(t *testing.T) {
	CreateLogger()
	dto := &datatransferobjects.LineAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           1,
			OrderID:      1,
			ReleaseCount: 1,
			Cost:         12,
			Text:         "asd",
			Tags:         []string{"b", "c", "a"},
			ExtraCharges: []string{"b", "c", "a"},
			ReleaseDates: []time.Time{time.Now()},
		},
	}
	got, err := DtoToAdvertisementLine(dto)
	if dto.ID != got.Id() || dto.Cost != got.Cost() || dto.OrderID != got.OrderId() ||
		dto.Text != got.Content() || err != nil {
		t.Fatalf("values not equal. error: %v", err)
	}
	for i := range dto.ExtraCharges {
		if dto.ExtraCharges[i] != got.ExtraCharge()[i] {
			t.Fatal("values not equal")
		}
	}
	for i := range dto.Tags {
		if dto.Tags[i] != got.Tags()[i] {
			t.Fatal("values not equal")
		}
	}
}

func TestDtoToTag(t *testing.T) {
	CreateLogger()
	dto := datatransferobjects.TagDTO{TagName: "asd", TagCost: 123}
	got, err := DtoToTag(&dto)
	if err != nil || dto.TagCost != got.Cost() || dto.TagName != got.Name() {
		t.Fatalf("values not equal or got error %v", err)
	}
	dto = datatransferobjects.TagDTO{TagName: "", TagCost: 123}
	_, err = DtoToTag(&dto)
	if err.Error() != "имя тэга должно быть заполнено" {
		t.Fatalf("want err, got %v", err)
	}

}

func TestDtoToExtraCharge(t *testing.T) {
	CreateLogger()
	dto := datatransferobjects.ExtraChargeDTO{ChargeName: "asd", Multiplier: 123}
	got, err := DtoToExtraCharge(&dto)
	if err != nil || dto.Multiplier != got.Multiplier() || dto.ChargeName != got.Name() {
		t.Fatalf("values not equal or got error %v", err)
	}
	dto = datatransferobjects.ExtraChargeDTO{ChargeName: "", Multiplier: 123}
	_, err = DtoToExtraCharge(&dto)
	if err.Error() != "имя наценки должно быть заполнено" {
		t.Fatalf("want err, got %v", err)
	}
}
