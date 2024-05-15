package todatabasewin

import (
	"errors"
	"io/fs"
	"os"
	"strings"

	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func (tkh *ConnectToDatabaseWindow) folderPath() string {
	return tkh.LocalDatabasePathEntry.GetLayout().GetText()
}

func (tkh *ConnectToDatabaseWindow) changeStack(stackName string) {
	tkh.LocalDatabaseStack.SetVisibleChildName(stackName)
}

func (tkh *ConnectToDatabaseWindow) selectedLocalDatabaseType() string {
	return tkh.LocalDataBaseComboBox.GetActiveID()
}

func (tkh *ConnectToDatabaseWindow) selectedLocalDatabase() string {
	return tkh.LocalDatabaseNameComboBox.GetActiveID()
}

func (tkh *ConnectToDatabaseWindow) databaseList() ([]string, error) {
	selectedpath := tkh.folderPath()
	stat, err := os.Stat(selectedpath)
	if err != nil {
		return nil, err
	}
	if !stat.IsDir() {
		return nil, errors.New("выбранный файл не является каталогом")
	}
	dir := os.DirFS(selectedpath)
	return fs.Glob(dir, "*"+tkh.localDatabaseSuffix())
}

func (tkh *ConnectToDatabaseWindow) fillDatabaseList(databases []string) {
	for i := range databases {
		tkh.databaseLists.AppendLocalDatabase(databases[i])
	}
}

func (tkh *ConnectToDatabaseWindow) localDatabaseSuffix() string {
	return strings.ToLower("." + tkh.selectedLocalDatabaseType())
}

func (tkh *ConnectToDatabaseWindow) SetLocalfolderPath(path string) {
	tkh.LocalDatabasePathEntry.SetText(path)
}

func (tkh *ConnectToDatabaseWindow) setLocalDatabaseRemoveLabel(dbName string) {
	tkh.LocalDatabaseRemoveLabel.SetText("Удалить " + dbName + " ?")
}

func (tkh *ConnectToDatabaseWindow) SetLastLocalConnection(dbName, dbType string) {
	iter, err := tkh.tools.FindValue(dbType, tkh.LocalDatabasesListStore, 0)
	if err != nil {
		logging.Logger.Error("connectToDatabaseWindow.SetLastLocalConnection: cannot find dbType", "error", err)
		tkh.LocalDataBaseComboBox.SetActive(0)
		return
	}
	tkh.LocalDataBaseComboBox.SetActiveIter(iter)
	tkh.RefreshButtonPressed()

	iter, err = tkh.tools.FindValue(dbName, tkh.FoundLocalDatabasesListStore, 0)
	if err != nil {
		logging.Logger.Error("connectToDatabaseWindow.SetLastLocalConnection: cannot find dbName", "error", err)
		tkh.LocalDatabaseNameComboBox.SetActive(0)
		return
	}
	tkh.LocalDatabaseNameComboBox.SetActiveIter(iter)
}

func (tkh *ConnectToDatabaseWindow) newDatabaseName() string {
	return tkh.LocalDatabaseCreateEntry.GetLayout().GetText()
}

func (tkh *ConnectToDatabaseWindow) cleanNewDatabaseName() {
	tkh.LocalDatabaseCreateEntry.SetText("")
}
