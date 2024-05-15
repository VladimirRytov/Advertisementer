package orm

import (
	"encoding/json"
	"errors"
	"log/slog"
	"os"
	"testing"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func CreateLogger() {
	logging.CreateLogger(".", 0, &slog.HandlerOptions{Level: slog.LevelInfo}, false, os.Stderr)
}

func connectToDBForTests(dbName string) (*DataStorageOrm, error) {
	CreateLogger()
	switch dbName {
	case "Postgres":
		param := &datatransferobjects.NetworkDataBaseDSN{
			Source: "127.0.0.1", DataBase: "gorm_test", UserName: "gorm_test", Password: "gorm_test", SSLMode: false, Port: 5432}
		mar, err := json.Marshal(&param)
		if err != nil {
			return nil, err
		}
		database, err := ConnectToPostgres(mar)
		if err != nil {
			return nil, err
		}
		return &DataStorageOrm{db: database}, nil
	case "Sqlite":
		param := &datatransferobjects.LocalDSN{Name: "testing.sqlite"}

		mar, err := json.Marshal(&param)
		if err != nil {
			return nil, err
		}
		database, err := ConnectToSqlite(mar)
		if err != nil {
			return nil, err
		}
		return &DataStorageOrm{db: database}, nil
	case "Mysql":
		param := &datatransferobjects.NetworkDataBaseDSN{
			Source: "127.0.0.1", DataBase: "gorm_test", UserName: "gorm_test", Password: "gorm_test", SSLMode: false, Port: 3306}
		mar, err := json.Marshal(&param)
		if err != nil {
			return nil, err
		}
		database, err := ConnectToMysql(mar)
		if err != nil {
			return nil, err
		}
		return &DataStorageOrm{db: database}, nil
	case "Sql Server":
		param := &datatransferobjects.NetworkDataBaseDSN{
			Source: "127.0.0.1", DataBase: "master", UserName: "sa", Password: "Gorm_Test", SSLMode: false, Port: 1433}
		mar, err := json.Marshal(&param)
		if err != nil {
			return nil, err
		}
		database, err := ConnectToSqlServer(mar)
		if err != nil {
			return nil, err
		}
		return &DataStorageOrm{db: database}, nil
	}
	return nil, errors.New("unexpected error")
}

func TestConnectToSqlite(t *testing.T) {
	CreateLogger()
	param := &datatransferobjects.LocalDSN{Name: ":memory:"}
	mar, err := json.Marshal(&param)
	if err != nil {
		t.Error(err)
	}
	var db *DataStorageOrm = new(DataStorageOrm)
	db.db, err = ConnectToSqlite(mar)
	if err != nil {
		t.Fatal(err)
	}
	db.Close()
	_, err = ConnectToSqlite([]byte("asdas"))
	if err != nil {
		switch err.(type) {
		case *json.SyntaxError:
			t.Logf("got syntax error")
		default:
			t.Fatalf("unexpected error - %T", err)
		}
	} else {
		t.Fatalf("want error")
	}
	param = &datatransferobjects.LocalDSN{Path: "/"}
	mar, err = json.Marshal(&param)
	if err != nil {
		t.Error(err)
	}
	_, err = ConnectToSqlite(mar)
	if err == nil {
		t.Fatal("want error")
	}
}
func TestConnectToPostgres(t *testing.T) {
	CreateLogger()
	param := &datatransferobjects.NetworkDataBaseDSN{
		Source: "127.0.0.1", DataBase: "gorm_test", UserName: "gorm_test", Password: "gorm_test", SSLMode: false, Port: 5432}

	mar, err := json.Marshal(&param)
	if err != nil {
		t.Error(err)
	}
	var db *DataStorageOrm = new(DataStorageOrm)
	db.db, err = ConnectToPostgres(mar)
	if err != nil {
		t.Error(err)
	}
	db.Close()
	_, err = ConnectToPostgres([]byte("asdas"))
	if err != nil {
		switch err.(type) {
		case *json.SyntaxError:
			t.Logf("got syntax error")
		default:
			t.Fatalf("unexpected error - %T", err)
		}
	} else {
		t.Fatalf("want error")
	}

	param = &datatransferobjects.NetworkDataBaseDSN{
		Source: "127.0.0.1", DataBase: "gorn_test", UserName: "gorn_test", Password: "gorn_test", SSLMode: false, Port: 33333}
	mar, err = json.Marshal(&param)
	if err != nil {
		t.Error(err)
	}
	_, err = ConnectToPostgres(mar)
	if err == nil {
		t.Fatal("want error")
	}
}
func TestConnectToMysql(t *testing.T) {
	CreateLogger()
	param := &datatransferobjects.NetworkDataBaseDSN{
		Source: "127.0.0.1", DataBase: "gorm_test", UserName: "gorm_test", Password: "gorm_test", SSLMode: false, Port: 3306}

	mar, err := json.Marshal(&param)
	if err != nil {
		t.Error(err)
	}
	var db *DataStorageOrm = new(DataStorageOrm)
	db.db, err = ConnectToMysql(mar)
	if err != nil {
		t.Error(err)
	}
	db.Close()
	_, err = ConnectToMysql([]byte("asdas"))
	if err != nil {
		switch err.(type) {
		case *json.SyntaxError:
			t.Logf("got syntax error")
		default:
			t.Fatalf("unexpected error - %T", err)
		}
	} else {
		t.Fatalf("want error")
	}

	param = &datatransferobjects.NetworkDataBaseDSN{
		Source: "127.0.0.1", DataBase: "gorn_test", UserName: "gorn_test", Password: "gorn_test", SSLMode: false, Port: 33333}
	mar, err = json.Marshal(&param)
	if err != nil {
		t.Error(err)
	}
	_, err = ConnectToMysql(mar)
	if err == nil {
		t.Fatal("want error")
	}
}

func TestConnectToSQLServer(t *testing.T) {
	CreateLogger()
	param := &datatransferobjects.NetworkDataBaseDSN{
		Source: "127.0.0.1", DataBase: "master", UserName: "sa", Password: "Gorm_Test", SSLMode: false, Port: 1433}

	mar, err := json.Marshal(&param)
	if err != nil {
		t.Error(err)
	}
	var db *DataStorageOrm = new(DataStorageOrm)
	db.db, err = ConnectToSqlServer(mar)
	if err != nil {
		t.Fatal(err)
	}
	db.Close()
	_, err = ConnectToSqlServer([]byte("asdas"))
	if err != nil {
		switch err.(type) {
		case *json.SyntaxError:
			t.Logf("got syntax error")
		default:
			t.Fatalf("unexpected error - %T", err)
		}
	} else {
		t.Fatalf("want error")
	}

	param = &datatransferobjects.NetworkDataBaseDSN{
		Source: "127.0.0.1", DataBase: "gorn_test", UserName: "gorn_test", Password: "gorn_test", SSLMode: false, Port: 33333}
	mar, err = json.Marshal(&param)
	if err != nil {
		t.Error(err)
	}
	_, err = ConnectToSqlServer(mar)
	if err == nil {
		t.Fatal("want error")
	}
}
