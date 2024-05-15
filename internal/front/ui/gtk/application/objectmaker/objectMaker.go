package objectmaker

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/advertisementwindow/advertisementcontainer/blockadvcontainer"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/advertisementwindow/advertisementcontainer/lineadvcontainer"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/advertisementwindow/advertisementform"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/advertisementwindow/advertisementform/blockform"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/advertisementwindow/advertisementform/lineform"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/advertisementwindow/clientscontainer"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/advertisementwindow/extrachargescontainer"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/advertisementwindow/liststores"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/advertisementwindow/notebook"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/advertisementwindow/orderscontainer"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/advertisementwindow/tagscontainer"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/application"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type ObjectMaker struct {
	req           RequestHandler
	app           *application.WindowController
	lists         application.ListStores
	dataConverter DataConverter
	tools         Tools
	icon          *gdk.Pixbuf
}

type FileDialog struct {
	*gtk.FileChooserNativeDialog
}

func NewObjectMaker(reqGate RequestHandler, converter DataConverter, tools Tools) application.ObjectMaker {
	obm := &ObjectMaker{req: reqGate, dataConverter: converter, tools: tools}
	obm.NewListStores()
	obm.setIcon()
	return obm
}

func (wm *ObjectMaker) SetApplication(app *application.WindowController) {
	wm.app = app
}

func (wm *ObjectMaker) NewListStores() {
	wm.lists = liststores.Create(wm.tools)
}

func (wm *ObjectMaker) ListStores() application.ListStores {
	return wm.lists
}

func (wm *ObjectMaker) NewBlockForm() application.BlockForm {
	return blockform.CreateBlockAdvPage(wm.lists.OrdersList(), wm.lists.TagsList(), wm.lists.ExtraChargesList(),
		wm.tools, advertisementform.Create(wm.tools), wm.req, wm.dataConverter, wm.app, wm)
}

func (wm *ObjectMaker) NewBlockCopyForm() application.BlockForm {
	orderList, err := wm.lists.OrderListCopy()
	if err != nil {
		wm.app.NewErrorWindow(err)
		return nil
	}
	tagList, err := wm.lists.TagsListCopy()
	if err != nil {
		wm.app.NewErrorWindow(err)
		return nil
	}
	extraCharges, err := wm.lists.ExtraChargeListCopy()
	if err != nil {
		wm.app.NewErrorWindow(err)
		return nil
	}
	return blockform.CreateBlockAdvPage(orderList, tagList, extraCharges,
		wm.tools, advertisementform.Create(wm.tools), wm.req, wm.dataConverter, wm.app, wm)
}

func (wm *ObjectMaker) NewLineForm() application.LineForm {
	return lineform.CreateLineAdvPage(wm.lists.OrdersList(), wm.lists.TagsList(), wm.lists.ExtraChargesList(),
		wm.tools, advertisementform.Create(wm.tools), wm.req, wm.dataConverter, wm.app)
}

func (wm *ObjectMaker) NewLineCopyForm() application.LineForm {
	orderList, err := wm.lists.OrderListCopy()
	if err != nil {
		wm.app.NewErrorWindow(err)
		return nil
	}
	tagList, err := wm.lists.TagsListCopy()
	if err != nil {
		wm.app.NewErrorWindow(err)
		return nil
	}
	extraCharges, err := wm.lists.ExtraChargeListCopy()
	if err != nil {
		wm.app.NewErrorWindow(err)
		return nil
	}
	return lineform.CreateLineAdvPage(orderList, tagList, extraCharges,
		wm.tools, advertisementform.Create(wm.tools), wm.req, wm.dataConverter, wm.app)
}
func (wm *ObjectMaker) NewNotebook(b *builder.Builder, advWin application.AdvertisementsWindow) application.Notebook {
	return notebook.Create(b, wm, wm.app, wm.tools, wm.lists, advWin, wm.req)
}

func (wm *ObjectMaker) NewClientTab(b *builder.Builder) application.ClientTab {
	return clientscontainer.Create(b, wm.req, wm.tools, wm.lists)
}

func (wm *ObjectMaker) NewOrderTab(b *builder.Builder, advWin application.AdvertisementsWindow) application.OrderTab {
	return orderscontainer.Create(b, wm.req, wm.dataConverter, wm.tools, advWin, wm.lists, wm.app)
}

func (wm *ObjectMaker) NewBlockTab(b *builder.Builder, advWin application.AdvertisementsWindow) application.BlockTab {
	return blockadvcontainer.Create(b, wm.lists, wm, wm.tools, wm.req, advWin, wm.app)
}

func (wm *ObjectMaker) NewLineTab(b *builder.Builder, advWin application.AdvertisementsWindow) application.LineTab {
	return lineadvcontainer.Create(b, wm.lists, wm, wm.tools, wm.req, advWin, wm.app)
}

func (wm *ObjectMaker) NewTagTab(b *builder.Builder, advWin application.AdvertisementsWindow) application.TagTab {
	return tagscontainer.Create(b, wm.tools, wm.lists, advWin, wm.req)
}

func (wm *ObjectMaker) NewExtraChargeTab(b *builder.Builder, advWin application.AdvertisementsWindow) application.ChargeTab {
	return extrachargescontainer.Create(b, wm.req, advWin, wm.tools, wm.lists)
}

func (wm *ObjectMaker) NewSaveDialog(winLabel string, parent gtk.IWindow) (application.FileDialoger, error) {
	fileSaver, err := gtk.FileChooserNativeDialogNew(winLabel, parent, gtk.FILE_CHOOSER_ACTION_SAVE, "Сохранить", "Отмена")
	if err != nil {
		return nil, err
	}
	AllFilter, err := gtk.FileFilterNew()
	if err != nil {
		return nil, err
	}
	AllFilter.AddPattern("*")
	AllFilter.SetName("Все файлы")
	fileSaver.AddFilter(AllFilter)

	fc := &FileDialog{fileSaver}
	fc.SetModal(true)
	return fc, nil
}

func (wm *ObjectMaker) NewFolderChooseDialog(winLabel string, parent gtk.IWindow) (application.FileDialoger, error) {
	fileChooser, err := gtk.FileChooserNativeDialogNew(winLabel, parent, gtk.FILE_CHOOSER_ACTION_SELECT_FOLDER, "Выбрать папку", "Отмена")
	if err != nil {
		return nil, err
	}
	fc := &FileDialog{fileChooser}
	fc.SetModal(true)
	return fc, nil
}

func (wm *ObjectMaker) NewChooseDialog(winLabel string, parent gtk.IWindow) (application.FileDialoger, error) {
	fileChooser, err := gtk.FileChooserNativeDialogNew(winLabel, parent, gtk.FILE_CHOOSER_ACTION_OPEN, "Выбрать", "Отмена")
	if err != nil {
		return nil, err
	}
	AllFilter, err := gtk.FileFilterNew()
	if err != nil {
		return nil, err
	}
	AllFilter.AddPattern("*")
	AllFilter.SetName("Все файлы")
	fileChooser.AddFilter(AllFilter)

	fc := &FileDialog{fileChooser}
	fc.SetModal(true)
	return fc, nil
}

func (fc *FileDialog) BindResponseSignal(responseFunc func(self *glib.Object, responce int)) glib.SignalHandle {
	return fc.Connect("response", responseFunc)
}

func (fc *FileDialog) AddFileFilter(name, pattern string) error {
	filter, err := gtk.FileFilterNew()
	if err != nil {
		return err
	}
	filter.AddPattern(pattern)
	filter.SetName(name)
	fc.AddFilter(filter)
	fc.SetFilter(filter)
	return nil
}
