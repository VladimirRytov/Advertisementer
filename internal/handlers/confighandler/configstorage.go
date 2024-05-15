package confighandler

import (
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/VladimirRytov/advertisementer/internal/encodedecoder"
	"github.com/VladimirRytov/advertisementer/internal/encryptor"
)

type FileStorage interface {
	OpenForWrite(string) (io.WriteCloser, error)
	OpenForRead(string) (io.ReadSeekCloser, error)
}

type ConfigStorage struct {
	fileGate  FileStorage
	storage   map[string][]byte
	configDir string
}

func NewStorage(fileGate FileStorage) (*ConfigStorage, error) {
	cf := new(ConfigStorage)
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}
	cf.configDir = filepath.Join(configDir, "Advertisementer")

	_, err = os.Stat(cf.configDir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll(filepath.Join(cf.configDir), 0750)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	cf.storage = make(map[string][]byte)
	cf.fileGate = fileGate
	return cf, nil
}

func (cs *ConfigStorage) Load(name string) ([]byte, error) {
	val, ok := cs.storage[name]
	if ok {
		return val, nil
	}

	var b bytes.Buffer
	st, err := cs.fileGate.OpenForRead(filepath.Join(cs.configDir, ".config"))
	if err != nil {
		return nil, err
	}
	defer st.Close()

	err = encryptor.Aes.Decrypt(&b, st)
	if err != nil {
		return nil, err
	}

	err = encodedecoder.FromGob(&cs.storage, &b)
	if err != nil {
		return nil, err
	}

	return cs.storage[name], nil
}

func (cs *ConfigStorage) SaveConfig(name string, data []byte) error {
	var buf bytes.Buffer
	cs.storage[name] = data

	err := encodedecoder.ToGob(&buf, cs.storage)
	if err != nil {
		return err
	}

	err = cs.toFile(&buf)
	return err
}

func (cs *ConfigStorage) Remove(name string) error {
	var buf bytes.Buffer
	delete(cs.storage, name)

	err := encodedecoder.ToGob(&buf, cs.storage)
	if err != nil {
		return err
	}
	err = cs.toFile(&buf)
	return err
}

func (cs *ConfigStorage) toFile(b *bytes.Buffer) error {
	st, err := cs.fileGate.OpenForWrite(filepath.Join(cs.configDir, ".config"))
	if err != nil {
		return err
	}
	defer st.Close()
	err = encryptor.Aes.Ecrypt(st, b)
	if err != nil {
		return err
	}
	return nil
}
