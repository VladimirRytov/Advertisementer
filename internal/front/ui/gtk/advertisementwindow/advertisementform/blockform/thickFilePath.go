package blockform

import (
	"path/filepath"

	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func (thck *ThickFilePath) Box() *gtk.Box {
	return thck.box
}

func (thck *ThickFilePath) FilePath() string {
	return thck.filePathEntry.GetLayout().GetText()
}

func (thck *ThickFilePath) SetFilePath(path string) {
	b, err := thck.filePathEntry.GetBuffer()
	if err != nil {
		logging.Logger.Error("blockAdvPage.SetFilePath: cannot get entry buffer", "error", err)
		return
	}
	b.SetText(path)
}

func (thck *ThickFilePath) viewClicked() {
	fileChooser, err := thck.dialogMaker.NewChooseDialog("Выберите файл", thck.app.ActiveWin())
	if err != nil {
		logging.Logger.Error("thickFilePath.ViewClicked: cannot create fileCHoose window", "error", err)
		return
	}
	fileChooser.BindResponseSignal(func(self *glib.Object, responce int) {
		if gtk.RESPONSE_ACCEPT == gtk.ResponseType(responce) {
			thck.SetFilePath(filepath.Base(fileChooser.GetFilename()))
		}
	})
	fileChooser.Show()
}

type ThickSignals struct {
	viewButtonPressed glib.SignalHandle
}

func (thck *ThickFilePath) bindSignals() {
	thck.signals = ThickSignals{
		viewButtonPressed: thck.filePathView.Connect("clicked", thck.viewClicked),
	}
}
