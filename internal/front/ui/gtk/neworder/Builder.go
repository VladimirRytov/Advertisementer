package neworder

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/application"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"

	"github.com/gotk3/gotk3/gtk"
)

type OrderCreator interface {
	CreateOrder(*presenter.OrderDTO, []presenter.BlockAdvertisementDTO, []presenter.LineAdvertisementDTO) error
	CalculateOrderCost(*presenter.OrderDTO, []presenter.BlockAdvertisementDTO, []presenter.LineAdvertisementDTO) error
}

type AdvMaker interface {
	NewLineForm() application.LineForm
	NewLineCopyForm() application.LineForm

	NewBlockForm() application.BlockForm
	NewBlockCopyForm() application.BlockForm
}

type LineForm interface {
	SetNewNestedAdvMode(bool)
	FetchData() presenter.LineAdvertisementDTO
}

type BlockForm interface {
	SetNewNestedAdvMode(bool)
	FetchData() presenter.BlockAdvertisementDTO
}

type Tools interface {
	InterfaceFromIter(*gtk.TreeIter, *gtk.TreeModel, int) (interface{}, error)
	StringFromIter(*gtk.TreeIter, *gtk.TreeModel, int) (string, error)
	CheckCostString(*gtk.Entry, string)
}

type ApplicationGate interface {
	CreateAddClientWindow() application.AddClientWindow
	NewErrorWindow(error)
	RegisterReciever(*gtk.EntryBuffer)
}

type ListStores interface {
	OrderListCopy() (*gtk.ListStore, error)
	TagsListCopy() (*gtk.ListStore, error)
	ExtraChargeListCopy() (*gtk.ListStore, error)
	BlockAdvertisementsList() *gtk.ListStore
	LineAdvertisementsList() *gtk.ListStore
	ClientsList() *gtk.ListStore
}

type NewOrderWindow struct {
	req       OrderCreator
	app       ApplicationGate
	lists     ListStores
	tools     Tools
	formMaker AdvMaker

	window             *gtk.Window
	clientTreeView     *gtk.TreeView
	clientSelectToggle *gtk.CellRendererToggle
	clientListStore    *gtk.ListStore
	newClientButton    *gtk.Button

	paymentTypeLabel *gtk.Label
	paymentTypeEntry *gtk.Entry

	costLabel           *gtk.Label
	costEntry           *gtk.Entry
	calculateCostButton *gtk.Button

	paymentStatusLabel    *gtk.Label
	paymentStatusCheckBox *gtk.CheckButton

	selectExistedCheckButton *gtk.CheckButton
	appendExistRevealer      *gtk.Revealer

	lineAdvListStore      *gtk.ListStore
	lineAdvTreeView       *gtk.TreeView
	lineAdvTreeViewToggle *gtk.CellRendererToggle

	blockAdvListStore  *gtk.ListStore
	blockAdvTreeView   *gtk.TreeView
	blockAdvTreeToggle *gtk.CellRendererToggle

	newAdvertisementCheckButton *gtk.CheckButton
	newAdvertisementRevealer    *gtk.Revealer

	newAdvPopover            *gtk.Popover
	newAdvBox                *gtk.Box
	deleteAdvButton          *gtk.Button
	addBlockAdvButton        *gtk.Button
	addLineAdvButton         *gtk.Button
	newAdvertisementNotebook *gtk.Notebook

	cancelButton *gtk.Button
	createButton *gtk.Button

	advertisements []interface{}

	signalHandler signalHandler
}

