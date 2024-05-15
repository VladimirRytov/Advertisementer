package liststores

import (
	"errors"

	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func (ls *ListStores) newLocalDatabaseList() (*gtk.ListStore, error) {
	localDatabases, err := gtk.ListStoreNew(glib.TYPE_STRING)
	if err != nil {
		return nil, errors.Join(errors.New("listStores.newLocalDatabaseList: cannot create listStore"), err)
	}
	return localDatabases, err
}

func (ls *ListStores) AppendLocalDatabase(db string) {
	logging.Logger.Debug("Add localDb to listStore", "name", db)
	if ls.replaceMode {
		pos, err := ls.tools.FindValue(db, ls.localDatabases, 0)
		if err == nil {
			ls.localDatabases.SetValue(pos, 0, db)
			return
		}
	}
	ls.localDatabases.SetValue(ls.localDatabases.Append(), 0, db)
}

func (ls *ListStores) LocalDatabaseList() *gtk.ListStore {
	return ls.localDatabases
}

func (ls *ListStores) RemoveLocalDatabase(db string) {
	ls.tools.RemoveValue(db, 0, ls.localDatabases)
}

func (ls *ListStores) ClearLocalDatabaseList() {
	ls.localDatabases.Clear()
}
