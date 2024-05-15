package logging

import (
	"bytes"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCheckConfigFolder(t *testing.T) {
	log := makeTestLogger()
	err := log.checkConfigFolder()
	if err != nil {
		t.Fatal(err)
	}

}

func TestCheckAndCreateFile(t *testing.T) {
	log := makeTestLogger()
	err := log.checkAndCreateFile(log.filePattern + ".txt")
	if err != nil {
		t.Fatal(err)
	}
	log.logFile.Close()
}

func TestNewFile(t *testing.T) {
	log := makeTestLogger()
	err := log.newFile()
	if err != nil {
		t.Fatal(err)
	}
	log.logFile.Close()
}

func TestRotateFilesRemoveOldFiles(t *testing.T) {
	log := makeTestLogger()
	err := log.checkConfigFolder()
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 20; i++ {
		log.newFile()
		log.logFile.Close()
		time.Sleep(1 * time.Second)
	}
	err = log.rotateFiles()
	if err != nil {
		t.Fatal(err)
	}
	dir := os.DirFS(log.pathToWrite)

	files, err := fs.Glob(dir, "*")
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 9 {
		t.Fatalf("folder must contain 10 files, but got %d files", len(files))
	}
	if err = clean(log.pathToWrite); err != nil {
		t.Fatal(err)
	}
}

func TestWrite(t *testing.T) {
	makeLogger()
	Logger.Debug("test")
	f, err := os.Open(filepath.Join(Logger.pathToWrite, Logger.prefix+Logger.filePattern+".txt"))
	if err != nil {
		t.Fatal(err)
	}
	var b bytes.Buffer
	b.ReadFrom(f)
	if len(b.String()) == 0 {
		t.Fatalf("got empty file")
	}
	Logger.logFile.Close()
	f.Close()
	if err = clean(Logger.pathToWrite); err != nil {
		t.Fatal(err)
	}
}

func makeTestLogger() Logging {
	return Logging{
		filePattern:   time.Now().Format("02.01.2006"),
		pathToWrite:   filepath.Join("Advertisementer"),
		writer:        os.Stderr,
		lifeDuraction: time.Duration(10 * time.Second),
	}
}

func makeLogger() {
	CreateLogger(filepath.Join("Advertisementer"), 14*(24*time.Hour), &slog.HandlerOptions{Level: slog.LevelDebug}, true, os.Stderr)
}

func clean(path string) error {
	return os.RemoveAll(path)
}
