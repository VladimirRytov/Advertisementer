package newtag

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"

	"github.com/gotk3/gotk3/gtk"
)

type TagCreator interface {
	CreateTag(*presenter.TagDTO) error
}

type NewTagWindow struct {
	request TagCreator
	window  *gtk.Window

	nameLabel *gtk.Label
	nameEntry *gtk.Entry
	costLabel *gtk.Label
	costEntry *gtk.Entry

	createButton *gtk.Button

	errorPopover *gtk.Popover
	errorLabel   *gtk.Label
	signals      NewTagHandler
}

func Create(req TagCreator) *NewTagWindow {
	ntw := new(NewTagWindow)
	buildFile, err := builder.NewBuilderFromString(builder.AddTagWindow)
	if err != nil {
		panic(err)
	}
	ntw.request = req
	ntw.build(buildFile)
	ntw.costEntry.SetPlaceholderText("0.00")
	ntw.bindSignals()
	ntw.window.SetTitle("Создание тэга")
	return ntw
}

func (ntw *NewTagWindow) build(buildFile *builder.Builder) {
	ntw.window = buildFile.FetchWindow("AddTagWindow")
	ntw.nameLabel = buildFile.FetchLabel("TagNameLabel")
	ntw.nameEntry = buildFile.FetchEntry("TagNameEntry")
	ntw.costLabel = buildFile.FetchLabel("TagCostLabel")
	ntw.costEntry = buildFile.FetchEntry("TagCostEntry")
	ntw.createButton = buildFile.FetchButton("TagAddButton")
	ntw.errorPopover = buildFile.FetchPopover("ErrorPopover")
	ntw.errorLabel = buildFile.FetchLabel("ErrorLabel")
}
