package liststores

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func (ls *ListStores) costRatesNewListStore() (*gtk.ListStore, error) {
	costRate, err := gtk.ListStoreNew(glib.TYPE_BOOLEAN, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_BOOLEAN, glib.TYPE_BOOLEAN)
	if err != nil {
		return nil, err
	}
	return costRate, nil
}

func (ls *ListStores) AppendCostRate(costRate presenter.CostRateDTO) {
	logging.Logger.Debug("Add CostRate to listStore", "CostRate Name", costRate.Name)
	if ls.replaceMode {
		pos, err := ls.tools.FindValue(costRate.Name, ls.costRates, 1)
		if err == nil {
			ls.costRates.Set(pos, []int{0, 1, 2, 3, 4, 5},
				[]interface{}{false, costRate.Name, costRate.Onecm2, costRate.OneWordOrSymbol, costRate.CalcForOneWord, !costRate.CalcForOneWord})
			return
		}
	}

	err := ls.costRates.Set(ls.costRates.Append(), []int{0, 1, 2, 3, 4, 5},
		[]interface{}{false, costRate.Name, costRate.Onecm2, costRate.OneWordOrSymbol, costRate.CalcForOneWord, !costRate.CalcForOneWord})
	if err != nil {
		logging.Logger.Error("ListStores.AppendCostRate", err)
	}
}

func (ls *ListStores) ClearCostRateListStore() {
	ls.costRates.Clear()
}

func (ls *ListStores) RemoveCostRate(name string) {
	ls.tools.RemoveValue(name, 1, ls.costRates)
}

func (ls *ListStores) CostRatesListStore() *gtk.ListStore {
	return ls.costRates
}
