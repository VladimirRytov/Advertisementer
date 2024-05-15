package clientscontainer

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
)

type ClientTabSignalHandler struct {
	treeSelectionChanged glib.SignalHandle
	deleteButtonClicked  glib.SignalHandle
	resetButtonClicked   glib.SignalHandle
	applyButton          glib.SignalHandle
	clientsRowDeleted    glib.SignalHandle
	clientsRowInserted   glib.SignalHandle
	clientsRowChanged    glib.SignalHandle
}

func (ct *ClientsTab) bindSignals() {
	ct.signalHandler = ClientTabSignalHandler{
		treeSelectionChanged: ct.treeSelection.Connect("changed", ct.SelectedRowChanged),
		deleteButtonClicked:  ct.deleteButton.Connect("clicked", ct.removeButtonPressed),
		resetButtonClicked:   ct.resetButton.Connect("clicked", ct.resetButtonPressed),
		applyButton:          ct.applyButton.Connect("clicked", ct.updateButtonPressed),
		clientsRowDeleted:    ct.clients.Connect("row-deleted", ct.rowChanged),
		clientsRowInserted:   ct.clients.Connect("row-inserted", ct.rowChanged),
		clientsRowChanged:    ct.clients.Connect("row-changed", ct.rowChanged),
	}
}

func (ct *ClientsTab) SelectedRowChanged() {
	model, iter, ok := ct.treeSelection.GetSelected()
	if !ok {
		ct.resetRightSide()
		ct.SetSensitive(false)
		return
	}
	ct.SetSensitive(true)
	name, err := ct.tools.StringFromIter(iter, model.ToTreeModel(), 1)
	if err != nil {
		logging.Logger.Error("Cannot convert get string", "error message", err)
	}

	phone, err := ct.tools.StringFromIter(iter, model.ToTreeModel(), 2)
	if err != nil {
		logging.Logger.Error("Cannot convert get string", "error message", err)
	}

	email, err := ct.tools.StringFromIter(iter, model.ToTreeModel(), 3)
	if err != nil {
		logging.Logger.Error("Cannot convert get string", "error message", err)
	}
	addInfo, err := ct.tools.StringFromIter(iter, model.ToTreeModel(), 4)
	if err != nil {
		logging.Logger.Error("Cannot convert get string", "error message", err)
	}

	cli := &presenter.ClientDTO{
		Name:                  name,
		Phones:                phone,
		Email:                 email,
		AdditionalInformation: addInfo,
	}

	ct.fillRightSide(cli)
	ct.ordersFilter.Refilter()
}

func (ct *ClientsTab) removeButtonPressed() {
	ct.SetSensitive(false)
	go ct.req.RemoveClient(ct.SelectedName())
}

func (ct *ClientsTab) resetButtonPressed() {
	ct.SelectedRowChanged()
}

func (ct *ClientsTab) updateButtonPressed() {
	client := &presenter.ClientDTO{
		Name:                  ct.SelectedName(),
		Phones:                ct.SelectedPhone(),
		Email:                 ct.SelectedEmail(),
		AdditionalInformation: ct.SelectedAdditionalInformation(),
	}
	ct.req.UpdateClient(client)
}

func (ct *ClientsTab) rowChanged() {
	_, _, ok := ct.treeSelection.GetSelected()
	if !ok {
		ct.resetRightSide()
		ct.SetSensitive(false)
		return
	}
	ch := ct.clients.IterNChildren(nil)
	if ch == 0 {
		ct.resetRightSide()
		ct.SetSensitive(false)
		return
	}
	ct.SetSensitive(true)
}
