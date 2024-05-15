package thickclienthandler

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
	status    bool
	responcer Responcer
	req       handlers.DataBase
}

func NewAdvertisementController(responcer Responcer) *AdvertisementController {
	logging.Logger.Info("advertisementhandler: Initialize Advertisement controller")
	return &AdvertisementController{responcer: responcer}
}

func (ac *AdvertisementController) InitAdvertisementController(dbGateway handlers.DataBase) {
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
		ac.responcer.SendError(errors.New(CancelStr))
	case errors.Is(err, context.DeadlineExceeded):
		ac.responcer.SendError(errors.New(TimeoutStr))
	default:
		ac.responcer.SendError(err)
	}
}
