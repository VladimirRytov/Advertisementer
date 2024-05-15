package presenter

import (
	"reflect"
	"testing"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
)

func TestClientToView(t *testing.T) {
	dc := NewDataConverter()
	CreateLogger()
	want := ClientDTO{
		Name:                  "Вася",
		Phones:                "8-800-555-35-35",
		Email:                 "mail@mail.com",
		AdditionalInformation: "asdasd",
	}
	testCase := &datatransferobjects.ClientDTO{
		Name:                  "Вася",
		Phones:                "8-800-555-35-35",
		Email:                 "mail@mail.com",
		AdditionalInformation: "asdasd",
	}
	got := dc.ClientToViewDTO(testCase)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestOrderTiView(t *testing.T) {
	dc := NewDataConverter()
	CreateLogger()
	want := OrderDTO{
		ID:            1,
		ClientName:    "Вася",
		Cost:          "1234,55",
		PaymentType:   "zxc",
		CreatedDate:   "12.12.2023",
		PaymentStatus: true,
	}
	testCase := &datatransferobjects.OrderDTO{
		ID:            1,
		ClientName:    "Вася",
		Cost:          123455,
		PaymentType:   "zxc",
		CreatedDate:   time.Date(2023, 12, 12, 0, 0, 0, 0, time.UTC),
		PaymentStatus: true,
	}
	got := dc.OrderToViewDTO(testCase)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestBlockAdvertisementToView(t *testing.T) {
	dc := NewDataConverter()
	CreateLogger()
	want := BlockAdvertisementDTO{
		Advertisement: Advertisement{
			ID:           1,
			OrderID:      1,
			ReleaseCount: 2,
			ReleaseDates: string("15.12.2023, 12.12.2007, 22.05.2000, 02.05.2000"),
			Cost:         "22,23",
			Text:         "asd",
			Tags:         string("asd, dsa"),
			ExtraCharge:  string("asd, dsa"),
		},
		Size:     10,
		FileName: "asd.jpg",
	}
	testCase := &datatransferobjects.BlockAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           1,
			OrderID:      1,
			ReleaseCount: 2,
			Cost:         2223,
			Text:         "asd",
			Tags:         []string{"asd", "dsa"},
			ExtraCharges: []string{"asd", "dsa"},
			ReleaseDates: []time.Time{
				time.Date(2023, 12, 15, 0, 0, 0, 0, time.UTC),
				time.Date(2007, 12, 12, 0, 0, 0, 0, time.UTC),
				time.Date(2000, 5, 22, 0, 0, 0, 0, time.UTC),
				time.Date(2000, 5, 2, 0, 0, 0, 0, time.UTC)},
		},
		Size:     10,
		FileName: "asd.jpg",
	}
	got := dc.BlockAdvertisementToViewDTO(testCase)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestLineAdvertisementToView(t *testing.T) {
	dc := NewDataConverter()
	CreateLogger()
	want := LineAdvertisementDTO{
		Advertisement: Advertisement{
			ID:           1,
			OrderID:      1,
			ReleaseCount: 2,
			ReleaseDates: string("15.12.2023, 12.12.2007, 22.05.2000, 02.05.2000"),
			Cost:         "22,23",
			Text:         "asd",
			Tags:         string("asd, dsa"),
			ExtraCharge:  string("asd, dsa"),
		},
	}
	testCase := &datatransferobjects.LineAdvertisementDTO{
		Advertisement: datatransferobjects.Advertisement{
			ID:           1,
			OrderID:      1,
			ReleaseCount: 2,
			Cost:         2223,
			Text:         "asd",
			Tags:         []string{"asd", "dsa"},
			ExtraCharges: []string{"asd", "dsa"},
			ReleaseDates: []time.Time{
				time.Date(2023, 12, 15, 0, 0, 0, 0, time.UTC),
				time.Date(2007, 12, 12, 0, 0, 0, 0, time.UTC),
				time.Date(2000, 5, 22, 0, 0, 0, 0, time.UTC),
				time.Date(2000, 5, 2, 0, 0, 0, 0, time.UTC)},
		},
	}
	got := dc.LineAdvertisementToViewDTO(testCase)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestTagToView(t *testing.T) {
	dc := NewDataConverter()
	CreateLogger()
	want := TagDTO{
		TagName: "asd",
		TagCost: "22,23",
	}
	testCase := &datatransferobjects.TagDTO{
		TagName: "asd",
		TagCost: 2223,
	}
	got := dc.TagToViewDTO(testCase)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestExtraChargeToView(t *testing.T) {
	dc := NewDataConverter()
	CreateLogger()
	want := ExtraChargeDTO{
		ChargeName: "asd",
		Multiplier: "200",
	}
	testCase := &datatransferobjects.ExtraChargeDTO{
		ChargeName: "asd",
		Multiplier: 200,
	}
	got := dc.ExtraChargeToViewDTO(testCase)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestLocalDsnToView(t *testing.T) {
	CreateLogger()
}

func TestNetworkDsnToView(t *testing.T) {
	CreateLogger()
}

func TestCostRateToView(t *testing.T) {
	dc := NewDataConverter()
	CreateLogger()
	testCase := datatransferobjects.CostRateDTO{
		CalcForOneWord:   true,
		Name:             "asd",
		ForOneWordSymbol: 5000,
		ForOneSquare:     1000,
	}
	want := CostRateDTO{Name: "asd", OneWordOrSymbol: "50,00", Onecm2: "10,00", CalcForOneWord: true}
	got := dc.CostRateToViewDTO(&testCase)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v,got %v", want, got)
	}
}
func TestClosestRelease(t *testing.T) {
	CreateLogger()
	dc := NewDataConverter()
	arr := [][]string{
		{"22.02.2000", "12.11.2007", "17.11.2007", "25.01.2023", "25.11.2023"},
		{"01.04.2023", "01.04.2023", "01.05.2023", "13.10.2023"},
		{"06.09.2023", "06.10.2023", "05.11.2023"},
		{"01.04.2024", "01.05.2024", "01.06.2024", "01.07.2024", "01.08.2024", "13.10.2024"},
		{"06.09.2023"}, {"05.11.2023"}, {"02.05.2023", "01.06.2023", "12.08.2023", "10.09.2023", "18.11.2023"},
		{"12.01.2023", "02.03.2023", "09.04.2023", "19.05.2023", "19.07.2023", "14.09.2023", "19.10.2023", "06.11.2023"},
		{"22.02.2000", "25.11.2023", "26.11.2023", "25.12.2023", "26.12.2023", "27.12.2023"}}

	want := []time.Time{
		time.Date(2023, 11, 25, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 10, 26, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 11, 5, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 10, 26, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 11, 5, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 11, 18, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 11, 6, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 11, 25, 0, 0, 0, 0, time.UTC),
	}
	for i, v := range arr {
		tt := make([]time.Time, 0)
		for j := range v {
			parsed, err := time.Parse("02.01.2006", v[j])
			if err != nil {
				t.Fatalf("got err %v", err)
			}
			tt = append(tt, parsed)
		}
		got, _ := dc.ClosestRelease(tt, time.Date(2023, 10, 26, 0, 0, 0, 0, time.UTC))
		if got != want[i] {
			t.Fatalf("want %v,got %v", want[i], got)
		}
	}
}
func BenchmarkClosestRelease(b *testing.B) {
	dc := NewDataConverter()
	CreateLogger()
	arr := []time.Time{
		time.Date(2023, 1, 12, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 3, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 4, 9, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 5, 19, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 7, 19, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 9, 14, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 10, 19, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 11, 6, 0, 0, 0, 0, time.UTC)}
	for i := 0; i < b.N; i++ {
		dc.ClosestRelease(arr, time.Date(2023, 10, 26, 0, 0, 0, 0, time.UTC))
	}
}
