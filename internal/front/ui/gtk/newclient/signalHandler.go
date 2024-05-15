package newclient

import (
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
)

func (nc *NewClientWindow) bindSignals() {
	nc.createButton.Connect("clicked", nc.createButtonPressed)
	nc.window.Connect("destroy", nc.window.Destroy)
}

func (nc *NewClientWindow) createButtonPressed() {
	clientView := &presenter.ClientDTO{
		Name:                  nc.Name(),
		Phones:                nc.Phone(),
		Email:                 nc.Email(),
		AdditionalInformation: nc.AdditionalInformation(),
	}
	nc.req.CreateClient(clientView)
}
