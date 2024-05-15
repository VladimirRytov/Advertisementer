package liststores

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func (ls *ListStores) NewFilesList() {
	var err error
	ls.files, err = gtk.ListStoreNew(glib.TYPE_STRING, glib.TYPE_STRING, gdk.PixbufGetType())
	if err != nil {
		panic(err)
	}
}

func (ls *ListStores) FilesList() *gtk.ListStore {
	return ls.files
}

func (ls *ListStores) InsertFileFirstPlace(file presenter.File) {
	b, err := gdk.PixbufNewFromBytesOnly(file.Data)
	if err != nil {
		return
	}
	iter, ok := ls.files.GetIterFirst()
	if !ok {
		return
	}
	firstIter := ls.files.InsertBefore(iter)
	ls.files.Set(firstIter, []int{0, 1, 2}, []interface{}{file.Name, file.Size, b})
}

func (ls *ListStores) AppendFile(file presenter.File) {
	b, err := gdk.PixbufNewFromBytesOnly(file.Data)
	if err != nil {
		return
	}
	if ls.replaceMode {
		pos, err := ls.tools.FindValue(file.Name, ls.files, 0)
		if err == nil {
			ls.files.Set(pos, []int{0, 1, 2}, []interface{}{file.Name, file.Size, b})
			return
		}
	}
	ls.files.Set(ls.files.Append(), []int{0, 1, 2}, []interface{}{file.Name, file.Size, b})
}

func (ls *ListStores) RemoveFile(name string) {
	ls.tools.RemoveValue(name, 0, ls.files)

}
