package blockadvcontainer

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
)

type BlockAdvertisementsSignalHandler struct {
	blockadvertisementsSelectorChanged      glib.SignalHandle
	deleteButtonClicked                     glib.SignalHandle
	resetButtonClicked                      glib.SignalHandle
	applyButtonClicked                      glib.SignalHandle
	blockadvertisementsListStoreRowDeleted  glib.SignalHandle
	blockadvertisementsListStoreRowInserted glib.SignalHandle
	blockadvertisementsListStoreRowChanged  glib.SignalHandle
}

func (block *BlockAdvertisementsTab) bindSignals() {
	block.signalHandler = BlockAdvertisementsSignalHandler{
		blockadvertisementsSelectorChanged:      block.blockadvertisementsSelector.Connect("changed", block.SelectedRowChanged),
		deleteButtonClicked:                     block.deleteButton.Connect("clicked", block.RemoveButtonPressed),
		resetButtonClicked:                      block.resetButton.Connect("clicked", block.ResetButtonPressed),
		applyButtonClicked:                      block.applyButton.Connect("clicked", block.UpdateButtonPressed),
		blockadvertisementsListStoreRowDeleted:  block.blockadvertisementsListStoreSort.Connect("row-deleted", block.RowChanged),
		blockadvertisementsListStoreRowInserted: block.blockadvertisementsListStoreSort.Connect("row-inserted", block.RowChanged),
		blockadvertisementsListStoreRowChanged:  block.blockadvertisementsListStoreSort.Connect("row-changed", block.RowChanged),
	}
}

func (block *BlockAdvertisementsTab) SelectedRowChanged() {
	iModel, iter, ok := block.blockadvertisementsSelector.GetSelected()
	if !ok {
		block.blockAdvForm.Reset()
		block.blockAdvForm.SetSensetive(false)
		block.SetSensitive(false)
		return
	}
	block.blockAdvForm.SetSensetive(true)
	block.SetSensitive(true)
	model := iModel.ToTreeModel()
	id, err := block.tools.InterfaceFromIter(iter, model, 1)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	orderID, err := block.tools.InterfaceFromIter(iter, model, 2)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	releaseCount, err := block.tools.InterfaceFromIter(iter, model, 3)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	releaseDates, err := block.tools.StringFromIter(iter, model, 5)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	tags, err := block.tools.StringFromIter(iter, model, 6)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	extraCharges, err := block.tools.StringFromIter(iter, model, 7)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	size, err := block.tools.InterfaceFromIter(iter, model, 8)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	cost, err := block.tools.StringFromIter(iter, model, 9)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	comment, err := block.tools.StringFromIter(iter, model, 10)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	fileName, err := block.tools.StringFromIter(iter, model, 11)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}
	selectedBlock := &presenter.BlockAdvertisementDTO{
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
		Size:     size.(int),
		FileName: fileName,
	}
	block.blockAdvForm.FillData(selectedBlock)
}

func (block *BlockAdvertisementsTab) RemoveButtonPressed() {
	selected := block.blockAdvForm.FetchData()
	block.SetSensitive(false)
	block.req.RemoveBlockAdvertisement(&selected)
}

func (block *BlockAdvertisementsTab) ResetButtonPressed() {
	block.SelectedRowChanged()
}

func (block *BlockAdvertisementsTab) UpdateButtonPressed() {
	block.SetSensitive(false)
	selected := block.blockAdvForm.FetchData()
	err := block.req.UpdateBlockAdvertisement(&selected)
	if err != nil {
		block.errWin.NewErrorWindow(err)
	}
}

func (block *BlockAdvertisementsTab) RowChanged() {
	_, _, ok := block.blockadvertisementsSelector.GetSelected()
	if !ok {
		block.blockAdvForm.Reset()
		block.blockAdvForm.SetSensetive(false)
		block.SetSensitive(false)
		return
	}
	ch := block.blockadvertisementsListStoreFilter.IterNChildren(nil)
	if ch == 0 {
		block.blockAdvForm.Reset()
		block.blockAdvForm.SetSensetive(false)
		block.SetSensitive(false)
		return
	}
	block.SetSensitive(true)
	block.blockAdvForm.SetSensetive(true)
}
