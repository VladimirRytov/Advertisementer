package advertisements

import (
	"reflect"
	"testing"
)

func TestNewAdvertisementLine(t *testing.T) {
	al := NewAdvertisementLine()
	want := AdvertisementLine{advertisement: newAdvertisement()}
	if alk := reflect.ValueOf(al).Kind(); alk == reflect.Map && reflect.DeepEqual(al.tags, want.tags) {
		t.Errorf("want %v struct, got %v", want, al)
	}
}

func TestContent(t *testing.T) {
	want := "asdzxc"
	testadv := AdvertisementLine{}
	testadv.SetContent(want)
	got := testadv.Content()
	if want != got {
		t.Errorf("want content = %s, got content = %s", want, got)
	}
}

func TestSetContent(t *testing.T) {
	testCase := "   фывяч\nсч    dfd\nfgп   "
	ad := AdvertisementLine{}
	want := "фывяч сч dfd fgп"
	ad.SetContent(testCase)
	if ad.message != want {
		t.Errorf("got %s != %s", ad.message, want)
	}
}

func TestWordsCount(t *testing.T) {
	testCase := "   фывяч\nсч    dfd\nfgп   "
	ad := AdvertisementLine{}
	ad.SetContent(testCase)
	want := 4
	count := ad.WordsCount()
	if count != want {
		t.Errorf("want word count = %d, got %d", want, count)
	}
}

func TestZeroWordsCount(t *testing.T) {
	testCase := ""
	ad := AdvertisementLine{}
	ad.SetContent(testCase)
	want := 0
	count := ad.WordsCount()
	if count != want {
		t.Errorf("want word count = %d, got %d", want, count)
	}
}

func TestOneWordsCount(t *testing.T) {
	testCase := "G"
	ad := AdvertisementLine{}
	ad.SetContent(testCase)
	want := 1
	count := ad.WordsCount()
	if count != want {
		t.Errorf("want word count = %d, got %d", want, count)
	}
}
func TestSymbolsCount(t *testing.T) {
	testCase := "   фывяч\nсч   dfd\nfgп   "
	ad := AdvertisementLine{}
	ad.SetContent(testCase)
	want := 16
	count := ad.SymbolsCount()
	if count != want {
		t.Errorf("want word count = %d, got %d", want, count)
	}
}

func BenchmarkSetContentNew(b *testing.B) {
	testCase := "   фыdfgsdfgfgвяч\nсч dfd\nfgп   "
	ad := AdvertisementLine{}
	for i := 0; i < b.N; i++ {
		ad.SetContent(testCase)
	}
}

func BenchmarkSetContentOld(b *testing.B) {
	testCase := "   фыdfgsdfgfgвяч\nсч dfd\nfgп   "
	ad := AdvertisementLine{}
	for i := 0; i < b.N; i++ {
		ad.SetContentOld(testCase)
	}
}

func BenchmarkWordCount(b *testing.B) {
	testCase := "   фыdfgsdfgfgвяч\nсч dfd\nfgп   "
	ad := AdvertisementLine{}
	ad.SetContent(testCase)
	for i := 0; i < b.N; i++ {
		ad.WordsCount()
	}
}
