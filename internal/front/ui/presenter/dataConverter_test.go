package presenter

import (
	"testing"
	"time"
)

func TestCostToView(t *testing.T) {
	CreateLogger()
	dc := NewDataConverter()
	wantCases := []string{"123,15", "121231231231233,15", "0,15", "0,02", "0,40", "22,00"}
	testcases := []int{12315, 12123123123123315, 15, 2, 40, 2200}
	for i := range wantCases {
		got := dc.CostToView(testcases[i])
		if wantCases[i] != got {
			t.Fatalf("want %s, got %s", wantCases[i], got)
		}
	}
}

func TestDateToView(t *testing.T) {
	CreateLogger()
	dc := NewDataConverter()
	wantCases := []string{"15.12.2023", "12.12.2007", "22.05.2000", "02.05.2000"}
	testCases := []time.Time{
		time.Date(2023, 12, 15, 0, 0, 0, 0, time.UTC),
		time.Date(2007, 12, 12, 0, 0, 0, 0, time.UTC),
		time.Date(2000, 5, 22, 0, 0, 0, 0, time.UTC),
		time.Date(2000, 5, 2, 0, 0, 0, 0, time.UTC)}
	for i := range wantCases {
		got := dc.DateToView(testCases[i])
		if wantCases[i] != got {
			t.Fatalf("want %s, got %s", wantCases[i], got)
		}
	}
}

func TestDateToDto(t *testing.T) {
	CreateLogger()
	dc := NewDataConverter()
	want := []time.Time{
		time.Date(2023, 12, 15, 0, 0, 0, 0, time.UTC),
		time.Date(2007, 11, 12, 0, 0, 0, 0, time.UTC),
		time.Date(2000, 2, 22, 0, 0, 0, 0, time.UTC),
		time.Date(20335, 5, 2, 0, 0, 0, 0, time.UTC)}
	testCase := []string{"15.12.2023", "12.11.2007", "22.02.2000", "02.05.20335"}
	for i := range want {
		got, err := dc.DateToDto(testCase[i])
		if err != nil {
			t.Fatalf("got unexpected error: %v", err)
		}
		if !got.Equal(want[i]) {
			t.Errorf("want %v, got %v", want[i], got)
		}
	}
	_, err := dc.DateToDto("2.3.2023")
	if err.Error() != "dateToDto: date pattern not found" {
		t.Fatalf("want pattern error, got %v", err)
	}
}

func TestCostToDto(t *testing.T) {
	CreateLogger()
	dc := NewDataConverter()
	want := []int{12312312, 345345345, 67345348678678, 0, -12, -12, -0, 10000, 40}
	testCase := []string{"123123.12", "3453453,45", "673453486786.78", "0", "-0.12", "-0,12", "+0", "100", "0.4"}
	for i := range want {
		got, err := dc.CostToDto(testCase[i])
		if err != nil {
			t.Fatalf("got unexpected error: %v,iterator %d", err, i)
		}
		if got != want[i] {
			t.Errorf("want %v, got %v", want[i], got)
		}
	}
	_, err := dc.CostToDto("123.231")
	if err.Error() != "error: not found cost pattern in variable" {
		t.Fatalf("want error \"error: not found cost pattern in variable\", got %v", err)
	}
	err = nil
	_, err = dc.CostToDto("")
	if err.Error() != "error: not found cost pattern in variable" {
		t.Fatalf("want error \"error: not found cost pattern in variable\", got %v", err)
	}
}

func TestConvertFileSize(t *testing.T) {
	testCases := []int{1024, 1111, 111100}
	want := []string{"1", "1.08", "108.5"}
	dc := NewDataConverter()
	for i := range testCases {
		dig := dc.convertFileSize(testCases[i], KiB)
		if dig != want[i] {
			t.Fatalf("want %v, got digit = %s", want[i], dig)
		}
	}
}

func TestCalcFileSize(t *testing.T) {
	testCases := []int{1301 * KiB, 111 * MiB, 1113 * KiB << 10}
	want := []string{"1.27 Mib", "111 Mib", "1.09 Gib"}
	dc := NewDataConverter()
	for i := range testCases {
		got := dc.calcFileSize(testCases[i])
		if got != want[i] {
			t.Fatalf("want %v, got  %s", want[i], got)
		}
	}
}

func BenchmarkDateToview(b *testing.B) {
	CreateLogger()
	dc := NewDataConverter()
	for i := 0; i < b.N; i++ {
		dc.DateToView(time.Date(2023, 12, 15, 0, 0, 0, 0, time.UTC))
	}
}

func BenchmarkCostToView(b *testing.B) {
	dc := NewDataConverter()
	CreateLogger()
	for i := 0; i < b.N; i++ {
		dc.CostToView(1234567)
	}
}

func BenchmarkCalcFileSize(b *testing.B) {
	dc := NewDataConverter()
	CreateLogger()
	for i := 0; i < b.N; i++ {
		dc.calcFileSize(9876543)
	}
}

func BenchmarkConvertFileSize(b *testing.B) {
	dc := NewDataConverter()
	CreateLogger()
	for i := 0; i < b.N; i++ {
		dc.convertFileSize(9876543, 1)
	}
}

func BenchmarkConvertFileSizeOld(b *testing.B) {
	dc := NewDataConverter()
	CreateLogger()
	for i := 0; i < b.N; i++ {
		dc.convertFileSizeOld(9876543, 1)
	}
}
