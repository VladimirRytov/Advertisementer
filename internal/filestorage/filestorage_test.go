package filestorage

import (
	"bytes"
	"errors"
	"log/slog"
	"os"
	"slices"
	"testing"

	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func TestOpenForWrite(t *testing.T) {
	CreateLogger()
	fileName := "test"
	Files := new(Storage)
	f, err := Files.OpenForWrite(fileName)
	if err != nil {
		t.Fatal(err)
	}
	f.Write([]byte("Hello"))
	f.Write([]byte("World"))
	f.Close()
	os.Remove(fileName)
}

func TestOpenForRead(t *testing.T) {
	CreateLogger()
	fileName := "test"
	Files := new(Storage)

	writeFile, err := Files.OpenForWrite(fileName)
	if err != nil {
		t.Fatal(err)
	}
	writeFile.Write([]byte("Hello"))
	writeFile.Write([]byte("World"))
	writeFile.Close()

	f, err := Files.OpenForRead(fileName)
	if err != nil {
		t.Fatal(err)
	}
	testCase := []byte("HelloWorld")
	var b bytes.Buffer
	_, err = b.ReadFrom(f)
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(b.Bytes(), testCase) {
		t.Fatalf("want %v,got %v", testCase, b)
	}
	f.Close()
	os.Remove(fileName)
}

func TestNew(t *testing.T) {
	CreateLogger()
	fileName := "test"
	Files := new(Storage)
	f, err := Files.New(fileName)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	_, err = Files.New(fileName)
	if err != nil && !errors.Is(err, errExist) {
		t.Fatal(err)
	}
	os.Remove(fileName)
}

func CreateLogger() {
	logging.CreateLogger(".", 0, &slog.HandlerOptions{Level: slog.LevelInfo}, false, os.Stderr)
}
