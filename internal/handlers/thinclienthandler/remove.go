package thinclienthandler

import (
	"context"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func (a *AdvertisementController) RemoveClientByName(name string) {
	logging.Logger.Debug("advertisementhandler: remove request. Removing Client by name", "Client", name)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	err := a.req.RemoveClientByName(ctx, name)
	if err != nil {
		a.handleError(err)
	}
}

func (a *AdvertisementController) RemoveOrderByID(id int) {
	logging.Logger.Debug("advertisementhandler: remove request. Removing Order by id", "Order", id)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	err := a.req.RemoveOrderByID(ctx, id)
	if err != nil {
		a.handleError(err)
	}
}

func (a *AdvertisementController) RemoveLineAdvertisementByID(id int) {
	logging.Logger.Debug("advertisementhandler: remove request. Removing LineAdvertisement by id", "LineAdvertisement", id)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	err := a.req.RemoveLineAdvertisementByID(ctx, id)
	if err != nil {
		a.handleError(err)
	}
}

func (a *AdvertisementController) RemoveBlockAdvertisementByID(id int) {
	logging.Logger.Debug("advertisementhandler: remove request. Removing BlockAdvertisement by id", "BlockAdvertisement", id)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	err := a.req.RemoveBlockAdvertisementByID(ctx, id)
	if err != nil {
		a.handleError(err)
	}
}

func (a *AdvertisementController) RemoveTagByName(name string) {
	logging.Logger.Debug("advertisementhandler: remove request. Removing Tag by name", "Tag", name)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	err := a.req.RemoveTagByName(ctx, name)
	if err != nil {
		a.handleError(err)
	}
}

func (a *AdvertisementController) RemoveExtraChargeByName(name string) {
	logging.Logger.Debug("advertisementhandler: remove request. Removing ExtraCharge by name", "ExtraCharge", name)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	err := a.req.RemoveExtraChargeByName(ctx, name)
	if err != nil {
		a.handleError(err)
		return
	}
}

func (a *AdvertisementController) RemoveCostRateByName(name string) {
	logging.Logger.Debug("advertisementhandler: remove request. Removing CostRate by name", "CostRate", name)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	err := a.req.RemoveCostRateByName(ctx, name)
	if err != nil {
		a.handleError(err)
		return
	}
}

func (a *AdvertisementController) RemoveFileByName(name string) {
	logging.Logger.Debug("advertisementhandler.RemoveFileByName: remove request. Removing file by name", "File", name)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	err := a.req.RemoveFileByName(ctx, name)
	if err != nil {
		a.handleError(err)
		return
	}
	a.app.RemoveFileByName(name)
}
