package progress

import (
	"context"

	"github.com/gotk3/gotk3/gtk"
)

func (pw *ProgressWindow) Window() *gtk.Window {
	return pw.window
}

func (pw *ProgressWindow) Show() {
	if pw.beforeFunc != nil {
		pw.beforeFunc()
	}
	pw.window.Show()
}

func (pw *ProgressWindow) Hide() {
	pw.window.Hide()
}

func (pw *ProgressWindow) Close() {
	pw.window.Destroy()
	pw.messages.ClearMessageList()
}

func (pw *ProgressWindow) disableDeletable() {
	pw.signals.onDeleteEvent = pw.window.Connect("delete-event", pw.blockdeleteEvent)
}

func (pw *ProgressWindow) enableDeletable() {
	pw.window.HandlerDisconnect(pw.signals.onDeleteEvent)
	pw.signals.onDeleteEvent = pw.window.Connect("delete-event", pw.unblockdeleteEvent)
}

func (pw *ProgressWindow) SetMessage(message string) {
	pw.progresLabel.SetText(message)
}

func (pw *ProgressWindow) SetCancelFunc(f context.CancelFunc) {
	pw.cancelProgress = f
}

func (pw *ProgressWindow) BlockCalcelButton(b bool) {
	pw.closeButton.SetSensitive(!b)
}

func (pw *ProgressWindow) ShowTreeview(b bool) {
	pw.logTreeView.SetVisible(b)
}

func (pw *ProgressWindow) SetCloseButtonMessage(mas string) {
	pw.closeButton.SetLabel(mas)
}

func (pw *ProgressWindow) SetBeforeFunc(f func()) {
	pw.beforeFunc = f
}

func (pw *ProgressWindow) SetAfterFunc(f func()) {
	pw.afterFunc = f
}

func (pw *ProgressWindow) ProgressDone(message string) {
	pw.enableDeletable()
	pw.currentProgress.SetVisible(false)
	pw.progresLabel.SetText(message)
	if pw.afterFunc != nil {
		pw.afterFunc()
	}
	pw.SetCloseButtonMessage("Закрыть")
	pw.BlockCalcelButton(false)
}
