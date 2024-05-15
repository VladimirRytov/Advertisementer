package application

import (
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const (
	ThickMode = int8(iota + 1)
	ThinMode
)

type ReqGate interface {
	CostRateCalculator
	LoadConfig(string, any) error
	SaveConfig(string, any) error
	Databases()
	CloseDatabaseConnection() error
}

type Tools interface {
	StringFromIter(*gtk.TreeIter, *gtk.TreeModel, int) (string, error)
}

type WindowController struct {
	app     *gtk.Application
	version string

	requstsGate ReqGate
	tools       Tools

	connectionStatus bool
	mode             int8
	objectMaker      ObjectMaker
	LoginWin         LoginWin
	AdvWin           AdvertisementsWindow
	ProgressWindow   ProgressWindow
	FileWindow       FileChooseWindow
	listStores       ListStores
	costReciever     *gtk.EntryBuffer
}

func CreateApplication(maker ObjectMaker, reqh ReqGate, tools Tools, version string) *WindowController {
	app, err := gtk.ApplicationNew("com.github.rytowladimir.advertisementer", glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		panic(err)
	}
	wc := &WindowController{app: app, objectMaker: maker, tools: tools, requstsGate: reqh, version: version}
	app.Connect("activate", func() {
		wc.LoadLoginWin()
		wc.LoginWin.Show()
	})
	app.Connect("shutdown", func() {
		logging.Logger.Info("starting shutdown gui")
		reqh.CloseDatabaseConnection()
	})
	wc.objectMaker.SetApplication(wc)
	return wc
}

func (wc *WindowController) LoadLoginWin() {
	wc.listStores = wc.objectMaker.ListStores()
	wc.LoginWin = wc.CreateLoginWindow()
	wc.requstsGate.Databases()
	wc.LoginWin.LoadConfigs()
}

func (wc *WindowController) Start(args []string) {
	wc.app.Run(args)
}

func (wc *WindowController) Stop() {
	wc.app.Quit()
}

func (wc *WindowController) Mode() int8 {
	return wc.mode
}

func (wc *WindowController) LoadNetworkDatabases(dbs []string) {
	wc.LoginWin.LoadNetworkDatabases(dbs)
}

func (wc *WindowController) LoadLocalDatabases(dbs []string) {
	wc.LoginWin.LoadLocalDatabases(dbs)
}

func (wc *WindowController) ShowDefaultPort(port string) {
	wc.LoginWin.ShowDefaultPort(port)
}

func (wc *WindowController) ShowError(err error) {
	wc.LoginWin.ShowError(err)
}

func (wc *WindowController) LockAll() {
	wc.LoginWin.LockAll()
}
func (wc *WindowController) UnlockAllLoginForm() {
	wc.LoginWin.UnlockAll()
}

func (wc *WindowController) UnlockAll() {
	wc.AdvWin.SetSensetive(true)
}

func (wc *WindowController) InitAdvertisement() {

	glib.IdleAdd(func() {
		wc.LoginWin.SaveConfigs()
		if wc.mode == ThinMode {
			wc.listStores.NewFilesList()
		}
		wc.AdvWin = wc.CreateAdvertisementWindow()
		wc.AdvWin.StartInitialization()
	})

	glib.IdleAddPriority(glib.PRIORITY_LOW, func() {
		wc.listStores.SetReplaceMode(true)
		wc.AdvWin.AttachAll()
		wc.setCostRateFromSave()
		wc.AdvWin.UnblockAllSignals()
		wc.LoginWin.Close()
		wc.AdvWin.Show()
	})

}

func (wc *WindowController) SetReplaceMode(val bool) {
	wc.listStores.SetReplaceMode(val)
}

func (wc *WindowController) setCostRateFromSave() {
	var costName string
	err := wc.requstsGate.LoadConfig("CostRate", &costName)
	if err != nil {
		return
	}
	costRateList := wc.listStores.CostRatesListStore()
	costRateList.ForEach(func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
		name, err := wc.tools.StringFromIter(iter, model, 1)
		if err != nil {
			return false
		}
		costRateList.SetValue(iter, 0, name == costName)
		if name == costName {
			err = wc.requstsGate.SetActiveCostRate(costName)
			if err != nil {
				return true
			}
			wc.AdvWin.SelectCostRate(costName)
			return true
		}
		return false
	})
}

func (wc *WindowController) SetActiveCostRate(activeCostcostRate string) {
	logging.Logger.Debug("windowController.SetActiveCostRate: setting active cost rate", "costRate name", activeCostcostRate)
	wc.AdvWin.SelectCostRate("")
	costRateList := wc.listStores.CostRatesListStore()
	costRateList.ForEach(func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
		name, err := wc.tools.StringFromIter(iter, model, 1)
		if err != nil {
			return false
		}
		costRateList.SetValue(iter, 0, name == activeCostcostRate)
		if name == activeCostcostRate {
			wc.AdvWin.SelectCostRate(activeCostcostRate)
		}
		return false
	})
}
func (wc *WindowController) ListStores() ListStores {
	return wc.listStores
}

func (wc *WindowController) SelectCostRate(name string) {
	wc.AdvWin.SelectCostRate(name)
}

func (wc *WindowController) BlockAllSignals() {
	wc.AdvWin.BlockAllSignals()
}

func (wc *WindowController) UnblockAllSignals() {
	wc.AdvWin.UnblockAllSignals()
}

func (wc *WindowController) DisableAllModels() {
	wc.AdvWin.DisableAllModels()
}

func (wc *WindowController) EnableAllModels() {
	wc.AdvWin.EnableAllModels()
}

func (wc *WindowController) UpdateAll() {
	wc.AdvWin.UpdateButtonPressed()
}

func (wc *WindowController) Version() string {
	return wc.version
}
