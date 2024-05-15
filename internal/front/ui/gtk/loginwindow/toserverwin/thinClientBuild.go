package toserverwin

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"

	"github.com/gotk3/gotk3/gtk"
)

type ConnectToServerWindow struct {
	AuthorizationTypeLabel    *gtk.Label
	AuthorizationTypeComboBox *gtk.ComboBox

	URLLabel *gtk.Label
	URLEntry *gtk.Entry

	LoginLabel *gtk.Label
	LoginEntry *gtk.Entry

	PortLabel *gtk.Label
	PortEntry *gtk.Entry

	PasswordLabel  *gtk.Label
	PasswordEntry  *gtk.Entry
	PasswordBuffer *gtk.EntryBuffer

	EncryptionLabel       *gtk.Label
	EncryptionCheckButton *gtk.CheckButton
}

func BuildConnectToServerWindow(builder *builder.Builder) *ConnectToServerWindow {
	return &ConnectToServerWindow{
		AuthorizationTypeLabel:    builder.FetchLabel("ConnectToServerAuthorizationTypeLabel"),
		AuthorizationTypeComboBox: builder.FetchComboBox("ConnectToServerAuthorizationTypeComboBox"),
		URLLabel:                  builder.FetchLabel("ConnectToServerURLLabel"),
		URLEntry:                  builder.FetchEntry("ConnectToServerURLEntry"),
		LoginLabel:                builder.FetchLabel("ConnectToServerLoginLabel"),
		LoginEntry:                builder.FetchEntry("ConnectToServerLoginEntry"),
		PortLabel:                 builder.FetchLabel("ConnectToServerPortLabel"),
		PortEntry:                 builder.FetchEntry("ConnectToServerPortEntry"),
		PasswordLabel:             builder.FetchLabel("ConnectToServerPasswordLabel"),
		PasswordEntry:             builder.FetchEntry("ConnectToServerPasswordEntry"),
		PasswordBuffer:            builder.FetchEntryBuffer("ConnectToServerPasswordBuffer"),
		EncryptionLabel:           builder.FetchLabel("ConnectToServerEncryptLabel"),
		EncryptionCheckButton:     builder.FetchCheckButton("ConnectToServerEncryptCheckButton"),
	}
}
