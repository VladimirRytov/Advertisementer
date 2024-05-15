package confighandler

import (
	"os"
	"slices"
	"testing"

	"github.com/VladimirRytov/advertisementer/internal/encryptor"
	"github.com/VladimirRytov/advertisementer/internal/filestorage"
)

func TestSaveConfig(t *testing.T) {
	encryptor.Aes, _ = encryptor.AesInit()
	dataToSave, err := NewStorage(&filestorage.Storage{})
	if err != nil {
		t.Fatalf("cannot create storage, error %v", err)
	}
	err = dataToSave.SaveConfig("Test", []byte("hello"))
	if err != nil {
		t.Fatalf("cannot save config, error %v", err)
	}
}

func TestLoadConfig(t *testing.T) {
	encryptor.Aes, _ = encryptor.AesInit()
	dataToSave, err := NewStorage(&filestorage.Storage{})
	if err != nil {
		t.Fatalf("cannot create storage, error %v", err)
	}
	data, err := dataToSave.Load("Test")
	if err != nil {
		t.Fatalf("cannot save config, error %v", err)
	}
	if slices.Compare(data, []byte("hello")) != 0 {
		t.Fatalf("want %v, got %v", []byte("hello"), data)

	}
	defer os.Remove(".config")
}
