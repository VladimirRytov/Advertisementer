package todatabasewin

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/application"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"

	"github.com/gotk3/gotk3/gtk"
)

const (
	selectDb = "SelectDatabase"
	createDb = "CreateDatabase"
	removeDb = "RemoveDatabase"
)

type Requests interface {
	CreateFile(string) error
	DefaultNetworkPort(string)
}

type LoginWindow interface {
	Window() *gtk.Window
	ShowError(err error)
}

type PathParser interface {
	ParsePath(string) string
}

type Tools interface {
	FindValue(string, *gtk.ListStore, int) (*gtk.TreeIter, error)
}

type FolderSelector interface {
	NewFolderChooseDialog(winLabel string, parent gtk.IWindow) (application.FileDialoger, error)
}

type ListStores interface {
	LocalDatabaseList() *gtk.ListStore
	AppendLocalDatabase(string)
	RemoveLocalDatabase(string)
	ClearLocalDatabaseList()
}

type ConnectToDatabaseWindow struct {
	req            Requests
	databaseLists  ListStores
	folderSelector FolderSelector
	tools          Tools
	loginWin       LoginWindow
	pathParser     PathParser

	DataStorageTypeStack         *gtk.Stack
	DataStorageTypeStackSwitcher *gtk.StackSwitcher

	NetworkDataBaseLabel    *gtk.Label
	NetworkDataBaseComboBox *gtk.ComboBox

	LocalDataBaseLabel    *gtk.Label
	LocalDataBaseComboBox *gtk.ComboBox

	LocalDatabasePathLabel  *gtk.Label
	LocalDatabasePathEntry  *gtk.Entry
	LocalDatabasePathButton *gtk.Button

	LocalDatabaseStack         *gtk.Stack
	LocalDatabaseNameLabel     *gtk.Label
	LocalDatabaseNameComboBox  *gtk.ComboBox
	LocalDatabaseCreateButton  *gtk.Button
	LocalDatabaseRefreshButton *gtk.Button
	LocalDatabaseRemoveButton  *gtk.Button

	LocalDatabaseCreateEntry         *gtk.Entry
	LocalDatabaseApplyCreationButton *gtk.Button
	LocalDatabaseCancelCreateButton  *gtk.Button

	LocalDatabaseRemoveLabel        *gtk.Label
	LocalDatabaseApplyRemoveButton  *gtk.Button
	LocalDatabaseCancelRemoveButton *gtk.Button

	URLLabel *gtk.Label
	URLEntry *gtk.Entry

	PortLabel *gtk.Label
	PortEntry *gtk.Entry

	DatabaseNameLabel *gtk.Label
	DatabaseNameEntry *gtk.Entry

	LoginLabel *gtk.Label
	LoginEntry *gtk.Entry

	PasswordLabel  *gtk.Label
	PasswordEntry  *gtk.Entry
	PasswordBuffer *gtk.EntryBuffer

	NetworkDatabasesListStore    *gtk.ListStore
	LocalDatabasesListStore      *gtk.ListStore
	FoundLocalDatabasesListStore *gtk.ListStore
	signals                      ConnectToDatabaseWindowHandler
}

