package newextracharge

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
)

func (nexw *NewExtrachargeWindow) bindSignals() {
	nexw.createButton.Connect("clicked", nexw.createButtonPressed)
	nexw.multiplierEntry.Connect("insert-text", nexw.tools.CheckExtraChargeString)
	nexw.window.Connect("destroy", nexw.window.Destroy)
}

func (nexw *NewExtrachargeWindow) createButtonPressed() {
	extraCharge := &presenter.ExtraChargeDTO{
		ChargeName: nexw.ChargeName(),
		Multiplier: nexw.Multiplier(),
	}
	err := nexw.req.CreateExtraCharge(extraCharge)
	if err != nil {
		nexw.errorLabel.SetText(err.Error())
		nexw.errorPopover.SetRelativeTo(nexw.multiplierEntry)
		nexw.errorPopover.Popup()
	}
}
