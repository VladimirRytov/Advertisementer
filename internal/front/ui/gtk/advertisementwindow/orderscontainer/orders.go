package orderscontainer

import (
	"errors"
	"strconv"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/gtk"
)

func (or *OrdersTab) id() int {
	id, _ := strconv.Atoi(or.orderNumberEntry.GetLayout().GetText())
	return id
}

func (or *OrdersTab) setId(id string) {
	or.orderNumberEntry.SetText(id)
}

func (or *OrdersTab) client() string {
	var client string
	or.clientsListStore.ForEach(
		func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
			selected, err := or.tools.InterfaceFromIter(iter, &or.clientsListStore.TreeModel, 0)
			if err != nil {
				logging.Logger.Error("client.Client: cannot get bool from client listStore", "error", err)
				return false
			}
			if sel := selected.(bool); sel {
				name, err := or.tools.StringFromIter(iter, &or.clientsListStore.TreeModel, 1)
				if err != nil {
					logging.Logger.Error("client.Client: cannot get name from client listStore", "error", err)
					return false
				}
				client = name
				return true
			}
			return false
		})

	return client
}

func (or *OrdersTab) cost() string {
	return or.costEntry.GetLayout().GetText()
}

func (or *OrdersTab) setCost(cost string) {
	or.costEntry.SetText(cost)
}

func (or *OrdersTab) paymentType() string {
	return or.paymentTypeEntry.GetLayout().GetText()
}

func (or *OrdersTab) setPaymentType(payType string) {
	or.paymentTypeEntry.SetText(payType)
}

func (or *OrdersTab) paymentStatus() bool {
	return or.paymentStatusCheckButton.GetActive()
}

func (or *OrdersTab) setPaymentStatus(status bool) {
	or.paymentStatusCheckButton.SetActive(status)
}
func (or *OrdersTab) createdDate() string {
	return or.createdDateEntry.GetLayout().GetText()
}

func (or *OrdersTab) setCreatedDate(year, month, day uint) {
	or.createdDateEntry.SetText(or.conv.YearMonthDayToString(year, month+1, day))
}

func (or *OrdersTab) resetSidebar() {
	or.orderNumberEntry.SetText("")
	or.costEntry.SetText("")
	or.createdDateEntry.SetText("")
	or.paymentTypeEntry.SetText("")
	or.paymentStatusCheckButton.SetActive(false)
	or.SetSensitive(false)
}

func (or *OrdersTab) SetSensitive(lock bool) {
	or.clientTreeView.SetSensitive(lock)
	or.blocksAdvTreeView.SetSensitive(lock)
	or.linesAdvTreeView.SetSensitive(lock)
	or.orderNumberEntry.SetSensitive(lock)
	or.costEntry.SetSensitive(lock)
	or.paymentTypeEntry.SetSensitive(lock)
	or.paymentStatusCheckButton.SetSensitive(lock)
	or.costCalculateCostButton.SetSensitive(lock)
	or.deleteButton.SetSensitive(lock)
	or.resetButton.SetSensitive(lock)
	or.applyButton.SetSensitive(lock)
	or.createdDateEntry.SetSensitive(lock)
}

func (or *OrdersTab) fillRightSide(order *presenter.OrderDTO) {

	logging.Logger.Debug("fillRightSide: Setting texts in right side")
	or.setId(strconv.Itoa(order.ID))
	path, err := or.MarkSelectedClient(order.ClientName)
	if err == nil {
		or.ToSelectedClient(path)
	}
	or.createdDateEntry.SetText(order.CreatedDate)
	or.setCost(order.Cost)
	or.setPaymentType(order.PaymentType)
	or.setPaymentStatus(order.PaymentStatus)
}

func (or *OrdersTab) filterAdvertisements(model *gtk.TreeModel, iter *gtk.TreeIter) bool {
	if !or.enableFilters {
		return true
	}

	blockOrderID, err := or.tools.InterfaceFromIter(iter, model, 2)
	if err != nil {
		logging.Logger.Warn("filterAdvertisements: got empty string. Skipping", "error", err)
		return false
	}

	return blockOrderID.(int) == or.id()
}

func (or *OrdersTab) ToSelectedClient(path *gtk.TreePath) {
	or.clientTreeView.ScrollToCell(path, nil, true, 0.5, 0)
}

