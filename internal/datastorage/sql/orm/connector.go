package orm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

const (
	Postgres  = "Postgres"
	Mysql     = "MySql"
	Sqlserver = "Sql Server"
	Sqlite    = "Sqlite"
)

type DataStorageOrm struct {
	dbName string
	db     *gorm.DB
}

func NewDataStorageOrm(dbName string) *DataStorageOrm {
	return &DataStorageOrm{dbName: dbName}
}

func (ds *DataStorageOrm) ConnectToDatabase(p []byte) error {
	var err error
	switch ds.dbName {
	case Postgres:
		ds.db, err = ConnectToPostgres(p)
	case Mysql:
		ds.db, err = ConnectToMysql(p)
	case Sqlite:
		ds.db, err = ConnectToSqlite(p)
	case Sqlserver:
		ds.db, err = ConnectToSqlServer(p)
	default:
		return errors.New("unexpected database")
	}
	return err
}

func ConnectToSqlite(p []byte) (*gorm.DB, error) {
	logging.Logger.Info("orm: Start Connecting to Sqlite")
	var err error
	reader := bytes.NewReader(p)
	param := &datatransferobjects.LocalDSN{}
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&param)
	if err != nil {
		logging.Logger.Error("orm: Decoding failed", "error", err)
		return nil, err
	}
	db, err := gorm.Open(sqlite.Open(filepath.Join(param.Path, param.Name+"?_foreign_keys=1")), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		logging.Logger.Warn("orm: Can`t open database file", "error", err)
		return nil, err
	}
	err = db.AutoMigrate(&Client{}, &Order{}, &ExtraCharge{}, &AdvertisementBlock{},
		&AdvertisementLine{}, &ReleaseDates{}, &CostRate{})
	if err != nil {
		logging.Logger.Error("orm: Migration error", "error", err)
		return nil, err
	}
	return db, nil
}

func ConnectToPostgres(p []byte) (*gorm.DB, error) {
	logging.Logger.Info("orm: Start Connecting to Postgresql")
	var err error
	dsnParam := &datatransferobjects.NetworkDataBaseDSN{}
	reader := bytes.NewReader(p)
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&dsnParam)
	if err != nil {
		logging.Logger.Error("orm: Decoding failed", "error", err)
		return nil, err
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dsnParam.Source, dsnParam.UserName, dsnParam.Password, dsnParam.DataBase, dsnParam.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		logging.Logger.Warn("orm: Connection failed", "error", err)
		return nil, errors.Unwrap(err)
	}
	err = db.AutoMigrate(&Client{}, &Order{}, &ExtraCharge{}, &AdvertisementBlock{},
		&AdvertisementLine{}, &ReleaseDates{}, &CostRate{})
	if err != nil {
		logging.Logger.Error("orm: Migration error", "error", err)
		return nil, err
	}
	return db, nil
}

func ConnectToSqlServer(p []byte) (*gorm.DB, error) {
	logging.Logger.Info("orm: Start Connecting to Sql Server")

	var err error
	dsnParam := &datatransferobjects.NetworkDataBaseDSN{}
	reader := bytes.NewReader(p)
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&dsnParam)
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		dsnParam.UserName, dsnParam.Password, dsnParam.Source, dsnParam.Port, dsnParam.DataBase)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		logging.Logger.Warn("orm: Connection failed", "error", err)
		return nil, err
	}
	err = db.AutoMigrate(&Client{}, &Order{}, &ExtraCharge{}, &AdvertisementBlock{},
		&AdvertisementLine{}, &ReleaseDates{}, &CostRate{})
	if err != nil {
		logging.Logger.Error("orm: Migration error", "error", err)
		return nil, err
	}
	return db, nil
}

func ConnectToMysql(p []byte) (*gorm.DB, error) {
	logging.Logger.Info("orm: Start Connecting to MySql")
	var err error
	dsnParam := &datatransferobjects.NetworkDataBaseDSN{}
	reader := bytes.NewReader(p)
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&dsnParam)
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		dsnParam.UserName, dsnParam.Password, dsnParam.Source, dsnParam.Port, dsnParam.DataBase)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		logging.Logger.Warn("orm: Connection failed", "error", err)
		return nil, err
	}

	err = db.AutoMigrate(&Client{}, &Order{}, &ExtraCharge{}, &AdvertisementBlock{},
		&AdvertisementLine{}, &ReleaseDates{}, &CostRate{})
	if err != nil {
		logging.Logger.Error("orm: Migration error", "error", err)
		return nil, err
	}
	return db, nil
}

func (ds *DataStorageOrm) Close() error {
	logging.Logger.Info("orm: Closing connection")
	sql, err := ds.db.DB()
	if err != nil {
		return err
	}
	return sql.Close()
}

func (ds *DataStorageOrm) ConnectionInfo() map[string]string {
	m := make(map[string]string)
	m["databaseName"] = ds.dbName
	return m
}
