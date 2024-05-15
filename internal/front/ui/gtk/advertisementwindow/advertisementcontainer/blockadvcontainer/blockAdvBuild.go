package blockadvcontainer

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/application"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"

	"github.com/gotk3/gotk3/gtk"
)

type FilePath interface {
	FilePath() string
	SetFilePath(string)
	ViewClicked()
	Box() *gtk.Box
}

type DateRangeFetcher interface {
	FromDate() string
	ToDate() string
}

type BlockFormMaker interface {
	NewBlockForm() application.BlockForm
}

type BlockList interface {
	BlockAdvertisementsList() *gtk.ListStore
}

type BlockForm interface {
	Widget() *gtk.Widget
	FillData(selectedBlock *presenter.BlockAdvertisementDTO)
	FetchData() presenter.BlockAdvertisementDTO
	SetSensetive(bool)
	SetNewAdvMode(bool)
	SetNewNestedAdvMode(bool)
	Reset()
	UnsetModel()
	SetModel()
}

type Tools interface {
	CompareInts(model *gtk.TreeModel, a, b *gtk.TreeIter, col int) int
	CompareTime(model *gtk.TreeModel, a, b *gtk.TreeIter, col int) int
	CompareCosts(model *gtk.TreeModel, a, b *gtk.TreeIter, col int) int
	StringFromIter(*gtk.TreeIter, *gtk.TreeModel, int) (string, error)
	InterfaceFromIter(*gtk.TreeIter, *gtk.TreeModel, int) (interface{}, error)
}

type BlockRequests interface {
	ReleasesInTimeRange(string, string, string) bool
	RemoveBlockAdvertisement(*presenter.BlockAdvertisementDTO)
	UpdateBlockAdvertisement(*presenter.BlockAdvertisementDTO) error
}

type ErrorWindowCreator interface {
	NewErrorWindow(err error)
}

type BlockAdvertisementsTab struct {
	tools      Tools
	req        BlockRequests
	dateRanges DateRangeFetcher
	errWin     ErrorWindowCreator

	blockadvertisementsListStoreFilterEnabled bool
	blockadvertisementsListStoreFilter        *gtk.TreeModelFilter
	blockadvertisementsListStoreSort          *gtk.TreeModelSort
	blockAdvertisementsTreeView               *gtk.TreeView
	blockadvertisementsSelector               *gtk.TreeSelection
	blockAdvViewPort                          *gtk.Viewport
	blockAdvForm                              BlockForm

	deleteButton *gtk.Button
	resetButton  *gtk.Button
	applyButton  *gtk.Button

	signalHandler BlockAdvertisementsSignalHandler
}

func Create(bldFile *builder.Builder, lists BlockList, blockForm BlockFormMaker, tools Tools, requests BlockRequests,
	dates DateRangeFetcher, errWin ErrorWindowCreator) *BlockAdvertisementsTab {
	var err error
	blockAdv := new(BlockAdvertisementsTab)
	blockAdv.tools = tools
	blockAdv.req = requests
	blockAdv.errWin = errWin
	blockAdv.dateRanges = dates
	blockAdv.build(bldFile)
	blockAdv.blockadvertisementsListStoreFilter, err = lists.BlockAdvertisementsList().ToTreeModel().FilterNew(nil)
	if err != nil {
		panic(err)
	}
	blockAdv.blockadvertisementsListStoreFilter.SetVisibleFunc(blockAdv.filterList)
	blockAdv.blockadvertisementsSelector, err = blockAdv.blockAdvertisementsTreeView.GetSelection()
	if err != nil {
		panic(err)
	}

	blockAdv.blockadvertisementsListStoreSort, err = gtk.TreeModelSortNew(blockAdv.blockadvertisementsListStoreFilter)
	if err != nil {
		panic(err)
	}
	blockAdv.blockadvertisementsListStoreSort.SetSortFunc(1, func(model *gtk.TreeModel, a, b *gtk.TreeIter) int {
		return blockAdv.tools.CompareInts(model, a, b, 1)
	})
	blockAdv.blockadvertisementsListStoreSort.SetSortFunc(2, func(model *gtk.TreeModel, a, b *gtk.TreeIter) int {
		return blockAdv.tools.CompareInts(model, a, b, 2)
	})
	blockAdv.blockadvertisementsListStoreSort.SetSortFunc(3, func(model *gtk.TreeModel, a, b *gtk.TreeIter) int {
		return blockAdv.tools.CompareInts(model, a, b, 3)
	})
	blockAdv.blockadvertisementsListStoreSort.SetSortFunc(4, func(model *gtk.TreeModel, a, b *gtk.TreeIter) int {
		return blockAdv.tools.CompareTime(model, a, b, 4)
	})
	blockAdv.blockadvertisementsListStoreSort.SetSortFunc(8, func(model *gtk.TreeModel, a, b *gtk.TreeIter) int {
		return blockAdv.tools.CompareInts(model, a, b, 8)
	})
	blockAdv.blockadvertisementsListStoreSort.SetSortFunc(9, func(model *gtk.TreeModel, a, b *gtk.TreeIter) int {
		return blockAdv.tools.CompareCosts(model, a, b, 9)
	})

	blockAdv.blockAdvForm = blockForm.NewBlockForm()
	blockAdv.blockAdvViewPort.Add(blockAdv.blockAdvForm.Widget())
	blockAdv.bindSignals()
	blockAdv.BlockSignals()
	blockAdv.SetSensitive(false)
	return blockAdv
}

func (blk *BlockAdvertisementsTab) build(bldFile *builder.Builder) {
	blk.blockAdvertisementsTreeView = bldFile.FetchTreeView("BlockAdvertisementsTreeView")
	blk.deleteButton = bldFile.FetchButton("DeleteBlockAdvertisementButton")
	blk.blockAdvViewPort = bldFile.FetchViewPort("BlockViewPort")
	blk.resetButton = bldFile.FetchButton("BlockResetButton")
	blk.applyButton = bldFile.FetchButton("BlockApplyButton")
}
