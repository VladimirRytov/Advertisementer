package neworder

import (
	"slices"
	"strconv"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/advertisementwindow/advertisementform/blockform"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/advertisementwindow/advertisementform/lineform"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type signalHandler struct {
	addExistingAdvertisementCheckButtonToggled glib.SignalHandle
	clientSelectToggleToggled                  glib.SignalHandle
	blockAdvTreeToggleToggled                  glib.SignalHandle
	lineAdvTreeViewToggled                     glib.SignalHandle
	addNewAdvertisementCheckButtonToggled      glib.SignalHandle
	addNewBlockAdvertisementButtonclicked      glib.SignalHandle
	addNewLineAdvertisementButtonclicked       glib.SignalHandle
	removeNewAdvertisementButtonclicked        glib.SignalHandle
	newClientButtonClicked                     glib.SignalHandle
	notebookPageAdded                          glib.SignalHandle
	notebookPageRemoved                        glib.SignalHandle
	applyButtonClicked                         glib.SignalHandle
	cancelButtonClicked                        glib.SignalHandle
	costChanged                                glib.SignalHandle
	windowDestroyed                            glib.SignalHandle
	calculateCostPressed                       glib.SignalHandle
}

func (now *NewOrderWindow) bindSignals() {
	now.signalHandler.addExistingAdvertisementCheckButtonToggled = now.selectExistedCheckButton.Connect("toggled", now.addExistingAdvertisementCheckButtonToggled)
	now.signalHandler.clientSelectToggleToggled = now.clientSelectToggle.Connect("toggled", now.ClientsToggled)
	now.signalHandler.blockAdvTreeToggleToggled = now.blockAdvTreeToggle.Connect("toggled", now.BlockAdvToggled)
	now.signalHandler.lineAdvTreeViewToggled = now.lineAdvTreeViewToggle.Connect("toggled", now.LineAdvToggled)
	now.signalHandler.addNewAdvertisementCheckButtonToggled = now.newAdvertisementCheckButton.Connect("toggled", now.addNewAdvertisementCheckButtonToggled)

	now.signalHandler.addNewBlockAdvertisementButtonclicked = now.addBlockAdvButton.Connect("clicked", now.addNewBlockAdvertisementButtonclicked)
	now.signalHandler.addNewLineAdvertisementButtonclicked = now.addLineAdvButton.Connect("clicked", now.addNewLineAdvertisementButtonclicked)
	now.signalHandler.removeNewAdvertisementButtonclicked = now.deleteAdvButton.Connect("clicked", now.removeNewAdvertisementButtonclicked)
	now.signalHandler.newClientButtonClicked = now.newClientButton.Connect("clicked", now.newClientButtonClicked)

	now.signalHandler.notebookPageAdded = now.newAdvertisementNotebook.Connect("page-added", now.notebookPageAdded)
	now.signalHandler.notebookPageRemoved = now.newAdvertisementNotebook.Connect("page-removed", now.notebookPageRemoved)
	now.signalHandler.applyButtonClicked = now.createButton.Connect("clicked", now.applyButtonClicked)
	now.signalHandler.cancelButtonClicked = now.cancelButton.Connect("clicked", now.cancelButtonClicked)
	now.signalHandler.costChanged = now.costEntry.Connect("insert-text", now.tools.CheckCostString)
	now.signalHandler.windowDestroyed = now.window.Connect("destroy", now.winDestroyed)
	now.signalHandler.calculateCostPressed = now.calculateCostButton.Connect("clicked", now.calculateCost)
}

func (now *NewOrderWindow) winDestroyed() {
	for i := range now.advertisements {
		switch adv := now.advertisements[i].(type) {
		case *blockform.BlockAdvPage:
			adv.Destroy()
			now.advertisements[i] = nil
		case *lineform.LineAdvPage:
			adv.Destroy()
			now.advertisements[i] = nil
		}
	}
	now.window.Destroy()
}

func (now *NewOrderWindow) newClientButtonClicked() {
	clientWindow := now.app.CreateAddClientWindow()
	clientWindow.Show()
}

func (now *NewOrderWindow) addExistingAdvertisementCheckButtonToggled() {
	now.appendExistRevealer.SetRevealChild(now.selectExistedCheckButton.GetActive())
}

func (now *NewOrderWindow) addNewAdvertisementCheckButtonToggled() {
	now.newAdvertisementRevealer.SetRevealChild(now.newAdvertisementCheckButton.GetActive())
}

func (now *NewOrderWindow) addNewBlockAdvertisementButtonclicked() {
	blockPage := now.formMaker.NewBlockCopyForm()
	blockPage.SetNewNestedAdvMode(true)
	now.advertisements = append(now.advertisements, blockPage)
	label, _ := gtk.LabelNew("Блоч." + strconv.Itoa(len(now.advertisements)))
	now.newAdvertisementNotebook.AppendPage(blockPage.Widget(), label)
	now.newAdvertisementNotebook.SetCurrentPage(now.newAdvertisementNotebook.GetNPages() - 1)
	now.newAdvPopover.Popdown()
}

func (now *NewOrderWindow) addNewLineAdvertisementButtonclicked() {
	linePage := now.formMaker.NewLineCopyForm()
	linePage.SetNewNestedAdvMode(true)
	now.advertisements = append(now.advertisements, linePage)
	label, _ := gtk.LabelNew("Стр." + strconv.Itoa(len(now.advertisements)))
	now.newAdvertisementNotebook.AppendPage(linePage.Widget(), label)
	now.newAdvertisementNotebook.SetCurrentPage(now.newAdvertisementNotebook.GetNPages() - 1)
	now.newAdvPopover.Popdown()
}

func (now *NewOrderWindow) removeNewAdvertisementButtonclicked() {
	val := now.newAdvertisementNotebook.GetCurrentPage()
	now.newAdvertisementNotebook.RemovePage(val)
	now.advertisements = slices.Delete(now.advertisements, val, val+1)
}

func (now *NewOrderWindow) cancelButtonClicked() {
	now.window.Destroy()
}

func (now *NewOrderWindow) applyButtonClicked() {
	var (
		bloks []presenter.BlockAdvertisementDTO = make([]presenter.BlockAdvertisementDTO, 0, 5)
		lines []presenter.LineAdvertisementDTO  = make([]presenter.LineAdvertisementDTO, 0, 5)
	)
	order := presenter.OrderDTO{
		ID:            0,
		ClientName:    now.client(),
		Cost:          now.cost(),
		PaymentType:   now.paymentType(),
		PaymentStatus: now.paymentStatus(),
	}
	if now.selectExistedCheckButton.GetActive() {
		bloks = append(bloks, now.CollectExistingBlockAdvertisements()...)
		lines = append(lines, now.CollectExistingLineAdvertisements()...)
	}
	if now.newAdvertisementCheckButton.GetActive() {
		newBloks, newLines := now.CollectNewAdvertisements()
		bloks = append(bloks, newBloks...)
		lines = append(lines, newLines...)
	}
	err := now.req.CreateOrder(&order, bloks, lines)
	if err != nil {
		now.app.NewErrorWindow(err)
	}
}

func (now *NewOrderWindow) CollectNewAdvertisements() ([]presenter.BlockAdvertisementDTO, []presenter.LineAdvertisementDTO) {
	var (
		blocks []presenter.BlockAdvertisementDTO = make([]presenter.BlockAdvertisementDTO, 0, 5)
		lines  []presenter.LineAdvertisementDTO  = make([]presenter.LineAdvertisementDTO, 0, 5)
	)
	for i := range now.advertisements {
		switch adv := now.advertisements[i].(type) {
		case *blockform.BlockAdvPage:
			blocks = append(blocks, adv.FetchData())
		case *lineform.LineAdvPage:
			lines = append(lines, adv.FetchData())
		}
	}
	return blocks, lines
}

func (now *NewOrderWindow) notebookPageAdded() {
	if now.newAdvertisementNotebook.GetNPages() > 0 {
		now.newAdvertisementNotebook.Show()
		now.deleteAdvButton.SetSensitive(true)
	}
}

func (now *NewOrderWindow) notebookPageRemoved() {
	if now.newAdvertisementNotebook.GetNPages() == 0 {
		now.newAdvertisementNotebook.Hide()
		now.deleteAdvButton.SetSensitive(false)
	}
}

func (now *NewOrderWindow) ClientsToggled(self *gtk.CellRendererToggle, pathStr string) {
	now.clientListStore.ForEach(func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
		now.clientListStore.SetValue(iter, 0, path.String() == pathStr)
		return false
	})
}

