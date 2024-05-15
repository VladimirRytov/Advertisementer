package newcostratewindow

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"

	"github.com/gotk3/gotk3/gtk"
)

type CostRateCreator interface {
	CreateCostRate(*presenter.CostRateDTO) error
	UpdateCostRate(*presenter.CostRateDTO) error
}

type Tools interface {
	CheckCostString(*gtk.Entry, string)
}

type NewCostRateWindow struct {
	req   CostRateCreator
	tools Tools

	signals signalHandler
	window  *gtk.Window

	newCostRateLabel *gtk.Label

	costNameLabel *gtk.Label
	costNameEntry *gtk.Entry

	costOneSquareLabel *gtk.Label
	costOneSquareEntry *gtk.Entry

	costForWordSymbolLabel *gtk.Label
	costForWordSymbolEntry *gtk.Entry

	costForOneWordRadioButton *gtk.RadioButton
	costForSymbolRadioButton  *gtk.RadioButton

	createButton *gtk.Button
}

func Create(updateMode bool, req CostRateCreator, tools Tools) *NewCostRateWindow {
	buildfile, err := builder.NewBuilderFromString(builder.CostRateWindow)
	if err != nil {
		panic(err)
	}
	ncs := new(NewCostRateWindow)
	ncs.req = req
	ncs.tools = tools
	ncs.build(*buildfile)
	ncs.costForWordSymbolEntry.SetPlaceholderText("0,00")
	ncs.costOneSquareEntry.SetPlaceholderText("0,00")
	ncs.window.SetTitle("Создание тарифа")
	if updateMode {
		ncs.costNameEntry.SetEditable(false)
		ncs.window.SetTitle("Обновление тарифа")
		ncs.newCostRateLabel.SetText("Обновление тарифа")
		ncs.createButton.SetLabel("Обновить")
	}
	ncs.bindSignals(updateMode)
	return ncs
}

func (ncs *NewCostRateWindow) build(buildFile builder.Builder) {
	ncs.window = buildFile.FetchWindow("NewCostRateWindow")
	ncs.newCostRateLabel = buildFile.FetchLabel("NewCostRateLabel")
	ncs.costNameLabel = buildFile.FetchLabel("CostNameLabel")
	ncs.costNameEntry = buildFile.FetchEntry("CostNameEntry")
	ncs.costOneSquareLabel = buildFile.FetchLabel("CostOneSquareLabel")
	ncs.costOneSquareEntry = buildFile.FetchEntry("CostOneSquareEntry")
	ncs.costForWordSymbolLabel = buildFile.FetchLabel("CostForWordSymbolLabel")
	ncs.costForWordSymbolEntry = buildFile.FetchEntry("CostForWordSymbolEntry")
	ncs.costForOneWordRadioButton = buildFile.FetchRadioButton("CostForOneWordRadioButton")
	ncs.costForSymbolRadioButton = buildFile.FetchRadioButton("CostForSymbolRadioButton")
	ncs.createButton = buildFile.FetchButton("CreateButton")
}
