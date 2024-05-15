package newcostratewindow

import "github.com/gotk3/gotk3/gtk"

func (ncs *NewCostRateWindow) Close() {
	ncs.window.Destroy()
}

func (ncs *NewCostRateWindow) Window() *gtk.Window {
	return ncs.window
}

func (ncs *NewCostRateWindow) Show() {
	ncs.window.Show()
}

func (ncs *NewCostRateWindow) SetName(name string) {
	ncs.costNameEntry.SetText(name)
}

func (ncs *NewCostRateWindow) Name() string {
	return ncs.costNameEntry.GetLayout().GetText()
}

func (ncs *NewCostRateWindow) SetCostForWordSymbol(c string) {
	ncs.costForWordSymbolEntry.HandlerBlock(ncs.signals.forWordSymbolCostInserted)
	ncs.costForWordSymbolEntry.SetText(c)
	ncs.costForWordSymbolEntry.HandlerUnblock(ncs.signals.forWordSymbolCostInserted)
}

func (ncs *NewCostRateWindow) CostForWordSymbol() string {
	return ncs.costForWordSymbolEntry.GetLayout().GetText()
}

func (ncs *NewCostRateWindow) SetCostForOneSquare(c string) {
	ncs.costOneSquareEntry.HandlerBlock(ncs.signals.forOneSquareInserted)
	ncs.costOneSquareEntry.SetText(c)
	ncs.costOneSquareEntry.HandlerUnblock(ncs.signals.forOneSquareInserted)
}

func (ncs *NewCostRateWindow) CostForOneSquare() string {
	return ncs.costOneSquareEntry.GetLayout().GetText()
}

func (ncs *NewCostRateWindow) SetCostForOneWord(b bool) {
	ncs.costForOneWordRadioButton.SetActive(b)
}

func (ncs *NewCostRateWindow) CostForOneWord() bool {
	return ncs.costForOneWordRadioButton.GetActive()
}
