package errorwindow

import "github.com/VladimirRytov/advertisementer/internal/logging"

type ErrorWindowHandler struct {
	*ErrorWindow
}

func (ew *ErrorWindow) bindSignals() {
	handler := &ErrorWindowHandler{ew}
	ew.okButton.Connect("clicked", handler.okButtonClicked)
	ew.window.Connect("destroy", ew.window.Destroy)
}

func (ewh *ErrorWindowHandler) okButtonClicked() {
	logging.Logger.Debug("okButtonClicked")
	ewh.window.Close()
}
