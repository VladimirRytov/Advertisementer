package filemanager

import (
	"context"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func (fm *FileManager) Close() {
	fm.window.Close()
}

func (fm *FileManager) Show() {
	fm.window.Show()
}

func (fm *FileManager) Window() *gtk.Window {
	return fm.window
}

func (fm *FileManager) LoadFiles() {
	var ctx context.Context
	fm.fileList.Clear()
	ctx, fm.cancelFunc = context.WithCancel(context.Background())
	fm.buttonStack.SetVisibleChildName("CancelLoadButtonStack")
	fm.LockButtons(true)
	fm.reqGate.AllFiles(ctx)
}

func (fm *FileManager) ShowSelectedFile(file presenter.File) {
	pixBufLoader, err := gdk.PixbufLoaderNew()
	if err != nil {
		logging.Logger.Error("FileManager.ShowSelectedFile: cannot create pixbuf loader", "error", err)
		return
	}
	img, err := pixBufLoader.WriteAndReturnPixbuf(file.Data)
	if err != nil {
		logging.Logger.Error("FileManager.ShowSelectedFile: cannot create pixbuf", "error", err)
		return
	}
	fm.nameLabel.SetText(file.Name)
	fm.sizeLabel.SetText(file.Size)
	fm.previewImage.SetFromPixbuf(img)
}

func (fm *FileManager) LockButtons(lock bool) {
	fm.addFileButton.SetSensitive(!lock)
	fm.selectButton.SetSensitive(!lock)
	fm.removeButton.SetSensitive(!lock)
}

func (fm *FileManager) LockAddRemoveButtons(lock bool) {
	fm.selectButton.SetSensitive(!lock)
	fm.removeButton.SetSensitive(!lock)
}

func (fm *FileManager) LoadFilesComplete() {
	fm.cancelButtonClicked()
}
