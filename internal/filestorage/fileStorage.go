package filestorage

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

var (
	errIsDir = errors.New("файл является директорией")
	errExist = errors.New("файл существует")
)

func ErrIsDir() error { return errIsDir }

func ErrExist() error { return errExist }

type Storage struct{}

func (s *Storage) OpenForWrite(name string) (io.WriteCloser, error) {
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (s *Storage) New(name string) (io.WriteCloser, error) {
	stat, err := os.Stat(name)
	switch {
	case errors.Is(err, os.ErrNotExist):
		return s.OpenForWrite(name)

	case stat.IsDir():
		return nil, errors.Join(errors.New("выбранный путь указывает на директорию"), errIsDir)
	default:
		return nil, errors.Join(errors.New("выбранный файл уже существует"), errExist)
	}
}

func (s *Storage) OpenForRead(name string) (io.ReadSeekCloser, error) {
	f, err := os.OpenFile(name, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (s *Storage) CopyFile(source, destination string) error {
	logging.Logger.Debug("copyFile: start copying file")
	sourceFile, err := s.OpenForRead(source)
	if err != nil {
		logging.Logger.Error("storage.CopyFile: cannot open source file", "error", err)
		return err
	}
	defer sourceFile.Close()

	destFile, err := s.OpenForWrite(destination)
	if err != nil {
		logging.Logger.Error("storage.CopyFile: cannot open destination file", "error", err)
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		logging.Logger.Error("storage.CopyFile: cannot copy file", "error", err)
		return err
	}
	return nil
}

func (s *Storage) FileByName(ctx context.Context, name string) (datatransferobjects.FileDTO, error) {
	select {
	case <-ctx.Done():
		return datatransferobjects.FileDTO{}, ctx.Err()
	default:
		file, err := os.Open(name)
		if err != nil {
			return datatransferobjects.FileDTO{}, err
		}
		defer file.Close()

		fileStat, err := file.Stat()
		if err != nil {
			return datatransferobjects.FileDTO{}, err
		}

		buf := bytes.NewBuffer(make([]byte, 0, fileStat.Size()))
		written, err := io.Copy(buf, file)
		switch {
		case err != nil:
			return datatransferobjects.FileDTO{}, err
		case written != fileStat.Size():
			return datatransferobjects.FileDTO{}, errors.New("при копировании фалйа возникла ошибка")
		}

		return datatransferobjects.FileDTO{
			Name: file.Name(),
			Size: fileStat.Size(),
			Data: buf.Bytes(),
		}, nil

	}
}
