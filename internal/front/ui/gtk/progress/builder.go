package progress

import (
	"context"

	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"

	"github.com/gotk3/gotk3/gtk"
)

type MessageHandler interface {
	MessageList() *gtk.ListStore
	ClearMessageList()
}
type Tools interface {
	StringFromIter(*gtk.TreeIter, *gtk.TreeModel, int) (string, error)
}

type ProgressWindow struct {
	messages        MessageHandler
	tools           Tools
	window          *gtk.Window
	progresLabel    *gtk.Label
	currentProgress *gtk.Label
	logList         *gtk.ListStore
	scroller        *gtk.ScrolledWindow
	logTreeView     *gtk.TreeView
	scrollWindow    *gtk.ScrolledWindow
	closeButton     *gtk.Button
	signals         SignalHandler

	beforeFunc     func()
	afterFunc      func()
	cancelProgress context.CancelFunc
}

func Create(msgHandler MessageHandler, tools Tools) *ProgressWindow {
	build, err := builder.NewBuilderFromString(builder.Progress)
	if err != nil {
		panic(err)
	}
	pw := new(ProgressWindow)
	pw.messages = msgHandler
	pw.tools = tools
	pw.Build(build)
	pw.logList = pw.messages.MessageList()
	pw.logTreeView.SetModel(pw.logList)
	pw.bindSignals()
	pw.disableDeletable()
	pw.window.SetModal(true)
	pw.window.SetTitle("Выполнение операции")
	return pw
}

func (pw *ProgressWindow) Build(b *builder.Builder) {
	pw.window = b.FetchWindow("ProgressWindow")
	pw.progresLabel = b.FetchLabel("ProgressLabel")
	pw.currentProgress = b.FetchLabel("CurrentProgress")
	pw.scroller = b.FetchScrolledWindow("ScrollWindow")
	pw.logTreeView = b.FetchTreeView("logTreeView")
	pw.scrollWindow = b.FetchScrolledWindow("ScrollWindow")
	pw.closeButton = b.FetchButton("CloseButton")
}
