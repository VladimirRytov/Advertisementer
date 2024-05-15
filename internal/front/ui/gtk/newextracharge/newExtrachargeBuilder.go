package newextracharge

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"

	"github.com/gotk3/gotk3/gtk"
)

type ExtraChargeCreator interface {
	CreateExtraCharge(*presenter.ExtraChargeDTO) error
}

type Tools interface {
	CheckExtraChargeString(*gtk.Entry, string)
}

type NewExtrachargeWindow struct {
	req    ExtraChargeCreator
	tools  Tools
	window *gtk.Window

	nameLabel       *gtk.Label
	nameEntry       *gtk.Entry
	multiplierLabel *gtk.Label
	multiplierEntry *gtk.Entry

	createButton *gtk.Button

	errorPopover *gtk.Popover
	errorLabel   *gtk.Label
}

func Create(req ExtraChargeCreator, tools Tools) *NewExtrachargeWindow {
	nexw := new(NewExtrachargeWindow)
	buildFile, err := builder.NewBuilderFromString(builder.AddExtraChargeWindow)
	if err != nil {
		panic(err)
	}
	nexw.req = req
	nexw.tools = tools
	nexw.build(buildFile)
	nexw.bindSignals()
	nexw.multiplierEntry.SetPlaceholderText("0")
	nexw.window.SetTitle("Создание наценки")
	return nexw
}

func (nexw *NewExtrachargeWindow) build(buildFile *builder.Builder) {
	nexw.window = buildFile.FetchWindow("AddExtraChargeWindow")
	nexw.nameLabel = buildFile.FetchLabel("ExtraChargeNameLabel")
	nexw.nameEntry = buildFile.FetchEntry("ExtraChargeNameEntry")
	nexw.multiplierLabel = buildFile.FetchLabel("ExtraChargeMultiplierLabel")
	nexw.multiplierEntry = buildFile.FetchEntry("ExtraChargeMultiplierEntry")
	nexw.createButton = buildFile.FetchButton("ExtraChargeAddButton")

	nexw.errorPopover = buildFile.FetchPopover("ErrorPopover")
	nexw.errorLabel = buildFile.FetchLabel("ErrorLabel")
}
