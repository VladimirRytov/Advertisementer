package application

import (
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/gtk"
)

func (wc *WindowController) ActiveWin() gtk.IWindow {
	return wc.app.GetActiveWindow()
}

func (wc *WindowController) CreateAddOrderWindow() AddOrderWindow {
	newOrder := wc.objectMaker.AddOrderWindow()
	wc.app.AddWindow(newOrder.Window())
	return newOrder
}

func (wc *WindowController) CreateAddAdvertisementWindow() AddAdvertisementWindow {
	newAdvWin := wc.objectMaker.AddAdvertisementWindow()
	wc.app.AddWindow(newAdvWin.Window())
	return newAdvWin
}

func (wc *WindowController) CreateLoginWindow() LoginWin {
	loginWin := wc.objectMaker.LoginWindow()
	wc.app.AddWindow(loginWin.Window())
	return loginWin
}

func (wc *WindowController) CloseLoginWindow() {
	wc.LoginWin.Close()
}

func (wc *WindowController) CreateAddTagWindow() AddTagWindow {
	newTag := wc.objectMaker.AddTagWindow()
	wc.app.AddWindow(newTag.Window())
	return newTag
}

func (wc *WindowController) CreateAddExtraChargeWindow() AddExtraChargeWindow {
	newExtraCharge := wc.objectMaker.AddExtraChargeWindow()
	wc.app.AddWindow(newExtraCharge.Window())
	return newExtraCharge
}

func (wc *WindowController) CreateAddClientWindow() AddClientWindow {
	newClient := wc.objectMaker.AddClientWindow()
	wc.app.AddWindow(newClient.Window())
	return newClient
}

func (wc *WindowController) CreateAdvertisementWindow() AdvertisementsWindow {
	advWin := wc.objectMaker.AdvertisementsWindow()
	logging.Logger.Debug("application: creating advertisement window")
	wc.app.AddWindow(advWin.Window())
	return advWin
}

func (wc *WindowController) AdvertisementWindow() AdvertisementsWindow {
	return wc.AdvWin
}

func (wc *WindowController) CloseAdvertisementWindow() {
	wc.AdvWin.Close()
}

func (wc *WindowController) CreateImportDataWindow() ImportDataWindow {
	logging.Logger.Debug("application: creating ImportData window")
	importDataWin := wc.objectMaker.ImportDataWindow()
	wc.app.AddWindow(importDataWin.Window())
	return importDataWin
}

func (wc *WindowController) CreateProgressWindow() ProgressWindow {
	logging.Logger.Debug("application: creating progress window")
	wc.ProgressWindow = wc.objectMaker.ProgressWindow()
	wc.app.AddWindow(wc.ProgressWindow.Window())
	return wc.ProgressWindow
}

func (wc *WindowController) CreateAndShowProgressWindow(message string) {
	logging.Logger.Debug("application: creating progress window")
	wc.ProgressWindow = wc.objectMaker.ProgressWindow()
	wc.ProgressWindow.SetMessage(message)
	wc.app.AddWindow(wc.ProgressWindow.Window())
	wc.ProgressWindow.Show()
}

func (wc *WindowController) CreateCostRatesWindow() CostRateWindow {
	logging.Logger.Debug("application: creating CostRate window")
	costRateWin := wc.objectMaker.CostRateWindow()
	wc.app.AddWindow(costRateWin.Window())
	return costRateWin
}

func (wc *WindowController) CreateNewCostRatesWindow(updateMode bool) NewCostRateWindow {
	logging.Logger.Debug("application: creating NewCostRate window")
	newCostRateWin := wc.objectMaker.NewCostRateWindow(updateMode)
	wc.app.AddWindow(newCostRateWin.Window())
	return newCostRateWin
}

func (wc *WindowController) NewAdvertisementReportWindow() AdvertisementReportWindow {
	logging.Logger.Debug("application: creating AdvertisementReport Window")
	advReportWin := wc.objectMaker.NewAdvertisementReportWindow()
	wc.app.AddWindow(advReportWin.Window())
	return advReportWin
}

func (wc *WindowController) NewThinImageChooserWindow(fileChooser FileChooser) FileChooseWindow {
	logging.Logger.Debug("application: creating ThinImageChooser Window")
	wc.FileWindow = wc.objectMaker.NewThinImageChooserWindow(fileChooser)
	wc.app.AddWindow(wc.FileWindow.Window())
	return wc.FileWindow
}
