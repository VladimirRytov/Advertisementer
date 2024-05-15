package costcalculationhandler

import (
	"context"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/advertisements"
	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/handlers"
	"github.com/VladimirRytov/advertisementer/internal/logging"
	"github.com/VladimirRytov/advertisementer/internal/mapper"
)

type CostRateCalculator struct {
	reqGate        handlers.ThinCostRateRequests
	activeCostRate advertisements.CostRate
	app            Responcer
	enabled        bool
}

func NewCostRateCalculator() *CostRateCalculator {
	return &CostRateCalculator{}
}

func (c *CostRateCalculator) InitCostRateCalculator(app handlers.CostRateResponcer, reqGate handlers.ThinCostRateRequests) {
	c.app = app
	c.reqGate = reqGate
}

func (c *CostRateCalculator) SetActiveCostRate(costRateName string) error {
	context, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	costRateDto, err := c.reqGate.CostRateByName(context, costRateName)
	if err != nil {
		c.enabled = false
		logging.Logger.Debug("costRateCalculator.SetActiveCostRate: cannot set ActiveCostRate", "error", err)
		c.app.SendActiveCostRate("")
		return err
	}
	costRateAdv, err := mapper.DtoToCostRate(&costRateDto)
	if err != nil {
		c.enabled = false
		c.app.SendError(err)
		c.app.SendActiveCostRate("")
		return err
	}
	c.activeCostRate = costRateAdv
	logging.Logger.Debug("costRateCalculator.SetActiveCostRate: set active costRate", "costRate", c.activeCostRate)

	c.enabled = true
	c.app.SendActiveCostRate(c.activeCostRate.Name())
	return nil
}

func (c *CostRateCalculator) SelectedCostRate() {
	if c.enabled {
		c.SetActiveCostRate(c.activeCostRate.Name())
	}
}

func (c *CostRateCalculator) ActiveCostRate() {
	c.app.SendActiveCostRate(c.activeCostRate.Name())
}

func (c *CostRateCalculator) CalculateBlockAdvertisementCost(adv datatransferobjects.BlockAdvertisementDTO) {
	if c.enabled {
		logging.Logger.Debug("costRateCalculator.CalculateBlockAdvertisementCost: start calculating block Advertisement cost")
		blockAdv, err := mapper.DtoToAdvertisementBlock(&adv)
		if err != nil {
			c.app.SendError(err)
			return
		}

		tags, err := c.collectTags(adv.Tags)
		if err != nil {
			c.app.SendError(err)
			return
		}
		charges, err := c.collectExtraCharges(adv.ExtraCharges)
		if err != nil {
			c.app.SendError(err)
			return
		}
		cost, err := c.activeCostRate.CalculateBlockCost(blockAdv, tags, charges)
		if err != nil {
			c.app.SendError(err)
			return
		}
		err = blockAdv.SetCost(cost)
		if err != nil {
			c.app.SendError(err)
			return
		}
		c.app.SendBlockAdvertisementCost(mapper.BlockAdvertisementToDTO(&blockAdv))
	}
}

func (c *CostRateCalculator) CalculateLineAdvertisementCost(adv datatransferobjects.LineAdvertisementDTO) {
	if c.enabled {
		logging.Logger.Debug("costRateCalculator.CalculateLineAdvertisementCost: start calculating line Advertisement cost")

		lineAdv, err := mapper.DtoToAdvertisementLine(&adv)
		if err != nil {
			c.app.SendError(err)
			return
		}

		tags, err := c.collectTags(adv.Tags)
		if err != nil {
			c.app.SendError(err)
			return
		}
		charges, err := c.collectExtraCharges(adv.ExtraCharges)
		if err != nil {
			c.app.SendError(err)
			return
		}
		cost, err := c.activeCostRate.CalculateLineCost(lineAdv, tags, charges)
		if err != nil {
			c.app.SendError(err)
			return
		}
		err = lineAdv.SetCost(cost)
		if err != nil {
			c.app.SendError(err)
			return
		}
		c.app.SendLineAdvertisementCost(mapper.LineAdvertisementToDTO(&lineAdv))
	}
}

