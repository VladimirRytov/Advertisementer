package lineform

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"

	"github.com/gotk3/gotk3/gtk"
)

type Tools interface {
	CompareTime(*gtk.TreeModel, *gtk.TreeIter, *gtk.TreeIter, int) int
	InterfaceFromIter(*gtk.TreeIter, *gtk.TreeModel, int) (interface{}, error)
	StringFromIter(*gtk.TreeIter, *gtk.TreeModel, int) (string, error)
	FindValue(string, *gtk.ListStore, int) (*gtk.TreeIter, error)
	FindIntValue(int, *gtk.ListStore, int) (*gtk.TreeIter, error)
	CompareNewTime(*gtk.TreeModel, *gtk.TreeIter, string, int) int
	CheckCostAdvString(*gtk.Entry, string)
}

type ListFiller interface {
	MarkSelectedOrder(*gtk.ListStore, int) (*gtk.TreePath, error)
	FillReleaseDates(*gtk.ListStore, string)
	UnMarkSelectedOrder(*gtk.ListStore)
	MarkExtraCost(*gtk.ListStore, string)
	UnMarkExtraCost(*gtk.ListStore)
}

type ListStores interface {
	OrdersList() *gtk.ListStore
	ExtraChargesList() *gtk.ListStore
	TagsList() *gtk.ListStore
}

type Converter interface {
	SelectedReleaseDatesToString([]string) string
	SelectedTagsToString([]presenter.SelectedTagDTO) string
	SelectedExtraChargeToString([]presenter.SelectedExtraChargeDTO) string
	YearMonthDayToString(uint, uint, uint) string
}

type Application interface {
	RegisterReciever(*gtk.EntryBuffer)
}

type Requests interface {
	CalculateLineAdvertisementCost(*presenter.LineAdvertisementDTO) error
}

type LineAdvPage struct {
	req     Requests
	app     Application
	tools   Tools
	listFil ListFiller
	conv    Converter

	box *gtk.Box

	idBox   *gtk.Box
	idLabel *gtk.Label
	idEntry *gtk.Entry

	orderBox              *gtk.Box
	orderLabel            *gtk.Label
	orderTreeView         *gtk.TreeView
	orderCellrenderToggle *gtk.CellRendererToggle

	releaseDatesLabel        *gtk.Label
	releaseDatesMenuButton   *gtk.MenuButton
	releaseDatesSelection    *gtk.TreeSelection
	releaseDatesTreeview     *gtk.TreeView
	releaseDatesListStore    *gtk.ListStore
	releaseDatesPopover      *gtk.Popover
	releaseDatesCalendar     *gtk.Calendar
	releaseDatesDeleteButton *gtk.Button
	releaseDatesAppendButton *gtk.Button

	releaseCountLabel *gtk.Label
	releaseCountEntry *gtk.Entry

	tagsLabel           *gtk.Label
	tagsTreeview        *gtk.TreeView
	tagCellrenderToggle *gtk.CellRendererToggle

	extraChargeLabel            *gtk.Label
	extraChargeTreeview         *gtk.TreeView
	extraChargeCellrenderToggle *gtk.CellRendererToggle

	costLabel           *gtk.Label
	costEntry           *gtk.Entry
	costCalculateButton *gtk.Button

	textLabel      *gtk.Label
	textTextView   *gtk.TextView
	textTextBuffer *gtk.TextBuffer

	orderListStore       *gtk.ListStore
	tagsListStore        *gtk.ListStore
	extraChargeListStore *gtk.ListStore
	signalHandler        advSignalHandler
}

func (adwp *LineAdvPage) BuildAdvForm() {

}
func (adwp *LineAdvPage) SetOrderModel(model *gtk.ListStore) {
	adwp.orderTreeView.SetModel(model)
}

func (adwp *LineAdvPage) SetTagModel(model *gtk.ListStore) {
	adwp.tagsTreeview.SetModel(model)
}

func (adwp *LineAdvPage) SetExtraChargeModel(model *gtk.ListStore) {
	adwp.extraChargeTreeview.SetModel(model)
}

func (adwp *LineAdvPage) Widget() *gtk.Widget {
	return adwp.box.ToWidget()
}

func (adwp *LineAdvPage) SetVisibleID(hide bool) {
	adwp.idEntry.SetVisible(hide)
	adwp.idLabel.SetVisible(hide)
}

