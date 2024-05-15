package lineadvcontainer

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/application"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"

	"github.com/gotk3/gotk3/gtk"
)

type DateRangeFetcher interface {
	FromDate() string
	ToDate() string
}

type LineFormMaker interface {
	NewLineForm() application.LineForm
}

type LineList interface {
	LineAdvertisementsList() *gtk.ListStore
}

type LineForm interface {
	Widget() *gtk.Widget
	FillData(selectedLine *presenter.LineAdvertisementDTO)
	FetchData() presenter.LineAdvertisementDTO
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

type LineRequests interface {
	ReleasesInTimeRange(string, string, string) bool
	RemoveLineAdvertisement(*presenter.LineAdvertisementDTO)
	UpdateLineAdvertisement(*presenter.LineAdvertisementDTO) error
}

type ErrorWindowCreator interface {
	NewErrorWindow(err error)
}

type LineAdvertisementsTab struct {
	tools  Tools
	dates  DateRangeFetcher
	req    LineRequests
	errWin ErrorWindowCreator

	lineadvertisementsListStoreFilterEnable bool
	lineadvertisementsListStoreFilter       *gtk.TreeModelFilter
	lineadvertisementsListStoreSort         *gtk.TreeModelSort
	lineAdvertisementsTreeView              *gtk.TreeView
	lineadvertisementsSelector              *gtk.TreeSelection
	lineAdvViewPort                         *gtk.Viewport
	lineAdvForm                             LineForm

	deleteButton *gtk.Button
	resetButton  *gtk.Button
	applyButton  *gtk.Button

	signalHandler LineAdvertisementsSignalHandler
}

func Create(bldFile *builder.Builder, lists LineList, lineForm LineFormMaker, tools Tools, requests LineRequests,
	dates DateRangeFetcher, errWin ErrorWindowCreator) *LineAdvertisementsTab {
	var err error
	lineAdv := new(LineAdvertisementsTab)
	lineAdv.tools = tools
	lineAdv.req = requests
	lineAdv.dates = dates
	lineAdv.errWin = errWin
	lineAdv.build(bldFile)
	lineAdv.lineadvertisementsListStoreFilter, err = lists.LineAdvertisementsList().FilterNew(nil)
	if err != nil {
		panic(err)
	}
	lineAdv.lineadvertisementsListStoreFilter.SetVisibleFunc(lineAdv.filterList)

	lineAdv.lineadvertisementsListStoreSort, err = gtk.TreeModelSortNew(lineAdv.lineadvertisementsListStoreFilter)
	if err != nil {
		panic(err)
	}
	lineAdv.lineadvertisementsListStoreSort.SetSortFunc(1, func(model *gtk.TreeModel, a, b *gtk.TreeIter) int {
		return lineAdv.tools.CompareInts(model, a, b, 1)
	})
	lineAdv.lineadvertisementsListStoreSort.SetSortFunc(2, func(model *gtk.TreeModel, a, b *gtk.TreeIter) int {
		return lineAdv.tools.CompareInts(model, a, b, 2)
	})
	lineAdv.lineadvertisementsListStoreSort.SetSortFunc(3, func(model *gtk.TreeModel, a, b *gtk.TreeIter) int {
		return lineAdv.tools.CompareInts(model, a, b, 3)
	})
	lineAdv.lineadvertisementsListStoreSort.SetSortFunc(4, func(model *gtk.TreeModel, a, b *gtk.TreeIter) int {
		return lineAdv.tools.CompareTime(model, a, b, 4)
	})
	lineAdv.lineadvertisementsListStoreSort.SetSortFunc(8, func(model *gtk.TreeModel, a, b *gtk.TreeIter) int {
		return lineAdv.tools.CompareCosts(model, a, b, 8)
	})

	lineAdv.bindSignals()
	lineAdv.lineAdvForm = lineForm.NewLineForm()
	lineAdv.lineAdvViewPort.Add(lineAdv.lineAdvForm.Widget())
	lineAdv.BlockSignals()
	lineAdv.SetSensitive(false)
	return lineAdv
}

func (line *LineAdvertisementsTab) build(bldFile *builder.Builder) {
	line.lineAdvViewPort = bldFile.FetchViewPort("LineViewPort")
	line.lineAdvertisementsTreeView = bldFile.FetchTreeView("LineAdvertisementsTreeView")
	line.lineadvertisementsSelector, _ = line.lineAdvertisementsTreeView.GetSelection()
	line.deleteButton = bldFile.FetchButton("DeleteLineAdvertisementButton")
	line.resetButton = bldFile.FetchButton("LineResetButton")
	line.applyButton = bldFile.FetchButton("LineApplyButton")
}
