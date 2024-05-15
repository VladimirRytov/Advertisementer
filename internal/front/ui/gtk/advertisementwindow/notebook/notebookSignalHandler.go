package notebook

import (
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
)

type notebookSignalHandler struct {
	updateButtonClicked               glib.SignalHandle
	createAdvertisementsButtonClicked glib.SignalHandle
	createClientButtonClicked         glib.SignalHandle
	createOrderButtonClicked          glib.SignalHandle
	createTagButtonClicked            glib.SignalHandle
	createExtraChargeButtonClicked    glib.SignalHandle
	notebookCurrendPageChanged        glib.SignalHandle
}

func (n *Notebook) bindSignals() {
	n.SignalHandler = notebookSignalHandler{
		updateButtonClicked:               n.updateButton.Connect("clicked", n.UpdateButtonPressed),
		createAdvertisementsButtonClicked: n.createAdvertisements.Connect("clicked", n.CreateAdvertisementsPressed),
		createClientButtonClicked:         n.createClient.Connect("clicked", n.CreateClientPressed),
		createOrderButtonClicked:          n.createOrder.Connect("clicked", n.CreateOrderPressed),
		createTagButtonClicked:            n.createTag.Connect("clicked", n.CreateTagPressed),
		createExtraChargeButtonClicked:    n.createExtraCharge.Connect("clicked", n.CreateExtraChargePressed),
		notebookCurrendPageChanged:        n.notebook.ConnectAfter("switch-page", n.CurrentPageChanged),
	}
}

func (n *Notebook) CurrentPageChanged() {
	switch n.notebook.GetCurrentPage() {
	case 0:
		n.BlockAdvertisementsTab.SelectedRowChanged()
	case 1:
		n.LineAdvertisementsTab.SelectedRowChanged()
	case 2:
		n.OrdersTab.SelectedRowChanged()
	}
}

func (n *Notebook) UpdateButtonPressed() {
	glib.IdleAddPriority(glib.PRIORITY_HIGH_IDLE, func() {
		n.req.LockReciever(true)
		n.lists.SetReplaceMode(false)
		n.SetSensetive(false)
		n.EnableSidebarsFilters(false)
		n.BlockAllSignals()
		n.DisableAllModels()
		n.ResetSorts()
		n.lists.SetReplaceMode(false)
	})

	glib.IdleAdd(func() {
		n.ClearAllListStores()
		n.advWin.UpdateLists()
	})

	glib.IdleAddPriority(glib.PRIORITY_DEFAULT_IDLE+1, func() {
		n.EnableAllModels()
		n.UnblockAllSignals()
		n.req.ActiveCostRate()
		n.advWin.SelectCostRate("")
		n.EnableSidebarsFilters(true)
		n.SetSensetive(true)
		n.lists.SetReplaceMode(true)
		n.req.LockReciever(false)
	})
}

func (n *Notebook) CreateAdvertisementsPressed() {
	logging.Logger.Debug("createAdvertisementsPressed")
	n.addEntryPopover.Popdown()
	n.application.CreateAddAdvertisementWindow().Show()
}

func (n *Notebook) CreateOrderPressed() {
	logging.Logger.Debug("createOrderPressed")
	n.addEntryPopover.Popdown()
	n.application.CreateAddOrderWindow().Show()

}

func (n *Notebook) CreateClientPressed() {
	logging.Logger.Debug("createClientPressed")
	n.addEntryPopover.Popdown()
	n.application.CreateAddClientWindow().Show()

}

func (n *Notebook) CreateTagPressed() {
	logging.Logger.Debug("createTagPressed")
	n.addEntryPopover.Popdown()
	n.application.CreateAddTagWindow().Show()
}

func (n *Notebook) CreateExtraChargePressed() {
	logging.Logger.Debug("createExtraChargePressed")
	n.addEntryPopover.Popdown()
	n.application.CreateAddExtraChargeWindow().Show()
}
