package objectmaker

import (
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
)

func (wm *ObjectMaker) setIcon() {}

func (wm *ObjectMaker) AddAdvertisementWindow() application.AddAdvertisementWindow {
	return newadvertisement.Create(wm.req, wm.app.AdvertisementWindow(), wm.NewLineForm(), wm.NewBlockForm())
}

func (wm *ObjectMaker) AddClientWindow() application.AddClientWindow {
	return newclient.Create(wm.req)
}

func (wm *ObjectMaker) AddExtraChargeWindow() application.AddExtraChargeWindow {
	return newextracharge.Create(wm.req, wm.tools)
}

func (wm *ObjectMaker) AddOrderWindow() application.AddOrderWindow {
	return neworder.Create(wm.req, wm.app, wm.app.ListStores(), wm.tools, wm)
}

func (wm *ObjectMaker) AddTagWindow() application.AddTagWindow {
	return newtag.Create(wm.req)
}

func (wm *ObjectMaker) ErrorWindow() application.ErrorWindow {
	win := new(errorwindow.ErrorWindow)
	win.Create()
	return win
}

func (wm *ObjectMaker) ImportDataWindow() application.ImportDataWindow {
	return importjsonwindow.Create(wm.req, wm.app, wm)
}

func (wm *ObjectMaker) LoginWindow() application.LoginWin {
	return loginwindow.Create(wm.req, wm.app, wm.tools, wm.lists, wm, wm.dataConverter)
}

func (wm *ObjectMaker) ProgressWindow() application.ProgressWindow {
	return progress.Create(wm.app.ListStores(), wm.tools)
}

func (wm *ObjectMaker) AdvertisementsWindow() application.AdvertisementsWindow {
	return advertisementwindow.Create(wm.req, wm.app, wm, wm.dataConverter, wm)
}

func (wm *ObjectMaker) CostRateWindow() application.CostRateWindow {
	return costratewindow.Create(wm.req, wm.app, wm.app.ListStores(), wm.tools)
}

func (wm *ObjectMaker) NewCostRateWindow(updateMode bool) application.NewCostRateWindow {
	return newcostratewindow.Create(updateMode, wm.req, wm.tools)
}

func (wm *ObjectMaker) NewAdvertisementReportWindow() application.AdvertisementReportWindow {
	return newadvreport.Create(wm.req, wm.dataConverter, wm.app, wm)
}

func (wm *ObjectMaker) NewThinImageChooserWindow(fileChooser application.FileChooser) application.FileChooseWindow {
	return filemanager.Create(wm.req, fileChooser, wm.lists.FilesList(), wm.tools, wm, wm.app)
}
