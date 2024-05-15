//go:build !windows

package objectmaker

import (
	_ "embed"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/advertisementwindow"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/application"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/costratewindow"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/errorwindow"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/filemanager"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/importjsonwindow"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/loginwindow"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/newadvertisement"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/newadvreport"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/newclient"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/newcostratewindow"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/newextracharge"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/neworder"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/newtag"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/progress"
	"github.com/VladimirRytov/advertisementer/internal/logging"
	"github.com/gotk3/gotk3/gdk"
)

//go:embed advertisementer128.png
var rawIcon []byte

func (wm *ObjectMaker) setIcon() {
	var err error
	wm.icon, err = gdk.PixbufNewFromBytesOnly(rawIcon)
	if err != nil {
		logging.Logger.Error("objectMaker.setIcon: cannot create pixbuf", "error", err)
		return
	}
}

func (wm *ObjectMaker) AddAdvertisementWindow() application.AddAdvertisementWindow {
	window := newadvertisement.Create(wm.req, wm.app.AdvertisementWindow(), wm.NewLineForm(), wm.NewBlockForm())
	window.Window().SetIcon(wm.icon)
	return window
}

func (wm *ObjectMaker) AddClientWindow() application.AddClientWindow {
	window := newclient.Create(wm.req)
	window.Window().SetIcon(wm.icon)
	return window
}

func (wm *ObjectMaker) AddExtraChargeWindow() application.AddExtraChargeWindow {
	window := newextracharge.Create(wm.req, wm.tools)
	window.Window().SetIcon(wm.icon)
	return window
}

func (wm *ObjectMaker) AddOrderWindow() application.AddOrderWindow {
	window := neworder.Create(wm.req, wm.app, wm.app.ListStores(), wm.tools, wm)
	window.Window().SetIcon(wm.icon)
	return window
}

func (wm *ObjectMaker) AddTagWindow() application.AddTagWindow {
	window := newtag.Create(wm.req)
	window.Window().SetIcon(wm.icon)
	return window
}

func (wm *ObjectMaker) ErrorWindow() application.ErrorWindow {
	win := new(errorwindow.ErrorWindow)
	win.Create()
	win.Window().SetIcon(wm.icon)
	return win
}

func (wm *ObjectMaker) ImportDataWindow() application.ImportDataWindow {
	window := importjsonwindow.Create(wm.req, wm.app, wm)
	window.Window().SetIcon(wm.icon)
	return window
}

func (wm *ObjectMaker) LoginWindow() application.LoginWin {
	loginWin := loginwindow.Create(wm.req, wm.app, wm.tools, wm.lists, wm, wm.dataConverter)
	loginWin.Window().SetIcon(wm.icon)
	return loginWin
}

func (wm *ObjectMaker) ProgressWindow() application.ProgressWindow {
	window := progress.Create(wm.app.ListStores(), wm.tools)
	window.Window().SetIcon(wm.icon)
	return window
}

func (wm *ObjectMaker) AdvertisementsWindow() application.AdvertisementsWindow {
	window := advertisementwindow.Create(wm.req, wm.app, wm, wm.dataConverter, wm)
	window.Window().SetIcon(wm.icon)
	return window
}

func (wm *ObjectMaker) CostRateWindow() application.CostRateWindow {
	window := costratewindow.Create(wm.req, wm.app, wm.app.ListStores(), wm.tools)
	window.Window().SetIcon(wm.icon)
	return window
}

func (wm *ObjectMaker) NewCostRateWindow(updateMode bool) application.NewCostRateWindow {
	window := newcostratewindow.Create(updateMode, wm.req, wm.tools)
	window.Window().SetIcon(wm.icon)
	return window
}

func (wm *ObjectMaker) NewAdvertisementReportWindow() application.AdvertisementReportWindow {
	window := newadvreport.Create(wm.req, wm.dataConverter, wm.app, wm)
	window.Window().SetIcon(wm.icon)
	return window
}

func (wm *ObjectMaker) NewThinImageChooserWindow(fileChooser application.FileChooser) application.FileChooseWindow {
	window := filemanager.Create(wm.req, fileChooser, wm.lists.FilesList(), wm.tools, wm, wm.app)
	window.Window().SetIcon(wm.icon)
	return window
}
