package encodedecoder

import (
	"slices"
	"testing"
)

func BenchmarkToBase64URL(b *testing.B) {
	a := []byte("asdasdaszxczxczxcbnvbnvcbfghjfghfdghdfgfcvbcxvbxcvbxcvvvbndfghdsrtyetyretyrt")
	for i := 0; i < b.N; i++ {
		ToBase64URL(a)
	}
}
func BenchmarkToBase64URLNew(b *testing.B) {
	b64 := NewBase64Encoder()
	test := []byte("asdasdaszxczxczxcbnvbnvcbfghjfghfdghdfgfcvbcxvbxcvbxcvvvbndfghdsrtyetyretyrt")
	for i := 0; i < b.N; i++ {
		b64.ToBase64(test)
	}
}

func TestToBase64(t *testing.T) {
	val := ToBase64([]byte("admin:admin"))
	if !slices.Equal(val, []byte("YWRtaW46YWRtaW4=")) {
		t.Fatalf("%s", val)
	}
}

func BenchmarkToBase64(b *testing.B) {
	b64 := NewBase64Encoder()
	test := []byte("admin:admin")
	for i := 0; i < b.N; i++ {
		b64.ToBase64(test)
	}
}

func BenchmarkToBase64String(b *testing.B) {
	b64 := NewBase64Encoder()
	test := []byte("admin:admin")
	for i := 0; i < b.N; i++ {
		b64.ToBase64String(test)
	}
}
func BenchmarkToBase64Old(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToBase64([]byte("admin:admin"))
	}
}
