package advertisements

import (
	"reflect"
	"testing"
	"time"
)

// Общие параметры объявлений
func TestNewAdvertisement(t *testing.T) {
	want := advertisement{}
	got := newAdvertisement()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestSetReleaseOld(t *testing.T) {
	a := advertisement{}
	releaseDate := []time.Time{time.Date(2024, 2, 22, 0, 0, 0, 0, time.Local),
		time.Date(2023, 2, 22, 0, 0, 0, 0, time.Local), time.Date(2024, 2, 22, 0, 0, 0, 0, time.Local), time.Date(2023, 2, 22, 0, 0, 0, 0, time.Local)}

	want := []time.Time{time.Date(2023, 2, 22, 0, 0, 0, 0, time.Local), time.Date(2024, 2, 22, 0, 0, 0, 0, time.Local)}
	for i := range releaseDate {
		a.AppendReleaseDate(releaseDate[i])
	}
	if !reflect.DeepEqual(want, a.releaseDates) {
		t.Errorf("want %v,got %v", want, a.releaseDates)
	}
}

func TestSetReleaseDate(t *testing.T) {
	a := advertisement{}
	releaseDate := []time.Time{time.Date(2024, 2, 22, 0, 0, 0, 0, time.Local),
		time.Date(2023, 2, 22, 0, 0, 0, 0, time.Local), time.Date(2024, 2, 22, 0, 0, 0, 0, time.Local), time.Date(2023, 2, 22, 0, 0, 0, 0, time.Local)}

	want := []time.Time{time.Date(2023, 2, 22, 0, 0, 0, 0, time.Local), time.Date(2024, 2, 22, 0, 0, 0, 0, time.Local)}

	a.SetReleaseDates(releaseDate)
	if !reflect.DeepEqual(want, a.releaseDates) {
		t.Errorf("want %v,got %v", want, a.releaseDates)
	}
}

func TestSetAdvOrderId(t *testing.T) {
	want := advertisement{orderId: 2345}
	got := newAdvertisement()
	got.SetOrderId(2345)
	if want.orderId != got.orderId {
		t.Errorf("got %d != %d", got.orderId, want.orderId)
	}
}

func TestSetCost(t *testing.T) {
	ad := advertisement{}
	want := 640
	ad.SetCost(want)
	if ad.cost != want {
		t.Errorf("got %d,but want %d", ad.cost, want)
	}
}
func TestSetId(t *testing.T) {
	want := advertisement{id: 10}
	got := advertisement{}
	got.SetId(10)
	if want.id != got.id {
		t.Errorf("want id = %d, got id = %d", want.id, got.id)
	}
}

func TestId(t *testing.T) {
	want := 123
	testadv := advertisement{}
	testadv.SetId(123)
	got := testadv.Id()
	if want != got {
		t.Errorf("want id = %d, got id = %d", want, got)
	}
}

func TestGetOrderId(t *testing.T) {
	want := 432
	ad := newAdvertisement()
	ad.orderId = want
	got := ad.OrderId()
	if want != got {
		t.Errorf("want content = %d, got content = %d", want, got)
	}
}
func TestAppendReleaseDates(t *testing.T) {
	want := [][]time.Time{
		{
			time.Date(2023, 2, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 3, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 5, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 6, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 7, 22, 0, 0, 0, 0, time.Local)},
		{
			time.Date(2021, 7, 22, 0, 0, 0, 0, time.Local)},
	}
	cases := [][]time.Time{
		{
			time.Date(2023, 2, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 3, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 5, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 6, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 6, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 7, 22, 0, 0, 0, 0, time.Local)},
		{
			time.Date(2021, 7, 22, 0, 0, 0, 0, time.Local)},
	}
	for i := range want {
		testadv := newAdvertisement()
		for _, v := range cases[i] {
			testadv.AppendReleaseDate(v)
		}
		got := testadv.ReleaseDates()
		if len(want[i]) != len(got) {
			t.Error("different length")
		}
		if !reflect.DeepEqual(want[i], got) {
			t.Errorf("got date %v has count %v", got, want[i])
		}
	}
}
func TestReleaseCount(t *testing.T) {
	want := int16(3)
	adv := newAdvertisement()
	adv.reseaseCount = 3
	if got := adv.ReseaseCount(); got != want {
		t.Errorf("want %v,got %v", want, got)
	}
}
func TestRemoveReleaseDates(t *testing.T) {
	testSlice := []time.Time{
		time.Date(2023, 2, 22, 0, 0, 0, 0, time.Local),
		time.Date(2023, 3, 22, 0, 0, 0, 0, time.Local),
		time.Date(2023, 5, 22, 0, 0, 0, 0, time.Local),
		time.Date(2023, 6, 22, 0, 0, 0, 0, time.Local),
		time.Date(2023, 7, 22, 0, 0, 0, 0, time.Local),
	}

	want := [][]time.Time{{},
		{
			time.Date(2023, 2, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 3, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 5, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 7, 22, 0, 0, 0, 0, time.Local),
		},
		{
			time.Date(2023, 2, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 3, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 5, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 6, 22, 0, 0, 0, 0, time.Local),
		},
		{
			time.Date(2023, 3, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 5, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 6, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 7, 22, 0, 0, 0, 0, time.Local),
		},
		{},
	}

	remove := [][]time.Time{
		{
			time.Date(2023, 2, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 3, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 5, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 6, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 7, 22, 0, 0, 0, 0, time.Local),
		},
		{
			time.Date(2023, 6, 22, 0, 0, 0, 0, time.Local),
		},
		{
			time.Date(2023, 7, 22, 0, 0, 0, 0, time.Local),
		},
		{
			time.Date(2023, 2, 22, 0, 0, 0, 0, time.Local),
		},
		{
			time.Date(2023, 2, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 3, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 5, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 6, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 7, 22, 0, 0, 0, 0, time.Local),
			time.Date(2023, 2, 22, 0, 0, 0, 0, time.Local),
		},
	}

	for i := range want {
		testadv := &advertisement{}
		testadv.releaseDates = testSlice
		testadv.reseaseCount = int16(len(testSlice))
		for _, v := range remove[i] {
			testadv.RemoveReleaseDate(v)
		}
		if !reflect.DeepEqual(testadv.releaseDates, want[i]) && len(want[i]) != int(testadv.reseaseCount) {
			t.Errorf("want %v, got %v", want[i], testadv.releaseDates)
		}
	}

}
func TestAdvCost(t *testing.T) {
	want := 640
	ad := &advertisement{cost: want}
	got := ad.Cost()
	if got != want {
		t.Errorf("got %d,but want %d", ad.cost, want)
	}
}

func TestTags(t *testing.T) {
	want := []string{"a", "b"}
	al := NewAdvertisementLine()
	al.tags = want
	got := al.Tags()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestAppendTags(t *testing.T) {
	want := [][]string{{"a", "b"}, {"a", "b"}, {"a", "b", "c"}}
	testCases := [][]string{{"a", "b"}, {"b"}, {"b", "c"}}
	advLines := newAdvertisement()
	for i, mapWant := range testCases {
		for _, v := range mapWant {
			advLines.AppendTag(v)
		}
		if !reflect.DeepEqual(advLines.tags, want[i]) {
			t.Errorf("len %v not equal %v", want[i], advLines.tags)
		}
	}
}

func TestRemoveTags(t *testing.T) {
	testTags := []string{"a", "b", "c"}
	removeTags := [][]string{{"a"}, {"a", "b"}, {}, {"a", "a", "b"}, {"a", "f", "b"}}
	want := [][]string{{"b", "c"}, {"c"}, {"a", "b", "c"}, {"c"}, {"c"}}
	for i := range removeTags {
		testAdv := newAdvertisement()
		testAdv.tags = testTags
		for _, v := range removeTags[i] {
			testAdv.RemoveTag(v)
		}
		got := testAdv.Tags()
		if !reflect.DeepEqual(got, want[i]) {
			t.Errorf("want %v,got %v", want[i], got)
		}
	}
}

func TestExtraCharge(t *testing.T) {
	want := []string{"a", "b"}
	al := NewAdvertisementLine()

	for i := range want {
		al.AppendExtraCharge(want[i])
	}
	got := al.ExtraCharge()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("len %v not equal %v", want, got)
	}
}

func TestAppendCostRate(t *testing.T) {
	want := [][]string{{"a", "b"}, {"a", "b"}, {"a", "b", "c"}}
	testCases := [][]string{{"a", "b"}, {"b"}, {"b", "c"}}
	advLines := newAdvertisement()
	for i, mapWant := range testCases {
		for _, v := range mapWant {
			advLines.AppendExtraCharge(v)
		}
		if !reflect.DeepEqual(advLines.ExtraCharge(), want[i]) {
			t.Errorf("len %v not equal %v", want[i], advLines.tags)
		}
	}
}
func TestRemoveExtraCharge(t *testing.T) {
	testExtraCharge := []string{"a", "b", "c"}
	removeExtraCharge := [][]string{{"a"}, {"a", "b"}, {}, {"a", "a", "b"}, {"a", "f", "b"}}
	want := [][]string{{"b", "c"}, {"c"}, {"a", "b", "c"}, {"c"}, {"c"}}
	for i := range removeExtraCharge {
		testAdv := newAdvertisement()
		testAdv.extraCharge = testExtraCharge
		for _, v := range removeExtraCharge[i] {
			testAdv.RemoveExtraCharge(v)
		}
		got := testAdv.ExtraCharge()
		if !reflect.DeepEqual(got, want[i]) {
			t.Errorf("want %v,got %v", want[i], got)
		}
	}
}

func BenchmarkAppendReleaseDate(b *testing.B) {
	arr := []time.Time{
		time.Date(2025, 2, 22, 0, 0, 0, 0, time.Local),
		time.Date(2023, 3, 22, 0, 0, 0, 0, time.Local),
		time.Date(2023, 5, 22, 0, 0, 0, 0, time.Local),
		time.Date(2023, 6, 22, 0, 0, 0, 0, time.Local),
		time.Date(2023, 7, 22, 0, 0, 0, 0, time.Local)}
	for i := 0; i < b.N; i++ {
		testadv := newAdvertisement()
		testadv.releaseDates = arr
		testadv.AppendReleaseDate(time.Now())
	}
}
func BenchmarkAppendReleaseDateNew(b *testing.B) {
	arr := []time.Time{
		time.Date(2025, 2, 22, 0, 0, 0, 0, time.Local),
		time.Date(2023, 3, 22, 0, 0, 0, 0, time.Local),
		time.Date(2023, 5, 22, 0, 0, 0, 0, time.Local),
		time.Date(2023, 6, 22, 0, 0, 0, 0, time.Local),
		time.Date(2023, 7, 22, 0, 0, 0, 0, time.Local)}
	for i := 0; i < b.N; i++ {
		testadv := newAdvertisement()
		testadv.releaseDates = arr
		testadv.AppendReleaseDateNew(time.Now())
	}
}

func BenchmarkSetTags(b *testing.B) {
	testCases := [][]string{{"a", "b"}, {"b"}, {"b", "c"}}
	a := newAdvertisement()
	for i := 0; i < b.N; i++ {
		for _, mapWant := range testCases {
			for _, v := range mapWant {
				a.AppendTag(v)
			}
		}

	}
}

func BenchmarkAppendTags(b *testing.B) {
	testCases := [][]string{{"a", "b"}, {"b"}, {"b", "c"}}
	advLines := newAdvertisement()

	for i := 0; i < b.N; i++ {
		for _, mapWant := range testCases {
			for _, v := range mapWant {
				advLines.AppendTag(v)
			}
		}
	}
}

func BenchmarkSetReleaseDate(b *testing.B) {
	a := advertisement{}
	releaseDate := []time.Time{time.Date(2024, 2, 22, 0, 0, 0, 0, time.Local),
		time.Date(2023, 2, 22, 0, 0, 0, 0, time.Local), time.Date(2044, 2, 22, 0, 0, 0, 0, time.Local), time.Date(2123, 2, 22, 0, 0, 0, 0, time.Local),
		time.Date(2024, 2, 22, 0, 0, 0, 0, time.Local), time.Date(2045, 2, 22, 0, 0, 0, 0, time.Local), time.Date(2223, 2, 22, 0, 0, 0, 0, time.Local),
		time.Date(2025, 2, 22, 0, 0, 0, 0, time.Local), time.Date(2046, 2, 22, 0, 0, 0, 0, time.Local), time.Date(2323, 2, 22, 0, 0, 0, 0, time.Local),
		time.Date(2026, 2, 22, 0, 0, 0, 0, time.Local), time.Date(2074, 2, 22, 0, 0, 0, 0, time.Local), time.Date(2443, 2, 22, 0, 0, 0, 0, time.Local),
		time.Date(2027, 2, 22, 0, 0, 0, 0, time.Local), time.Date(2064, 2, 22, 0, 0, 0, 0, time.Local), time.Date(2453, 2, 22, 0, 0, 0, 0, time.Local),
		time.Date(2028, 2, 22, 0, 0, 0, 0, time.Local), time.Date(2044, 2, 22, 0, 0, 0, 0, time.Local), time.Date(2553, 2, 22, 0, 0, 0, 0, time.Local),
		time.Date(2029, 2, 22, 0, 0, 0, 0, time.Local), time.Date(2224, 2, 22, 0, 0, 0, 0, time.Local), time.Date(2023, 2, 22, 0, 0, 0, 0, time.Local),
		time.Date(2030, 2, 22, 0, 0, 0, 0, time.Local), time.Date(2324, 2, 22, 0, 0, 0, 0, time.Local), time.Date(3023, 2, 22, 0, 0, 0, 0, time.Local),
		time.Date(2031, 2, 22, 0, 0, 0, 0, time.Local), time.Date(2424, 2, 22, 0, 0, 0, 0, time.Local), time.Date(4023, 2, 22, 0, 0, 0, 0, time.Local),
		time.Date(2032, 2, 22, 0, 0, 0, 0, time.Local), time.Date(2524, 2, 22, 0, 0, 0, 0, time.Local), time.Date(5023, 2, 22, 0, 0, 0, 0, time.Local),
		time.Date(2033, 2, 22, 0, 0, 0, 0, time.Local), time.Date(2624, 2, 22, 0, 0, 0, 0, time.Local), time.Date(6023, 2, 22, 0, 0, 0, 0, time.Local)}
	for i := 0; i < b.N; i++ {
		a.SetReleaseDates(releaseDate)
	}
}
