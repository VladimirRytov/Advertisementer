package windowscontroller

import (
	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func (g *GUI) SendCostRate(costRate *datatransferobjects.CostRateDTO) {
	cost := g.converter.CostRateToViewDTO(costRate)
	g.front.AppendCostRate(&cost)
}

func (g *GUI) RemoveCostRate(costRateName string) {
	g.front.RemoveCostRate(costRateName)
}

func (g *GUI) SendBlockAdvertisementCost(blkAdv datatransferobjects.BlockAdvertisementDTO) {
	blockView := g.converter.BlockAdvertisementToViewDTO(&blkAdv)
	logging.Logger.Debug("GUI.SendBlockAdvertisementCost: start Recieving value", "cost ", blockView.Cost)
	g.front.RecieveValue(blockView.Cost)
}

func (g *GUI) SendLineAdvertisementCost(lineAdv datatransferobjects.LineAdvertisementDTO) {
	lineView := g.converter.LineAdvertisementToViewDTO(&lineAdv)
	logging.Logger.Debug("GUI.SendLineAdvertisementCost: start Recieving value", "cost ", lineAdv.Cost)
	g.front.RecieveValue(lineView.Cost)

}

func (g *GUI) SendOrderCost(order datatransferobjects.OrderDTO) {
	orderView := g.converter.OrderToViewDTO(&order)
	logging.Logger.Debug("GUI.SendOrderCost: start Recieving value", "cost ", orderView.Cost)
	g.front.RecieveValue(orderView.Cost)
}

func (g *GUI) SendActiveCostRate(name string) {
	logging.Logger.Debug("GUI.SendActiveCostRate: sending costRateName", "name", name)
	g.front.SetActiveCostRate(name)
}
