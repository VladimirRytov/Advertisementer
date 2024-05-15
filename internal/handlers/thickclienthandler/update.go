package thickclienthandler

import (
	"context"
	"errors"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/logging"
	"github.com/VladimirRytov/advertisementer/internal/mapper"
)

func (a *AdvertisementController) UpdateClient(client datatransferobjects.ClientDTO) {
	logging.Logger.Debug("advertisementhandler: update request. Updating Client", "Client", client)
	clientAdv, err := mapper.DtoToClient(&client)
	if err != nil {
		logging.Logger.Warn("advertisementhandler: update request. Client entity is not created", "err", err)
		a.responcer.SendError(err)
		return
	}
	client = mapper.ClientToDTO(&clientAdv)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	err = a.req.UpdateClient(ctx, &client)
	if err != nil {
		logging.Logger.Error("advertisementhandler.UpdateClient: cannot update client", "error", err)
		switch {
		case errors.Is(err, context.Canceled):
			a.responcer.SendError(errors.New(CancelStr))
		case errors.Is(err, context.DeadlineExceeded):
			a.responcer.SendError(errors.New(TimeoutStr))
		default:
			a.responcer.SendError(err)
		}
		return

	}
	a.responcer.SendClient(&client)
}

func (a *AdvertisementController) UpdateOrder(order datatransferobjects.OrderDTO) {
	logging.Logger.Debug("advertisementhandler: update request. Updating Order", "Order", order)
	orderAdv, err := mapper.DtoToOrder(&order)
	if err != nil {
		logging.Logger.Warn("advertisementhandler: update request. Order entity is not created", "err", err)
		a.responcer.SendError(err)
		return
	}
	order = mapper.OrderToDTO(&orderAdv)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	err = a.req.UpdateOrder(ctx, &order)
	if err != nil {
		logging.Logger.Error("advertisementhandler.UpdateOrder: cannot update order", "error", err)
		a.handleError(err)
		return

	}
	a.responcer.SendAdvertisementsOrder(&order)
}

func (a *AdvertisementController) UpdateLineAdvertisement(line datatransferobjects.LineAdvertisementDTO) {
	logging.Logger.Debug("advertisementhandler: update request. Updating LineAdvertisement", "LineAdvertisement", line)
	lineAdv, err := mapper.DtoToAdvertisementLine(&line)
	if err != nil {
		logging.Logger.Warn("advertisementhandler: update request. LineAdvertisement entity is not created", "err", err)
		a.responcer.SendError(err)
		return
	}
	line = mapper.LineAdvertisementToDTO(&lineAdv)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	err = a.req.UpdateLineAdvertisement(ctx, &line)
	if err != nil {
		logging.Logger.Error("advertisementhandler.UpdateLineAdvertisement: cannot update LineAdvertisement", "error", err)
		a.handleError(err)
		return

	}

	a.responcer.SendLineAdvertisement(&line)
}

func (a *AdvertisementController) UpdateBlockAdvertisement(block datatransferobjects.BlockAdvertisementDTO) {
	logging.Logger.Debug("advertisementhandler: update request. Updating BlockAdvertisement", "BlockAdvertisement", block)
	blockAdv, err := mapper.DtoToAdvertisementBlock(&block)
	if err != nil {
		logging.Logger.Warn("advertisementhandler: update request. BlockAdvertisement entity is not created", "err", err)
		a.responcer.SendError(err)
		return
	}
	block = mapper.BlockAdvertisementToDTO(&blockAdv)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	err = a.req.UpdateBlockAdvertisement(ctx, &block)
	if err != nil {
		logging.Logger.Error("advertisementhandler.UpdateBlockAdvertisement: cannot update BlockAdvertisement", "error", err)
		a.handleError(err)
		return
	}
	a.responcer.SendBlockAdvertisement(&block)
}

func (a *AdvertisementController) UpdateExtraCharge(extraCharge datatransferobjects.ExtraChargeDTO) {
	logging.Logger.Debug("advertisementhandler: update request. Updating ExtraCharge", "ExtraCharge", extraCharge)
	chargeAdv, err := mapper.DtoToExtraCharge(&extraCharge)
	if err != nil {
		logging.Logger.Warn("advertisementhandler: update request. ExtraCharge entity is not created", "err", err)
		a.responcer.SendError(err)
		return
	}
	extraCharge = mapper.ExtraChargeToDTO(&chargeAdv)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	err = a.req.UpdateExtraCharge(ctx, &extraCharge)
	if err != nil {
		logging.Logger.Error("advertisementhandler.UpdateExtraCharge: cannot update ExtraCharge", "error", err)
		a.handleError(err)
		return
	}

	a.responcer.SendExtraCharge(&extraCharge)
}

func (a *AdvertisementController) UpdateTag(tag datatransferobjects.TagDTO) {
	logging.Logger.Debug("advertisementhandler: update request. Updating Tag", "Tag", tag)
	tagAdv, err := mapper.DtoToTag(&tag)
	if err != nil {
		logging.Logger.Warn("advertisementhandler: update request. Tag entity is not created", "err", err)
		a.responcer.SendError(err)
		return
	}
	tag = mapper.TagToDTO(&tagAdv)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	err = a.req.UpdateTag(ctx, &tag)
	if err != nil {
		logging.Logger.Error("advertisementhandler.UpdateTag: cannot update Tag", "error", err)
		a.handleError(err)
		return
	}

	a.responcer.SendTag(&tag)
}

func (a *AdvertisementController) UpdateCostRate(costRate datatransferobjects.CostRateDTO) {
	logging.Logger.Debug("advertisementhandler: update request. Updating CostRate", "name", costRate)
	costRateAdv, err := mapper.DtoToCostRate(&costRate)
	if err != nil {
		logging.Logger.Warn("advertisementhandler: update request. CostRate entity is not created", "err", err)
		a.responcer.SendError(err)
		return
	}
	costRate = mapper.CostRateToDTO(&costRateAdv)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	err = a.req.UpdateCostRate(ctx, &costRate)
	if err != nil {
		logging.Logger.Error("advertisementhandler.UpdateCostRate: cannot update CostRate", "error", err)
		a.handleError(err)
		return
	}

	a.responcer.SendCostRate(&costRate)
}
