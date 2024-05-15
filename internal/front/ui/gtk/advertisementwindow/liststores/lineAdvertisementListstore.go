package liststores

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func (ls *ListStores) newLineAdvertisementsList() (*gtk.ListStore, error) {
	lineAdv, err := gtk.ListStoreNew(glib.TYPE_BOOLEAN, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_STRING,
		glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING)
	if err != nil {
		logging.Logger.Error("listStores.NewBlockAdvertisementsList: cannot create listStore")
		return nil, err
	}
	return lineAdv, nil
}

func (ls *ListStores) LineAdvertisementsList() *gtk.ListStore {
	return ls.lineAdvertisements
}

func (ls *ListStores) AppendLineAdvertisement(lineAdv presenter.LineAdvertisementDTO) {
	logging.Logger.Debug("Add lineAdvertisement to listStore", "lineAdvertisement ID", lineAdv.ID)
	if ls.replaceMode {
		pos, err := ls.tools.FindIntValue(lineAdv.ID, ls.lineAdvertisements, 1)
		if err == nil {
			ls.lineAdvertisements.Set(pos, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				[]interface{}{false, lineAdv.ID, lineAdv.OrderID, lineAdv.ReleaseCount, lineAdv.ClosestRelease,
					lineAdv.ReleaseDates, lineAdv.Tags, lineAdv.ExtraCharge, lineAdv.Cost, lineAdv.Text})
			return
		}
	}
	err := ls.lineAdvertisements.Set(ls.lineAdvertisements.Append(), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]interface{}{false, lineAdv.ID, lineAdv.OrderID, lineAdv.ReleaseCount, lineAdv.ClosestRelease,
			lineAdv.ReleaseDates, lineAdv.Tags, lineAdv.ExtraCharge, lineAdv.Cost, lineAdv.Text})
	if err != nil {
		logging.Logger.Error("ListStores.AppendLineAdvertisement", err)
	}
}

func (ls *ListStores) RemoveLineAdvertisement(id int) {
	ls.tools.RemoveIntValue(id, 1, ls.lineAdvertisements)
}

func (ls *ListStores) RemoveLineAdvertisementByOrderID(id int) {
	for {
		_, err := ls.tools.FindIntValue(id, ls.lineAdvertisements, 2)
		if err != nil {
			break
		}
		ls.tools.RemoveIntValue(id, 2, ls.lineAdvertisements)
	}
}

func (ls *ListStores) ClearLineAdvertisementList() {
	ls.lineAdvertisements.Clear()
}
