package requestscontroller

import (
	"testing"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

func TestCheckCostString(t *testing.T) {
	dc := NewRequestsHandler(nil, nil, nil)
	testCases := []string{"123", "123.22", "32144,052", "15", "15,25", "123asd,45", "123,esda45"}
	want := []bool{true, true, false, true, true, false, false}
	for i := range testCases {
		got := dc.CheckCostString(testCases[i])
		if got != want[i] {
			t.Fatalf("got %s", testCases[i])
		}
	}
}

func TestCheckAdvCostString(t *testing.T) {
	dc := NewRequestsHandler(nil, nil, nil)
	testCases := []string{"123", "123.22", "32144,052", "-15", "-15,25", "123asd,45", "123,esda45"}
	want := []bool{true, true, false, false, false, false, false}
	for i := range testCases {
		got := dc.CheckAdvCostString(testCases[i])
		if got != want[i] {
			t.Fatalf("got %s", testCases[i])
		}
	}
}

func TestCompareStrings(t *testing.T) {
	dc := NewRequestsHandler(nil, nil, nil)
	testCases := [][]string{{"Вася", "Петя"}, {"Серёжа", "Игорь"}, {"", "фыв"}, {"фыв", "asd"}, {"123,12", "113,11"}}
	want := []int{-1, 1, -1, 1, 1}
	for i := range testCases {
		got := dc.CompareStrings(testCases[i][0], testCases[i][1])
		if got != want[i] {
			t.Fatalf("got %s", testCases[i])
		}
	}
}

func TestReleasesInTimeRange(t *testing.T) {
	CreateLogger()
	dc := NewRequestsHandler(nil, nil, nil)
	dc.converter = presenter.NewDataConverter()
	testCases := []string{
		"22.02.2000, 12.11.2007, 17.11.2007, 25.01.2023, 25.11.2023",
		"01.04.2023, 01.04.2023, 01.05.2023, 13.10.2023",
		"06.09.2023, 06.10.2023, 05.11.2023",
		"01.04.2024, 01.05.2024, 01.06.2024, 01.07.2024, 01.08.2024, 13.10.2024",
		"02.05.2023, 01.06.2023, 12.08.2023, 10.09.2023, 18.11.2023",
		"12.01.2023, 02.03.2023, 09.04.2023, 19.05.2023, 19.07.2023, 14.09.2023, 19.10.2023, 06.11.2023",
		"22.02.2000, 12.11.2007, 17.11.2007, 25.01.2023, 25.11.2023"}

	dateRange := [][]string{{"17.11.2007", "25.01.2023"},
		{"25.01.2023", ""}, {"", "05.11.2023"}, {"01.04.2025", "01.04.2027"}, {"02.05.2023", "03.05.2023"},
		{"", ""}, {"25.01.2023", "25.01.2023"}}

	want := []bool{true, true, true, false, true, true, false}
	for i := range testCases {
		got := dc.ReleasesInTimeRange(testCases[i], dateRange[i][0], dateRange[i][1])
		if got != want[i] {
			t.Fatalf("got %v want %v,iter %d", got, want[i], i)
		}
	}
}
func BenchmarkDateInTimeRange(t *testing.B) {
	CreateLogger()
	dc := NewRequestsHandler(nil, nil, nil)
	for i := 0; i < t.N; i++ {
		dc.ReleasesInTimeRange("12.01.2023, 02.03.2023, 09.04.2023, 19.05.2023, 19.07.2023, 14.09.2023, 19.10.2023, 06.11.2023",
			"02.03.2023", "19.05.2023")
	}
}

func BenchmarkCorrectCostString(b *testing.B) {
	dc := NewRequestsHandler(nil, nil, nil)
	for i := 0; i < b.N; i++ {
		dc.CheckCostString("-1255675675334,26")
	}
}

func BenchmarkCompareStrings(b *testing.B) {
	dc := NewRequestsHandler(nil, nil, nil)
	for i := 0; i < b.N; i++ {
		dc.CompareStrings("Игорь", "Петя")
	}
}

func BenchmarkCompareStringsRu(b *testing.B) {
	r := collate.New(language.Russian, collate.IgnoreCase)
	for i := 0; i < b.N; i++ {
		r.CompareString("Игорь", "Петя")
	}
}
