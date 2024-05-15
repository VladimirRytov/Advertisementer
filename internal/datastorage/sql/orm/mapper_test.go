package orm

import (
	"reflect"
	"testing"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
)

func TestConvertClientToDTO(t *testing.T) {
	CreateLogger()
	got := convertClientToDTO(&clientModel)
	if got.Name != clientDtoForTest.Name || got.Email != clientDtoForTest.Email || got.Phones != clientDtoForTest.Phones ||
		got.AdditionalInformation != clientDtoForTest.AdditionalInformation {
		t.Fatalf("want %v,got %v", clientDtoForTest, got)
	}
}
func TestConvertClientToModel(t *testing.T) {
	CreateLogger()
	got := convertClientToModel(&clientDtoForTest)
	if got.Name != clientDtoForTest.Name || got.Email != clientDtoForTest.Email || got.Phones != clientDtoForTest.Phones {
		t.Fatalf("want %v,got %v", clientDtoForTest, got)
	}
}

func TestConvertOrderToDTO(t *testing.T) {
	CreateLogger()
	got := convertOrderToDTO(&orderModel)
	if got.ID != orderDtoForTest.ID || got.ClientName != orderDtoForTest.ClientName || orderDtoForTest.PaymentStatus != got.PaymentStatus ||
		got.Cost != orderDtoForTest.Cost || !got.CreatedDate.Equal(orderDtoForTest.CreatedDate) {

		t.Fatalf("want %v,got %v", orderDtoForTest, got)
	}
}

func TestConvertOrderToModel(t *testing.T) {
	CreateLogger()
	got := convertOrderToModel(&orderDtoForTest)
	if got.ClientName != orderModel.ClientName || orderModel.PaymentStatus != got.PaymentStatus ||
		got.Cost != orderModel.Cost || got.CreatedDate != orderModel.CreatedDate {

		t.Fatalf("want %v,got %v", orderDtoForTest, got)
	}
}

func TestConvertTagToDTO(t *testing.T) {
	CreateLogger()
	got := make([]datatransferobjects.TagDTO, 0, len(tagsModel))
	for _, v := range tagsModel {
		got = append(got, convertTagToDTO(&v))
	}
	for i := range got {
		if got[i].TagName != tagsDtoForTest[i].TagName || got[i].TagCost != tagsDtoForTest[i].TagCost {
			t.Fatalf("want %v,got %v", tagsDtoForTest, got)
		}
	}
}

func TestConvertTagsToModel(t *testing.T) {
	CreateLogger()
	got := make([]Tag, 0, len(tagsDtoForTest))
	for _, v := range tagsDtoForTest {
		got = append(got, convertTagToModel(&v))
	}
	for i := range got {
		if got[i].Name != tagsModel[i].Name || got[i].Cost != tagsModel[i].Cost {
			t.Fatalf("want %v,got %v", tagsModel, got)
		}
	}
}

func TestConvertExtraChargeToDTO(t *testing.T) {
	CreateLogger()
	got := make([]datatransferobjects.ExtraChargeDTO, 0, len(extraChargesModel))
	for _, v := range extraChargesModel {
		got = append(got, convertExtraChargeToDTO(&v))
	}
	for i := range got {
		if got[i].ChargeName != extraChargesDtoForTest[i].ChargeName || got[i].Multiplier != extraChargesDtoForTest[i].Multiplier {
			t.Fatalf("want %v,got %v", clientDtoForTest, got)
		}
	}
}

func TestConvertExtraChargeToModel(t *testing.T) {
	CreateLogger()
	got := make([]ExtraCharge, 0, len(extraChargesDtoForTest))
	for _, v := range extraChargesDtoForTest {
		got = append(got, convertExtraChargeToModel(&v))
	}
	for i := range got {
		if got[i].Name != extraChargesModel[i].Name || got[i].Multiplier != extraChargesModel[i].Multiplier {
			t.Fatalf("want %v,got %v", extraChargesModel, got)
		}
	}
}

func TestConvertReleaseDateToModel(t *testing.T) {
	CreateLogger()
	got := convertReleaseDateToModel(releaseDate)
	if !reflect.DeepEqual(got, *releaseDateModel) {
		t.Fatalf("want %v,got %v", *releaseDateModel, got)
	}
}

func TestConvertBlockAdvertisementToDTO(t *testing.T) {
	CreateLogger()
	got := convertBlockAdvertisementToDTO(&blockModel)
	if !reflect.DeepEqual(got, blockDtoForTest) {

		t.Fatalf("want %v,got %v", blockDtoForTest, got)
	}
}

