package toserverwin

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func (tn *ConnectToServerWindow) serverAddress() string {
	return tn.URLEntry.GetLayout().GetText()
}

func (tn *ConnectToServerWindow) SetserverAddress(addr string) {
	tn.URLEntry.SetText(addr)
}

func (tn *ConnectToServerWindow) login() string {
	return tn.LoginEntry.GetLayout().GetText()
}

func (tn *ConnectToServerWindow) Setlogin(login string) {
	tn.LoginEntry.SetText(login)
}

func (tn *ConnectToServerWindow) port() string {
	return tn.PortEntry.GetLayout().GetText()
}

func (tn *ConnectToServerWindow) Setport(port string) {
	b, err := tn.PortEntry.GetBuffer()
	if err != nil {
		logging.Logger.Error("connectToServerWindow.Setport: cannot get port buffer", "error", err)
	}
	b.SetText(port)
}

func (tn *ConnectToServerWindow) password() string {
	text, _ := tn.PasswordBuffer.GetText()
	return text
}

func (tn *ConnectToServerWindow) Setpassword(pass string) {
	tn.PasswordBuffer.SetText(pass)
}

func (tn *ConnectToServerWindow) AuthorizationForm() presenter.ServerDSN {
	logging.Logger.Debug("ConnectToDatabaseWindow: fetching network database authorization form")
	return presenter.ServerDSN{
		DatabaseName: "Server",
		Source:       tn.serverAddress(),
		Port:         tn.port(),
		UserName:     tn.login(),
		Password:     tn.password(),
	}
}
