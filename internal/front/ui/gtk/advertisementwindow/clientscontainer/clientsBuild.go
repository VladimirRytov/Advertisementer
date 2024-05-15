package clientscontainer

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"

	"github.com/gotk3/gotk3/gtk"
)

type Tools interface {
	CompareStrings(*gtk.TreeModel, *gtk.TreeIter, *gtk.TreeIter, int) int
	StringFromIter(*gtk.TreeIter, *gtk.TreeModel, int) (string, error)
}

type Requests interface {
	ValueInString(string, string) bool
	UpdateClient(*presenter.ClientDTO) error
	RemoveClient(string)
}

type ListStores interface {
	ClientsList() *gtk.ListStore
	OrdersList() *gtk.ListStore
}

type ClientsTab struct {
	tools Tools
	req   Requests

	clients       *gtk.TreeModelSort
	treeView      *gtk.TreeView
	treeSelection *gtk.TreeSelection
	rightSide     *gtk.Box

	nameLabel *gtk.Label
	nameEntry *gtk.Entry

	contactNumberLabel      *gtk.Label
	contactNumberEntry      *gtk.Entry
	contactNumberErrorLabel *gtk.Label

	emailLabel      *gtk.Label
	emailEntry      *gtk.Entry
	emailErrorLabel *gtk.Label

	additionalInfoLabel  *gtk.Label
	additionalInfoEntry  *gtk.TextView
	additionalInfoBuffer *gtk.TextBuffer

	deleteButton *gtk.Button
	resetButton  *gtk.Button
	applyButton  *gtk.Button

	ordersLabel        *gtk.Label
	ordersFilterEnable bool
	ordersFilter       *gtk.TreeModelFilter
	ordersView         *gtk.TreeView

	signalHandler ClientTabSignalHandler
}

func Create(bldFile *builder.Builder, reqGate Requests, tools Tools, lists ListStores) *ClientsTab {
	var err error

	client := new(ClientsTab)
	client.tools = tools
	client.req = reqGate
	client.build(bldFile)
	client.clients, err = gtk.TreeModelSortNew(lists.ClientsList())
	if err != nil {
		panic(err)
	}
	client.clients.SetSortFunc(1, func(model *gtk.TreeModel, a, b *gtk.TreeIter) int {
		return client.tools.CompareStrings(model, a, b, 1)
	})

	client.treeSelection, err = client.treeView.GetSelection()
	if err != nil {
		panic(err)
	}
	client.ordersFilter, err = lists.OrdersList().FilterNew(nil)
	if err != nil {
		panic(err)
	}
	client.ordersFilter.SetVisibleFunc(client.filterOrders)
	client.ordersView.SetModel(client.ordersFilter)
	client.bindSignals()
	client.BlockSignals()
	client.SetSensitive(false)
	return client
}

func (cl *ClientsTab) build(bldFile *builder.Builder) {
	cl.treeView = bldFile.FetchTreeView("ClientsTreeView")
	cl.nameLabel = bldFile.FetchLabel("ClientNameLabel")
	cl.nameEntry = bldFile.FetchEntry("ClientNameEntry")
	cl.rightSide = bldFile.FetchBox("ClientSidebar")
	cl.contactNumberLabel = bldFile.FetchLabel("ClientContactNumbersLabel")
	cl.contactNumberEntry = bldFile.FetchEntry("ClientContactNumbersEntry")
	cl.contactNumberErrorLabel = bldFile.FetchLabel("ClientPhoneNumberError")
	cl.emailLabel = bldFile.FetchLabel("ClientEmailLabel")
	cl.emailEntry = bldFile.FetchEntry("ClientEmailEntry")
	cl.emailErrorLabel = bldFile.FetchLabel("ClientEmailError")
	cl.additionalInfoLabel = bldFile.FetchLabel("ClientAdditionalInformationLabel")
	cl.additionalInfoEntry = bldFile.FetchTextView("ClientAdditionalInformationTextView")
	cl.additionalInfoBuffer = bldFile.FetchTextBuffer("ClientAdditionalInformationBuffer")
	cl.deleteButton = bldFile.FetchButton("DeleteClientButton")
	cl.resetButton = bldFile.FetchButton("ClientsResetButton")
	cl.applyButton = bldFile.FetchButton("ClientsApplyButton")
	cl.ordersLabel = bldFile.FetchLabel("ClientOrdersLabel")
	cl.ordersView = bldFile.FetchTreeView("ClientOrdersTreeView")
}
