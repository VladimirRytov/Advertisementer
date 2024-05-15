package server

import (
	"reflect"
	"testing"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/encodedecoder"
)

func TestConvertClientToDTO(t *testing.T) {
	CreateLogger()
	ds := NewServerStorage(nil, encodedecoder.NewBase64Encoder())
	got := ds.convertClientToDTO(&clientModel)
	if got.Name != clientDtoForTest.Name || got.Email != clientDtoForTest.Email || got.Phones != clientDtoForTest.Phones ||
		got.AdditionalInformation != clientDtoForTest.AdditionalInformation {
		t.Fatalf("want %v,got %v", clientDtoForTest, got)
	}
}
func TestConvertClientToModel(t *testing.T) {
	CreateLogger()
	ds := NewServerStorage(nil, encodedecoder.NewBase64Encoder())
	got := ds.convertClientToModel(&clientDtoForTest)
	if got.Name != clientDtoForTest.Name || got.Email != clientDtoForTest.Email || got.Phones != clientDtoForTest.Phones {
		t.Fatalf("want %v,got %v", clientDtoForTest, got)
	}
}

func TestConvertOrderToDTO(t *testing.T) {
	CreateLogger()
	ds := NewServerStorage(nil, encodedecoder.NewBase64Encoder())
	got := ds.convertOrderToDTO(&orderModel)
	if got.ID != orderDtoForTest.ID || got.ClientName != orderDtoForTest.ClientName || orderDtoForTest.PaymentStatus != got.PaymentStatus ||
		got.Cost != orderDtoForTest.Cost || !got.CreatedDate.Equal(orderDtoForTest.CreatedDate) {

		t.Fatalf("want %v,got %v", orderDtoForTest, got)
	}
}

func TestConvertOrderToModel(t *testing.T) {
	CreateLogger()
	ds := NewServerStorage(nil, encodedecoder.NewBase64Encoder())
	got := ds.convertOrderToModel(&orderDtoForTest)
	if got.ClientName != orderModel.ClientName || orderModel.PaymentStatus != got.PaymentStatus ||
		got.Cost != orderModel.Cost || got.CreatedDate != orderModel.CreatedDate {

		t.Fatalf("want %v,got %v", orderDtoForTest, got)
	}
}

func TestConvertTagToDTO(t *testing.T) {
	CreateLogger()
	ds := NewServerStorage(nil, encodedecoder.NewBase64Encoder())
	got := make([]datatransferobjects.TagDTO, 0, len(tagsModel))
	for _, v := range tagsModel {
		got = append(got, ds.convertTagToDTO(&v))
	}
	for i := range got {
		if got[i].TagName != tagsDtoForTest[i].TagName || got[i].TagCost != tagsDtoForTest[i].TagCost {
			t.Fatalf("want %v,got %v", tagsDtoForTest, got)
		}
	}
}

func TestConvertTagsToModel(t *testing.T) {
	CreateLogger()
	ds := NewServerStorage(nil, encodedecoder.NewBase64Encoder())
	got := make([]TagFront, 0, len(tagsDtoForTest))
	for _, v := range tagsDtoForTest {
		got = append(got, ds.convertTagToModel(&v))
	}
	for i := range got {
		if got[i].TagName != tagsModel[i].TagName || got[i].TagCost != tagsModel[i].TagCost {
			t.Fatalf("want %v,got %v", tagsModel, got)
		}
	}
}

func TestConvertExtraChargeToDTO(t *testing.T) {
	CreateLogger()
	ds := NewServerStorage(nil, encodedecoder.NewBase64Encoder())
	got := make([]datatransferobjects.ExtraChargeDTO, 0, len(extraChargesModel))
	for _, v := range extraChargesModel {
		got = append(got, ds.convertExtraChargeToDTO(&v))
	}
	for i := range got {
		if got[i].ChargeName != extraChargesDtoForTest[i].ChargeName || got[i].Multiplier != extraChargesDtoForTest[i].Multiplier {
			t.Fatalf("want %v,got %v", clientDtoForTest, got)
		}
	}
}

func TestConvertExtraChargeToModel(t *testing.T) {
	CreateLogger()
	ds := NewServerStorage(nil, encodedecoder.NewBase64Encoder())
	got := make([]ExtraChargeFront, 0, len(extraChargesDtoForTest))
	for _, v := range extraChargesDtoForTest {
		got = append(got, ds.convertExtraChargeToModel(&v))
	}
	for i := range got {
		if got[i].ChargeName != extraChargesModel[i].ChargeName || got[i].Multiplier != extraChargesModel[i].Multiplier {
			t.Fatalf("want %v,got %v", extraChargesModel, got)
		}
	}
}
func TestConvertBlockAdvertisementToDTO(t *testing.T) {
	CreateLogger()
	ds := NewServerStorage(nil, encodedecoder.NewBase64Encoder())
	got := ds.convertBlockAdvertisementToDTO(&blockModel)
	if !reflect.DeepEqual(got, blockDtoForTest) {

		t.Fatalf("want %v,got %v", blockDtoForTest, got)
	}
}

