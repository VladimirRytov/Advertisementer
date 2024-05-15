package orderscontainer

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"

	"github.com/gotk3/gotk3/gtk"
)

type Tools interface {
	CheckCostAdvString(*gtk.Entry, string)
	CompareInts(model *gtk.TreeModel, a, b *gtk.TreeIter, col int) int
	CompareCosts(model *gtk.TreeModel, a, b *gtk.TreeIter, col int) int
	CompareBools(model *gtk.TreeModel, a, b *gtk.TreeIter, col int) int
	CompareStrings(model *gtk.TreeModel, a, b *gtk.TreeIter, col int) int
	StringFromIter(*gtk.TreeIter, *gtk.TreeModel, int) (string, error)

	InterfaceFromIter(*gtk.TreeIter, *gtk.TreeModel, int) (interface{}, error)
}

type ListStores interface {
	ClientsList() *gtk.ListStore
	OrdersList() *gtk.ListStore
	BlockAdvertisementsList() *gtk.ListStore
	LineAdvertisementsList() *gtk.ListStore
}

type DataConverter interface {
	YearMonthDayToString(uint, uint, uint) string
}

type OrderHandler interface {
	CalculateOrderCost(order *presenter.OrderDTO, blkAdv []presenter.BlockAdvertisementDTO, lineAdv []presenter.LineAdvertisementDTO) error
	UpdateOrder(*presenter.OrderDTO) error
	RemoveOrder(*presenter.OrderDTO)
	ValueInString(string, string) bool
}

type Application interface {
	RegisterReciever(*gtk.EntryBuffer)
}

type AdvWin interface {
	ShowPopover(string, gtk.IWidget)
}

type OrdersTab struct {
	tools  Tools
	req    OrderHandler
	conv   DataConverter
	advWin AdvWin
	lists  ListStores
	app    Application

	sidebar                  *gtk.Box
	enableFilters            bool
	ordersTreeView           *gtk.TreeView
	OrdersListStore          *gtk.TreeModelSort
	ordersListStoreSelection *gtk.TreeSelection

	orderNumberLabel *gtk.Label
	orderNumberEntry *gtk.Entry

	clientLabel      *gtk.Label
	clientTreeView   *gtk.TreeView
	clientsListStore *gtk.ListStore
	clientCellRender *gtk.CellRendererToggle

	costLabel               *gtk.Label
	costEntry               *gtk.Entry
	costCalculateCostButton *gtk.Button

	paymentTypeLabel *gtk.Label
	paymentTypeEntry *gtk.Entry

	paymentStatusLabel       *gtk.Label
	paymentStatusCheckButton *gtk.CheckButton

	createdDateLabel         *gtk.Label
	createdDateEntry         *gtk.Entry
	createdDateEntryBPopover *gtk.Popover
	createdDateCalendar      *gtk.Calendar

	deleteButton *gtk.Button
	resetButton  *gtk.Button
	applyButton  *gtk.Button

	blocksAdvLabel    *gtk.Label
	blocksAdvFilter   *gtk.TreeModelFilter
	blocksAdvTreeView *gtk.TreeView

	linesAdvLabel    *gtk.Label
	linesAdvFilter   *gtk.TreeModelFilter
	linesAdvTreeView *gtk.TreeView

	signalHandler OrdersSignalHandler
}

