package errorwindow

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/gtk"
)

type ErrorWindow struct {
	window *gtk.Window

	errorMessage *gtk.Label
	okButton     *gtk.Button
}

func (ew *ErrorWindow) Create() {
	ew.build()
	ew.bindSignals()
	ew.window.SetTitle("Ошибка")
}

func (ew *ErrorWindow) build() {
	buildFile, err := builder.NewBuilderFromString(builder.ErrorMessagesWindow)
	if err != nil {
		logging.Logger.Error("errorWindow: got panic")
		panic(err)
	}
	ew.window = buildFile.FetchWindow("ErrorWindowDialog")
	ew.errorMessage = buildFile.FetchLabel("ErrorLabel")
	ew.okButton = buildFile.FetchButton("OkButton")
}
