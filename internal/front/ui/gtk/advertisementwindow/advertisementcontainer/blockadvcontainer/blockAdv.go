package blockadvcontainer

import (
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/gtk"
)

func (blk *BlockAdvertisementsTab) SetSensitive(s bool) {
	blk.deleteButton.SetSensitive(s)
	blk.resetButton.SetSensitive(s)
	blk.applyButton.SetSensitive(s)
	blk.blockAdvForm.SetSensetive(s)
}

func (blk *BlockAdvertisementsTab) BlockSignals() {
	blk.blockadvertisementsSelector.HandlerBlock(blk.signalHandler.blockadvertisementsSelectorChanged)
	blk.deleteButton.HandlerBlock(blk.signalHandler.deleteButtonClicked)
	blk.resetButton.HandlerBlock(blk.signalHandler.resetButtonClicked)
	blk.applyButton.HandlerBlock(blk.signalHandler.applyButtonClicked)
	blk.blockadvertisementsListStoreSort.HandlerBlock(blk.signalHandler.blockadvertisementsListStoreRowDeleted)
	blk.blockadvertisementsListStoreSort.HandlerBlock(blk.signalHandler.blockadvertisementsListStoreRowInserted)
	blk.blockadvertisementsListStoreSort.HandlerBlock(blk.signalHandler.blockadvertisementsListStoreRowChanged)
}

func (blk *BlockAdvertisementsTab) UnblockSignals() {
	blk.blockadvertisementsSelector.HandlerUnblock(blk.signalHandler.blockadvertisementsSelectorChanged)
	blk.deleteButton.HandlerUnblock(blk.signalHandler.deleteButtonClicked)
	blk.resetButton.HandlerUnblock(blk.signalHandler.resetButtonClicked)
	blk.applyButton.HandlerUnblock(blk.signalHandler.applyButtonClicked)
	blk.blockadvertisementsListStoreSort.HandlerUnblock(blk.signalHandler.blockadvertisementsListStoreRowDeleted)
	blk.blockadvertisementsListStoreSort.HandlerUnblock(blk.signalHandler.blockadvertisementsListStoreRowInserted)
	blk.blockadvertisementsListStoreSort.HandlerUnblock(blk.signalHandler.blockadvertisementsListStoreRowChanged)
}

func (blk *BlockAdvertisementsTab) Refilter() {
	blk.blockadvertisementsListStoreFilter.Refilter()
}

func (blk *BlockAdvertisementsTab) SetEnableFilters(enable bool) {
	blk.blockadvertisementsListStoreFilterEnabled = enable
}

func (blk *BlockAdvertisementsTab) filterList(model *gtk.TreeModel, iter *gtk.TreeIter) bool {
	if !blk.blockadvertisementsListStoreFilterEnabled {
		return true
	}
	releaseDates, err := blk.tools.StringFromIter(iter, model.ToTreeModel(), 5)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.filterList: cannot get releaseDates from listStore", "error", err)
		return false
	}
	return blk.req.ReleasesInTimeRange(releaseDates,
		blk.dateRanges.FromDate(), blk.dateRanges.ToDate())
}

func (blk *BlockAdvertisementsTab) AttachList() {
	blk.blockAdvertisementsTreeView.SetModel(blk.blockadvertisementsListStoreSort)
}

func (blk *BlockAdvertisementsTab) DetachList() {
	blk.blockAdvertisementsTreeView.SetModel(nil)
}

func (blk *BlockAdvertisementsTab) UnsetModel() {
	blk.SetSensitive(false)
	blk.DetachList()
	blk.blockadvertisementsSelector.UnselectAll()
	blk.blockAdvertisementsTreeView.Hide()
	blk.blockAdvForm.UnsetModel()
}

func (blk *BlockAdvertisementsTab) SetModel() {
	blk.AttachList()
	blk.blockAdvertisementsTreeView.Show()
	blk.blockAdvForm.SetModel()
}

func (blk *BlockAdvertisementsTab) ResetSort() {
	blk.blockadvertisementsListStoreSort.ResetDefaultSortFunc()
}
