package tools

import (
	"errors"
	"unicode"

	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/gtk"
)

type Comparer interface {
	CompareStringTime(string, string) int
	CompareStrings(string, string) int
	CompareSelected(bool, bool) int
	CheckCostString(string) bool
	CheckAdvCostString(string) bool
	CompareCosts(string, string) int
}

type Tools struct {
	comparer Comparer
}

func NewTools(comp Comparer) *Tools {
	return &Tools{comparer: comp}
}

func (t *Tools) FindValue(id string, list *gtk.ListStore, column int) (*gtk.TreeIter, error) {
	var foundIter *gtk.TreeIter
	list.ForEach(
		func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
			value, err := list.GetValue(iter, column)
			if err != nil {
				logging.Logger.Error("an error occured while getting gVal from ListStore", "error", err.Error())
				return true
			}

			str, err := value.GetString()
			if err != nil {
				logging.Logger.Error("an error occured while converting gVal to string", "error", err.Error())
				return true
			}

			if str == id {
				foundIter = iter
				return true
			}
			return false
		})

	if foundIter != nil {
		return foundIter, nil
	}
	return nil, errors.New("value not found")
}

func (t *Tools) RemoveValue(id string, column int, list *gtk.ListStore) {
	pos, err := t.FindValue(id, list, column)
	if err != nil {
		logging.Logger.Error(err.Error())
		return
	}
	list.Remove(pos)
}

func (t *Tools) FindIntValue(id int, list *gtk.ListStore, column int) (*gtk.TreeIter, error) {
	var foundIter *gtk.TreeIter
	list.ForEach(
		func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
			value, err := list.GetValue(iter, column)
			if err != nil {
				logging.Logger.Error("an error occured while getting gVal from ListStore", "error", err.Error())
				return true
			}

			rawInt, err := value.GoValue()
			if err != nil {
				logging.Logger.Error("an error occured while converting gVal to string", "error", err.Error())
				return true
			}
			if rawInt.(int) == id {
				foundIter = iter
				return true
			}
			return false
		})

	if foundIter != nil {
		return foundIter, nil
	}
	return nil, errors.New("value not found")
}

func (t *Tools) RemoveIntValue(id int, column int, list *gtk.ListStore) {
	pos, err := t.FindIntValue(id, list, column)
	if err != nil {
		logging.Logger.Error(err.Error())
		return
	}
	list.Remove(pos)
}

func (t *Tools) CompareNewTime(model *gtk.TreeModel, a *gtk.TreeIter, newTime string, col int) int {
	valA, err := t.StringFromIter(a, model, col)
	if err != nil {
		logging.Logger.Error("tools.CompareTime: cannot get time A string", "error", err)
		return 0
	}

	return t.comparer.CompareStringTime(valA, newTime)
}

func (t *Tools) CompareTime(model *gtk.TreeModel, a, b *gtk.TreeIter, col int) int {
	valA, err := t.StringFromIter(a, model, col)
	if err != nil {
		logging.Logger.Error("tools.CompareTime: cannot get time A string", "error", err)
		return 0
	}
	valB, err := t.StringFromIter(b, model, col)
	if err != nil {
		logging.Logger.Error("tools.CompareTime: cannot get time B string", "error", err)
		return 0
	}

	return t.comparer.CompareStringTime(valA, valB)
}

func (t *Tools) CompareStrings(model *gtk.TreeModel, a, b *gtk.TreeIter, col int) int {
	valA, err := t.StringFromIter(a, model, col)
	if err != nil {
		logging.Logger.Error("tools.CompareStrings: cannot get string A from model", "error", err)
		return 0
	}
	valB, err := t.StringFromIter(b, model, col)
	if err != nil {
		logging.Logger.Error("tools.CompareStrings: cannot get string B from model", "error", err)
		return 0
	}
	return t.comparer.CompareStrings(valA, valB)
}

func (t *Tools) CompareBools(model *gtk.TreeModel, a, b *gtk.TreeIter, col int) int {
	gvalA, err := t.InterfaceFromIter(a, model, col)
	if err != nil {
		logging.Logger.Error("tools.CompareBools: cannot get value A from listStore", "error", err)
		return 0
	}
	gvalB, err := t.InterfaceFromIter(b, model, col)
	if err != nil {
		logging.Logger.Error("tools.CompareBools: cannot get value B from listStore", "error", err)
		return 0
	}
	selectedA, ok := gvalA.(bool)
	if !ok {
		logging.Logger.Error("cannot convert gvalA to bool")
		return 0
	}
	selectedB, ok := gvalB.(bool)
	if !ok {
		logging.Logger.Error("cannot convert gvalB to bool")
		return 0
	}

	return t.comparer.CompareSelected(selectedA, selectedB)
}

func (t *Tools) CompareInts(model *gtk.TreeModel, a, b *gtk.TreeIter, col int) int {
	intA, err := t.InterfaceFromIter(a, model, col)
	if err != nil {
		logging.Logger.Error("tools.CompareInts: cannot get strA string from list", "error", err)
		return 0
	}

	intB, err := t.InterfaceFromIter(b, model, col)
	if err != nil {
		logging.Logger.Error("tools.CompareInts: cannot get strB string from list", "error", err)
		return 0
	}

	return intA.(int) - intB.(int)
}

func (t *Tools) CompareCosts(model *gtk.TreeModel, a, b *gtk.TreeIter, col int) int {
	strA, err := t.StringFromIter(a, model, col)
	if err != nil {
		logging.Logger.Error("tools.CompareInts: cannot get strA string from list", "error", err)
		return 0
	}

	strB, err := t.StringFromIter(b, model, col)
	if err != nil {
		logging.Logger.Error("tools.CompareInts: cannot get strB string from list", "error", err)
		return 0
	}
	return t.comparer.CompareCosts(strA, strB)
}

func (t *Tools) StringFromIter(iter *gtk.TreeIter, model *gtk.TreeModel, col int) (string, error) {
	gval, err := model.GetValue(iter, col)
	if err != nil {
		logging.Logger.Error("stringFromIter: got error while getting vals from ListStore", "err", err)
		return "", err
	}
	str, err := gval.GetString()
	if err != nil {
		logging.Logger.Debug("stringFromIter: got error while converting glib val to string", "err", err)
		return "", err
	}
	return str, nil
}

func (t *Tools) InterfaceFromIter(iter *gtk.TreeIter, model *gtk.TreeModel, col int) (interface{}, error) {
	gval, err := model.GetValue(iter, col)
	if err != nil {
		logging.Logger.Error("interfaceFromIter: got error while getting vals from ListStore", "err", err)
		return nil, err
	}
	str, err := gval.GoValue()
	if err != nil {

		logging.Logger.Debug("interfaceFromIter: got error while converting glib val to interface", "err", err)
		return nil, err
	}
	return str, nil
}

func (t *Tools) CheckCostString(self *gtk.Entry, new string) {
	oldText := self.GetLayout().GetText()
	position := self.GetPosition()
	if !t.comparer.CheckCostString(oldText[:position] + new + oldText[position:]) {
		self.StopEmission("insert-text")
		return
	}
}

func (t *Tools) CheckCostAdvString(self *gtk.Entry, new string) {
	oldText := self.GetLayout().GetText()
	position := self.GetPosition()
	if !t.comparer.CheckAdvCostString(oldText[:position] + new + oldText[position:]) {
		self.StopEmission("insert-text")
		return
	}
}

func (t *Tools) CheckExtraChargeString(self *gtk.Entry, new string) {
	for _, v := range new {
		if !unicode.IsDigit(v) {
			self.StopEmission("insert-text")
			return
		}
	}
}
