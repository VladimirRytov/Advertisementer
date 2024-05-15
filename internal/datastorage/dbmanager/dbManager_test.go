package dbmanager

import (
	"log/slog"
	"os"
	"reflect"
	"testing"

	"github.com/VladimirRytov/advertisementer/internal/datastorage/sql/orm"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

var connectedDbName string

func CreateLogger() {
	logging.CreateLogger(".", 0, &slog.HandlerOptions{Level: slog.LevelInfo}, false, os.Stderr)
}

func TestAvailableDatabases(t *testing.T) {
	CreateLogger()
	wantNetwork := []string{"Mysql", "Postgres"}
	wantLocal := []string{"sqlite"}
	db := NewDatabaseManager()
	for i := range wantLocal {
		db.Register(wantLocal[i], 1111, LocalDB, nil)
	}
	got := db.AvailableLocalDatabases()
	if !reflect.DeepEqual(wantLocal, got) {
		t.Errorf("want %v, got %v", wantLocal, got)
	}
	for i := range wantNetwork {
		db.Register(wantNetwork[i], 1111, NetworkDB, nil)
	}
	got = db.AvailableNetworkDatabases()
	if !reflect.DeepEqual(wantNetwork, got) {
		t.Errorf("want %v, got %v", wantNetwork, got)
	}
}

func TestConnectToDataBase(t *testing.T) {
	CreateLogger()
	dbm := NewDatabaseManager()
	dbm.Register("sqlite", 0, LocalDB, orm.NewDataStorageOrm("sqlite"))

	dbm.Register("Postgres", 5432, LocalDB, orm.NewDataStorageOrm("Postgres"))
	dbm.Register("Mysql", 3306, LocalDB, orm.NewDataStorageOrm("Mysql"))

	localDB := []string{"sqlite"}
	networkDB := []string{"Postgres", "Mysql"}
	for _, v := range localDB {
		dbm.ConnectToDatabase(v, []byte(v))
		if connectedDbName != v {
			t.Fatal("Connect to wrong DB")
		}
	}
	for _, v := range networkDB {
		dbm.ConnectToDatabase(v, []byte(v))
		if connectedDbName != v {
			t.Fatal("Connect to wrong DB")
		}
	}
	_, err := dbm.ConnectToDatabase("asdasd", []byte("asd"))
	if err.Error() != "error: this database not supported" {
		t.Fatalf("want error \"error: this database not supported\",got %v", err)
	}
}
