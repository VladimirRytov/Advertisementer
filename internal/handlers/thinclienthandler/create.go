package thinclienthandler

import (
	"context"
	"os"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

func (a *AdvertisementController) NewClient(client datatransferobjects.ClientDTO) {
	logging.Logger.Debug("advertisementhandler: create request. Checking and sending client DTO to database", "Client", client)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	_, err := a.req.NewClient(ctx, &client)
	if err != nil {
		logging.Logger.Error("advertisementhandler: create request. Client is not created")
		a.handleError(err)
	}
	a.app.RequestComplete()
}

func (a *AdvertisementController) NewAdvertisementsOrder(order datatransferobjects.OrderDTO) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	logging.Logger.Debug("advertisementhandler: create request. Checking and sending Order DTO to database", "Order", order)
	_, err := a.req.NewAdvertisementsOrder(ctx, &order)
	if err != nil {
		logging.Logger.Error("advertisementhandler: create request. Order is not created")
		a.handleError(err)
	}
	a.app.RequestComplete()
}

func (a *AdvertisementController) NewLineAdvertisement(line datatransferobjects.LineAdvertisementDTO) {
	logging.Logger.Debug("advertisementhandler: create request. Checking and sending LineAdvertisement DTO to database", "LineAdvertisement", line)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	_, err := a.req.NewLineAdvertisement(ctx, &line)
	if err != nil {
		logging.Logger.Error("advertisementhandler: create request. LineAdvertisement is not created")
		a.handleError(err)
	}
	a.app.RequestComplete()
}

func (a *AdvertisementController) NewBlockAdvertisement(block datatransferobjects.BlockAdvertisementDTO) {
	logging.Logger.Debug("advertisementhandler: create request. Checking and sending BlockAdvertisement DTO to database", "LineAdvertisement", block)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	_, err := a.req.NewBlockAdvertisement(ctx, &block)
	if err != nil {
		logging.Logger.Error("advertisementhandler: create request. BlockAdvertisement is not created")
		a.handleError(err)
	}
	a.app.RequestComplete()
}

func (a *AdvertisementController) NewExtraCharge(extraCharge datatransferobjects.ExtraChargeDTO) {
	logging.Logger.Debug("advertisementhandler: create request. Checking and sending Tag DTO to database", "ExtraCharge", extraCharge)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	_, err := a.req.NewExtraCharge(ctx, &extraCharge)
	if err != nil {
		logging.Logger.Error("advertisementhandler: create request. ExtraCharge is not created")
		a.handleError(err)
	}
	a.app.RequestComplete()
}

func (a *AdvertisementController) NewTag(tag datatransferobjects.TagDTO) {
	logging.Logger.Debug("advertisementhandler: create request. Checking and sending ExtraCharge DTO to database", "Tag", tag)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	_, err := a.req.NewTag(ctx, &tag)
	if err != nil {
		logging.Logger.Error("advertisementhandler: create request. Tag is not created")
		a.handleError(err)
	}
	a.app.RequestComplete()
}

func (a *AdvertisementController) NewCostRate(costRate datatransferobjects.CostRateDTO) {
	logging.Logger.Debug("advertisementhandler: create request. Checking and sending CostRate DTO to database", "CostRate", costRate)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	_, err := a.req.NewCostRate(ctx, &costRate)
	if err != nil {
		logging.Logger.Error("advertisementhandler.NewCostRate: create request. CostRate is not created", "error", err)
		a.handleError(err)
	}
	a.app.RequestComplete()
}

func (a *AdvertisementController) NewFile(ctx context.Context, file datatransferobjects.FileDTO) {
	logging.Logger.Debug("advertisementhandler.NewFile: create request")
	_, err := a.req.NewFile(ctx, &file)
	if err != nil {
		logging.Logger.Error("advertisementhandler.NewFile: create request. File is not created", "error", err)
		a.handleError(err)
		a.app.CancelProgressWithError(err)
		return
	}
	a.app.ProgressComplete()
}

func (a *AdvertisementController) NewFileMultipart(ctx context.Context, file string) {
	logging.Logger.Debug("advertisementhandler.NewFileMultipart: create request")
	f, err := os.OpenFile(file, os.O_RDONLY, 0644)
	if err != nil {
		a.app.CancelProgressWithError(err)
		return
	}

	fStat, err := f.Stat()
	if err != nil {
		a.app.CancelProgressWithError(err)
		return
	}
	fileDto := datatransferobjects.FileStream{
		Name: f.Name(),
		Size: fStat.Size(),
		Data: f,
	}
	fileName, err := a.req.NewFileUpload(ctx, fileDto)
	if err != nil {
		logging.Logger.Error("advertisementhandler.NewFile: create request. File is not created", "error", err)
		a.app.CancelProgressWithError(err)
		return
	}
	for i := range fileName {
		a.NewFileMiniatureByName(fileName[i], "miniature")
	}
	a.app.ProgressComplete()
}
