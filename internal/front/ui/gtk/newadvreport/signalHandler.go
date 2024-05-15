package newadvreport

import (
	"context"
	"errors"
	"net/url"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type SignalHandler struct {
	sourceDirButtonPressed  glib.SignalHandle
	destDirButtonPressed    glib.SignalHandle
	dateSelected            glib.SignalHandle
	dateCalendarScrollEvent glib.SignalHandle

	stackChildChanged glib.SignalHandle

	fromDateIconPressed         glib.SignalHandle
	fromDateSelected            glib.SignalHandle
	fromDateReset               glib.SignalHandle
	fromDateCalendarScrollEvent glib.SignalHandle

	toDateIconPressed         glib.SignalHandle
	toDateSelected            glib.SignalHandle
	todateReset               glib.SignalHandle
	toDateCalendarScrollEvent glib.SignalHandle

	createButtonPressed glib.SignalHandle
}

func (nar *NewAdvertisementReport) bindSignals() {
	nar.signals.sourceDirButtonPressed = nar.sourceFolderButton.Connect("clicked", nar.sourceDirButtonPressed)
	nar.signals.destDirButtonPressed = nar.deployFolderButton.Connect("clicked", nar.destDirButtonPressed)

	nar.signals.dateSelected = nar.selectedDateCalendar.Connect("day-selected", nar.dateSelected)
	nar.signals.dateCalendarScrollEvent = nar.selectedDateCalendar.Connect("scroll-event", nar.disableScroll)

	nar.signals.stackChildChanged = nar.dateFilterStack.Connect("notify::visible-child", nar.stackChildChanged)

	nar.signals.fromDateReset = nar.fromDateResetButton.Connect("clicked", nar.fromDateResetPressed)
	nar.signals.fromDateIconPressed = nar.fromDateEntry.Connect("icon-press", nar.fromDateIconPressed)
	nar.signals.fromDateSelected = nar.fromDateCalendar.Connect("day-selected", nar.fromDateSelected)
	nar.signals.fromDateCalendarScrollEvent = nar.fromDateCalendar.Connect("scroll-event", nar.disableScroll)

	nar.signals.todateReset = nar.toDateResetButton.Connect("clicked", nar.toDateResetPressed)
	nar.signals.toDateIconPressed = nar.toDateEntry.Connect("icon-press", nar.toDateIconPressed)
	nar.signals.toDateSelected = nar.toDateCalendar.Connect("day-selected", nar.toDateSelected)
	nar.signals.toDateCalendarScrollEvent = nar.toDateCalendar.Connect("scroll-event", nar.disableScroll)
	nar.signals.createButtonPressed = nar.createButton.Connect("clicked", nar.createButtonPressed)
}

func (nar *NewAdvertisementReport) stackChildChanged() {
	logging.Logger.Debug("advertisementsWindow.stackChildChanged: focus changed")
	nar.fromDateRevealer.SetRevealChild(false)
	nar.toDateRevealer.SetRevealChild(false)
}

func (nar *NewAdvertisementReport) sourceDirButtonPressed() {
	folderChooser, err := nar.dialogCreator.NewFolderChooseDialog("Папка с блочными объявлениями", nar.window)
	if err != nil {
		logging.Logger.Error("newAdvertisementReport.sourceDirButtonPressed: cannot create choode dialog", "error", err)
		return
	}

	folderChooser.BindResponseSignal(func(self *glib.Object, responce int) {
		if gtk.ResponseType(responce) == gtk.RESPONSE_ACCEPT {
			folderURI, err := url.Parse(folderChooser.GetURI())
			if err != nil {
				return
			}
			folder := nar.conv.ParsePath(folderURI.Path)
			err = nar.req.SaveConfig("sourcePath", &folder)
			if err != nil {
				logging.Logger.Error("newAdvertisementReport.sourceDirButtonPressed: cannot save config \"sourcePath\"", "error", err)
			}
			nar.setSourceFolder(folder)
		}
	})
	folderChooser.Show()
}

func (nar *NewAdvertisementReport) destDirButtonPressed() {
	folderChooser, err := nar.dialogCreator.NewFolderChooseDialog("Папка для выгрузки", nar.window)
	if err != nil {
		logging.Logger.Error("newAdvertisementReport.sourceDirButtonPressed: cannot create choode dialog", "error", err)
		return
	}

	folderChooser.BindResponseSignal(func(self *glib.Object, responce int) {
		if gtk.ResponseType(responce) == gtk.RESPONSE_ACCEPT {
			folderURI, err := url.Parse(folderChooser.GetURI())
			if err != nil {
				return
			}
			folder := nar.conv.ParsePath(folderURI.Path)
			err = nar.req.SaveConfig("deployPath", &folder)
			if err != nil {
				logging.Logger.Error("newAdvertisementReport.sourceDirButtonPressed: cannot save config \"deployPath\"", "error", err)
			}
			nar.setDeployFolder(folder)
		}
	})
	folderChooser.Show()
}

func (nar *NewAdvertisementReport) dateSelected() {
	logging.Logger.Debug("newAdvertisementReport.dateSelected: fill entry")
	y, m, d := nar.selectedDateCalendar.GetDate()
	nar.setSelectedDate(nar.conv.YearMonthDayToString(y, m+1, d))
}

func (nar *NewAdvertisementReport) fromDateIconPressed() {
	logging.Logger.Debug("advertisementsWindow.fromDateIconPressed: reveal calendar")
	nar.fromDateCalendar.HandlerBlock(nar.signals.fromDateSelected)
	defer nar.fromDateCalendar.HandlerUnblock(nar.signals.fromDateSelected)

	if len(nar.fromDate()) == 0 {
		nar.fromDateCalendar.SelectDay(uint(time.Now().Day()))
		nar.fromDateCalendar.SelectMonth(uint(time.Now().Month()-1), uint(time.Now().Year()))
	} else {
		t, err := time.Parse("02.01.2006", nar.fromDate())
		if err != nil {
			logging.Logger.Error("advertisementsWindow.fromDateIconPressed: cannot convert string date to time type", "error", err)
			return
		}
		nar.fromDateCalendar.SelectDay(uint(t.Day()))
		nar.fromDateCalendar.SelectMonth(uint(t.Month()-1), uint(t.Year()))
	}
	nar.fromDateRevealer.SetRevealChild(!nar.fromDateRevealer.GetRevealChild())
}

func (nar *NewAdvertisementReport) fromDateSelected() {
	logging.Logger.Debug("newAdvertisementReport.dateSelected: fill entry")
	y, m, d := nar.fromDateCalendar.GetDate()
	nar.setFromDate(nar.conv.YearMonthDayToString(y, m+1, d))
}

func (nar *NewAdvertisementReport) fromDateResetPressed() {
	logging.Logger.Debug("advertisementsWindow.fromDateResetPressed: reset entry")
	nar.setFromDate("")
}

func (nar *NewAdvertisementReport) toDateIconPressed() {
	logging.Logger.Debug("advertisementsWindow.toDateIconPressed: reveal calendar")
	nar.toDateCalendar.HandlerBlock(nar.signals.toDateSelected)
	defer nar.toDateCalendar.HandlerUnblock(nar.signals.toDateSelected)

	if len(nar.toDate()) == 0 {
		nar.toDateCalendar.SelectDay(uint(time.Now().Day()))
		nar.toDateCalendar.SelectMonth(uint(time.Now().Month()-1), uint(time.Now().Year()))
	} else {
		t, err := time.Parse("02.01.2006", nar.toDate())
		if err != nil {
			logging.Logger.Error("advertisementsWindow.fromDateIconPressed: cannot convert string date to time type", "error", err)
			return
		}
		nar.toDateCalendar.SelectDay(uint(t.Day()))
		nar.toDateCalendar.SelectMonth(uint(t.Month()-1), uint(t.Year()))
	}
	nar.toDateRevealer.SetRevealChild(!nar.toDateRevealer.GetRevealChild())
}

func (nar *NewAdvertisementReport) toDateResetPressed() {
	logging.Logger.Debug("advertisementsWindow.toDateResetPressed: reset entry")
	nar.setToDate("")
}

func (nar *NewAdvertisementReport) toDateSelected() {
	logging.Logger.Debug("newAdvertisementReport.dateSelected: fill entry")
	y, m, d := nar.toDateCalendar.GetDate()
	nar.setToDate(nar.conv.YearMonthDayToString(y, m+1, d))

}

func (nar *NewAdvertisementReport) createButtonPressed() {
	var dateRange string
	rep := presenter.ReportParams{
		ReportType:       "Exel",
		BlocksFolderPath: nar.sourceFolder(),
		DeployPath:       nar.deployFolder(),
	}

	switch nar.selectedDateMode() {
	case "SelectDate":
		rep.FromDate = nar.selectedDate()
		rep.ToDate = nar.selectedDate()
		dateRange = nar.selectedDate()
	case "SelectDateRange":
		rep.FromDate = nar.fromDate()
		rep.ToDate = nar.toDate()
		dateRange = rep.FromDate + "-" + rep.ToDate

	default:
		nar.app.NewErrorWindow(errors.New("выбран неизвестный тип дат"))
		return
	}

	progresWin := nar.app.CreateProgressWindow()
	progresWin.SetMessage("Генерируется сводка за " + dateRange)
	ctx, cancel := context.WithCancel(context.Background())
	progresWin.SetCancelFunc(cancel)
	err := nar.req.CreateAfvertisementReport(ctx, &rep)
	if err != nil {
		progresWin.Close()
		nar.app.NewErrorWindow(err)
		return
	}
	progresWin.Show()
}

func (nar *NewAdvertisementReport) disableScroll() bool {
	return true
}
