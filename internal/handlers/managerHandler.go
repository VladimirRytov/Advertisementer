package handlers

import (
	"context"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/filestorage"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/handlers/advreport/exelexporter"
	"github.com/VladimirRytov/advertisementer/internal/handlers/exporthandler"
	"github.com/VladimirRytov/advertisementer/internal/handlers/importhandler"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

const (
	ThickClientMode = "thickClient"
	ThinClientMode  = "thinClient"
)

type DatabaseRequester interface {
	SetRequestsGateway(req HandlerRequests)
	SetReciever(rec Reciever)
	SetFileStorage(storage FileStorage)
	SetFileManager(filesReq ServerRequests)
	SetSecondaryCompopnents(jsExp JSONExporter, jsImp JSONImporter,
		repGen Reports, costCalc CostRateCalculator)
}

type FilesRequests interface {
	NewFile(context.Context, *datatransferobjects.FileDTO) error
	AllFiles()
	FileByName(string) error
	RemoveFile(string) error
}

type DbManagerHandler struct {
	connected  bool
	app        Responcer
	requests   DatabaseRequester
	dbHandler  DataBaseConnector
	subHandler SubHandler

	thickCostCalculator ThickCostRateCalculator
	thickMode           ThickAdvertisementHandler
	thinMode            ThinAdvertisementHandler
	thinCostCalculator  ThinCostRateCalculator
	repGenerator        ReportHandler
}

func NewDatabaseController(dbm DataBaseConnector, app Responcer,
	sub SubHandler, thick ThickAdvertisementHandler, thin ThinAdvertisementHandler,
	reqGateway DatabaseRequester, thinCostCalc ThinCostRateCalculator,
	thickCostCalc ThickCostRateCalculator, repGen ReportHandler) *DbManagerHandler {

	logging.Logger.Debug("dbManagerHandler: Initialize Database Controller")
	return &DbManagerHandler{
		dbHandler: dbm, app: app, subHandler: sub,
		thickMode: thick, thinMode: thin, requests: reqGateway,
		thickCostCalculator: thickCostCalc, thinCostCalculator: thinCostCalc, repGenerator: repGen}
}

func (dbm *DbManagerHandler) AvailableLocalDatabases() {
	logging.Logger.Debug("dbManagerHandler: Requesting available local databases")
	dbm.app.SendLocalDatabases(
		dbm.dbHandler.AvailableLocalDatabases())
}

func (dbm *DbManagerHandler) AvailableNetworkDatabases() {
	logging.Logger.Debug("dbManagerHandler: Requesting available network databases")
	dbm.app.SendNetworkDatabases(
		dbm.dbHandler.AvailableNetworkDatabases())
}

func (dbm *DbManagerHandler) DefaultPort(dbName string) {
	logging.Logger.Debug("dbManagerHandler: Requesting default network port for database", "database", dbName)
	dbm.app.SendDefaultPort(dbm.dbHandler.DefaultPort(dbName))
}

func (dbm *DbManagerHandler) DatabaseConnected() bool {
	return dbm.connected
}

func (dbm *DbManagerHandler) ConnectToDatabase(dbName string, params []byte) {
	logging.Logger.Debug("dbManagerHandler: Requesting connect to database", "dbName", dbName)
	db, err := dbm.dbHandler.ConnectToDatabase(dbName, params)
	if err != nil {
		dbm.app.SendConnectionError(err)
		return
	}
	files := new(filestorage.Storage)
	dbm.initReportGenerator(db, files, files)
	dbm.thickCostCalculator.InitCostRateCalculator(dbm.app, db)
	dbm.initSecondaryComponents(db, dbm.repGenerator, files, dbm.thickCostCalculator)
	dbm.thickMode.InitAdvertisementController(db)
	dbm.requests.SetRequestsGateway(dbm.thickMode)
	dbm.app.SendSuccesConnection()
}

func (dbm *DbManagerHandler) ConnectToServer(connType string, params []byte) {
	logging.Logger.Debug("dbManagerHandler: connecting to server", "connType", connType)
	db, err := dbm.dbHandler.ConnectToServer(params)
	if err != nil {
		dbm.app.SendConnectionError(err)
		return
	}
	conInfo := db.ConnectionInfo()
	logging.Logger.Debug("dbManagerHandler: subscriging to", "addr", conInfo["adress"], "apiPath", conInfo["apiPath"])

	err = dbm.subHandler.Subscribe(conInfo["token"], conInfo["adress"], conInfo["apiPath"])
	if err != nil {
		logging.Logger.Error("dbManagerHandler: cannot subscribe", "error", err)
	}
	files := new(filestorage.Storage)
	dbm.initReportGenerator(db, db, files)
	dbm.thinCostCalculator.InitCostRateCalculator(dbm.app, db)
	dbm.initSecondaryComponents(db, dbm.repGenerator, files, dbm.thinCostCalculator)
	logging.Logger.Debug("dbManagerHandler: connecting to server", "connType", connType)
	dbm.thinMode.InitAdvertisementController(db)
	dbm.requests.SetRequestsGateway(dbm.thinMode)
	dbm.requests.SetFileManager(dbm.thinMode)
	dbm.requests.SetReciever(dbm.subHandler)
	dbm.connected = true
	dbm.app.SendSuccesConnection()
}

func (dbm *DbManagerHandler) initSecondaryComponents(requests DataBase, reportHandler Reports, fileStorage FileStorage, costRate CostRateCalculator) {
	Export := exporthandler.NewExportHandler(dbm.app, requests, fileStorage)
	Import := importhandler.CreateImporter(dbm.app, presenter.NewDataConverter(), fileStorage, requests)

	dbm.requests.SetSecondaryCompopnents(Export, Import, reportHandler, costRate)
}

func (dbm *DbManagerHandler) initReportGenerator(requests DataBase, file AdvFilesGateway, fileStorage FileStorage) {
	repGen := exelexporter.NewExelReporter(fileStorage, file)
	dbm.repGenerator.InitReportGenerator(dbm.app, requests)
	dbm.repGenerator.RegisterGenerator("Exel", repGen)
}
