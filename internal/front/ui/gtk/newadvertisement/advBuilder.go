package newadvertisement

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"

	"github.com/gotk3/gotk3/gtk"
)

type AdvCreator interface {
	CreateBlockAdvertisement(*presenter.BlockAdvertisementDTO) error
	CreateLineAdvertisement(*presenter.LineAdvertisementDTO) error
}

type advForms interface {
	Widget() *gtk.Widget
	ToSelectedOrder()
	SetNewAdvMode(bool)
	SetSensetive(bool)
}

type LineForm interface {
	advForms
	FetchData() presenter.LineAdvertisementDTO
}

type BlockForm interface {
	advForms
	FetchData() presenter.BlockAdvertisementDTO
}

type ListStores interface {
	OrdersList() *gtk.ListStore
	TagsList() *gtk.ListStore
	ExtraChargesList() *gtk.ListStore
}

type Application interface {
	RefilterLists()
}

type AddAdvertisementWindow struct {
	req    AdvCreator
	app    Application
	window *gtk.Window

	advStackSwitcher *gtk.StackSwitcher
	advStack         *gtk.Stack
	blockPage        BlockForm
	linePage         LineForm

	createAdvButton *gtk.Button
}

func Create(reqGate AdvCreator, app Application, lineForm LineForm, blockForm BlockForm) *AddAdvertisementWindow {
	bldFile, err := builder.NewBuilderFromString(builder.AddAdvertisementWindow)
	if err != nil {
		panic(err)
	}
	adw := new(AddAdvertisementWindow)
	adw.req = reqGate
	adw.app = app
	adw.blockPage = blockForm
	adw.blockPage.SetNewAdvMode(true)

	adw.linePage = lineForm
	adw.linePage.SetNewAdvMode(true)
	adw.build(bldFile)
	adw.advStack.AddTitled(adw.blockPage.Widget(), "BlockAdvertisementStack", "Блочное объявление")
	adw.advStack.AddTitled(adw.linePage.Widget(), "LineAdvertisementStack", "Строковое объявление")
	adw.window.SetTitle("Создание объявлений")
	return adw
}

func (adw *AddAdvertisementWindow) build(buildFile *builder.Builder) {
	adw.window = buildFile.FetchWindow("AddAdvertisementWindow")
	adw.advStackSwitcher = buildFile.FetchStackSwitcher("AddAdvertisementStackSwitcher")
	adw.advStack = buildFile.FetchStack("AddAdvertisementStack")
	adw.createAdvButton = buildFile.FetchButton("AdvertisementAddButton")
	adw.bindSignals()
}
