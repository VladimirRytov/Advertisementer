package blockform

import (
	"strconv"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/gtk"
)

func (blockAdv *BlockAdvPage) FillData(block *presenter.BlockAdvertisementDTO) {
	blockAdv.idEntry.SetText(strconv.Itoa(block.ID))
	path, err := blockAdv.listFil.MarkSelectedOrder(blockAdv.orderListStore, block.OrderID)
	if err == nil {
		blockAdv.toSelectedOrder(path)
	}
	blockAdv.releaseCountEntry.SetText(strconv.Itoa(block.ReleaseCount))
	blockAdv.listFil.FillReleaseDates(blockAdv.releaseDatesListStore, block.ReleaseDates)
	blockAdv.listFil.MarkExtraCost(blockAdv.tagsListStore, block.Tags)
	blockAdv.listFil.MarkExtraCost(blockAdv.extraChargeListStore, block.ExtraCharge)
	blockAdv.setSize(block.Size)
	blockAdv.setCost(block.Cost)
	blockAdv.textTextBuffer.SetText(block.Text)
	blockAdv.setPath(block.FileName)
}

func (blockAdv *BlockAdvPage) Reset() {
	blockAdv.idEntry.SetText("")
	blockAdv.releaseCountEntry.SetText("")
	blockAdv.costEntry.SetText("")
	blockAdv.textTextBuffer.SetText("")
	blockAdv.sizeEntry.SetText("")
	blockAdv.setPath("")
	blockAdv.SetSensetive(false)
}

func (blockAdv *BlockAdvPage) SetSensetive(s bool) {
	blockAdv.box.SetSensitive(s)
}

func (blockAdv *BlockAdvPage) FetchData() presenter.BlockAdvertisementDTO {
	var (
		orderid = 0
		id      = 0
	)
	if blockAdv.orderBox.GetVisible() {
		orderid = blockAdv.orderID()
	}
	if blockAdv.idBox.GetVisible() {
		id = blockAdv.id()
	}
	return presenter.BlockAdvertisementDTO{
		Advertisement: presenter.Advertisement{
			ID:           id,
			OrderID:      orderid,
			ReleaseDates: blockAdv.conv.SelectedReleaseDatesToString(blockAdv.releaseDates()),
			ReleaseCount: blockAdv.releaseCount(),
			Cost:         blockAdv.cost(),
			Tags:         blockAdv.selectedTags(),
			ExtraCharge:  blockAdv.selectedExtraCharges(),
			Text:         blockAdv.text(),
		},
		Size:     blockAdv.size(),
		FileName: blockAdv.path(),
	}
}

func (blockAdv *BlockAdvPage) SetNewAdvMode(mode bool) {
	blockAdv.idBox.SetVisible(!mode)
}

func (blockAdv *BlockAdvPage) SetNewNestedAdvMode(mode bool) {
	blockAdv.idBox.SetVisible(!mode)
	blockAdv.orderBox.SetVisible(!mode)
}

func (blockAdv *BlockAdvPage) Destroy() {
	blockAdv.box.Destroy()
}

func (blockAdv *BlockAdvPage) Widget() *gtk.Widget {
	return blockAdv.box.ToWidget()
}

func (blockAdv *BlockAdvPage) ToSelectedOrder() {
	tPath, err := blockAdv.listFil.MarkSelectedOrder(blockAdv.orderListStore, blockAdv.orderID())
	if err != nil {
		logging.Logger.Warn("advPage.TreePath: cannot get path From listStore", "error", err)
		return
	}
	blockAdv.orderTreeView.ScrollToCell(tPath, nil, true, 0.5, 0)
}
