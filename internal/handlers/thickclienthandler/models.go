package thickclienthandler

import "github.com/VladimirRytov/advertisementer/internal/datatransferobjects"

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

type FileHandler interface {
	SendFile(*datatransferobjects.FileDTO)
	RemoveFileByName(string)
}
