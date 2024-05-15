package windowscontroller

import (
	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func (g *GUI) SendDatabaseJournal(hournalsName []string) {}

func (g *GUI) SendExtraCosts() {}

func (g *GUI) SendClient(client *datatransferobjects.ClientDTO) {
	logging.Logger.Debug("responceController: Converting and sending Clients array to GUI")
	clientDto := g.converter.ClientToViewDTO(client)
	g.front.AppendClient(&clientDto)
}

func (g *GUI) SendAdvertisementsOrder(order *datatransferobjects.OrderDTO) {
	logging.Logger.Debug("responceController: Converting and sending Orders array to GUI")
	orderDto := g.converter.OrderToViewDTO(order)

	g.front.AppendOrder(&orderDto)
	for i := range order.LineAdvertisements {
		line := g.converter.LineAdvertisementToViewDTO(&order.LineAdvertisements[i])
		g.front.AppendLineAdvertisement(&line)
	}
	for i := range order.BlockAdvertisements {
		block := g.converter.BlockAdvertisementToViewDTO(&order.BlockAdvertisements[i])
		g.front.AppendBlockAdvertisement(&block)
	}
}

func (g *GUI) SendLineAdvertisement(line *datatransferobjects.LineAdvertisementDTO) {
	logging.Logger.Debug("responceController: Converting and sending LineAdvertisements array to GUI")
	lineAdv := g.converter.LineAdvertisementToViewDTO(line)
	g.front.AppendLineAdvertisement(&lineAdv)
}

func (g *GUI) SendBlockAdvertisement(block *datatransferobjects.BlockAdvertisementDTO) {
	logging.Logger.Debug("responceController: Converting and sending BlockAdvertisements array to GUI")
	blockAdv := g.converter.BlockAdvertisementToViewDTO(block)
	g.front.AppendBlockAdvertisement(&blockAdv)
}

func (g *GUI) SendTag(tag *datatransferobjects.TagDTO) {
	logging.Logger.Debug("responceController: Converting and sending Tags array to GUI")
	tagDto := g.converter.TagToViewDTO(tag)
	g.front.AppendTag(&tagDto)
}

func (g *GUI) SendExtraCharge(charge *datatransferobjects.ExtraChargeDTO) {
	logging.Logger.Debug("responceController: Converting and sending ExtraCharges array to GUI")
	extraChargeDto := g.converter.ExtraChargeToViewDTO(charge)
	g.front.AppendExtraCharge(&extraChargeDto)
}

func (g *GUI) RemoveClientByName(name string) {
	logging.Logger.Debug("responceController: sending Client name to GUI for delete")
	g.front.RemoveClient(name)
}

func (g *GUI) RemoveOrderByID(id int) {
	logging.Logger.Debug("responceController: sending Order id to GUI for delete")
	g.front.RemoveOrder(id)
}

func (g *GUI) RemoveLineAdvertisementByID(id int) {
	logging.Logger.Debug("responceController: sending LinesAdvertisements id to GUI for delete")
	g.front.RemoveLineAdvertisement(id)
}

func (g *GUI) RemoveBlockAdvertisementByID(id int) {
	logging.Logger.Debug("responceController: sending BlockAdvertisements id to GUI for delete")
	g.front.RemoveBlockAdvertisement(id)
}

func (g *GUI) RemoveTagByName(name string) {
	logging.Logger.Debug("responceController: sending Tags name to GUI for delete")
	g.front.RemoveTag(name)
}

func (g *GUI) RemoveExtraChargeByName(charge string) {
	logging.Logger.Debug("responceController: sending ExtraCharges name to GUI for delete")
	g.front.RemoveExtraCharge(charge)
}

func (g *GUI) SendError(err error) {
	logging.Logger.Debug("responceController: sending error to GUI", "error", err)
	g.front.NewErrorWindow(err)
}

func (g *GUI) RequestComplete() {
	g.front.RequestCompleted()
}

func (g *GUI) ProgressComplete() {
	g.front.ProgressComplete()
}

func (g *GUI) CancelProgressWithError(err error) {
	g.front.ProgressCompleteWithError(err)
}

func (g *GUI) SendMessage(msg string) {
	g.front.AppendMessage(msg)
}

func (g *GUI) SetConnectionStatus(status bool) {
	g.front.SetConnectionStatus(status)
}

func (g *GUI) RemoveFileByName(name string) {
	g.front.RemoveFileByName(name)
}

func (g *GUI) SendFile(file *datatransferobjects.FileDTO) {
	convertedFile := g.converter.FileToViewDTO(file)
	g.front.AppendFile(&convertedFile)
}

func (g *GUI) SendFileFirstPlace(file *datatransferobjects.FileDTO) {
	convertedFile := g.converter.FileToViewDTO(file)
	g.front.AppendFileFirstPlace(&convertedFile)
}

func (g *GUI) ShowFile(file *datatransferobjects.FileDTO) {
	convertedFile := g.converter.FileToViewDTO(file)
	g.front.AppendSelectedFile(&convertedFile)
}

func (g *GUI) UnlockFilesWindow() {
	g.front.UnlockFilesWindow()
}
