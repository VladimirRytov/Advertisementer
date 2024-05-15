package thinclienthandler

import (
	"context"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func (a *AdvertisementController) UpdateClient(client datatransferobjects.ClientDTO) {
	logging.Logger.Debug("advertisementhandler: update request. Updating Client", "Client", client)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	err := a.req.UpdateClient(ctx, &client)
	if err != nil {
		logging.Logger.Error("advertisementhandler.UpdateClient: cannot update client", "error", err)
		a.handleError(err)
	}
}

func (a *AdvertisementController) UpdateOrder(order datatransferobjects.OrderDTO) {
	logging.Logger.Debug("advertisementhandler: update request. Updating Order", "Order", order)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	err := a.req.UpdateOrder(ctx, &order)
	if err != nil {
		logging.Logger.Error("advertisementhandler.UpdateOrder: cannot update order", "error", err)
		a.handleError(err)
	}
}

func (a *AdvertisementController) UpdateLineAdvertisement(line datatransferobjects.LineAdvertisementDTO) {
	logging.Logger.Debug("advertisementhandler: update request. Updating LineAdvertisement", "LineAdvertisement", line)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	err := a.req.UpdateLineAdvertisement(ctx, &line)
	if err != nil {
		logging.Logger.Error("advertisementhandler.UpdateLineAdvertisement: cannot update LineAdvertisement", "error", err)
		a.handleError(err)
	}
}

func (a *AdvertisementController) UpdateBlockAdvertisement(block datatransferobjects.BlockAdvertisementDTO) {
	logging.Logger.Debug("advertisementhandler: update request. Updating BlockAdvertisement", "BlockAdvertisement", block)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	err := a.req.UpdateBlockAdvertisement(ctx, &block)
	if err != nil {
		logging.Logger.Error("advertisementhandler.UpdateBlockAdvertisement: cannot update BlockAdvertisement", "error", err)
		a.handleError(err)
	}
}

func (a *AdvertisementController) UpdateExtraCharge(extraCharge datatransferobjects.ExtraChargeDTO) {
	logging.Logger.Debug("advertisementhandler: update request. Updating ExtraCharge", "ExtraCharge", extraCharge)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	err := a.req.UpdateExtraCharge(ctx, &extraCharge)
	if err != nil {
		logging.Logger.Error("advertisementhandler.UpdateExtraCharge: cannot update ExtraCharge", "error", err)
		a.handleError(err)
	}
}

func (a *AdvertisementController) UpdateTag(tag datatransferobjects.TagDTO) {
	logging.Logger.Debug("advertisementhandler: update request. Updating Tag", "Tag", tag)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	err := a.req.UpdateTag(ctx, &tag)
	if err != nil {
		logging.Logger.Error("advertisementhandler.UpdateTag: cannot update Tag", "error", err)
		a.handleError(err)
	}
}

func (a *AdvertisementController) UpdateCostRate(costRate datatransferobjects.CostRateDTO) {
	logging.Logger.Debug("advertisementhandler: update request. Updating CostRate", "name", costRate)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	err := a.req.UpdateCostRate(ctx, &costRate)
	if err != nil {
		logging.Logger.Error("advertisementhandler.UpdateCostRate: cannot update CostRate", "error", err)
		a.handleError(err)
	}
}