func Create(bldFile *builder.Builder, req OrderHandler, conv DataConverter, tools Tools, advWin AdvWin, lists ListStores, app Application) *OrdersTab {
	var err error
	orders := new(OrdersTab)
	orders.app = app
	orders.tools = tools
	orders.req = req
	orders.lists = lists
	orders.conv = conv
	orders.build(bldFile)

	orders.OrdersListStore, err = gtk.TreeModelSortNew(orders.lists.OrdersList())
	if err != nil {
		panic(err)
	}
	orders.OrdersListStore.SetSortFunc(1, func(model *gtk.TreeModel, a, b *gtk.TreeIter) int {
		return orders.tools.CompareInts(model, a, b, 1)
	})
	orders.OrdersListStore.SetSortFunc(2, func(model *gtk.TreeModel, a, b *gtk.TreeIter) int {
		return orders.tools.CompareStrings(model, a, b, 2)
	})
	orders.OrdersListStore.SetSortFunc(3, func(model *gtk.TreeModel, a, b *gtk.TreeIter) int {
		return orders.tools.CompareStrings(model, a, b, 3)
	})
	orders.OrdersListStore.SetSortFunc(4, func(model *gtk.TreeModel, a, b *gtk.TreeIter) int {
		return orders.tools.CompareCosts(model, a, b, 4)
	})
	orders.OrdersListStore.SetSortFunc(5, func(model *gtk.TreeModel, a, b *gtk.TreeIter) int {
		return orders.tools.CompareStrings(model, a, b, 5)
	})
	orders.OrdersListStore.SetSortFunc(6, func(model *gtk.TreeModel, a, b *gtk.TreeIter) int {
		return orders.tools.CompareBools(model, a, b, 6)
	})

	orders.clientsListStore = orders.lists.ClientsList()
	orders.clientTreeView.SetModel(orders.clientsListStore)
	orders.ordersListStoreSelection, err = orders.ordersTreeView.GetSelection()
	if err != nil {
		panic(err)
	}
	orders.blocksAdvFilter, err = orders.lists.BlockAdvertisementsList().FilterNew(nil)
	if err != nil {
		panic(err)
	}
	orders.blocksAdvFilter.SetVisibleFunc(orders.filterAdvertisements)
	orders.blocksAdvTreeView.SetModel(orders.blocksAdvFilter)
	orders.linesAdvFilter, err = orders.lists.LineAdvertisementsList().FilterNew(nil)
	if err != nil {
		panic(err)
	}
	orders.linesAdvFilter.SetVisibleFunc(orders.filterAdvertisements)
	orders.linesAdvTreeView.SetModel(orders.linesAdvFilter)

	orders.bindSignals()
	orders.BlockSignals()
	return orders
}

func (or *OrdersTab) build(bldFile *builder.Builder) {
	or.sidebar = bldFile.FetchBox("OrderSidebar")
	or.ordersTreeView = bldFile.FetchTreeView("OrdersTreeView")
	or.orderNumberLabel = bldFile.FetchLabel("OrdersNumberLabel")
	or.orderNumberEntry = bldFile.FetchEntry("OrdersNumberEntry")
	or.clientLabel = bldFile.FetchLabel("OrdersClientLabel")
	or.clientTreeView = bldFile.FetchTreeView("SelectedClientTreeView")
	or.clientCellRender = bldFile.FetchCellRendererToggle("ClientSelectedToggle")
	or.costLabel = bldFile.FetchLabel("OrdersTotalPaymentLabel")
	or.costEntry = bldFile.FetchEntry("OrdersTotalPaymentEntr")
	or.costCalculateCostButton = bldFile.FetchButton("CalculateCostOrderButton")
	or.paymentTypeLabel = bldFile.FetchLabel("OrdersPaymentTypeLabel")
	or.paymentTypeEntry = bldFile.FetchEntry("OrdersPaymentTypeEntry")
	or.paymentStatusLabel = bldFile.FetchLabel("OrdersPaymentStatusLabel")
	or.createdDateLabel = bldFile.FetchLabel("CreatedDateLabel")
	or.createdDateEntry = bldFile.FetchEntry("CreatedDateEntry")
	or.createdDateEntryBPopover = bldFile.FetchPopover("OrderSetReleaseDatePopover")
	or.createdDateCalendar = bldFile.FetchCalendar("OrderSetReleaseDateCalendar")
	or.paymentStatusCheckButton = bldFile.FetchCheckButton("OrdersPaymentStatusCheckButton")
	or.deleteButton = bldFile.FetchButton("DeleteOrderButton")
	or.resetButton = bldFile.FetchButton("OrdersResetButton")
	or.applyButton = bldFile.FetchButton("OrdersApplyButton")
	or.blocksAdvLabel = bldFile.FetchLabel("OrderBlockAdvLabel")
	or.blocksAdvTreeView = bldFile.FetchTreeView("OrderBlockAdvertisementsTreeView")
	or.linesAdvLabel = bldFile.FetchLabel("OrderLineAdvLabel")
	or.linesAdvTreeView = bldFile.FetchTreeView("OrderLineAdvertisementsTreeView")
}
