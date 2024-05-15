package costratewindow

import "github.com/gotk3/gotk3/gtk"

func (cs *CostRateWindow) Close() {
	cs.window.Destroy()
}

func (cs *CostRateWindow) Window() *gtk.Window {
	return cs.window
}

func (cs *CostRateWindow) Show() {
	cs.window.Show()
}

func (cs *CostRateWindow) UnsetModel() {
	cs.treeview.SetModel(nil)
}

func (cs *CostRateWindow) SetModel() {
	cs.treeview.SetModel(cs.list)
}

func (cs *CostRateWindow) SetSensitiveAll(s bool) {
	cs.addButton.SetSensitive(s)
	cs.refreshButton.SetSensitive(s)
	cs.delButton.SetSensitive(s)
	cs.editButton.SetSensitive(s)
}
