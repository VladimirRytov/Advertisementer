package orm

import (
	"reflect"
	"testing"
	"time"
)

func TestFetchTagsName(t *testing.T) {
	CreateLogger()
	testCase := []Tag{{Name: "a", Cost: 1}, {Name: "b", Cost: 2}}
	want := []string{}
	for _, v := range testCase {
		want = append(want, v.Name)
	}
	got := fetchTagsName(testCase)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v,got %v", want, got)
	}
}

func TestFetchExtraChargeName(t *testing.T) {
	CreateLogger()
	testCase := []ExtraCharge{{Name: "a", Multiplier: 1}, {Name: "b", Multiplier: 2}}
	want := []string{}
	for _, v := range testCase {
		want = append(want, v.Name)
	}
	got := fetchExtraChargesName(testCase)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v,got %v", want, got)
	}
}
func TestReleaseDates(t *testing.T) {
	CreateLogger()
	testCase := []ReleaseDates{{ReleaseDate: time.Date(2023, 12, 12, 0, 0, 0, 0, time.Local)},
		{ReleaseDate: time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local)},
	}
	want := []time.Time{}
	for _, v := range testCase {
		want = append(want, v.ReleaseDate)
	}
	got := fetchReleaseDates(testCase)
	for i := range got {
		if got[i].Day() != testCase[i].ReleaseDate.Day() || got[i].Month() != testCase[i].ReleaseDate.Month() ||
			got[i].Year() != testCase[i].ReleaseDate.Year() {
			t.Fatalf("want %v,got %v", want, got)
		}
	}
}
