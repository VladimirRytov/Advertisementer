package importjsonwindow

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/application"
	"github.com/gotk3/gotk3/gtk"
)

func (ijw *ImportJSONWindow) Window() *gtk.Window {
	return ijw.window
}

func (ijw *ImportJSONWindow) Show() {
	ijw.window.Show()
}

func (ijw *ImportJSONWindow) Close() {
	ijw.window.Destroy()
}

func (ijw *ImportJSONWindow) SetFilePath(path string) {
	ijw.filePathEntry.SetText(path)
}
func (ijw *ImportJSONWindow) FilePath() string {
	return ijw.filePathEntry.GetLayout().GetText()
}

func (ijw *ImportJSONWindow) AllBlockAdv() bool {
	return ijw.blockAdvRadioButton.GetActive()
}

func (ijw *ImportJSONWindow) AllLineAdv() bool {
	return ijw.lineAdvRadioButton.GetActive()
}

func (ijw *ImportJSONWindow) Path() string {
	return ijw.filePathEntry.GetLayout().GetText()
}

func (ijw *ImportJSONWindow) ActualClients() bool {
	return ijw.clientOnlyActualRadioButton.GetActive()
}

func (ijw *ImportJSONWindow) IgnoreErrors() bool {
	return ijw.ignoreErrorsCheckButton.GetActive()
}

func (ijw *ImportJSONWindow) Tags() bool {
	return ijw.tagsCheckButton.GetActive()
}

func (ijw *ImportJSONWindow) ExtraCharges() bool {
	return ijw.extraChargesCheckButton.GetActive()
}

func (ijw *ImportJSONWindow) CostRates() bool {
	return ijw.costRateCheckButton.GetActive()
}

func (ijw *ImportJSONWindow) Mode() bool {
	return ijw.app.Mode() == application.ThickMode
}
