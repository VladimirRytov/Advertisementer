package handlers

import (
	"context"
	"io"
	"time"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
)

type Responcer interface {
	SendConnectionError(error)
	SendNetworkDatabases([]string)
	SendLocalDatabases([]string)
	SendDefaultPort(uint)
	SendSuccesConnection()
	Application
}

type ThickAdvertisementHandler interface {
	InitAdvertisementController(DataBase)
	HandlerRequests
}

type ThinAdvertisementHandler interface {
	InitAdvertisementController(Server)
	HandlerRequests
	ServerRequests
}

type HandlerRequests interface {
	HandleCreator
	HandleGetter
	HandleSearcher
	HandleRemover
	HandleUpdater
	HandleCloser
	ConnectionInfo() map[string]string
}

type ServerRequests interface {
	AllFiles(context.Context)
	NewFile(context.Context, datatransferobjects.FileDTO)
	FileByName(string)
	FileMiniatureByName(name string, size string)
	NewFileMultipart(ctx context.Context, file string)
	RemoveFileByName(string)
}

type FileStorage interface {
	New(name string) (io.WriteCloser, error)
	OpenForWrite(name string) (io.WriteCloser, error)
	OpenForRead(name string) (io.ReadSeekCloser, error)
	CopyFile(source, destination string) error
	FileByName(context.Context, string) (datatransferobjects.FileDTO, error)
}

type AdvFilesGateway interface {
	FileByName(context.Context, string) (datatransferobjects.FileDTO, error)
}

type SubHandler interface {
	Subscribe(token, adress, apiPath string) error
	Reciever
}

type Reciever interface {
	LockRecieving(lock bool)
	IgnoreMessages(lock bool)
}

type HandleCreator interface {
	NewClient(datatransferobjects.ClientDTO)
	NewAdvertisementsOrder(datatransferobjects.OrderDTO)
	NewLineAdvertisement(datatransferobjects.LineAdvertisementDTO)
	NewBlockAdvertisement(datatransferobjects.BlockAdvertisementDTO)
	NewExtraCharge(datatransferobjects.ExtraChargeDTO)
	NewTag(datatransferobjects.TagDTO)
	NewCostRate(datatransferobjects.CostRateDTO)
}

type HandleGetter interface {
	ClientByName(string)
	OrderByID(int)
	LineAdvertisementByID(int)
	BlockAdvertisementByID(int)
	TagByName(string)
	ExtraChargeByName(string)
}

type HandleSearcher interface {
	OrdersByClientName(string)

	BlockAdvertisementsByOrderID(int)
	LineAdvertisementsByOrderID(int)

	BlockAdvertisementsBetweenReleaseDates(time.Time, time.Time)
	LineAdvertisementsBetweenReleaseDates(time.Time, time.Time)

	BlockAdvertisementsActualReleaseDate()
	LineAdvertisementsActualReleaseDate()

	BlockAdvertisementsFromReleaseDate(time.Time)
	LineAdvertisementsFromReleaseDate(time.Time)

	AllTags()
	AllExtraCharges()
	AllClients()
	AllOrders()
	AllLineAdvertisements()
	AllBlockAdvertisements()
	AllCostRates()
}

type HandleRemover interface {
	RemoveClientByName(string)
	RemoveOrderByID(int)
	RemoveLineAdvertisementByID(int)
	RemoveBlockAdvertisementByID(int)
	RemoveTagByName(string)
	RemoveExtraChargeByName(string)
	RemoveCostRateByName(string)
}

type HandleUpdater interface {
	UpdateClient(datatransferobjects.ClientDTO)
	UpdateOrder(datatransferobjects.OrderDTO)
	UpdateLineAdvertisement(datatransferobjects.LineAdvertisementDTO)
	UpdateBlockAdvertisement(datatransferobjects.BlockAdvertisementDTO)
	UpdateExtraCharge(datatransferobjects.ExtraChargeDTO)
	UpdateTag(datatransferobjects.TagDTO)
	UpdateCostRate(datatransferobjects.CostRateDTO)
}

type ReportGenerator interface {
	NewReport(ctx context.Context, blocks []datatransferobjects.BlockAdvertisementReport, lines []datatransferobjects.LineAdvertisementReport,
		data datatransferobjects.ReportParams) error
}

type ReportHandler interface {
	RegisterGenerator(string, ReportGenerator)
	InitReportGenerator(progr Progresser, reqGate ReportRequestGate)
	GenerateReport(ctx context.Context, params datatransferobjects.ReportParams)
}

type HandleCloser interface {
	Close() error
}

type Application interface {
	AdvertisementSender
	SendErrors
	Progresser
	RequestComplete()
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

type ReportRequestGate interface {
	OrderByID(context.Context, int) (datatransferobjects.OrderDTO, error)
	BlockAdvertisementBetweenReleaseDates(context.Context, time.Time, time.Time) ([]datatransferobjects.BlockAdvertisementDTO, error)
	LineAdvertisementBetweenReleaseDates(context.Context, time.Time, time.Time) ([]datatransferobjects.LineAdvertisementDTO, error)
}
