package liststores

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func (ls *ListStores) newBlockAdvertisementsList() (*gtk.ListStore, error) {
	blockAdv, err := gtk.ListStoreNew(glib.TYPE_BOOLEAN, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_STRING,
		glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING)
	if err != nil {
		logging.Logger.Error("listStores.NewBlockAdvertisementsList: cannot create listStore")
		return nil, err
	}
	return blockAdv, nil
}

func (ls *ListStores) BlockAdvertisementsList() *gtk.ListStore {
	return ls.blockAdvertisements
}

func (ls *ListStores) AppendBlockAdvertisement(blockAdv presenter.BlockAdvertisementDTO) {
	logging.Logger.Debug("Add blockAdvertisement to listStore", "blockAdvertisement", blockAdv)
	if ls.replaceMode {
		pos, err := ls.tools.FindIntValue(blockAdv.ID, ls.blockAdvertisements, 1)
		if err == nil {
			logging.Logger.Debug("ListStores.AppendBlockAdvertisement: replacing Block advertisement", "id", blockAdv.ID)
			err := ls.blockAdvertisements.Set(pos, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
				[]interface{}{false, blockAdv.ID, blockAdv.OrderID, blockAdv.ReleaseCount, blockAdv.ClosestRelease,
					blockAdv.ReleaseDates, blockAdv.Tags, blockAdv.ExtraCharge, blockAdv.Size, blockAdv.Cost, blockAdv.Text, blockAdv.FileName})
			if err != nil {
				logging.Logger.Error("ListStores.AppendBlockAdvertisement: cannot set data to liststore", "error", err)
			}
			return
		}
	}
	logging.Logger.Debug("ListStores.AppendBlockAdvertisement: append Block advertisement", "id", blockAdv.ID)
	ls.blockAdvertisements.Set(ls.blockAdvertisements.Append(), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		[]interface{}{false, blockAdv.ID, blockAdv.OrderID, blockAdv.ReleaseCount, blockAdv.ClosestRelease,
			blockAdv.ReleaseDates, blockAdv.Tags, blockAdv.ExtraCharge, blockAdv.Size, blockAdv.Cost, blockAdv.Text, blockAdv.FileName})

}

func (ls *ListStores) RemoveBlockAdvertisement(id int) {
	ls.tools.RemoveIntValue(id, 1, ls.blockAdvertisements)
}

func (ls *ListStores) RemoveBlockAdvertisementByOrderID(id int) {
	for {
		_, err := ls.tools.FindIntValue(id, ls.blockAdvertisements, 2)
		if err != nil {
			break
		}
		ls.tools.RemoveIntValue(id, 2, ls.blockAdvertisements)
	}
}

func (ls *ListStores) ClearBlockAdvertisementList() {
	ls.blockAdvertisements.Clear()
}
