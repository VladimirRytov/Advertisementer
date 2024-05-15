package newadvertisement

import (
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func (adadw *AddAdvertisementWindow) bindSignals() {
	adadw.window.Connect("destroy", adadw.windowDestroyed)
	adadw.advStack.Connect("notify::visible-child", adadw.pageChanged)
	adadw.createAdvButton.Connect("clicked", adadw.createButtonClicked)
}

func (adadw *AddAdvertisementWindow) windowDestroyed() {
	adadw.app.RefilterLists()
	adadw.window.Destroy()
}

func (adadw *AddAdvertisementWindow) createButtonClicked() {
	switch adadw.advStack.GetVisibleChildName() {
	case "BlockAdvertisementStack":
		blockAdv := adadw.blockPage.FetchData()
		logging.Logger.Debug("BlockAdvertisement.createButtonClicked: sending BlockAdv dto", "data", blockAdv)
		adadw.req.CreateBlockAdvertisement(&blockAdv)
	case "LineAdvertisementStack":
		lineAdv := adadw.linePage.FetchData()
		logging.Logger.Debug("BlockAdvertisement.createButtonClicked: sending LinekAdv dto", "data", lineAdv)
		adadw.req.CreateLineAdvertisement(&lineAdv)
	}
}

func (adadw *AddAdvertisementWindow) pageChanged() {
	switch adadw.advStack.GetVisibleChildName() {
	case "BlockAdvertisementStack":
		adadw.blockPage.ToSelectedOrder()
	case "LineAdvertisementStack":
		adadw.linePage.ToSelectedOrder()
	}

}
