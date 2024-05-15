package thickclienthandler

import (
	"context"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/logging"
	"github.com/VladimirRytov/advertisementer/internal/mapper"
)

func (a *AdvertisementController) ClientByName(name string) {
	logging.Logger.Debug("advertisementhandler: get request. Get Client by Name", "name", name)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	client, err := a.req.ClientByName(ctx, name)
	if err != nil {
		a.handleError(err)
		return
	}
	clientModel, err := mapper.DtoToClient(&client)
	if err != nil {
		a.responcer.SendError(err)
		return
	}
	newClient := mapper.ClientToDTO(&clientModel)
	a.responcer.SendClient(&newClient)
}

func (a *AdvertisementController) OrderByID(id int) {
	logging.Logger.Debug("advertisementhandler: get request. Get Order by id", "id", id)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	order, err := a.req.OrderByID(ctx, id)
	if err != nil {
		a.handleError(err)
		return
	}

	orderModel, err := mapper.DtoToOrder(&order)
	if err != nil {
		a.responcer.SendError(err)
		return
	}
	order = mapper.OrderToDTO(&orderModel)
	a.responcer.SendAdvertisementsOrder(&order)
}

func (a *AdvertisementController) LineAdvertisementByID(id int) {
	logging.Logger.Debug("advertisementhandler: get request. get LineAdvertisement by id", "LineAdvertisement", id)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	line, err := a.req.LineAdvertisementByID(ctx, id)
	if err != nil {
		a.handleError(err)
		return
	}

	lineModel, err := mapper.DtoToAdvertisementLine(&line)
	if err != nil {
		a.responcer.SendError(err)
		return
	}
	line = mapper.LineAdvertisementToDTO(&lineModel)
	a.responcer.SendLineAdvertisement(&line)
}

func (a *AdvertisementController) BlockAdvertisementByID(id int) {
	logging.Logger.Debug("advertisementhandler: get request. get BlockAdvertisement by id", "BlockAdvertisement", id)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	block, err := a.req.BlockAdvertisementByID(ctx, id)
	if err != nil {
		a.handleError(err)
		return
	}

	blockModel, err := mapper.DtoToAdvertisementBlock(&block)
	if err != nil {
		a.responcer.SendError(err)
		return
	}
	block = mapper.BlockAdvertisementToDTO(&blockModel)
	a.responcer.SendBlockAdvertisement(&block)
}

func (a *AdvertisementController) TagByName(name string) {
	logging.Logger.Debug("advertisementhandler: get request. get Tag by name", "Tag", name)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	tag, err := a.req.TagByName(ctx, name)
	if err != nil {
		a.handleError(err)
		return
	}

	tagModel, err := mapper.DtoToTag(&tag)
	if err != nil {
		a.responcer.SendError(err)
		return
	}
	tag = mapper.TagToDTO(&tagModel)
	a.responcer.SendTag(&tag)
}

func (a *AdvertisementController) ExtraChargeByName(name string) {
	logging.Logger.Debug("advertisementhandler: get request. get ExtraCharge by name", "ExtraCharge", name)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	extraCharge, err := a.req.ExtraChargeByName(ctx, name)
	if err != nil {
		a.handleError(err)
		return
	}

	extraChargeModel, err := mapper.DtoToExtraCharge(&extraCharge)
	if err != nil {
		a.responcer.SendError(err)
		return
	}
	extraCharge = mapper.ExtraChargeToDTO(&extraChargeModel)
	a.responcer.SendExtraCharge(&extraCharge)
}

func (a *AdvertisementController) CostRateByName(name string) {
	logging.Logger.Debug("advertisementhandler.CostRateByName: get request. get CostRate by name", "CostRate", name)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	costRate, err := a.req.CostRateByName(ctx, name)
	if err != nil {
		a.handleError(err)
		return
	}

	costRateModel, err := mapper.DtoToCostRate(&costRate)
	if err != nil {
		a.responcer.SendError(err)
		return
	}
	costRate = mapper.CostRateToDTO(&costRateModel)
	a.responcer.SendCostRate(&costRate)
}
