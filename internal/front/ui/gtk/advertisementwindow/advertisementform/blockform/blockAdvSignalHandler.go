package blockform

import (
	"time"

	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type BlockAdvertisementsSignalHandler struct {
	selectedTagSwitcherToggled         glib.SignalHandle
	selectedExtraChargeSwitcherToggled glib.SignalHandle
	orderCellrenderToggleToggled       glib.SignalHandle
	releaseDatesMenuButtonClicked      glib.SignalHandle
	releaseDatesAddButtonClicked       glib.SignalHandle
	releaseDatesRemoveButtonClicked    glib.SignalHandle
	releaseDatesSelectorChanged        glib.SignalHandle
	costChanged                        glib.SignalHandle
	costCalculateButtonPressed         glib.SignalHandle
}

func (block *BlockAdvPage) bindSignals() {
	block.signalHandler = BlockAdvertisementsSignalHandler{
		selectedTagSwitcherToggled:         block.tagCellrenderToggle.Connect("toggled", block.TagToggled),
		selectedExtraChargeSwitcherToggled: block.extraChargeCellrenderToggle.Connect("toggled", block.ExtraChargeToggled),
		orderCellrenderToggleToggled:       block.orderCellrenderToggle.Connect("toggled", block.OrderToggled),
		releaseDatesMenuButtonClicked:      block.releaseDatesMenuButton.Connect("clicked", block.ChangeReleaseDatesButtonClicked),
		releaseDatesAddButtonClicked:       block.releaseDatesAppendButton.Connect("clicked", block.ReleaseDatesAppendButtonClicked),
		releaseDatesRemoveButtonClicked:    block.releaseDatesDeleteButton.Connect("clicked", block.ReleaseDatesDeleteButtonClicked),
		releaseDatesSelectorChanged:        block.releaseDatesSelection.Connect("changed", block.ReleaseDateChanged),
		costChanged:                        block.costEntry.Connect("insert-text", block.tools.CheckCostAdvString),
		costCalculateButtonPressed:         block.costCalculateButton.Connect("clicked", block.costCalculateButtonPressed),
	}
}

func (block *BlockAdvPage) OrderToggled(self *gtk.CellRendererToggle, pathStr string) {
	iter, valsExist := block.orderListStore.GetIterFirst()
	for valsExist {

		path, err := block.orderListStore.GetPath(iter)
		if err != nil {
			logging.Logger.Error("orderToggled: cannot get path", "error", err)
		}

		block.orderListStore.SetValue(iter, 0, pathStr == path.String())
		valsExist = block.orderListStore.IterNext(iter)
	}
}

func (block *BlockAdvPage) TagToggled(self *gtk.CellRendererToggle, path string) {
	iter, err := block.tagsListStore.GetIterFromString(path)
	if err != nil {
		logging.Logger.Error("tagToggled: cannot get iter from string", "error", err)
	}

	rawVal, _ := block.tools.InterfaceFromIter(iter, block.tagsListStore.ToTreeModel(), 0)
	if err != nil {
		logging.Logger.Error("tagToggled: cannot get interface val from treeModel", "error", err)
	}

	if val, ok := rawVal.(bool); ok {
		block.tagsListStore.SetValue(iter, 0, !val)
	} else {
		logging.Logger.Error("tagToggled: got not boolean value")
	}
}

func (block *BlockAdvPage) ExtraChargeToggled(self *gtk.CellRendererToggle, path string) {
	iter, err := block.extraChargeListStore.GetIterFromString(path)
	if err != nil {
		logging.Logger.Error("extraChargeToggled: cannot get iter from string", "error", err)
	}

	rawVal, err := block.tools.InterfaceFromIter(iter, block.extraChargeListStore.ToTreeModel(), 0)
	if err != nil {
		logging.Logger.Error("extraChargeToggled: cannot get interface val from treeModel", "error", err)
	}

	if val, ok := rawVal.(bool); ok {
		block.extraChargeListStore.SetValue(iter, 0, !val)
	} else {
		logging.Logger.Error("extraChargeToggled: got not boolean value")
	}
}

func (block *BlockAdvPage) ReleaseDatesDeleteButtonClicked() {
	y, m, d := block.releaseDatesCalendar.GetDate()
	releaseDate := block.conv.YearMonthDayToString(y, m+1, d)
	block.RemoveReleaseDate(releaseDate)
	selection, err := block.releaseDatesTreeview.GetSelection()
	if err != nil {
		logging.Logger.Error("releaseDatesDeleteButtonClicked: an error occured while getting releaseDates selector", "error", err)
		return
	}
	model, iter, _ := selection.GetSelected()

	date, err := block.tools.StringFromIter(iter, model.ToTreeModel(), 0)
	if err != nil {
		logging.Logger.Error("releaseDatesDeleteButtonClicked: an error occured while converting GVal to string", "error", err)
		return
	}
	selectedDate, _ := time.Parse("02.01.2006", date)
	block.releaseDatesCalendar.SelectMonth(uint(selectedDate.Month()-1), uint(selectedDate.Year()))
	block.releaseDatesCalendar.SelectDay(uint(selectedDate.Day()))
}

func (block *BlockAdvPage) ReleaseDatesAppendButtonClicked() {
	y, m, d := block.releaseDatesCalendar.GetDate()
	releaseDate := block.conv.YearMonthDayToString(y, m+1, d)
	block.appendReleaseDate(releaseDate)
}

func (block *BlockAdvPage) ChangeReleaseDatesButtonClicked() {
	_, _, ok := block.releaseDatesSelection.GetSelected()
	if !ok {
		block.releaseDatesCalendar.SelectMonth(uint(time.Now().Month()-1), uint(time.Now().Year()))
		block.releaseDatesCalendar.SelectDay(uint(time.Now().Day()))
	}
}

func (block *BlockAdvPage) ReleaseDateChanged(self *gtk.TreeSelection) {
	iModel, iter, _ := block.releaseDatesSelection.GetSelected()
	date, err := block.tools.StringFromIter(iter, iModel.ToTreeModel(), 0)
	if err != nil {
		logging.Logger.Error("releaseDateChanged: an error occured while getting releaseDates", "error", err)
		return
	}
	selectedDate, _ := time.Parse("02.01.2006", date)
	block.releaseDatesCalendar.SelectMonth(uint(selectedDate.Month()-1), uint(selectedDate.Year()))
	block.releaseDatesCalendar.SelectDay(uint(selectedDate.Day()))
}

func (block *BlockAdvPage) costCalculateButtonPressed() {
	b, err := block.costEntry.GetBuffer()
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.costCalculateButtonPressed: cannot get cost entry buffer", "error", err)
		return
	}
	block.app.RegisterReciever(b)
	blockAdv := block.FetchData()
	block.req.CalculateBlockAdvertisementCost(&blockAdv)
}
