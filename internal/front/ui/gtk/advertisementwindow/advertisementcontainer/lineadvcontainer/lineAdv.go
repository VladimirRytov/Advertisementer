package lineadvcontainer

import (
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/gtk"
)

func (line *LineAdvertisementsTab) SetSensitive(locked bool) {
	line.lineAdvForm.SetSensetive(locked)
	line.deleteButton.SetSensitive(locked)
	line.resetButton.SetSensitive(locked)
	line.applyButton.SetSensitive(locked)
	line.lineAdvForm.SetSensetive(locked)
}

func (line *LineAdvertisementsTab) BlockSignals() {
	line.lineadvertisementsSelector.HandlerBlock(line.signalHandler.lineadvertisementsSelectorChanged)
	line.deleteButton.HandlerBlock(line.signalHandler.deleteButtonClicked)
	line.resetButton.HandlerBlock(line.signalHandler.resetButtonClicked)
	line.applyButton.HandlerBlock(line.signalHandler.applyButtonClicked)
	line.lineadvertisementsListStoreSort.HandlerBlock(line.signalHandler.lineadvertisementsListStoreRowDeleted)
	line.lineadvertisementsListStoreSort.HandlerBlock(line.signalHandler.lineadvertisementsListStoreRowInserted)
	line.lineadvertisementsListStoreSort.HandlerBlock(line.signalHandler.lineadvertisementsListStoreRowChanged)
}

func (line *LineAdvertisementsTab) UnblockSignals() {
	line.lineadvertisementsSelector.HandlerUnblock(line.signalHandler.lineadvertisementsSelectorChanged)
	line.deleteButton.HandlerUnblock(line.signalHandler.deleteButtonClicked)
	line.resetButton.HandlerUnblock(line.signalHandler.resetButtonClicked)
	line.applyButton.HandlerUnblock(line.signalHandler.applyButtonClicked)
	line.lineadvertisementsListStoreSort.HandlerUnblock(line.signalHandler.lineadvertisementsListStoreRowDeleted)
	line.lineadvertisementsListStoreSort.HandlerUnblock(line.signalHandler.lineadvertisementsListStoreRowInserted)
	line.lineadvertisementsListStoreSort.HandlerUnblock(line.signalHandler.lineadvertisementsListStoreRowChanged)
}

func (line *LineAdvertisementsTab) Refilter() {
	line.lineadvertisementsListStoreFilter.Refilter()
}

func (line *LineAdvertisementsTab) filterList(model *gtk.TreeModel, iter *gtk.TreeIter) bool {
	if !line.lineadvertisementsListStoreFilterEnable {
		return true
	}
	releaseDates, err := line.tools.StringFromIter(iter, model.ToTreeModel(), 5)
	if err != nil {
		logging.Logger.Error("lineAdvertisementsTab.filterList: cannot get releaseDates from listStore", "error", err)
		return false
	}
	return line.req.ReleasesInTimeRange(releaseDates,
		line.dates.FromDate(), line.dates.ToDate())
}

func (line *LineAdvertisementsTab) SetEnableFilters(enable bool) {
	line.lineadvertisementsListStoreFilterEnable = enable
}

func (line *LineAdvertisementsTab) AttachList() {
	line.lineAdvertisementsTreeView.SetModel(line.lineadvertisementsListStoreSort)
}

func (line *LineAdvertisementsTab) DetachList() {
	line.lineAdvertisementsTreeView.SetModel(nil)
}

func (line *LineAdvertisementsTab) UnsetModel() {
	line.SetSensitive(false)
	line.lineadvertisementsSelector.UnselectAll()
	line.lineAdvertisementsTreeView.Hide()
	line.DetachList()
	line.lineAdvForm.UnsetModel()
}

func (line *LineAdvertisementsTab) SetModel() {
	line.AttachList()
	line.lineAdvertisementsTreeView.Show()
	line.lineAdvForm.SetModel()
}

func (line *LineAdvertisementsTab) ResetSort() {
	line.lineadvertisementsListStoreSort.ResetDefaultSortFunc()
}
