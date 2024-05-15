package costratewindow

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/application"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"

	"github.com/gotk3/gotk3/gtk"
)

var CostRate string = "resources/CostRateWindow.glade"

type CostRateRequests interface {
	LoadConfig(string, any) error
	SaveConfig(string, any) error
	SetActiveCostRate(string) error
	AllCostRates()
	RemoveCostRate(string)
}

type Tools interface {
	StringFromIter(*gtk.TreeIter, *gtk.TreeModel, int) (string, error)
	InterfaceFromIter(*gtk.TreeIter, *gtk.TreeModel, int) (interface{}, error)
}

type Application interface {
	CreateNewCostRatesWindow(bool) application.NewCostRateWindow
	SelectCostRate(string)
	NewErrorWindow(error)
}

type ListStores interface {
	CostRatesListStore() *gtk.ListStore
}

type CostRateWindow struct {
	req   CostRateRequests
	app   Application
	tools Tools
	lists ListStores

	signals CostRateSignals
	window  *gtk.Window

	list               *gtk.ListStore
	treeview           *gtk.TreeView
	chooseSelector     *gtk.CellRendererToggle
	forWordSelection   *gtk.CellRendererToggle
	forSymbolSelection *gtk.CellRendererToggle

	addButton     *gtk.Button
	delButton     *gtk.Button
	editButton    *gtk.Button
	refreshButton *gtk.Button

	closeButton *gtk.Button
}

func Create(reqGate CostRateRequests, app Application, lists ListStores, tools Tools) *CostRateWindow {
	buildfile, err := builder.NewBuilderFromString(builder.CostRateWindow)
	if err != nil {
		panic(err)
	}
	cs := new(CostRateWindow)
	cs.req = reqGate
	cs.app = app
	cs.tools = tools
	cs.lists = lists
	cs.build(buildfile)
	cs.list = lists.CostRatesListStore()
	cs.bindSignals()
	cs.window.SetModal(true)
	cs.treeview.SetModel(lists.CostRatesListStore())
	cs.window.SetTitle("Управление тарифами")
	return cs
}

func (cs *CostRateWindow) build(buildFile *builder.Builder) {
	cs.window = buildFile.FetchWindow("CostRateWindow")
	cs.treeview = buildFile.FetchTreeView("CostRateTreeView")
	cs.chooseSelector = buildFile.FetchCellRendererToggle("CelectedCellRender")
	cs.forWordSelection = buildFile.FetchCellRendererToggle("WordCostCellRender")
	cs.forSymbolSelection = buildFile.FetchCellRendererToggle("SymbolCostCellRender")
	cs.addButton = buildFile.FetchButton("AddRowButton")
	cs.delButton = buildFile.FetchButton("DeleteRowButton")
	cs.refreshButton = buildFile.FetchButton("RefreshCostRates")
	cs.editButton = buildFile.FetchButton("EditButton")
	cs.closeButton = buildFile.FetchButton("CloseButton")
}
