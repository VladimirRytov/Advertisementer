package extrachargescontainer

import (
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/gtk"
)

func (ex *ExtraChargeTab) Name() string {
	return ex.extraChargeNameEntry.GetLayout().GetText()
}

func (ex *ExtraChargeTab) SetName(name string) {
	ex.extraChargeNameEntry.SetText(name)
}

func (ex *ExtraChargeTab) Mulplier() string {
	return ex.extraChargeMultiplierEntry.GetLayout().GetText()
}

func (ex *ExtraChargeTab) SetSensitive(lock bool) {
	ex.extraChargeMultiplierEntry.SetSensitive(lock)
	ex.removeExtraChargeButton.SetSensitive(lock)
	ex.resetExtraChargeButton.SetSensitive(lock)
	ex.applyExtraChargeButton.SetSensitive(lock)
}

func (ex *ExtraChargeTab) resetSidebar() {
	ex.extraChargeNameEntry.SetText("")
	ex.extraChargeMultiplierEntry.SetText("")
}

func (ex *ExtraChargeTab) SetMultiplier(cost string) {
	ex.extraChargeMultiplierEntry.SetText(cost)
}

func (ex *ExtraChargeTab) filterAdvertisements(model *gtk.TreeModel, iter *gtk.TreeIter) bool {
	if !ex.enableFilters {
		return false
	}
	charges, err := ex.tools.StringFromIter(iter, model, 7)
	if err != nil {
		logging.Logger.Warn("chargesFilterAdvertisements: got empty string. Skipping", "error", err)
		return false
	}
	return ex.req.ValueInString(charges, ex.Name())
}

func (ex *ExtraChargeTab) BlockSignals() {
	ex.treeSelector.HandlerBlock(ex.signalHandler.treeSelectorChanged)
	ex.removeExtraChargeButton.HandlerBlock(ex.signalHandler.removeExtraChargeButtonClicked)
	ex.resetExtraChargeButton.HandlerBlock(ex.signalHandler.resetExtraChargeButtonClicked)
	ex.applyExtraChargeButton.HandlerBlock(ex.signalHandler.applyExtraChargeButtonClicked)
	ex.ExtraCharges.HandlerBlock(ex.signalHandler.extraChargesRowDeleted)
	ex.ExtraCharges.HandlerBlock(ex.signalHandler.extraChargesRowChanged)
	ex.ExtraCharges.HandlerBlock(ex.signalHandler.extraChargesRowInserted)
	ex.extraChargeMultiplierEntry.HandlerBlock(ex.signalHandler.multiplierChanged)
}

func (ex *ExtraChargeTab) UnblockSignals() {
	ex.treeSelector.HandlerUnblock(ex.signalHandler.treeSelectorChanged)
	ex.removeExtraChargeButton.HandlerUnblock(ex.signalHandler.removeExtraChargeButtonClicked)
	ex.resetExtraChargeButton.HandlerUnblock(ex.signalHandler.resetExtraChargeButtonClicked)
	ex.applyExtraChargeButton.HandlerUnblock(ex.signalHandler.applyExtraChargeButtonClicked)
	ex.ExtraCharges.HandlerUnblock(ex.signalHandler.extraChargesRowDeleted)
	ex.ExtraCharges.HandlerUnblock(ex.signalHandler.extraChargesRowChanged)
	ex.ExtraCharges.HandlerUnblock(ex.signalHandler.extraChargesRowInserted)
	ex.extraChargeMultiplierEntry.HandlerUnblock(ex.signalHandler.multiplierChanged)
}

func (ex *ExtraChargeTab) SetEnableFilters(enable bool) {
	ex.enableFilters = enable
}

func (ex *ExtraChargeTab) AttachList() {
	ex.treeView.SetModel(ex.ExtraCharges)
	ex.blockAdvsTreeView.SetModel(ex.blockAdvsTreeFilter)
	ex.lineAdvsTreeView.SetModel(ex.lineAdvsTreeFilter)
	ex.SetEnableFilters(true)

}

func (ex *ExtraChargeTab) DetachList() {
	ex.treeView.SetModel(nil)
	ex.blockAdvsTreeView.SetModel(nil)
	ex.lineAdvsTreeView.SetModel(nil)
	ex.SetEnableFilters(false)
}

func (ex *ExtraChargeTab) UnsetModel() {
	ex.resetRightSide()
	bloks, _ := ex.blockAdvsTreeView.GetSelection()
	bloks.UnselectAll()
	lines, _ := ex.lineAdvsTreeView.GetSelection()
	lines.UnselectAll()
	ex.treeSelector.UnselectAll()
	ex.treeView.Hide()
	ex.DetachList()

}

func (ex *ExtraChargeTab) SetModel() {
	ex.AttachList()
	ex.treeView.Show()
}

func (ex *ExtraChargeTab) ResetSort() {
	ex.ExtraCharges.ResetDefaultSortFunc()
}
