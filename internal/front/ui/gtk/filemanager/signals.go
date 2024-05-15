package filemanager

import (
	"context"

	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type FileManagerSignals struct {
	windowDestroyed         glib.SignalHandle
	iconClicked             glib.SignalHandle
	searchChanged           glib.SignalHandle
	removeButtonClicked     glib.SignalHandle
	selectFileButtonClicked glib.SignalHandle
	cancelFileButtonClicked glib.SignalHandle
	updateListButtonClicked glib.SignalHandle
	addFileButtonClicked    glib.SignalHandle
}

func (fm *FileManager) bindSignals() {
	fm.signals.windowDestroyed = fm.window.Connect("destroy", fm.cancelButtonClicked)
	fm.signals.iconClicked = fm.iconView.Connect("item-activated", fm.iconClicked)
	fm.signals.searchChanged = fm.searchEntry.Connect("search-changed", fm.fileFilterList.Refilter)
	fm.signals.removeButtonClicked = fm.removeButton.Connect("clicked", fm.removeButtonClicked)
	fm.signals.selectFileButtonClicked = fm.selectButton.Connect("clicked", fm.selectButtonClicked)
	fm.signals.updateListButtonClicked = fm.reloadButton.Connect("clicked", fm.updateButtonClicked)
	fm.signals.cancelFileButtonClicked = fm.cancelButton.Connect("clicked", fm.cancelButtonClicked)
	fm.signals.addFileButtonClicked = fm.addFileButton.Connect("clicked", fm.addFileButtonClicked)
}

func (fm *FileManager) iconClicked(self *gtk.IconView, path *gtk.TreePath) {
	iter, err := fm.fileFilterList.GetIter(path)
	if err != nil {
		logging.Logger.Error("FileManager.iconClicked: cannot get iter from path", "error", err)
		return
	}
	fm.LockAddRemoveButtons(false)
	fileName, err := fm.tools.StringFromIter(iter, fm.fileFilterList.ToTreeModel(), 0)
	if err != nil {
		logging.Logger.Error("FileManager.iconClicked: cannot find file in listStore", "error", err)
		return
	}
	err = fm.reqGate.LargeFileByName(fileName)
	if err != nil {
		logging.Logger.Error("FileManager.iconClicked: request is incorrect ", "error", err)
		return
	}
	fm.selectedFile = fileName
	if fm.fileChooser.FilePath() == fileName {
		fm.selectButton.SetSensitive(false)
		return
	}
	fm.selectButton.SetSensitive(true)
	fm.removeButton.SetSensitive(true)
}

func (fm *FileManager) selectButtonClicked() {
	fm.fileChooser.SetFilePath(fm.selectedFile)
	fm.window.Destroy()
}

func (fm *FileManager) removeButtonClicked() {
	fm.reqGate.RemoveFile(fm.selectedFile)
	fm.LockAddRemoveButtons(true)
}

func (fm *FileManager) updateButtonClicked() {
	fm.fileList.Clear()
	ctx, cancel := context.WithCancel(context.Background())
	fm.cancelFunc = cancel
	fm.buttonStack.SetVisibleChildName("CancelLoadButtonStack")
	fm.LockButtons(true)
	fm.reqGate.AllFiles(ctx)
}

func (fm *FileManager) cancelButtonClicked() {
	fm.cancelFunc()
	fm.buttonStack.SetVisibleChildName("ReloadButtonStack")
	fm.LockButtons(false)
}

func (fm *FileManager) addFileButtonClicked() {
	dialog, err := fm.dialogMaker.NewChooseDialog("Выберите файл для загрузки", fm.window)
	if err != nil {
		logging.Logger.Error("FileManager.addFileButtonClicked: cannot create file choose dialog", "error", err)
		return
	}
	dialog.BindResponseSignal(func(self *glib.Object, responce int) {
		if gtk.RESPONSE_ACCEPT == gtk.ResponseType(responce) {
			ctx, cancel := context.WithCancel(context.Background())
			progresWin := fm.app.CreateProgressWindow()
			progresWin.SetCancelFunc(cancel)
			progresWin.SetAfterFunc(fm.selectFirstImage)
			progresWin.Show()
			fm.reqGate.UploadFiles(ctx, dialog.GetFilename())
		}
	})
	dialog.Show()
}

func (fm *FileManager) selectFirstImage() {
	iter, ok := fm.fileFilterList.GetIterFirst()
	if !ok {
		return
	}
	path, err := fm.fileFilterList.GetPath(iter)
	if err != nil {
		return
	}
	fm.iconView.SelectPath(path)
	fm.iconView.ItemActivated(path)
}
