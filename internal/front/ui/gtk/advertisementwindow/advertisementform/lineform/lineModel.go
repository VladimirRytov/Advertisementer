package lineform

import (
	"strconv"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func (line *LineAdvPage) FillData(selectedLine *presenter.LineAdvertisementDTO) {
	line.releaseDatesSelection.UnselectAll()
	line.idEntry.SetText(strconv.Itoa(selectedLine.ID))
	path, err := line.listFil.MarkSelectedOrder(line.orderListStore, selectedLine.OrderID)
	if err == nil {
		line.toSelectedOrder(path)
	}
	line.setReleaseCount(selectedLine.ReleaseCount)
	line.setCost(selectedLine.Cost)
	line.textTextBuffer.SetText(selectedLine.Text)
	line.listFil.FillReleaseDates(line.releaseDatesListStore, selectedLine.ReleaseDates)
	line.listFil.MarkExtraCost(line.tagsListStore, selectedLine.Tags)
	line.listFil.MarkExtraCost(line.extraChargeListStore, selectedLine.ExtraCharge)
}

func (line *LineAdvPage) SetSensetive(s bool) {
	line.box.SetSensitive(s)
}

func (line *LineAdvPage) FetchData() presenter.LineAdvertisementDTO {
	var (
		orderid = 0
		id      = 0
	)
	if line.orderBox.GetVisible() {
		orderid = line.orderID()
	}
	if line.idBox.GetVisible() {
		id = line.id()
	}
	return presenter.LineAdvertisementDTO{
		Advertisement: presenter.Advertisement{
			ID:           id,
			OrderID:      orderid,
			ReleaseDates: line.conv.SelectedReleaseDatesToString(line.releaseDates()),
			ReleaseCount: line.releaseCount(),
			Cost:         line.cost(),
			Tags:         line.selectedTags(),
			ExtraCharge:  line.selectedExtraCharges(),
			Text:         line.text(),
		},
	}
}

func (line *LineAdvPage) Reset() {
	line.idEntry.SetText("")
	line.releaseCountEntry.SetText("")
	line.costEntry.SetText("")
	line.textTextBuffer.SetText("")
	line.SetSensetive(false)
}

func (line *LineAdvPage) SetNewAdvMode(mode bool) {
	line.idBox.SetVisible(!mode)
}

func (line *LineAdvPage) SetNewNestedAdvMode(mode bool) {
	line.idBox.SetVisible(!mode)
	line.orderBox.SetVisible(!mode)
}

func (line *LineAdvPage) Destroy() {
	line.box.Destroy()
}

func (line *LineAdvPage) ToSelectedOrder() {
	tPath, err := line.listFil.MarkSelectedOrder(line.orderListStore, line.orderID())
	if err != nil {
		logging.Logger.Warn("advPage.TreePath: cannot get path From listStore", "error", err)
		return
	}
	line.orderTreeView.ScrollToCell(tPath, nil, true, 0.5, 0)
}
