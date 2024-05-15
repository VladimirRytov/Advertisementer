package blockform

import (
	"strconv"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/gtk"
)

func (blk *BlockAdvPage) toSelectedOrder(path *gtk.TreePath) {
	blk.orderTreeView.ScrollToCell(path, nil, true, 0.5, 0)
}

func (blk *BlockAdvPage) path() string {
	return blk.file.FilePath()
}

func (blk *BlockAdvPage) setPath(file string) {
	blk.file.SetFilePath(file)
}

func (blkpg *BlockAdvPage) size() int {
	i, _ := strconv.Atoi(blkpg.sizeEntry.GetLayout().GetText())
	return i
}

func (blkpg *BlockAdvPage) setSize(size int) {
	b, err := blkpg.sizeEntry.GetBuffer()
	if err != nil {
		logging.Logger.Error("blockAdvPage.setSize: cannot get entry buffer", "error", err)
		return
	}

	b.SetText(strconv.Itoa(size))
}

func (blkpg *BlockAdvPage) setCost(cost string) {
	b, err := blkpg.costEntry.GetBuffer()
	if err != nil {
		logging.Logger.Error("blockAdvPage.setCost: cannot get entry buffer", "error", err)
		return
	}
	b.SetText(cost)
}

type FilePath interface {
	FilePath() string
	SetFilePath(string)
	Box() *gtk.Box
}

type ThinFilePath struct {
	app          Application
	req          Requests
	dialogMaker  ChooseSaveFileDialoger
	box          *gtk.Box
	filePathLink *gtk.LinkButton
	filePathView *gtk.Button
	signals      ThinSignals
}

func (blk *BlockAdvPage) attachThinPath(bld *builder.Builder, app Application, objMaker ChooseSaveFileDialoger, req Requests) {
	thn := &ThinFilePath{
		app:          app,
		req:          req,
		dialogMaker:  objMaker,
		box:          bld.FetchBox("ThinFileChooser"),
		filePathView: bld.FetchButton("ThinFilePathButton"),
		filePathLink: bld.FetchLinkButton("ThinFilePathLink"),
	}
	thn.BindSignals()
	blk.file = thn
	blk.fileNameBox.PackEnd(blk.file.Box(), true, false, 0)

}

type ThickFilePath struct {
	app           Application
	dialogMaker   ChooseSaveFileDialoger
	box           *gtk.Box
	filePathEntry *gtk.Entry
	filePathView  *gtk.Button
	signals       ThickSignals
}

func (blk *BlockAdvPage) attachThickPath(bld *builder.Builder, app Application, objMaker ChooseSaveFileDialoger) {
	thk := &ThickFilePath{
		dialogMaker:   objMaker,
		app:           app,
		box:           bld.FetchBox("ThickFileChooser"),
		filePathView:  bld.FetchButton("ThickFilePathButton"),
		filePathEntry: bld.FetchEntry("ThickFilePathEntry"),
	}
	thk.bindSignals()
	blk.file = thk
	blk.fileNameBox.PackEnd(blk.file.Box(), true, false, 0)
}

func (blkpg *BlockAdvPage) id() int {
	i, _ := strconv.Atoi(blkpg.idEntry.GetLayout().GetText())
	return i
}

func (blkpg *BlockAdvPage) orderID() int {
	var orderID interface{}
	imodel, err := blkpg.orderTreeView.GetModel()
	if err != nil {
		logging.Logger.Error("advertisementLineForm.orderID: cannot get treeModel", "error", err)
		return 0
	}
	model := imodel.ToTreeModel()

	model.ForEach(
		func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {

			selected, err := blkpg.tools.InterfaceFromIter(iter, model, 0)
			if err != nil {
				logging.Logger.Error("advertisementLineForm.orderID: cant get selected val from liststore", "error", err)
				return true
			}

			if selected.(bool) {
				orderID, err = blkpg.tools.InterfaceFromIter(iter, model, 1)
				if err != nil {
					logging.Logger.Error("advertisementLineForm.orderID: cant get orderID from liststore", "error", err)
					return true
				}
				return false
			}

			return false
		})

	return orderID.(int)
}

func (blkpg *BlockAdvPage) releaseCount() int {
	i, _ := strconv.Atoi(blkpg.releaseCountEntry.GetLayout().GetText())
	return i
}

func (blkpg *BlockAdvPage) setReleaseCount(count string) {
	b, err := blkpg.releaseCountEntry.GetBuffer()
	if err != nil {
		logging.Logger.Error("blockAdvPage.setReleaseCount: cannot get entry buffer", "error", err)
		return
	}
	b.SetText(count)
}

func (blkpg *BlockAdvPage) increaseReleaseCount() {
	countStr := blkpg.releaseCountEntry.GetLayout().GetText()
	countInt, err := strconv.Atoi(countStr)
	if err != nil {
		logging.Logger.Error("advertisementLineForm.increaseReleaseCount: cannot conver count to int", "error", err)
	}
	blkpg.releaseCountEntry.SetText(strconv.Itoa(countInt + 1))
}

func (blkpg *BlockAdvPage) decreaseReleaseCount() {
	countStr := blkpg.releaseCountEntry.GetLayout().GetText()
	countInt, err := strconv.Atoi(countStr)
	if err != nil {
		logging.Logger.Error("advertisementLineForm.decreaseReleaseCount: cannot conver count to int", "error", err)
	}
	blkpg.releaseCountEntry.SetText(strconv.Itoa(countInt - 1))
}

