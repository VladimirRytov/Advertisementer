package advertisements

import (
	"testing"
)

func TestSetSize(t *testing.T) {
	bl := NewAdvertisementBlock()
	want := AdvertisementBlock{size: 123}
	bl.SetSize(123)
	if bl.size != want.size {
		t.Errorf("want %v struct, got %v", want, bl)
	}
}
func TestSize(t *testing.T) {
	bl := &AdvertisementBlock{size: 123}
	size := bl.Size()
	if bl.size != size {
		t.Errorf("want size: %d, got size: %d", size, bl.size)
	}
}

func TestComment(t *testing.T) {
	want := "asdzxc"
	testadv := AdvertisementBlock{}
	testadv.SetComment(want)
	got := testadv.Comment()
	if want != got {
		t.Errorf("want content = %s, got content = %s", want, got)
	}
}

func TestSetComment(t *testing.T) {
	ad := AdvertisementBlock{}
	want := "фывячсчп"
	ad.SetComment(want)
	if ad.comment != want {
		t.Errorf("got %s != %s", ad.comment, want)
	}
}

func TestFilePath(t *testing.T) {
	want := "asdzxc"
	testadv := AdvertisementBlock{}
	err := testadv.SetFileName(want)
	if err != nil {
		t.Fatal(err)
	}
	got := testadv.FileName()
	if want != got {
		t.Errorf("want filePath  = %s, got filePath = %s", want, got)
	}
}

func TestSetFilePath(t *testing.T) {
	ad := AdvertisementBlock{}
	want := "фывячсчп"
	err := ad.SetFileName(want)
	if err != nil {
		t.Fatal(err)
	}
	if ad.fileName != want {
		t.Errorf("got %s != %s", ad.comment, want)
	}
}

func TestSetFilePathError(t *testing.T) {
	ad := AdvertisementBlock{}
	err := ad.SetFileName("/фывячсчп.df")
	if err == nil {
		t.Fatal("want error, got nil")
	}
	err = ad.SetFileName("\\фывячсчп.df")
	if err == nil {
		t.Fatal("want error, got nil")
	}
}