func TestConvertBlockAdvertisementToModel(t *testing.T) {
	CreateLogger()
	ds := NewServerStorage(nil, encodedecoder.NewBase64Encoder())
	got := ds.convertBlockAdvertisementToModel(&blockDtoForTest)
	if got.ID != blockDtoForTest.ID || got.OrderID != blockDtoForTest.OrderID || got.Cost != blockDtoForTest.Cost ||
		got.FileName != blockDtoForTest.FileName || got.ReleaseCount != blockDtoForTest.ReleaseCount ||
		len(got.Tags) != len(blockDtoForTest.Tags) || len(got.ExtraCharges) != len(blockDtoForTest.ExtraCharges) {

		t.Fatalf("want %v,got %v", blockModel, got)
	}
	for i, v := range got.ReleaseDates {
		if !v.Equal(blockModel.ReleaseDates[i]) {
			t.Fatalf("want %v,got %v", blockModel, got)
		}
	}
}

func TestConvertLineAdvertisementToDTO(t *testing.T) {
	CreateLogger()
	ds := NewServerStorage(nil, encodedecoder.NewBase64Encoder())
	got := ds.convertLineAdvertisementToDTO(&lineModel)
	if got.ID != lineDtoForTest.ID || got.OrderID != lineDtoForTest.OrderID || got.Cost != lineDtoForTest.Cost ||
		got.ReleaseCount != lineDtoForTest.ReleaseCount ||
		len(got.Tags) != len(lineDtoForTest.Tags) || len(got.ExtraCharges) != len(lineDtoForTest.ExtraCharges) {
		t.Fatalf("want %v,got %v", lineDtoForTest, got)
	}
	for i, v := range got.ReleaseDates {
		if !v.Equal(lineDtoForTest.ReleaseDates[i]) {
			t.Fatalf("want %v,got %v", lineModel, got)
		}
	}
}

func TestConvertLineAdvertisementToModel(t *testing.T) {
	CreateLogger()
	ds := NewServerStorage(nil, encodedecoder.NewBase64Encoder())
	got := ds.convertLineAdvertisementToModel(&lineDtoForTest)
	if got.ID != lineDtoForTest.ID || got.OrderID != lineDtoForTest.OrderID || got.Cost != lineDtoForTest.Cost ||
		got.ReleaseCount != lineDtoForTest.ReleaseCount ||
		len(got.Tags) != len(lineDtoForTest.Tags) || len(got.ExtraCharges) != len(lineDtoForTest.ExtraCharges) {
		t.Fatalf("want %v,got %v", lineModel, got)
	}
	for i, v := range got.ReleaseDates {
		if !v.Equal(lineModel.ReleaseDates[i]) {
			t.Fatalf("want %v,got %v", lineModel, got)
		}
	}
}

func TestCostRateToModel(t *testing.T) {
	CreateLogger()
	ds := NewServerStorage(nil, encodedecoder.NewBase64Encoder())
	testcase := datatransferobjects.CostRateDTO{CalcForOneWord: true, Name: "a", ForOneWordSymbol: 2, ForOneSquare: 3}
	got := ds.convertCostRateToModel(&testcase)
	if got.CalcForOneWord != testcase.CalcForOneWord || got.Name != testcase.Name ||
		got.ForOnecm2 != testcase.ForOneSquare || got.ForOneWordSymbol != testcase.ForOneWordSymbol {
		t.Fatalf("want %v,got %v", testcase, got)
	}
}

func TestCostRateToDto(t *testing.T) {
	CreateLogger()
	ds := NewServerStorage(nil, encodedecoder.NewBase64Encoder())
	testcase := CostRateFront{CalcForOneWord: true, Name: "a", ForOneWordSymbol: 2, ForOnecm2: 3}
	got := ds.convertCostRateToDto(&testcase)
	if testcase.CalcForOneWord != got.CalcForOneWord || testcase.Name != got.Name ||
		testcase.ForOnecm2 != got.ForOneSquare || testcase.ForOneWordSymbol != got.ForOneWordSymbol {
		t.Fatalf("want %v,got %v", testcase, got)
	}
}

func BenchmarkCostRateToModel(b *testing.B) {
	CreateLogger()
	ds := NewServerStorage(nil, encodedecoder.NewBase64Encoder())
	testcase := datatransferobjects.CostRateDTO{CalcForOneWord: true, Name: "a", ForOneWordSymbol: 2, ForOneSquare: 3}
	for i := 0; i < b.N; i++ {
		ds.convertCostRateToModel(&testcase)
	}
}
