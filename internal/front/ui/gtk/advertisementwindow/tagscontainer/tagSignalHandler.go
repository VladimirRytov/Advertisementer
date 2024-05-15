package tagscontainer

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
)

type TagsTabHandler struct {
	treeSelectionChanged     glib.SignalHandle
	removeTagButtonClicked   glib.SignalHandle
	resetTagButtonClicked    glib.SignalHandle
	applyTagButtonClicked    glib.SignalHandle
	tagsListStoreRowDeleted  glib.SignalHandle
	tagsListStoreRowInserted glib.SignalHandle
	tagsListStoreRowChanged  glib.SignalHandle
	costChanged              glib.SignalHandle
}

func (tt *TagsTab) bindSignals() {
	tt.signalHandler = TagsTabHandler{
		treeSelectionChanged:     tt.treeSelection.Connect("changed", tt.SelectedRowChanged),
		removeTagButtonClicked:   tt.removeTagButton.Connect("clicked", tt.removeButtonPressed),
		resetTagButtonClicked:    tt.resetTagButton.Connect("clicked", tt.resetButtonPressed),
		applyTagButtonClicked:    tt.applyTagButton.Connect("clicked", tt.updateButtonPressed),
		tagsListStoreRowDeleted:  tt.TagsListStore.Connect("row-deleted", tt.rowChanged),
		tagsListStoreRowInserted: tt.TagsListStore.Connect("row-inserted", tt.rowChanged),
		tagsListStoreRowChanged:  tt.TagsListStore.Connect("row-changed", tt.rowChanged),
		costChanged:              tt.tagCostEntry.Connect("insert-text", tt.tools.CheckCostString),
	}

}

func (tt *TagsTab) SelectedRowChanged() {
	_, iter, ok := tt.treeSelection.GetSelected()
	if !ok {
		tt.resetRightSide()
		return
	}
	tagName, err := tt.tools.StringFromIter(iter, tt.TagsListStore.ToTreeModel(), 1)
	if err != nil {
		logging.Logger.Error("Cannot get TagName", "error message", err)
	}
	tagCost, err := tt.tools.StringFromIter(iter, tt.TagsListStore.ToTreeModel(), 2)
	if err != nil {
		logging.Logger.Error("Cannot get TagCost", "error message", err)
	}
	tt.tagNameEntry.SetText(tagName)
	tt.tagCostEntry.SetText(tagCost)
	tt.blockAdvsTreeFilter.Refilter()
	tt.lineAdvsTreeFilter.Refilter()
	tt.SetSensitive(true)
}

func (tt *TagsTab) resetRightSide() {
	tt.tagNameEntry.SetText("")
	tt.tagCostEntry.SetText("")
	tt.SetSensitive(false)
}

func (tt *TagsTab) removeButtonPressed() {
	tt.SetSensitive(false)
	go tt.req.RemoveTag(tt.SelectedName())
}

func (tt *TagsTab) resetButtonPressed() {
	tt.SelectedRowChanged()
}

func (tt *TagsTab) updateButtonPressed() {
	tt.SetSensitive(false)
	tag := &presenter.TagDTO{
		TagName: tt.SelectedName(),
		TagCost: tt.SelectedCost(),
	}
	err := tt.req.UpdateTag(tag)
	if err != nil {
		tt.SetSensitive(true)
		tt.advWin.ShowPopover(err.Error(), tt.tagCostEntry)
	}
}

func (tt *TagsTab) rowChanged() {
	logging.Logger.Debug("row changed")
	_, _, ok := tt.treeSelection.GetSelected()
	if !ok {
		tt.resetRightSide()
		return
	}
	ch := tt.TagsListStore.IterNChildren(nil)
	if ch == 0 {
		tt.ResetSidebar()
		tt.SetSensitive(false)
		return
	}
	tt.rightSide.SetSensitive(true)
}
