package objectmaker

import (
	"context"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/application"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"

	"github.com/gotk3/gotk3/gtk"
)

type Application interface {
	AdvertisementWindow() application.AdvertisementsWindow
	SelectCostRate(name string)
	ActiveWin() *gtk.Window
	AppendBlockAdvertisement(blockAdv *presenter.BlockAdvertisementDTO)
	AppendClient(client *presenter.ClientDTO)
	AppendCostRate(costRate *presenter.CostRateDTO)
	AppendExtraCharge(charge *presenter.ExtraChargeDTO)
	AppendLineAdvertisement(lineAdv *presenter.LineAdvertisementDTO)
	AppendMessage(msg string)
	AppendOrder(order *presenter.OrderDTO)
	AppendTag(tag *presenter.TagDTO)
	CloseAdvertisementWindow()
	CloseLoginWindow()
	ConnectionStatus() bool
	CreateAddAdvertisementWindow() application.AddAdvertisementWindow
	CreateAddClientWindow() application.AddClientWindow
	CreateAddExtraChargeWindow() application.AddExtraChargeWindow
	CreateAddOrderWindow() application.AddOrderWindow
	CreateAddTagWindow() application.AddTagWindow
	CreateAdvertisementWindow() application.AdvertisementsWindow
	CreateAndShowProgressWindow(message string)
	CreateCostRatesWindow() application.CostRateWindow
	CreateImportDataWindow() application.ImportDataWindow
	CreateLoginWindow() application.LoginWin
	CreateNewCostRatesWindow(updateMode bool) application.NewCostRateWindow
	CreateProgressWindow() application.ProgressWindow
	InitAdvertisement()
	LoadLocalDatabases(dbs []string)
	LoadLoginWin()
	LoadNetworkDatabases(dbs []string)
	LockAll()
	Mode() int8
	NewAdvertisementReportWindow() application.AdvertisementReportWindow
	NewErrorWindow(err error)
	ProgressComplete()
	ProgressCompleteWithError(err error)
	RecieveValue(str string)
	RegisterReciever(entry *gtk.EntryBuffer)
	RemoveBlockAdvertisement(blockAdv string)
	RemoveClient(client string)
	RemoveCostRate(name string)
	RemoveExtraCharge(charge string)
	RemoveLineAdvertisement(lineAdvsID string)
	RemoveOrder(order string)
	RemoveTag(tag string)
	RequestCompleted()
	SetActiveCostRate(activeCostcostRate string)
	SetConnectionStatus(status bool)
	SetMode(mode int8)
	SetReplaceMode(val bool)
	ShowDefaultPort(port string)
	ShowError(err error)
	Start(args []string)
	Stop()
	UnlockAll()
	UnlockAllLoginForm()
	ListStores() application.ListStores
	RefilterLists()
	BlockAllSignals()
	UnblockAllSignals()
	Version() string
}

type RequestHandler interface {
	FileRequests
	Comparer
	Creator
	Updater
	Searcher
	Remover
	Calculator
	CostRateHandler
	ReportHandler
	ConfigHandler
	ImportExporter
	Connector
	DbRequest
	FileStorage
	LockReciever(lock bool)
	IgnoreMessages(lock bool)
}

type FileStorage interface {
	CreateFile(string) error
}

type Creator interface {
	CreateClient(*presenter.ClientDTO) error
	CreateOrder(*presenter.OrderDTO, []presenter.BlockAdvertisementDTO, []presenter.LineAdvertisementDTO) error
	CreateBlockAdvertisement(*presenter.BlockAdvertisementDTO) error
	CreateLineAdvertisement(*presenter.LineAdvertisementDTO) error
	CreateTag(*presenter.TagDTO) error
	CreateExtraCharge(*presenter.ExtraChargeDTO) error
	CreateCostRate(*presenter.CostRateDTO) error
}

type Updater interface {
	UpdateClient(*presenter.ClientDTO) error
	UpdateOrder(*presenter.OrderDTO) error
	UpdateLineAdvertisement(*presenter.LineAdvertisementDTO) error
	UpdateBlockAdvertisement(*presenter.BlockAdvertisementDTO) error
	UpdateTag(*presenter.TagDTO) error
	UpdateExtraCharge(*presenter.ExtraChargeDTO) error
	UpdateCostRate(*presenter.CostRateDTO) error
}

