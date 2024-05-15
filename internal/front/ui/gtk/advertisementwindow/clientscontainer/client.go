package clientscontainer

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/gtk"
)

func (ct *ClientsTab) SelectedName() string {
	return ct.nameEntry.GetLayout().GetText()
}

func (ct *ClientsTab) SelectedPhone() string {
	return ct.contactNumberEntry.GetLayout().GetText()
}
func (ct *ClientsTab) SelectedEmail() string {
	return ct.emailEntry.GetLayout().GetText()
}
func (ct *ClientsTab) SelectedAdditionalInformation() string {
	startIter := ct.additionalInfoBuffer.GetStartIter()
	endIter := ct.additionalInfoBuffer.GetEndIter()
	additionalInformation, err := ct.additionalInfoBuffer.GetText(startIter, endIter, true)
	if err != nil {
		logging.Logger.Error("cannot get string from text buffer", "error", err)
	}
	return additionalInformation
}

func (ct *ClientsTab) SetSelectedName(name string) {
	ct.nameEntry.SetText(name)
}

func (ct *ClientsTab) SetSelectedPhone(phone string) {
	ct.contactNumberEntry.SetText(phone)
}

func (ct *ClientsTab) SetSelectedEmail(email string) {
	ct.emailEntry.SetText(email)
}

func (ct *ClientsTab) SetSelectedAdditionalInformation(adi string) {
	ct.additionalInfoBuffer.SetText(adi)
}

func (ct *ClientsTab) SetSensitive(s bool) {
	ct.nameEntry.SetSensitive(s)
	ct.contactNumberEntry.SetSensitive(s)
	ct.emailEntry.SetSensitive(s)
	ct.additionalInfoEntry.SetSensitive(s)
	ct.deleteButton.SetSensitive(s)
	ct.resetButton.SetSensitive(s)
	ct.applyButton.SetSensitive(s)
}

func (ct *ClientsTab) fillRightSide(cli *presenter.ClientDTO) {
	logging.Logger.Debug("Setting texts in right side")
	ct.nameEntry.SetText(cli.Name)
	ct.contactNumberEntry.SetText(cli.Phones)
	ct.emailEntry.SetText(cli.Email)
	ct.additionalInfoBuffer.SetText(cli.AdditionalInformation)
}

func (ct *ClientsTab) resetRightSide() {
	ct.nameEntry.SetText("")
	ct.contactNumberEntry.SetText("")
	ct.emailEntry.SetText("")
	ct.additionalInfoBuffer.SetText("")
	ct.SetSensitive(false)
}

func (ct *ClientsTab) filterOrders(model *gtk.TreeModel, iter *gtk.TreeIter) bool {
	if !ct.ordersFilterEnable {
		return false
	}
	orderClientName, err := ct.tools.StringFromIter(iter, model, 2)
	if err != nil {
		logging.Logger.Warn("filterOrders: got empty string. Skipping", "error", err)
		return false
	}
	return ct.req.ValueInString(orderClientName, ct.SelectedName())
}

func (ct *ClientsTab) BlockSignals() {
	ct.treeSelection.HandlerBlock(ct.signalHandler.treeSelectionChanged)
	ct.deleteButton.HandlerBlock(ct.signalHandler.deleteButtonClicked)
	ct.resetButton.HandlerBlock(ct.signalHandler.resetButtonClicked)
	ct.applyButton.HandlerBlock(ct.signalHandler.applyButton)
	ct.clients.HandlerBlock(ct.signalHandler.clientsRowDeleted)
	ct.clients.HandlerBlock(ct.signalHandler.clientsRowInserted)
	ct.clients.HandlerBlock(ct.signalHandler.clientsRowChanged)
}

func (ct *ClientsTab) UnblockSignals() {
	ct.treeSelection.HandlerUnblock(ct.signalHandler.treeSelectionChanged)
	ct.deleteButton.HandlerUnblock(ct.signalHandler.deleteButtonClicked)
	ct.resetButton.HandlerUnblock(ct.signalHandler.resetButtonClicked)
	ct.applyButton.HandlerUnblock(ct.signalHandler.applyButton)
	ct.clients.HandlerUnblock(ct.signalHandler.clientsRowDeleted)
	ct.clients.HandlerUnblock(ct.signalHandler.clientsRowInserted)
	ct.clients.HandlerUnblock(ct.signalHandler.clientsRowChanged)
}

func (ct *ClientsTab) SetEnableFilters(enable bool) {
	ct.ordersFilterEnable = enable
}

func (ct *ClientsTab) AttachList() {
	ct.treeView.SetModel(ct.clients)
	ct.ordersView.SetModel(ct.ordersFilter)
	ct.SetEnableFilters(true)
}

func (ct *ClientsTab) DetachList() {
	ct.treeView.SetModel(nil)
	ct.ordersView.SetModel(nil)
	ct.SetEnableFilters(false)
}

func (ct *ClientsTab) UnsetModel() {
	ct.resetRightSide()
	orde, _ := ct.ordersView.GetSelection()
	orde.UnselectAll()
	ct.treeSelection.UnselectAll()
	ct.treeView.Hide()
	ct.DetachList()
}

func (ct *ClientsTab) SetModel() {
	ct.AttachList()
	ct.treeView.Show()
}

func (ct *ClientsTab) ResetSort() {
	ct.clients.ResetDefaultSortFunc()
}
