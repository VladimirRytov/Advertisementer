package application

import (
	"context"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type Window interface {
	Window() *gtk.Window
	Close()
	Show()
}

type LoginWin interface {
	Window
	LoadNetworkDatabases([]string)
	LoadLocalDatabases([]string)
	ShowDefaultPort(string)
	ShowError(error)
	LockAll()
	UnlockAll()
	SaveConfigs()
	LoadConfigs()
}

type ErrorWindow interface {
	Window
	SetErrorMessage(string)
}

type ProgressWindow interface {
	Window
	SetMessage(string)
	SetCancelFunc(context.CancelFunc)
	ShowTreeview(bool)
	BlockCalcelButton(bool)
	SetCloseButtonMessage(string)
	ProgressDone(string)
	SetBeforeFunc(func())
	SetAfterFunc(func())
}

type AdvertisementsWindow interface {
	Window
	ActualSelected() bool
	SetConnectionStatus(bool)
	Notebook
	ShowPopover(string, gtk.IWidget)
	StartInitialization()
	UpdateLists()
	FromDate() string
	ToDate() string
	Update()
	AttachAll()
	SelectCostRate(string)
	RefilterLists()
}

type Notebook interface {
	TabsLocker
	UpdateButtonPressed()
	AttachAll()
	SetSensetive(bool)
	CurrentPageChanged()
	BlockAllSignals()
	UnblockAllSignals()
	EnableAllModels()
	DisableAllModels()
	ClearAllListStores()
	EnableSidebarsFilters(bool)
	EnableAdvertisementFilters(bool)
	ResetSorts()
	RefilterBlock()
	RefilterLine()
}

type Tabs interface {
	SelectedRowChanged()
	SetSensitive(bool)
	SetEnableFilters(bool)
	AttachList()
	ResetSort()
	SetModel()
	UnsetModel()
	BlockSignals()
	UnblockSignals()
}

type TabsLocker interface {
	LockClientsTab(bool)
	LockOrdersTab(bool)
	LockBlockAdvertisementsTab(bool)
	LockLineAdvertisementsTab(bool)
	LockTagsTab(bool)
	LockExtraChargesTab(bool)
}

type TabRefilter interface {
	Refilter()
}

type ClientTab interface {
	Tabs
}

type OrderTab interface {
	Tabs
}

type BlockTab interface {
	Tabs
	TabRefilter
}

type LineTab interface {
	Tabs
	TabRefilter
}

type TagTab interface {
	Tabs
}

type ChargeTab interface {
	Tabs
}

type AddTagWindow interface {
	Window
	TagName() string
	TagCost() string
}

type AddExtraChargeWindow interface {
	Window
	ChargeName() string
	Multiplier() string
}

type AddClientWindow interface {
	Window
	Name() string
	Phone() string
	Email() string
	AdditionalInformation() string
}

type AddAdvertisementWindow interface {
	Window() *gtk.Window
	Show()
	Close()
}

type AddOrderWindow interface {
	Window
}

type ImportDataWindow interface {
	Window
	SetFilePath(path string)
}

type ListStores interface {
	TagsList() *gtk.ListStore
	TagsListCopy() (*gtk.ListStore, error)
	AppendTag(presenter.TagDTO)
	RemoveTag(string)
	ClearTagList()

	ExtraChargesList() *gtk.ListStore
	ExtraChargeListCopy() (*gtk.ListStore, error)
	AppendExtraCharge(presenter.ExtraChargeDTO)
	RemoveExtraCharge(string)
	ClearExtraChargeList()

	ClientsList() *gtk.ListStore
	ClientListCopy() (*gtk.ListStore, error)
	AppendClient(presenter.ClientDTO)
	RemoveClient(string)
	ClearClientList()

	OrdersList() *gtk.ListStore
	OrderListCopy() (*gtk.ListStore, error)
	AppendOrder(presenter.OrderDTO)
	RemoveOrder(int)
	ClearOrderList()

	BlockAdvertisementsList() *gtk.ListStore
	AppendBlockAdvertisement(presenter.BlockAdvertisementDTO)
	RemoveBlockAdvertisement(int)
	ClearBlockAdvertisementList()

	LineAdvertisementsList() *gtk.ListStore
	AppendLineAdvertisement(presenter.LineAdvertisementDTO)
	RemoveLineAdvertisement(int)
	ClearLineAdvertisementList()

	CostRatesListStore() *gtk.ListStore
	AppendCostRate(presenter.CostRateDTO)
	RemoveCostRate(string)
	ClearCostRateListStore()

	ReplaceMode() bool
	SetReplaceMode(bool)

	MessageList() *gtk.ListStore
	AppendMessage(string)
	ClearMessageList()

	NewFilesList()
	FilesList() *gtk.ListStore
	AppendFile(presenter.File)
	InsertFileFirstPlace(presenter.File)
	RemoveFile(string)

	LocalDatabaseList() *gtk.ListStore
	AppendLocalDatabase(string)
	RemoveLocalDatabase(string)
	ClearLocalDatabaseList()
}

type ObjectMaker interface {
	SetApplication(app *WindowController)
	LoginWindow() LoginWin
	AdvertisementsWindow() AdvertisementsWindow

	AddTagWindow() AddTagWindow
	AddExtraChargeWindow() AddExtraChargeWindow
	AddClientWindow() AddClientWindow
	AddOrderWindow() AddOrderWindow
	AddAdvertisementWindow() AddAdvertisementWindow
	ErrorWindow() ErrorWindow
	ImportDataWindow() ImportDataWindow
	ProgressWindow() ProgressWindow
	ListStores() ListStores
	CostRateWindow() CostRateWindow
	NewCostRateWindow(bool) NewCostRateWindow
	NewAdvertisementReportWindow() AdvertisementReportWindow
	NewLineForm() LineForm
	NewBlockForm() BlockForm
	NewThinImageChooserWindow(fileChooser FileChooser) FileChooseWindow
	NewSaveDialog(winLabel string, parent gtk.IWindow) (FileDialoger, error)
	NewFolderChooseDialog(winLabel string, parent gtk.IWindow) (FileDialoger, error)
	NewChooseDialog(winLabel string, parent gtk.IWindow) (FileDialoger, error)
}

type CostRateWindow interface {
	Window
}

type FileDialoger interface {
	AddFileFilter(name, pattern string) error
	BindResponseSignal(responseFunc func(self *glib.Object, responce int)) glib.SignalHandle
	GetURI() string
	Show()
	GetFilename() string
	SetCurrentName(string)
}

type FileChooseWindow interface {
	LoadFiles()
	ShowSelectedFile(presenter.File)
	LoadFilesComplete()
	Window
}

type AdvertisementReportWindow interface {
	Window
}

type NewCostRateWindow interface {
	Window() *gtk.Window
	Show()
	Close()

	SetName(string)
	SetCostForWordSymbol(string)
	SetCostForOneSquare(string)
	SetCostForOneWord(bool)
}

type advForms interface {
	Widget() *gtk.Widget
	ToSelectedOrder()
	SetSensetive(bool)
	SetNewAdvMode(bool)
	SetNewNestedAdvMode(bool)
	Reset()
	UnsetModel()
	SetModel()
}

type CostRateCalculator interface {
	SetActiveCostRate(string) error
	ActiveCostRate()
	CalculateBlockAdvertisementCost(*presenter.BlockAdvertisementDTO) error
	CalculateLineAdvertisementCost(*presenter.LineAdvertisementDTO) error
	CalculateOrderCost(order *presenter.OrderDTO, blkAdv []presenter.BlockAdvertisementDTO, lineAdv []presenter.LineAdvertisementDTO) error
	CheckCostRate()
}

type LineForm interface {
	advForms
	FillData(selectedLine *presenter.LineAdvertisementDTO)
	FetchData() presenter.LineAdvertisementDTO
}

type BlockForm interface {
	advForms
	FillData(block *presenter.BlockAdvertisementDTO)
	FetchData() presenter.BlockAdvertisementDTO
}

type FileChooser interface {
	FilePath() string
	SetFilePath(string)
	Box() *gtk.Box
}
