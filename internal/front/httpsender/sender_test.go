package httpsender

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/encodedecoder"
)

func TestAuthenticate(t *testing.T) {
	dsn := datatransferobjects.ServerDSN{
		Source:   "127.0.0.1",
		UserName: "admin",
		Password: "admin",
		Port:     8080,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s := NewSender()
	err := s.Initialize(ctx, dsn)
	if err == nil {
		t.Fatal(err)
	}
}
func TestGetRequest(t *testing.T) {
	dsn := datatransferobjects.ServerDSN{
		Source:   "127.0.0.1",
		UserName: "admin",
		Password: "admin",
		Port:     8080,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s := NewSender()
	err := s.Initialize(ctx, dsn)
	if err != nil {
		t.Fatal(err)
	}
	got, err := s.GetRequest(ctx, &datatransferobjects.RequestDTO{Kind: "clients", Name: "", Queries: nil}, nil)
	if err != nil {
		t.Fatalf("error: %v\n body: %s", err, string(got))
	}
}

func TestPostRequest(t *testing.T) {
	dsn := datatransferobjects.ServerDSN{
		Source:   "127.0.0.1",
		UserName: "admin",
		Password: "admin",
		Port:     8080,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s := NewSender()
	err := s.Initialize(ctx, dsn)
	if err != nil {
		t.Fatal(err)
	}
	cli := &datatransferobjects.ClientDTO{
		Name:                  "Артур",
		Phones:                "88005553535",
		Email:                 "asdada@asdasd.com",
		AdditionalInformation: "zxczxc",
	}
	var b bytes.Buffer
	err = encodedecoder.ToJSON(&b, &cli, false)
	if err != nil {
		t.Fatal(err)
	}
	got, err := s.CreateRequest(ctx, &datatransferobjects.RequestDTO{Kind: "clients", Name: "", Queries: nil}, &b)
	if err != nil {
		t.Fatalf("error: %v\n body: %s", err, string(got))
	}
}

func TestPutRequest(t *testing.T) {
	dsn := datatransferobjects.ServerDSN{
		Source:   "127.0.0.1",
		UserName: "admin",
		Password: "admin",
		Port:     8080,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s := NewSender()
	err := s.Initialize(ctx, dsn)
	if err != nil {
		t.Fatal(err)
	}
	cli := &datatransferobjects.ClientDTO{
		Name:                  "Артур",
		Phones:                "88005553535",
		Email:                 "asdada@asdasd.com",
		AdditionalInformation: "обновлён",
	}
	var b bytes.Buffer
	err = encodedecoder.ToJSON(&b, &cli, false)
	if err != nil {
		t.Fatal(err)
	}
	got, err := s.UpdateRequest(ctx, &datatransferobjects.RequestDTO{Kind: "clients", Name: cli.Name, Queries: nil}, &b)
	if err != nil {
		t.Fatalf("error: %v\n body: %s", err, string(got))
	}
}

func TestDeleteRequest(t *testing.T) {
	dsn := datatransferobjects.ServerDSN{
		Source:   "127.0.0.1",
		UserName: "admin",
		Password: "admin",
		Port:     8080,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s := NewSender()
	err := s.Initialize(ctx, dsn)
	if err != nil {
		t.Fatal(err)
	}
	cli := &datatransferobjects.ClientDTO{
		Name:                  "Артур",
		Phones:                "88005553535",
		Email:                 "asdada@asdasd.com",
		AdditionalInformation: "обновлён",
	}
	var b bytes.Buffer
	err = encodedecoder.ToJSON(&b, &cli, false)
	if err != nil {
		t.Fatal(err)
	}
	got, err := s.DeleteRequest(ctx, &datatransferobjects.RequestDTO{Kind: "clients", Name: cli.Name, Queries: nil})
	if err != nil {
		t.Fatalf("error: %v\n body: %s", err, string(got))
	}
}
