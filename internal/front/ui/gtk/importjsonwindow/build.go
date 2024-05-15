package importjsonwindow

import (
	"context"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/application"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/gtk"
)

type Requests interface {
	LoadConfig(string, any) error
	SaveConfig(string, any) error
	StartImportingJson(context.Context, string, *presenter.ImportParams)
	IgnoreMessages(bool)
}

type Application interface {
	BlockAllSignals()
	CreateProgressWindow() application.ProgressWindow
	UnblockAllSignals()
	DisableAllModels()
	EnableAllModels()
	UpdateAll()
	Mode() int8
}

type ChooseFileDialoger interface {
	NewChooseDialog(winLabel string, parent gtk.IWindow) (application.FileDialoger, error)
}

type ImportJSONWindow struct {
	req         Requests
	app         Application
	window      *gtk.Window
	dialogMaker ChooseFileDialoger

	filePathLabel  *gtk.Label
	filePathEntry  *gtk.Entry
	filePathButton *gtk.Button

	blockAdvRadioButton           *gtk.RadioButton
	blockAdvOnlyActualRadioButton *gtk.RadioButton

	lineAdvRadioButton           *gtk.RadioButton
	lineAdvOnlyActualRadioButton *gtk.RadioButton

	clientRadioButton           *gtk.RadioButton
	clientOnlyActualRadioButton *gtk.RadioButton

	tagsCheckButton         *gtk.CheckButton
	extraChargesCheckButton *gtk.CheckButton
	costRateCheckButton     *gtk.CheckButton
	ignoreErrorsCheckButton *gtk.CheckButton

	cancelButton *gtk.Button
	importButton *gtk.Button

	signals signalHandler
}

func Create(reqGate Requests, app Application, objMaker ChooseFileDialoger) *ImportJSONWindow {
	build, err := builder.NewBuilderFromString(builder.ImportJSONWindow)
	if err != nil {
		panic(err)
	}
	ijw := new(ImportJSONWindow)
	ijw.req = reqGate
	ijw.app = app
	ijw.dialogMaker = objMaker
	ijw.Build(build)
	var selected string
	err = ijw.req.LoadConfig("lastImportPath", &selected)
	if err != nil {
		logging.Logger.Error("importJSONWindow.filePathButtonClicked: cant load lastImportPath", "error", err)
	}
	ijw.SetFilePath(selected)
	ijw.window.SetTitle("Импорт базы")
	ijw.bindSignals()
	return ijw
}

func (ijw *ImportJSONWindow) Build(buildFile *builder.Builder) {
	ijw.window = buildFile.FetchWindow("ImportJSONWindow")
	ijw.filePathLabel = buildFile.FetchLabel("FilePathLabel")
	ijw.filePathEntry = buildFile.FetchEntry("FilePathEntry")
	ijw.filePathButton = buildFile.FetchButton("FilePathButton")
	ijw.blockAdvRadioButton = buildFile.FetchRadioButton("BlockAdvRadioButton")
	ijw.blockAdvOnlyActualRadioButton = buildFile.FetchRadioButton("BlockAdvActualRadioButton")
	ijw.lineAdvRadioButton = buildFile.FetchRadioButton("LineAdvRadioButton")
	ijw.lineAdvOnlyActualRadioButton = buildFile.FetchRadioButton("LineAdvActualRadioButton")
	ijw.clientRadioButton = buildFile.FetchRadioButton("AllClientsRadioButton")
	ijw.clientOnlyActualRadioButton = buildFile.FetchRadioButton("ActualClientsRadioButton")
	ijw.tagsCheckButton = buildFile.FetchCheckButton("TagsCheckButton")
	ijw.extraChargesCheckButton = buildFile.FetchCheckButton("ExtraChargesCheckButton")
	ijw.costRateCheckButton = buildFile.FetchCheckButton("CostRateCheckButton")
	ijw.ignoreErrorsCheckButton = buildFile.FetchCheckButton("IgnoreErrorsCheckButton")
	ijw.cancelButton = buildFile.FetchButton("CancelButton")
	ijw.importButton = buildFile.FetchButton("StartImportButton")
}
