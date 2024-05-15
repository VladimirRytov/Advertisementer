package extrachargescontainer

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"

	"github.com/gotk3/gotk3/gtk"
)

type Tools interface {
	CompareStrings(*gtk.TreeModel, *gtk.TreeIter, *gtk.TreeIter, int) int
	CompareCosts(*gtk.TreeModel, *gtk.TreeIter, *gtk.TreeIter, int) int
	CheckExtraChargeString(self *gtk.Entry, new string)
	StringFromIter(*gtk.TreeIter, *gtk.TreeModel, int) (string, error)
}

type AdvWin interface {
	ShowPopover(string, gtk.IWidget)
}

type Requests interface {
	ValueInString(string, string) bool
	UpdateExtraCharge(*presenter.ExtraChargeDTO) error
	RemoveExtraCharge(string)
}

type ListStores interface {
	BlockAdvertisementsList() *gtk.ListStore
	LineAdvertisementsList() *gtk.ListStore
	ExtraChargesList() *gtk.ListStore
}

type ExtraChargeTab struct {
	tools  Tools
	advWin AdvWin
	lists  ListStores
	req    Requests

	enableFilters        bool
	treeView             *gtk.TreeView
	ExtraCharges         *gtk.TreeModelSort
	treeSelector         *gtk.TreeSelection
	rightSide            *gtk.Box
	extraChargeNameLabel *gtk.Label
	extraChargeNameEntry *gtk.Entry

	extraChargeMultiplierLabel *gtk.Label
	extraChargeMultiplierEntry *gtk.Entry

	removeExtraChargeButton *gtk.Button
	resetExtraChargeButton  *gtk.Button
	applyExtraChargeButton  *gtk.Button

	blockAdvsLabel      *gtk.Label
	blockAdvsTreeView   *gtk.TreeView
	blockAdvsTreeFilter *gtk.TreeModelFilter

	lineAdvsLabel      *gtk.Label
	lineAdvsTreeView   *gtk.TreeView
	lineAdvsTreeFilter *gtk.TreeModelFilter

	signalHandler ExtraChargeSignalHander
}

func Create(bldFile *builder.Builder, reqGate Requests, advWin AdvWin, tools Tools, lists ListStores) *ExtraChargeTab {
	var err error
	extraChargeTab := new(ExtraChargeTab)
	extraChargeTab.tools = tools
	extraChargeTab.advWin = advWin
	extraChargeTab.req = reqGate
	extraChargeTab.lists = lists
	extraChargeTab.build(bldFile)
	extraChargeTab.ExtraCharges, err = gtk.TreeModelSortNew(lists.ExtraChargesList())
	if err != nil {
		panic(err)
	}
	extraChargeTab.ExtraCharges.SetSortFunc(1, func(model *gtk.TreeModel, a, b *gtk.TreeIter) int {
		return tools.CompareStrings(model, a, b, 1)
	})
	extraChargeTab.ExtraCharges.SetSortFunc(2, func(model *gtk.TreeModel, a, b *gtk.TreeIter) int {
		return tools.CompareCosts(model, a, b, 2)
	})
	extraChargeTab.treeSelector, err = extraChargeTab.treeView.GetSelection()
	if err != nil {
		panic(err)
	}
	extraChargeTab.blockAdvsTreeFilter, err = lists.BlockAdvertisementsList().FilterNew(nil)
	if err != nil {
		panic(err)
	}
	extraChargeTab.blockAdvsTreeFilter.SetVisibleFunc(extraChargeTab.filterAdvertisements)
	extraChargeTab.blockAdvsTreeView.SetModel(extraChargeTab.blockAdvsTreeFilter)

	extraChargeTab.lineAdvsTreeFilter, err = lists.LineAdvertisementsList().FilterNew(nil)
	if err != nil {
		panic(err)
	}
	extraChargeTab.extraChargeMultiplierEntry.SetPlaceholderText("0")
	extraChargeTab.lineAdvsTreeFilter.SetVisibleFunc(extraChargeTab.filterAdvertisements)
	extraChargeTab.lineAdvsTreeView.SetModel(extraChargeTab.lineAdvsTreeFilter)
	extraChargeTab.bindSignals()
	extraChargeTab.BlockSignals()
	extraChargeTab.SetSensitive(false)
	return extraChargeTab
}
func (ex *ExtraChargeTab) build(bldFile *builder.Builder) {
	ex.treeView = bldFile.FetchTreeView("ExtraChargesTreeView")
	ex.extraChargeNameLabel = bldFile.FetchLabel("ExtraChargeNameLabel")
	ex.extraChargeNameEntry = bldFile.FetchEntry("ExtraChargeNameEntry")
	ex.rightSide = bldFile.FetchBox("ExtraChargeSidebar")
	ex.extraChargeMultiplierLabel = bldFile.FetchLabel("ExtraChargeMultiplierLabel")
	ex.extraChargeMultiplierEntry = bldFile.FetchEntry("ExtraChargeMultiplierEntry")
	ex.removeExtraChargeButton = bldFile.FetchButton("ExtraChargeDeleteButton")
	ex.resetExtraChargeButton = bldFile.FetchButton("ExtraChargeResetButton")
	ex.applyExtraChargeButton = bldFile.FetchButton("ExtraChargeApplyButton")
	ex.blockAdvsLabel = bldFile.FetchLabel("ExtraChargeBlockAdvLabel")
	ex.blockAdvsTreeView = bldFile.FetchTreeView("ExtraChargeBlockAdvTreeView")
	ex.lineAdvsLabel = bldFile.FetchLabel("ExtraChargeLineAdvLabel")
	ex.lineAdvsTreeView = bldFile.FetchTreeView("ExtraChargeLineAdvTreeView")
}