func BuildConnectToDatabaseWindow(builder *builder.Builder, req Requests, tools Tools, lists ListStores, folderSelector FolderSelector,
	loginWin LoginWindow, path PathParser) *ConnectToDatabaseWindow {
	localDBSelector := &ConnectToDatabaseWindow{
		req:                          req,
		tools:                        tools,
		pathParser:                   path,
		DataStorageTypeStack:         builder.FetchStack("DataStorageTypeStack"),
		DataStorageTypeStackSwitcher: builder.FetchStackSwitcher("DataStorageTypeStackSwitcher"),

		LocalDataBaseLabel:    builder.FetchLabel("ConnectToDatabaseLocalDataBaseLabel"),
		LocalDataBaseComboBox: builder.FetchComboBox("ConnectToDatabaseLocalDataBaseComboBox"),

		LocalDatabasePathLabel:  builder.FetchLabel("ConnectToDatabasePathToFolderLabel"),
		LocalDatabasePathEntry:  builder.FetchEntry("ConnectToDatabasePathToFolderEntry"),
		LocalDatabasePathButton: builder.FetchButton("ConnectToDatabasePathToFolderButton"),

		LocalDatabaseStack:         builder.FetchStack("LocalDatabaseCasesStack"),
		LocalDatabaseNameLabel:     builder.FetchLabel("ConnectToDatabaseNameLocalDatabaseLabel"),
		LocalDatabaseNameComboBox:  builder.FetchComboBox("ConnectToDatabaseSelectLocalDatabaseCombobox"),
		LocalDatabaseCreateButton:  builder.FetchButton("ConnectToDatabaseCreateLocalDatabaseButton"),
		LocalDatabaseRefreshButton: builder.FetchButton("LocalDatabaseRefreshButton"),
		LocalDatabaseRemoveButton:  builder.FetchButton("LocalDatabaseRemoveButton"),

		LocalDatabaseCreateEntry:         builder.FetchEntry("LocalDatabaseCreateEntry"),
		LocalDatabaseApplyCreationButton: builder.FetchButton("LocalDatabaseApplyCreationButton"),
		LocalDatabaseCancelCreateButton:  builder.FetchButton("LocalDatabaseCancelCreationButton"),

		LocalDatabaseRemoveLabel:        builder.FetchLabel("LocalDatabaseRemoveLabel"),
		LocalDatabaseApplyRemoveButton:  builder.FetchButton("LocalDatabaseApplyRemovingButton"),
		LocalDatabaseCancelRemoveButton: builder.FetchButton("LocalDatabaseCancelRemovingButton"),

		NetworkDataBaseLabel:    builder.FetchLabel("ConnectToDatabaseNetworkDataBaseLabel"),
		NetworkDataBaseComboBox: builder.FetchComboBox("ConnectToDatabaseNetworkDataBaseComboBox"),

		NetworkDatabasesListStore: builder.FetchListStore("ConnectToDatabaseNetworkDatabasesListStore"),
		LocalDatabasesListStore:   builder.FetchListStore("ConnectToDatabaseLocalDatabasesListStore"),

		URLLabel: builder.FetchLabel("ConnectToDatabaseURLLabel"),
		URLEntry: builder.FetchEntry("ConnectToDatabaseURLEntry"),

		PortLabel: builder.FetchLabel("ConnectToDatabasePortLabel"),
		PortEntry: builder.FetchEntry("ConnectToDatabasePortEntry"),

		DatabaseNameLabel: builder.FetchLabel("ConnectToDatabaseDatabaseNameLabel"),
		DatabaseNameEntry: builder.FetchEntry("ConnectToDatabaseDatabaseNameEntry"),

		LoginLabel: builder.FetchLabel("ConnectToDatabaseLoginLabel"),
		LoginEntry: builder.FetchEntry("ConnectToDatabaseLoginEntry"),

		PasswordLabel: builder.FetchLabel("ConnectToDatabasePasswordLabel"),
		PasswordEntry: builder.FetchEntry("ConnectToDatabasePasswordEntry"),

		PasswordBuffer: builder.FetchEntryBuffer("ConnectToDatabasePasswordBuffer"),
	}
	localDBSelector.databaseLists = lists
	localDBSelector.folderSelector = folderSelector
	localDBSelector.req = req
	localDBSelector.tools = tools
	localDBSelector.loginWin = loginWin
	localDBSelector.FoundLocalDatabasesListStore = lists.LocalDatabaseList()
	localDBSelector.LocalDatabaseNameComboBox.SetModel(localDBSelector.FoundLocalDatabasesListStore)
	localDBSelector.LocalDatabaseRemoveButton.SetSensitive(false)
	return localDBSelector
}
