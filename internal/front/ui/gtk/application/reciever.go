package application

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func (wc *WindowController) AppendTag(tag *presenter.TagDTO) {
	glib.IdleAdd(func() {
		wc.listStores.AppendTag(*tag)
	})
}

func (wc *WindowController) RegisterReciever(entry *gtk.EntryBuffer) {
	logging.Logger.Debug("WindowController.RegisterReciever: costRate registred", "costRate", entry)
	wc.costReciever = entry
}

func (wc *WindowController) resetReciever() {
	wc.costReciever = nil
}

func (wc *WindowController) RecieveValue(str string) {
	logging.Logger.Debug("WindowController.RecieveValue: start Recieving value", "cost ", str)

	if wc.costReciever == nil {
		return
	}
	wc.costReciever.SetText(str)
	logging.Logger.Debug("WindowController.RecieveValue: value recieved", "value ", str)
	wc.resetReciever()
}

func (wc *WindowController) RemoveTag(tag string) {
	glib.IdleAdd(func() {
		wc.listStores.RemoveTag(tag)
	})
}

func (wc *WindowController) NewErrorWindow(err error) {

	glib.IdleAdd(func() {
		errWin := wc.objectMaker.ErrorWindow()
		errWin.SetErrorMessage(err.Error())
		errWin.Show()
		wc.UnlockAll()
	})
}

func (wc *WindowController) AppendExtraCharge(charge *presenter.ExtraChargeDTO) {
	glib.IdleAdd(func() {
		wc.listStores.AppendExtraCharge(*charge)
	})
}

func (wc *WindowController) RemoveExtraCharge(charge string) {
	glib.IdleAdd(func() {
		wc.listStores.RemoveExtraCharge(charge)
	})
}

func (wc *WindowController) AppendOrder(order *presenter.OrderDTO) {
	glib.IdleAdd(func() {
		wc.listStores.AppendOrder(*order)
	})
}

func (wc *WindowController) RemoveOrder(order int) {
	glib.IdleAdd(func() {
		wc.listStores.RemoveOrder(order)
	})
}

func (wc *WindowController) AppendClient(client *presenter.ClientDTO) {
	glib.IdleAdd(func() {
		wc.listStores.AppendClient(*client)
	})
}

func (wc *WindowController) RemoveClient(client string) {
	glib.IdleAdd(func() {
		wc.listStores.RemoveClient(client)
	})
}

func (wc *WindowController) AppendBlockAdvertisement(blockAdv *presenter.BlockAdvertisementDTO) {
	glib.IdleAdd(func() {
		wc.listStores.AppendBlockAdvertisement(*blockAdv)
	})
}

func (wc *WindowController) RemoveBlockAdvertisement(blockAdv int) {
	glib.IdleAdd(func() {
		wc.listStores.RemoveBlockAdvertisement(blockAdv)
	})
}

func (wc *WindowController) AppendLineAdvertisement(lineAdv *presenter.LineAdvertisementDTO) {
	glib.IdleAdd(func() {
		wc.listStores.AppendLineAdvertisement(*lineAdv)
	})
}

func (wc *WindowController) RemoveLineAdvertisement(lineAdvsID int) {
	glib.IdleAdd(func() {
		wc.listStores.RemoveLineAdvertisement(lineAdvsID)
	})
}

func (wc *WindowController) ProgressComplete() {
	glib.IdleAdd(func() {
		wc.ProgressWindow.ProgressDone("Операция выполнена")
	})
}

func (wc *WindowController) ProgressCompleteWithError(err error) {
	glib.IdleAdd(func() {
		wc.ProgressWindow.ProgressDone("Операция произошла с ошибкой: \n" + err.Error())
	})
}

func (wc *WindowController) AppendCostRate(costRate *presenter.CostRateDTO) {
	glib.IdleAdd(func() {
		wc.listStores.AppendCostRate(*costRate)
		wc.requstsGate.CheckCostRate()
	})
}

func (wc *WindowController) RemoveCostRate(name string) {
	glib.IdleAdd(func() {
		wc.listStores.RemoveCostRate(name)
		wc.requstsGate.CheckCostRate()
	})
}

func (wc *WindowController) AppendMessage(msg string) {
	glib.IdleAdd(func() {
		wc.listStores.AppendMessage(msg)
	})
}

func (wc *WindowController) SetMode(mode int8) {
	glib.IdleAdd(func() {
		wc.mode = mode
	})
}

func (wc *WindowController) SetConnectionStatus(status bool) {
	glib.IdleAdd(func() {
		wc.connectionStatus = status
	})
}

func (wc *WindowController) ConnectionStatus() bool {
	return wc.connectionStatus
}

func (wc *WindowController) RemoveFileByName(name string) {
	glib.IdleAdd(func() {
		wc.listStores.RemoveFile(name)
	})
}

func (wc *WindowController) AppendFile(file *presenter.File) {
	glib.IdleAdd(func() {
		wc.listStores.AppendFile(*file)
	})
}

func (wc *WindowController) AppendFileFirstPlace(file *presenter.File) {
	glib.IdleAdd(func() {
		wc.listStores.InsertFileFirstPlace(*file)
	})
}

func (wc *WindowController) AppendSelectedFile(file *presenter.File) {
	glib.IdleAdd(func() {
		wc.FileWindow.ShowSelectedFile(*file)
	})
}

func (wc *WindowController) UnlockFilesWindow() {
	glib.IdleAdd(func() {
		wc.FileWindow.LoadFilesComplete()
	})
}

func (wc *WindowController) RequestCompleted() {
	glib.IdleAdd(func() {
		wc.app.GetActiveWindow().Destroy()
	})
}
