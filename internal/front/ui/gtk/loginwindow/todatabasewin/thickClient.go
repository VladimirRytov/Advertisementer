package todatabasewin

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func (thk *ConnectToDatabaseWindow) serverAddress() string {
	logging.Logger.Debug("ConnectToDatabaseWindow: getting server address")
	return thk.URLEntry.GetLayout().GetText()
}

func (thk *ConnectToDatabaseWindow) SetServerAddress(addr string) {
	logging.Logger.Debug("ConnectToDatabaseWindow: setting server address")
	thk.URLEntry.SetText(addr)
}

func (thk *ConnectToDatabaseWindow) port() string {
	logging.Logger.Debug("ConnectToDatabaseWindow: getting port number")
	return thk.PortEntry.GetLayout().GetText()
}

func (thk *ConnectToDatabaseWindow) SetPort(port string) {
	logging.Logger.Debug("ConnectToDatabaseWindow: setting port number")
	thk.PortEntry.SetText(port)
}

func (thk *ConnectToDatabaseWindow) networkDataBase() string {
	logging.Logger.Debug("ConnectToDatabaseWindow: getting database name")
	return thk.DatabaseNameEntry.GetLayout().GetText()
}

func (thk *ConnectToDatabaseWindow) SetDataBase(db string) {
	logging.Logger.Debug("ConnectToDatabaseWindow: setting database name")
	thk.DatabaseNameEntry.SetText(db)
}

func (thk *ConnectToDatabaseWindow) login() string {
	logging.Logger.Debug("ConnectToDatabaseWindow: getting username")
	return thk.LoginEntry.GetLayout().GetText()
}

func (thk *ConnectToDatabaseWindow) SetLogin(login string) {
	logging.Logger.Debug("ConnectToDatabaseWindow: setting username")
	thk.LoginEntry.SetText(login)
}

func (thk *ConnectToDatabaseWindow) password() string {
	logging.Logger.Debug("ConnectToDatabaseWindow: getting password")
	text, _ := thk.PasswordBuffer.GetText()
	return text
}

func (thk *ConnectToDatabaseWindow) SetPassword(password string) {
	logging.Logger.Debug("ConnectToDatabaseWindow: setting password")
	thk.PasswordBuffer.SetText(password)
}

func (thk *ConnectToDatabaseWindow) SelectNetworkDatabase(db string) {
	iter, err := thk.tools.FindValue(db, thk.NetworkDatabasesListStore, 0)
	if err != nil {
		logging.Logger.Error("toDatabaseMode.SelectDatabase: database not found")
		return
	}
	thk.NetworkDataBaseComboBox.SetActiveIter(iter)
}

func (thk *ConnectToDatabaseWindow) SelectedDatabase() string {
	switch thk.SelectedDataStorageType() {
	case "NetworkDB":
		logging.Logger.Debug("ConnectToDatabaseWindow: getting Network database type")
		return thk.NetworkDataBaseComboBox.GetActiveID()
	case "LocalDB":
		logging.Logger.Debug("ConnectToDatabaseWindow: getting Local database type")
		return thk.LocalDataBaseComboBox.GetActiveID()
	}
	logging.Logger.Error("ConnectToDatabaseWindow: cannot get database type")
	return ""
}

func (thk *ConnectToDatabaseWindow) SelectedDataStorageType() string {
	logging.Logger.Debug("ConnectToDatabaseWindow: getting storage type")
	return thk.DataStorageTypeStack.GetVisibleChildName()
}

func (thk *ConnectToDatabaseWindow) LocalDBAuthorizationForm() *presenter.LocalDSN {
	logging.Logger.Debug("ConnectToDatabaseWindow: fetching local database authorization form")
	return &presenter.LocalDSN{
		Path: thk.folderPath(),
		Name: thk.selectedLocalDatabase(),
		Type: thk.SelectedDatabase(),
	}
}

func (thk *ConnectToDatabaseWindow) NetworkDBAuthorizationForm() *presenter.NetworkDataBaseDSN {
	logging.Logger.Debug("ConnectToDatabaseWindow: fetching network database authorization form")
	return &presenter.NetworkDataBaseDSN{
		DatabaseName: thk.SelectedDatabase(),
		Source:       thk.serverAddress(),
		Port:         thk.port(),
		DataBase:     thk.networkDataBase(),
		UserName:     thk.login(),
		Password:     thk.password(),
	}
}

func (thk *ConnectToDatabaseWindow) LockAll() {
	thk.DataStorageTypeStack.SetSensitive(false)
	thk.DataStorageTypeStackSwitcher.SetSensitive(false)
}

func (thk *ConnectToDatabaseWindow) UnlockAll() {
	thk.DataStorageTypeStack.SetSensitive(true)
	thk.DataStorageTypeStackSwitcher.SetSensitive(true)
}