func TestConvertBlockAdvertisementToModel(t *testing.T) {
	CreateLogger()
	got := convertBlockAdvertisementToModel(&blockDtoForTest)
	if got.ID != blockDtoForTest.ID || got.OrderID != blockDtoForTest.OrderID || got.Cost != blockDtoForTest.Cost ||
		got.FileName != blockDtoForTest.FileName || got.ReleaseCount != blockDtoForTest.ReleaseCount ||
		len(got.Tags) != len(blockDtoForTest.Tags) || len(got.ExtraCharges) != len(blockDtoForTest.ExtraCharges) {

		t.Fatalf("want %v,got %v", blockModel, got)
	}
	for i, v := range got.ReleaseDates {
		if !v.ReleaseDate.Equal(blockModel.ReleaseDates[i].ReleaseDate) {
			t.Fatalf("want %v,got %v", blockModel, got)
		}
	}
}

func TestConvertLineAdvertisementToDTO(t *testing.T) {
	CreateLogger()
	got := convertLineAdvertisementToDTO(&lineModel)
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
	got := convertLineAdvertisementToModel(&lineDtoForTest)
	if got.ID != lineDtoForTest.ID || got.OrderID != lineDtoForTest.OrderID || got.Cost != lineDtoForTest.Cost ||
		got.ReleaseCount != lineDtoForTest.ReleaseCount ||
		len(got.Tags) != len(lineDtoForTest.Tags) || len(got.ExtraCharges) != len(lineDtoForTest.ExtraCharges) {
		t.Fatalf("want %v,got %v", lineModel, got)
	}
	for i, v := range got.ReleaseDates {
		if !v.ReleaseDate.Equal(lineModel.ReleaseDates[i].ReleaseDate) {
			t.Fatalf("want %v,got %v", lineModel, got)
		}
	}
}

func TestTagNamesToModel(t *testing.T) {
	CreateLogger()
	testCase := []string{"a", "b", "c"}
	want := []*Tag{}
	for _, v := range testCase {
		want = append(want, &Tag{Name: v})
	}
	got := tagNamesToModel(testCase)
	for i := range got {
		if got[i].Name != want[i].Name {
			t.Fatalf("want %v,got %v", want[i], got[i])
		}
	}
}

func TestEextraChargesToModel(t *testing.T) {
	CreateLogger()
	testCase := []string{"a", "b", "c"}
	want := []*ExtraCharge{}
	for _, v := range testCase {
		want = append(want, &ExtraCharge{Name: v})
	}
	got := extraChargesToModel(testCase)
	for i := range got {
		if got[i].Name != want[i].Name {
			t.Fatalf("want %v,got %v", want[i], got[i])
		}
	}
}

func TestReleaseDatesToModel(t *testing.T) {
	CreateLogger()
	testCase := []time.Time{time.Date(2023, 12, 12, 0, 0, 0, 0, time.Local),
		time.Date(2024, 11, 15, 0, 0, 0, 0, time.Local),
		time.Date(2023, 1, 12, 0, 0, 0, 0, time.Local)}
	want := []*ReleaseDates{}
	for _, v := range testCase {
		want = append(want, &ReleaseDates{ReleaseDate: v})
	}
	got := releaseDatesToModel(testCase)

	for i := range got {
		if got[i].ReleaseDate.Day() != testCase[i].Day() || got[i].ReleaseDate.Month() != testCase[i].Month() ||
			got[i].ReleaseDate.Year() != testCase[i].Year() {
			t.Fatalf("want %v,got %v", want, got)
		}
	}
}
func TestCostRateToModel(t *testing.T) {
	CreateLogger()
	testcase := datatransferobjects.CostRateDTO{CalcForOneWord: true, Name: "a", ForOneWordSymbol: 2, ForOneSquare: 3}
	got := convertCostRateToModel(&testcase)
	if got.CalcForOneWord != testcase.CalcForOneWord || got.Name != testcase.Name ||
		got.OneSquare != testcase.ForOneSquare || got.OneWordOrSymbol != testcase.ForOneWordSymbol {
		t.Fatalf("want %v,got %v", testcase, got)
	}
}

func TestCostRateToDto(t *testing.T) {
	CreateLogger()
	testcase := CostRate{CalcForOneWord: true, Name: "a", OneWordOrSymbol: 2, OneSquare: 3}
	got := convertCostRateToDto(&testcase)
	if testcase.CalcForOneWord != got.CalcForOneWord || testcase.Name != got.Name ||
		testcase.OneSquare != got.ForOneSquare || testcase.OneWordOrSymbol != got.ForOneWordSymbol {
		t.Fatalf("want %v,got %v", testcase, got)
	}
}

func BenchmarkCostRateToModel(b *testing.B) {
	CreateLogger()
	testcase := datatransferobjects.CostRateDTO{CalcForOneWord: true, Name: "a", ForOneWordSymbol: 2, ForOneSquare: 3}
	for i := 0; i < b.N; i++ {
		convertCostRateToModel(&testcase)
	}
}