func (blkpg *BlockAdvPage) cost() string {
	return blkpg.costEntry.GetLayout().GetText()
}

func (blkpg *BlockAdvPage) text() string {
	text, err := blkpg.textTextBuffer.GetText(blkpg.textTextBuffer.GetStartIter(), blkpg.textTextBuffer.GetEndIter(), true)
	if err != nil {
		logging.Logger.Error("advertisementLineForm.text: an error occurred while fetching text from buffer", "error", err)
	}
	return text
}

func (blkpg *BlockAdvPage) selectedTags() string {
	tags := make([]presenter.SelectedTagDTO, 0)
	blkpg.tagsListStore.ForEach(func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
		str, err := blkpg.tools.StringFromIter(iter, model, 1)
		if err != nil {
			logging.Logger.Error("advertisementLineForm.selectedTags: got error while getting string value from treeModel", "err", err)
			return false
		}

		selected, err := blkpg.tools.InterfaceFromIter(iter, model, 0)
		if err != nil {
			logging.Logger.Error("advertisementLineForm.selectedTags: got error while getting interface value from treeModel", "err", err)
			return false
		}
		if b, ok := selected.(bool); ok {
			tags = append(tags, presenter.SelectedTagDTO{TagName: str, Selected: b})
		}
		return false
	})
	return blkpg.conv.SelectedTagsToString(tags)
}

func (blkpg *BlockAdvPage) selectedExtraCharges() string {
	extraCharges := make([]presenter.SelectedExtraChargeDTO, 0)
	blkpg.extraChargeListStore.ForEach(func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
		str, err := blkpg.tools.StringFromIter(iter, model, 1)
		if err != nil {
			logging.Logger.Error("advertisementLineForm.selectedExtraCharges: got error while getting string val from treemovel", "err", err)
			return false
		}

		selected, err := blkpg.tools.InterfaceFromIter(iter, model, 0)
		if err != nil {
			logging.Logger.Error("advertisementLineForm.selectedExtraCharges: got error while getting interface val from treemovel", "err", err)
			return false
		}
		if b, ok := selected.(bool); ok {
			extraCharges = append(extraCharges, presenter.SelectedExtraChargeDTO{ChargeName: str, Selected: b})
		}
		return false
	})
	return blkpg.conv.SelectedExtraChargeToString(extraCharges)
}

func (blkpg *BlockAdvPage) releaseDates() []string {
	releaseDatesArr := make([]string, 0)
	iter, valsExist := blkpg.releaseDatesListStore.GetIterFirst()

	for valsExist {
		str, err := blkpg.tools.StringFromIter(iter, blkpg.releaseDatesListStore.ToTreeModel(), 0)
		if err != nil {
			logging.Logger.Error("advertisementLineForm.releaseDates: got error while getting vals from tagsSelectorListStore", "err", err)
		}
		releaseDatesArr = append(releaseDatesArr, str)
		valsExist = blkpg.releaseDatesListStore.IterNext(iter)
	}
	return releaseDatesArr
}

func (blkpg *BlockAdvPage) RemoveReleaseDate(releaseDate string) {
	iter, err := blkpg.tools.FindValue(releaseDate, blkpg.releaseDatesListStore, 0)
	if err != nil {
		logging.Logger.Error("advertisementLineForm.RemoveReleaseDate: an error occured while while removing Line advertisement`s release date", "error", err)
		return
	}
	blkpg.releaseDatesListStore.Remove(iter)
	if blkpg.releaseCount() != 0 {
		blkpg.decreaseReleaseCount()
	}
}

func (blkpg *BlockAdvPage) appendReleaseDate(releaseDate string) {
	iter, valsExist := blkpg.releaseDatesListStore.GetIterFirst()
	for valsExist {
		comp := blkpg.tools.CompareNewTime(&blkpg.releaseDatesListStore.TreeModel, iter, releaseDate, 0)
		if comp == 0 {
			err := blkpg.releaseDatesListStore.SetValue(iter, 0, interface{}(releaseDate))
			if err != nil {
				logging.Logger.Error("advertisementLineForm.AppendReleaseDate: an error occured while while appending Line advertisement`s release date", "error", err)
				return
			}
			return
		}
		valsExist = blkpg.releaseDatesListStore.IterNext(iter)
	}
	err := blkpg.releaseDatesListStore.SetValue(blkpg.releaseDatesListStore.Append(), 0, interface{}(releaseDate))
	if err != nil {
		logging.Logger.Error("advertisementLineForm.AppendReleaseDate: an error occured while while appending Line advertisement`s release date", "error", err)
		return
	}
	blkpg.increaseReleaseCount()
}

func (blkpg *BlockAdvPage) UnsetModel() {
	blkpg.orderTreeView.SetModel(nil)
	blkpg.tagsTreeview.SetModel(nil)
	blkpg.extraChargeTreeview.SetModel(nil)
}

func (blkpg *BlockAdvPage) SetModel() {
	blkpg.orderTreeView.SetModel(blkpg.orderListStore)
	blkpg.tagsTreeview.SetModel(blkpg.tagsListStore)
	blkpg.extraChargeTreeview.SetModel(blkpg.extraChargeListStore)
}
