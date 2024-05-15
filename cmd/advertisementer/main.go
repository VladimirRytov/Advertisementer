package main

import (
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"runtime/debug"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datastorage/server"
	"github.com/VladimirRytov/advertisementer/internal/datastorage/sql/orm"
	"github.com/VladimirRytov/advertisementer/internal/encodedecoder"
	"github.com/VladimirRytov/advertisementer/internal/encryptor"
	"github.com/VladimirRytov/advertisementer/internal/filestorage"
	"github.com/VladimirRytov/advertisementer/internal/front/httpsender"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/advertisementwindow/tools"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/application"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/application/objectmaker"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter/requestscontroller"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter/windowscontroller"
	"github.com/VladimirRytov/advertisementer/internal/handlers"
	"github.com/VladimirRytov/advertisementer/internal/handlers/advreport"
	"github.com/VladimirRytov/advertisementer/internal/handlers/confighandler"
	"github.com/VladimirRytov/advertisementer/internal/handlers/costcalculationhandler"
	"github.com/VladimirRytov/advertisementer/internal/handlers/thickclienthandler"
	"github.com/VladimirRytov/advertisementer/internal/handlers/thinclienthandler"
	"github.com/VladimirRytov/advertisementer/internal/handlers/thincostratecalculator"
	"github.com/VladimirRytov/advertisementer/internal/logging"
	"github.com/VladimirRytov/advertisementer/internal/reciever"
	"github.com/VladimirRytov/advertisementer/internal/reciever/wsreciever"

	"github.com/VladimirRytov/advertisementer/internal/datastorage/dbmanager"
)

var version = ""

func main() {
	defer func() {
		if p := recover(); p != nil {
			logging.Logger.Error("got panic", "message", p)
			logging.Logger.RawWrite("stack", debug.Stack())
		}
	}()
	var err error
	cd, err := os.UserConfigDir()
	if err != nil {
		log.Println("у пользователя отсуствует папка для конфигураций")
	}
	logging.CreateLogger(filepath.Join(cd, "Advertisementer"), 14*(24*time.Hour), &slog.HandlerOptions{Level: slog.LevelError}, true, os.Stderr)
	//

	b64Enc := encodedecoder.NewBase64Encoder()
	encryptor.Aes, err = encryptor.AesInit()
	if err != nil {
		panic(err)
	}
	//
	dbm := dbmanager.NewDatabaseManager()
	dbm.Register(orm.Sqlite, 0, dbmanager.LocalDB, orm.NewDataStorageOrm(orm.Sqlite))
	dbm.Register(orm.Postgres, 5432, dbmanager.NetworkDB, orm.NewDataStorageOrm(orm.Postgres))
	dbm.Register(orm.Mysql, 3306, dbmanager.NetworkDB, orm.NewDataStorageOrm(orm.Mysql))
	dbm.Register(orm.Sqlserver, 1433, dbmanager.NetworkDB, orm.NewDataStorageOrm(orm.Sqlserver))
	dbm.RegisterServerGateway(8080, dbmanager.Server, server.NewServerStorage(httpsender.NewSender(), b64Enc))
	//
	fileStorage := &filestorage.Storage{}
	ConfigsAccesor, err := confighandler.NewStorage(fileStorage)
	if err != nil {
		panic(err)
	}
	dataConverter := presenter.NewDataConverter()
	requestGate := requestscontroller.NewRequestsHandler(ConfigsAccesor, dataConverter, b64Enc)
	requestGate.SetFileStorage(fileStorage)
	tools := tools.NewTools(requestGate)
	//
	objectMaker := objectmaker.NewObjectMaker(requestGate, dataConverter, tools)
	app := application.CreateApplication(objectMaker, requestGate, tools, version)
	//
	winController := windowscontroller.NewWindowsController(app, dataConverter)
	recHandler := reciever.NewRecieveHandler(winController)
	recHandler.RegisterReciever("webSock", wsreciever.New(recHandler, b64Enc))

	thickMode := thickclienthandler.NewAdvertisementController(winController)
	thinMode := thinclienthandler.NewAdvertisementController(winController, fileStorage)
	handler := handlers.NewDatabaseController(dbm, winController, recHandler, thickMode, thinMode, requestGate, thincostratecalculator.NewCostRateCalculator(),
		costcalculationhandler.NewCostRateCalculator(), advreport.NewReportGenerator())
	requestGate.SetDatabaseGateway(handler)

	//
	app.Start(os.Args)
}
