package logging

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type Logging struct {
	logger        *slog.Logger
	writer        io.Writer
	logFile       io.WriteCloser
	lifeDuraction time.Duration
	prefix        string
	filePattern   string
	pathToWrite   string
}

var (
	Logger Logging
)

func CreateLogger(logsFolder string, lifeDur time.Duration, handlers *slog.HandlerOptions, toFile bool, writers ...io.Writer) {
	Logger = Logging{
		writer:        io.MultiWriter(writers...),
		prefix:        "log_",
		filePattern:   time.Now().Format("02.01.2006"),
		pathToWrite:   logsFolder,
		lifeDuraction: lifeDur,
	}
	if toFile {
		err := Logger.checkConfigFolder()
		if err != nil {
			Logger.writer = os.Stderr
			Logger.logger = slog.New(slog.NewTextHandler(Logger.writer, handlers))
			return
		}
		err = Logger.newFile()
		if err != nil {
			Logger.writer = os.Stderr
			Logger.logger = slog.New(slog.NewTextHandler(Logger.writer, handlers))
			return
		}
		err = Logger.rotateFiles()
		if err != nil {
			Logger.writer = os.Stderr
			Logger.logger = slog.New(slog.NewTextHandler(Logger.writer, handlers))
			return
		}
	}
	Logger.logger = slog.New(slog.NewTextHandler(Logger.writer, handlers))
}

func (l *Logging) checkConfigFolder() error {
	l.pathToWrite = filepath.Join(l.pathToWrite)
	_, err := os.Stat(l.pathToWrite)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = os.Mkdir(filepath.Join(l.pathToWrite), 0750)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func (l *Logging) rotateFiles() error {
	dir := os.DirFS(l.pathToWrite)
	files, err := fs.Glob(dir, l.prefix+"*")
	if err != nil {
		return err
	}
	for _, v := range files {
		st, err := fs.Stat(dir, v)
		if err != nil {
			return err
		}
		if time.Since(st.ModTime()) > l.lifeDuraction {
			err := os.Remove(filepath.Join(l.pathToWrite, v))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (l *Logging) checkAndCreateFile(fileName string) error {

	_, err := os.Stat(filepath.Join(l.pathToWrite, fileName))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			f, err := os.OpenFile(filepath.Join(l.pathToWrite, fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return err
			}
			l.logFile = f
			l.writer = io.MultiWriter(l.writer, f)
			return nil
		}
	}
	return os.ErrExist
}

func (l *Logging) newFile() error {
	err := l.checkAndCreateFile(l.prefix + l.filePattern + ".txt")
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			for i := 1; ; i++ {
				err := l.checkAndCreateFile(l.prefix + l.filePattern + "_" + "(" + strconv.Itoa(i) + ")" + ".txt")
				if err == nil {
					return err
				}

			}
		}
		return err
	}
	return nil
}

func (l *Logging) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l *Logging) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *Logging) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

func (l *Logging) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

func (l *Logging) RawWrite(msg string, data []byte) {
	l.writer.Write([]byte(msg))
	s := bufio.NewScanner(bytes.NewReader(data))
	for s.Scan() {
		_, err := l.writer.Write([]byte(s.Text() + "\n"))
		if err != nil {
			return
		}
	}
}
