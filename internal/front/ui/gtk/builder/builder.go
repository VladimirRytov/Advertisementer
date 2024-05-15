package builder

import (
	"embed"

	"github.com/gotk3/gotk3/gtk"
)

//go:embed resources/*
var files embed.FS

const (
	AddAdvertisementWindow = "resources/AddAdvertisementWindow.glade"
	AddClientWindow        = "resources/AddClientWindow.glade"
	AddExtraChargeWindow   = "resources/AddExtraChargeWindow.glade"
	AddOrderWindow         = "resources/AddOrderWindow.glade"
	AddTagWindow           = "resources/AddTagWindow.glade"
	AdvertisementForm      = "resources/AdvertisementForm.glade"
	AdvReport              = "resources/AdvReport.glade"
	CostRateWindow         = "resources/CostRateWindow.glade"
	ErrorMessagesWindow    = "resources/ErrorMessagesWindow.glade"
	ErrorPopover           = "resources/ErrorPopover.glade"
	Files                  = "resources/Files.glade"
	ImportJSONWindow       = "resources/ImportJSONWindow.glade"
	Journals               = "resources/Journals.glade"
	LoginWindow            = "resources/LoginWindow.glade"
	MainWindow             = "resources/MainWindow.glade"
	Progress               = "resources/Progress.glade"
	SettingsWindow         = "resources/SettingsWindow.glade"
	ToJSOnWindow           = "resources/ToJSOnWindow.glade"
)

type Builder struct {
	Builder *gtk.Builder
}

func NewBuilderFromFile(filePath string) (*Builder, error) {
	b, err := gtk.BuilderNewFromFile(filePath)
	if err != nil {
		return nil, err
	}
	return &Builder{Builder: b}, nil
}

func NewBuilderFromString(name string) (*Builder, error) {
	file, err := files.ReadFile(name)
	if err != nil {
		return nil, err
	}
	b, err := gtk.BuilderNewFromString(string(file))
	if err != nil {
		return nil, err
	}
	return &Builder{Builder: b}, nil
}

func (b *Builder) FetchButton(obj string) *gtk.Button {
	buttonObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return buttonObj.(*gtk.Button)
}

func (b *Builder) FetchLabel(obj string) *gtk.Label {
	labelObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return labelObj.(*gtk.Label)
}

func (b *Builder) FetchEntry(obj string) *gtk.Entry {
	entryObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return entryObj.(*gtk.Entry)
}

func (b *Builder) FetchStack(obj string) *gtk.Stack {
	stackObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return stackObj.(*gtk.Stack)
}

func (b *Builder) FetchStackSwitcher(obj string) *gtk.StackSwitcher {
	stackSwitcherObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return stackSwitcherObj.(*gtk.StackSwitcher)
}

func (b *Builder) FetchComboBox(obj string) *gtk.ComboBox {
	comboBoxObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return comboBoxObj.(*gtk.ComboBox)
}

func (b *Builder) FetchCheckButton(obj string) *gtk.CheckButton {
	checkButtonObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return checkButtonObj.(*gtk.CheckButton)
}

func (b *Builder) FetchRevealer(obj string) *gtk.Revealer {
	revealerObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return revealerObj.(*gtk.Revealer)
}

func (b *Builder) FetchListStore(obj string) *gtk.ListStore {
	listStoreObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return listStoreObj.(*gtk.ListStore)
}

func (b *Builder) FetchTreeView(obj string) *gtk.TreeView {
	treeViewObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return treeViewObj.(*gtk.TreeView)
}

func (b *Builder) FetchNoteBook(obj string) *gtk.Notebook {
	noteBookObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return noteBookObj.(*gtk.Notebook)
}

func (b *Builder) FetchRadioButton(obj string) *gtk.RadioButton {
	radioButtonkObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return radioButtonkObj.(*gtk.RadioButton)
}

func (b *Builder) FetchCalendar(obj string) *gtk.Calendar {
	calendarObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return calendarObj.(*gtk.Calendar)
}

func (b *Builder) FetchTextView(obj string) *gtk.TextView {
	textViewObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return textViewObj.(*gtk.TextView)
}

func (b *Builder) FetchMenuButton(obj string) *gtk.MenuButton {
	menuButtonObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return menuButtonObj.(*gtk.MenuButton)
}

func (b *Builder) FetchSpinButton(obj string) *gtk.SpinButton {
	spinButtonObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return spinButtonObj.(*gtk.SpinButton)
}

func (b *Builder) FetchSearchEntry(obj string) *gtk.SearchEntry {
	searchEntryObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return searchEntryObj.(*gtk.SearchEntry)
}

func (b *Builder) FetchPopover(obj string) *gtk.Popover {
	popoverObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return popoverObj.(*gtk.Popover)
}

func (b *Builder) FetchPopoverMenu(obj string) *gtk.PopoverMenu {
	popoverMenuObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return popoverMenuObj.(*gtk.PopoverMenu)
}

func (b *Builder) FetchImage(obj string) *gtk.Image {
	imageObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return imageObj.(*gtk.Image)
}

func (b *Builder) FetchEntryBuffer(obj string) *gtk.EntryBuffer {
	entryBufferObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return entryBufferObj.(*gtk.EntryBuffer)
}

func (b *Builder) FetchWindow(obj string) *gtk.Window {
	windowObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return windowObj.(*gtk.Window)
}

func (b *Builder) FetchTreeModelFilter(obj string) *gtk.TreeModelFilter {
	treeModelFilterObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return treeModelFilterObj.(*gtk.TreeModelFilter)
}

func (b *Builder) FetchTextBuffer(obj string) *gtk.TextBuffer {
	textBufferObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return textBufferObj.(*gtk.TextBuffer)
}

func (b *Builder) FetchScrolledWindow(obj string) *gtk.ScrolledWindow {
	scrolledWindowObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return scrolledWindowObj.(*gtk.ScrolledWindow)
}

func (b *Builder) FetchCellRendererToggle(obj string) *gtk.CellRendererToggle {
	cellToggleObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return cellToggleObj.(*gtk.CellRendererToggle)
}

func (b *Builder) FetchBox(obj string) *gtk.Box {
	boxObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return boxObj.(*gtk.Box)
}

func (b *Builder) FetchViewPort(obj string) *gtk.Viewport {
	viewPortObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return viewPortObj.(*gtk.Viewport)
}

func (b *Builder) FetchCellRendererText(obj string) *gtk.CellRendererText {
	cellTextObj, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return cellTextObj.(*gtk.CellRendererText)
}

func (b *Builder) FetchLinkButton(obj string) *gtk.LinkButton {
	linkButton, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return linkButton.(*gtk.LinkButton)
}

func (b *Builder) FetchIconView(obj string) *gtk.IconView {
	iconView, err := b.Builder.GetObject(obj)
	if err != nil {
		panic(err)
	}
	return iconView.(*gtk.IconView)
}
