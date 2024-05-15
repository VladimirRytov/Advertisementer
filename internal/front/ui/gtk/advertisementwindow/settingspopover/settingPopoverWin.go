package advertisementwindow

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/gtk/builder"

	"github.com/gotk3/gotk3/gtk"
)

type SettingsPopover struct {
	popover               *gtk.Popover
	journalsButton        *gtk.Button
	paymentSettingsButton *gtk.Button
	settingsButton        *gtk.Button
	manualButton          *gtk.Button
	exitButton            *gtk.Button
}

func Build(bldFile *builder.Builder) *SettingsPopover {
	return &SettingsPopover{
		popover:               bldFile.FetchPopover("SettingsPopover"),
		journalsButton:        bldFile.FetchButton("AdvertisementsJournalSettingsButton"),
		paymentSettingsButton: bldFile.FetchButton("CalculatePaymentSettingsButton"),
		settingsButton:        bldFile.FetchButton("OpenSettingsButton"),
		manualButton:          bldFile.FetchButton("ManualButton"),
		exitButton:            bldFile.FetchButton("LogOutButton"),
	}
}
