package newadvreport

import (
	"github.com/gotk3/gotk3/gtk"
)

func (nar *NewAdvertisementReport) Window() *gtk.Window {
	return nar.window
}

func (nar *NewAdvertisementReport) Show() {
	nar.window.Show()
}

func (nar *NewAdvertisementReport) Close() {
	nar.window.Destroy()
}

func (nar *NewAdvertisementReport) setSourceFolder(folder string) {
	nar.sourceFolderEntry.SetText(folder)
}

func (nar *NewAdvertisementReport) sourceFolder() string {
	return nar.sourceFolderEntry.GetLayout().GetText()
}

func (nar *NewAdvertisementReport) deployFolder() string {
	return nar.deployFolderEntry.GetLayout().GetText()
}

func (nar *NewAdvertisementReport) setDeployFolder(folder string) {
	nar.deployFolderEntry.SetText(folder)
}

func (nar *NewAdvertisementReport) setSelectedDate(date string) {
	nar.selectedDateEntry.SetText(date)
}

func (nar *NewAdvertisementReport) selectedDate() string {
	return nar.selectedDateEntry.GetLayout().GetText()
}

func (nar *NewAdvertisementReport) setFromDate(date string) {
	nar.fromDateEntry.SetText(date)
}

func (nar *NewAdvertisementReport) fromDate() string {
	return nar.fromDateEntry.GetLayout().GetText()
}

func (nar *NewAdvertisementReport) setToDate(date string) {
	nar.toDateEntry.SetText(date)
}

func (nar *NewAdvertisementReport) toDate() string {
	return nar.toDateEntry.GetLayout().GetText()
}

func (nar *NewAdvertisementReport) selectedDateMode() string {
	return nar.dateFilterStack.GetVisibleChildName()
}
