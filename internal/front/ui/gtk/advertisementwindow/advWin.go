package advertisementwindow

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const (
	ShowData = "ShowData"
)

func (aw *AdvertisementsWindow) LockTagsTab(lock bool) {
	aw.NoteBook.LockTagsTab(lock)
}

func (aw *AdvertisementsWindow) LockExtraChargesTab(lock bool) {
	aw.NoteBook.LockExtraChargesTab(lock)
}

func (aw *AdvertisementsWindow) LockClientsTab(lock bool) {
	aw.NoteBook.LockClientsTab(lock)
}

func (aw *AdvertisementsWindow) LockOrdersTab(lock bool) {
	aw.NoteBook.LockOrdersTab(lock)
}

func (aw *AdvertisementsWindow) LockBlockAdvertisementsTab(lock bool) {
	aw.NoteBook.LockBlockAdvertisementsTab(lock)
}

func (aw *AdvertisementsWindow) LockLineAdvertisementsTab(lock bool) {
	aw.NoteBook.LockLineAdvertisementsTab(lock)
}

func (aw *AdvertisementsWindow) ActualSelected() bool {
	return aw.filterShowActualRadioButton.GetActive()
}
func (aw *AdvertisementsWindow) Window() *gtk.Window {
	return aw.MainWindow
}

func (aw *AdvertisementsWindow) RefilterLists() {
	aw.NoteBook.CurrentPageChanged()
}

func (aw *AdvertisementsWindow) Update() {
	aw.NoteBook.UpdateButtonPressed()
}

func (aw *AdvertisementsWindow) StartInitialization() {
	aw.req.AllTags()
	aw.req.AllExtraCharges()
	aw.req.AllClients()
	aw.req.AllOrders()
	aw.req.AllCostRates()
	if aw.filterShowActualRadioButton.GetActive() {
		aw.req.BlockAdvertisementsActual()
		aw.req.LineAdvertisementsActual()
	} else {
		aw.req.AllBlockAdvertisements()
		aw.req.AllLineAdvertisements()
	}
}

func (aw *AdvertisementsWindow) UpdateLists() {
	aw.req.AllTags()
	aw.req.AllExtraCharges()
	aw.req.AllClients()
	aw.req.AllOrders()
	aw.req.AllCostRates()
	if aw.filterShowActualRadioButton.GetActive() {
		aw.req.BlockAdvertisementsActual()
		aw.req.LineAdvertisementsActual()
	} else {
		aw.req.AllBlockAdvertisements()
		aw.req.AllLineAdvertisements()
	}
}

func (aw *AdvertisementsWindow) AttachAll() {
	aw.NoteBook.AttachAll()
}

func (aw *AdvertisementsWindow) BlockAllSignals() {
	aw.NoteBook.BlockAllSignals()
}

func (aw *AdvertisementsWindow) UnblockAllSignals() {
	aw.NoteBook.UnblockAllSignals()
}

func (aw *AdvertisementsWindow) SelectCostRate(name string) {
	glib.IdleAddPriority(glib.PRIORITY_HIGH, func() {
		aw.costRateEntry.SetText(name)
	})
}

func (aw *AdvertisementsWindow) SetConnectionStatus(status bool) {
	if status {
		aw.connectionStatusButton.SetImage(aw.establiShedImage)
		return
	}
	aw.connectionStatusButton.SetImage(aw.lostConnImage)
}

func (aw *AdvertisementsWindow) ClearAllListStores() {
	aw.NoteBook.ClearAllListStores()
}

func (aw *AdvertisementsWindow) EnableAdvertisementFilters(enable bool) {
	aw.NoteBook.EnableAdvertisementFilters(enable)
}

func (aw *AdvertisementsWindow) EnableSidebarsFilters(enable bool) {
	aw.NoteBook.EnableSidebarsFilters(enable)
}

func (aw *AdvertisementsWindow) DisableAllModels() {
	aw.NoteBook.DisableAllModels()
}

func (aw *AdvertisementsWindow) EnableAllModels() {
	aw.NoteBook.EnableAllModels()
}

func (aw *AdvertisementsWindow) ResetSorts() {
	aw.NoteBook.ResetSorts()
}

func (aw *AdvertisementsWindow) SetSensetive(b bool) {
	aw.NoteBook.SetSensetive(b)
}

func (aw *AdvertisementsWindow) CurrentPageChanged() {
	aw.NoteBook.CurrentPageChanged()
}

func (aw *AdvertisementsWindow) RefilterBlock() {
	aw.NoteBook.RefilterBlock()
}

func (aw *AdvertisementsWindow) RefilterLine() {
	aw.NoteBook.RefilterLine()
}

func (aw *AdvertisementsWindow) UpdateButtonPressed() {
	aw.NoteBook.UpdateButtonPressed()
}
