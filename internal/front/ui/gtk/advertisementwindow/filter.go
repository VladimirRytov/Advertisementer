package advertisementwindow

import (
	"time"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"

	"github.com/gotk3/gotk3/gtk"
)

type Filter struct {
	from string
	to   string

	filterPopover       *gtk.Popover
	filterStack         *gtk.Stack
	filterStackSwitcher *gtk.StackSwitcher

	selectDateCalendar *gtk.Calendar

	fromDateLabel       *gtk.Label
	fromDateEntry       *gtk.Entry
	fromDateResetButton *gtk.Button
	fromDateRevealer    *gtk.Revealer
	fromDateCalendar    *gtk.Calendar

	toDateLabel       *gtk.Label
	toDateEntry       *gtk.Entry
	toDateResetButton *gtk.Button
	toDateRevealer    *gtk.Revealer
	toDateCalendar    *gtk.Calendar

	resetFiltersButton *gtk.Button
	applyFiltersButtom *gtk.Button
}

func (aw *AdvertisementsWindow) BuildFilter(buildFile *builder.Builder) {
	aw.filterPopover = buildFile.FetchPopover("filterPop")
	aw.filterStack = buildFile.FetchStack("filterPopStack")
	aw.filterStackSwitcher = buildFile.FetchStackSwitcher("filterPopStackSwitcher")

	aw.selectDateCalendar = buildFile.FetchCalendar("SelectDateCalendar")
	aw.selectDateCalendar.SelectDay(uint(time.Now().Day()))
	aw.selectDateCalendar.SelectMonth(uint(time.Now().Month()-1), uint(time.Now().Year()))

	aw.fromDateLabel = buildFile.FetchLabel("SinceDateLabel")
	aw.fromDateEntry = buildFile.FetchEntry("SinceDateEntry")
	aw.fromDateResetButton = buildFile.FetchButton("SinceDateResetButton")
	aw.fromDateRevealer = buildFile.FetchRevealer("SinceDateRevealer")
	aw.fromDateCalendar = buildFile.FetchCalendar("SinceDateCalendar")

	aw.toDateLabel = buildFile.FetchLabel("BeforeDateLabel")
	aw.toDateEntry = buildFile.FetchEntry("BeforeDateEntry")
	aw.toDateResetButton = buildFile.FetchButton("BeforeDateResetButton")
	aw.toDateRevealer = buildFile.FetchRevealer("BeforeDateRevealer")
	aw.toDateCalendar = buildFile.FetchCalendar("BeforeDateCalendar")

	aw.resetFiltersButton = buildFile.FetchButton("ResetDatesButton")
	aw.applyFiltersButtom = buildFile.FetchButton("ApplyDatesFilterButton")
}

func (aw *AdvertisementsWindow) SetFromDateEntry(date string) {
	aw.fromDateEntry.SetText(date)
}

func (aw *AdvertisementsWindow) FromDateEntry() string {
	return aw.fromDateEntry.GetLayout().GetText()
}

func (aw *AdvertisementsWindow) FromDate() string {
	return aw.from
}

func (aw *AdvertisementsWindow) setFromDate(date string) {
	aw.from = date
}

func (aw *AdvertisementsWindow) SetToDateEntry(date string) {
	aw.toDateEntry.SetText(date)
}

func (aw *AdvertisementsWindow) ToDateEntry() string {
	return aw.toDateEntry.GetLayout().GetText()
}

func (aw *AdvertisementsWindow) SetToDate(date string) {
	aw.to = date
}

func (aw *AdvertisementsWindow) ToDate() string {
	return aw.to
}
