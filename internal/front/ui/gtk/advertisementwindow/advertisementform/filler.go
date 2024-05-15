package advertisementform

import (
	"errors"
	"strings"

	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/gtk"
)

type Tools interface {
	StringFromIter(*gtk.TreeIter, *gtk.TreeModel, int) (string, error)
	InterfaceFromIter(*gtk.TreeIter, *gtk.TreeModel, int) (interface{}, error)
}

type ListMarker struct {
	tools Tools
}

func Create(tools Tools) *ListMarker {
	lm := new(ListMarker)
	lm.tools = tools
	return lm
}

func (lm *ListMarker) UnMarkSelectedOrder(list *gtk.ListStore) {
	list.ForEach(
		func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
			list.SetValue(iter, 0, false)
			return false
		})
}

func (lm *ListMarker) MarkSelectedOrder(list *gtk.ListStore, orderID int) (*gtk.TreePath, error) {
	var selectedPath *gtk.TreePath
	list.ForEach(
		func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
			id, err := lm.tools.InterfaceFromIter(iter, &list.TreeModel, 1)
			if err != nil {
				logging.Logger.Error("markSelectedOrder: an error occured while getting order id from liststore", "error", err)
				return false
			}

			list.SetValue(iter, 0, id.(int) == orderID)
			if id == orderID {
				selectedPath, err = list.GetPath(iter)
				if err != nil {
					logging.Logger.Error("filler.markSelectedOrder: an error accured while getting path from listStore", "error", err)
				}
			}
			return false
		})
	if selectedPath != nil {
		return selectedPath, nil
	}

	return nil, errors.New("cant get path")
}

func (lm *ListMarker) MarkExtraCost(list *gtk.ListStore, costsName string) {
	marks := strings.Split(costsName, ", ")
	list.ForEach(
		func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
			name, err := lm.tools.StringFromIter(iter, &list.TreeModel, 1)
			if err != nil {
				logging.Logger.Error("filler.MarkExtraCost: an error occured while getting name from liststore", "error", err)
				return false
			}

			for _, v := range marks {
				if v == name {
					list.SetValue(iter, 0, true)
					break
				}
				list.SetValue(iter, 0, false)
			}
			return false
		})
}

func (lm *ListMarker) UnMarkExtraCost(list *gtk.ListStore) {
	list.ForEach(
		func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
			list.SetValue(iter, 0, false)
			return false
		})
}

func (lm *ListMarker) FillReleaseDates(list *gtk.ListStore, releaseDates string) {
	list.Clear()
	datesSlice := strings.Split(releaseDates, ", ")
	if len(releaseDates) == 0 {
		return
	}
	for _, v := range datesSlice {
		list.SetValue(list.Append(), 0, v)
	}
	list.SetSortColumnId(0, gtk.SORT_DESCENDING)
}
