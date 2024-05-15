package extrachargescontainer

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
)

type ExtraChargeSignalHander struct {
	treeSelectorChanged            glib.SignalHandle
	removeExtraChargeButtonClicked glib.SignalHandle
	resetExtraChargeButtonClicked  glib.SignalHandle
	applyExtraChargeButtonClicked  glib.SignalHandle
	extraChargesRowDeleted         glib.SignalHandle
	extraChargesRowInserted        glib.SignalHandle
	extraChargesRowChanged         glib.SignalHandle
	multiplierChanged              glib.SignalHandle
}

func (ex *ExtraChargeTab) bindSignals() {
	ex.signalHandler = ExtraChargeSignalHander{
		treeSelectorChanged:            ex.treeSelector.Connect("changed", ex.SelectedRowChanged),
		removeExtraChargeButtonClicked: ex.removeExtraChargeButton.Connect("clicked", ex.removeButtonPressed),
		resetExtraChargeButtonClicked:  ex.resetExtraChargeButton.Connect("clicked", ex.resetButtonPressed),
		applyExtraChargeButtonClicked:  ex.applyExtraChargeButton.Connect("clicked", ex.updateButtonPressed),
		extraChargesRowDeleted:         ex.ExtraCharges.Connect("row-deleted", ex.rowChanged),
		extraChargesRowInserted:        ex.ExtraCharges.Connect("row-inserted", ex.rowChanged),
		extraChargesRowChanged:         ex.ExtraCharges.Connect("row-changed", ex.rowChanged),
		multiplierChanged:              ex.extraChargeMultiplierEntry.Connect("insert-text", ex.tools.CheckExtraChargeString),
	}
}

func (ex *ExtraChargeTab) SelectedRowChanged() {
	_, iter, ok := ex.treeSelector.GetSelected()
	if !ok {
		ex.resetSidebar()
		ex.SetSensitive(false)
		return
	}
	ex.SetSensitive(true)
	chargeName, err := ex.tools.StringFromIter(iter, &ex.ExtraCharges.TreeModel, 1)
	if err != nil {
		logging.Logger.Error("Cannot get ExtraCharge name", "error message", err)
	}
	chargeCost, _ := ex.tools.StringFromIter(iter, &ex.ExtraCharges.TreeModel, 2)
	if err != nil {
		logging.Logger.Error("Cannot get ExtraCharge cost", "error message", err)
	}

	ex.extraChargeNameEntry.SetText(chargeName)
	ex.extraChargeMultiplierEntry.SetText(chargeCost)
	ex.blockAdvsTreeFilter.Refilter()
	ex.lineAdvsTreeFilter.Refilter()
	ex.rightSide.SetSensitive(true)
}

func (ex *ExtraChargeTab) resetRightSide() {
	ex.extraChargeNameEntry.SetText("")
	ex.extraChargeMultiplierEntry.SetText("")
	ex.rightSide.SetSensitive(false)
}

func (ex *ExtraChargeTab) removeButtonPressed() {
	ex.SetSensitive(false)
	ex.req.RemoveExtraCharge(ex.Name())
}

func (ex *ExtraChargeTab) resetButtonPressed() {
	ex.SelectedRowChanged()
}

func (ex *ExtraChargeTab) updateButtonPressed() {
	ex.SetSensitive(false)
	extraCharge := &presenter.ExtraChargeDTO{
		ChargeName: ex.Name(),
		Multiplier: ex.Mulplier(),
	}
	err := ex.req.UpdateExtraCharge(extraCharge)
	if err != nil {
		ex.SetSensitive(true)
		ex.advWin.ShowPopover(err.Error(), ex.extraChargeMultiplierEntry)
	}
}

func (ex *ExtraChargeTab) rowChanged() {
	logging.Logger.Debug("row changed")
	_, _, ok := ex.treeSelector.GetSelected()
	if !ok {
		ex.resetSidebar()
		ex.SetSensitive(false)
		return
	}
	ch := ex.ExtraCharges.IterNChildren(nil)
	if ch == 0 {
		ex.SetSensitive(false)
		ex.resetSidebar()
		return
	}
	ex.SetSensitive(true)
}
