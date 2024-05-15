package windowscontroller

import (
	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/front/ui/presenter"
)

type GUI struct {
	converter ToViewConverter
	front     App
}

type App interface {
	LoadLoginWin()
	Start([]string)
	Stop()
	RequestCompleted()
	ProgressComplete()
	NewErrorWindow(error)
	LoginWin
	AdvWin
	ProgressCompleteWithError(error)
	AppendCostRate(*presenter.CostRateDTO)
	RemoveCostRate(string)
	RecieveValue(string)
	SetActiveCostRate(string)
	AppendMessage(string)
	SetConnectionStatus(bool)
	AppendSelectedFile(*presenter.File)
}

type LoginWin interface {
	LoadNetworkDatabases([]string)
	LoadLocalDatabases([]string)
	ShowDefaultPort(string)
	ShowError(error)
	UnlockFilesWindow()
	LockAll()
	UnlockAllLoginForm()
}

type AdvWin interface {
	InitAdvertisement()
	AppendTag(*presenter.TagDTO)
	RemoveTag(string)
	AppendExtraCharge(*presenter.ExtraChargeDTO)
	RemoveExtraCharge(string)
	AppendClient(*presenter.ClientDTO)
	RemoveClient(string)
	AppendOrder(*presenter.OrderDTO)
	RemoveOrder(int)
	AppendBlockAdvertisement(*presenter.BlockAdvertisementDTO)
	RemoveBlockAdvertisement(int)
	AppendLineAdvertisement(*presenter.LineAdvertisementDTO)
	RemoveLineAdvertisement(int)
	AppendFile(*presenter.File)
	AppendFileFirstPlace(*presenter.File)
	RemoveFileByName(string)
}

func NewWindowsController(app App, conv ToViewConverter) *GUI {
	return &GUI{front: app, converter: conv}
}

type ToViewConverter interface {
	ClientToViewDTO(*datatransferobjects.ClientDTO) presenter.ClientDTO
	OrderToViewDTO(*datatransferobjects.OrderDTO) presenter.OrderDTO
	LineAdvertisementToViewDTO(*datatransferobjects.LineAdvertisementDTO) presenter.LineAdvertisementDTO
	BlockAdvertisementToViewDTO(*datatransferobjects.BlockAdvertisementDTO) presenter.BlockAdvertisementDTO
	TagToViewDTO(*datatransferobjects.TagDTO) presenter.TagDTO
	ExtraChargeToViewDTO(*datatransferobjects.ExtraChargeDTO) presenter.ExtraChargeDTO
	CostRateToViewDTO(*datatransferobjects.CostRateDTO) presenter.CostRateDTO
	FileToViewDTO(*datatransferobjects.FileDTO) presenter.File
}
