package advertisementwindow

import (
	"time"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/application"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type NotebookCreator interface {
	NewNotebook(n *builder.Builder, advWin application.AdvertisementsWindow) application.Notebook
}

type SaveFileDialoger interface {
	NewSaveDialog(winLabel string, parent gtk.IWindow) (application.FileDialoger, error)
}
type AdvertisementsWindow struct {
	req         RequestGate
	conv        Converter
	app         Application
	dialogMaker SaveFileDialoger

	signalHandler AdvertisementsWindowHandler
	MainWindow    *gtk.Window

	journalLabel    *gtk.Label
	journalCombobox *gtk.ComboBox

	searchEntry *gtk.SearchEntry

	filterLabel                 *gtk.Label
	filterShowAllRadioButton    *gtk.RadioButton
	filterShowActualRadioButton *gtk.RadioButton
	filterDataRangeLabel        *gtk.Label
	filterMenuButton            *gtk.MenuButton
	Filter

	resetFilterIcoButton *gtk.Button

	costRateEntry          *gtk.Entry
	costRateChangeButton   *gtk.Button
	costRateSettingsButton *gtk.Button
	advReportLabel         *gtk.Label
	advReportButton        *gtk.Button

	settingButton   *gtk.MenuButton
	settingsPopover *gtk.PopoverMenu

	NoteBook application.Notebook

	exportDataButton *gtk.Button
	importDataButton *gtk.Button
	exitButton       *gtk.Button

	connectionStatusBox    *gtk.Box
	connectionStatusButton *gtk.Button
	establiShedImage       *gtk.Image
	lostConnImage          *gtk.Image

	errorPopover *gtk.Popover
	errorLabel   *gtk.Label
}

func Create(reqGate RequestGate, app Application, notebook NotebookCreator, conv Converter, objMaker SaveFileDialoger) *AdvertisementsWindow {
	logging.Logger.Debug("advertisementWindow: creating window")
	buildFile, err := builder.NewBuilderFromString(builder.MainWindow)
	if err != nil {
		logging.Logger.Error("advertisementWindow: got panic")
		panic(err)
	}

	aw := new(AdvertisementsWindow)
	aw.req = reqGate
	aw.dialogMaker = objMaker
	aw.NoteBook = notebook.NewNotebook(buildFile, aw)
	aw.app = app
	aw.conv = conv
	aw.build(buildFile)
	if aw.app.Mode() == application.ThinMode {
		aw.ShowConnecionStatus()
		go aw.CheckingConnection()
	}

	show := &presenter.ShowData{}
	err = aw.req.LoadConfig("ShowData", show)
	if err != nil {
		logging.Logger.Error("AdvertisementsWindow.Initialization: cannot load config", "error", err)
	}
	aw.filterShowActualRadioButton.SetActive(show.Actual)
	aw.bindSignals()
	aw.Window().Connect("destroy", aw.Window().Destroy)
	aw.MainWindow.SetTitle("Advertisementer")
	return aw
}

func (aw *AdvertisementsWindow) build(buildFile *builder.Builder) {
	logging.Logger.Debug("advertisementWindow: building window")
	aw.MainWindow = buildFile.FetchWindow("mainWindow")
	aw.journalLabel = buildFile.FetchLabel("JournalLabel")
	aw.journalCombobox = buildFile.FetchComboBox("JournalCombobox")
	aw.searchEntry = buildFile.FetchSearchEntry("SearchEntry")
	aw.filterLabel = buildFile.FetchLabel("FilterLabel")
	aw.filterShowAllRadioButton = buildFile.FetchRadioButton("ShowAllRadioBtn")
	aw.filterShowActualRadioButton = buildFile.FetchRadioButton("ShowActualRadioBtn")
	aw.filterDataRangeLabel = buildFile.FetchLabel("FilterDateRangeLabel")
	aw.filterMenuButton = buildFile.FetchMenuButton("FilterMenuButton")
	aw.resetFilterIcoButton = buildFile.FetchButton("ResetFilterButton")
	aw.settingButton = buildFile.FetchMenuButton("SettingsButton")
	aw.settingsPopover = buildFile.FetchPopoverMenu("SettingsPopover")
	aw.errorPopover = buildFile.FetchPopover("ErrorPopover")
	aw.errorLabel = buildFile.FetchLabel("ErrorPopoverLabel")
	aw.costRateEntry = buildFile.FetchEntry("SelectedCostRateEntry")
	aw.costRateChangeButton = buildFile.FetchButton("ChangeCostRateButton")
	aw.costRateSettingsButton = buildFile.FetchButton("CalculatePaymentSettingsButton")
	aw.advReportLabel = buildFile.FetchLabel("ExportLabel")
	aw.advReportButton = buildFile.FetchButton("AdvertisementReportButton")
	aw.exportDataButton = buildFile.FetchButton("AdvertisementsExportData")
	aw.importDataButton = buildFile.FetchButton("AdvertisementsImportData")
	aw.exitButton = buildFile.FetchButton("LogOutButton")
	aw.connectionStatusBox = buildFile.FetchBox("ConnectionStatusBox")
	aw.connectionStatusButton = buildFile.FetchButton("ConnectionStatusButton")
	aw.establiShedImage = buildFile.FetchImage("ConnectionEstablished")
	aw.lostConnImage = buildFile.FetchImage("ConnectionLost")
	aw.BuildFilter(buildFile)
}

func (aw *AdvertisementsWindow) ShowConnecionStatus() {
	aw.connectionStatusBox.SetVisible(true)
}

func (aw *AdvertisementsWindow) Show() {
	aw.NoteBook.EnableSidebarsFilters(true)
	aw.MainWindow.Show()
}

func (aw *AdvertisementsWindow) Close() {
	aw.MainWindow.Close()
}

func (aw *AdvertisementsWindow) Hide() {
	aw.MainWindow.Hide()
}

func (aw *AdvertisementsWindow) ShowPopover(errText string, point gtk.IWidget) {
	aw.errorLabel.SetText(errText)
	aw.errorPopover.SetRelativeTo(point)
	aw.errorPopover.Popup()
}

func (aw *AdvertisementsWindow) CheckingConnection() {
	glib.IdleAdd(func() {
		aw.SetConnectionStatus(aw.app.ConnectionStatus())
	})
	t := time.NewTicker(5 * time.Second)
	defer t.Stop()
	for range t.C {
		glib.IdleAdd(func() {
			aw.SetConnectionStatus(aw.app.ConnectionStatus())
		})
	}
}
