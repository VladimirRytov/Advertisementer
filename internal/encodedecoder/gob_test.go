package encodedecoder

import (
	"bytes"
	"testing"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
)

func TestToGob(t *testing.T) {
	CreateLogger()
	var f bytes.Buffer
	err := ToGob(&f, testModel)
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkToGob(b *testing.B) {
	CreateLogger()

	var f bytes.Buffer
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ToGob(&f, testModel)
		f.Reset()
	}
}

func TestFromGob(t *testing.T) {
	CreateLogger()
	var b bytes.Buffer
	ToGob(&b, testModel)
	err := FromGob(&datatransferobjects.JsonStr{}, &b)
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkFromGob(b *testing.B) {
	CreateLogger()

	var f bytes.Buffer
	ToGob(&f, testModel)

	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		FromGob(&datatransferobjects.JsonStr{}, &f)
	}
}
