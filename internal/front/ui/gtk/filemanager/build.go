package filemanager

import (
	"context"
	"strings"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/application"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"

	"github.com/gotk3/gotk3/gtk"
)

type LoadRemoverer interface {
	AllFiles(context.Context)
	NewFile(context.Context, *presenter.File) error
	UploadFiles(ctx context.Context, file string)

	RemoveFile(string) error
	FileByName(string) error
	LargeFileByName(string) error
}

type FileChooser interface {
	FilePath() string
	SetFilePath(string)
}

type Tools interface {
	StringFromIter(iter *gtk.TreeIter, model *gtk.TreeModel, col int) (string, error)
}

type SelectFileDialoger interface {
	NewChooseDialog(winLabel string, parent gtk.IWindow) (application.FileDialoger, error)
}

type ProgresHandler interface {
	CreateProgressWindow() application.ProgressWindow
}

type FileManager struct {
	app         ProgresHandler
	reqGate     LoadRemoverer
	fileChooser FileChooser
	tools       Tools
	signals     FileManagerSignals
	dialogMaker SelectFileDialoger
	cancelFunc  context.CancelFunc

	window        *gtk.Window
	addFileButton *gtk.Button
	searchEntry   *gtk.SearchEntry
	searchBuffer  *gtk.EntryBuffer

	buttonStack  *gtk.Stack
	reloadButton *gtk.Button
	cancelButton *gtk.Button

	iconView       *gtk.IconView
	selectedFile   string
	fileList       *gtk.ListStore
	fileFilterList *gtk.TreeModelFilter

	previewImage *gtk.Image
	nameLabel    *gtk.Label
	sizeLabel    *gtk.Label
	removeButton *gtk.Button
	selectButton *gtk.Button
}

func Create(reqGate LoadRemoverer, chooser FileChooser, fileList *gtk.ListStore, tools Tools, objMaker SelectFileDialoger, app ProgresHandler) *FileManager {
	builder, err := builder.NewBuilderFromString(builder.Files)
	if err != nil {
		panic(err)
	}
	fileManager := new(FileManager)
	fileManager.fileList = fileList
	fileManager.app = app
	fileManager.dialogMaker = objMaker
	fileManager.reqGate = reqGate
	fileManager.fileChooser = chooser
	fileManager.tools = tools
	fileManager.window = builder.FetchWindow("FilesWindow")
	fileManager.addFileButton = builder.FetchButton("FilesAddFileButton")
	fileManager.searchEntry = builder.FetchSearchEntry("SearchFileEntry")
	fileManager.searchBuffer, err = fileManager.searchEntry.GetBuffer()
	if err != nil {
		panic(err)
	}
	fileManager.buttonStack = builder.FetchStack("ReloadStack")
	fileManager.reloadButton = builder.FetchButton("ReloadImagesButton")
	fileManager.cancelButton = builder.FetchButton("CancelLoadButton")
	fileManager.iconView = builder.FetchIconView("FilesIconView")
	fileManager.previewImage = builder.FetchImage("FilesPreviewGtkImage")
	fileManager.nameLabel = builder.FetchLabel("FilesFileNameLabel")
	fileManager.sizeLabel = builder.FetchLabel("FilesSizeLabel")
	fileManager.removeButton = builder.FetchButton("FilesDeleteButton")
	fileManager.selectButton = builder.FetchButton("FilesSelectButton")
	fileManager.fileFilterList, err = fileManager.fileList.FilterNew(nil)
	if err != nil {
		panic(err)
	}
	fileManager.fileFilterList.SetVisibleFunc(func(model *gtk.TreeModel, iter *gtk.TreeIter) bool {
		if fileManager.searchBuffer.GetLength() == 0 {
			return true
		}
		file, err := fileManager.tools.StringFromIter(iter, model, 0)
		if err != nil {
			return false
		}
		msg, err := fileManager.searchBuffer.GetText()
		if err != nil {
			return false
		}
		return strings.Contains(file, msg)
	})
	fileManager.iconView.SetModel(fileManager.fileFilterList)
	fileManager.bindSignals()
	fileManager.window.SetTitle("Менеждер файлов")
	return fileManager
}
