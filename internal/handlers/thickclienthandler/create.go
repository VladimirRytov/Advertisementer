package thickclienthandler

import (
	"context"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/logging"
	"github.com/VladimirRytov/advertisementer/internal/mapper"
)

func (a *AdvertisementController) NewClient(client datatransferobjects.ClientDTO) {
	logging.Logger.Debug("advertisementhandler: create request. Checking and sending client DTO to database", "Client", client)
	clientAdv, err := mapper.DtoToClient(&client)
	if err != nil {
		logging.Logger.Warn("advertisementhandler: create request. Client is not created", "err", err)
		a.responcer.SendError(err)
		return
	}
	client = mapper.ClientToDTO(&clientAdv)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	_, err = a.req.NewClient(ctx, &client)
	if err != nil {
		logging.Logger.Error("advertisementhandler: create request. Client is not created")
		a.handleError(err)
		return
	}

	a.responcer.SendClient(&client)
	a.responcer.RequestComplete()
}

func (a *AdvertisementController) NewAdvertisementsOrder(order datatransferobjects.OrderDTO) {
	var (
		bloks []datatransferobjects.BlockAdvertisementDTO = make([]datatransferobjects.BlockAdvertisementDTO, 0, len(order.BlockAdvertisements))
		lines []datatransferobjects.LineAdvertisementDTO  = make([]datatransferobjects.LineAdvertisementDTO, 0, len(order.LineAdvertisements))
	)
	logging.Logger.Debug("advertisementhandler: create request. Checking and sending Order DTO to database", "Order", order)
	orderAdv, err := mapper.DtoToOrder(&order)
	if err != nil {
		logging.Logger.Warn("advertisementhandler: create request. Order is not created", "err", err)
		a.responcer.SendError(err)
		return
	}

	for i := range order.LineAdvertisements {
		line, err := mapper.DtoToAdvertisementLine(&order.LineAdvertisements[i])
		if err != nil {
			logging.Logger.Warn("advertisementhandler.NewAdvertisementsOrder: create request. LineAdvertisement is not created", "err", err)
			a.responcer.SendError(err)
			return
		}
		lines = append(lines, mapper.LineAdvertisementToDTO(&line))
	}
	for i := range order.BlockAdvertisements {
		block, err := mapper.DtoToAdvertisementBlock(&order.BlockAdvertisements[i])
		if err != nil {
			logging.Logger.Warn("advertisementhandler.NewAdvertisementsOrder: create request. Block is not created", "err", err)
			a.responcer.SendError(err)
			return
		}
		bloks = append(bloks, mapper.BlockAdvertisementToDTO(&block))
	}
	order = mapper.OrderToDTO(&orderAdv)
	order.BlockAdvertisements = bloks
	order.LineAdvertisements = lines

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	logging.Logger.Debug("advertisementhandler: create request. Checking and sending Order DTO to database", "Order", order)
	orders, err := a.req.NewAdvertisementsOrder(ctx, &order)
	if err != nil {
		logging.Logger.Error("advertisementhandler: create request. Order is not created")
		a.handleError(err)
		return
	}

	a.responcer.SendAdvertisementsOrder(&orders)
	a.responcer.RequestComplete()
}

func (a *AdvertisementController) NewLineAdvertisement(line datatransferobjects.LineAdvertisementDTO) {
	logging.Logger.Debug("advertisementhandler: create request. Checking and sending LineAdvertisement DTO to database", "LineAdvertisement", line)
	lineAdv, err := mapper.DtoToAdvertisementLine(&line)
	if err != nil {
		logging.Logger.Warn("advertisementhandler: create request. LineAdvertisement is not created", "err", err)
		a.responcer.SendError(err)
		return
	}
	line = mapper.LineAdvertisementToDTO(&lineAdv)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	id, err := a.req.NewLineAdvertisement(ctx, &line)
	if err != nil {
		logging.Logger.Error("advertisementhandler: create request. LineAdvertisement is not created")
		a.handleError(err)
		return
	}

	lineAdv.SetId(id)
	line = mapper.LineAdvertisementToDTO(&lineAdv)
	a.responcer.SendLineAdvertisement(&line)
	a.responcer.RequestComplete()
}

