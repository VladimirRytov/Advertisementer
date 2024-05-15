package newadvreport

import (
	"context"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/application"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/gtk"
)

type ReportRequests interface {
	LoadConfig(string, any) error
	SaveConfig(string, any) error
	CreateAfvertisementReport(context.Context, *presenter.ReportParams) error
}

type DataConverter interface {
	YearMonthDayToString(uint, uint, uint) string
	ParsePath(string) string
}

type Application interface {
	Mode() int8
	NewErrorWindow(error)
	CreateProgressWindow() application.ProgressWindow
}

type FileSaveDialog interface {
	NewFolderChooseDialog(winLabel string, parent gtk.IWindow) (application.FileDialoger, error)
}

type NewAdvertisementReport struct {
	app           Application
	req           ReportRequests
	conv          DataConverter
	dialogCreator FileSaveDialog

	window             *gtk.Window
	sourceFolderBox    *gtk.Box
	sourceFolderLabel  *gtk.Label
	sourceFolderEntry  *gtk.Entry
	sourceFolderButton *gtk.Button

	deployFolderLabel  *gtk.Label
	deployFolderEntry  *gtk.Entry
	deployFolderButton *gtk.Button

	fileFormatLabel     *gtk.Label
	fileFormarExelRadio *gtk.RadioButton

	dateFilterStackSwitcher *gtk.StackSwitcher
	dateFilterStack         *gtk.Stack

	selectedDateLabel       *gtk.Label
	selectedDateEntry       *gtk.Entry
	selectedDateResetButton *gtk.Button
	selectedDateCalendar    *gtk.Calendar

	fromDateLabel       *gtk.Label
	fromDateEntry       *gtk.Entry
	fromDateResetButton *gtk.Button
	fromDateRevealer    *gtk.Revealer
	fromDateCalendar    *gtk.Calendar

	toDateLabel       *gtk.Label
	toDateEntry       *gtk.Entry
	toDateResetButton *gtk.Button
	toDateRevealer    *gtk.Revealer
	toDateCalendar    *gtk.Calendar

	createButton *gtk.Button
	signals      SignalHandler
}

func Create(reqGate ReportRequests, conv DataConverter, app Application, objmaker FileSaveDialog) *NewAdvertisementReport {
	builder, err := builder.NewBuilderFromString(builder.AdvReport)
	if err != nil {
		panic(err)
	}
	nar := new(NewAdvertisementReport)
	nar.req = reqGate
	nar.dialogCreator = objmaker
	nar.app = app
	nar.conv = conv
	nar.build(builder)
	nar.bindSignals()
	var sourcePath string
	if app.Mode() == application.ThickMode {
		err = nar.req.LoadConfig("sourcePath", &sourcePath)
		if err != nil {
			logging.Logger.Error("newAdvertisementReport.sourceDirButtonPressed: cannot load config 'sourcePath'", "error", err)
		}
		nar.setSourceFolder(sourcePath)
		nar.sourceFolderBox.SetVisible(true)
	}

	var deployPath string
	err = nar.req.LoadConfig("deployPath", &deployPath)
	if err != nil {
		logging.Logger.Error("newAdvertisementReport.sourceDirButtonPressed: cannot load config 'deployPath'", "error", err)
	}
	nar.setDeployFolder(deployPath)
	nar.selectedDateCalendar.SelectMonth(uint(time.Now().Month()-1), uint(time.Now().Year()))
	nar.selectedDateCalendar.SelectDay(uint(time.Now().Day()))
	nar.dateSelected()
	nar.window.SetTitle("Экспорт объявлений")
	return nar
}

func (nar *NewAdvertisementReport) build(builder *builder.Builder) {
	nar.window = builder.FetchWindow("AdvertisementReportWindow")
	nar.sourceFolderBox = builder.FetchBox("SourceFolderBox")
	nar.sourceFolderLabel = builder.FetchLabel("SourceDirLabel")
	nar.sourceFolderEntry = builder.FetchEntry("SourceDirEntry")
	nar.sourceFolderButton = builder.FetchButton("SourceDirButton")
	nar.deployFolderLabel = builder.FetchLabel("DestDirLabel")
	nar.deployFolderEntry = builder.FetchEntry("DestDirEntry")
	nar.deployFolderButton = builder.FetchButton("DestDirButton")
	nar.fileFormatLabel = builder.FetchLabel("OutputFileTypeLabel")
	nar.fileFormarExelRadio = builder.FetchRadioButton("OutputFileTypeExelRadio")
	nar.dateFilterStackSwitcher = builder.FetchStackSwitcher("filterPopStackSwitcher")
	nar.dateFilterStack = builder.FetchStack("filterPopStack")
	nar.selectedDateLabel = builder.FetchLabel("SelectDateLabel")
	nar.selectedDateEntry = builder.FetchEntry("SelectDateEntry")
	nar.selectedDateResetButton = builder.FetchButton("SelectDateButton")
	nar.selectedDateCalendar = builder.FetchCalendar("SelectDateCalendar")
	nar.fromDateRevealer = builder.FetchRevealer("SinceDateRevealer")
	nar.fromDateLabel = builder.FetchLabel("SinceDateLabel")
	nar.fromDateEntry = builder.FetchEntry("SinceDateEntry")
	nar.fromDateResetButton = builder.FetchButton("SinceDateResetButton")
	nar.fromDateCalendar = builder.FetchCalendar("SinceDateCalendar")
	nar.toDateRevealer = builder.FetchRevealer("BeforeDateRevealer")
	nar.toDateLabel = builder.FetchLabel("BeforeDateLabel")
	nar.toDateEntry = builder.FetchEntry("BeforeDateEntry")
	nar.toDateResetButton = builder.FetchButton("BeforeDateResetButton")
	nar.toDateCalendar = builder.FetchCalendar("BeforeDateCalendar")
	nar.createButton = builder.FetchButton("CreateButton")
}
