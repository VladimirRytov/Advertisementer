package progress

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type SignalHandler struct {
	onDeleteEvent      glib.SignalHandle
	rowAdded           glib.SignalHandle
	closeButtonClicked glib.SignalHandle
}

func (pw *ProgressWindow) bindSignals() {
	pw.signals.closeButtonClicked = pw.closeButton.Connect("clicked", pw.closeButtonClicked)
	pw.signals.rowAdded = pw.logList.ConnectAfter("row-changed", pw.rowAdded)
}

func (pw *ProgressWindow) rowAdded(self *gtk.ListStore, path *gtk.TreePath, iter *gtk.TreeIter) {
	msg, err := pw.tools.StringFromIter(iter, self.ToTreeModel(), 0)
	if err != nil {
		return
	}
	pw.currentProgress.SetText(msg)
}

func (pw *ProgressWindow) closeButtonClicked() {
	pw.enableDeletable()
	if pw.cancelProgress != nil {
		pw.cancelProgress()
	}
	pw.logList.Clear()
	pw.window.Close()
}

func (pw *ProgressWindow) blockdeleteEvent() bool {
	return true
}

func (pw *ProgressWindow) unblockdeleteEvent() bool {
	return false
}
