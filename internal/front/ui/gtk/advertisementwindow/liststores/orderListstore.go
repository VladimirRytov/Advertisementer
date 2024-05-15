package liststores

import (
	"errors"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func (ls *ListStores) OrdersList() *gtk.ListStore {
	return ls.orders
}

func (ls *ListStores) newOrderList() (*gtk.ListStore, error) {
	orders, err := gtk.ListStoreNew(glib.TYPE_BOOLEAN, glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING,
		glib.TYPE_STRING, glib.TYPE_BOOLEAN)
	if err != nil {
		return nil, errors.Join(errors.New("listStores.ClientListCopy: cannot create listStore"), err)
	}
	return orders, nil
}

func (ls *ListStores) OrderListCopy() (*gtk.ListStore, error) {
	orders, err := gtk.ListStoreNew(glib.TYPE_BOOLEAN, glib.TYPE_INT, glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING,
		glib.TYPE_STRING, glib.TYPE_BOOLEAN)
	if err != nil {
		return nil, errors.Join(errors.New("listStores.ClientListCopy: cannot create listStore"), err)
	}
	ls.orders.ForEach(func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
		id, err := ls.tools.InterfaceFromIter(iter, &ls.orders.TreeModel, 1)
		if err != nil {
			logging.Logger.Error("listStores.OrderListCopy: cannt get id from listStore", "error", err)
			return false
		}
		clientName, err := ls.tools.StringFromIter(iter, &ls.orders.TreeModel, 2)
		if err != nil {
			logging.Logger.Error("listStores.OrderListCopy: cannt get id from listStore", "error", err)
			return false
		}
		createdDate, err := ls.tools.StringFromIter(iter, &ls.orders.TreeModel, 3)
		if err != nil {
			logging.Logger.Error("listStores.OrderListCopy: cannt get id from listStore", "error", err)
			return false
		}
		cost, err := ls.tools.StringFromIter(iter, &ls.orders.TreeModel, 4)
		if err != nil {
			logging.Logger.Error("listStores.OrderListCopy: cannt get id from listStore", "error", err)
			return false
		}
		paymentType, err := ls.tools.StringFromIter(iter, &ls.orders.TreeModel, 5)
		if err != nil {
			logging.Logger.Error("listStores.OrderListCopy: cannt get id from listStore", "error", err)
			return false
		}
		paymentStatus, err := ls.tools.InterfaceFromIter(iter, &ls.orders.TreeModel, 6)
		if err != nil {
			logging.Logger.Error("listStores.OrderListCopy: cannt get id from listStore", "error", err)
			return false
		}
		orders.Set(orders.Append(), []int{0, 1, 2, 3, 4, 5, 6}, []interface{}{false, id.(int), clientName, createdDate, cost, paymentType, paymentStatus.(bool)})
		return false
	})
	return orders, nil
}

func (ls *ListStores) AppendOrder(order presenter.OrderDTO) {
	logging.Logger.Debug("Add order to listStore", "Order ID", order.ID)
	if ls.replaceMode {
		pos, err := ls.tools.FindIntValue(order.ID, ls.orders, 1)
		if err == nil {
			ls.orders.Set(pos, []int{0, 1, 2, 3, 4, 5, 6},
				[]interface{}{false, order.ID, order.ClientName, order.CreatedDate, order.Cost, order.PaymentType, order.PaymentStatus})
			return
		}
	}
	ls.orders.Set(ls.orders.Append(), []int{0, 1, 2, 3, 4, 5, 6},
		[]interface{}{false, order.ID, order.ClientName, order.CreatedDate, order.Cost, order.PaymentType, order.PaymentStatus})
}

func (ls *ListStores) RemoveOrder(id int) {
	ls.tools.RemoveIntValue(id, 1, ls.orders)
	ls.RemoveBlockAdvertisementByOrderID(id)
	ls.RemoveLineAdvertisementByOrderID(id)
}

func (ls *ListStores) RemoveOrderByClientID(clientId string) {
	for {
		iter, err := ls.tools.FindValue(clientId, ls.orders, 2)
		if err != nil {
			break
		}
		id, err := ls.tools.InterfaceFromIter(iter, &ls.orders.TreeModel, 1)
		if err != nil {
			logging.Logger.Warn("listStores.RemoveOrderByClientID: order not found")
			break
		}
		ls.RemoveOrder(id.(int))
	}
}

func (ls *ListStores) ClearOrderList() {
	ls.orders.Clear()
}
