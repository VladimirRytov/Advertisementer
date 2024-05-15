package newextracharge

import "github.com/gotk3/gotk3/gtk"

func (nexw *NewExtrachargeWindow) ChargeName() string {
	return nexw.nameEntry.GetLayout().GetText()
}

func (nexw *NewExtrachargeWindow) Multiplier() string {
	return nexw.multiplierEntry.GetLayout().GetText()
}

func (nexw *NewExtrachargeWindow) Close() {
	nexw.window.Close()
}

func (nexw *NewExtrachargeWindow) Show() {
	nexw.window.Show()
}

func (nexw *NewExtrachargeWindow) Window() *gtk.Window {
	return nexw.window
}
