package liststores

import (
	"errors"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func (ls *ListStores) ExtraChargesList() *gtk.ListStore {
	return ls.extraCharge
}

func (ls *ListStores) AppendExtraCharge(charge presenter.ExtraChargeDTO) {
	logging.Logger.Debug("Add extraCharge to listStore", "ExtraCharge name", charge.ChargeName)
	if ls.replaceMode {
		pos, err := ls.tools.FindValue(charge.ChargeName, ls.extraCharge, 1)
		if err == nil {
			ls.extraCharge.Set(pos, []int{0, 1, 2}, []interface{}{false, charge.ChargeName, charge.Multiplier})
			return
		}
	}
	ls.extraCharge.Set(ls.extraCharge.Append(), []int{0, 1, 2}, []interface{}{false, charge.ChargeName, charge.Multiplier})

}

func (ls *ListStores) newExtraChargeList() (*gtk.ListStore, error) {
	charges, err := gtk.ListStoreNew(glib.TYPE_BOOLEAN, glib.TYPE_STRING, glib.TYPE_STRING)
	if err != nil {
		return nil, errors.Join(errors.New("listStores.ExtraChargeListCopy: cannot create listStore"), err)
	}
	return charges, nil
}

func (ls *ListStores) ExtraChargeListCopy() (*gtk.ListStore, error) {
	charges, err := gtk.ListStoreNew(glib.TYPE_BOOLEAN, glib.TYPE_STRING, glib.TYPE_STRING)
	if err != nil {
		return nil, errors.Join(errors.New("listStores.ExtraChargeListCopy: cannot create listStore"), err)
	}
	ls.extraCharge.ForEach(func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
		chargeName, err := ls.tools.StringFromIter(iter, &ls.extraCharge.TreeModel, 1)
		if err != nil {
			logging.Logger.Error("listStores.ExtraChargeListCopy: cannt get name from listStore", "error", err)
			return false
		}
		miltiplier, err := ls.tools.StringFromIter(iter, &ls.extraCharge.TreeModel, 2)
		if err != nil {
			logging.Logger.Error("listStores.ExtraChargeListCopy: cannt get multiplier from listStore", "error", err)
			return false
		}
		charges.Set(charges.Append(), []int{0, 1, 2}, []interface{}{false, chargeName, miltiplier})
		return false
	})
	return charges, nil
}

func (ls *ListStores) RemoveExtraCharge(chargeName string) {
	ls.tools.RemoveValue(chargeName, 1, ls.extraCharge)
}

func (ls *ListStores) ClearExtraChargeList() {
	ls.extraCharge.Clear()
}
