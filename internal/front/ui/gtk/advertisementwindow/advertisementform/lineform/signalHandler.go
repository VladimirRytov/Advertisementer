package lineform

import (
	"time"

	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type advSignalHandler struct {
	orderToggled                    glib.SignalHandle
	tagToggled                      glib.SignalHandle
	extraChargeToggled              glib.SignalHandle
	releaseDatesMenuButtonClicked   glib.SignalHandle
	releaseDatesAddButtonClicked    glib.SignalHandle
	releaseDatesRemoveButtonClicked glib.SignalHandle
	releaseDatesSelectorChanged     glib.SignalHandle
	costChanged                     glib.SignalHandle
	calculateButtonPressed          glib.SignalHandle
}

func (line *LineAdvPage) bindSignals() {
	line.signalHandler = advSignalHandler{
		orderToggled:                    line.orderCellrenderToggle.Connect("toggled", line.OrderToggled),
		tagToggled:                      line.tagCellrenderToggle.Connect("toggled", line.TagToggled),
		extraChargeToggled:              line.extraChargeCellrenderToggle.Connect("toggled", line.ExtraChargeToggled),
		releaseDatesMenuButtonClicked:   line.releaseDatesMenuButton.Connect("clicked", line.changeReleaseDatesButtonClicked),
		releaseDatesAddButtonClicked:    line.releaseDatesAppendButton.Connect("clicked", line.releaseDatesAppendButtonClicked),
		releaseDatesRemoveButtonClicked: line.releaseDatesDeleteButton.Connect("clicked", line.ReleaseDatesDeleteButtonClicked),
		releaseDatesSelectorChanged:     line.releaseDatesSelection.Connect("changed", line.releaseDateChanged),
		costChanged:                     line.costEntry.Connect("insert-text", line.tools.CheckCostAdvString),
		calculateButtonPressed:          line.costCalculateButton.Connect("clicked", line.costCalculateButtonPressed),
	}
}

func (line *LineAdvPage) TagToggled(self *gtk.CellRendererToggle, path string) {
	iter, err := line.tagsListStore.GetIterFromString(path)
	if err != nil {
		logging.Logger.Error("tagToggled: cannot get iter from string", "error", err)
	}
	rawVal, err := line.tools.InterfaceFromIter(iter, line.tagsListStore.ToTreeModel(), 0)
	if err != nil {
		logging.Logger.Error("tagToggled: cannot get value liststore", "error", err)
	}
	if val, ok := rawVal.(bool); ok {
		line.tagsListStore.SetValue(iter, 0, !val)
	} else {
		logging.Logger.Error("tagToggled: got not boolean value")
	}
}

func (line *LineAdvPage) ExtraChargeToggled(self *gtk.CellRendererToggle, path string) {
	iter, err := line.extraChargeListStore.GetIterFromString(path)
	if err != nil {
		logging.Logger.Error("extraChargeToggled: cannot get iter from string", "error", err)
	}

	rawVal, err := line.tools.InterfaceFromIter(iter, line.extraChargeListStore.ToTreeModel(), 0)
	if err != nil {
		logging.Logger.Error("extraChargeToggled: cannot get val from liststore", "error", err)
	}

	if val, ok := rawVal.(bool); ok {
		line.extraChargeListStore.SetValue(iter, 0, !val)
	} else {
		logging.Logger.Error("extraChargeToggled: got not boolean value")
	}
}

func (line *LineAdvPage) OrderToggled(self *gtk.CellRendererToggle, pathStr string) {
	iter, valsExist := line.orderListStore.GetIterFirst()
	for valsExist {

		path, err := line.orderListStore.GetPath(iter)
		if err != nil {
			logging.Logger.Error("orderToggled: cannot get path", "error", err)
		}

		line.orderListStore.SetValue(iter, 0, pathStr == path.String())
		valsExist = line.orderListStore.IterNext(iter)
	}
}

func (line *LineAdvPage) ReleaseDatesDeleteButtonClicked() {
	y, m, d := line.releaseDatesCalendar.GetDate()
	releaseDate := line.conv.YearMonthDayToString(y, m+1, d)
	line.removeReleaseDate(releaseDate)
}

func (line *LineAdvPage) releaseDatesAppendButtonClicked() {
	y, m, d := line.releaseDatesCalendar.GetDate()
	releaseDate := line.conv.YearMonthDayToString(y, m+1, d)
	line.appendReleaseDate(releaseDate)
}

func (line *LineAdvPage) changeReleaseDatesButtonClicked() {
	_, _, ok := line.releaseDatesSelection.GetSelected()
	if !ok {
		line.releaseDatesCalendar.SelectMonth(uint(time.Now().Month()-1), uint(time.Now().Year()))
		line.releaseDatesCalendar.SelectDay(uint(time.Now().Day()))
	}
}

func (line *LineAdvPage) releaseDateChanged(self *gtk.TreeSelection) {
	iModel, iter, _ := line.releaseDatesSelection.GetSelected()
	date, err := line.tools.StringFromIter(iter, iModel.ToTreeModel(), 0)
	if err != nil {
		logging.Logger.Error("releaseDateChanged: an error occured while getting releaseDates", "error", err)
		return
	}
	selectedDate, _ := time.Parse("02.01.2006", date)
	line.releaseDatesCalendar.SelectMonth(uint(selectedDate.Month()-1), uint(selectedDate.Year()))
	line.releaseDatesCalendar.SelectDay(uint(selectedDate.Day()))
}

func (line *LineAdvPage) costCalculateButtonPressed() {
	b, err := line.costEntry.GetBuffer()
	if err != nil {
		logging.Logger.Error("lineAdvertisementsTab.costCalculateButtonPressed: cannot get cost entry buffer", "error", err)
		return
	}
	line.app.RegisterReciever(b)
	lineAdv := line.FetchData()
	line.req.CalculateLineAdvertisementCost(&lineAdv)
}
