package lineform

import (
	"strconv"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/gtk"
)

func (line *LineAdvPage) setReleaseCount(count int) {
	b, err := line.releaseCountEntry.GetBuffer()
	if err != nil {
		logging.Logger.Error("lineAdvPage.setSize: cannot get entry buffer", "error", err)
		return
	}
	b.SetText(strconv.Itoa(count))
}

func (line *LineAdvPage) id() int {
	id, _ := strconv.Atoi(line.idEntry.GetLayout().GetText())
	return id
}

func (line *LineAdvPage) orderID() int {
	iter, rowsExist := line.orderListStore.GetIterFirst()

	for rowsExist {

		selected, err := line.tools.InterfaceFromIter(iter, line.orderListStore.ToTreeModel(), 0)
		if err != nil {
			logging.Logger.Error("lineAdv orderID: cant get selected val from liststore", "error", err)
			return 0
		}

		if selected.(bool) {
			orderID, err := line.tools.InterfaceFromIter(iter, line.orderListStore.ToTreeModel(), 1)
			if err != nil {
				logging.Logger.Error("lineAdv orderID: cant get orderID from liststore", "error", err)
				return 0
			}
			return orderID.(int)
		}
		rowsExist = line.orderListStore.IterNext(iter)
	}
	return 0
}

func (line *LineAdvPage) toSelectedOrder(path *gtk.TreePath) {
	line.orderTreeView.ScrollToCell(path, nil, true, 0.5, 0)
}

func (line *LineAdvPage) releaseCount() int {
	i, _ := strconv.Atoi(line.releaseCountEntry.GetLayout().GetText())
	return i
}

func (line *LineAdvPage) cost() string {
	return line.costEntry.GetLayout().GetText()
}

func (line *LineAdvPage) setCost(cost string) {
	b, err := line.costEntry.GetBuffer()
	if err != nil {
		logging.Logger.Error("lineAdvPage.setCost: cannot get entry buffer", "error", err)
		return
	}
	b.SetText(cost)
}

func (line *LineAdvPage) text() string {
	text, err := line.textTextBuffer.GetText(line.textTextBuffer.GetStartIter(), line.textTextBuffer.GetEndIter(), true)
	if err != nil {
		logging.Logger.Error("lineAdvtext: an error occurred while fetching text from buffer", "error", err)
	}
	return text
}

func (line *LineAdvPage) releaseDates() []string {
	releaseDatesArr := make([]string, 0)
	iter, valsExist := line.releaseDatesListStore.GetIterFirst()
	for valsExist {
		str, err := line.tools.StringFromIter(iter, line.releaseDatesListStore.ToTreeModel(), 0)
		if err != nil {
			logging.Logger.Error("releaseDates: got error while getting vals from tagsSelectorListStore", "err", err)
		}
		releaseDatesArr = append(releaseDatesArr, str)
		valsExist = line.releaseDatesListStore.IterNext(iter)
	}
	return releaseDatesArr
}

func (line *LineAdvPage) selectedTags() string {
	tags := make([]presenter.SelectedTagDTO, 0)
	iter, valsExist := line.tagsListStore.GetIterFirst()
	for valsExist {
		str, err := line.tools.StringFromIter(iter, line.tagsListStore.ToTreeModel(), 1)
		if err != nil {
			logging.Logger.Error("selectedtags: got error while getting string val from tagsSelectorListStore", "err", err)
		}

		selected, err := line.tools.InterfaceFromIter(iter, line.tagsListStore.ToTreeModel(), 0)
		if err != nil {
			logging.Logger.Error("selectedtags: got error while getting interface val from tagsSelectorListStore", "err", err)
		}
		if b, ok := selected.(bool); ok {
			tags = append(tags, presenter.SelectedTagDTO{TagName: str, Selected: b})
		}
		valsExist = line.tagsListStore.IterNext(iter)
	}
	return line.conv.SelectedTagsToString(tags)
}

func (line *LineAdvPage) selectedExtraCharges() string {
	tags := make([]presenter.SelectedExtraChargeDTO, 0)
	iter, valsExist := line.extraChargeListStore.GetIterFirst()
	for valsExist {

		str, err := line.tools.StringFromIter(iter, line.extraChargeListStore.ToTreeModel(), 1)
		if err != nil {
			logging.Logger.Error("selectedExtraCharges: got error while getting string extraChargeSelectorListStore", "err", err)
		}

		selected, err := line.tools.InterfaceFromIter(iter, line.extraChargeListStore.ToTreeModel(), 0)
		if err != nil {
			logging.Logger.Error("selectedExtraCharges: got error while getting interface extraChargeSelectorListStore", "err", err)
		}

		if b, ok := selected.(bool); ok {
			tags = append(tags, presenter.SelectedExtraChargeDTO{ChargeName: str, Selected: b})
		}
		valsExist = line.extraChargeListStore.IterNext(iter)
	}
	return line.conv.SelectedExtraChargeToString(tags)
}

func (line *LineAdvPage) removeReleaseDate(releaseDate string) {
	iter, err := line.tools.FindValue(releaseDate, line.releaseDatesListStore, 0)
	if err != nil {
		logging.Logger.Error("removeReleaseDate: an error occured while while removing Line advertisement`s release date", "error", err)
		return
	}
	line.releaseDatesListStore.Remove(iter)
	line.decreaseReleaseCount()
}

func (line *LineAdvPage) appendReleaseDate(releaseDate string) {
	iter, valsExist := line.releaseDatesListStore.GetIterFirst()
	for valsExist {
		comp := line.tools.CompareNewTime(&line.releaseDatesListStore.TreeModel, iter, releaseDate, 0)
		if comp == 0 {
			err := line.releaseDatesListStore.SetValue(iter, 0, interface{}(releaseDate))
			if err != nil {
				logging.Logger.Error("LineAdvPage.AppendReleaseDate: an error occured while while appending Line advertisement`s release date", "error", err)
				return
			}
			return
		}
		valsExist = line.releaseDatesListStore.IterNext(iter)
	}
	err := line.releaseDatesListStore.SetValue(line.releaseDatesListStore.Append(), 0, interface{}(releaseDate))
	if err != nil {
		logging.Logger.Error("LineAdvPage.AppendReleaseDate: an error occured while while appending Line advertisement`s release date", "error", err)
		return
	}
	line.increaseReleaseCount()
}

func (line *LineAdvPage) increaseReleaseCount() {
	countStr := line.releaseCountEntry.GetLayout().GetText()
	countInt, err := strconv.Atoi(countStr)
	if err != nil {
		logging.Logger.Error("LineAdvPage.increaseReleaseCount: cannot conver count to int", "error", err)
	}
	line.releaseCountEntry.SetText(strconv.Itoa(countInt + 1))
}

func (line *LineAdvPage) decreaseReleaseCount() {
	countStr := line.releaseCountEntry.GetLayout().GetText()
	countInt, err := strconv.Atoi(countStr)
	if err != nil {
		logging.Logger.Error("LineAdvPage.decreaseReleaseCount: cannot conver count to int", "error", err)
	}
	line.releaseCountEntry.SetText(strconv.Itoa(countInt - 1))
}

func (line *LineAdvPage) UnsetModel() {
	line.orderTreeView.SetModel(nil)
	line.tagsTreeview.SetModel(nil)
	line.extraChargeTreeview.SetModel(nil)
}

func (line *LineAdvPage) SetModel() {
	line.orderTreeView.SetModel(line.orderListStore)
	line.tagsTreeview.SetModel(line.tagsListStore)
	line.extraChargeTreeview.SetModel(line.extraChargeListStore)
}
