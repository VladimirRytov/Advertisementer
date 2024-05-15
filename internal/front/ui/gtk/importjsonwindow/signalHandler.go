package importjsonwindow

import (
	"context"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type signalHandler struct {
	filePathButtonClicked glib.SignalHandle
	cancelButtonPressed   glib.SignalHandle
	importButtonPressed   glib.SignalHandle
	windowDestroyed       glib.SignalHandle
}

func (ijw *ImportJSONWindow) bindSignals() {
	ijw.signals.filePathButtonClicked = ijw.filePathButton.Connect("clicked", ijw.filePathButtonClicked)
	ijw.signals.cancelButtonPressed = ijw.cancelButton.Connect("clicked", ijw.cancelButtonPressed)
	ijw.signals.importButtonPressed = ijw.importButton.Connect("clicked", ijw.importButtonPressed)
	ijw.signals.windowDestroyed = ijw.window.Connect("destroy", ijw.window.Destroy)
}

func (ijw *ImportJSONWindow) filePathButtonClicked() {

	fileChooser, err := ijw.dialogMaker.NewChooseDialog("Импорт", ijw.window)
	if err != nil {
		logging.Logger.Error("importJSONWindow.filePathButtonClicked: cant create fileChooserDialog", "error", err)
		return
	}

	err = fileChooser.AddFileFilter("Json", "*.json")
	if err != nil {
		logging.Logger.Error("ImportJSONWindow.filePathButtonClicked: cannot bind json filter to filechooser", "error", err)
	}

	fileChooser.BindResponseSignal(func(self *glib.Object, responce int) {
		if gtk.RESPONSE_ACCEPT == gtk.ResponseType(responce) {
			selectedDir := fileChooser.GetFilename()

			err := ijw.req.SaveConfig("lastImportPath", &selectedDir)
			if err != nil {
				logging.Logger.Error("importJSONWindow.filePathButtonClicked: cant save lastImportPath", "error", err)
			}
			ijw.SetFilePath(fileChooser.GetFilename())
			return
		}
	})
	fileChooser.Show()
}

func (ijw *ImportJSONWindow) cancelButtonPressed() {
	ijw.window.Close()
}

func (ijw *ImportJSONWindow) importButtonPressed() {
	params := &presenter.ImportParams{
		AllBlocks:       ijw.AllBlockAdv(),
		AlllLines:       ijw.AllLineAdv(),
		ActualClients:   ijw.ActualClients(),
		AllTags:         ijw.Tags(),
		AllExtraCharges: ijw.ExtraCharges(),
		AllCostRates:    ijw.CostRates(),
		IgnoreErrors:    ijw.IgnoreErrors(),
	}

	ctx, cancel := context.WithCancel(context.Background())
	ijw.app.BlockAllSignals()
	ijw.app.DisableAllModels()
	progresWin := ijw.app.CreateProgressWindow()
	progresWin.SetMessage("Выолняется импорт")
	progresWin.SetAfterFunc(func() {
		ijw.req.IgnoreMessages(false)
		ijw.app.UnblockAllSignals()
		ijw.app.EnableAllModels()
		ijw.app.UpdateAll()
	})
	progresWin.SetCancelFunc(cancel)
	if ijw.IgnoreErrors() {
		progresWin.ShowTreeview(true)
	}
	progresWin.Show()
	ijw.req.StartImportingJson(ctx, ijw.Path(), params)
	ijw.Close()
}
