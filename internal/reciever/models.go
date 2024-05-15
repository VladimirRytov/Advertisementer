package reciever

import (
	"context"
	"io"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
)

type SubscribeParams struct {
	UserID string `json:"userID,omitempty"`
	URL    string `json:"url"`
}

type Responcer interface {
	AdvertisementSender
	AdvertisementRemover
	SendErrors
	RequestComplete()
	SetConnectionStatus(bool)
}

type AdvertisementSender interface {
	SendClient(*datatransferobjects.ClientDTO)
	SendAdvertisementsOrder(*datatransferobjects.OrderDTO)
	SendLineAdvertisement(*datatransferobjects.LineAdvertisementDTO)
	SendBlockAdvertisement(*datatransferobjects.BlockAdvertisementDTO)
	SendExtraCharge(*datatransferobjects.ExtraChargeDTO)
	SendTag(*datatransferobjects.TagDTO)
	SendCostRate(*datatransferobjects.CostRateDTO)
	SendActiveCostRate(string)
	SendBlockAdvertisementCost(datatransferobjects.BlockAdvertisementDTO)
	SendLineAdvertisementCost(datatransferobjects.LineAdvertisementDTO)
	SendOrderCost(datatransferobjects.OrderDTO)
}

type AdvertisementRemover interface {
	RemoveClientByName(string)
	RemoveOrderByID(int)
	RemoveLineAdvertisementByID(int)
	RemoveBlockAdvertisementByID(int)
	RemoveTagByName(string)
	RemoveExtraChargeByName(string)
	RemoveCostRate(string)
}

type SendErrors interface {
	SendError(error)
	SendConnectionError(error)
}

type Reciever interface {
	Subscribe(token, adress string) error
	LockRecieve(lock bool)
	IgnoreMessages(lock bool)
	CloseConn() error
}

type Sender interface {
	Initialize(ctx context.Context, dsn datatransferobjects.ServerDSN) error

	GetRequest(ctx context.Context, params *datatransferobjects.RequestDTO) ([]byte, error)
	CreateRequest(ctx context.Context, params *datatransferobjects.RequestDTO, data io.Reader) ([]byte, error)
	DeleteRequest(ctx context.Context, params *datatransferobjects.RequestDTO) ([]byte, error)

	SetAuthToken(token string)
	SetAdress(addr string)
	SetApiPath(path string)
	ConnectionInfo() map[string]string
}
