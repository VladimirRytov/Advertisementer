package requestscontroller

import (
	"encoding/json"
	"log/slog"
	"os"
	"testing"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func CreateLogger() {
	logging.CreateLogger(".", 0, &slog.HandlerOptions{Level: slog.LevelInfo}, false, os.Stderr)
}

func BenchmarkDecode(b *testing.B) {
	dc := NewRequestsHandler(nil, nil, nil)
	data := &presenter.NetworkDataBaseDSN{
		DatabaseName: "Mysql",
		Source:       "127.0.0.1",
		DataBase:     "test",
		UserName:     "tester",
		Password:     "1234",
		SSLMode:      false,
		Port:         "11223",
	}
	raw, err := json.Marshal(data)
	if err != nil {
		b.Fatalf("got error during encoding: %v", err)
	}
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	got := &presenter.NetworkDataBaseDSN{}
	for i := 0; i < b.N; i++ {
		err = dc.decodeData(raw, got)
		if err != nil {
			b.Fatalf("got error during Decoding: %v", err)
		}
	}
}

func BenchmarkEncodeData(b *testing.B) {
	dc := NewRequestsHandler(nil, nil, nil)
	data := &presenter.NetworkDataBaseDSN{
		DatabaseName: "Mysql",
		Source:       "127.0.0.1",
		DataBase:     "test",
		UserName:     "tester",
		Password:     "1234",
		SSLMode:      false,
		Port:         "11223",
	}
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, err := dc.encodeData(data)
		if err != nil {
			b.Fatalf("got error during encoding: %v", err)
		}
	}
}

func TestEncode(t *testing.T) {
	dc := NewRequestsHandler(nil, nil, nil)
	data := &presenter.NetworkDataBaseDSN{
		DatabaseName: "Mysql",
		Source:       "127.0.0.1",
		DataBase:     "test",
		UserName:     "tester",
		Password:     "1234",
		SSLMode:      false,
		Port:         "11223",
	}
	rawData, err := dc.encodeData(data)
	if err != nil {
		t.Fatalf("got error during encoding: %v", err)
	}
	got := &presenter.NetworkDataBaseDSN{}
	err = json.Unmarshal(rawData, &got)
	if err != nil {
		t.Fatalf("got error during unmarshal: %v", err)
	}
	if *data != *got {
		t.Fatalf("got %v, want %v", got, data)
	}
}

func TestDecodeData(t *testing.T) {
	dc := NewRequestsHandler(nil, nil, nil)
	data := &presenter.NetworkDataBaseDSN{
		DatabaseName: "Mysql",
		Source:       "127.0.0.1",
		DataBase:     "test",
		UserName:     "tester",
		Password:     "1234",
		SSLMode:      false,
		Port:         "11223",
	}
	raw, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("got error during encoding: %v", err)
	}
	got := &presenter.NetworkDataBaseDSN{}
	err = dc.decodeData(raw, got)
	if err != nil {
		t.Fatalf("got error during Decoding: %v", err)
	}
	if *data != *got {
		t.Fatalf("got %v, want %v", got, data)
	}
}
