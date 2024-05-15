package errorwindow

import "github.com/gotk3/gotk3/gtk"

func (ew *ErrorWindow) ErrorMessage() string {
	text, _ := ew.errorMessage.GetText()
	return text
}

func (ew *ErrorWindow) SetErrorMessage(s string) {
	ew.errorMessage.SetText(s)
}

func (ew *ErrorWindow) Window() *gtk.Window {
	return ew.window
}

func (ew *ErrorWindow) Show() {
	ew.window.Show()
}

func (ew *ErrorWindow) Close() {
	ew.window.Destroy()
}
