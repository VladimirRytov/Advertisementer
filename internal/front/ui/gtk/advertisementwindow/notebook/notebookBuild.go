package notebook

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/application"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"

	"github.com/gotk3/gotk3/gtk"
)

type windowMaker interface {
	CreateAddClientWindow() application.AddClientWindow
	CreateAddOrderWindow() application.AddOrderWindow
	CreateAddAdvertisementWindow() application.AddAdvertisementWindow
	CreateAddExtraChargeWindow() application.AddExtraChargeWindow
	CreateAddTagWindow() application.AddTagWindow
}

type App interface {
	windowMaker
	ListStores() application.ListStores
	AdvertisementWindow() application.AdvertisementsWindow
}

type Requests interface {
	ActiveCostRate()
	LockReciever(lock bool)
}

type ObjectMaker interface {
	NewClientTab(b *builder.Builder) application.ClientTab
	NewOrderTab(b *builder.Builder, advWin application.AdvertisementsWindow) application.OrderTab
	NewBlockTab(b *builder.Builder, advWin application.AdvertisementsWindow) application.BlockTab
	NewLineTab(b *builder.Builder, advWin application.AdvertisementsWindow) application.LineTab
	NewTagTab(b *builder.Builder, advWin application.AdvertisementsWindow) application.TagTab
	NewExtraChargeTab(b *builder.Builder, advWin application.AdvertisementsWindow) application.ChargeTab

	NewLineForm() application.LineForm
	NewBlockForm() application.BlockForm
}

type Notebook struct {
	req         Requests
	tools       application.Tools
	application App
	lists       application.ListStores
	objMaker    ObjectMaker
	advWin      application.AdvertisementsWindow

	notebook             *gtk.Notebook
	addEntryPopover      *gtk.PopoverMenu
	updateButton         *gtk.Button
	createAdvertisements *gtk.Button
	createClient         *gtk.Button
	createOrder          *gtk.Button
	createTag            *gtk.Button
	createExtraCharge    *gtk.Button
	SignalHandler        notebookSignalHandler

	BlockAdvertisementsTab application.BlockTab
	LineAdvertisementsTab  application.LineTab
	OrdersTab              application.Tabs
	ClientsTab             application.ClientTab
	TagsTab                application.TagTab
	ExtrachargesTab        application.ChargeTab
}

func Create(bldFile *builder.Builder, objectMaker ObjectMaker, application App, tools application.Tools, lists application.ListStores, advWin application.AdvertisementsWindow, reqGate Requests) *Notebook {
	notebook := new(Notebook)
	notebook.tools = tools
	notebook.application = application
	notebook.objMaker = objectMaker
	notebook.lists = lists
	notebook.advWin = advWin
	notebook.req = reqGate
	notebook.build(bldFile)

	notebook.BlockAdvertisementsTab = objectMaker.NewBlockTab(bldFile, advWin)
	notebook.LineAdvertisementsTab = objectMaker.NewLineTab(bldFile, advWin)
	notebook.OrdersTab = objectMaker.NewOrderTab(bldFile, notebook.advWin)
	notebook.ClientsTab = objectMaker.NewClientTab(bldFile)
	notebook.TagsTab = objectMaker.NewTagTab(bldFile, advWin)
	notebook.ExtrachargesTab = objectMaker.NewExtraChargeTab(bldFile, advWin)

	notebook.bindSignals()
	notebook.EnableAllModels()
	return notebook
}
func (n *Notebook) build(bldFile *builder.Builder) {
	n.addEntryPopover = bldFile.FetchPopoverMenu("AddEntryPopover")
	n.notebook = bldFile.FetchNoteBook("NotebookWithTables")
	n.updateButton = bldFile.FetchButton("UpdateEntriesButton")
	n.createAdvertisements = bldFile.FetchButton("CreateAdvertisementsButton")
	n.createOrder = bldFile.FetchButton("CreateOrderButton")
	n.createClient = bldFile.FetchButton("CreateClientButton")
	n.createTag = bldFile.FetchButton("CreateTagButton")
	n.createExtraCharge = bldFile.FetchButton("CreateExtraChargeButton")

}