func (now *NewOrderWindow) BlockAdvToggled(self *gtk.CellRendererToggle, pathStr string) {

	iter, err := now.blockAdvListStore.GetIterFromString(pathStr)
	if err != nil {
		logging.Logger.Error("extraChargeToggled: cannot get iter from string", "error", err)
	}

	rawVal, err := now.tools.InterfaceFromIter(iter, &now.blockAdvListStore.TreeModel, 0)
	if err != nil {
		logging.Logger.Error("extraChargeToggled: cannot get val from liststore", "error", err)
	}

	if val, ok := rawVal.(bool); ok {
		now.blockAdvListStore.SetValue(iter, 0, !val)
	} else {
		logging.Logger.Error("extraChargeToggled: got not boolean value")
	}
}

func (now *NewOrderWindow) LineAdvToggled(self *gtk.CellRendererToggle, pathStr string) {

	iter, err := now.lineAdvListStore.GetIterFromString(pathStr)
	if err != nil {
		logging.Logger.Error("extraChargeToggled: cannot get iter from string", "error", err)
	}

	rawVal, err := now.tools.InterfaceFromIter(iter, &now.lineAdvListStore.TreeModel, 0)
	if err != nil {
		logging.Logger.Error("extraChargeToggled: cannot get val from liststore", "error", err)
	}

	if val, ok := rawVal.(bool); ok {
		now.lineAdvListStore.SetValue(iter, 0, !val)
	} else {
		logging.Logger.Error("extraChargeToggled: got not boolean value")
	}
}

func (now *NewOrderWindow) calculateCost() {
	b, err := now.costEntry.GetBuffer()
	if err != nil {
		logging.Logger.Error("newOrderWindow.calculateCost: cannot get cost entry buffer", "error", err)
		return
	}
	now.app.RegisterReciever(b)
	var (
		bloks []presenter.BlockAdvertisementDTO = make([]presenter.BlockAdvertisementDTO, 0, 5)
		lines []presenter.LineAdvertisementDTO  = make([]presenter.LineAdvertisementDTO, 0, 5)
	)
	order := presenter.OrderDTO{
		ID:            0,
		ClientName:    now.client(),
		Cost:          now.cost(),
		PaymentType:   now.paymentType(),
		PaymentStatus: now.paymentStatus(),
	}
	if now.selectExistedCheckButton.GetActive() {
		bloks = append(bloks, now.CollectExistingBlockAdvertisements()...)
		lines = append(lines, now.CollectExistingLineAdvertisements()...)
	}
	if now.newAdvertisementCheckButton.GetActive() {
		newBloks, newLines := now.CollectNewAdvertisements()
		bloks = append(bloks, newBloks...)
		lines = append(lines, newLines...)
	}
	err = now.req.CalculateOrderCost(&order, bloks, lines)
	if err != nil {
		now.app.NewErrorWindow(err)
	}
}
