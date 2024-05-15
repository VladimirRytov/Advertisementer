package tagscontainer

import (
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/gtk"
)

func (tt *TagsTab) SelectedName() string {
	return tt.tagNameEntry.GetLayout().GetText()
}

func (tt *TagsTab) SetSelectedName(name string) {
	tt.tagNameEntry.SetText(name)
}

func (tt *TagsTab) SelectedCost() string {
	return tt.tagCostEntry.GetLayout().GetText()
}

func (tt *TagsTab) SetSelectedCost(cost string) {
	tt.tagCostEntry.SetText(cost)
}

func (tt *TagsTab) UpdateLinks() {
	tt.blockAdvsTreeFilter.Refilter()
	tt.lineAdvsTreeFilter.Refilter()
}

func (tt *TagsTab) SetSensitive(lock bool) {
	tt.tagCostEntry.SetSensitive(lock)
	tt.removeTagButton.SetSensitive(lock)
	tt.resetTagButton.SetSensitive(lock)
	tt.applyTagButton.SetSensitive(lock)
}

func (tt *TagsTab) ResetSidebar() {
	tt.tagCostEntry.SetText("")
	tt.tagNameEntry.SetText("")
}

func (tt *TagsTab) filterAdvertisements(model *gtk.TreeModel, iter *gtk.TreeIter) bool {
	if !tt.enableFilters {
		return false
	}

	val, err := tt.tools.StringFromIter(iter, model, 6)
	if err != nil {
		logging.Logger.Warn("tagFilterAdvertisements: got empty string. Skipping", "error", err)
		return false
	}
	return tt.req.ValueInString(val, tt.SelectedName())
}

func (tt *TagsTab) BlockSignals() {
	tt.treeSelection.HandlerBlock(tt.signalHandler.treeSelectionChanged)
	tt.removeTagButton.HandlerBlock(tt.signalHandler.removeTagButtonClicked)
	tt.resetTagButton.HandlerBlock(tt.signalHandler.resetTagButtonClicked)
	tt.applyTagButton.HandlerBlock(tt.signalHandler.applyTagButtonClicked)
	tt.TagsListStore.HandlerBlock(tt.signalHandler.tagsListStoreRowDeleted)
	tt.TagsListStore.HandlerBlock(tt.signalHandler.tagsListStoreRowInserted)
	tt.TagsListStore.HandlerBlock(tt.signalHandler.tagsListStoreRowChanged)
	tt.tagCostEntry.HandlerBlock(tt.signalHandler.costChanged)
}

func (tt *TagsTab) UnblockSignals() {
	tt.treeSelection.HandlerUnblock(tt.signalHandler.treeSelectionChanged)
	tt.removeTagButton.HandlerUnblock(tt.signalHandler.removeTagButtonClicked)
	tt.resetTagButton.HandlerUnblock(tt.signalHandler.resetTagButtonClicked)
	tt.applyTagButton.HandlerUnblock(tt.signalHandler.applyTagButtonClicked)
	tt.TagsListStore.HandlerUnblock(tt.signalHandler.tagsListStoreRowDeleted)
	tt.TagsListStore.HandlerUnblock(tt.signalHandler.tagsListStoreRowInserted)
	tt.TagsListStore.HandlerUnblock(tt.signalHandler.tagsListStoreRowChanged)
	tt.tagCostEntry.HandlerUnblock(tt.signalHandler.costChanged)
}

func (tt *TagsTab) SetEnableFilters(enable bool) {
	tt.enableFilters = enable
}

func (tt *TagsTab) AttachList() {
	tt.treeView.SetModel(tt.TagsListStore)
	tt.blockAdvsTreeView.SetModel(tt.blockAdvsTreeFilter)
	tt.lineAdvsTreeView.SetModel(tt.lineAdvsTreeFilter)
	tt.SetEnableFilters(true)
}

func (tt *TagsTab) DetachList() {
	tt.treeView.SetModel(nil)
	tt.blockAdvsTreeView.SetModel(nil)
	tt.lineAdvsTreeView.SetModel(nil)
	tt.SetEnableFilters(false)
}

func (tt *TagsTab) UnsetModel() {
	tt.resetRightSide()
	bloks, _ := tt.blockAdvsTreeView.GetSelection()
	bloks.UnselectAll()
	lines, _ := tt.lineAdvsTreeView.GetSelection()
	lines.UnselectAll()

	tt.treeSelection.UnselectAll()
	tt.treeView.Hide()
	tt.DetachList()
}

func (tt *TagsTab) SetModel() {
	tt.AttachList()
	tt.treeView.Show()
}

func (tt *TagsTab) ResetSort() {
	tt.TagsListStore.ResetDefaultSortFunc()
}
