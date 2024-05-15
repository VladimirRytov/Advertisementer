package windowscontroller

import (
	"strconv"

	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func (g *GUI) SendNetworkDatabases(netdb []string) {
	logging.Logger.Debug("responceController: Sending list of network databasese to the UI")
	g.front.LoadNetworkDatabases(netdb)
}

func (g *GUI) SendLocalDatabases(localdb []string) {
	logging.Logger.Debug("responceController: Sending list of local databasese to the UI")
	g.front.LoadLocalDatabases(localdb)
}

func (g *GUI) SendConnectionError(err error) {
	logging.Logger.Debug("responceController: Sending connection error to the UI", "error", err)
	g.front.UnlockAllLoginForm()
	g.front.ShowError(err)
}

func (g *GUI) SendDefaultPort(port uint) {
	logging.Logger.Debug("responceController: Sending default port to the UI", "port", port)
	val := strconv.Itoa(int(port))
	g.front.ShowDefaultPort(val)
}

func (g *GUI) SendSuccesConnection() {
	g.front.InitAdvertisement()
}