func (c *CostRateCalculator) CalculateOrderCost(adv datatransferobjects.OrderDTO) {
	if c.enabled {
		logging.Logger.Debug("costRateCalculator.CalculateBlockAdvertisementCost: start calculating order cost")

		orderAdv, err := mapper.DtoToOrder(&adv)
		if err != nil {
			c.app.SendError(err)
			return
		}
		if orderAdv.OrderId() != 0 {
			adv.BlockAdvertisements, err = c.collectBlockAdvertisements(orderAdv.OrderId())
			if err != nil {
				c.app.SendError(err)
				return
			}
			adv.LineAdvertisements, err = c.collectLineAdvertisements(orderAdv.OrderId())
			if err != nil {
				c.app.SendError(err)
				return
			}
		}
		blocksAdv, err := c.dtoToBlockAdvertisements(adv.BlockAdvertisements)
		if err != nil {
			c.app.SendError(err)
			return
		}
		linesAdv, err := c.dtoToLinekAdvertisements(adv.LineAdvertisements)
		if err != nil {
			c.app.SendError(err)
			return
		}

		cost := c.activeCostRate.CalculateOrderCost(orderAdv, blocksAdv, linesAdv)
		err = orderAdv.SetOrderCost(cost)
		if err != nil {
			c.app.SendError(err)
			return
		}
		c.app.SendOrderCost(mapper.OrderToDTO(&orderAdv))
	}

}

func (c *CostRateCalculator) collectTags(tags []string) ([]advertisements.Tag, error) {
	tagArr := make([]advertisements.Tag, 0, len(tags))
	for _, v := range tags {
		context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		tag, err := c.reqGate.TagByName(context, v)
		if err != nil {
			c.app.SendError(err)
			cancel()
			return tagArr, err
		}
		cancel()
		tagAdv, err := mapper.DtoToTag(&tag)
		if err != nil {
			c.app.SendError(err)
			cancel()
			return tagArr, err
		}
		tagArr = append(tagArr, tagAdv)
	}
	return tagArr, nil
}

func (c *CostRateCalculator) collectExtraCharges(extraCharges []string) ([]advertisements.ExtraCharge, error) {
	chargeArr := make([]advertisements.ExtraCharge, 0, len(extraCharges))
	for _, v := range extraCharges {
		context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		charge, err := c.reqGate.ExtraChargeByName(context, v)
		if err != nil {
			c.app.SendError(err)
			cancel()
			return chargeArr, err
		}
		cancel()

		chargeAdv, err := mapper.DtoToExtraCharge(&charge)
		if err != nil {
			c.app.SendError(err)
			return chargeArr, err
		}
		chargeArr = append(chargeArr, chargeAdv)
	}
	return chargeArr, nil
}

func (c *CostRateCalculator) collectBlockAdvertisements(orderID int) ([]datatransferobjects.BlockAdvertisementDTO, error) {
	context, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	blocksDto, err := c.reqGate.BlockAdvertisementsByOrderID(context, orderID)
	if err != nil {
		return nil, err
	}
	return blocksDto, nil
}

func (c *CostRateCalculator) dtoToBlockAdvertisements(blocksDto []datatransferobjects.BlockAdvertisementDTO) ([]advertisements.AdvertisementBlock, error) {
	blocksAdv := make([]advertisements.AdvertisementBlock, 0, len(blocksDto))
	for i := range blocksDto {
		blockAdv, err := mapper.DtoToAdvertisementBlock(&blocksDto[i])
		if err != nil {
			return nil, err
		}
		blocksAdv = append(blocksAdv, blockAdv)
	}
	return blocksAdv, nil
}

func (c *CostRateCalculator) collectLineAdvertisements(orderID int) ([]datatransferobjects.LineAdvertisementDTO, error) {
	context, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	lineDto, err := c.reqGate.LineAdvertisementsByOrderID(context, orderID)
	if err != nil {
		return nil, err
	}
	return lineDto, nil
}

func (c *CostRateCalculator) dtoToLinekAdvertisements(lineDto []datatransferobjects.LineAdvertisementDTO) ([]advertisements.AdvertisementLine, error) {
	linesAdv := make([]advertisements.AdvertisementLine, 0, len(lineDto))
	for i := range lineDto {
		lineAdv, err := mapper.DtoToAdvertisementLine(&lineDto[i])
		if err != nil {
			return nil, err
		}
		linesAdv = append(linesAdv, lineAdv)
	}
	return linesAdv, nil
}