func (or *OrdersTab) MarkSelectedClient(selected string) (*gtk.TreePath, error) {
	var selectedPath *gtk.TreePath
	or.clientsListStore.ForEach(
		func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
			clientName, err := or.tools.StringFromIter(iter, &or.clientsListStore.TreeModel, 1)
			if err != nil {
				logging.Logger.Error("markSelectedOrder: an error occured while getting order id from liststore", "error", err)
				return false
			}

			or.clientsListStore.SetValue(iter, 0, clientName == selected)
			if clientName == selected {
				selectedPath, err = or.clientsListStore.GetPath(iter)
				if err != nil {
					logging.Logger.Error("filler.markSelectedOrder: an error accured while getting path from listStore", "error", err)
				}
			}
			return false
		})
	if selectedPath != nil {
		return selectedPath, nil
	}

	return nil, errors.New("cant get path")
}

func (or *OrdersTab) BlockSignals() {
	or.ordersListStoreSelection.HandlerBlock(or.signalHandler.selectedRowChanged)
	or.clientCellRender.HandlerBlock(or.signalHandler.selectedClientToggled)
	or.deleteButton.HandlerBlock(or.signalHandler.deleteButtonClicked)
	or.resetButton.HandlerBlock(or.signalHandler.resetButtonClicked)
	or.applyButton.HandlerBlock(or.signalHandler.applyButtonClicked)
	or.createdDateEntry.HandlerBlock(or.signalHandler.createdDateEntryIconPress)
	or.createdDateCalendar.HandlerBlock(or.signalHandler.createdDateCalendarDaySelected)
	or.OrdersListStore.HandlerBlock(or.signalHandler.ordersListStoreRowDeleted)
	or.OrdersListStore.HandlerBlock(or.signalHandler.ordersListStoreRowInserted)
	or.OrdersListStore.HandlerBlock(or.signalHandler.ordersListStoreRowChanged)
}

func (or *OrdersTab) UnblockSignals() {
	or.ordersListStoreSelection.HandlerUnblock(or.signalHandler.selectedRowChanged)
	or.clientCellRender.HandlerUnblock(or.signalHandler.selectedClientToggled)
	or.deleteButton.HandlerUnblock(or.signalHandler.deleteButtonClicked)
	or.resetButton.HandlerUnblock(or.signalHandler.resetButtonClicked)
	or.applyButton.HandlerUnblock(or.signalHandler.applyButtonClicked)
	or.createdDateEntry.HandlerUnblock(or.signalHandler.createdDateEntryIconPress)
	or.createdDateCalendar.HandlerUnblock(or.signalHandler.createdDateCalendarDaySelected)
	or.OrdersListStore.HandlerUnblock(or.signalHandler.ordersListStoreRowDeleted)
	or.OrdersListStore.HandlerUnblock(or.signalHandler.ordersListStoreRowInserted)
	or.OrdersListStore.HandlerUnblock(or.signalHandler.ordersListStoreRowChanged)
}

func (or *OrdersTab) SetEnableFilters(enable bool) {
	or.enableFilters = enable
}

func (or *OrdersTab) AttachList() {
	or.ordersTreeView.SetModel(or.OrdersListStore)
	or.clientTreeView.SetModel(or.clientsListStore)
	or.blocksAdvTreeView.SetModel(or.blocksAdvFilter)
	or.linesAdvTreeView.SetModel(or.linesAdvFilter)
	or.SetEnableFilters(true)

}

func (or *OrdersTab) DetachList() {
	or.ordersTreeView.SetModel(nil)
	or.clientTreeView.SetModel(nil)
	or.blocksAdvTreeView.SetModel(nil)
	or.linesAdvTreeView.SetModel(nil)
	or.SetEnableFilters(false)
}

func (or *OrdersTab) UnsetModel() {
	or.resetSidebar()
	bloks, _ := or.blocksAdvTreeView.GetSelection()
	bloks.UnselectAll()
	lines, _ := or.linesAdvTreeView.GetSelection()
	lines.UnselectAll()

	or.ordersListStoreSelection.UnselectAll()
	or.ordersTreeView.Hide()
	or.DetachList()
}

func (or *OrdersTab) SetModel() {
	or.AttachList()
	or.ordersTreeView.Show()
}

func (or *OrdersTab) ResetSort() {
	or.OrdersListStore.ResetDefaultSortFunc()
}
