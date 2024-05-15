package advertisementwindow

import (
	"context"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type AdvertisementsWindowHandler struct {
	mainWindowActivateDefault          glib.SignalHandle
	filterShowActualRadioButtonToggled glib.SignalHandle

	stackChildChanged    glib.SignalHandle
	fromDateIconPressed  glib.SignalHandle
	fromDateResetPressed glib.SignalHandle
	fromDateDateSelected glib.SignalHandle

	toDateIconPressed  glib.SignalHandle
	toDateResetPressed glib.SignalHandle
	toDateDateSelected glib.SignalHandle

	advReportClicked glib.SignalHandle

	applyFilterPressed       glib.SignalHandle
	resetFilterButtonPressed glib.SignalHandle
	resetFilterIcoPressed    glib.SignalHandle

	costRatesButtonClicked      glib.SignalHandle
	costRateChangeButtonClicked glib.SignalHandle

	exportDataButtonPressed glib.SignalHandle
	importDataButtonPressed glib.SignalHandle

	exitButtonPressed glib.SignalHandle
}

func (aw *AdvertisementsWindow) bindSignals() {
	logging.Logger.Debug("advertisementsWindow: binding loginwin signals")

	aw.signalHandler = AdvertisementsWindowHandler{
		mainWindowActivateDefault:          aw.MainWindow.Connect("activate-default", aw.activate),
		filterShowActualRadioButtonToggled: aw.filterShowActualRadioButton.Connect("toggled", aw.ActualReleasesToggled),

		stackChildChanged:    aw.filterStack.Connect("notify::visible-child", aw.stackChildChanged),
		fromDateIconPressed:  aw.fromDateEntry.Connect("icon-press", aw.fromDateIconPressed),
		fromDateResetPressed: aw.fromDateResetButton.Connect("clicked", aw.fromDateResetPressed),
		fromDateDateSelected: aw.fromDateCalendar.Connect("day-selected", aw.fromDateDaySelected),

		toDateIconPressed:  aw.toDateEntry.Connect("icon-press", aw.toDateIconPressed),
		toDateResetPressed: aw.toDateResetButton.Connect("clicked", aw.toDateResetPressed),
		toDateDateSelected: aw.toDateCalendar.Connect("day-selected", aw.toDateDaySelected),

		advReportClicked: aw.advReportButton.Connect("clicked", aw.advReportClicked),

		resetFilterIcoPressed:       aw.resetFilterIcoButton.Connect("clicked", aw.resetFilterPressed),
		resetFilterButtonPressed:    aw.resetFiltersButton.Connect("clicked", aw.resetFilterPressed),
		applyFilterPressed:          aw.applyFiltersButtom.Connect("clicked", aw.applyFilterPressed),
		costRatesButtonClicked:      aw.costRateSettingsButton.Connect("clicked", aw.costRateSettingsClicked),
		costRateChangeButtonClicked: aw.costRateChangeButton.Connect("clicked", aw.costRateSettingsClicked),

		exportDataButtonPressed: aw.exportDataButton.Connect("clicked", aw.exportToJsonClicked),
		importDataButtonPressed: aw.importDataButton.Connect("clicked", aw.importDataClicked),

		exitButtonPressed: aw.exitButton.Connect("clicked", aw.exitButtonPressed),
	}
}

func (aw *AdvertisementsWindow) activate() {
	logging.Logger.Debug("advertisementsWindow: activating window")
}

func (aw *AdvertisementsWindow) ActualReleasesToggled() {
	logging.Logger.Debug("advertisementsWindowHandler.ActualReleasesToggled: saving radioButtonStatus")
	act := &presenter.ShowData{
		Actual: aw.filterShowActualRadioButton.GetActive(),
	}
	err := aw.req.SaveConfig("ShowData", act)
	if err != nil {
		logging.Logger.Error("advertisementsWindowHandler.ActualReleasesToggled: cannot save radioButtonStatus", "error", err)
	}
}
func (aw *AdvertisementsWindow) stackChildChanged() {
	logging.Logger.Debug("advertisementsWindow.stackChildChanged: focus changed")
	aw.fromDateRevealer.SetRevealChild(false)
	aw.toDateRevealer.SetRevealChild(false)
}

func (aw *AdvertisementsWindow) applyFilterPressed() {
	logging.Logger.Debug("advertisementsWindow.applyFilterPressed: startFiltering")
	aw.NoteBook.BlockAllSignals()
	aw.NoteBook.SetSensetive(false)
	defer aw.NoteBook.UnblockAllSignals()
	defer aw.NoteBook.SetSensetive(true)
	switch aw.filterStack.GetVisibleChildName() {
	case "SelectDate":
		y, m, d := aw.selectDateCalendar.GetDate()
		date := aw.conv.YearMonthDayToString(y, m+1, d)
		aw.setFromDate(date)
		aw.SetToDate(aw.conv.YearMonthDayToString(y, m+1, d+1))
		aw.filterDataRangeLabel.SetText(date)
	case "SelectDateRange":
		if len(aw.FromDateEntry()) == 0 && len(aw.ToDateEntry()) == 0 {
			aw.resetFilterPressed()
			return
		}

		aw.setFromDate(aw.FromDateEntry())
		aw.SetToDate(aw.ToDateEntry())

		if len(aw.from) == 0 || len(aw.to) == 0 {
			aw.filterDataRangeLabel.SetText(aw.FromDateEntry() + "..." + aw.ToDateEntry())
		} else {
			aw.filterDataRangeLabel.SetText(aw.FromDateEntry() + "-" + aw.ToDateEntry())
		}
	}
	aw.NoteBook.EnableAdvertisementFilters(true)
	aw.filterPopover.Popdown()
	aw.NoteBook.RefilterBlock()
	aw.NoteBook.RefilterLine()
}

func (aw *AdvertisementsWindow) resetFilterPressed() {
	logging.Logger.Debug("advertisementsWindow.resetFilterPressed: resetFilters")
	aw.setFromDate("")
	aw.SetToDate("")
	aw.filterDataRangeLabel.SetText("Дата-Диапазон дат")
	aw.NoteBook.EnableAdvertisementFilters(false)
	aw.NoteBook.RefilterBlock()
	aw.NoteBook.RefilterLine()
}

func (aw *AdvertisementsWindow) fromDateIconPressed() {
	logging.Logger.Debug("advertisementsWindow.fromDateIconPressed: reveal calendar")
	aw.fromDateCalendar.HandlerBlock(aw.signalHandler.fromDateDateSelected)
	defer aw.fromDateCalendar.HandlerUnblock(aw.signalHandler.fromDateDateSelected)

	if len(aw.FromDateEntry()) == 0 {
		aw.fromDateCalendar.SelectDay(uint(time.Now().Day()))
		aw.fromDateCalendar.SelectMonth(uint(time.Now().Month()-1), uint(time.Now().Year()))
	} else {
		t, err := time.Parse("02.01.2006", aw.FromDateEntry())
		if err != nil {
			logging.Logger.Error("advertisementsWindow.fromDateIconPressed: cannot convert string date to time type", "error", err)
			return
		}
		aw.fromDateCalendar.SelectDay(uint(t.Day()))
		aw.fromDateCalendar.SelectMonth(uint(t.Month()-1), uint(t.Year()))
	}
	aw.fromDateRevealer.SetRevealChild(!aw.fromDateRevealer.GetRevealChild())
}

func (aw *AdvertisementsWindow) fromDateDaySelected() {
	logging.Logger.Debug("advertisementsWindow.fromDateDaySelected: fill entry")
	y, m, d := aw.fromDateCalendar.GetDate()
	aw.SetFromDateEntry(aw.conv.YearMonthDayToString(y, m+1, d))

}

func (aw *AdvertisementsWindow) fromDateResetPressed() {
	logging.Logger.Debug("advertisementsWindow.fromDateResetPressed: reset entry")
	aw.SetFromDateEntry("")
}

func (aw *AdvertisementsWindow) toDateIconPressed() {
	logging.Logger.Debug("advertisementsWindow.toDateIconPressed: reveal calendar")
	aw.toDateCalendar.HandlerBlock(aw.signalHandler.toDateDateSelected)
	defer aw.toDateCalendar.HandlerUnblock(aw.signalHandler.toDateDateSelected)

	if len(aw.ToDateEntry()) == 0 {
		aw.toDateCalendar.SelectDay(uint(time.Now().Day()))
		aw.toDateCalendar.SelectMonth(uint(time.Now().Month()-1), uint(time.Now().Year()))
	} else {
		t, err := time.Parse("02.01.2006", aw.ToDateEntry())
		if err != nil {
			logging.Logger.Error("advertisementsWindow.fromDateIconPressed: cannot convert string date to time type", "error", err)
			return
		}
		aw.toDateCalendar.SelectDay(uint(t.Day()))
		aw.toDateCalendar.SelectMonth(uint(t.Month()-1), uint(t.Year()))
	}
	aw.toDateRevealer.SetRevealChild(!aw.toDateRevealer.GetRevealChild())
}

func (aw *AdvertisementsWindow) toDateResetPressed() {
	logging.Logger.Debug("advertisementsWindow.toDateResetPressed: reset entry")

	aw.SetToDateEntry("")
}

func (aw *AdvertisementsWindow) toDateDaySelected() {
	logging.Logger.Debug("advertisementsWindow.toDateDaySelected: fill entry")
	y, m, d := aw.toDateCalendar.GetDate()
	aw.SetToDateEntry(aw.conv.YearMonthDayToString(y, m+1, d))
}

func (aw *AdvertisementsWindow) exportToJsonClicked() {
	aw.settingsPopover.Popdown()
	fileChooser, err := aw.dialogMaker.NewSaveDialog("Экспорт", aw.Window())
	if err != nil {
		logging.Logger.Error("advertisementsWindow.exportToJsonClicked: cannot create file diablog", "error", err)
		return
	}

	err = fileChooser.AddFileFilter("Json", "*.json")
	if err != nil {
		logging.Logger.Error("ImportJSONWindow.filePathButtonClicked: cannot bind json filter to filechooser", "error", err)
	}
	fileChooser.SetCurrentName("Выгрузка базы данных за " + time.Now().Format("02.01.2006") + ".json")
	fileChooser.BindResponseSignal(func(self *glib.Object, responce int) {
		if gtk.RESPONSE_ACCEPT == gtk.ResponseType(responce) {
			context, cancel := context.WithCancel(context.Background())
			progrWin := aw.app.CreateProgressWindow()
			progrWin.SetCancelFunc(cancel)
			progrWin.Show()
			progrWin.SetMessage("Выполняется экспорт данных")
			aw.req.StartExportingJson(context, fileChooser.GetFilename())
		}
	})
	fileChooser.Show()
}

func (aw *AdvertisementsWindow) importDataClicked() {
	aw.settingsPopover.Popdown()
	aw.app.CreateImportDataWindow().Show()
}

func (aw *AdvertisementsWindow) costRateSettingsClicked() {
	aw.settingsPopover.Popdown()
	aw.app.CreateCostRatesWindow().Show()
}

func (aw *AdvertisementsWindow) advReportClicked() {
	aw.app.NewAdvertisementReportWindow().Show()
}

func (aw *AdvertisementsWindow) exitButtonPressed() {
	aw.settingsPopover.Popdown()
}
