package loginwindow

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/application"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
)

type loginWinHandler struct {
	clickStartConnection glib.SignalHandle
	saveconfigToggled    glib.SignalHandle
	autoLoginToggled     glib.SignalHandle
}

func (lw *LoginWindow) bindSignals() {
	logging.Logger.Debug("loginWindow: binding loginwin signals")
	lw.signalHandler.clickStartConnection = lw.StartConnectionButton.Connect("clicked", lw.clickStartConnection)
	lw.signalHandler.saveconfigToggled = lw.SaveConfigCheckButton.Connect("toggled", lw.saveconfigToggled)
	lw.signalHandler.autoLoginToggled = lw.AutoLoginCheckButton.Connect("toggled", lw.autoLoginToggled)
}

func (lw *LoginWindow) clickStartConnection() {
	logging.Logger.Info("loginWindow: start connecting to DataBase")
	lw.LockAll()
	switch lw.selectedClientType() {
	case "ConnectToServer":

		lw.app.SetMode(application.ThinMode)
		logging.Logger.Debug("loginWindow: start connecting to DataBase as Thin Client")
		dsn := lw.ConnectToServerWindow.AuthorizationForm()
		lw.req.ConnectToServer(&dsn)
	case "ConnectToDatabase":
		logging.Logger.Debug("loginWindow: start connecting to DataBase as Thick Client")
		lw.app.SetMode(application.ThickMode)
		switch lw.ConnectToDatabaseWindow.SelectedDataStorageType() {
		case NetworkDatabase:
			logging.Logger.Debug("loginWindow: fetching authorization date from network form")
			dsn := lw.ConnectToDatabaseWindow.NetworkDBAuthorizationForm()
			lw.req.ConnectToNetworkDatabase(dsn)
		case LocalDatabase:
			logging.Logger.Debug("loginWindow: fetching authorization date from local form")
			dsn := lw.ConnectToDatabaseWindow.LocalDBAuthorizationForm()
			lw.req.ConnectToLocalDatabase(dsn)
		}
	}

}
func (lw *LoginWindow) autoLoginToggled() {
	lw.SaveConfigCheckButton.SetSensitive(!lw.AutoLoginCheckButton.GetActive())
	if !lw.SaveConfigCheckButton.GetActive() {
		lw.SaveConfigCheckButton.SetActive(lw.AutoLoginCheckButton.GetActive())
	}
}

func (lw *LoginWindow) saveconfigToggled() {
	logging.Logger.Debug("saveconfigToggled: fetching authorization date from local form")
	if !lw.SaveConfigCheck() {
		err := lw.req.RemoveConfig(LastSuccefullConnection)
		if err != nil {
			logging.Logger.Debug("saveconfigToggled: cannot remove LastConnectionInfo")
		}
	}
}
