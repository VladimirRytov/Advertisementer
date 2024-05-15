package loginwindow

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/loginwindow/todatabasewin"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/loginwindow/toserverwin"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"

	"github.com/gotk3/gotk3/gtk"
)

type RequestGate interface {
	LoadConfig(string, any) error
	SaveConfig(string, any) error
	RemoveConfig(string) error

	todatabasewin.Requests
	DefaultNetworkPort(string)
	ConnectToLocalDatabase(*presenter.LocalDSN) error
	ConnectToNetworkDatabase(*presenter.NetworkDataBaseDSN) error
	ConnectToServer(*presenter.ServerDSN) error
}

type LoginwinBuilder interface {
}

type Tools interface {
	FindValue(string, *gtk.ListStore, int) (*gtk.TreeIter, error)
}

type Application interface {
	SetMode(int8)
	Version() string
}

type LoginWindow struct {
	app     Application
	req     RequestGate
	window  *gtk.Window
	version *gtk.Label

	SelectModeStack         *gtk.Stack
	SelectModeStackSwitcher *gtk.StackSwitcher

	ConnectToDatabaseWindow *todatabasewin.ConnectToDatabaseWindow
	ConnectToServerWindow   *toserverwin.ConnectToServerWindow

	SaveConfigCheckButton *gtk.CheckButton
	AutoLoginCheckButton  *gtk.CheckButton

	ErrorReleaver         *gtk.Revealer
	ErrorLabel            *gtk.Label
	StartConnectionButton *gtk.Button
	signalHandler         loginWinHandler
}

func Create(req RequestGate, app Application, tools Tools, DBLists todatabasewin.ListStores, folderSelector todatabasewin.FolderSelector,
	path todatabasewin.PathParser) *LoginWindow {
	loginWinBuilder, err := builder.NewBuilderFromString(builder.LoginWindow)
	if err != nil {
		panic(err)
	}
	lw := new(LoginWindow)
	lw.app = app
	lw.req = req
	lw.buildLoginWindow(loginWinBuilder)
	lw.version.SetText(app.Version())
	lw.bindSignals()

	lw.ConnectToServerWindow = toserverwin.BuildConnectToServerWindow(loginWinBuilder)
	lw.ConnectToServerWindow.PortEntry.SetPlaceholderText("8080")
	lw.ConnectToDatabaseWindow = todatabasewin.BuildConnectToDatabaseWindow(loginWinBuilder, req, tools, DBLists, folderSelector, lw, path)
	lw.ConnectToDatabaseWindow.BindSignals()
	lw.req = req
	lw.SelectModeStack.SetVisibleChildName("ConnectToDatabase")
	lw.ErrorReleaver.SetRevealChild(false)
	lw.window.SetTitle("Окно входа")
	lw.window.SetIconName("advertisementer.png")
	return lw
}

func (lw *LoginWindow) Close() {
	lw.window.Close()
}

func (lw *LoginWindow) buildLoginWindow(builder *builder.Builder) {
	lw.SelectModeStack = builder.FetchStack("SelectModeStack")
	lw.version = builder.FetchLabel("VersionLabel")
	lw.SelectModeStackSwitcher = builder.FetchStackSwitcher("SelectModeStackSwitcher")
	lw.window = builder.FetchWindow("LoginWindow")
	lw.SaveConfigCheckButton = builder.FetchCheckButton("SaveConfigCheckButton")
	lw.AutoLoginCheckButton = builder.FetchCheckButton("AutoLOginCheckButton")
	lw.ErrorReleaver = builder.FetchRevealer("ErrorReleaver")
	lw.ErrorLabel = builder.FetchLabel("ErrorLabel")
	lw.StartConnectionButton = builder.FetchButton("StartConnectionButton")

}
