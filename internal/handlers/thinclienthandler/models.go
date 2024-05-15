package thinclienthandler

import (
	"io"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/handlers"
)

type SubHandler interface {
	Subscribe(token, adress, apiPath string) error
}
type DataBase interface {
	handlers.Server
}

type Responcer interface {
	App
	AdvertisementSender
	DatabaseManagerSender
	AdvertisementRemover
	SendErrors
	Progresser
	RequestComplete()
}

type App interface {
	SetConnectionStatus(bool)
	UnlockFilesWindow()
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
	SendFile(*datatransferobjects.FileDTO)
	SendFileFirstPlace(*datatransferobjects.FileDTO)
	ShowFile(*datatransferobjects.FileDTO)
}

type AdvertisementRemover interface {
	RemoveClientByName(string)
	RemoveOrderByID(int)
	RemoveLineAdvertisementByID(int)
	RemoveBlockAdvertisementByID(int)
	RemoveTagByName(string)
	RemoveExtraChargeByName(string)
	RemoveCostRate(string)
	RemoveFileByName(string)
}

type DatabaseManagerSender interface {
	SendNetworkDatabases([]string)
	SendLocalDatabases([]string)
	SendDefaultPort(uint)
	SendSuccesConnection()
}

type SendErrors interface {
	SendError(error)
	SendConnectionError(error)
}

type Progresser interface {
	ProgressComplete()
	CancelProgressWithError(error)
	SendMessage(string)
}

type FileStorage interface {
	OpenForRead(name string) (io.ReadSeekCloser, error)
}
