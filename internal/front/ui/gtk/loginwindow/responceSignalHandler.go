package loginwindow

import "github.com/VladimirRytov/advertisementer/internal/logging"

func (lw *LoginWindow) LoadNetworkDatabases(db []string) {
	logging.Logger.Debug("thickClient: filling network database combobox")
	for _, v := range db {
		iter := lw.ConnectToDatabaseWindow.NetworkDatabasesListStore.Append()
		lw.ConnectToDatabaseWindow.NetworkDatabasesListStore.SetValue(iter, 0, v)
	}
	lw.ConnectToDatabaseWindow.NetworkDataBaseComboBox.SetActive(0)
}

func (lw *LoginWindow) LoadLocalDatabases(db []string) {
	logging.Logger.Debug("thickClient: filling local database combobox")
	for _, v := range db {
		iter := lw.ConnectToDatabaseWindow.LocalDatabasesListStore.Append()
		lw.ConnectToDatabaseWindow.LocalDatabasesListStore.SetValue(iter, 0, v)
	}
	lw.ConnectToDatabaseWindow.LocalDataBaseComboBox.SetActive(0)
}

func (lw *LoginWindow) ShowError(err error) {
	logging.Logger.Debug("thickClient: show error")
	lw.ErrorLabel.SetText(err.Error())
	lw.ErrorReleaver.SetRevealChild(true)
}

func (lw *LoginWindow) ShowDefaultPort(port string) {
	logging.Logger.Debug("thickClient: set default port for network database")
	lw.ConnectToDatabaseWindow.PortEntry.SetPlaceholderText(port)
}
