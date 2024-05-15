package loginwindow

import (
	"os"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/gtk"
)

const (
	LastSuccefullConnection = "LastConnectionInfo"
	NetworkDatabase         = "NetworkDB"
	LocalDatabase           = "LocalDB"
	Server                  = "Server"
	DbType                  = "DatabaseType"
)

func (lw *LoginWindow) selectedClientType() string {
	return lw.SelectModeStack.GetVisibleChildName()
}
func (lw *LoginWindow) Window() *gtk.Window {
	return lw.window
}
func (lw *LoginWindow) LockAll() {
	lw.StartConnectionButton.SetSensitive(false)
}

func (lw *LoginWindow) UnlockAll() {
	lw.StartConnectionButton.SetSensitive(true)
}

func (lw *LoginWindow) Hide() {
	lw.window.SetVisible(false)
}

func (lw *LoginWindow) Show() {
	lw.window.SetVisible(true)
}

func (lw *LoginWindow) SaveConfigCheck() bool {
	return lw.SaveConfigCheckButton.GetActive()
}

func (lw *LoginWindow) SaveConfigs() {
	logging.Logger.Debug("loginWindow.SaveConfigs: StartSavingConfigs")
	if !lw.SaveConfigCheck() {
		return
	}
	connectionType := &presenter.LastDatabaseConnection{
		DatabaseType: lw.selectedClientType(),
		ConfigSaved:  lw.SaveConfigCheck(),
		AutoLogin:    lw.AutoLoginCheck(),
	}
	err := lw.req.SaveConfig(LastSuccefullConnection, connectionType)
	if err != nil {
		logging.Logger.Error("loginWindow.SaveConfig: cannot save config", "error", err)
	}
	switch lw.selectedClientType() {
	case "ConnectToServer":
		logging.Logger.Debug("loginWindow.SaveConfig: saving server authorization form")
		dsn := lw.ConnectToServerWindow.AuthorizationForm()
		err := lw.req.SaveConfig(Server, &dsn)
		if err != nil {
			logging.Logger.Error("loginWindow.SaveConfig: cannot save config", "error", err)
		}

	case "ConnectToDatabase":
		logging.Logger.Debug("loginWindow: saving Database authorization form")

		err := lw.req.SaveConfig(DbType, lw.ConnectToDatabaseWindow.SelectedDataStorageType())
		if err != nil {
			logging.Logger.Error("loginWindow.SaveConfig: cannot save config", "error", err)
		}
	}

	switch lw.ConnectToDatabaseWindow.SelectedDataStorageType() {
	case NetworkDatabase:
		logging.Logger.Debug("loginWindow: fetching authorization date from network form")
		dsn := lw.ConnectToDatabaseWindow.NetworkDBAuthorizationForm()
		err := lw.req.SaveConfig(NetworkDatabase, dsn)
		if err != nil {
			logging.Logger.Error("loginWindow.SaveConfig: cannot save config", "error", err)
		}
	case LocalDatabase:
		logging.Logger.Debug("loginWindow: fetching authorization date from local form")
		dsn := lw.ConnectToDatabaseWindow.LocalDBAuthorizationForm()
		err := lw.req.SaveConfig(LocalDatabase, dsn)
		if err != nil {
			logging.Logger.Error("loginWindow.SaveConfig: cannot save config", "error", err)
		}
	}
}

func (lw *LoginWindow) LoadConfigs() {
	connectionType := &presenter.LastDatabaseConnection{}
	err := lw.req.LoadConfig(LastSuccefullConnection, connectionType)
	if err != nil {
		logging.Logger.Error("loginwindow.LoadConfigs: cannot load configs", "error", err)
		lw.LoadDefaultLocal()
		return
	}

	lw.SelectModeStack.SetVisibleChildName(connectionType.DatabaseType)
	lw.LoadDefaultLocal()
	switch connectionType.DatabaseType {
	case "ConnectToServer":
		lw.LoadServerLastConnection()

	case "ConnectToDatabase":
		var dbType string
		err := lw.req.LoadConfig(DbType, &dbType)
		if err != nil {
			logging.Logger.Error("loginwindow.LoadConfigs: cannot load dbType config", "error", err)
			return
		}
		lw.ConnectToDatabaseWindow.DataStorageTypeStack.SetVisibleChildName(dbType)
		switch dbType {
		case NetworkDatabase:
			lw.LoadNetworkLastConnection()
		case LocalDatabase:
			lw.LoadLocalLastConnection()
		}
	}
	lw.SaveConfigCheckButton.SetActive(connectionType.ConfigSaved)
}

func (lw *LoginWindow) AutoLoginCheck() bool {
	return lw.AutoLoginCheckButton.GetActive()
}

func (lw *LoginWindow) LoadServerLastConnection() {
	dsn := presenter.ServerDSN{}
	err := lw.req.LoadConfig(Server, &dsn)
	if err != nil {
		logging.Logger.Error("loginwindow.LoadServerLastConnection: cannot load Server configs", "error", err)
		return
	}
	lw.ConnectToServerWindow.Setlogin(dsn.UserName)
	lw.ConnectToServerWindow.Setpassword(dsn.Password)
	lw.ConnectToServerWindow.Setport(dsn.Port)
	lw.ConnectToServerWindow.SetserverAddress(dsn.Source)
}

func (lw *LoginWindow) LoadNetworkLastConnection() {
	dsn := presenter.NetworkDataBaseDSN{}
	err := lw.req.LoadConfig(NetworkDatabase, &dsn)
	if err != nil {
		logging.Logger.Error("loginwindow.LoadNetworkLastConnection: cannot load NetworkDB configs", "error", err)
		return
	}
	lw.ConnectToDatabaseWindow.SelectNetworkDatabase(dsn.DatabaseName)
	lw.ConnectToDatabaseWindow.SetServerAddress(dsn.Source)
	lw.ConnectToDatabaseWindow.SetPort(dsn.Port)
	lw.ConnectToDatabaseWindow.SetDataBase(dsn.DataBase)
	lw.ConnectToDatabaseWindow.SetLogin(dsn.UserName)
	lw.ConnectToDatabaseWindow.SetPassword(dsn.Password)
}

func (lw *LoginWindow) LoadLocalLastConnection() {
	dsn := presenter.LocalDSN{}
	err := lw.req.LoadConfig(LocalDatabase, &dsn)
	if err != nil {
		logging.Logger.Error("loginwindow.LoadLocalLastConnection: cannot load LocalDB configs", "error", err)
		lw.LoadDefaultLocal()
	}
	lw.ConnectToDatabaseWindow.SetLocalfolderPath(dsn.Path)
	lw.ConnectToDatabaseWindow.SetLastLocalConnection(dsn.Name, dsn.Type)
}

func (lw *LoginWindow) LoadDefaultLocal() {
	dbDir, err := os.UserHomeDir()
	if err != nil {
		dbDir, _ = os.Getwd()
	}
	lw.ConnectToDatabaseWindow.SetLocalfolderPath(dbDir)
	lw.ConnectToDatabaseWindow.RefreshButtonPressed()
}
