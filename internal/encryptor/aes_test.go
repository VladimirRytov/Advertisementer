package encryptor

import (
	"bytes"
	"crypto/cipher"
	"io"
	"slices"
	"testing"
)

func TestEncrypt(t *testing.T) {
	aes, _ := AesInit()
	testCase := []byte("hello world")
	r := bytes.NewReader(testCase)
	var got bytes.Buffer
	aes.Ecrypt(&got, r)

	stream := cipher.NewOFB(aes.key, aes.iv[:])
	var dec bytes.Buffer

	writer := &cipher.StreamWriter{S: stream, W: &dec}
	io.Copy(writer, &got)

	if slices.Compare(testCase, dec.Bytes()) != 0 {
		t.Fatalf("want %v,got %v", testCase, dec.Bytes())
	}
}

func TestDencrypt(t *testing.T) {
	aes, _ := AesInit()
	want := []byte("hello world")
	testCase := bytes.NewBuffer(want)

	var test bytes.Buffer
	aes.Ecrypt(&test, testCase)

	var got bytes.Buffer
	aes.Decrypt(&got, &test)

	if slices.Compare(got.Bytes(), want) != 0 {
		t.Fatalf("want %s,got %s", got.Bytes(), want)
	}
}
func BenchmarkEncrypt(b *testing.B) {
	b.StopTimer()
	aes, _ := AesInit()
	var by bytes.Buffer
	var testCases []*bytes.Buffer = make([]*bytes.Buffer, b.N)
	for i := 0; i < b.N; i++ {
		testCases[i] = bytes.NewBuffer([]byte("helloWorld"))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		aes.Ecrypt(&by, testCases[i])
	}
}

func BenchmarkDecrypt(b *testing.B) {
	b.StopTimer()
	aes, _ := AesInit()
	var by bytes.Buffer
	var testCases []*bytes.Buffer = make([]*bytes.Buffer, b.N)
	for i := 0; i < b.N; i++ {
		var got bytes.Buffer
		r := bytes.NewReader([]byte("helloWorld"))
		aes.Ecrypt(&got, r)
		testCases[i] = bytes.NewBuffer(got.Bytes())
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		aes.Decrypt(&by, testCases[i])
	}
}
