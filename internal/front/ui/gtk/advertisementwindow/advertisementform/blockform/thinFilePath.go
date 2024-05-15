package blockform

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func (thn *ThinFilePath) Box() *gtk.Box {
	return thn.box
}

func (thn *ThinFilePath) FilePath() string {
	lab, _ := thn.filePathLink.GetLabel()
	return lab
}

func (thn *ThinFilePath) SetFilePath(path string) {
	uri, _ := thn.req.GetFileURI(path)
	thn.filePathLink.SetUri(uri)
	thn.filePathLink.SetLabel(path)
}

func (thn *ThinFilePath) viewClicked() {
	filechooser := thn.app.NewThinImageChooserWindow(thn)
	filechooser.Show()
	filechooser.LoadFiles()
}

func (thn *ThinFilePath) linkClicked() {
}

type ThinSignals struct {
	viewButtonPressed glib.SignalHandle
	linkClicked       glib.SignalHandle
}

func (thn *ThinFilePath) BindSignals() {
	thn.signals.viewButtonPressed = thn.filePathView.Connect("clicked", thn.viewClicked)
	thn.signals.linkClicked = thn.filePathLink.Connect("clicked", thn.linkClicked)
}