func Create(req OrderCreator, app ApplicationGate, lists ListStores, tools Tools, formMaker AdvMaker) *NewOrderWindow {
	now := new(NewOrderWindow)
	build, err := builder.NewBuilderFromString(builder.AddOrderWindow)
	if err != nil {
		panic(err)
	}
	now.req = req
	now.formMaker = formMaker
	now.app = app
	now.tools = tools
	now.lists = lists
	now.Build(build)
	now.clientTreeView.SetModel(now.clientListStore)
	now.blockAdvTreeView.SetModel(now.blockAdvListStore)
	now.lineAdvTreeView.SetModel(now.lineAdvListStore)
	now.appendExistRevealer.SetRevealChild(now.selectExistedCheckButton.GetActive())
	now.newAdvertisementRevealer.SetRevealChild(now.newAdvertisementCheckButton.GetActive())
	now.newAdvertisementNotebook, err = gtk.NotebookNew()
	if err != nil {
		panic(err)
	}
	now.costEntry.SetPlaceholderText("0,00")
	now.newAdvertisementNotebook.SetScrollable(true)
	now.newAdvBox.PackEnd(now.newAdvertisementNotebook.ToWidget(), true, true, 0)
	now.deleteAdvButton.SetSensitive(false)
	now.advertisements = make([]interface{}, 0)
	now.bindSignals()
	now.window.SetTitle("Создание заказа")
	return now
}

func (now *NewOrderWindow) Build(buildFile *builder.Builder) {
	now.window = buildFile.FetchWindow("AddOrderWindow")

	now.clientTreeView = buildFile.FetchTreeView("SelectedClientTreeView")
	now.clientSelectToggle = buildFile.FetchCellRendererToggle("ClientSelectedToggle")
	now.clientListStore = now.lists.ClientsList()
	now.newClientButton = buildFile.FetchButton("AddClientButton")
	now.paymentTypeLabel = buildFile.FetchLabel("OrdersPaymentTypeLabel")
	now.paymentTypeEntry = buildFile.FetchEntry("OrdersPaymentTypeEntry")
	now.costLabel = buildFile.FetchLabel("OrdersTotalPaymentLabel")
	now.costEntry = buildFile.FetchEntry("OrdersTotalPaymentEntry")
	now.calculateCostButton = buildFile.FetchButton("CalculateCostButton")
	now.paymentStatusLabel = buildFile.FetchLabel("OrdersPaymentStatusLabel")
	now.paymentStatusCheckBox = buildFile.FetchCheckButton("OrdersPaymentStatusCheckButton")

	now.selectExistedCheckButton = buildFile.FetchCheckButton("SelectExistedCheckButton")
	now.appendExistRevealer = buildFile.FetchRevealer("AppendExistsAdvertisementsReleaver")
	now.lineAdvListStore = now.lists.LineAdvertisementsList()
	now.lineAdvTreeView = buildFile.FetchTreeView("LineAdvertisementsTreeView")
	now.lineAdvTreeViewToggle = buildFile.FetchCellRendererToggle("LineCellRenderToggle")

	now.blockAdvListStore = now.lists.BlockAdvertisementsList()
	now.blockAdvTreeView = buildFile.FetchTreeView("BlockAdvertisementsTreeView")
	now.blockAdvTreeToggle = buildFile.FetchCellRendererToggle("BlockCellRenderToggle")

	now.newAdvPopover = buildFile.FetchPopover("AddAdvertisementsPopover")
	now.newAdvertisementRevealer = buildFile.FetchRevealer("AddNewAdvertisementsReleaver")
	now.newAdvertisementCheckButton = buildFile.FetchCheckButton("AddNewAdvertisementsCheckButton")
	now.newAdvBox = buildFile.FetchBox("AddNewAdvertisementBox")
	now.deleteAdvButton = buildFile.FetchButton("RemoveAdvertisementBtn")
	now.addBlockAdvButton = buildFile.FetchButton("AddBlockAdvertisementButton")
	now.addLineAdvButton = buildFile.FetchButton("AddLineAdvertisementButton")
	now.cancelButton = buildFile.FetchButton("CancelButton")
	now.createButton = buildFile.FetchButton("CreateButton")
}

func (now *NewOrderWindow) Window() *gtk.Window {
	return now.window
}

func (now *NewOrderWindow) Show() {
	now.window.Show()
}

func (now *NewOrderWindow) Close() {
	now.window.Close()
}
