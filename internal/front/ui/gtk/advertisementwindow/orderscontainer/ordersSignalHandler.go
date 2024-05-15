package orderscontainer

import (
	"time"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type OrdersSignalHandler struct {
	selectedRowChanged             glib.SignalHandle
	selectedClientToggled          glib.SignalHandle
	deleteButtonClicked            glib.SignalHandle
	resetButtonClicked             glib.SignalHandle
	applyButtonClicked             glib.SignalHandle
	createdDateEntryIconPress      glib.SignalHandle
	createdDateCalendarDaySelected glib.SignalHandle
	ordersListStoreRowDeleted      glib.SignalHandle
	ordersListStoreRowInserted     glib.SignalHandle
	ordersListStoreRowChanged      glib.SignalHandle
	orderCostTextInserted          glib.SignalHandle
	calculateOrderButtonPressed    glib.SignalHandle
}

func (or *OrdersTab) bindSignals() {
	or.signalHandler = OrdersSignalHandler{
		selectedRowChanged:             or.ordersListStoreSelection.Connect("changed", or.SelectedRowChanged),
		selectedClientToggled:          or.clientCellRender.Connect("toggled", or.ClientsToggled),
		deleteButtonClicked:            or.deleteButton.Connect("clicked", or.removeButtonPressed),
		resetButtonClicked:             or.resetButton.Connect("clicked", or.resetButtonPressed),
		applyButtonClicked:             or.applyButton.Connect("clicked", or.updateButtonPressed),
		createdDateEntryIconPress:      or.createdDateEntry.Connect("icon-press", or.createDateEntryIcoPressed),
		createdDateCalendarDaySelected: or.createdDateCalendar.Connect("day-selected", or.calendarDaySelected),
		ordersListStoreRowDeleted:      or.OrdersListStore.Connect("row-deleted", or.rowChanged),
		ordersListStoreRowInserted:     or.OrdersListStore.Connect("row-inserted", or.rowChanged),
		ordersListStoreRowChanged:      or.OrdersListStore.Connect("row-changed", or.rowChanged),
		orderCostTextInserted:          or.costEntry.Connect("insert-text", or.tools.CheckCostAdvString),
		calculateOrderButtonPressed:    or.costCalculateCostButton.Connect("clicked", or.calculateOrderPressed),
	}
}

func (or *OrdersTab) SelectedRowChanged() {

	model, iter, ok := or.ordersListStoreSelection.GetSelected()
	if !ok {
		or.resetSidebar()
		return
	}
	or.SetSensitive(true)
	rawId, err := or.tools.InterfaceFromIter(iter, model.ToTreeModel(), 1)
	if err != nil {
		logging.Logger.Error("SelectedRowChanged: Cannot get value", "error message", err)
	}

	clientName, err := or.tools.StringFromIter(iter, model.ToTreeModel(), 2)
	if err != nil {
		logging.Logger.Error("SelectedRowChanged: Cannot get value", "error message", err)
	}

	createDate, err := or.tools.StringFromIter(iter, model.ToTreeModel(), 3)
	if err != nil {
		logging.Logger.Error("SelectedRowChanged: Cannot get value", "error message", err)
	}

	cost, err := or.tools.StringFromIter(iter, model.ToTreeModel(), 4)
	if err != nil {
		logging.Logger.Error("SelectedRowChanged: Cannot get value", "error message", err)
	}

	paymentType, err := or.tools.StringFromIter(iter, model.ToTreeModel(), 5)
	if err != nil {
		logging.Logger.Error("SelectedRowChanged: Cannot get value", "error message", err)
	}

	paymentStatusRaw, err := or.tools.InterfaceFromIter(iter, model.ToTreeModel(), 6)
	if err != nil {
		logging.Logger.Error("SelectedRowChanged: Cannot get value", "error message", err)
	}

	paymentStatus, ok := paymentStatusRaw.(bool)
	if !ok {
		return
	}

	order := &presenter.OrderDTO{
		ID:            rawId.(int),
		ClientName:    clientName,
		Cost:          cost,
		PaymentType:   paymentType,
		CreatedDate:   createDate,
		PaymentStatus: paymentStatus,
	}
	or.fillRightSide(order)
	or.blocksAdvFilter.Refilter()
	or.linesAdvFilter.Refilter()
}

func (or *OrdersTab) removeButtonPressed() {
	or.SetSensitive(false)
	order := &presenter.OrderDTO{
		ID:            or.id(),
		ClientName:    or.client(),
		Cost:          or.cost(),
		PaymentType:   or.paymentType(),
		PaymentStatus: or.paymentStatus(),
		CreatedDate:   or.createdDate(),
	}
	or.req.RemoveOrder(order)
}

func (or *OrdersTab) resetButtonPressed() {
	or.SelectedRowChanged()
}

func (or *OrdersTab) updateButtonPressed() {
	or.SetSensitive(false)
	order := &presenter.OrderDTO{
		ID:            or.id(),
		ClientName:    or.client(),
		Cost:          or.cost(),
		PaymentType:   or.paymentType(),
		PaymentStatus: or.paymentStatus(),
		CreatedDate:   or.createdDate(),
	}
	err := or.req.UpdateOrder(order)
	if err != nil {
		or.SetSensitive(true)
		or.advWin.ShowPopover(err.Error(), or.costEntry)
	}
}

func (or *OrdersTab) createDateEntryIcoPressed() {
	t, err := time.Parse("02.01.2006", or.createdDate())
	if err != nil {
		logging.Logger.Error("createDateEntryIcoPressed: cannot convert string date to time type", "error", err)
	}
	or.createdDateCalendar.SelectMonth(uint(t.Month()-1), uint(t.Year()))
	or.createdDateCalendar.SelectDay(uint(t.Day()))
	or.createdDateEntryBPopover.Popup()
}

func (or *OrdersTab) calendarDaySelected() {
	y, m, d := or.createdDateCalendar.GetDate()
	or.setCreatedDate(y, m, d)
}

func (or *OrdersTab) rowChanged() {
	_, _, ok := or.ordersListStoreSelection.GetSelected()
	if !ok {
		or.resetSidebar()
		or.SetSensitive(false)
		return
	}
	if or.OrdersListStore.IterNChildren(nil) == 0 {
		or.resetSidebar()
		or.SetSensitive(false)
		return
	}
	or.SetSensitive(true)
}

func (or *OrdersTab) ClientsToggled(self *gtk.CellRendererToggle, pathStr string) {
	or.clientsListStore.ForEach(func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
		or.clientsListStore.SetValue(iter, 0, path.String() == pathStr)
		return false
	})
}

func (or *OrdersTab) calculateOrderPressed() {
	b, err := or.costEntry.GetBuffer()
	if err != nil {
		logging.Logger.Error("ordersTab.calculateOrderPressed: cannot get cost entry buffer", "error", err)
		return
	}
	or.app.RegisterReciever(b)
	order := &presenter.OrderDTO{
		ID:            or.id(),
		ClientName:    or.client(),
		Cost:          or.cost(),
		PaymentType:   or.paymentType(),
		PaymentStatus: or.paymentStatus(),
		CreatedDate:   or.createdDate(),
	}
	or.req.CalculateOrderCost(order, nil, nil)
}
