package liststores

import (
	"errors"

	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func (ls *ListStores) ReplaceMode() bool {
	return ls.replaceMode
}

func (ls *ListStores) SetReplaceMode(mode bool) {
	ls.replaceMode = mode
}

func (ls *ListStores) MessageList() *gtk.ListStore {
	return ls.messageList
}

func (ls *ListStores) newMessageList() (*gtk.ListStore, error) {
	errList, err := gtk.ListStoreNew(glib.TYPE_STRING)
	if err != nil {
		return nil, errors.Join(errors.New("tagsListCopy: cannot create tag listStore"), err)
	}
	return errList, nil
}

func (ls *ListStores) AppendMessage(msg string) {
	logging.Logger.Debug("Add message to listStore", "Message", msg)
	err := ls.messageList.SetValue(ls.messageList.Append(), 0, msg)
	if err != nil {
		logging.Logger.Error("ListStores.AppendError", err)
	}
}

func (ls *ListStores) ClearMessageList() {
	ls.messageList.Clear()
}
