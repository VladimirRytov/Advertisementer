package presenter

import (
	"testing"
)

func BenchmarkSelectedReleaseDatesToStringOld(b *testing.B) {
	dc := NewDataConverter()
	ar := []SelectedTagDTO{{true, "asd"}, {false, "as2d"}, {true, "asd"}, {false, "asdfsd"}, {false, "as342d"}, {false, "avbnsd"}, {true, "asd"}, {false, "asdfsd"}, {false, "as342d"}, {false, "avbnsd"}}
	for i := 0; i < b.N; i++ {
		dc.SelectedTagsToStringOld(ar)
	}
}

func BenchmarkSelectedReleaseDatesToString(b *testing.B) {
	dc := NewDataConverter()
	ar := []SelectedTagDTO{{true, "asd"}, {false, "as2d"}, {true, "asd"}, {false, "asdfsd"}, {false, "as342d"}, {false, "avbnsd"}, {true, "asd"}, {false, "asdfsd"}, {false, "as342d"}, {false, "avbnsd"}}
	for i := 0; i < b.N; i++ {
		dc.SelectedTagsToString(ar)
	}
}

func BenchmarkYearMonthDayToString(b *testing.B) {
	dc := NewDataConverter()
	for i := 0; i < b.N; i++ {
		dc.YearMonthDayToString(2023, 12, 11)
	}
}