type Searcher interface {
	AllTags()
	AllExtraCharges()
	AllCostRates()
	AllClients()
	AllOrders()
	AllBlockAdvertisements()
	BlockAdvertisementsActual()
	AllLineAdvertisements()
	LineAdvertisementsActual()
}

type Remover interface {
	RemoveClient(string)
	RemoveOrder(*presenter.OrderDTO)
	RemoveLineAdvertisement(*presenter.LineAdvertisementDTO)
	RemoveBlockAdvertisement(*presenter.BlockAdvertisementDTO)
	RemoveTag(string)
	RemoveExtraCharge(string)
	RemoveCostRate(string)
}

type Calculator interface {
	CalculateOrderCost(*presenter.OrderDTO, []presenter.BlockAdvertisementDTO, []presenter.LineAdvertisementDTO) error
	CalculateLineAdvertisementCost(*presenter.LineAdvertisementDTO) error
	CalculateBlockAdvertisementCost(*presenter.BlockAdvertisementDTO) error
}

type ImportExporter interface {
	StartExportingJson(context.Context, string)
	StartImportingJson(context.Context, string, *presenter.ImportParams)
}

type CostRateHandler interface {
	ActiveCostRate()
	SetActiveCostRate(activeCostcostRate string) error
	CheckCostRate()
}

type ReportHandler interface {
	CreateAfvertisementReport(context.Context, *presenter.ReportParams) error
}

type Comparer interface {
	ValueInString(string, string) bool
	CompareStringTime(string, string) int
	ReleasesInTimeRange(string, string, string) bool
	CompareSelected(bool, bool) int
	CheckCostString(string) bool
	CheckAdvCostString(string) bool
	CompareInts(int, int) int
	CompareStrings(string, string) int
	CompareCosts(string, string) int
}

type ConfigHandler interface {
	LoadConfig(string, any) error
	SaveConfig(string, any) error
	RemoveConfig(string) error
}

type DbRequest interface {
	DefaultNetworkPort(string)
	Databases()
}

type FileRequests interface {
	AllFiles(context.Context)
	NewFile(context.Context, *presenter.File) error
	UploadFiles(ctx context.Context, file string)
	FileByName(string) error
	LargeFileByName(string) error
	RemoveFile(string) error
	GetFileURI(string) (string, error)
}

type Connector interface {
	ConnectToServer(*presenter.ServerDSN) error
	ConnectToNetworkDatabase(*presenter.NetworkDataBaseDSN) error
	ConnectToLocalDatabase(*presenter.LocalDSN) error
}

type Tools interface {
	FindValue(id string, list *gtk.ListStore, column int) (*gtk.TreeIter, error)
	RemoveValue(id string, column int, list *gtk.ListStore)
	FindIntValue(int, *gtk.ListStore, int) (*gtk.TreeIter, error)
	RemoveIntValue(id int, column int, list *gtk.ListStore)
	CompareTime(model *gtk.TreeModel, a, b *gtk.TreeIter, col int) int
	CompareStrings(model *gtk.TreeModel, a, b *gtk.TreeIter, col int) int
	CompareBools(model *gtk.TreeModel, a, b *gtk.TreeIter, col int) int
	CompareInts(model *gtk.TreeModel, a, b *gtk.TreeIter, col int) int
	CompareCosts(model *gtk.TreeModel, a, b *gtk.TreeIter, col int) int
	StringFromIter(iter *gtk.TreeIter, model *gtk.TreeModel, col int) (string, error)
	InterfaceFromIter(iter *gtk.TreeIter, model *gtk.TreeModel, col int) (interface{}, error)
	CheckCostString(self *gtk.Entry, new string)
	CheckCostAdvString(self *gtk.Entry, new string)
	CompareNewTime(*gtk.TreeModel, *gtk.TreeIter, string, int) int
	CheckExtraChargeString(self *gtk.Entry, new string)
}

type DataConverter interface {
	ParsePath(path string) string
	ArrayToString(arr []string) string
	SelectedReleaseDatesToString([]string) string
	SelectedTagsToString([]presenter.SelectedTagDTO) string
	SelectedExtraChargeToString([]presenter.SelectedExtraChargeDTO) string
	YearMonthDayToString(uint, uint, uint) string
}