func (a *AdvertisementController) NewBlockAdvertisement(block datatransferobjects.BlockAdvertisementDTO) {
	logging.Logger.Debug("advertisementhandler: create request. Checking and sending BlockAdvertisement DTO to database", "LineAdvertisement", block)
	blockAdv, err := mapper.DtoToAdvertisementBlock(&block)
	logging.Logger.Warn("advertisementhandler: create request. BlockAdvertisement is not created", "err", err)
	if err != nil {
		a.responcer.SendError(err)
		return
	}
	block = mapper.BlockAdvertisementToDTO(&blockAdv)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	id, err := a.req.NewBlockAdvertisement(ctx, &block)
	if err != nil {
		logging.Logger.Error("advertisementhandler: create request. BlockAdvertisement is not created")
		a.handleError(err)
		return
	}

	err = blockAdv.SetId(id)
	if err != nil {
		a.responcer.SendError(err)
		return
	}

	block = mapper.BlockAdvertisementToDTO(&blockAdv)
	a.responcer.SendBlockAdvertisement(&block)
	a.responcer.RequestComplete()
}

func (a *AdvertisementController) NewExtraCharge(extraCharge datatransferobjects.ExtraChargeDTO) {
	logging.Logger.Debug("advertisementhandler: create request. Checking and sending Tag DTO to database", "ExtraCharge", extraCharge)

	chargeAdv, err := mapper.DtoToExtraCharge(&extraCharge)
	if err != nil {
		logging.Logger.Error("newExtraCharge: create request. ExtraCharge is not created", "err", err)
		a.responcer.SendError(err)
		return
	}
	extraCharge = mapper.ExtraChargeToDTO(&chargeAdv)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	_, err = a.req.NewExtraCharge(ctx, &extraCharge)
	if err != nil {
		logging.Logger.Error("advertisementhandler: create request. ExtraCharge is not created")
		a.handleError(err)
		return
	}

	a.responcer.SendExtraCharge(&extraCharge)
	a.responcer.RequestComplete()
}

func (a *AdvertisementController) NewTag(tag datatransferobjects.TagDTO) {
	logging.Logger.Debug("advertisementhandler: create request. Checking and sending ExtraCharge DTO to database", "Tag", tag)
	tagAdv, err := mapper.DtoToTag(&tag)
	if err != nil {
		logging.Logger.Warn("advertisementhandler: create request. ExtraCharge is not created", "err", err)
		a.responcer.SendError(err)
		return
	}
	tag = mapper.TagToDTO(&tagAdv)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	_, err = a.req.NewTag(ctx, &tag)
	if err != nil {
		logging.Logger.Error("advertisementhandler: create request. Tag is not created")
		a.handleError(err)
		return
	}

	a.responcer.SendTag(&tag)
	a.responcer.RequestComplete()
}

func (a *AdvertisementController) NewCostRate(costRate datatransferobjects.CostRateDTO) {
	logging.Logger.Debug("advertisementhandler: create request. Checking and sending CostRate DTO to database", "CostRate", costRate)
	costRateAdv, err := mapper.DtoToCostRate(&costRate)
	if err != nil {
		logging.Logger.Warn("advertisementhandler.NewCostRate: create request. CostRate is not created", "err", err)
		a.responcer.SendError(err)
		return
	}
	costRate = mapper.CostRateToDTO(&costRateAdv)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	_, err = a.req.NewCostRate(ctx, &costRate)
	if err != nil {
		logging.Logger.Error("advertisementhandler.NewCostRate: create request. CostRate is not created", "error", err)
		a.handleError(err)
		return
	}

	a.responcer.SendCostRate(&costRate)
	a.responcer.RequestComplete()
}
