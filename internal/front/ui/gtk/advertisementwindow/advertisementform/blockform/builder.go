package blockform

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/application"
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
	FilesList() *gtk.ListStore
}

type Converter interface {
	SelectedReleaseDatesToString([]string) string
	SelectedTagsToString([]presenter.SelectedTagDTO) string
	SelectedExtraChargeToString([]presenter.SelectedExtraChargeDTO) string
	YearMonthDayToString(uint, uint, uint) string
}

type Application interface {
	ActiveWin() gtk.IWindow
	RegisterReciever(*gtk.EntryBuffer)
	NewThinImageChooserWindow(application.FileChooser) application.FileChooseWindow
	Mode() int8
}

type ChooseSaveFileDialoger interface {
	NewChooseDialog(winLabel string, parent gtk.IWindow) (application.FileDialoger, error)
	NewSaveDialog(winLabel string, parent gtk.IWindow) (application.FileDialoger, error)
}
type Requests interface {
	CalculateBlockAdvertisementCost(*presenter.BlockAdvertisementDTO) error
	GetFileURI(string) (string, error)
}

type BlockAdvPage struct {
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
	signalHandler        BlockAdvertisementsSignalHandler

	sizeLabel   *gtk.Label
	sizeEntry   *gtk.Entry
	fileNameBox *gtk.Box
	file        FilePath
}

func CreateBlockAdvPage(orderList, tagList, chargeList *gtk.ListStore, tools Tools,
	listFiller ListFiller, reqGate Requests, converter Converter, app Application, objMaker ChooseSaveFileDialoger) *BlockAdvPage {

	blockAdv := new(BlockAdvPage)
	blockAdv.tools = tools
	blockAdv.app = app
	blockAdv.listFil = listFiller
	blockAdv.req = reqGate
	blockAdv.conv = converter
	blockAdv.orderListStore = orderList
	blockAdv.tagsListStore = tagList
	blockAdv.extraChargeListStore = chargeList

	advBuilder, err := builder.NewBuilderFromString(builder.AdvertisementForm)
	if err != nil {
		panic(err)
	}
	blockAdv.BuildBlockPage(advBuilder)
	blockAdv.costEntry.SetPlaceholderText("0,00")
	blockAdv.sizeEntry.SetPlaceholderText("0")
	if blockAdv.app.Mode() == application.ThickMode {
		blockAdv.attachThickPath(advBuilder, blockAdv.app, objMaker)
	} else {
		blockAdv.attachThinPath(advBuilder, blockAdv.app, objMaker, reqGate)
	}

	blockAdv.listFil.UnMarkSelectedOrder(blockAdv.orderListStore)
	blockAdv.listFil.UnMarkExtraCost(blockAdv.tagsListStore)
	blockAdv.listFil.UnMarkExtraCost(blockAdv.extraChargeListStore)
	blockAdv.releaseDatesListStore.SetSortFunc(0, func(model *gtk.TreeModel, a, b *gtk.TreeIter) int {
		return tools.CompareTime(model, a, b, 0)
	})
	blockAdv.bindSignals()
	return blockAdv
}

func (blkpg *BlockAdvPage) BuildBlockPage(buildFile *builder.Builder) {
	blkpg.box = buildFile.FetchBox("AddBlockAdvertisementBox")
	blkpg.idBox = buildFile.FetchBox("BlockIdBox")
	blkpg.idLabel = buildFile.FetchLabel("BlockIDLabel")
	blkpg.idEntry = buildFile.FetchEntry("BlockIDEntry")
	blkpg.orderBox = buildFile.FetchBox("BlockOrderSelectorBox")
	blkpg.orderLabel = buildFile.FetchLabel("BlockOrderLabel")
	blkpg.orderTreeView = buildFile.FetchTreeView("BlockOrderTreeView")
	blkpg.orderTreeView.SetModel(blkpg.orderListStore)
	blkpg.orderCellrenderToggle = buildFile.FetchCellRendererToggle("BlockSelectedOrderToggle")
	blkpg.sizeLabel = buildFile.FetchLabel("BlockSezeLabel")
	blkpg.sizeEntry = buildFile.FetchEntry("BlockSizeEntry")
	blkpg.releaseDatesLabel = buildFile.FetchLabel("BlockReleaseDatesLabel")
	blkpg.releaseDatesMenuButton = buildFile.FetchMenuButton("BlockReleaseDateMenuButton")
	blkpg.releaseDatesTreeview = buildFile.FetchTreeView("BlockReleaseDatesTreeView")
	blkpg.releaseDatesSelection, _ = blkpg.releaseDatesTreeview.GetSelection()
	blkpg.releaseDatesListStore = buildFile.FetchListStore("BlockReleaseDatesListStore")
	blkpg.releaseDatesPopover = buildFile.FetchPopover("BlockSetReleaseDatePopover")
	blkpg.releaseDatesCalendar = buildFile.FetchCalendar("BlockSetReleaseDateCalendar")
	blkpg.releaseDatesDeleteButton = buildFile.FetchButton("BlockDeleteReleaseDateButton")
	blkpg.releaseDatesAppendButton = buildFile.FetchButton("BlockAddReleaseDateButton")
	blkpg.releaseCountLabel = buildFile.FetchLabel("BlockReleaseCountsLabel")
	blkpg.releaseCountEntry = buildFile.FetchEntry("BlockReleaseCountsEntry")
	blkpg.setReleaseCount("0")
	blkpg.tagsLabel = buildFile.FetchLabel("BlockTagsLabel")
	blkpg.tagsTreeview = buildFile.FetchTreeView("BlockTagSelectorTreeView")
	blkpg.tagsTreeview.SetModel(blkpg.tagsListStore)
	blkpg.tagCellrenderToggle = buildFile.FetchCellRendererToggle("BlockSelectedTagSwitcher")
	blkpg.extraChargeLabel = buildFile.FetchLabel("BlockCostRateMultiplayerLabel")
	blkpg.extraChargeTreeview = buildFile.FetchTreeView("BlockExtraChargeSelectorTreeView")
	blkpg.extraChargeTreeview.SetModel(blkpg.extraChargeListStore)
	blkpg.extraChargeCellrenderToggle = buildFile.FetchCellRendererToggle("BlockSelectedExtraChargeSwitcher")
	blkpg.costLabel = buildFile.FetchLabel("BlockTotalPaymentLabel")
	blkpg.costEntry = buildFile.FetchEntry("BlockTotalPaymentEntry")
	blkpg.fileNameBox = buildFile.FetchBox("FileNameBox")
	blkpg.costCalculateButton = buildFile.FetchButton("CalculateCostBlockAdvertisementButton")
	blkpg.textLabel = buildFile.FetchLabel("BlockCommentLabel")
	blkpg.textTextView = buildFile.FetchTextView("BlockCommentTextView")
	blkpg.textTextBuffer = buildFile.FetchTextBuffer("BlockTextBuffer")
}
