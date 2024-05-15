package tagscontainer

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"

	"github.com/gotk3/gotk3/gtk"
)

type Tools interface {
	CheckCostString(self *gtk.Entry, new string)
	CompareCosts(model *gtk.TreeModel, a, b *gtk.TreeIter, col int) int
	CompareStrings(model *gtk.TreeModel, a, b *gtk.TreeIter, col int) int
	StringFromIter(*gtk.TreeIter, *gtk.TreeModel, int) (string, error)
}

type ListStores interface {
	TagsList() *gtk.ListStore
	BlockAdvertisementsList() *gtk.ListStore
	LineAdvertisementsList() *gtk.ListStore
}

type TagRequests interface {
	UpdateTag(*presenter.TagDTO) error
	RemoveTag(string)
	ValueInString(string, string) bool
}

type AdvWin interface {
	ShowPopover(string, gtk.IWidget)
}

type TagsTab struct {
	enableFilters bool
	req           TagRequests
	tools         Tools
	lists         ListStores
	advWin        AdvWin

	TagsListStore *gtk.TreeModelSort
	treeView      *gtk.TreeView
	treeSelection *gtk.TreeSelection

	rightSide    *gtk.Box
	tagNameLabel *gtk.Label
	tagNameEntry *gtk.Entry

	tagCostLabel *gtk.Label
	tagCostEntry *gtk.Entry

	removeTagButton *gtk.Button
	resetTagButton  *gtk.Button
	applyTagButton  *gtk.Button

	blockAdvsLabel      *gtk.Label
	blockAdvsTreeView   *gtk.TreeView
	blockAdvsTreeFilter *gtk.TreeModelFilter

	lineAdvsLabel      *gtk.Label
	lineAdvsTreeView   *gtk.TreeView
	lineAdvsTreeFilter *gtk.TreeModelFilter

	signalHandler TagsTabHandler
}

func Create(bldFile *builder.Builder, tools Tools, lists ListStores, advWin AdvWin, req TagRequests) *TagsTab {
	var err error
	tagTab := new(TagsTab)
	tagTab.req = req
	tagTab.lists = lists
	tagTab.tools = tools
	tagTab.advWin = advWin
	tagTab.build(bldFile)
	tagTab.TagsListStore, err = gtk.TreeModelSortNew(tagTab.lists.TagsList())
	if err != nil {
		panic(err)
	}
	tagTab.TagsListStore.SetSortFunc(1, func(model *gtk.TreeModel, a, b *gtk.TreeIter) int {
		return tools.CompareStrings(model, a, b, 1)
	})
	tagTab.TagsListStore.SetSortFunc(2, func(model *gtk.TreeModel, a, b *gtk.TreeIter) int {
		return tools.CompareCosts(model, a, b, 2)
	})

	tagTab.treeSelection, err = tagTab.treeView.GetSelection()
	if err != nil {
		panic(err)
	}

	tagTab.blockAdvsTreeFilter, err = tagTab.lists.BlockAdvertisementsList().FilterNew(nil)
	if err != nil {
		panic(err)
	}
	tagTab.blockAdvsTreeFilter.SetVisibleFunc(tagTab.filterAdvertisements)
	tagTab.blockAdvsTreeView.SetModel(tagTab.blockAdvsTreeFilter)
	tagTab.lineAdvsTreeFilter, err = tagTab.lists.LineAdvertisementsList().FilterNew(nil)
	if err != nil {
		panic(err)
	}
	tagTab.lineAdvsTreeFilter.SetVisibleFunc(tagTab.filterAdvertisements)
	tagTab.lineAdvsTreeView.SetModel(tagTab.lineAdvsTreeFilter)
	tagTab.bindSignals()
	tagTab.BlockSignals()
	tagTab.SetSensitive(false)
	return tagTab
}

func (tt *TagsTab) build(bldFile *builder.Builder) {
	tt.treeView = bldFile.FetchTreeView("TagsTreeView")
	tt.rightSide = bldFile.FetchBox("TagSidebar")
	tt.tagNameLabel = bldFile.FetchLabel("TagNameLabel")
	tt.tagNameEntry = bldFile.FetchEntry("TagNameEntry")
	tt.tagCostLabel = bldFile.FetchLabel("TagCostLabel")
	tt.tagCostEntry = bldFile.FetchEntry("TagCostEntry")
	tt.removeTagButton = bldFile.FetchButton("TagDeleteButton")
	tt.resetTagButton = bldFile.FetchButton("TagResetButton")
	tt.applyTagButton = bldFile.FetchButton("TagApplyButton")
	tt.blockAdvsLabel = bldFile.FetchLabel("TagBlockAdvLabel")
	tt.blockAdvsTreeView = bldFile.FetchTreeView("TagBlockAdvTreeView")
	tt.lineAdvsLabel = bldFile.FetchLabel("TagLineAdvLabel")
	tt.lineAdvsTreeView = bldFile.FetchTreeView("TagLineAdvTreeView")
}
