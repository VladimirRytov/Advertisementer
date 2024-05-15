package newclient

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"

	"github.com/gotk3/gotk3/gtk"
)

type ClientCreator interface {
	CreateClient(*presenter.ClientDTO) error
}

type NewClientWindow struct {
	req ClientCreator

	window *gtk.Window

	nameLabel *gtk.Label
	nameEntry *gtk.Entry

	phoneLabel *gtk.Label
	phoneEntry *gtk.Entry

	emailLabel *gtk.Label
	emailEntry *gtk.Entry

	additionalInformationLabel      *gtk.Label
	additionalInformationTextBuffer *gtk.TextBuffer
	additionalInformationTextView   *gtk.TextView

	createButton *gtk.Button
}

func Create(req ClientCreator) *NewClientWindow {
	nc := new(NewClientWindow)
	build, err := builder.NewBuilderFromString(builder.AddClientWindow)
	if err != nil {
		panic(err)
	}
	nc.req = req
	nc.build(build)
	nc.bindSignals()
	nc.window.SetTitle("Создание контрагента")
	return nc
}

func (nc *NewClientWindow) build(buildFile *builder.Builder) {
	nc.window = buildFile.FetchWindow("AddClientWindow")
	nc.nameLabel = buildFile.FetchLabel("ClientNameLabel")
	nc.nameEntry = buildFile.FetchEntry("ClientNameEntry")
	nc.phoneLabel = buildFile.FetchLabel("ClientContactNumbersLabel")
	nc.phoneEntry = buildFile.FetchEntry("ClientContactNumbersEntry")
	nc.emailLabel = buildFile.FetchLabel("ClientEmailLabel")
	nc.emailEntry = buildFile.FetchEntry("ClientEmailEntry")
	nc.additionalInformationLabel = buildFile.FetchLabel("ClientAdditionalInformationLabel")
	nc.additionalInformationTextBuffer = buildFile.FetchTextBuffer("ClientTextBuffer")
	nc.additionalInformationTextView = buildFile.FetchTextView("ClientAdditionalInformationTextView")
	nc.createButton = buildFile.FetchButton("AddClientButton")
}
