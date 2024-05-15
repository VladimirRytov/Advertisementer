package liststores

import (
	"github.com/gotk3/gotk3/gtk"
)

type Tools interface {
	StringFromIter(*gtk.TreeIter, *gtk.TreeModel, int) (string, error)
	FindValue(string, *gtk.ListStore, int) (*gtk.TreeIter, error)
	FindIntValue(int, *gtk.ListStore, int) (*gtk.TreeIter, error)

	RemoveValue(string, int, *gtk.ListStore)
	RemoveIntValue(int, int, *gtk.ListStore)

	InterfaceFromIter(*gtk.TreeIter, *gtk.TreeModel, int) (interface{}, error)
}

type ListStores struct {
	tools       Tools
	replaceMode bool

	tags                *gtk.ListStore
	extraCharge         *gtk.ListStore
	clients             *gtk.ListStore
	orders              *gtk.ListStore
	blockAdvertisements *gtk.ListStore
	lineAdvertisements  *gtk.ListStore
	costRates           *gtk.ListStore
	messageList         *gtk.ListStore
	files               *gtk.ListStore
	localDatabases      *gtk.ListStore
}

func Create(tools Tools) *ListStores {
	ls := new(ListStores)
	ls.tools = tools
	ls.baseLists()
	return ls
}

func (ls *ListStores) baseLists() {
	var err error
	ls.tags, err = ls.newTagsList()
	if err != nil {
		panic(err)
	}
	ls.extraCharge, err = ls.newExtraChargeList()
	if err != nil {
		panic(err)
	}
	ls.clients, err = ls.newClientList()
	if err != nil {
		panic(err)
	}
	ls.orders, err = ls.newOrderList()
	if err != nil {
		panic(err)
	}
	ls.blockAdvertisements, err = ls.newBlockAdvertisementsList()
	if err != nil {
		panic(err)
	}
	ls.lineAdvertisements, err = ls.newLineAdvertisementsList()
	if err != nil {
		panic(err)
	}
	ls.costRates, err = ls.costRatesNewListStore()
	if err != nil {
		panic(err)
	}

	ls.messageList, err = ls.newMessageList()
	if err != nil {
		panic(err)
	}
	ls.localDatabases, err = ls.newLocalDatabaseList()
	if err != nil {
		panic(err)
	}
}
