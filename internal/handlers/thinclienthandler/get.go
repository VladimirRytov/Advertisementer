package thinclienthandler

import (
	"context"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/logging"
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
	a.app.SendClient(&client)
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
	a.app.SendAdvertisementsOrder(&order)
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
	a.app.SendLineAdvertisement(&line)
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
	a.app.SendBlockAdvertisement(&block)
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
	a.app.SendTag(&tag)
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
	a.app.SendExtraCharge(&extraCharge)
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
	a.app.SendCostRate(&costRate)
}

func (a *AdvertisementController) FileByName(name string) {
	logging.Logger.Debug("advertisementhandler.FileByName: get request. get FileByName by name")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	file, err := a.req.FileByName(ctx, name)
	if err != nil {
		a.handleError(err)
		return
	}
	a.app.ShowFile(&file)
}

func (a *AdvertisementController) FileMiniatureByName(name string, size string) {
	logging.Logger.Debug("advertisementhandler.FileMiniatureByName: get request. get FileByName by name")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	file, err := a.req.FileMiniatureByName(ctx, name, size)
	if err != nil {
		a.handleError(err)
		return
	}
	a.app.ShowFile(&file)
}

func (a *AdvertisementController) NewFileMiniatureByName(name string, size string) {
	logging.Logger.Debug("advertisementhandler.FileMiniatureByName: get request. get FileByName by name")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	file, err := a.req.FileMiniatureByName(ctx, name, size)
	if err != nil {
		a.handleError(err)
		return
	}
	a.app.SendFileFirstPlace(&file)
}
