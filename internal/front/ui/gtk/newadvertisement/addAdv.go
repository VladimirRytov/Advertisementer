package newadvertisement

import "github.com/gotk3/gotk3/gtk"

func (adw *AddAdvertisementWindow) Window() *gtk.Window {
	return adw.window
}

func (adw *AddAdvertisementWindow) Show() {
	adw.window.Show()
}

func (adw *AddAdvertisementWindow) Close() {
	adw.window.Close()
}

func (adw *AddAdvertisementWindow) SelectedPage() string {
	return adw.advStack.GetVisibleChildName()
}
