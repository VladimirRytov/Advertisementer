package thinclienthandler

import (
	"context"
	"errors"

	"github.com/VladimirRytov/advertisementer/internal/handlers"
	"github.com/VladimirRytov/advertisementer/internal/logging"
)

const (
	CancelStr  = "операция была отменена пользователем"
	TimeoutStr = "превышено время ожидания"
)

type AdvertisementController struct {
	status bool
	req    DataBase
	app    Responcer
}

func NewAdvertisementController(app Responcer, files FileStorage) *AdvertisementController {
	return &AdvertisementController{app: app}
}

func (ac *AdvertisementController) InitAdvertisementController(dbGateway handlers.Server) {
	logging.Logger.Info("advertisementhandler: Initialize Advertisement controller")
	ac.req = dbGateway
	ac.status = true

}

func (ac *AdvertisementController) Close() error {
	logging.Logger.Info("advertisementhandler: Initialize Advertisement controller")
	return ac.req.Close()
}

func (ac *AdvertisementController) ConnectionInfo() map[string]string {
	return ac.req.ConnectionInfo()
}

func (ac *AdvertisementController) handleError(err error) {
	switch {
	case errors.Is(err, context.Canceled):
		ac.app.SendError(errors.New(CancelStr))
	case errors.Is(err, context.DeadlineExceeded):
		ac.app.SendError(errors.New(TimeoutStr))
	default:
		ac.app.SendError(err)
	}
}
