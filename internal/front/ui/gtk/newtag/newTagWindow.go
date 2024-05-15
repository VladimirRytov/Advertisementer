package newtag

import "github.com/gotk3/gotk3/gtk"

func (ntw *NewTagWindow) TagName() string {
	return ntw.nameEntry.GetLayout().GetText()
}

func (ntw *NewTagWindow) TagCost() string {
	return ntw.costEntry.GetLayout().GetText()
}

func (ntw *NewTagWindow) Close() {
	ntw.window.Close()
}

func (ntw *NewTagWindow) Show() {
	ntw.window.Show()
}

func (ntw *NewTagWindow) Window() *gtk.Window {
	return ntw.window
}
