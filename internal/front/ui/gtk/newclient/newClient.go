package newclient

import (
	"github.com/VladimirRytov/advertisementer/internal/logging"

	"github.com/gotk3/gotk3/gtk"
)

func (nc *NewClientWindow) Name() string {
	return nc.nameEntry.GetLayout().GetText()
}

func (nc *NewClientWindow) SetName(name string) {
	nc.nameEntry.SetText(name)
}

func (nc *NewClientWindow) Phone() string {
	return nc.phoneEntry.GetLayout().GetText()
}

func (nc *NewClientWindow) SetPhone(phone string) {
	nc.phoneEntry.SetText(phone)
}

func (nc *NewClientWindow) Email() string {
	return nc.emailEntry.GetLayout().GetText()
}

func (nc *NewClientWindow) SetEmail(mail string) {
	nc.emailEntry.SetText(mail)
}

func (nc *NewClientWindow) AdditionalInformation() string {
	text, err := nc.additionalInformationTextBuffer.GetText(
		nc.additionalInformationTextBuffer.GetStartIter(), nc.additionalInformationTextBuffer.GetEndIter(), true)
	if err != nil {
		logging.Logger.Error("NewClient AdditionalInfo: cannot get text from buffer", "error", err)
	}
	return text
}

func (nc *NewClientWindow) SetAdditionalInfo(text string) {
	nc.additionalInformationTextBuffer.SetText(text)
}

func (nc *NewClientWindow) Window() *gtk.Window {
	return nc.window
}

func (nc *NewClientWindow) Show() {
	nc.window.Show()
}

func (nc *NewClientWindow) Close() {
	nc.window.Close()
}
