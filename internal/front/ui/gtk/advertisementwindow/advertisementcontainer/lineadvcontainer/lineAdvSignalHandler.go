package lineadvcontainer

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
)

type LineAdvertisementsSignalHandler struct {
	lineadvertisementsSelectorChanged      glib.SignalHandle
	deleteButtonClicked                    glib.SignalHandle
	resetButtonClicked                     glib.SignalHandle
	applyButtonClicked                     glib.SignalHandle
	lineadvertisementsListStoreRowDeleted  glib.SignalHandle
	lineadvertisementsListStoreRowInserted glib.SignalHandle
	lineadvertisementsListStoreRowChanged  glib.SignalHandle
}

func (line *LineAdvertisementsTab) bindSignals() {
	line.signalHandler = LineAdvertisementsSignalHandler{
		lineadvertisementsSelectorChanged:      line.lineadvertisementsSelector.Connect("changed", line.SelectedRowChanged),
		deleteButtonClicked:                    line.deleteButton.Connect("clicked", line.RemoveButtonPressed),
		resetButtonClicked:                     line.resetButton.Connect("clicked", line.ResetButtonPressed),
		applyButtonClicked:                     line.applyButton.Connect("clicked", line.UpdateButtonPressed),
		lineadvertisementsListStoreRowDeleted:  line.lineadvertisementsListStoreSort.Connect("row-deleted", line.RowChanged),
		lineadvertisementsListStoreRowInserted: line.lineadvertisementsListStoreSort.Connect("row-inserted", line.RowChanged),
		lineadvertisementsListStoreRowChanged:  line.lineadvertisementsListStoreSort.Connect("row-changed", line.RowChanged),
	}
}

func (line *LineAdvertisementsTab) SelectedRowChanged() {
	iModel, iter, ok := line.lineadvertisementsSelector.GetSelected()
	if !ok {
		line.lineAdvForm.Reset()
		line.SetSensitive(false)
		return
	}

	model := iModel.ToTreeModel()
	id, err := line.tools.InterfaceFromIter(iter, model, 1)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	orderID, err := line.tools.InterfaceFromIter(iter, model, 2)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	releaseCount, err := line.tools.InterfaceFromIter(iter, model, 3)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	releaseDates, err := line.tools.StringFromIter(iter, model, 5)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	tags, err := line.tools.StringFromIter(iter, model, 6)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	extraCharges, err := line.tools.StringFromIter(iter, model, 7)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	cost, err := line.tools.StringFromIter(iter, model, 8)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	comment, err := line.tools.StringFromIter(iter, model, 9)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	selectedLine := &presenter.LineAdvertisementDTO{
		Advertisement: presenter.Advertisement{
			ID:           id.(int),
			OrderID:      orderID.(int),
			ReleaseCount: releaseCount.(int),
			ReleaseDates: releaseDates,
			Cost:         cost,
			Text:         comment,
			Tags:         tags,
			ExtraCharge:  extraCharges,
		},
	}
	line.lineAdvForm.FillData(selectedLine)
	line.lineAdvForm.SetSensetive(true)
	line.SetSensitive(true)
}

func (line *LineAdvertisementsTab) RemoveButtonPressed() {
	line.SetSensitive(false)
	lineAdv := line.lineAdvForm.FetchData()
	line.req.RemoveLineAdvertisement(&lineAdv)
}

func (line *LineAdvertisementsTab) ResetButtonPressed() {
	line.SelectedRowChanged()
}

func (line *LineAdvertisementsTab) UpdateButtonPressed() {
	line.SetSensitive(false)
	lineAdv := line.lineAdvForm.FetchData()
	err := line.req.UpdateLineAdvertisement(&lineAdv)
	if err != nil {
		line.errWin.NewErrorWindow(err)
	}
}

func (line *LineAdvertisementsTab) RowChanged() {
	logging.Logger.Debug("row changed")
	_, _, ok := line.lineadvertisementsSelector.GetSelected()
	if !ok {
		line.lineAdvForm.Reset()
		line.SetSensitive(false)
		return
	}
	if line.lineadvertisementsListStoreFilter.IterNChildren(nil) == 0 {
		line.lineAdvForm.Reset()
		line.SetSensitive(false)
		return
	}
	line.lineAdvForm.SetSensetive(true)
}