func CreateLineAdvPage(orderList, tagList, chargeList *gtk.ListStore, tools Tools, listFiller ListFiller,
	reqGate Requests, converter Converter, app Application) *LineAdvPage {

	lineAdv := new(LineAdvPage)
	lineAdv.tools = tools
	lineAdv.listFil = listFiller
	lineAdv.req = reqGate
	lineAdv.conv = converter
	lineAdv.app = app
	lineAdv.orderListStore = orderList
	lineAdv.tagsListStore = tagList
	lineAdv.extraChargeListStore = chargeList
	advBuilder, err := builder.NewBuilderFromString(builder.AdvertisementForm)
	if err != nil {
		panic(err)
	}
	lineAdv.BuildLinePage(advBuilder)
	lineAdv.bindSignals()
	lineAdv.costEntry.SetPlaceholderText("0,00")
	lineAdv.listFil.UnMarkSelectedOrder(lineAdv.orderListStore)
	lineAdv.listFil.UnMarkExtraCost(lineAdv.tagsListStore)
	lineAdv.listFil.UnMarkExtraCost(lineAdv.extraChargeListStore)
	lineAdv.releaseDatesListStore.SetSortFunc(0, func(model *gtk.TreeModel, a, b *gtk.TreeIter) int {
		return lineAdv.tools.CompareTime(model, a, b, 0)
	})
	return lineAdv
}

func (linepg *LineAdvPage) BuildLinePage(buildFile *builder.Builder) {
	linepg.box = buildFile.FetchBox("AddLineAdvertisementBox")
	linepg.idBox = buildFile.FetchBox("LineIdBox")
	linepg.idLabel = buildFile.FetchLabel("LineIDLabel")
	linepg.idEntry = buildFile.FetchEntry("LineIDEntry")
	linepg.orderLabel = buildFile.FetchLabel("LineOrderLabel")
	linepg.orderBox = buildFile.FetchBox("LineOrderSelectorBox")
	linepg.orderTreeView = buildFile.FetchTreeView("LineOrderTreeView")
	linepg.orderTreeView.SetModel(linepg.orderListStore)
	linepg.orderCellrenderToggle = buildFile.FetchCellRendererToggle("LineSelectedOrderToggle")
	linepg.releaseDatesLabel = buildFile.FetchLabel("LineReleaseDatesLabel")
	linepg.releaseDatesMenuButton = buildFile.FetchMenuButton("LineReleaseDatesMenuButton")
	linepg.releaseDatesTreeview = buildFile.FetchTreeView("LineReleaseDatesTreeView")
	linepg.releaseDatesSelection, _ = linepg.releaseDatesTreeview.GetSelection()
	linepg.releaseDatesListStore = buildFile.FetchListStore("LineReleaseDatesListStore")
	linepg.releaseDatesPopover = buildFile.FetchPopover("LineSetReleaseDatePopover")
	linepg.releaseDatesCalendar = buildFile.FetchCalendar("LineSetReleaseDateCalendar")
	linepg.releaseDatesDeleteButton = buildFile.FetchButton("LineDeleteReleaseDateButton")
	linepg.releaseDatesAppendButton = buildFile.FetchButton("LineAddReleaseDateButton")
	linepg.releaseCountLabel = buildFile.FetchLabel("LineReleaseCountsLabel")
	linepg.releaseCountEntry = buildFile.FetchEntry("LineReleaseCountsEntry")
	linepg.setReleaseCount(0)
	linepg.tagsLabel = buildFile.FetchLabel("LineTagsLabel")
	linepg.tagsTreeview = buildFile.FetchTreeView("LineTagSelectorTreeView")
	linepg.tagsTreeview.SetModel(linepg.tagsListStore)
	linepg.tagCellrenderToggle = buildFile.FetchCellRendererToggle("LineSelectedTagSwitcher")
	linepg.extraChargeLabel = buildFile.FetchLabel("LineCostRateMultiplayerLabel")
	linepg.extraChargeTreeview = buildFile.FetchTreeView("LineExtraChargeSelectorTreeView")
	linepg.extraChargeTreeview.SetModel(linepg.extraChargeListStore)
	linepg.extraChargeCellrenderToggle = buildFile.FetchCellRendererToggle("LineSelectedExtraChargeSwitcher")
	linepg.costLabel = buildFile.FetchLabel("LineTotalPaymentLabel")
	linepg.costEntry = buildFile.FetchEntry("LineTotalPaymentEntry")
	linepg.costCalculateButton = buildFile.FetchButton("CalculateCostLineAdvertisementButton")
	linepg.textLabel = buildFile.FetchLabel("LineCommentLabel")
	linepg.textTextView = buildFile.FetchTextView("LineCommentTextView")
	linepg.textTextBuffer = buildFile.FetchTextBuffer("LineTextBuffer")
}
