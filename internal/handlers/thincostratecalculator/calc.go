package thincostratecalculator

import (
	"context"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/handlers"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

type CostRateCalculator struct {
	reqGate        handlers.CostRequests
	activeCostRate string
	app            Responcer
	enabled        bool
}

func NewCostRateCalculator() *CostRateCalculator {
	return &CostRateCalculator{}
}

func (c *CostRateCalculator) InitCostRateCalculator(app handlers.Responcer, reqGate handlers.CostRequests) {
	c.app = app
	c.reqGate = reqGate
}

func (c *CostRateCalculator) SetActiveCostRate(costRateName string) error {
	c.activeCostRate = costRateName
	c.enabled = true
	return nil
}

func (c *CostRateCalculator) SelectedCostRate() {
	if c.enabled {
		c.SetActiveCostRate(c.activeCostRate)
	}
}

func (c *CostRateCalculator) ActiveCostRate() {
	c.app.SendActiveCostRate(c.activeCostRate)
}

func (c *CostRateCalculator) CalculateBlockAdvertisementCost(adv datatransferobjects.BlockAdvertisementDTO) {
	if c.enabled {
		logging.Logger.Debug("costRateCalculator.CalculateBlockAdvertisementCost: start calculating block Advertisement cost")
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
		defer cancel()
		blockAdv, err := c.reqGate.CalculateBlockAdvertisementCost(ctx, adv, c.activeCostRate)
		if err != nil {
			c.app.SendError(err)
			return
		}
		c.app.SendBlockAdvertisementCost(blockAdv)
	}
}

func (c *CostRateCalculator) CalculateLineAdvertisementCost(adv datatransferobjects.LineAdvertisementDTO) {
	if c.enabled {
		logging.Logger.Debug("costRateCalculator.CalculateLineAdvertisementCost: start calculating line Advertisement cost")
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
		defer cancel()
		lineAdv, err := c.reqGate.CalculateLineAdvertisementCost(ctx, adv, c.activeCostRate)
		if err != nil {
			c.app.SendError(err)
			return
		}
		c.app.SendLineAdvertisementCost(lineAdv)
	}
}

func (c *CostRateCalculator) CalculateOrderCost(adv datatransferobjects.OrderDTO) {
	if c.enabled {
		logging.Logger.Debug("costRateCalculator.CalculateBlockAdvertisementCost: start calculating order cost")
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
		defer cancel()

		order, err := c.reqGate.CalculateOrderCost(ctx, adv, c.activeCostRate)
		if err != nil {
			c.app.SendError(err)
		}
		c.app.SendOrderCost(order)
	}
}
