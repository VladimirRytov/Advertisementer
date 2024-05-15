package newcostratewindow

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"

	"github.com/gotk3/gotk3/glib"
)

type signalHandler struct {
	windowClosed                glib.SignalHandle
	createOrUpdateButtonClicked glib.SignalHandle
	forWordSymbolCostInserted   glib.SignalHandle
	forOneSquareInserted        glib.SignalHandle
}

func (ncs *NewCostRateWindow) bindSignals(updateMode bool) {
	ncs.signals.windowClosed = ncs.window.Connect("destroy", ncs.window.Destroy)
	ncs.signals.forWordSymbolCostInserted = ncs.costForWordSymbolEntry.Connect("insert-text", ncs.tools.CheckCostString)
	ncs.signals.forOneSquareInserted = ncs.costOneSquareEntry.Connect("insert-text", ncs.tools.CheckCostString)
	if updateMode {
		ncs.signals.createOrUpdateButtonClicked = ncs.createButton.Connect("clicked", ncs.updateButtonPressed)
		return
	}
	ncs.signals.createOrUpdateButtonClicked = ncs.createButton.Connect("clicked", ncs.createButtonPressed)
}

func (ncs *NewCostRateWindow) createButtonPressed() {
	costRateForm := presenter.CostRateDTO{
		Name:            ncs.Name(),
		OneWordOrSymbol: ncs.CostForWordSymbol(),
		Onecm2:          ncs.CostForOneSquare(),
		CalcForOneWord:  ncs.CostForOneWord(),
	}
	err := ncs.req.CreateCostRate(&costRateForm)
	if err != nil {
		return
	}
}

func (ncs *NewCostRateWindow) updateButtonPressed() {
	costRateForm := presenter.CostRateDTO{
		Name:            ncs.Name(),
		OneWordOrSymbol: ncs.CostForWordSymbol(),
		Onecm2:          ncs.CostForOneSquare(),
		CalcForOneWord:  ncs.CostForOneWord(),
	}
	err := ncs.req.UpdateCostRate(&costRateForm)
	if err != nil {
		return
	}
	ncs.window.Close()
}
