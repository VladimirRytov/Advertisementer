package neworder

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/gtk"
)

func (now *NewOrderWindow) client() string {
	var client string
	now.clientListStore.ForEach(func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
		selected, err := now.tools.InterfaceFromIter(iter, &now.clientListStore.TreeModel, 0)
		if err != nil {
			logging.Logger.Error("newOrderWindown.client: cant get interfaceVal from listStore", "error", err)
			return false
		}
		if selected.(bool) {
			client, err = now.tools.StringFromIter(iter, &now.clientListStore.TreeModel, 1)
			if err != nil {
				logging.Logger.Error("newOrderWindown.client: cant get interfaceVal from listStore", "error", err)
				return false
			}
		}
		return false
	})
	return client
}

func (now *NewOrderWindow) paymentStatus() bool {
	return now.paymentStatusCheckBox.GetActive()
}

func (now *NewOrderWindow) paymentType() string {
	return now.paymentTypeEntry.GetLayout().GetText()
}

func (now *NewOrderWindow) cost() string {
	return now.costEntry.GetLayout().GetText()
}

func (now *NewOrderWindow) collectMarkedBlockAdvertisementValues(iter *gtk.TreeIter, list *gtk.ListStore) presenter.BlockAdvertisementDTO {
	id, err := now.tools.InterfaceFromIter(iter, &list.TreeModel, 1)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	orderID, err := now.tools.InterfaceFromIter(iter, &list.TreeModel, 2)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	releaseCount, err := now.tools.InterfaceFromIter(iter, &list.TreeModel, 3)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	releaseDates, err := now.tools.StringFromIter(iter, &list.TreeModel, 5)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	tags, err := now.tools.StringFromIter(iter, &list.TreeModel, 6)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	extraCharges, err := now.tools.StringFromIter(iter, &list.TreeModel, 7)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	size, err := now.tools.InterfaceFromIter(iter, &list.TreeModel, 8)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	cost, err := now.tools.StringFromIter(iter, &list.TreeModel, 9)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	comment, err := now.tools.StringFromIter(iter, &list.TreeModel, 10)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	fileName, err := now.tools.StringFromIter(iter, &list.TreeModel, 11)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}
	selectedBlock := presenter.BlockAdvertisementDTO{
		Advertisement: presenter.Advertisement{
			ID:           id.(int),
			OrderID:      orderID.(int),
			ReleaseCount: releaseCount.(int),
			ReleaseDates: releaseDates,
			Cost:         cost,
			Text:         comment,
			Tags:         tags,
			ExtraCharge:  extraCharges,
		},
		Size:     size.(int),
		FileName: fileName,
	}
	return selectedBlock
}

func (now *NewOrderWindow) CollectExistingBlockAdvertisements() []presenter.BlockAdvertisementDTO {
	blocks := make([]presenter.BlockAdvertisementDTO, 0)
	now.blockAdvListStore.ForEach(func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
		marked, err := now.tools.InterfaceFromIter(iter, model, 0)
		if err != nil {
			logging.Logger.Error("newOrderWindow.CollectExistingAdvertisements: Cannow get interface from listStore", "error", err)
			return false
		}
		if marked.(bool) {
			blocks = append(blocks, now.collectMarkedBlockAdvertisementValues(iter, now.blockAdvListStore))
		}
		return false
	})
	return blocks
}

func (now *NewOrderWindow) collectMarkedLineAdvertisementValues(iter *gtk.TreeIter, list *gtk.ListStore) presenter.LineAdvertisementDTO {
	id, err := now.tools.InterfaceFromIter(iter, &list.TreeModel, 1)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	orderID, err := now.tools.InterfaceFromIter(iter, &list.TreeModel, 2)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	releaseCount, err := now.tools.InterfaceFromIter(iter, &list.TreeModel, 3)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	releaseDates, err := now.tools.StringFromIter(iter, &list.TreeModel, 5)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	tags, err := now.tools.StringFromIter(iter, &list.TreeModel, 6)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	extraCharges, err := now.tools.StringFromIter(iter, &list.TreeModel, 7)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	cost, err := now.tools.StringFromIter(iter, &list.TreeModel, 8)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	comment, err := now.tools.StringFromIter(iter, &list.TreeModel, 9)
	if err != nil {
		logging.Logger.Error("blockAdvertisementsTab.SelectedRowChangedNew: Cannot get id value from treeModel", "error message", err)
	}

	selectedLine := presenter.LineAdvertisementDTO{
		Advertisement: presenter.Advertisement{
			ID:           id.(int),
			OrderID:      orderID.(int),
			ReleaseCount: releaseCount.(int),
			ReleaseDates: releaseDates,
			Cost:         cost,
			Text:         comment,
			Tags:         tags,
			ExtraCharge:  extraCharges,
		},
	}

	return selectedLine
}

func (now *NewOrderWindow) CollectExistingLineAdvertisements() []presenter.LineAdvertisementDTO {
	lines := make([]presenter.LineAdvertisementDTO, 0)
	now.lineAdvListStore.ForEach(func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
		marked, err := now.tools.InterfaceFromIter(iter, model, 0)
		if err != nil {
			logging.Logger.Error("newOrderWindow.CollectExistingAdvertisements: Cannow get interface from listStore", "error", err)
			return false
		}
		if marked.(bool) {
			lines = append(lines, now.collectMarkedLineAdvertisementValues(iter, now.lineAdvListStore))
		}
		return false
	})
	return lines
}
