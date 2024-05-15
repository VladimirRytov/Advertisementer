package todatabasewin

import (
	"net/url"
	"os"
	"path/filepath"

	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type ConnectToDatabaseWindowHandler struct {
	networkDatabaseChanged    glib.SignalHandle
	chooseFolderButtonPressed glib.SignalHandle
	refreshButtonPressed      glib.SignalHandle
	createButtonPressed       glib.SignalHandle

	createConfimButtonPressed glib.SignalHandle
	createCancelButtonPressed glib.SignalHandle

	deleteButtonPressed       glib.SignalHandle
	deleteConfimButtonPressed glib.SignalHandle
	deleteCancelButtonPressed glib.SignalHandle

	localDatabasesRowInserted glib.SignalHandle
	localDatabasesRowRemoved  glib.SignalHandle
}

func (tkh *ConnectToDatabaseWindow) BindSignals() {
	tkh.signals.networkDatabaseChanged = tkh.NetworkDataBaseComboBox.Connect("changed", tkh.NetworkDatabaseChanged)
	tkh.signals.chooseFolderButtonPressed = tkh.LocalDatabasePathButton.Connect("clicked", tkh.chooseFolderButtonPressed)
	tkh.signals.refreshButtonPressed = tkh.LocalDatabaseRefreshButton.Connect("clicked", tkh.RefreshButtonPressed)

	tkh.signals.createButtonPressed = tkh.LocalDatabaseCreateButton.Connect("clicked", tkh.createButtonPressed)
	tkh.signals.createConfimButtonPressed = tkh.LocalDatabaseApplyCreationButton.Connect("clicked", tkh.applyCreationButtonPressed)
	tkh.signals.createCancelButtonPressed = tkh.LocalDatabaseCancelCreateButton.Connect("clicked", tkh.cancelButtonPressed)

	tkh.signals.deleteButtonPressed = tkh.LocalDatabaseRemoveButton.Connect("clicked", tkh.removeButtonPressed)
	tkh.signals.deleteConfimButtonPressed = tkh.LocalDatabaseApplyRemoveButton.Connect("clicked", tkh.applyRemovingButtonPressed)
	tkh.signals.deleteCancelButtonPressed = tkh.LocalDatabaseCancelRemoveButton.Connect("clicked", tkh.cancelButtonPressed)
	tkh.signals.localDatabasesRowInserted = tkh.FoundLocalDatabasesListStore.Connect("row-inserted", tkh.selectedDatabasesChanged)
	tkh.signals.localDatabasesRowRemoved = tkh.FoundLocalDatabasesListStore.Connect("row-deleted", tkh.selectedDatabasesChanged)

}

func (tkh *ConnectToDatabaseWindow) NetworkDatabaseChanged() {
	logging.Logger.Debug("thickClient: network database combobox changed")
	tkh.req.DefaultNetworkPort(
		tkh.NetworkDataBaseComboBox.GetActiveID())
}

func (tkh *ConnectToDatabaseWindow) chooseFolderButtonPressed() {
	d, err := tkh.folderSelector.NewFolderChooseDialog("Выберите каталог с базами данных", tkh.loginWin.Window())
	if err != nil {
		tkh.loginWin.ShowError(err)
		return
	}
	d.BindResponseSignal(func(self *glib.Object, responce int) {
		if gtk.ResponseType(responce) == gtk.RESPONSE_ACCEPT {
			folderURI, err := url.Parse(d.GetURI())
			if err != nil {
				return
			}
			folder := tkh.pathParser.ParsePath(folderURI.Path)
			tkh.SetLocalfolderPath(folder)
			tkh.RefreshButtonPressed()
		}
	})
	d.Show()
}

func (tkh *ConnectToDatabaseWindow) RefreshButtonPressed() {
	dbList, err := tkh.databaseList()
	if err != nil {
		return
	}
	tkh.databaseLists.ClearLocalDatabaseList()
	tkh.fillDatabaseList(dbList)
	tkh.LocalDatabaseNameComboBox.SetActive(0)
}

func (tkh *ConnectToDatabaseWindow) createButtonPressed() {
	tkh.LocalDataBaseComboBox.SetSensitive(false)
	tkh.changeStack(createDb)
}

func (tkh *ConnectToDatabaseWindow) applyCreationButtonPressed() {
	err := tkh.req.CreateFile(filepath.Join(tkh.folderPath(), tkh.newDatabaseName()+tkh.localDatabaseSuffix()))
	if err != nil {
		tkh.loginWin.ShowError(err)
		return
	}
	tkh.RefreshButtonPressed()
	tkh.selectCreatedDatabase(tkh.newDatabaseName() + tkh.localDatabaseSuffix())
	tkh.cancelButtonPressed()
}

func (tkh *ConnectToDatabaseWindow) selectCreatedDatabase(newDb string) {
	iter, err := tkh.tools.FindValue(newDb, tkh.databaseLists.LocalDatabaseList(), 0)
	if err != nil {
		logging.Logger.Error("connectToDatabaseWindow.selectCreatedDatabase: cannot find value in ListStore", "error", err, "value", newDb)
		return
	}
	tkh.LocalDatabaseNameComboBox.SetActiveIter(iter)
}

func (tkh *ConnectToDatabaseWindow) removeButtonPressed() {
	tkh.LocalDataBaseComboBox.SetSensitive(false)
	tkh.setLocalDatabaseRemoveLabel(tkh.selectedLocalDatabase())
	tkh.changeStack(removeDb)
}

func (tkh *ConnectToDatabaseWindow) applyRemovingButtonPressed() {
	err := os.Remove(filepath.Join(tkh.folderPath(), tkh.selectedLocalDatabase()))
	if err != nil {
		tkh.loginWin.ShowError(err)
	}
	tkh.cancelButtonPressed()
}

func (tkh *ConnectToDatabaseWindow) cancelButtonPressed() {
	tkh.LocalDataBaseComboBox.SetSensitive(true)
	tkh.changeStack(selectDb)
	tkh.RefreshButtonPressed()
	tkh.cleanNewDatabaseName()
}

func (tkh *ConnectToDatabaseWindow) selectedDatabasesChanged() {
	ch := tkh.FoundLocalDatabasesListStore.IterNChildren(nil)
	if ch == 0 {
		tkh.LocalDatabaseRemoveButton.SetSensitive(false)
		return
	}
	tkh.LocalDatabaseRemoveButton.SetSensitive(true)
}
