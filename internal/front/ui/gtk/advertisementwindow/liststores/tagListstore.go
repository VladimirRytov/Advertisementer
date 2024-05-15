package liststores

import (
	"errors"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func (ls *ListStores) TagsList() *gtk.ListStore {
	return ls.tags
}

func (ls *ListStores) newTagsList() (*gtk.ListStore, error) {
	tag, err := gtk.ListStoreNew(glib.TYPE_BOOLEAN, glib.TYPE_STRING, glib.TYPE_STRING)
	if err != nil {
		return nil, errors.Join(errors.New("tagsListCopy: cannot create tag listStore"), err)
	}
	return tag, nil
}

func (ls *ListStores) TagsListCopy() (*gtk.ListStore, error) {
	tag, err := gtk.ListStoreNew(glib.TYPE_BOOLEAN, glib.TYPE_STRING, glib.TYPE_STRING)
	if err != nil {
		return nil, errors.Join(errors.New("tagsListCopy: cannot create tag listStore"), err)
	}
	ls.tags.ForEach(func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
		tagname, err := ls.tools.StringFromIter(iter, &ls.tags.TreeModel, 1)
		if err != nil {
			logging.Logger.Error("listStores.TagsListCopy: cannt get tag name from listStore", "error", err)
			return false
		}
		tagCost, err := ls.tools.StringFromIter(iter, &ls.tags.TreeModel, 2)
		if err != nil {
			logging.Logger.Error("listStores.TagsListCopy: cannt get tag cost from listStore", "error", err)
			return false
		}
		tag.Set(tag.Append(), []int{0, 1, 2}, []interface{}{false, tagname, tagCost})
		return false
	})
	return tag, nil
}

func (ls *ListStores) AppendTag(tag presenter.TagDTO) {
	logging.Logger.Debug("Add tag to listStore", "Tag name", tag.TagName)
	if ls.replaceMode {
		pos, err := ls.tools.FindValue(tag.TagName, ls.tags, 1)
		if err == nil {
			ls.tags.Set(pos, []int{0, 1, 2}, []interface{}{false, tag.TagName, tag.TagCost})
			return
		}
	}
	ls.tags.Set(ls.tags.Append(), []int{0, 1, 2}, []interface{}{false, tag.TagName, tag.TagCost})
}

func (ls *ListStores) RemoveTag(tagName string) {
	ls.tools.RemoveValue(tagName, 1, ls.tags)
}
func (ls *ListStores) ClearTagList() {
	ls.tags.Clear()
}
