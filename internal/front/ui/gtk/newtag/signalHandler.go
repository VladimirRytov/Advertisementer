package newtag

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"

	"github.com/gotk3/gotk3/glib"
)

type NewTagHandler struct {
	createButtonPressed glib.SignalHandle
	windowClosed        glib.SignalHandle
}

func (ntw *NewTagWindow) bindSignals() {
	ntw.signals.createButtonPressed = ntw.createButton.Connect("clicked", ntw.createButtonPressed)
	ntw.signals.windowClosed = ntw.window.Connect("destroy", ntw.windowClosed)
}

func (ntw *NewTagWindow) createButtonPressed() {
	tag := &presenter.TagDTO{
		TagName: ntw.TagName(),
		TagCost: ntw.TagCost(),
	}
	err := ntw.request.CreateTag(tag)
	if err != nil {
		ntw.errorLabel.SetText(err.Error())
		ntw.errorPopover.SetRelativeTo(ntw.costEntry)
		ntw.errorPopover.Popup()
	}
}

func (ntw *NewTagWindow) windowClosed() {
	ntw.window.Destroy()
}
