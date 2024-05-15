package liststores

import (
	"errors"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func (ls *ListStores) ClientsList() *gtk.ListStore {
	return ls.clients
}

func (ls *ListStores) AppendClient(client presenter.ClientDTO) {
	logging.Logger.Debug("Add client to listStore", "client name", client.Name)
	if ls.replaceMode {
		pos, err := ls.tools.FindValue(client.Name, ls.clients, 1)
		if err == nil {
			ls.clients.Set(pos, []int{0, 1, 2, 3, 4},
				[]interface{}{false, client.Name, client.Phones, client.Email, client.AdditionalInformation})
			return
		}
	}
	ls.clients.Set(ls.clients.Append(),
		[]int{0, 1, 2, 3, 4}, []interface{}{false, client.Name, client.Phones, client.Email, client.AdditionalInformation})
}

func (ls *ListStores) newClientList() (*gtk.ListStore, error) {
	clients, err := gtk.ListStoreNew(glib.TYPE_BOOLEAN, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING)
	if err != nil {
		return nil, errors.Join(errors.New("listStores.NewClientList: cannot create listStore"), err)
	}
	return clients, nil
}

func (ls *ListStores) ClientListCopy() (*gtk.ListStore, error) {
	clients, err := gtk.ListStoreNew(glib.TYPE_BOOLEAN, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING)
	if err != nil {
		return nil, errors.Join(errors.New("listStores.ClientListCopy: cannot create listStore"), err)
	}
	ls.clients.ForEach(func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
		clientName, err := ls.tools.StringFromIter(iter, &ls.clients.TreeModel, 1)
		if err != nil {
			logging.Logger.Error("listStores.ClientListCopy: cannt get name from listStore", "error", err)
			return false
		}
		phone, err := ls.tools.StringFromIter(iter, &ls.clients.TreeModel, 2)
		if err != nil {
			logging.Logger.Error("listStores.ClientListCopy: cannt get phone Number from listStore", "error", err)
			return false
		}
		email, err := ls.tools.StringFromIter(iter, &ls.clients.TreeModel, 3)
		if err != nil {
			logging.Logger.Error("listStores.ClientListCopy: cannt get email from listStore", "error", err)
			return false
		}
		addInfo, err := ls.tools.StringFromIter(iter, &ls.clients.TreeModel, 4)
		if err != nil {
			logging.Logger.Error("listStores.ClientListCopy: cannt get additional information from listStore", "error", err)
			return false
		}

		clients.Set(clients.Append(), []int{0, 1, 2, 3, 4}, []interface{}{false, clientName, phone, email, addInfo})
		return false
	})
	return clients, nil
}

func (ls *ListStores) RemoveClient(client string) {
	ls.tools.RemoveValue(client, 1, ls.clients)
	ls.RemoveOrderByClientID(client)
}

func (ls *ListStores) ClearClientList() {
	ls.clients.Clear()
}
