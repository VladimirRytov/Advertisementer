package costratewindow

import (
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type CostRateSignals struct {
	costRateSelected glib.SignalHandle

	addButtonClicked     glib.SignalHandle
	delButtonBlicked     glib.SignalHandle
	closeButtonClicked   glib.SignalHandle
	refreshButtonClicked glib.SignalHandle
	editButtonClicked    glib.SignalHandle
}

func (cs *CostRateWindow) bindSignals() {
	cs.signals.closeButtonClicked = cs.closeButton.Connect("clicked", cs.closeButtonClicked)
	cs.signals.addButtonClicked = cs.addButton.Connect("clicked", cs.addButtonClicked)
	cs.signals.delButtonBlicked = cs.delButton.Connect("clicked", cs.deleteButtonClicked)
	cs.signals.refreshButtonClicked = cs.refreshButton.Connect("clicked", cs.refreshButtonClicked)
	cs.signals.editButtonClicked = cs.editButton.Connect("clicked", cs.editButtonClicked)
	cs.signals.costRateSelected = cs.chooseSelector.Connect("toggled", cs.costRateSelected)

}

func (cs *CostRateWindow) deleteButtonClicked() {
	selected, err := cs.treeview.GetSelection()
	if err != nil {
		logging.Logger.Warn("CostRateWindow.deleteButtonClicked: cannot get selection", "error", err)
		return
	}
	model, iter, exist := selected.GetSelected()
	if !exist {
		logging.Logger.Warn("CostRateWindow.deleteButtonClicked: cannot get iter from selection", "error", err)
		return
	}
	costName, err := cs.tools.StringFromIter(iter, model.ToTreeModel(), 1)
	if err != nil {
		logging.Logger.Warn("CostRateWindow.deleteButtonClicked: cannot get costName from model", "error", err)
		return
	}
	cs.req.RemoveCostRate(costName)
}

func (cs *CostRateWindow) addButtonClicked() {
	newCostRate := cs.app.CreateNewCostRatesWindow(false)
	newCostRate.Show()
}

func (cs *CostRateWindow) editButtonClicked() {
	selection, err := cs.treeview.GetSelection()
	if err != nil {
		logging.Logger.Warn("CostRateWindow.editButtonClicked: cannot get selection", "error", err)
		return
	}
	model, iter, ok := selection.GetSelected()
	if !ok {
		return
	}

	name, err := cs.tools.StringFromIter(iter, model.ToTreeModel(), 1)
	if err != nil {
		logging.Logger.Error("CostRateWindow.editButtonClicked: cannot get name from model", "error", err)
		return
	}

	forSquare, err := cs.tools.StringFromIter(iter, model.ToTreeModel(), 2)
	if err != nil {
		logging.Logger.Error("CostRateWindow.editButtonClicked: cannot get cost for square from model", "error", err)
		return
	}

	forSymbolWord, err := cs.tools.StringFromIter(iter, model.ToTreeModel(), 3)
	if err != nil {
		logging.Logger.Error("CostRateWindow.editButtonClicked: cannot get cost for symbol or word from model", "error", err)
		return
	}
	icalculateForWord, err := cs.tools.InterfaceFromIter(iter, model.ToTreeModel(), 4)
	if err != nil {
		logging.Logger.Error("CostRateWindow.editButtonClicked: cannot get calculateForWord interface from model", "error", err)
		return
	}
	calculateForWord, ok := icalculateForWord.(bool)
	if !ok {
		logging.Logger.Error("CostRateWindow.editButtonClicked: cannot convert interface to bool")
		return
	}
	newCostRate := cs.app.CreateNewCostRatesWindow(true)
	newCostRate.SetName(name)
	newCostRate.SetCostForOneSquare(forSquare)
	newCostRate.SetCostForWordSymbol(forSymbolWord)
	newCostRate.SetCostForOneWord(calculateForWord)
	newCostRate.Show()
}

func (cs *CostRateWindow) refreshButtonClicked() {
	glib.IdleAdd(func() {
		cs.SetSensitiveAll(false)
		cs.UnsetModel()
		cs.req.AllCostRates()
		cs.SetModel()
	})
	glib.IdleAddPriority(glib.PRIORITY_DEFAULT_IDLE+1, func() {
		cs.SetSensitiveAll(true)
	})
}

func (cs *CostRateWindow) closeButtonClicked() {
	cs.window.Destroy()
}

func (cs *CostRateWindow) costRateSelected(self *gtk.CellRendererToggle, StrPath string) {
	var name string
	cs.list.ForEach(func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
		err := cs.list.SetValue(iter, 0, StrPath == path.String())
		if err != nil {
			logging.Logger.Error("costRateWindow.costRateSelected: cannot set value to listStore", "error", err)
			return true
		}
		if StrPath == path.String() {
			name, err = cs.tools.StringFromIter(iter, model, 1)
			if err != nil {
				logging.Logger.Error("costRateWindow.costRateSelected: cannot set costRateName from listStore", "error", err)
			}
		}
		return false
	})
	cs.app.SelectCostRate(name)
	cs.selectCostRate(name)
}

func (cs *CostRateWindow) costForWordSelected(self *gtk.CellRendererToggle, path string) {
	iter, err := cs.list.GetIterFromString(path)
	if err != nil {
		logging.Logger.Error("costRateWindow.costForWordSelected: cannot get iter from string", "error", err)
		return
	}
	cs.list.Set(iter, []int{4, 5}, []interface{}{true, false})
}

func (cs *CostRateWindow) costForSymbolSelected(self *gtk.CellRendererToggle, path string) {
	iter, err := cs.list.GetIterFromString(path)
	if err != nil {
		logging.Logger.Error("costRateWindow.costForSymbolSelected: cannot get iter from string", "error", err)
		return
	}
	cs.list.Set(iter, []int{4, 5}, []interface{}{false, true})
}

func (cs *CostRateWindow) selectCostRate(name string) {
	err := cs.req.SetActiveCostRate(name)
	if err != nil {
		cs.app.NewErrorWindow(err)
		return
	}
	err = cs.req.SaveConfig("CostRate", &name)
	if err != nil {
		logging.Logger.Error("costRateWindow.selectCostRate: cannow save costRate to configStorage", "error", err)
		return
	}
}
